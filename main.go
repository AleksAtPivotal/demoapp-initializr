package main

import (
	"log"
	"net/http"
	"os"

	"github.com/alekssaul/demoapp-initializr/pkg/github"
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

	// Create a repo request
	newrepo := github.RepositoryRequest{
		Name:              "demo",
		GithubAccessToken: os.Getenv("DEMOAPP_INITIALIZR_GITHUBTOKEN"),
		AutoInit:          true,
	}

	// Setup Github Connection

	// Create repository
	bRepoCreate := r.Header.Get("REPOCREATE")
	if bRepoCreate != "" {

	}
	err := newrepo.CreateRepository()
	if err != nil {
		log.Printf("Failed to create repo: %s\n", newrepo.Name)
	}

	w.WriteHeader(http.StatusOK)

}
