package xy

import (
	"errors"

	"github.com/Crablicious/tspctl/client"
	"github.com/Crablicious/tspctl/cmd/cmdutil"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func newModelCmd() *cobra.Command {
	var rawUUID string
	var outputId string
	var items []int32
	var nbTimes int32
	var timerange cmdutil.TimeRange

	modelCmd := &cobra.Command{
		Use:   "model",
		Short: "Get the XY model",
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
			params := client.RequestedParameters{RequestedItems: items, RequestedTimerange: client.TimeRange(timerange)}
			body := client.GetXYJSONRequestBody{Parameters: params}
			resp, err := cfg.Tspc.GetXYWithResponse(cmd.Context(), UUID, outputId, body)
			if err != nil {
				return err
			}
			return client.Handle(resp.HTTPResponse, resp.JSON200, resp.Body, cfg)
		},
	}
	modelCmd.Flags().StringVarP(&rawUUID, "uuid", "u", "", "uuid of the experiment")
	modelCmd.MarkFlagRequired("uuid")
	modelCmd.Flags().StringVarP(&outputId, "outputid", "i", "", "id of the output provider")
	modelCmd.MarkFlagRequired("outputid")
	modelCmd.Flags().Int32SliceVarP(&items, "items", "t", nil, "the entryids or seriesids to request")
	modelCmd.MarkFlagRequired("items")
	modelCmd.Flags().VarP(&timerange, "timerange", "r", "the start and end of the timerange")
	modelCmd.Flags().Int32VarP(&nbTimes, "nbtimes", "n", 0, "the number of timestamps to be sampled (1-65336) in the timerange")
	return modelCmd
}

func init() {}
