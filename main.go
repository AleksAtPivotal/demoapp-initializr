package main

import (
	"log"
	"net/http"

	"github.com/concourse/go-concourse/concourse"
)

const (
	appVersion = "0.1"
)

func main() {
	log.Printf("Started the Application... Version %v\n", appVersion)
	http.HandleFunc("/", handleRoot)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	log.Printf("Recieved a request %v", r)

	w.WriteHeader(http.StatusOK)
	team := concourse.Team{}

}
