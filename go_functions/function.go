// Package p contains an HTTP Cloud Function.
package p

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"log"
)

type Event struct {
	EventClass 	string
	TerminalID	string 
	Notes		string
}

// HelloWorld prints the JSON encoded "message" field in the body
// of the request or "Hello, World!" if there isn't one.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	
	log.Println("Inside HelloWorld")

	
	var d struct {
		Message string `json:"message"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(w, "No Message")
		log.Println("No Message")
		return
	}
	if d.Message == "" {
		fmt.Fprint(w, "dMsg:" +r)
		log.Println(r)
		return
	}
	fmt.Fprint(w, html.EscapeString(d.Message))
}
