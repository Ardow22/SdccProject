package service

import (
	"SdccProject/matchingService/pattern"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	amqp "github.com/rabbitmq/amqp091-go"
	"os"
	"sort"

	"log"
	"time"
)

// struct per dati anagrafici e accademici di un utente
type DatiUtente struct {
	Nome         string
	Cognome      string
	Eta          int
	Email        string
	VotoLaurea   int
	Universita   string
	TipoLaurea   string
	Password     string
	Assegnamento string
}

type DatiUtente2 struct {
	Email    string
	Password string
}

type ScelteUtente struct {
	Email        string
	Scelta1      string
	Scelta2      string
	Scelta3      string
	Voto         int
	TipoDiLaurea string
}

type DatiAzienda struct {
	Nome       string
	Preferenza string
}

type MatchingResult string

type CompanyResult []string

type MatchingService int

type Result int

// struct che rappresenta un vertice nel grafo per l'algoritmo TTC
type Vertex struct {
	graph         *Graph
	vertexID      string
	outgoingEdges map[[2]string]bool
	incomingEdges map[[2]string]bool
}

// -----------funzioni per la connessione al database-------------------
// recupero della stringa di connessione per il database
func getDatabaseDSN(database string) string {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)
}

// apertura della connessione verso l'apposito database
func OpenDB(database string) (*sql.DB, error) {
	dsn := getDatabaseDSN(database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("errore nella connessione al DB: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("errore nel ping del DB: %w", err)
	}

	return db, nil
}

// -------------funzioni di supporto all'esecuzione dell'algoritmo TTC-------------------
// funzione ausiliaria per creare un nuovo vertice del grafo
func newVertex(graph *Graph, vertexID string) *Vertex {
	return &Vertex{
		graph:         graph,
		vertexID:      vertexID,
		outgoingEdges: make(map[[2]string]bool),
		incomingEdges: make(map[[2]string]bool),
	}
}

// funzione ausiliaria per restituire il grado di uscita del vertice.
func (v *Vertex) outdegree() int {
	return len(v.outgoingEdges)
}

// funzione ausiliaria per restituire un vertice adiacente al vertice corrente.
func (v *Vertex) anyNext() *Vertex {
	for edge := range v.outgoingEdges {
		return v.graph.vertices[edge[1]]
	}
	return nil
}

// rappresentazione dwll'intero grafo orientato
type Graph struct {
	vertices map[string]*Vertex
	edges    map[[2]string]bool
}

// funzione ausiliaria per la creazione di un nuovo grafo.
func newGraph(vertexIDs []string) *Graph {
	g := &Graph{
		vertices: make(map[string]*Vertex),
		edges:    make(map[[2]string]bool),
	}
	for _, id := range vertexIDs {
		g.vertices[id] = newVertex(g, id)
	}
	return g
}

// funzione ausiliaria per aggiungere un arco al grafo.
func (g *Graph) addEdge(source, target string) {
	g.edges[[2]string{source, target}] = true
	g.vertices[source].outgoingEdges[[2]string{source, target}] = true
	g.vertices[target].incomingEdges[[2]string{source, target}] = true
}

// funzione ausiliaria per eliminare un vertice dal grafo.
func (g *Graph) deleteVertex(vertexID string) {
	vertex := g.vertices[vertexID]
	involvedEdges := make(map[[2]string]bool)
	for edge := range vertex.outgoingEdges {
		involvedEdges[edge] = true
	}
	for edge := range vertex.incomingEdges {
		involvedEdges[edge] = true
	}

	for edge := range involvedEdges {
		delete(g.vertices[edge[1]].incomingEdges, edge)
		delete(g.vertices[edge[0]].outgoingEdges, edge)
		delete(g.edges, edge)
	}
	delete(g.vertices, vertexID)
}

// funzione ausiliaria per restituire un vertice arbitrario del grafo.
func (g *Graph) anyVertex() *Vertex {
	for _, v := range g.vertices {
		return v
	}
	return nil
}

// funzione ausiliaria per restituire l'insieme degli agenti in un ciclo.
func getAgents(g *Graph, cycle *Vertex, agents map[string]bool) map[string]bool {
	if _, ok := agents[cycle.vertexID]; ok {
		cycle = cycle.anyNext()
	}
	startingHouse := cycle
	currentVertex := startingHouse.anyNext()
	theAgents := make(map[string]bool)

	for !theAgents[currentVertex.vertexID] {
		theAgents[currentVertex.vertexID] = true
		currentVertex = currentVertex.anyNext()
		currentVertex = currentVertex.anyNext()
	}
	return theAgents
}

// funzione ausiliaria per trovare un vertice in un ciclo.
func anyCycle(g *Graph) *Vertex {
	visited := make(map[string]bool)
	v := g.anyVertex()

	if v == nil {
		//se il grafo è vuoto, ritorna nil
		return nil
	}

	//se il vertice è stato già visitato, significa che non c'è un ciclo
	for !visited[v.vertexID] {
		visited[v.vertexID] = true
		v = v.anyNext()

		if v == nil {
			//se nessun altro vertice è disponibile, ritorna nil
			return nil
		}
	}

	return v
}

// implementazione effettiva dell'algoritmo TTC
func topTradingCycles(agents, houses []string, agentPreferences map[string][]string, initialOwnership map[string]string) map[string]string {
	fmt.Println("\n\n")
	agentSet := make(map[string]bool)
	vertexSet := make([]string, 0)

	for _, a := range agents {
		agentSet[a] = true
		vertexSet = append(vertexSet, a)
	}
	for _, h := range houses {
		vertexSet = append(vertexSet, h)
	}
	fmt.Println("\nL'elenco dei vertici è: ", vertexSet)
	//creazione del grafo iniziale con tutti i vertici
	g := newGraph(vertexSet)
	fmt.Println("\nIl grafo iniziale: ")
	//stampo il grafo indicando per ogni vertice gli archi entranti ed uscenti
	g.printGraph()
	currentPreferenceIndex := make(map[string]int)

	preferredHouse := func(a string) (string, bool) {
		prefs := agentPreferences[a]
		if currentPreferenceIndex[a] >= len(prefs) {
			return "", false // non ci sono più preferenze disponibili
		}
		return prefs[currentPreferenceIndex[a]], true
	}

	//aggiunta degli archi con le preferenze degli agenti
	for _, a := range agents {
		preference, valid := preferredHouse(a)
		if valid {
			g.addEdge(a, preference)
		}
		currentPreferenceIndex[a] = 0
	}
	for _, h := range houses {
		g.addEdge(h, initialOwnership[h])
	}
	fmt.Println("\nAggiornamento grafo: ")
	g.printGraph()
	allocation := make(map[string]string)

	for len(g.vertices) > 0 {
		cycle := anyCycle(g)
		if cycle == nil {
			//se non ci sono più cicli, l'algoritmo si interrompe
			break
		}
		cycleAgents := getAgents(g, cycle, agentSet)
		fmt.Println("\nAgenti nel ciclo:")
		for agent := range cycleAgents {
			fmt.Println("- ", agent, "\n")
		}

		//eliminazione di una coppia dal ciclo, la quale diventa un'allocazione
		for a := range cycleAgents {
			h := g.vertices[a].anyNext().vertexID
			allocation[a] = h
			g.deleteVertex(a)
			g.deleteVertex(h)
			fmt.Println("\nEliminazione della coppia: ", a, ",", h)
			fmt.Println("\nAggiornamento grafo:")
			g.printGraph()
		}

		for _, a := range agents {
			if v, ok := g.vertices[a]; ok && v.outdegree() == 0 {
				for {
					pref, valid := preferredHouse(a)
					if !valid {
						break
					}
					if _, ok := g.vertices[pref]; ok {
						g.addEdge(a, pref)
						break
					}
					currentPreferenceIndex[a]++
				}
			}
		}
	}
	//Fallback: assegnazione manuale della coppia rimasta nel grafo
	for _, a := range agents {
		if _, ok := allocation[a]; !ok {
			if _, exists := g.vertices[a]; exists {
				for _, h := range houses {
					if _, exists := g.vertices[h]; exists {
						allocation[a] = h
						fmt.Printf("\n[Fallback] Assegnazione residua: %s -> %s\n", a, h)
						g.deleteVertex(a)
						g.deleteVertex(h)
						break
					}
				}
			}
		}
	}

	//stampa il risultato finale
	fmt.Println("\nAllocazione finale:")
	for agente, casa := range allocation {
		fmt.Printf("%s -> %s\n", agente, casa)
	}
	return allocation
}

// funzione ausiliaria per stampare il grafo.
func (g *Graph) printGraph() {
	for _, v := range g.vertices {
		fmt.Printf("Vertice: %s\n", v.vertexID)
		fmt.Printf("  Archi uscenti:\n")
		for edge := range v.outgoingEdges {
			fmt.Printf("    %s -> %s\n", edge[0], edge[1])
		}
		fmt.Printf("  Archi entranti:\n")
		for edge := range v.incomingEdges {
			fmt.Printf("    %s -> %s\n", edge[0], edge[1])
		}
	}
}

// -------------------------------------------------------------------------------------------
// funzione ausiliaria per calcolare le preferenze di ciascun gruppo di utenti
func calcolaPreferenzeGruppo(db *sql.DB, query string, args ...interface{}) (string, string, string) {
	rows, err := db.Query(query, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	preferenzaCounts := make(map[string]int)

	for rows.Next() {
		var pref1, pref2, pref3 string
		if err := rows.Scan(&pref1, &pref2, &pref3); err != nil {
			log.Fatal(err)
		}
		preferenzaCounts[pref1]++
		preferenzaCounts[pref2]++
		preferenzaCounts[pref3]++
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	//ordinamento delle preferenze in base alla frequenza
	type kv struct {
		Key   string
		Value int
	}
	var ss []kv
	for k, v := range preferenzaCounts {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	//selezione delle prime tre preferenze uniche
	preferenzeUniche := make([]string, 0)
	for _, kv := range ss {
		if len(preferenzeUniche) < 3 {
			unico := true
			for _, pref := range preferenzeUniche {
				if pref == kv.Key {
					unico = false
					break
				}
			}
			if unico {
				preferenzeUniche = append(preferenzeUniche, kv.Key)
			}
		} else {
			break
		}
	}

	pref1Gruppo := ""
	pref2Gruppo := ""
	pref3Gruppo := ""
	if len(preferenzeUniche) > 0 {
		pref1Gruppo = preferenzeUniche[0]
	}
	if len(preferenzeUniche) > 1 {
		pref2Gruppo = preferenzeUniche[1]
	}
	if len(preferenzeUniche) > 2 {
		pref3Gruppo = preferenzeUniche[2]
	}

	return pref1Gruppo, pref2Gruppo, pref3Gruppo
}

// funzione principale per eseguire il matching
func (t *MatchingService) Match(scelteUtente ScelteUtente, mResult *MatchingResult) error {
	fmt.Println("ScelteUtente:", scelteUtente)
	db, err := OpenDB("matchingservice")
	if err != nil {
		*mResult = "Error open database"
		panic(err.Error())
	}
	defer db.Close()
	_, err2 := db.Exec(`
        UPDATE datiutente
        SET preferenza1 = ?, preferenza2 = ?, preferenza3 = ?
        WHERE emailUtente = ?;`,
		scelteUtente.Scelta1,
		scelteUtente.Scelta2,
		scelteUtente.Scelta3,
		scelteUtente.Email,
	)
	if err2 != nil {
		*mResult = "Error update database"
		log.Println("Errore:", err2)
	}
	fmt.Println("AVVIO ALGORITMO TTC")
	//vengono raccolti dal database i 3 gruppi di utenti: triennale, magistrale <100, magistrale >100
	//verranno calcolate per ciascuno di questi gruppi le preferenze di gruppo
	pref1Triennale, pref2Triennale, pref3Triennale := calcolaPreferenzeGruppo(db, `
                SELECT preferenza1, preferenza2, preferenza3
                FROM datiutente
                WHERE tipoDiLaureaUtente = ? AND
                     preferenza1 IS NOT NULL AND
                     preferenza2 IS NOT NULL AND
                     preferenza3 IS NOT NULL;`,
		"triennale",
	)
	fmt.Println("Preferenze triennale:", pref1Triennale, pref2Triennale, pref3Triennale)

	//calcolo delle preferenze per magistrale >= 100
	pref1MagistraleGT100, pref2MagistraleGT100, pref3MagistraleGT100 := calcolaPreferenzeGruppo(db, `
                SELECT preferenza1, preferenza2, preferenza3
                FROM datiutente
                WHERE tipoDiLaureaUtente = ? AND votoDiLaureaUtente > ? AND
                     preferenza1 IS NOT NULL AND
                     preferenza2 IS NOT NULL AND
                     preferenza3 IS NOT NULL;`,
		"magistrale", 100,
	)
	fmt.Println("Preferenze magistrale > 100:", pref1MagistraleGT100, pref2MagistraleGT100, pref3MagistraleGT100)

	//calcolo delle preferenze per magistrale < 100
	pref1MagistraleLT100, pref2MagistraleLT100, pref3MagistraleLT100 := calcolaPreferenzeGruppo(db, `
                SELECT preferenza1, preferenza2, preferenza3
                FROM datiutente
                WHERE tipoDiLaureaUtente = ? AND votoDiLaureaUtente < ? AND
                     preferenza1 IS NOT NULL AND
                     preferenza2 IS NOT NULL AND
                     preferenza3 IS NOT NULL;`,
		"magistrale", 100,
	)
	fmt.Println("Preferenze magistrale < 100:", pref1MagistraleLT100, pref2MagistraleLT100, pref3MagistraleLT100)

	agents := []string{"Triennale", "Magistrale100m", "Magistrale100p"}
	fmt.Println("Categorie utenti: ", agents)
	houses := []string{"Cybersecurity", "Software", "Data Science"}
	fmt.Println("Categorie aziende: ", houses)
	agentPreferences := map[string][]string{
		"Triennale":      {pref1Triennale, pref2Triennale, pref3Triennale},
		"Magistrale100m": {pref1MagistraleLT100, pref2MagistraleLT100, pref3MagistraleLT100},
		"Magistrale100p": {pref1MagistraleGT100, pref2MagistraleGT100, pref3MagistraleGT100},
	}
	fmt.Println("\nPreferenze delle categorie di utenti:")
	for agent, preferences := range agentPreferences {
		fmt.Printf("Categoria %s: %v\n", agent, preferences)
	}
	initialOwnership := map[string]string{
		"Cybersecurity": "Triennale",
		"Software":      "Magistrale100m",
		"Data Science":  "Magistrale100p",
	}
	fmt.Println("\nAllocazione iniziale:")
	for casa, agente := range initialOwnership {
		fmt.Printf("Categoria di azienda %s: %v\n", casa, agente)
	}
	finalAllocation := make(map[string]string)
	//esecuzione effettiva del TTC
	finalAllocation = topTradingCycles(agents, houses, agentPreferences, initialOwnership)
	if scelteUtente.TipoDiLaurea == "triennale" {
		*mResult = MatchingResult(finalAllocation["Triennale"])
	} else if scelteUtente.TipoDiLaurea == "magistrale" {
		if scelteUtente.Voto < 100 {
			*mResult = MatchingResult(finalAllocation["Magistrale100m"])
		} else {
			*mResult = MatchingResult(finalAllocation["Magistrale100p"])
		}
	}

	_, err3 := db.Exec(`
        UPDATE datiutente
        SET ambitoAssegnato = ?
        WHERE emailUtente = ?;`,
		*mResult,
		scelteUtente.Email,
	)
	if err3 != nil {
		*mResult = "Error update database"
		log.Println("Errore:", err3)
	}

	companies, err4 := getCompaniesByResult(db, string(*mResult))
	if err4 != nil {
		log.Println("Errore: ", err4)
	}

	err5 := notifyWithRabbitMq(companies, scelteUtente.Email)
	if err5 != nil {
		log.Println("Errore: ", err5)
	}
	return nil
}

func getCompaniesByResult(db *sql.DB, mResult string) ([]string, error) {
	query := `SELECT DISTINCT nomeAzienda FROM aziende WHERE ambitoRichiesto = ?`
	rows, err := db.Query(query, mResult)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var companies []string
	for rows.Next() {
		var nomeAzienda string
		if err := rows.Scan(&nomeAzienda); err != nil {
			return nil, err
		}
		companies = append(companies, nomeAzienda)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return companies, nil
}

// funzione per notificare la coda di RabbitMQ
func notifyWithRabbitMq(companies []string, email string) error {
	_, err := pattern.CBRabbitMQ.Execute(func() (interface{}, error) {
		conn, err := amqp.Dial("amqp://myuser:mypassword@rabbitmq:5672/")
		if err != nil {
			return nil, fmt.Errorf("Errore nella connessione a RabbitMQ: %v", err)
		}
		defer conn.Close()

		ch, err := conn.Channel()
		if err != nil {
			return nil, fmt.Errorf("Errore nell'apertura del canale: %v", err)
		}
		defer ch.Close()

		q, err := ch.QueueDeclare(
			"matching", //nome della coda
			false,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			return nil, fmt.Errorf("Errore nella dichiarazione della coda: %v", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		message := struct {
			Email     string   `json:"email"`
			Companies []string `json:"companies"`
		}{
			Email:     email,
			Companies: companies,
		}

		body, err := json.Marshal(message)
		if err != nil {
			return nil, fmt.Errorf("Errore nel marshaling del messaggio: %v", err)
		}

		err = ch.PublishWithContext(ctx,
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        body,
			},
		)
		if err != nil {
			return nil, fmt.Errorf("Errore nel messaggio: %v", err)
		}

		return nil, nil
	})
	return err
}

// funzione per inserire un nuovo utente nel DB
func (t *MatchingService) Insert(datiUtente DatiUtente, Result *int) error {
	fmt.Println("Connessione al database...")
	db, err := OpenDB("matchingservice")

	if err != nil {
		*Result = -1
		fmt.Println(err.Error())
	}
	//chisura del database al termine di tutto
	defer db.Close()
	_, err2 := db.Exec(`
		    INSERT INTO datiutente (nomeUtente, cognomeUtente, emailUtente, votoDiLaureaUtente, universitàUtente, tipoDiLaureaUtente, password, etaUtente)
		    VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		datiUtente.Nome,
		datiUtente.Cognome,
		datiUtente.Email,
		datiUtente.VotoLaurea,
		datiUtente.Universita,
		datiUtente.TipoLaurea,
		datiUtente.Password,
		datiUtente.Eta,
	)
	*Result = 1
	if err2 != nil {
		*Result = -2
		log.Println("Errore: ", err2)
	}
	return nil
}

// funzione per il recupero dei dati di un utente autenticato
func (t *MatchingService) Retrieve(datiUtente2 DatiUtente2, utente *DatiUtente) error {
	db, err := OpenDB("matchingservice")
	if err != nil {
		log.Print(err.Error())
		*utente = DatiUtente{}
	}
	defer db.Close()
	// esecuzione della query
	row := db.QueryRow("SELECT nomeUtente, cognomeUtente, emailUtente, votoDiLaureaUtente, universitàUtente, tipoDiLaureaUtente, etaUtente, ambitoAssegnato FROM datiutente WHERE emailUtente = ? AND password = ?", datiUtente2.Email, datiUtente2.Password)
	err2 := row.Scan(&utente.Nome, &utente.Cognome, &utente.Email, &utente.VotoLaurea, &utente.Universita, &utente.TipoLaurea, &utente.Eta, &utente.Assegnamento)
	if err2 != nil {
		if err2 == sql.ErrNoRows {
			*utente = DatiUtente{}
		}
	}
	return nil
}

// funzione per l'inserimento di nuove aziende nel sistema
func (t *MatchingService) InsertCompanies(datiAzienda DatiAzienda, Result *int) error {
	db, err := OpenDB("matchingservice")
	if err != nil {
		*Result = -1
		panic(err.Error())
	}
	defer db.Close()
	_, err2 := db.Exec(`
		    INSERT INTO aziende (nomeAzienda, ambitoRichiesto)
		    VALUES (?, ?)`,
		datiAzienda.Nome,
		datiAzienda.Preferenza,
	)
	*Result = 1
	if err2 != nil {
		*Result = -2
	}
	return nil
}

// funzione per il recupero delle aziende nel sistema
func (t *MatchingService) RetrieveCompanies(datiAziendaInput DatiAzienda, companies2 *CompanyResult) error {
	db, err := OpenDB("matchingservice")
	if err != nil {
		return nil
	}
	defer db.Close()

	//query per ottenere i nomi univoci delle aziende
	query := `
		SELECT DISTINCT nomeAzienda
		FROM aziende
	`
	//esecuzione della query
	rows, err := db.Query(query)
	if err != nil {
		return nil
	}
	defer rows.Close()

	//memorizzazione dei nomi delle aziende
	var companies []string

	//scorrimento dei risultati e popolazione della lista delle aziende
	for rows.Next() {
		var nomeAzienda string

		//scansionamento del nome dell'azienda dalla query
		err := rows.Scan(&nomeAzienda)
		if err != nil {
			return nil
		}
		//aggiunta dell'azienda alla lista
		companies = append(companies, nomeAzienda)
	}
	*companies2 = companies
	//gestione di eventuali errori durante l'iterazione
	if err := rows.Err(); err != nil {
		return nil
	}
	return nil
}
