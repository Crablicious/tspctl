package vtable

import (
	"github.com/Crablicious/tspctl/client"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func newLinesCmd() *cobra.Command {
	var rawUUID string
	var outputId string
	var tableCount int32
	var tableIndex int64
	var fromTime int64
	var columnIds []int64

	linesCmd := &cobra.Command{
		Use:     "lines",
		Aliases: []string{"getlines", "lin"},
		Short:   "Get the lines of the table",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := client.MustFromContext(cmd.Context())
			UUID, err := uuid.Parse(rawUUID)
			if err != nil {
				return err
			}
			var reqColumnIds *[]int64
			if columnIds != nil {
				reqColumnIds = &columnIds
			}
			var reqTableIndex *int64
			var reqTimes *[]int64
			if cmd.Flags().Changed("index") {
				reqTableIndex = &tableIndex
			} else { // mutually exclusive flags => we know time is set
				reqTimes = &[]int64{fromTime}
			}
			params := client.LinesParameters{
				RequestedTableColumnIds: reqColumnIds,
				RequestedTableCount:     tableCount,
				RequestedTableIndex:     reqTableIndex,
				RequestedTimes:          reqTimes,
				TableSearchDirection:    nil,
				TableSearchExpressions:  nil,
			}
			body := client.GetLinesJSONRequestBody{Parameters: params}
			resp, err := cfg.Tspc.GetLinesWithResponse(cmd.Context(), UUID, outputId, body)
			if err != nil {
				return err
			}
			return client.Handle(resp.HTTPResponse, resp.JSON200, resp.Body, cfg)
		},
	}
	linesCmd.Flags().StringVarP(&rawUUID, "uuid", "u", "", "uuid of the experiment")
	linesCmd.MarkFlagRequired("uuid")
	linesCmd.Flags().StringVarP(&outputId, "outputid", "i", "", "id of the output provider")
	linesCmd.MarkFlagRequired("outputid")
	linesCmd.Flags().Int32VarP(&tableCount, "count", "c", 32, "number of lines to request")
	linesCmd.Flags().Int64VarP(&tableIndex, "index", "x", 0, "starting index of requested lines")
	linesCmd.Flags().Int64VarP(&fromTime, "time", "t", 0, "starting time of requested lines")
	linesCmd.MarkFlagsMutuallyExclusive("index", "time")
	linesCmd.MarkFlagsOneRequired("index", "time")
	linesCmd.Flags().Int64SliceVarP(&columnIds, "columnIds", "l", nil, "the columns to request, not set means all")
	// table search expression + table search direction
	return linesCmd
}

func init() {}
