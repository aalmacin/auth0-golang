package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// Response Used to marshal the json respons
type Response struct {
	Message string `json:"message"`
}

func main() {
	r := mux.NewRouter()

	// This route is always accessible
	r.Handle("/api/public", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		message := "Hello from a public endpoint! You don't need to be authenticated to see this."
		responseJSON(message, w, http.StatusOK)
	}))

	// This route is only accessible if the user has a valid Access Token
	// We are chaining the jwtmiddleware middleware into the negroni handler function which will check
	// for a valid token.
	r.Handle("/api/private", negroni.New(
		negroni.HandlerFunc(JwtMiddleware().HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			message := "Hello from a private endpoint! You need to be authenticated to see this."
			responseJSON(message, w, http.StatusOK)
		}))))

	http.ListenAndServe(":3441", r)
}

func responseJSON(message string, w http.ResponseWriter, statusCode int) {
	response := Response{message}
	jsonResponse, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}
