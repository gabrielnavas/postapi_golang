package helpers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type responseHTTP struct {
	Message string      `json:"message"`
	Body    interface{} `json:"body"`
}

func Parse(w http.ResponseWriter, r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func GetVarsRoute(r *http.Request) map[string]string {
	return mux.Vars(r)
}

func SendResponse(w http.ResponseWriter, r *http.Request, message string, body interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(responseHTTP{
		Message: message,
		Body:    body,
	})
	if err != nil {
		log.Printf("Cannot format json. err=%v", err)
	}
}
