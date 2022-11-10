package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
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
		baseURL: "https://gateway.marvel.com/v1/public/",
		pubKey:  pubKey,
		privKey: privKey,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}

	events, err := client.getEvents(3)
	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(events)
}

type marvelClient struct {
	baseURL    string
	pubKey     string
	privKey    string
	httpClient *http.Client
}

func (c *marvelClient) md5Hash(ts int64) string {
	tsForHash := strconv.Itoa(int(ts))
	hash := md5.Sum([]byte(tsForHash + c.privKey + c.pubKey))
	return hex.EncodeToString(hash[:])
}

func (c *marvelClient) signURL(url string) string {
	ts := time.Now().Unix()
	hash := c.md5Hash(ts)
	return fmt.Sprintf("%s&ts=%d&apikey=%s&hash=%s", url, ts, c.pubKey, hash)
}

func (c *marvelClient) getEvents(limit int) ([]Event, error) {
	url := c.baseURL + fmt.Sprintf("/events?limit=%d", limit)
	url = c.signURL(url)

	res, err := c.httpClient.Get(url)
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
