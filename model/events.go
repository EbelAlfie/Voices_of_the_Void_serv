package model

type IncomingMessage struct {
	Type string `json:"type"`
	Message string `json:"message"`
}