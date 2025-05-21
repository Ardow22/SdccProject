package main

import (
	"SdccProject/matchingService/pattern"
	"SdccProject/matchingService/service"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

const (
	port       = ":8081" //porta sulla quale il server RPC ascolterà
	rpcPath    = "/"
	debugPath  = "/debug"
	serviceKey = "Matching"
)

func main() {
	//inizializzazione del circuitBreaker
	pattern.InitCircuitBreakers()
	startRPCServer()
}

func startRPCServer() {
	//creazione di un'istanza del MatchingService
	matchingService := new(service.MatchingService)
	//inizializzazione di un nuovo server RPC
	server := rpc.NewServer()
	if err := server.RegisterName(serviceKey, matchingService); err != nil {
		log.Fatalf("Failed to register RPC service '%s': %v", serviceKey, err)
	}

	//aggiunta dell'handler per le richieste RPC e il debug
	server.HandleHTTP(rpcPath, debugPath)
	//listener per ascoltare sulla porta specificata
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}
	log.Printf("RPC server is running on port %s", port)
	//avvio del server HTTP
	if err := http.Serve(listener, nil); err != nil {
		log.Fatalf("HTTP server error: %v", err)
	}
}
