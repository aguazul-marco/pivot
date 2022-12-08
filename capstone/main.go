package main

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const radConv = math.Pi / 180.0

type Tube struct {
	Name        string `json:"name"`
	MarkerColor string `json:"marker-color"`
	Zone        string `json:"zone"`
}

type Location struct {
	Coordinates []float64 `json:"coordinates"`
}

type TubeStation struct {
	Tube     Tube     `json:"properties"`
	Location Location `json:"geometry"`
}

type DataResponse struct {
	Stations []TubeStation `json:"features"`
}

type StationProxy struct {
	Station   Tube
	Longitude float64
	Latitude  float64
}

type Position struct {
	Longitude float64
	Latitude  float64
}

var response []TubeStation

func InitStations(path string) []TubeStation {
	var response DataResponse
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(data, &response); err != nil {
		log.Fatal(err)
	}

	return response.Stations

}

func GetStations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("error occured while encoding: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func ClosetProxy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func Distance(x, y Position) float64 {
	long := (y.Longitude - x.Longitude) * radConv
	lat := (y.Latitude - y.Latitude) * radConv
	a := math.Pow(math.Sin(lat/2.0), 2) + math.Cos(x.Latitude*radConv)*math.Cos(y.Latitude*radConv)*math.Pow(math.Sin(long/2.0), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return 6357 * c
}

func main() {

	response = InitStations("london_stations.json")

	r := mux.NewRouter()
	r.HandleFunc("/stations", GetStations).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
