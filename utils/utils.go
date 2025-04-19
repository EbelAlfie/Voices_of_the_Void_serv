package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"votv.co/model"
)

func GetRequest(r *http.Request) (*model.RequestModel, error) {
	var request model.RequestModel
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&request); err != nil {
		fmt.Println("read error " + err.Error())
		return nil, err
	}
	return &request, nil
}