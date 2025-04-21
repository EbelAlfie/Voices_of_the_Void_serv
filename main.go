package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"votv.co/model"
	"votv.co/utils"
)

var upgrader = websocket.Upgrader { }
var sockets map[string]*websocket.Conn = make(map[string]*websocket.Conn)

func main() {
	server := http.NewServeMux()

	server.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("uid")
		wsConn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Print("WS connection error " + err.Error())
			return 
		}

		if sockets[id] != nil {
			delete(sockets, id)
		}

		fmt.Println("New Socket Registered " + id)
		sockets[id] = wsConn
	})

	server.HandleFunc("/postEvent", func(w http.ResponseWriter, r *http.Request) {
		request, err := utils.GetRequest(r)

		if err != nil {
			fmt.Println("Uid not found")
			return 
		}

		uid := request.Uid
		if socket := sockets[uid]; socket != nil {
			println("Sending Request")
			socket.WriteJSON(
				model.IncomingMessage {
					Type: "incoming.message",
					Message: request.Message,
				},
			)
		}
	
	})

	server.HandleFunc("/route-call", func(w http.ResponseWriter, r *http.Request) {
		request, err := utils.GetRequest(r)

		if err != nil {
			fmt.Println("Uid not found")
			return 
		}

		uid := request.Uid
		if socket := sockets[uid]; socket != nil {
			socket.WriteJSON(model.CreateCall {
				Type: "call.create",
			})
		}
	})

	addr := "0.0.0.0:3001"

	fmt.Println("Listening on " + addr)
	err := http.ListenAndServe(addr, server)
	if err != nil {
		fmt.Println("Listen error")
	}
}