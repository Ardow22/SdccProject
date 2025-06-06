package service2

import (
	"context"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"net/smtp"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Costanti di configurazione
const (
	smtpServer   = "smtp.gmail.com"
	smtpPort     = "587"
	smtpUser     = "niente"
	smtpPassword = "niente"
	queueName    = "matching"
)

type DatiNotifica struct {
	Email string
}

type DatiMail struct {
	Email string
}

type RabbitMQMessage struct {
	Email     string   `json:"email"`
	Companies []string `json:"companies"`
}

type NotificationService int
type NotifyResult int

// -----------funzioni per la connessione al database-------------------
// funzione per il recupero URI per MongoDB dalle variabili d'ambiente
func getMongoURI() string {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	return fmt.Sprintf("mongodb://%s:%s@%s:%s", user, password, host, port)
}

// funzione per l'apertura di una connessione al database MongoDB
func OpenMongoDB() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(getMongoURI()))
	if err != nil {
		return nil, fmt.Errorf("errore nella connessione a MongoDB: %w", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("errore nel ping di MongoDB: %w", err)
	}

	return client, nil
}

// funzione per l'invio effettivo della mail
func sendEmail(to string, body string) error {
	log.Printf("Invio email a: %s\n", to)
	log.Printf("Corpo dell'email: %s\n", body)
	auth := smtp.PlainAuth("", smtpUser, smtpPassword, smtpServer)
	return smtp.SendMail(smtpServer+":"+smtpPort, auth, smtpUser, []string{to}, []byte(body))
}

// funzione per salvare una nuova notifica o aggiornare l'ultima notifica di un utente
func saveNotification(email string, content string) error {
	client, err := OpenMongoDB()
	if err != nil {
		log.Println("Errore1: ", err)
		return err
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("notificationservice").Collection("notificheinviate")

	filter := bson.M{"EmailUtente": email}
	update := bson.M{
		"$set": bson.M{
			"EmailUtente":   email,
			"ElencoAziende": content,
		},
	}
	//inserimento della notifica se non esiste
	opts := options.Update().SetUpsert(true)

	result, err := collection.UpdateOne(context.Background(), filter, update, opts)
	log.Printf("UpdateOne risultato: MatchedCount=%v, ModifiedCount=%v, UpsertedCount=%v\n",
		result.MatchedCount, result.ModifiedCount, result.UpsertedCount)
	log.Println("Errore2: ", err)
	return err
}

// funzione di recupero email ed invio all'utente
func (t *NotificationService) Notify(dNotifica DatiNotifica, nResult *NotifyResult) error {
	client, err := OpenMongoDB()
	if err != nil {
		log.Print("Errore connessione DB:", err)
		*nResult = -1
		return err
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("notificationservice").Collection("notificheinviate")

	var result struct {
		ContenutoEmail string `bson:"ElencoAziende"`
	}

	//ricerca della notifica salvata per la mail fornita
	filter := bson.M{"EmailUtente": dNotifica.Email}
	log.Printf("Esecuzione query per EmailUtente: %s", dNotifica.Email)
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Println("Errore nella query:", err)
		*nResult = -2
		return err
	}

	log.Printf("Notifica trovata per l'utente %s. ContenutoEmail: %s", dNotifica.Email, result.ContenutoEmail)
	//creazione del corpo della email
	emailBody2 := "Subject: Elenco aziende Sdcc\n" +
		"Il destinatario è " + dNotifica.Email + "\n" +
		"L'elenco delle aziende è: " + result.ContenutoEmail
	//invio della email
	err = sendEmail(dNotifica.Email, emailBody2)
	if err != nil {
		log.Println("Errore invio mail:", err)
		*nResult = -3
	} else {
		*nResult = 1
	}
	return err
}

// funzione che ascolta la coda RabbitMQ e gestisce ogni messaggio ricevuto
func ConsumeRabbitMQMessages() {
	var conn *amqp.Connection
	var err error

	//tentativi di connessione a RabbitMQ
	maxRetries := 20
	for i := 1; i <= maxRetries; i++ {
		conn, err = amqp.Dial("amqp://myuser:mypassword@rabbitmq:5672/")
		if err == nil {
			log.Println("Connessione a RabbitMQ riuscita")
			break
		}

		log.Printf("Tentativo %d/%d fallito: %v", i, maxRetries, err)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Fatalf("🚫 Impossibile connettersi a RabbitMQ dopo %d tentativi: %v", maxRetries, err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Errore apertura canale RabbitMQ: %v", err)
	}
	defer ch.Close()

	//assicurazione che la coda esista
	q, err := ch.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Errore dichiarazione coda RabbitMQ: %v", err)
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Errore registrazione consumer RabbitMQ: %v", err)
	}

	//ciclo di elaborazione dei messaggi ricevuti
	for d := range msgs {
		var message RabbitMQMessage
		if err := json.Unmarshal(d.Body, &message); err != nil {
			log.Printf("Errore parsing messaggio: %v", err)
			continue
		}

		log.Printf("Messaggio ricevuto: %+v", message)

		//costruzione del corpo della email
		companiesString := strings.Join(message.Companies, ", ")
		emailBody := "Subject: Elenco aziende Sdcc\n" +
			"Il destinatario è " + message.Email + "\n" +
			"L'elenco delle aziende è: " + companiesString

		//invio della email
		err = sendEmail(message.Email, emailBody)
		if err != nil {
			log.Printf("Errore invio mail: %v", err)
			continue
		}
		log.Printf("Mail inviata correttamente a %s", message.Email)

		//salvataggio della email nel database
		err = saveNotification(message.Email, companiesString)
		if err != nil {
			log.Printf("Errore salvataggio notifica DB: %v", err)
		} else {
			log.Printf("Notifica salvata per %s", message.Email)
		}
	}
}
