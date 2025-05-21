package handlers

import (
	"SdccProject/apiGateway/pattern"
	"SdccProject/notificationService/service2"
	"fmt"
	"log"
	"net/http"
	"os"
)

// email dell'utente da notificare
type recuperoNotifica struct {
	email string
}

// recupero l'indirizzo del servizio di notifica dalla variabile d'ambiente
func getNotificationServiceAddr() string {
	addr := os.Getenv("NOTIFICATION_SERVICE_ADDR")
	return addr
}

// Handler che gestisce la richiesta di reinviare una notifica
func SendNotification(w http.ResponseWriter, r *http.Request) {
	//Parsing del corpo della richiesta in formato JSON
	data, err := parseRequestBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//Estrazione dell'email dal payload
	input := recuperoNotifica{}
	if email, ok := data["Email"].(string); ok {
		input.email = email
	}
	fmt.Println("Dati notifica:", input)

	//Circuit breaker + chiamata RPC
	result, err := pattern.CBNotification.Execute(func() (interface{}, error) {
		addr := getNotificationServiceAddr()
		dati := &service2.DatiNotifica{input.email}
		var reply service2.NotifyResult
		//chiamata RPC al servizio di notifica
		err := callRPCservice("Notification", "Notify", addr, dati, &reply)
		if err != nil {
			return nil, err
		}
		return reply, nil
	})

	//risposta HTTP
	if err != nil {
		log.Printf("Errore nella notifica (o circuito aperto): %v", err)
		sendJSON(w, map[string]interface{}{
			"error":   "Servizio di notifica non disponibile",
			"details": err.Error(),
		})
		return
	}
	//invio della risposta
	sendJSON(w, map[string]interface{}{"result": result})
	return
}
