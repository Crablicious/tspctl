package vtable

import (
	"github.com/spf13/cobra"
)

var VtableCmd = &cobra.Command{
	Use:     "vtable",
	Aliases: []string{"vtbl", "vta"},
	Short:   "Query virtual tables, e.g. events table",
}

func init() {
	VtableCmd.AddCommand(newColumnsCmd())
	VtableCmd.AddCommand(newLinesCmd())
}
