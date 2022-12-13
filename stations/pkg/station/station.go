package station

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"sort"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Station struct {
	Name        string `json:"name"`
	MarkerColor string `json:"marker-color"`
	Zone        string `json:"zone"`
}

type Location struct {
	Coordinates []float64 `json:"coordinates"`
}

type TubeResponse struct {
	Tube     Station  `json:"properties"`
	Location Location `json:"geometry"`
}

type StationData struct {
	Stations []TubeResponse `json:"features"`
}

type StationLoc struct {
	Name      string
	Latitude  float64
	Longitude float64
}

type StationProxy struct {
	Name     string
	Distance float64
}

type Position struct {
	Latitude  float64
	Longitude float64
}

const radConv = math.Pi / 180.0

func initStations() []TubeResponse {
	var response StationData
	data, err := os.ReadFile("london_stations.json")
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(data, &response); err != nil {
		log.Fatal(err)
	}

	return response.Stations
}

func GetAllStations(flag bool) {
	response := initStations()

	if flag {
		for _, stations := range response {
			fmt.Printf("Name: %s | Zone: %s Marker-Color: %s\n", stations.Tube.Name, stations.Tube.Zone, stations.Tube.MarkerColor)
		}
	}

	if !flag {
		for _, stations := range response {
			fmt.Printf("Name: %s\n", stations.Tube.Name)
		}
	}
}

func GetStation(input string, flag bool) {
	response := initStations()
	sMap := stationsMap(response)

	caser := cases.Title(language.English)
	input = caser.String(input)

	v, ok := sMap[input]

	if ok {
		if flag {
			fmt.Printf("Name: %s | Zone: %s Marker-Color: %s\n", v.Name, v.Zone, v.MarkerColor)
		} else {
			fmt.Printf("Name: %s\n", v.Name)
		}
	}
	if !ok {
		fmt.Printf("%s doesn't exist, Please try again.", input)
	}

}

func GetClosestStations(p Position) []StationProxy {
	response := initStations()
	stations := getStationsForProxy(response)

	sp := make([]StationProxy, len(stations))
	for i, station := range stations {
		sp[i] = StationProxy{
			Name:     station.Name,
			Distance: convKmToMiles(distance(p, Position{station.Latitude, station.Longitude})),
		}
	}

	sort.SliceStable(sp, func(i, j int) bool {
		return sp[i].Distance < sp[j].Distance
	})

	return sp
}

// func caseSensitive(a, b string) bool {
// 	a, b = strings.ToUpper(a), strings.ToUpper(b)
// 	return strings.Contains(a, b)
// }

func getStationsForProxy(t []TubeResponse) (sp []StationLoc) {
	for _, stations := range t {
		s := StationLoc{
			Name:      stations.Tube.Name,
			Latitude:  stations.Location.Coordinates[1],
			Longitude: stations.Location.Coordinates[0],
		}
		sp = append(sp, s)
	}

	return sp
}

func distance(x, y Position) float64 {
	long := (y.Longitude - x.Longitude) * radConv
	lat := (y.Latitude - x.Latitude) * radConv
	a := math.Pow(math.Sin(lat/2.0), 2) + math.Cos(x.Latitude*radConv)*math.Cos(y.Latitude*radConv)*math.Pow(math.Sin(long/2.0), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return 6357 * c
}

func convKmToMiles(input float64) float64 {
	return input * 0.6214
}

func stationsMap(r []TubeResponse) map[string]Station {
	s := map[string]Station{}
	for _, station := range r {
		s[station.Tube.Name] = Station{
			Name:        station.Tube.Name,
			Zone:        station.Tube.Zone,
			MarkerColor: station.Tube.MarkerColor,
		}
	}
	return s
}
