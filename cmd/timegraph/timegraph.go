package timegraph

import (
	"github.com/spf13/cobra"
)

var TimegraphCmd = &cobra.Command{
	Use:   "timegraph",
	Short: "Commands to query timegraph models",
}

func init() {
	TimegraphCmd.AddCommand(newArrowsCmd())
	TimegraphCmd.AddCommand(newStatesCmd())
	TimegraphCmd.AddCommand(newTooltipCmd())
	TimegraphCmd.AddCommand(newTreeCmd())
}
