package service3

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type DatiVotazione struct {
	Email     string
	Corso     string
	Azienda   string
	Voto      int
	Votazione int
}

type ConsigliUtente struct {
	Email   string
	Corso   string
	Voto    int
	Azienda string
}

type DatiAzienda struct {
	Nome       string
	Preferenza string
}

type RecommendationService int

type Result int

type CompanyResult []string

// -----------funzioni per la connessione al database-------------------
func getDatabaseDSN(database string) string {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)
}

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

// -------------------funzioni offerte dal servizio--------------------------------------------------
// funzione per l'inserimento di una votazione per un'azienda da parte di un utente
func (t *RecommendationService) Insert(datiVotazione DatiVotazione, nResult *Result) error {
	log.Println("Sto inviando: ", datiVotazione)
	//apertura del database
	db, err := OpenDB("recommendationservice")
	*nResult = 1
	if err != nil {
		*nResult = -1
		panic(err.Error())
	}
	defer db.Close()
	//esecuzione della query
	_, err2 := db.Exec(`
		    INSERT INTO votazione (mailUtente, corsoUtente, nomeAzienda, votoUtente, votazione)
		    VALUES (?, ?, ?, ?, ?)
		    ON DUPLICATE KEY UPDATE 
		    votazione = VALUES(votazione)`,
		datiVotazione.Email,
		datiVotazione.Corso,
		datiVotazione.Azienda,
		datiVotazione.Voto,
		datiVotazione.Votazione,
	)
	if err2 != nil {
		*nResult = -2
		log.Println(err2)
	}
	return nil
}

// ----------------------------funzioni per l'algoritmo di raccomandazione----------------------------
// funzione ausiliaria per calcolare la similarità tra due utenti
func similarity(user1, user2 map[string]float64) float64 {
	var dotProduct, magnitude1, magnitude2 float64
	for item, rating1 := range user1 {
		if rating2, ok := user2[item]; ok {
			dotProduct += rating1 * rating2
		}
		magnitude1 += rating1 * rating1
	}
	for _, rating2 := range user2 {
		magnitude2 += rating2 * rating2
	}
	if magnitude1 == 0 || magnitude2 == 0 {
		return 0
	}
	return dotProduct / (math.Sqrt(magnitude1) * math.Sqrt(magnitude2))
}

// funzione di utilità per verificare se una stringa è presente in una slice
func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func getUserPreferences(corsoUtente string, votoUtente int) (map[string]map[string]float64, error) {
	db, err := OpenDB("recommendationservice")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	//definizione della query in base alla condizione
	var query string
	if corsoUtente == "triennale" {
		query = `
			SELECT mailUtente, nomeAzienda, votazione
			FROM votazione
			WHERE corsoUtente = "triennale"`
	} else if corsoUtente == "magistrale" && votoUtente < 100 {
		query = `
			SELECT mailUtente, nomeAzienda, votazione
			FROM votazione
			WHERE corsoUtente = "magistrale" AND votoUtente < 100`
	} else if corsoUtente == "magistrale" && votoUtente >= 100 {
		query = `
			SELECT mailUtente, nomeAzienda, votazione
			FROM votazione
			WHERE corsoUtente = "magistrale" AND votoUtente >= 100`
	}
	//esecuzione effettiva la query
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//map per memorizzare le preferenze
	preferences := make(map[string]map[string]float64)

	//scorrimento dei risultati e popolazione di map
	for rows.Next() {
		var emailUtente, nomeAzienda string
		var votazione float64

		//Scansionamento dei valori dalla query
		err := rows.Scan(&emailUtente, &nomeAzienda, &votazione)
		if err != nil {
			return nil, err
		}

		//se non esiste una map per questa email, creo una nuova mappa
		if _, exists := preferences[emailUtente]; !exists {
			preferences[emailUtente] = make(map[string]float64)
		}

		//aggiunta la votazione per l'azienda
		preferences[emailUtente][nomeAzienda] = votazione
	}
	//gestione di eventuali errori durante l'iterazione
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return preferences, nil
}

// funzione ausiliaria per ottenere i corsi e i voti degli utenti (una sola volta per utente)
func getCourses() (map[string]string, error) {
	db, err := OpenDB("recommendationservice")
	if err != nil {
		log.Println("Errore1:", err)
		return nil, err
	}
	defer db.Close()

	//query per ottenere emailUtente, corsoUtente e votoUtente
	query := `
		SELECT mailUtente, corsoUtente, votoUtente
		FROM votazione
		GROUP BY mailUtente
	`
	//esecuzione effettiva della query
	rows, err := db.Query(query)
	if err != nil {
		log.Println("Errore2:", err)
		return nil, err
	}
	defer rows.Close()

	//map per memorizzare i corsi e i voti
	userCourses := make(map[string]string)

	//scorrimento dei risultati e popolazione della map
	for rows.Next() {
		var emailUtente, corsoUtente string
		var votoUtente int

		//scansione dei valori dalla query
		err := rows.Scan(&emailUtente, &corsoUtente, &votoUtente)
		if err != nil {
			fmt.Println("Errore3:", err)
			return nil, err
		}

		//creazione di una stringa con corso e voto e inserimento nella map
		userCourses[emailUtente] = fmt.Sprintf("%s,%d", corsoUtente, votoUtente)
	}

	//gestione di eventuali errori durante l'iterazione
	if err := rows.Err(); err != nil {
		fmt.Println("Errore4:", err)
		return nil, err
	}

	return userCourses, nil
}

