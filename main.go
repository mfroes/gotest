package main

import (
	"log"
	"net/http"

	"github.com/mfroes/gotest/main/pkg"
	"github.com/mfroes/gotest/main/routes"
)

func main() {
	log.Print("Starting server at port 8090 ...\n")
	meta := pkg.GetMetadata()
	log.Print(meta.ToJSON())
	http.HandleFunc("/", routes.HelloWorld)
	http.HandleFunc("/status", routes.Status)
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal(err)
	}
}