package experiments

import (
	"github.com/Crablicious/tspctl/client"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func newDeleteCmd() *cobra.Command {
	var rawUUID string

	deleteCmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"del"},
		Short:   "Delete an experiment",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := client.MustFromContext(cmd.Context())
			UUID, err := uuid.Parse(rawUUID)
			if err != nil {
				return err
			}
			resp, err := cfg.Tspc.DeleteExperimentWithResponse(cmd.Context(), UUID)
			if err != nil {
				return err
			}
			return client.Handle(resp.HTTPResponse, resp.JSON200, resp.Body, cfg)
		},
	}
	deleteCmd.Flags().StringVarP(&rawUUID, "uuid", "u", "", "uuid of an experiment")
	deleteCmd.MarkFlagRequired("uuid")
	return deleteCmd
}

func init() {}
