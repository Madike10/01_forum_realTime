package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"realtimeforum/database"
	"realtimeforum/models"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	// specifier la route avec avec la bonne methode
	// if r.URL.Path != "/register" {
	// 	http.Error(w, "404 Not Found", http.StatusNotFound)
	// 	return
	// }
	// if r.Method != "POST" {
	// 	http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
	// 	return
	// }
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(user.Name, user.Age, user.Email, user.Password, user.FirstName, user.LastName, user.Email, user.Gender)

	err = database.AddUser("./database.db", user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := map[string]string{
		"message": "Registration successful",
	}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
