package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"realtimeforum/models"
)

func LikedHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var likeRequest models.LikeRequest
		err := json.NewDecoder(r.Body).Decode(&likeRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(likeRequest)
	}
}
