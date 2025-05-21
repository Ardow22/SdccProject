package handlers

import (
	"SdccProject/apiGateway/pattern"
	"SdccProject/matchingService/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"os"
	"strconv"
)

// le seguenti struct servono a preparare i dati dell'utente da inviare alle relative RPC
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

type DatiMatching struct {
	Email      string
	scelta1    string
	scelta2    string
	scelta3    string
	Voto       int
	TipoLaurea string
}

type DatiUtente2 struct {
	Email    string
	Password string
}

type InfoAzienda struct {
	Nome       string
	Preferenza string
}

// ottengo l'indirizzo del servizio Matching dalla variabile d'ambiente
func getMatchingServiceAddr() string {
	addr := os.Getenv("MATCHING_SERVICE_ADDR")
	return addr
}

// chiamata rpc generica ad un servizio remoto
func callRPCservice(serviceName, methodName, addr string, args interface{}, reply interface{}) error {
	client, err := rpc.DialHTTP("tcp", addr)
	if err != nil {
		return fmt.Errorf("Errore durante la connessione RPC al servizio %s: %v", serviceName, err)
	}
	defer client.Close()

	//creazione del nome completo del metodo
	fullMethod := fmt.Sprintf("%s.%s", serviceName, methodName)

	log.Printf("Chiamata sincrona al metodo %s del servizio %s", fullMethod, serviceName)
	err = client.Call(fullMethod, args, reply)
	if err != nil {
		return fmt.Errorf("Errore durante la chiamata RPC a %s: %v", fullMethod, err)
	}
	return nil
}

// Handler per l'inserimento di un nuovo utente nel sistema
func HandleMatchingInsert(w http.ResponseWriter, r *http.Request) {
	data, err := parseRequestBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//riempimento struct da inviare
	input := DatiUtente{}
	if nome, ok := data["name"].(string); ok {
		input.Nome = nome
	}
	if cognome, ok := data["surname"].(string); ok {
		input.Cognome = cognome
	}
	if email, ok := data["email_address"].(string); ok {
		input.Email = email
	}
	if universita, ok := data["university"].(string); ok {
		input.Universita = universita
	}
	if tipo, ok := data["type"].(string); ok {
		input.TipoLaurea = tipo
	}
	if password, ok := data["password"].(string); ok {
		input.Password = password
	}
	//conversione di "age" da stringa a intero
	ageStr, ok := data["age"].(string)
	if !ok {
		log.Println("age non trovato o tipo errato")
	}
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		log.Println("age non è un numero valido")
	}
	input.Eta = age
	//conversione di "Dgrade" da stringa a intero
	dgradeStr, ok := data["Dgrade"].(string)
	if !ok {
		log.Println("Dgrade non trovata o tipo errato")
	}
	dgrade, err := strconv.Atoi(dgradeStr)
	if err != nil {
		log.Println("Dgrade non è un numero valido")
	}
	input.VotoLaurea = dgrade
	input.Assegnamento = ""

	log.Println("Utente:", input)

	//circuit breaker + chiamata RPC
	result, err := pattern.CBMatching.Execute(func() (interface{}, error) {
		addr := getMatchingServiceAddr()
		dati := &service.DatiUtente{input.Nome, input.Cognome, input.Eta, input.Email,
			input.VotoLaurea, input.Universita, input.TipoLaurea,
			input.Password, input.Assegnamento}
		var reply service.Result
		//chiamata RPC al servizio
		err := callRPCservice("Matching", "Insert", addr, dati, &reply)
		if err != nil {
			return nil, err
		}
		return reply, nil
	})

	//risposta HTTP
	if err != nil {
		log.Printf("Errore nel matching (o circuito aperto): %v", err)
		sendJSON(w, map[string]interface{}{
			"error":   "Servizio di matching non disponibile",
			"details": err.Error(),
		})
		return
	}
	sendJSON(w, map[string]interface{}{"result": result})
	return
}

