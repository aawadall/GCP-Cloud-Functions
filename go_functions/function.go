// Package p contains an HTTP Cloud Function.
package p

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"log"
	"bytes"
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
	log.Println("Received HTTP Request: " + r.Method)
	log.Println("Protocol: "+r.Proto)
	body, err := r.GetBody()

	if err == nil {
		buf := new(bytes.Buffer)
		buf.ReadFrom(body)
		bdy := buf.String()
		log.Println("Body: " + bdy)
	}
	var d struct {
		Message string `json:"message"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(w, "No Message")
		log.Println("No Message")
		return
	}
	if d.Message == "" {
		fmt.Fprint(w, "dMsg empty" )
		
		return
	}
	log.Println(d.Message)
	fmt.Fprint(w, html.EscapeString(d.Message))
}
