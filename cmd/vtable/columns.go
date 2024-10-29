package vtable

import (
	"github.com/Crablicious/tspctl/client"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func newColumnsCmd() *cobra.Command {
	var rawUUID string
	var outputId string

	columnsCmd := &cobra.Command{
		Use:     "columns",
		Aliases: []string{"getcolumns", "getcol", "col"},
		Short:   "Get the column entries of the table",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := client.MustFromContext(cmd.Context())
			UUID, err := uuid.Parse(rawUUID)
			if err != nil {
				return err
			}
			params := client.OptionalParameters{}
			body := client.GetColumnsJSONRequestBody{Parameters: params}
			resp, err := cfg.Tspc.GetColumnsWithResponse(cmd.Context(), UUID, outputId, body)
			if err != nil {
				return err
			}
			return client.Handle(resp.HTTPResponse, resp.JSON200, resp.Body, cfg)
		},
	}
	columnsCmd.Flags().StringVarP(&rawUUID, "uuid", "u", "", "uuid of the experiment")
	columnsCmd.MarkFlagRequired("uuid")
	columnsCmd.Flags().StringVarP(&outputId, "outputid", "i", "", "id of the output provider")
	columnsCmd.MarkFlagRequired("outputid")
	return columnsCmd
}

func init() {}
