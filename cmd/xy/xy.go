package xy

import (
	"github.com/spf13/cobra"
)

var XYCmd = &cobra.Command{
	Use:   "xy",
	Short: "Commands to query XY charts",
}

func init() {
	XYCmd.AddCommand(newTreeCmd())
	XYCmd.AddCommand(newModelCmd())
}
