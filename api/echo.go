package api

import (
	"fmt"
	"net/http"
)

func EchoHandleFunc(w http.ResponseWriter, r *http.Request) {
	echoMessage := r.URL.Query().Get("message")
	if len(echoMessage) == 0 {
		echoMessage = "No message was found in URL"
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, echoMessage)
}