package marvel

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Character struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CharacterResponse struct {
	Code            int    `json:"code"`
	Status          string `json:"status"`
	Copyright       string `json:"copyright"`
	AttributionText string `json:"attributionText"`
	AttributionHTML string `json:"attributionHTML"`
	Etag            string `json:"etag"`
	Data            struct {
		Offset  int         `json:"offset"`
		Limit   int         `json:"limit"`
		Total   int         `json:"total"`
		Count   int         `json:"count"`
		Results []Character `json:"results"`
	} `json:"data"`
}

type Client struct {
	BaseURL    string
	PubKey     string
	PrivKey    string
	HttpClient *http.Client
}

func NewClient(pubKey, privKey string) Client {
	return Client{
		BaseURL: "https://gateway.marvel.com/v1/public/",
		PubKey:  pubKey,
		PrivKey: privKey,
		HttpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) md5Hash(ts int64) string {
	tsForHash := strconv.Itoa(int(ts))
	hash := md5.Sum([]byte(tsForHash + c.PrivKey + c.PubKey))
	return hex.EncodeToString(hash[:])
}

func (c *Client) signURL(url string) string {
	ts := time.Now().Unix()
	hash := c.md5Hash(ts)
	return fmt.Sprintf("%s&ts=%d&apikey=%s&hash=%s", url, ts, c.PubKey, hash)
}

func (c *Client) GetCharacter(limit int) ([]Character, error) {
	url := c.BaseURL + fmt.Sprintf("/characters?limit=%d", limit)
	url = c.signURL(url)

	res, err := c.HttpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var CharacterResponse CharacterResponse
	if err := json.NewDecoder(res.Body).Decode(&CharacterResponse); err != nil {
		return nil, err
	}
	return CharacterResponse.Data.Results, nil
}
