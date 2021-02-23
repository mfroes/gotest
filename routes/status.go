package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mfroes/gotest/main/pkg"
)

// Status will return the application metadata
func Status(w http.ResponseWriter, req *http.Request) {
	log.Print("server: Status handl	er started")

	metadata := pkg.GetMetadata()
	stringMetadata := metadata.ToJSON()
	response := fmt.Sprintf(`{"myapplication":[%v]}`, stringMetadata)
	fmt.Fprintf(w, response)

	log.Print("server: Status handler ended")
}