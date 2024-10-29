package experiments

import (
	"github.com/spf13/cobra"
)

var ExperimentsCmd = &cobra.Command{
	Use:     "experiments",
	Aliases: []string{"exp"},
	Short:   "Commands to manage experiments",
}

func init() {
	ExperimentsCmd.AddCommand(newListCmd())
	ExperimentsCmd.AddCommand(newGetCmd())
	ExperimentsCmd.AddCommand(newOpenCmd())
	ExperimentsCmd.AddCommand(newDeleteCmd())
	ExperimentsCmd.AddCommand(newGetOutputCmd())
	ExperimentsCmd.AddCommand(newListOutputsCmd())
}
