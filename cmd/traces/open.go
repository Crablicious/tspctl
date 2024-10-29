package traces

import (
	"github.com/Crablicious/tspctl/client"

	"github.com/spf13/cobra"
)

func newOpenCmd() *cobra.Command {
	var name string
	var typeID string
	var uri string

	openCmd := &cobra.Command{
		Use:   "open",
		Short: "Open a trace",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := client.MustFromContext(cmd.Context())
			params := client.TraceQueryParameters{
				Name:   &name,
				TypeID: &typeID,
				Uri:    uri,
			}
			body := client.PutTraceJSONRequestBody{
				Parameters: params,
			}
			resp, err := cfg.Tspc.PutTraceWithResponse(cmd.Context(), body)
			if err != nil {
				return err
			}
			return client.Handle(resp.HTTPResponse, resp.JSON200, resp.Body, cfg)
		},
	}
	openCmd.Flags().StringVarP(&uri, "uri", "u", "", "URI of the trace")
	openCmd.MarkFlagRequired("uri")
	openCmd.Flags().StringVarP(&typeID, "typeid", "t", "", `The trace type's ID, to force the use of a parser / disambiguate the trace type`)
	openCmd.Flags().StringVarP(&name, "name", "n", "", "The name of the trace in the server, to override the default name")
	return openCmd
}

func init() {}
