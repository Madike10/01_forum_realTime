package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"realtimeforum/models"
	"time"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	switch r.Method {
	case "GET":
		userData, ok := GetSessionData(w, r, db)
		if !ok {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		response := models.ResponseUser{
			Id_user_id:     userData.UserID,
			Username_email: userData.Username,
			NameUser:       userData.Username,
			Message:        "success Login",
			OK:             true,
		}
		jsonResponse, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	case "POST":
		// Intansiation du websocket server

		var credential models.Credential
		err = json.NewDecoder(r.Body).Decode(&credential)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}

		fmt.Println(credential.Username_email, credential.Password)

		// Exécuter la requête SQL pour vérifier les informations d'identification
		var user models.User
		err = db.QueryRow("SELECT * FROM users WHERE name_user = ? OR mail_user = ? AND password_user = ?", credential.Username_email, credential.Username_email, credential.Password).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Age, &user.FirstName, &user.LastName, &user.Gender)
		if err != nil {
			if err == sql.ErrNoRows {
				// Aucun utilisateur trouvé avec ces informations d'identification
				http.Error(w, "Invalid username/email or password", http.StatusUnauthorized)
				return
			}
			// Autre erreur lors de l'exécution de la requête SQL
			fmt.Println(err)
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		SessionData := models.SessionData{
			UserID:      user.Id,
			Username:    user.Name,
			IsLogged:    true,
			CreatedTime: time.Now(),
		}
		fmt.Println(SessionData)
		// Créer un nouveau cookie
		CreateSession(w, r, SessionData, db)
		// Faire quelque chose comme une réponse JSON avec le statut OK
		response := models.ResponseUser{
			Id_user_id:     user.Id,
			Username_email: credential.Username_email,
			NameUser:       user.Name,
			Password:       credential.Password,
			Message:        "success Login",
			OK:             true,
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

// http.ListenAndServe(":8081", http.HandlerFunc(HandleWebsocket))
// HandleWebsocket(w, r)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	cookie, err := r.Cookie("session")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	DeleteSession(w, r, db, cookie.Value)
}
