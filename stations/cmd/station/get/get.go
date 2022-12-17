package get

import (
	"github.com/spf13/cobra"
)

var InfoFlag bool
var GetCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"Get", "GET"},
	Short:   "Use the commands retrieve stations",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

	},
}
