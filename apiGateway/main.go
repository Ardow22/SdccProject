package main

import (
	"SdccProject/apiGateway/handlers"
	"SdccProject/apiGateway/pattern"
	"log"
	"net/http"
	"time"
)

// middleware che abilita il CORS per consentire richieste cross-origin da browser
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	//inizializzazione dei Circuit Breaker per ogni servizio
	pattern.InitCircuitBreakers()

	//Creazione di un router HTTP per gestire le richieste in entrata
	mux := http.NewServeMux()

	//Matching Service
	mux.HandleFunc("/service1/match/insert", handlers.HandleMatchingInsert)
	mux.HandleFunc("/service1/match/matching", handlers.HandleMatching)
	mux.HandleFunc("/service1/match/retrieve", handlers.HandleMatchingRetrieve)
	mux.HandleFunc("/service1/match/insertCompanies", handlers.HandleMatchingInsertCompanies)
	mux.HandleFunc("/service1/match/retrieveCompanies", handlers.HandleMatchingRetrieveCompanies)

	//Notification Service
	mux.HandleFunc("/service2/notify", handlers.SendNotification)

	//Recommendation Service
	mux.HandleFunc("/service3/recommendation/insert", handlers.HandleRecommendationInsert)
	mux.HandleFunc("/service3/recommendation/recommend", handlers.HandleRecommendation)
	mux.HandleFunc("/service3/recommendation/retrieveCompanies", handlers.HandleRecommendationRetrieveCompanies)

	//Configurazione del server HTTP con timeout e abilitazione CORS per tutte le routes
	server := &http.Server{
		Addr:           ":12345",
		Handler:        enableCORS(mux), //middleware CORS
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("API Gateway in ascolto sulla porta 12345...")
	log.Fatal(server.ListenAndServe())
}
