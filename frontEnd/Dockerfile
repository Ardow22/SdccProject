#Usa un'immagine base di Go (Alpine è leggera e ideale per Go)
FROM golang:1.24-alpine

#Imposta la directory di lavoro all'interno del container
WORKDIR /app

#Copia i file di dipendenza go.mod e go.sum prima per evitare di ricostruire le dipendenze ogni volta
COPY go.mod go.sum ./

#Installa le dipendenze
RUN go mod tidy

#Copia tutto il codice sorgente nella directory di lavoro del container
COPY . .

#Compila il programma Go
RUN go build -o frontendServer ./frontEnd/main.go

#Espone la porta su cui il microservizio ascolta
EXPOSE 8085

#Avvia il programma compilato
CMD ["./frontendServer"]
