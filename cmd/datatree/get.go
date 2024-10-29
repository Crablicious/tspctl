package datatree

import (
	"github.com/Crablicious/tspctl/client"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func newGetCmd() *cobra.Command {
	var rawUUID string
	var outputId string

	getCmd := &cobra.Command{
		Use:   "get",
		Short: "Get the visible entries of the datatree",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := client.MustFromContext(cmd.Context())
			UUID, err := uuid.Parse(rawUUID)
			if err != nil {
				return err
			}
			body := client.GetDataTreeJSONRequestBody{Parameters: client.TreeParameters{}}
			resp, err := cfg.Tspc.GetDataTreeWithResponse(cmd.Context(), UUID, outputId, body)
			if err != nil {
				return err
			}
			return client.Handle(resp.HTTPResponse, resp.JSON200, resp.Body, cfg)
		},
	}
	getCmd.Flags().StringVarP(&rawUUID, "uuid", "u", "", "uuid of the experiment")
	getCmd.MarkFlagRequired("uuid")
	getCmd.Flags().StringVarP(&outputId, "outputid", "i", "", "id of the output provider")
	getCmd.MarkFlagRequired("outputid")
	return getCmd
}

func init() {}
