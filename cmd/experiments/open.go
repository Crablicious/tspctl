package experiments

import (
	"github.com/Crablicious/tspctl/client"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func newOpenCmd() *cobra.Command {
	var name string
	var traces []string

	openCmd := &cobra.Command{
		Use:   "open",
		Short: "Open traces in an experiment",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := client.MustFromContext(cmd.Context())
			UUIDs := make([]uuid.UUID, len(traces))
			for i, rawUUID := range traces {
				UUID, err := uuid.Parse(rawUUID)
				if err != nil {
					return err
				}
				UUIDs[i] = UUID
			}
			params := client.ExperimentQueryParameters{
				Name:   name,
				Traces: UUIDs,
			}
			body := client.PostExperimentJSONRequestBody{
				Parameters: params,
			}
			resp, err := cfg.Tspc.PostExperimentWithResponse(cmd.Context(), body)
			if err != nil {
				return err
			}
			return client.Handle(resp.HTTPResponse, resp.JSON200, resp.Body, cfg)
		},
	}
	openCmd.Flags().StringSliceVarP(&traces, "traces", "t", nil, "uuids of the traces (-t <uuid>,<uuid> or -t <uuid> -t <uuid>)")
	openCmd.MarkFlagRequired("traces")
	openCmd.Flags().StringVarP(&name, "name", "n", "", "The name of the experiment")
	openCmd.MarkFlagRequired("name")
	return openCmd
}

func init() {}
