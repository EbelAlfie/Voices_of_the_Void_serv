package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func main() {
	server := http.NewServeMux()

	server.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		wsConn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Print("WS connection error " + err.Error())
			return 
		}

		go func () {
			for {
				_, message, err := wsConn.ReadMessage()	
				if (err != nil) {
					fmt.Println("Error reading message " + err.Error())
					return 
				}
				fmt.Print(message)
			}
		}()
	})
}