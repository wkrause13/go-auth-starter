package middleware

import (
	"net/http"
)

func AddContentTypeHeader(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Header().Set("Content-Type", "application/json")
	// Fetch API Spec sends OPTIONS request for CORS. Need to handle that case
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		return
	}
	next(w, r)
}
