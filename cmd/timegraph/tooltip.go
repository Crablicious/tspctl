package timegraph

import (
	"github.com/Crablicious/tspctl/client"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// TODO: Find a trace to test this on.
func newTooltipCmd() *cobra.Command {
	var rawUUID string
	var outputId string
	var tstamp int64
	var item int32
	var elemTime int64
	var elemType string
	var elemDuration int64
	var elemEntryId int64
	var elemDestId int64

	tooltipCmd := &cobra.Command{
		Use:     "tooltip",
		Aliases: []string{"gettooltip", "arr"},
		Short:   "Get the tooltip of the timegraph",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := client.MustFromContext(cmd.Context())
			UUID, err := uuid.Parse(rawUUID)
			if err != nil {
				return err
			}
			var elemEntryIdP *int64
			if cmd.Flags().Changed("elementryid") {
				elemEntryIdP = &elemEntryId
			}
			var elemDestIdP *int64
			if cmd.Flags().Changed("elemdestid") {
				elemDestIdP = &elemDestId
			}
			params := client.TooltipParameters{
				RequestedElement: client.Element{Time: elemTime, ElementType: client.ElementElementType(elemType),
					Duration: elemDuration, EntryId: elemEntryIdP, DestinationId: elemDestIdP},
				RequestedItems: []int32{item},
				RequestedTimes: []int64{tstamp}}
			body := client.GetTimeGraphTooltipJSONRequestBody{Parameters: params}
			resp, err := cfg.Tspc.GetTimeGraphTooltipWithResponse(cmd.Context(), UUID, outputId, body)
			if err != nil {
				return err
			}
			return client.Handle(resp.HTTPResponse, resp.JSON200, resp.Body, cfg)
		},
	}
	tooltipCmd.Flags().StringVarP(&rawUUID, "uuid", "u", "", "uuid of the experiment")
	tooltipCmd.MarkFlagRequired("uuid")
	tooltipCmd.Flags().StringVarP(&outputId, "outputid", "i", "", "id of the output provider")
	tooltipCmd.MarkFlagRequired("outputid")
	tooltipCmd.Flags().Int32VarP(&item, "item", "t", 0, "the entryid to request")
	tooltipCmd.MarkFlagRequired("item")
	tooltipCmd.Flags().Int64VarP(&tstamp, "time", "r", 0, "the timestamp to fetch tooltip at")
	tooltipCmd.MarkFlagRequired("time")

	tooltipCmd.Flags().Int64VarP(&elemTime, "elemtime", "e", 0, "the start time of the element")
	tooltipCmd.MarkFlagRequired("elemtime")
	tooltipCmd.Flags().StringVarP(&elemType, "elemtype", "y", "", "the type of the element, allowed values: STATE/ANNOTATION/ARROW")
	tooltipCmd.MarkFlagRequired("elemtype")
	tooltipCmd.Flags().Int64VarP(&elemDuration, "elemduration", "d", 0, "the duration of the element")
	tooltipCmd.MarkFlagRequired("elemduration")
	tooltipCmd.Flags().Int64VarP(&elemEntryId, "elementryid", "n", 0, "the ID of the entry (annotation/arrow)")
	tooltipCmd.Flags().Int64VarP(&elemDestId, "elemdestid", "s", 0, "destination entry's unique ID (arrow)")

	return tooltipCmd
}

func init() {}
