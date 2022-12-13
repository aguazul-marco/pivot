package get

import (
	"github.com/aguazul-marco/pivot/stations/pkg/station"
	"github.com/spf13/cobra"
)

var AllCmd = &cobra.Command{
	Use:     "all",
	Aliases: []string{"All", "a"},
	Short:   "get a list of all stations in London",
	Run: func(cmd *cobra.Command, args []string) {
		station.GetAllStations(InfoFlag)
	},
}

func init() {
	AllCmd.Flags().BoolVarP(&InfoFlag, "info", "i", false, "station information (zone & marker-color)")
	GetCmd.AddCommand(AllCmd)
}
