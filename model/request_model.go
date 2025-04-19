package model

type RequestModel struct {
	Uid string `json:"targetUid"`
	Message string `json:"message"`
}