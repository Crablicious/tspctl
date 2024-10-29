package traces

import (
	"github.com/spf13/cobra"
)

var TracesCmd = &cobra.Command{
	Use:     "traces",
	Aliases: []string{"tra"},
	Short:   "Commands to manage traces",
}

func init() {
	TracesCmd.AddCommand(newListCmd())
	TracesCmd.AddCommand(newGetCmd())
	TracesCmd.AddCommand(newOpenCmd())
	TracesCmd.AddCommand(newDeleteCmd())
}
