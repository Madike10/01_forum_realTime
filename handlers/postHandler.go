package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"realtimeforum/database"
	"realtimeforum/models"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		db, err := sql.Open("sqlite3", "./database.db")
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		defer db.Close()
		userData, ok := GetSessionData(w, r, db)
		if !ok {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		posts, err := database.GetAllPost("./database.db")
		if err != nil {
			fmt.Println("Erreur de recuperation des post")
			return
		}
		
		response := models.ResponsePost{
			Message:  "GET post",
			Post:     posts,
			UserData: userData,
		}
		jsonResponse, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)

	case "POST":
		var sessionData models.SessionData
		var post models.Post
		var posts []models.Post
		err := json.NewDecoder(r.Body).Decode(&post)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// fmt.Println(post.Content, post.Title, post.Date)

		err = database.AddPost("./database.db", post)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response := models.ResponsePost{
			Message:  "Post added",
			Post:     posts,
			UserData: sessionData,
		}
		jsonResponse, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
