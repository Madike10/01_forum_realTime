package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"realtimeforum/models"
	"time"

	"github.com/google/uuid"
)

type SessionData struct {
	UserID      int
	Username    string
	IsLoggedIn  bool
	CreatedTime time.Time
}

func CreateSession(w http.ResponseWriter, r *http.Request, data models.SessionData, db *sql.DB) {
	// Vérifiez si l'utilisateur a déjà une session active
	existingSessionID, err := r.Cookie("session")
	if err == nil {
		// L'utilisateur a une session active, mettez à jour la session existante
		newExpiration := time.Now().Add(5 * time.Hour) // Mise à jour avec votre durée de session désirée
		_, err := db.Exec("UPDATE sessions SET expiration = ? WHERE session_id = ? AND id_user = ?", newExpiration, existingSessionID.Value, data.UserID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Mettez à jour le cookie avec la nouvelle expiration
		cookie := http.Cookie{
			Name:     "session",
			Value:    existingSessionID.Value,
			Expires:  newExpiration,
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, &cookie)
		return
	}

	// L'utilisateur n'a pas de session active, créez une nouvelle session
	expiration := time.Now().Add(5 * time.Hour)
	data.CreatedTime = time.Now()
	token := uuid.NewString()
	dataJSON, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = db.Exec("INSERT INTO sessions (session_id, data, expiration, id_user, name_user) VALUES (?, ?, ?, ?, ?)", token, string(dataJSON), expiration, data.UserID, data.Username)
	cookie := http.Cookie{
		Name:     "session",
		Value:    token,
		Expires:  expiration,
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, &cookie)
}

func GetSessionData(w http.ResponseWriter, r *http.Request, db *sql.DB) (models.SessionData, bool) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return models.SessionData{}, false
	}
	var dataJSON string
	var expiration time.Time
	err = db.QueryRow("SELECT data, expiration FROM sessions WHERE session_id = ?", cookie.Value).Scan(&dataJSON, &expiration)
	if err != nil {
		// utils.Handle400Error(w)
		return models.SessionData{}, false
	}
	if time.Now().After(expiration) {
		_, err := db.Exec("DELETE FROM sessions WHERE session_id = ?", cookie.Value)
		if err != nil {
			http.Error(w, "404 Not Found", http.StatusNotFound)
		}
		cookie.Expires = time.Now().AddDate(0, 0, -1)
		http.SetCookie(w, cookie)
		return models.SessionData{}, false
	}
	var sessionData models.SessionData
	err = json.Unmarshal([]byte(dataJSON), &sessionData)
	if err != nil {
		return models.SessionData{}, false
	}
	return sessionData, true
}

func DeleteSession(w http.ResponseWriter, r *http.Request, db *sql.DB, sessionID string) {
	// Supprimez la session en se basant sur l'ID de session
	_, err := db.Exec("DELETE FROM sessions WHERE session_id = ?", sessionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Supprimez également le cookie côté client
	cookie := http.Cookie{
		Name:    "session",
		Value:   "",
		Expires: time.Now().AddDate(0, 0, -1),
		Path:    "/",
	}
	http.SetCookie(w, &cookie)
}

func IsValidSession(sessionData models.SessionData) bool {
	return time.Now().Before(sessionData.CreatedTime.Add(5 * time.Hour))
}
