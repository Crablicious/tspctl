package timegraph

import (
	"github.com/Crablicious/tspctl/client"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// TODO: Find a trace to test this on.
func newTreeCmd() *cobra.Command {
	var rawUUID string
	var outputId string

	treeCmd := &cobra.Command{
		Use:     "tree",
		Aliases: []string{"gettree", "arr"},
		Short:   "Get the tree of visible entries of the timegraph",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := client.MustFromContext(cmd.Context())
			UUID, err := uuid.Parse(rawUUID)
			if err != nil {
				return err
			}
			params := client.TreeParameters{RequestedTimes: nil}
			body := client.GetTimeGraphTreeJSONRequestBody{Parameters: params}
			resp, err := cfg.Tspc.GetTimeGraphTreeWithResponse(cmd.Context(), UUID, outputId, body)
			if err != nil {
				return err
			}
			return client.Handle(resp.HTTPResponse, resp.JSON200, resp.Body, cfg)
		},
	}
	treeCmd.Flags().StringVarP(&rawUUID, "uuid", "u", "", "uuid of the experiment")
	treeCmd.MarkFlagRequired("uuid")
	treeCmd.Flags().StringVarP(&outputId, "outputid", "i", "", "id of the output provider")
	treeCmd.MarkFlagRequired("outputid")
	return treeCmd
}

func init() {}
