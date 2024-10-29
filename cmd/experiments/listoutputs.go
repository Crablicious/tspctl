package experiments

import (
	"github.com/Crablicious/tspctl/client"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func newListOutputsCmd() *cobra.Command {
	var rawUUID string

	listOutputsCmd := &cobra.Command{
		Use:     "listoutputs",
		Aliases: []string{"listo", "lso"},
		Short:   "List all outputs of an experiment",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := client.MustFromContext(cmd.Context())
			UUID, err := uuid.Parse(rawUUID)
			if err != nil {
				return err
			}
			resp, err := cfg.Tspc.GetProvidersWithResponse(cmd.Context(), UUID)
			if err != nil {
				return err
			}
			var outputs []*client.DataProvider
			if resp.HTTPResponse.StatusCode == 200 {
				outputs = make([]*client.DataProvider, len(*resp.JSON200))
				for i, o := range *resp.JSON200 {
					outputs[i] = &o
				}
			}
			return client.HandleSlice(resp.HTTPResponse, outputs, resp.Body, cfg)
		},
	}
	listOutputsCmd.Flags().StringVarP(&rawUUID, "uuid", "u", "", "uuid of the experiment")
	listOutputsCmd.MarkFlagRequired("uuid")
	return listOutputsCmd
}

func init() {}
