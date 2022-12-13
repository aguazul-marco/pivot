package stations

import (
	"fmt"
	"os"

	"github.com/aguazul-marco/pivot/stations/cmd/station/get"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stations",
	Short: "stations - a simple CLI of London's Tube System",
	Long: `stations can provide station information such as zone and maker color.
User can also us stations CLI to find out the closest station from a certain location using latitude and longitude.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "There was an error while executing the CLI '%s'", err)
		os.Exit(1)
	}

}

func addSubCommand() {
	rootCmd.AddCommand(get.GetCmd)
}

func init() {
	addSubCommand()
}
