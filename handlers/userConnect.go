package handlers

// import (
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"realtimeforum/models"
// )

// func GetUsersInSession(w http.ResponseWriter, r *http.Request) {

// 	db, err := sql.Open("sqlite3", "./database.db")
// 	if err != nil {
// 		fmt.Println(err)
// 		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}
// 	defer db.Close()
// 	_, ok := GetSessionData(w, r, db)
// 	if !ok {
// 		http.Redirect(w, r, "/login", http.StatusFound)
// 		return
// 	}
// 	var users []models.User
// 	err = db.QueryRow("SELECT * FROM users").Scan(&users)
// 	if err != nil {
// 		fmt.Println(err)
// 		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}
// 	response := models.ResponseUser{
// 		Message: "GET users",
// 		Users:   users,
// 	}
// 	jsonResponse, _ := json.Marshal(response)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(jsonResponse)
// }
