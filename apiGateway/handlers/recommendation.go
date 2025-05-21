package handlers

import (
	"SdccProject/apiGateway/pattern"
	"SdccProject/matchingService/service"
	"SdccProject/recommendationService/service3"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

//le seguenti struct servono a preparare i dati da inviare alla RPC

// struct per contenere i dati di una votazione utente
type DatiVoto struct {
	Email     string
	Corso     string
	Voto      int
	Azienda   string
	Votazione int
}

// struct per contenere i dati necessari a generare consigli
type DatiConsigli struct {
	Email   string
	Corso   string
	Voto    int
	Azienda string
}

// l'indirizzo del servizio è recuperato dalla variabile d'ambiente
func getRecommendationServiceAddr() string {
	addr := os.Getenv("RECOMMENDATION_SERVICE_ADDR")
	return addr
}

// funzione che legge e decodifica il corpo della richiesta JSON in una map
func parseRequestBody(r *http.Request) (map[string]interface{}, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("errore nella lettura del corpo della richiesta: %w", err)
	}
	log.Printf("Corpo della richiesta JSON: %s\n", body)

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("errore nell'analisi del JSON: %w", err)
	}

	fmt.Println("I nuovi dati sono:", data)
	return data, nil
}

// Handler per inserire una valutazione da parte dell’utente
func HandleRecommendationInsert(w http.ResponseWriter, r *http.Request) {
	data, err := parseRequestBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//parsing dei dati dal body
	input := DatiVoto{}
	if email, ok := data["Email"].(string); ok {
		input.Email = email
	}
	if corso, ok := data["Corso"].(string); ok {
		input.Corso = corso
	}
	if azienda, ok := data["azienda"].(string); ok {
		input.Azienda = azienda
	}
	votazioneStr, ok := data["voto_esperienza"].(string)
	if !ok {
		log.Println("age non trovata o tipo errato")
	}
	votazione, err := strconv.Atoi(votazioneStr)
	if err != nil {
		log.Println("age non è un numero valido")
	}
	input.Votazione = votazione

	votoInterface, ok := data["Voto"]
	if !ok {
		log.Println("voto non trovato")
	}
	voto, ok := votoInterface.(float64)
	if !ok {
		log.Println("voto non è un numero valido")
	}
	input.Voto = int(voto)
	log.Println("Dati voto:", input)

	//Circuit breaker + chiamata RPC
	result, err := pattern.CBRecommendation.Execute(func() (interface{}, error) {
		addr := getRecommendationServiceAddr()
		dati := &service3.DatiVotazione{input.Email, input.Corso, input.Azienda,
			input.Voto, input.Votazione}
		var reply service.Result
		//chiamata RPC
		err := callRPCservice("Recommendation", "Insert", addr, dati, &reply)
		if err != nil {
			return nil, err
		}
		return reply, nil
	})

	//risposta HTTP
	if err != nil {
		log.Printf("Errore nella notifica (o circuito aperto): %v", err)
		sendJSON(w, map[string]interface{}{
			"error":   "Servizio di raccomandazione non disponibile",
			"details": err.Error(),
		})
		return
	}
	//invio della risposta a buon fine
	sendJSON(w, map[string]interface{}{"result": result})
	return
}

// Handler per ottenere consigli sulla base dei voti e delle aziende
func HandleRecommendation(w http.ResponseWriter, r *http.Request) {
	data, err := parseRequestBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//riempimento struct con i dati da inviare
	input := DatiConsigli{}
	if email, ok := data["email"].(string); ok {
		input.Email = email
	}
	if corso, ok := data["corso"].(string); ok {
		input.Corso = corso
	}
	if azienda, ok := data["companyName"].(string); ok {
		input.Azienda = azienda
	}
	votoInterface, ok := data["voto"]
	if !ok {
		log.Println("voto non trovato")
	}
	voto, ok := votoInterface.(float64)
	if !ok {
		log.Println("voto non è un numero valido")
	}
	input.Voto = int(voto)
	log.Println("Consigli utente:", input)

	//Circuit breaker + chiamata RPC
	result, err := pattern.CBRecommendation.Execute(func() (interface{}, error) {
		addr := getRecommendationServiceAddr()
		dati := &service3.ConsigliUtente{input.Email, input.Corso, input.Voto,
			input.Azienda}
		var reply service.Result
		err := callRPCservice("Recommendation", "Recommendations", addr, dati, &reply)
		if err != nil {
			return nil, err
		}
		return reply, nil
	})

	//risposta HTTP con errore
	if err != nil {
		log.Printf("Errore nella notifica (o circuito aperto): %v", err)
		sendJSON(w, map[string]interface{}{
			"error":   "Servizio di raccomandazione non disponibile",
			"details": err.Error(),
		})
		return
	}
	sendJSON(w, map[string]interface{}{"result": result})
	return
}

// Handler per recuperare aziende consigliate
func HandleRecommendationRetrieveCompanies(w http.ResponseWriter, r *http.Request) {
	//azienda fittizia come input
	dummyAzienda := InfoAzienda{
		Nome:       "ciao",
		Preferenza: "ciao",
	}

	//Circuit breaker + chiamata RPC
	result, err := pattern.CBRecommendation.Execute(func() (interface{}, error) {
		addr := getRecommendationServiceAddr()
		dati := &service3.DatiAzienda{dummyAzienda.Nome, dummyAzienda.Preferenza}
		var reply service.CompanyResult
		err := callRPCservice("Recommendation", "RetrieveCompanies", addr, dati, &reply)
		if err != nil {
			return nil, err
		}
		return reply, nil
	})

	//risposta HTTP
	if err != nil {
		log.Printf("Errore nella notifica (o circuito aperto): %v", err)
		sendJSON(w, map[string]interface{}{
			"error":   "Servizio di raccomandazione non disponibile",
			"details": err.Error(),
		})
		return
	}
	sendJSON(w, map[string]interface{}{"result": result})
	return
}
