package traces

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
		Short:   "Delete the trace with a specific UUID",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := client.MustFromContext(cmd.Context())
			UUID, err := uuid.Parse(rawUUID)
			if err != nil {
				return err
			}
			resp, err := cfg.Tspc.DeleteTraceWithResponse(cmd.Context(), UUID)
			if err != nil {
				return err
			}
			return client.Handle(resp.HTTPResponse, resp.JSON200, resp.Body, cfg)
		},
	}
	deleteCmd.Flags().StringVarP(&rawUUID, "uuid", "u", "", "uuid of the trace")
	deleteCmd.MarkFlagRequired("uuid")
	return deleteCmd
}

func init() {}
