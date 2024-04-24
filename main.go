package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"realtimeforum/database"
	"realtimeforum/handlers"
)

func main() {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	database.CreateDatabase(db)
	defer db.Close()

	fileServer := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.Handle("/utils/", http.StripPrefix("/utils/", http.FileServer(http.Dir("utils"))))
	http.HandleFunc("/", handlers.HomePageHandler)
	http.HandleFunc("/register-data", handlers.RegisterHandler)
	http.HandleFunc("/login-data", handlers.LoginHandler)
	http.HandleFunc("/posts-data", handlers.PostHandler)
	http.HandleFunc("/ws", handlers.Socket)

	// Démarrer le serveur sur le port 8080
	fmt.Println("Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
