package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients []websocket.Conn
var websocketupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Socket(w http.ResponseWriter, r *http.Request) {

	conn, err := websocketupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}

	clients = append(clients, *conn)

	for {
		msgType, msg, errr := conn.ReadMessage()
		if errr != nil {
			return
		}
		// Affichage du message sur la cosole
		fmt.Printf("%s envoyer: %s\n", conn.RemoteAddr().String(), msg)
		// fmt.Println("Response", msg)

		for _, client := range clients {
			if err = client.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	}
}
