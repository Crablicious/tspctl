package traces

import (
	"github.com/Crablicious/tspctl/client"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func newGetCmd() *cobra.Command {
	var rawUUID string

	getCmd := &cobra.Command{
		Use:   "get",
		Short: "Get the trace with a specific UUID",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := client.MustFromContext(cmd.Context())
			UUID, err := uuid.Parse(rawUUID)
			if err != nil {
				return err
			}
			resp, err := cfg.Tspc.GetTraceWithResponse(cmd.Context(), UUID)
			if err != nil {
				return err
			}
			return client.Handle(resp.HTTPResponse, resp.JSON200, resp.Body, cfg)
		},
	}
	getCmd.Flags().StringVarP(&rawUUID, "uuid", "u", "", "uuid of the trace")
	getCmd.MarkFlagRequired("uuid")
	return getCmd
}

func init() {}
