package model

import "github.com/gorilla/websocket"

type SocketModel struct {
	Id string 
	Ws *websocket.Conn
}