package stations

import (
	"fmt"
	"os"

	"github.com/aguazul-marco/pivot/stations/cmd/station/get"
	"github.com/spf13/cobra"
)

var version = "0.0.1"

var rootCmd = &cobra.Command{
	Use:     "stations",
	Version: version,
	Short:   "'stations' is a simple London Underground Tube system CLI help desk\n",
	Long: `Using the command and subcommands the user is able tp retrieve the list of all the tube stations or a specific station. Using the flag will also retrieve station information such as zone and marker-color.
User can also find out the closest station from a certain location using latitude and longitude.`,
	Run: func(cmd *cobra.Command, args []string) {},
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
