package respond

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Msg
type Msg struct {
	Error interface{} `json:"error"`
	Data  interface{} `json:"data"`
}

// DataMsg
type DataMsg struct {
	Data interface{} `json:"data"`
}

// ErrorMsg
type ErrorMsg struct {
	Error interface{} `json:"error"`
}

// With sends out a formated response
func With(w http.ResponseWriter, r *http.Request, status int, data interface{}, err interface{}) {
	var message interface{}

	if data != nil {
		message = &DataMsg{Data: data}
	}
	if err != nil {
		message = &ErrorMsg{Error: err}
	}
	if data != nil && err != nil {
		message = &Msg{Data: data, Error: err}
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(message); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	if _, err := io.Copy(w, &buf); err != nil {
		log.Println("respond:", err)
	}
}
