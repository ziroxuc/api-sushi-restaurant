package utils

import (
	"net/http"
	"encoding/json"
    "golang.org/x/crypto/bcrypt"

	"log"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}


func HashearPassword(password string) []byte{
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// TODO: Properly handle error
		log.Fatal(err)
	}
	return hash
}

func ComparePassword(hashFromDatabase []byte, password2 string) bool{
	if err := bcrypt.CompareHashAndPassword(hashFromDatabase, []byte(password2)); err != nil {
		// TODO: Properly handle error
		return false
	}
	return true
}