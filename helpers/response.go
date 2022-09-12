package helpers

import (
	"encoding/json"
	"net/http"
)

type ResponseFormat struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func FormattedResponse(w http.ResponseWriter, data *ResponseFormat) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&data)

}


