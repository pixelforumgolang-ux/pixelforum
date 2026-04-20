
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	initDB()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "index.html")
			http.HandleFunc("/redirect", backGo.redirectHandler)
			http.HandleFunc("/", backGo.homeHandler)
			http.HandleFunc("/login", backGo.homeHandler2)
			http.HandleFunc("/aboutus", backGo.homeHandler3)
			http.HandleFunc("/subject", backGo.homeHandler4)
			http.HandleFunc("/admin", backGo.homeHandler4)
			
		})

	fmt.Printf("Serveur lancé sur le port %s...\n", port)


	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Erreur serveur : ", err)
	}
}