// funzione ausiliaria per ottenere i nomi delle aziende una sola volta
func getUniqueCompanies() ([]string, error) {
	db, err := OpenDB("recommendationservice")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	//query per ottenere i nomi univoci delle aziende
	query := `
		SELECT DISTINCT nomeAzienda
		FROM votazione
	`
	//esecuzione della query
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//map per memorizzare i nomi delle aziende
	var companies []string

	//scorrimento dei risultati e popolazione della lista delle aziende
	for rows.Next() {
		var nomeAzienda string

		//scansionamento del nome dell'azienda dalla query
		err := rows.Scan(&nomeAzienda)
		if err != nil {
			return nil, err
		}

		//aggiunta dell'azienda alla lista
		companies = append(companies, nomeAzienda)
	}

	//gestione di eventuali errori durante l'iterazione
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return companies, nil
}

// funzione per generare raccomandazioni per un utente, implementa un algoritmo di raccomandazione
func (t *RecommendationService) Recommendations(consigliUtente ConsigliUtente, res *Result) error {
	fmt.Println("\nALGORITMO DI RACCOMANDAZIONE\n")
	*res = -10 //valore di default iniziale del valore di ritorno
	preferences, err := getUserPreferences(consigliUtente.Corso, consigliUtente.Voto)
	if err != nil {
		*res = -1
		log.Println("Errore: ", err)
	}
	for user, companies := range preferences {
		fmt.Printf("%s: ", user)
		for company, vote := range companies {
			fmt.Printf("%s: %.2f ", company, vote)
		}
		fmt.Println()
	}
	var userID = consigliUtente.Email
	userCourses, err := getCourses()
	if err != nil {
		*res = -2
		log.Println("Errore: ", err)
	}
	fmt.Println("Corso e voto dell'utente:")
	for email, course := range userCourses {
		fmt.Printf("%s: %s\n", email, course)
	}
	companies, err := getUniqueCompanies()
	if err != nil {
		*res = -3
		log.Println("Errore: ", err)
	}
	fmt.Println("Unique Companies:")
	for _, company := range companies {
		log.Println(company)
	}
	log.Printf("\nAnalisi dell'utente %s...\n", userID)
	userInfo := strings.Split(userCourses[userID], ",")
	if len(userInfo) != 2 {
		return nil //dati utente non validi
	}
	log.Printf("L'utente %s ha voto %s e tipo di laurea %s\n", userID, userInfo[1], userInfo[0])
	userDegree, err := strconv.Atoi(userInfo[1])
	if err != nil {
		log.Println("Errore: ", err)
		return nil //voto di laurea non valido
	}
	userType := userInfo[0]

	//calcolo del profilo medio degli utenti dello stesso tipo di laurea e voto
	courseUsers := make(map[string]bool)
	log.Println("\nRicerca di utenti con la stessa categoria di tipo di laure e voto...")
	for otherID, otherCourseInfo := range userCourses {
		if otherID != userID {
			otherUserInfo := strings.Split(otherCourseInfo, ",")
			if len(otherUserInfo) != 2 {
				continue //dati utente non validi, passa al prossimo utente
			}
			otherDegree, err := strconv.Atoi(otherUserInfo[1])
			if err != nil {
				continue //voto di laurea non valido
			}
			otherType := otherUserInfo[0]

			if len(preferences[otherID]) > 0 { //Controllo che l'utente abbia voti
				if userType == "triennale" && otherType == "triennale" {
					courseUsers[otherID] = true
				} else if userType == "magistrale" && otherType == "magistrale" {
					if userDegree < 100 && otherDegree < 100 {
						courseUsers[otherID] = true
					} else if userDegree >= 100 && otherDegree >= 100 {
						courseUsers[otherID] = true
					}
				}
			}
		}
	}

	log.Println("\nUtenti trovati della stessa categoria di tipo di laurea e voto:", courseUsers)
	if len(courseUsers) == 0 {
		fmt.Println("Nessun altro utente trovato della stessa categoria di tipo di laurea e voto")
		*res = -20 //nessun altro utente dello stesso tipo di laurea e voto
		return nil
	}

	courseAvgPrefs := make(map[string]float64)
	fmt.Println("\nCalcolo del profilo medio degli utenti...")
	for otherID := range courseUsers {
		otherPrefs := preferences[otherID]
		for company, rating := range otherPrefs {
			if contains(companies, company) {
				courseAvgPrefs[company] += rating
				fmt.Printf("  - Aggiungo il voto %f (voto dell'utente %s per l'azienda %s) al profilo medio.\n", rating, otherID, company)
			}
		}
	}
	fmt.Println("\nProfilo medio degli utenti (somma dei voti):", courseAvgPrefs)
	for company := range courseAvgPrefs {
		courseAvgPrefs[company] /= float64(len(courseUsers))
		fmt.Printf("  - Calcolo la media per l'azienda %s dividendo la somma dei voti per %d (numero di utenti).\n", company, len(courseUsers))
	}
	fmt.Println("\nProfilo medio normalizzato (media dei voti):", courseAvgPrefs)

	//calcolo delle similarità con gli altri utenti e i punteggi ponderati
	similarities := make(map[string]float64)
	fmt.Println("\nCalcolo delle similarità con gli altri utenti...")
	for otherID, otherPrefs := range preferences {
		otherUserInfo := strings.Split(userCourses[otherID], ",")
		otherDegree, err := strconv.Atoi(otherUserInfo[1])
		if err != nil {
			continue
		}
		otherType := otherUserInfo[0]
		if otherID != userID && len(preferences[otherID]) > 0 { //controllo che l'utente abbia voti
			if userType == "triennale" && otherType == "triennale" {
				similarities[otherID] = similarity(courseAvgPrefs, otherPrefs)
				fmt.Printf("  - Calcolo similarità tra %s e il profilo medio: %f (prodotto scalare / prodotto delle magnitudini).\n", otherID, similarities[otherID])
			} else if userType == "magistrale" && otherType == "magistrale" {
				if userDegree < 100 && otherDegree < 100 {
					similarities[otherID] = similarity(courseAvgPrefs, otherPrefs)
					fmt.Printf("  - Calcolo similarità tra %s e il profilo medio: %f (prodotto scalare / prodotto delle magnitudini).\n", otherID, similarities[otherID])
				} else if userDegree >= 100 && otherDegree >= 100 {
					similarities[otherID] = similarity(courseAvgPrefs, otherPrefs)
					fmt.Printf("  - Calcolo similarità tra %s e il profilo medio: %f (prodotto scalare / prodotto delle magnitudini).\n", otherID, similarities[otherID])
				}
			}
		}
	}
	fmt.Println("\nSimilarità calcolate:", similarities)

	type neighbor struct {
		id    string
		score float64
	}
	var neighbors []neighbor
	for id, score := range similarities {
		if score > 0 {
			neighbors = append(neighbors, neighbor{id, score})
		}
	}
	sort.Slice(neighbors, func(i, j int) bool {
		return neighbors[i].score > neighbors[j].score
	})
	companyScores := make(map[string]float64)
	fmt.Println("\nCalcolo dei punteggi delle aziende...")
	for _, neighbor := range neighbors {
		prefs := preferences[neighbor.id]
		for company, rating := range prefs {
			if contains(companies, company) {
				companyScores[company] += neighbor.score * rating
				fmt.Printf("  - Aggiungo %f (score di %s) * %f (voto di %s per %s).\n", neighbor.score, neighbor.id, rating, neighbor.id, company)
			}
		}
	}
	fmt.Println("\nPunteggi delle aziende:", companyScores)

	type recommendation struct {
		company string
		score   float64
	}
	var recs []recommendation
	for company, score := range companyScores {
		recs = append(recs, recommendation{company, score})
	}
	sort.Slice(recs, func(i, j int) bool {
		return recs[i].score > recs[j].score
	})
	if len(recs) > 0 {
		fmt.Printf("Aziende raccomandate per %s:\n", userID)
		for i, rec := range recs {
			fmt.Printf("%d. %s\n", i+1, rec.company)
		}
	} else {
		fmt.Printf("Nessuna raccomandazione per %s.\n", userID)
	}

	fmt.Println("\nRaccomandazioni finali:", recs)
	for i, rec := range recs {
		if rec.company == consigliUtente.Azienda {
			*res = Result(i) + 1
		}
	}
	return nil
}

// -------------------------------------------------------------------------------------------
// funzione per recuperare l'elenco delle aziende che sono state votate almeno una volta
func (t *RecommendationService) RetrieveCompanies(datiAziendaInput DatiAzienda, companies2 *CompanyResult) error {
	db, err := OpenDB("recommendationservice")
	if err != nil {
		return nil
	}
	defer db.Close()

	//query per ottenere i nomi univoci delle aziende
	query := `
		SELECT DISTINCT nomeAzienda
		FROM votazione
	`
	//esecuzione della query
	rows, err := db.Query(query)
	if err != nil {
		return nil
	}
	defer rows.Close()

	//variabile per memorizzare i nomi delle aziende
	var companies []string

	//scorrimento dei risultati del database e popolazione della lista delle aziende
	for rows.Next() {
		var nomeAzienda string

		//scansionamento del nome dell'azienda dalla query
		err := rows.Scan(&nomeAzienda)
		if err != nil {
			return nil
		}
		//si aggiunge l'azienda alla lista
		companies = append(companies, nomeAzienda)
	}
	*companies2 = companies
	//gestione di eventuali errori durante l'iterazione
	if err := rows.Err(); err != nil {
		return nil
	}

	return nil
}
