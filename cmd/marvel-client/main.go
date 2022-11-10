package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	pubKey := os.Getenv("MARVEL_PUBLIC_KEY")
	privKey := os.Getenv("MARVEL_PRIVATE_KEY")

	client := marvelClient{
		pubKey:  pubKey,
		privKey: privKey,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}

	events, err := client.getEvents()
	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(events)
}

type marvelClient struct {
	pubKey     string
	privKey    string
	httpClient *http.Client
}

func (c *marvelClient) getEvents() ([]Event, error) {
	res, err := c.httpClient.Get("https://gateway.marvel.com/v1/public/events")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var EventResponse EventResponse
	if err := json.NewDecoder(res.Body).Decode(&EventResponse); err != nil {
		return nil, err
	}

	return EventResponse.Data.Results, nil
}
