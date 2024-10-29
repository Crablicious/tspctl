package datatree

import (
	"github.com/spf13/cobra"
)

var DataTreeCmd = &cobra.Command{
	Use:     "datatree",
	Aliases: []string{"dat"},
	Short:   "Commands to query data tree",
}

func init() {
	DataTreeCmd.AddCommand(newGetCmd())
}
