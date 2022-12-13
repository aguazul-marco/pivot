package get

import (
	"fmt"
	"log"
	"strconv"

	"github.com/aguazul-marco/pivot/stations/pkg/station"
	"github.com/spf13/cobra"
)

var ClosestCmd = &cobra.Command{
	Use:     "closest",
	Aliases: []string{"clst", "Closest"},
	Short:   "lists the top 5 closest stations based on the latitude and longitude submitted",
	Example: "get closest <lat> -- <lng>",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		lat, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			log.Fatalf("most use a numeric value: %v", err)
		}

		lng, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
			log.Fatalf("most use a numeric value: %v", err)
		}

		userInput := station.Position{
			Latitude:  lat,
			Longitude: lng,
		}

		closestStations := station.GetClosestStations(userInput)

		for _, station := range closestStations[:5] {
			fmt.Printf("Name: %s Distance: %v miles\n", station.Name, station.Distance)
		}

	},
}

func init() {
	GetCmd.AddCommand(ClosestCmd)
}