// Handler per il matching tra utente e aziende basato sulle scelte
func HandleMatching(w http.ResponseWriter, r *http.Request) {
	data, err := parseRequestBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//riempimento struct da inviare
	input := DatiMatching{}
	if primaScelta, ok := data["Cfield1"].(string); ok {
		input.scelta1 = primaScelta
	}
	if secondaScelta, ok := data["Cfield2"].(string); ok {
		input.scelta2 = secondaScelta
	}
	if terzaScelta, ok := data["Cfield3"].(string); ok {
		input.scelta3 = terzaScelta
	}
	if email, ok := data["email"].(string); ok {
		input.Email = email
	}
	if corso, ok := data["corso"].(string); ok {
		input.TipoLaurea = corso
	}
	votoInterface, ok := data["voto"]
	if !ok {
		log.Fatal("voto non trovato")
	}
	voto, ok := votoInterface.(float64) //assumo che il voto ricevuto sia float64
	if !ok {
		log.Fatal("voto non è un numero valido")
	}
	input.Voto = int(voto)
	fmt.Println("UtenteMatching:", input)

	//circuit breaker + chiamata RPC
	result, err := pattern.CBMatching.Execute(func() (interface{}, error) {
		addr := getMatchingServiceAddr()
		dati := &service.ScelteUtente{input.Email, input.scelta1, input.scelta2,
			input.scelta3, input.Voto, input.TipoLaurea}
		var reply service.MatchingResult
		//chiamata RPC al servizio
		err := callRPCservice("Matching", "Match", addr, dati, &reply)
		if err != nil {
			return nil, err
		}
		return reply, nil
	})

	// Risposta HTTP
	if err != nil {
		log.Printf("Errore nel matching (o circuito aperto): %v", err)
		sendJSON(w, map[string]interface{}{
			"error":   "Servizio di matching non disponibile",
			"details": err.Error(),
		})
		return
	}
	sendJSON(w, map[string]interface{}{"result": result})
	return
}

// Handler per il recupero dati di un utente registrato
func HandleMatchingRetrieve(w http.ResponseWriter, r *http.Request) {
	data, err := parseRequestBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//riempimento struct con i dati da inviare
	input := DatiUtente2{}
	if email, ok := data["email_address"].(string); ok {
		input.Email = email
	}
	if password, ok := data["password"].(string); ok {
		input.Password = password
	}
	fmt.Println("Utente:", input)

	//Circuit breaker + chiamata RPC
	result, err := pattern.CBMatching.Execute(func() (interface{}, error) {
		addr := getMatchingServiceAddr()
		dati := &service.DatiUtente2{input.Email, input.Password}
		var reply service.DatiUtente
		//chiamata RPC
		err := callRPCservice("Matching", "Retrieve", addr, dati, &reply)
		if err != nil {
			return nil, err
		}
		return reply, nil
	})

	//risposta HTTP
	if err != nil {
		log.Printf("Errore nel matching (o circuito aperto): %v", err)
		sendJSON(w, map[string]interface{}{
			"error":   "Servizio di matching non disponibile",
			"details": err.Error(),
		})
		return
	}
	sendJSON(w, map[string]interface{}{"result": result})
	return
}

// Handler per l'inserimento di una nuova azienda nel sistema
func HandleMatchingInsertCompanies(w http.ResponseWriter, r *http.Request) {
	data, err := parseRequestBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	input := InfoAzienda{}
	if preferenza, ok := data["careerField"].(string); ok {
		input.Preferenza = preferenza
	}
	if nome, ok := data["companyName"].(string); ok {
		input.Nome = nome
	}
	fmt.Println("Azienda:", input)

	//Circuit breaker + chiamata RPC
	result, err := pattern.CBMatching.Execute(func() (interface{}, error) {
		addr := getMatchingServiceAddr()
		dati := &service.DatiAzienda{input.Nome, input.Preferenza}
		var reply service.Result

		err := callRPCservice("Matching", "InsertCompanies", addr, dati, &reply)
		if err != nil {
			return nil, err
		}
		return reply, nil
	})

	//risposta HTTP
	if err != nil {
		log.Printf("Errore nel matching (o circuito aperto): %v", err)
		sendJSON(w, map[string]interface{}{
			"error":   "Servizio di matching non disponibile",
			"details": err.Error(),
		})
		return
	}
	sendJSON(w, map[string]interface{}{"result": result})
	return
}

// Handler per il recupero delle aziende registrate (placeholder)
func HandleMatchingRetrieveCompanies(w http.ResponseWriter, r *http.Request) {
	dummy := InfoAzienda{
		Nome:       "ciao",
		Preferenza: "ciao",
	}

	//Circuit breaker + chiamata RPC
	result, err := pattern.CBMatching.Execute(func() (interface{}, error) {
		addr := getMatchingServiceAddr()
		dati := &service.DatiAzienda{dummy.Nome, dummy.Preferenza}
		var reply service.CompanyResult

		err := callRPCservice("Matching", "RetrieveCompanies", addr, dati, &reply)
		if err != nil {
			return nil, err
		}
		return reply, nil
	})

	//risposta HTTP
	if err != nil {
		log.Printf("Errore nel matching (o circuito aperto): %v", err)
		sendJSON(w, map[string]interface{}{
			"error":   "Servizio di matching non disponibile",
			"details": err.Error(),
		})
		return
	}
	sendJSON(w, map[string]interface{}{"result": result})
	return
}

// Funzione helper per inviare una risposta JSON
func sendJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(data)
}
