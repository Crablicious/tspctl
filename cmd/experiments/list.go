package experiments

import (
	"github.com/Crablicious/tspctl/client"

	"github.com/spf13/cobra"
)

func newListCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List all opened experiments",
		Long: `List all opened experiments.
If verbose is enabled, the row for each experiment will be printed followed by all its traces`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := client.MustFromContext(cmd.Context())
			resp, err := cfg.Tspc.GetExperimentsWithResponse(cmd.Context())
			if err != nil {
				return err
			}
			var exps []*client.Experiment
			if resp.HTTPResponse.StatusCode == 200 {
				exps = make([]*client.Experiment, len(*resp.JSON200))
				for i, e := range *resp.JSON200 {
					exps[i] = &e
				}
			}
			return client.HandleSlice(resp.HTTPResponse, exps, resp.Body, cfg)
		},
	}
}

func init() {}
