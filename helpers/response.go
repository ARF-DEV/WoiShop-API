package helpers

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ProductCategory struct {
	ProductID    int    `json:"product_id"`
	ProductName  string `json:"product_name"`
	ProductPrice int    `json:"product_price"`
	CategoryID   int    `json:"category_id,omitempty"`
	CategoryName string `json:"category_name"`
}

func SuccessResponseJSON(w http.ResponseWriter, msg string, data interface{}) {
	response := jsonResponse{Success: true, Message: msg, Data: data}
	responseByte, _ := json.Marshal(response)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseByte)
}

func ErrorResponseJSON(w http.ResponseWriter, msg string, statusCode int) {
	response := jsonResponse{
		Success: false,
		Message: msg,
		Data:    msg,
	}

	responseByte, _ := json.Marshal(response)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(responseByte)
}
