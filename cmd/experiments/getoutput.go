package experiments

import (
	"github.com/Crablicious/tspctl/client"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func newGetOutputCmd() *cobra.Command {
	var rawUUID string
	var outputId string

	getOutputCmd := &cobra.Command{
		Use:     "getoutput",
		Aliases: []string{"geto"},
		Short:   "Get a specific output descriptor for an experiment",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := client.MustFromContext(cmd.Context())
			UUID, err := uuid.Parse(rawUUID)
			if err != nil {
				return err
			}
			resp, err := cfg.Tspc.GetProviderWithResponse(cmd.Context(), UUID, outputId)
			if err != nil {
				return err
			}
			return client.Handle(resp.HTTPResponse, resp.JSON200, resp.Body, cfg)
		},
	}
	getOutputCmd.Flags().StringVarP(&rawUUID, "uuid", "u", "", "uuid of the experiment")
	getOutputCmd.MarkFlagRequired("uuid")
	getOutputCmd.Flags().StringVarP(&outputId, "outputid", "i", "", "id of the output provider")
	getOutputCmd.MarkFlagRequired("outputid")
	return getOutputCmd
}

func init() {}
