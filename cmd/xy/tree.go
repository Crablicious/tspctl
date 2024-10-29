package xy

import (
	"github.com/Crablicious/tspctl/client"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func newTreeCmd() *cobra.Command {
	var rawUUID string
	var outputId string

	treeCmd := &cobra.Command{
		Use:     "tree",
		Aliases: []string{"tree"},
		Short:   "Get the XY tree",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := client.MustFromContext(cmd.Context())
			UUID, err := uuid.Parse(rawUUID)
			if err != nil {
				return err
			}
			params := client.TreeParameters{}
			body := client.GetXYTreeJSONRequestBody{Parameters: params}
			resp, err := cfg.Tspc.GetXYTreeWithResponse(cmd.Context(), UUID, outputId, body)
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
