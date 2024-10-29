package cmd

import (
	"github.com/Crablicious/tspctl/client"

	"github.com/spf13/cobra"
)

func newHealthCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "health",
		Short: "Get health status of the server",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := client.MustFromContext(cmd.Context())
			resp, err := cfg.Tspc.GetHealthStatusWithResponse(cmd.Context())
			if err != nil {
				return err
			}
			return client.Handle(resp.HTTPResponse, resp.JSON200, resp.Body, cfg)
		},
	}
}

func init() {}
