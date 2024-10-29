package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/Crablicious/tspctl/client"
	"github.com/Crablicious/tspctl/cmd/datatree"
	"github.com/Crablicious/tspctl/cmd/experiments"
	"github.com/Crablicious/tspctl/cmd/timegraph"
	"github.com/Crablicious/tspctl/cmd/traces"
	"github.com/Crablicious/tspctl/cmd/vtable"
	"github.com/Crablicious/tspctl/cmd/xy"

	"github.com/spf13/cobra"
)

func setAcceptHeader(c *client.Client) error {
	c.RequestEditors = append(c.RequestEditors,
		func(ctx context.Context, req *http.Request) error {
			req.Header.Add("Accept", "application/json")
			return nil
		})
	return nil
}

var (
	rawUrl     string
	outputType client.OutputType = client.OutputTypeTable
	verbose    bool
	fail       bool

	rootCmd = &cobra.Command{
		Use:   "tspctl",
		Short: "tspctl is a trace-server-protocol CLI",
		Long:  `A command-line tool to query a trace-server using the trace-server-protocol (TSP).`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			tspc, err := client.NewClientWithResponses(rawUrl, setAcceptHeader)
			if err != nil {
				return err
			}
			ctx := client.NewContext(context.Background(), &client.GlobalConfig{
				Verbose: verbose, Tspc: tspc,
				OutputType: outputType, FailOnHTTPErr: fail})
			cmd.SetContext(ctx)
			return nil
		},
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&rawUrl, "url", "http://localhost:8080/tsp/api", "url to a trace-server")

	rootCmd.PersistentFlags().VarP(&outputType, "output", "o", fmt.Sprintf("type of output, one of %v", client.OutputTypes))
	rootCmd.RegisterFlagCompletionFunc("output", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return client.OutputTypes, cobra.ShellCompDirectiveDefault
	})
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().BoolVar(&fail, "fail", false, "fail immediately after printing on non-success HTTP status codes (<200 or >299)")

	rootCmd.AddCommand(traces.TracesCmd)
	rootCmd.AddCommand(experiments.ExperimentsCmd)
	rootCmd.AddCommand(newHealthCmd())
	rootCmd.AddCommand(datatree.DataTreeCmd)
	rootCmd.AddCommand(vtable.VtableCmd)
	rootCmd.AddCommand(xy.XYCmd)
	rootCmd.AddCommand(timegraph.TimegraphCmd)
}
