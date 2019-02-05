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
		Notes: "TBC"
	}

	var d struct {
		Message string `json:"message"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(w, "ERR:" +r.Method)
		return
	}
	if d.Message == "" {
		w, err := json.Marshal(event)	
		//fmt.Fprint(w, "dMsg:"+r.Method)
		return
	}
	fmt.Fprint(w, html.EscapeString(d.Message))
}
