package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"votv.co/model"
)

var upgrader = websocket.Upgrader { }
var sockets map[string]websocket.Conn = make(map[string]websocket.Conn)

func main() {
	server := http.NewServeMux()

	server.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		
		id := r.URL.Query().Get("uid")
		wsConn, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			fmt.Print("WS connection error " + err.Error())
			return 
		}

		newSocket := model.SocketModel{
			Id: id,
			Ws: wsConn,
		}

		sockets = append(sockets, newSocket)
	})

	server.HandleFunc("/respond", func(w http.ResponseWriter, r *http.Request) {
	
		sockets[0].Ws.WriteJSON(struct {
			From string `json:"from"`
		} {
			From: "Saya",
		})
	})

	server.HandleFunc("/route-call", func(w http.ResponseWriter, r *http.Request) {
		
		sockets[0].Ws.WriteJSON(struct {
			From string `json:"from"`
		} {
			From: "Saya",
		})
	})

	http.ListenAndServe("0.0.0.0:3001", server)
}