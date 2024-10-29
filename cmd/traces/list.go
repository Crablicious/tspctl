package traces

import (
	"github.com/Crablicious/tspctl/client"

	"github.com/spf13/cobra"
)

func newListCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List all opened traces",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := client.MustFromContext(cmd.Context())
			resp, err := cfg.Tspc.GetTracesWithResponse(cmd.Context())
			if err != nil {
				return err
			}
			var traces []*client.Trace
			if resp.HTTPResponse.StatusCode == 200 {
				traces = make([]*client.Trace, len(*resp.JSON200))
				for i, t := range *resp.JSON200 {
					traces[i] = &t
				}
			}
			return client.HandleSlice(resp.HTTPResponse, traces, resp.Body, cfg)
		},
	}
}

func init() {}
