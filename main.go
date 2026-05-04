
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var templateDir string

func main() {

	initDB()

	staticDir := "./front"
	templateDir = "./front/page"
	if exePath, err := os.Executable(); err == nil {
		candidate := filepath.Join(filepath.Dir(exePath), "front")
		if _, statErr := os.Stat(candidate); statErr == nil {
			staticDir = candidate
			templateDir = filepath.Join(candidate, "page")
		}
	}

	fsStatic := http.FileServer(http.Dir(staticDir))
	http.Handle("/static/", http.StripPrefix("/static/", fsStatic))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "index.html")
			http.HandleFunc("/redirect", redirectHandler)
			http.HandleFunc("/", homeHandler)
			http.HandleFunc("/account", homeHandler2)
			http.HandleFunc("/aboutus", homeHandler3)
			http.HandleFunc("/subject", homeHandler4)
			http.HandleFunc("/admin", homeHandler4)
			
		})

	fmt.Printf("Serveur lancé sur le port %s...\n", port)


	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Erreur serveur : ", err)
	}
}
