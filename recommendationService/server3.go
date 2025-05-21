package main

import (
	"SdccProject/recommendationService/service3"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

const (
	port       = ":8083" //porta sulla quale il server RPC ascolterà
	rpcPath    = "/"
	debugPath  = "/debug"
	serviceKey = "Recommendation"
)

func main() {
	startRPCServer()
}

func startRPCServer() {
	//creazione di un'istanza del RecommendationService
	recommendationService := new(service3.RecommendationService)

	//inizializzazione di un nuovo server RPC
	server := rpc.NewServer()

	//registrazione del servizio con un nome
	if err := server.RegisterName(serviceKey, recommendationService); err != nil {
		log.Fatalf("Registrazione del servizio '%s' fallita: %v", serviceKey, err)
	}

	//aggiunta dell'handler per le richieste RPC e il debug
	server.HandleHTTP(rpcPath, debugPath)

	//listener per ascoltare sulla porta specificata
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Errore di ascolto sulla porta %s: %v", port, err)
	}

	log.Println("Recommendation RPC server in ascolto sulla porta %s", port)

	//avvio del server HTTP
	if err := http.Serve(listener, nil); err != nil {
		log.Fatalf("Errore durante l'avvio del server HTTP: %v", err)
	}
}
