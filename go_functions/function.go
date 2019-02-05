// Package p contains an HTTP Cloud Function.
package p

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

// HelloWorld prints the JSON encoded "message" field in the body
// of the request or "Hello, World!" if there isn't one.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	
	type Event struct {
		EventClass 	string
		TerminalID	string 
		Notes		string
	}

	event := Event{
		EventClass: "cash_swap",
		TerminalID: "ATM0000",
		Notes: "TBC",
	}

	var d struct {
		Message string `json:"message"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(w, "ERR:" +r.Method)
		return
	}
	if d.Message == "" {
		pl, err := json.Marshal(event)	
		if err != nil {
			fmt.Fprint(w, pl)
		} else {
			fmt.Fprint(w, "DEFAULT")
		}
		return
	}
	fmt.Fprint(w, html.EscapeString(d.Message))
}
