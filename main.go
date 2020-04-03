package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Contact struct {
	Email string `json:"email"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Organization string `json:"organization"`
}

type Contacts struct {
	Contacts []Contact `json:"contacts"`
}

const baseURL = "https://lamppoststudios.api-us1.com/api/3/contacts"

var client = &http.Client{}

func GetContacts(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	activeCampaignKey := os.Getenv("ACTIVE_CAMPAIGN_KEY")

	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	req.Header.Add("Api-Token", activeCampaignKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	defer resp.Body.Close()

	var contacts Contacts

	if err := json.NewDecoder(resp.Body).Decode(&contacts); err != nil {
		log.Println(err)
	}

	if err := json.NewEncoder(w).Encode(contacts); err != nil {
		log.Println(err)
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/contacts", GetContacts)
	log.Fatal(http.ListenAndServe(":8080", router))
}
