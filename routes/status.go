package routes

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Status will return the application metadata
func Status(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	log.Print("server: hello handler started")
	defer log.Print("server: hello handler ended")

	select {
	case <-time.After(1 * time.Second):
		fmt.Fprintf(w, "Hello World")
	case <-ctx.Done():

		err := ctx.Err()
		log.Print("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}

}