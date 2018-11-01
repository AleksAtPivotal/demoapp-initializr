package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	appVersion = "1.0"
)

var (
	Posturl string
	Token   string
)

// App defines app settings
type App struct {
	Posturl   string
	Token     string
	RevertURL string
}

// Payload used to construct the payload
type Payload struct {
	Repo string `json:"repo"`
}

func main() {
	setting := App{
		Posturl:   os.Getenv("POST_URL"),
		Token:     os.Getenv("TOKEN"),
		RevertURL: os.Getenv("REVERT_URL"),
	}

	if setting.Token == "" {
		setting.Token = "42"
	}

	if setting.RevertURL == "" {
		setting.RevertURL = "https://www.google.com"
	}

	http.HandleFunc("/", setting.handleRoot)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (setting *App) handleRoot(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("name")
	log.Printf("Creating Project Name: %s", url)
	log.Printf("Post URL: %s, Token: %s, Revert URL: %s", setting.Posturl, setting.Token, setting.RevertURL)

	p := Payload{
		Repo: url,
	}

	j, err := json.Marshal(p)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("Payload: %s", j)
	log.Printf("Post URL: %s", setting.Posturl+"?token"+setting.Token)
	req, err := http.NewRequest("POST", setting.Posturl+"?token"+setting.Token, bytes.NewBuffer(j))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	http.Redirect(w, r, setting.RevertURL, 301)
}
