#utilizzo di un'immagine base di Go
FROM golang:1.24-alpine

#impostazione dela directory di lavoro all'interno del container
WORKDIR /app

#copio i file di dipendenza go.mod e go.sum prima per evitare di ricostruire le dipendenze ogni volta
COPY go.mod go.sum ./

#installazione delle dipendenze
RUN go mod tidy

#copio tutto il codice sorgente nella directory di lavoro del container
COPY . .

#compilazione del programma Go
RUN go build -o matchingServer ./matchingService/server.go

#esposizione della porta su cui il microservizio ascolta
EXPOSE 8081

#avvio del programma compilato
CMD ["./matchingServer"]
