package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Catch error
func Catch(err error) {
	if err != nil {
		panic(err)
	}
}

// respondwithError return error message
func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondwithJSON(w, code, map[string]string{"message": msg})
}

// respondwithJSON write json response format
func RespondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
