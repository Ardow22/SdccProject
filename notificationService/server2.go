package main

import (
	"SdccProject/notificationService/service2"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

const (
	port       = ":8082" //porta sulla quale il server RPC ascolterà
	rpcPath    = "/"
	debugPath  = "/debug"
	serviceKey = "Notification"
)

func main() {
	startRPCServer()
	startRabbitMQConsumer()

	//blocco la goroutine principale per mantenere il programma in esecuzione
	select {}
}

// avvio del server RPC
func startRPCServer() {
	//nuova istanza del servizio di notifica
	notificationService := new(service2.NotificationService)
	//nuovo server
	server := rpc.NewServer()
	if err := server.RegisterName(serviceKey, notificationService); err != nil {
		log.Fatalf("Errore durante la registrazione del servizio '%s': %v", serviceKey, err)
	}
	//aggiunta dell'handler per le richieste RPC e il debug
	server.HandleHTTP(rpcPath, debugPath)
	//listener per ascoltare sulla porta specificata
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Errore di ascolto sulla porta %s: %v", port, err)
	}

	log.Println("%s RPC server in ascolto sulla porta %s", serviceKey, port)
	//avvio del server HTTP in una goroutine separata
	go func() {
		if err := http.Serve(listener, nil); err != nil {
			log.Fatalf("Errore del server HTTP: %v", err)
		}
	}()
}

// funzione per l'avvio del consumer RabbitMQ
func startRabbitMQConsumer() {
	//avvio del consumer RabbitMQ in una goroutine separata
	go func() {
		log.Println("Avvio del consumer RabbitMQ...")
		//chiamo la funzione che consuma i messaggi RabbitMQ
		service2.ConsumeRabbitMQMessages()
	}()
}
