package main

import (
	"net/http"

	"github.com/gorilla/websocket"
	"votv.co/controller"
)

var upgrader = websocket.Upgrader{}

func main() {
	server := http.NewServeMux()

	server.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		controller.WsController(upgrader, w, r)
	})

	http.ListenAndServe("0.0.0.0:3001", server)
}