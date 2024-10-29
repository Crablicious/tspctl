package timegraph

import (
	"errors"
	"github.com/Crablicious/tspctl/client"
	"github.com/Crablicious/tspctl/cmd/cmdutil"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func newStatesCmd() *cobra.Command {
	var rawUUID string
	var outputId string
	var nbTimes int32
	var timerange cmdutil.TimeRange
	var items []int32

	statesCmd := &cobra.Command{
		Use:     "states",
		Aliases: []string{"getstates", "arr"},
		Short:   "Get the states of the timegraph",
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
			params := client.RequestedParameters{
				RequestedTimerange: client.TimeRange(timerange),
				RequestedItems:     items}
			body := client.GetStatesJSONRequestBody{Parameters: params}
			resp, err := cfg.Tspc.GetStatesWithResponse(cmd.Context(), UUID, outputId, body)
			if err != nil {
				return err
			}
			return client.Handle(resp.HTTPResponse, resp.JSON200, resp.Body, cfg)
		},
	}
	statesCmd.Flags().StringVarP(&rawUUID, "uuid", "u", "", "uuid of the experiment")
	statesCmd.MarkFlagRequired("uuid")
	statesCmd.Flags().StringVarP(&outputId, "outputid", "i", "", "id of the output provider")
	statesCmd.MarkFlagRequired("outputid")
	statesCmd.Flags().Int32SliceVarP(&items, "items", "t", nil, "the entryids to request")
	statesCmd.MarkFlagRequired("items")
	statesCmd.Flags().VarP(&timerange, "timerange", "r", "the start and end of the timerange")
	statesCmd.MarkFlagRequired("timerange")
	statesCmd.Flags().Int32VarP(&nbTimes, "nbtimes", "n", 0, "the number of timestamps to be sampled (1-65336) in the timerange")
	return statesCmd
}

func init() {}
