package handler

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type language struct {
	ID      uint   `json:"id"`
	Message string `json:"message"`
}

func HelloWorld(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	responseStruct := &language{ID: 1, Message: "hello world"}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(responseStruct); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	return
}
