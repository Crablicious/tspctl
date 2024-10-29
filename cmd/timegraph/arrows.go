package timegraph

import (
	"errors"
	"github.com/Crablicious/tspctl/client"
	"github.com/Crablicious/tspctl/cmd/cmdutil"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// TODO: Find a trace to test this on.
func newArrowsCmd() *cobra.Command {
	var rawUUID string
	var outputId string
	var nbTimes int32
	var timerange cmdutil.TimeRange

	arrowsCmd := &cobra.Command{
		Use:     "arrows",
		Aliases: []string{"getarrows", "arr"},
		Short:   "Get the arrows of the timegraph",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if nbTimes < 0 {
				return errors.New("nbtimes should not be a negative number")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := client.MustFromContext(cmd.Context())
			UUID, err := uuid.Parse(rawUUID)
			if err != nil {
				return err
			}
			if nbTimes > 0 {
				timerange.NbTimes = &nbTimes
			}
			params := client.ArrowsParameters{client.TimeRange(timerange)}
			body := client.GetArrowsJSONRequestBody{Parameters: params}
			resp, err := cfg.Tspc.GetArrowsWithResponse(cmd.Context(), UUID, outputId, body)
			if err != nil {
				return err
			}
			return client.Handle(resp.HTTPResponse, resp.JSON200, resp.Body, cfg)
		},
	}
	arrowsCmd.Flags().StringVarP(&rawUUID, "uuid", "u", "", "uuid of the trace")
	arrowsCmd.MarkFlagRequired("uuid")
	arrowsCmd.Flags().StringVarP(&outputId, "outputid", "i", "", "id of the output provider")
	arrowsCmd.MarkFlagRequired("outputid")
	arrowsCmd.Flags().VarP(&timerange, "timerange", "r", "the start and end of the timerange")
	arrowsCmd.Flags().Int32VarP(&nbTimes, "nbtimes", "n", 0, "the number of timestamps to be sampled (1-65336) in the timerange")
	return arrowsCmd
}

func init() {}
