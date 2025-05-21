package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//Serve i file statici dalla cartella ./html tramite il percorso /static/
	fs := http.FileServer(http.Dir("./frontEnd/html"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	//Gestisce tutte le richieste alla root e serve i file HTML richiesti
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if path == "/" {
			path = "/entryPage.html"
		}
		http.ServeFile(w, r, "./frontEnd/html"+path)
	})

	//Avvio del server sulla porta 8085
	fmt.Println("Server running on port 8085/")
	if err := http.ListenAndServe(":8085", nil); err != nil {
		log.Fatal(err)
	}
}
