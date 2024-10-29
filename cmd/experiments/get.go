package experiments

import (
	"github.com/Crablicious/tspctl/client"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func newGetCmd() *cobra.Command {
	var rawUUID string

	getCmd := &cobra.Command{
		Use:   "get",
		Short: "Get an experiment with a specific UUID",
		Long: `Get an experiment with a specific UUID.
If verbose is enabled, the row for the experiment will be printed followed by all its traces`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := client.MustFromContext(cmd.Context())
			UUID, err := uuid.Parse(rawUUID)
			if err != nil {
				return err
			}
			resp, err := cfg.Tspc.GetExperimentWithResponse(cmd.Context(), UUID)
			if err != nil {
				return err
			}
			return client.Handle(resp.HTTPResponse, resp.JSON200, resp.Body, cfg)
		},
	}
	getCmd.Flags().StringVarP(&rawUUID, "uuid", "u", "", "uuid of an experiment")
	getCmd.MarkFlagRequired("uuid")
	return getCmd
}

func init() {}
