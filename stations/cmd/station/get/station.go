package get

import (
	"github.com/aguazul-marco/pivot/stations/pkg/station"
	"github.com/spf13/cobra"
)

var StationCmd = &cobra.Command{
	Use:     "station",
	Aliases: []string{"Station", "stat"},
	Short:   "searching for station using the name or part of the name",
	Args:    cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]
		if len(args) == 2 {
			input = args[0] + " " + args[1]
		}
		station.GetStation(input, InfoFlag)
	},
}

func init() {
	StationCmd.Flags().BoolVarP(&InfoFlag, "info", "i", false, "station information (zone & marker-color)")
	GetCmd.AddCommand(StationCmd)
}
