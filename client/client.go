package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"text/tabwriter"
)

func newWriter() *tabwriter.Writer {
	const padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)
	return w
}

func printResponseHeader(r *http.Response) {
	fmt.Printf("%s %s\n", r.Status, r.Proto)
}

func printJSONIndent(b []byte) error {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", " ")
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", &out)
	return nil
}

type TabPrinter interface {
	// PrintHeader prints the tabwriter header to w. The pointer
	// receiver must be allowed to be nil.
	PrintHeader(w *tabwriter.Writer, verbose bool)
	// Print prints the data rows. The format must match what is
	// produced by PrintHeader for the same verbosity.
	Print(w *tabwriter.Writer, verbose bool)
}

func Handle[T TabPrinter](r *http.Response, printer T, body []byte, cfg *GlobalConfig) error {
	return HandleSlice(r, []T{printer}, body, cfg)
}

func HandleSlice[T TabPrinter](r *http.Response, printers []T, body []byte, cfg *GlobalConfig) error {
	if cfg.FailOnHTTPErr && (r.StatusCode < 200 || r.StatusCode > 299) {
		os.Exit(1)
	}
	switch cfg.OutputType {
	case OutputTypeTable:
		if r.StatusCode != 200 {
			printResponseHeader(r)
			fmt.Println()
			fmt.Println(string(body))
			return nil
		}
		printResponseHeader(r)
		fmt.Println()
		w := newWriter()
		var nilT T
		nilT.PrintHeader(w, cfg.Verbose)
		for _, p := range printers {
			p.Print(w, cfg.Verbose)
		}
		w.Flush()
	case OutputTypeJSON:
		if r.StatusCode != 200 {
			fmt.Println(string(body)) // not guaranteed to be JSON
			return nil
		} else {
			return printJSONIndent(body)
		}
	default:
		return NewUnsupportedError(cfg.OutputType)
	}
	return nil
}

func (t *Trace) PrintHeader(w *tabwriter.Writer, verbose bool) {
	if !verbose {
		fmt.Fprintln(w, "Name\tNbEvents\tIndexingStatus\tUUID\t")
	} else {
		fmt.Fprintln(w, "Name\tStart\tEnd\tNbEvents\tIndexingStatus\tUUID\tPath\t")
	}
}

func (t *Trace) Print(w *tabwriter.Writer, verbose bool) {
	if t.Name != nil {
		fmt.Fprintf(w, "%s", *t.Name)
	}
	fmt.Fprint(w, "\t")
	if verbose {
		if t.Start != nil {
			fmt.Fprintf(w, "%d", *t.Start)
		}
		fmt.Fprint(w, "\t")
		if t.End != nil {
			fmt.Fprintf(w, "%d", *t.End)
		}
		fmt.Fprint(w, "\t")
	}
	if t.NbEvents != nil {
		fmt.Fprintf(w, "%d", *t.NbEvents)
	}
	fmt.Fprint(w, "\t")
	if t.IndexingStatus != nil {
		fmt.Fprintf(w, "%s", *t.IndexingStatus)
	}
	fmt.Fprint(w, "\t")
	if t.UUID != nil {
		fmt.Fprintf(w, "%s", t.UUID)
	}
	fmt.Fprint(w, "\t")
	if verbose {
		if t.Path != nil {
			fmt.Fprintf(w, "%s", *t.Path)
		}
		fmt.Fprint(w, "\t")
	}
	fmt.Fprint(w, "\n")
}

func (e *Experiment) PrintHeader(w *tabwriter.Writer, verbose bool) {
	if !verbose {
		fmt.Fprintln(w, "Name\tNbEvents\tIndexingStatus\tUUID\tTraceUUIDs\t")
	} else {
		fmt.Fprintln(w, "Type\tName\tStart\tEnd\tNbEvents\tIndexingStatus\tUUID\tTraceUUIDs/Path\t")
	}
}

func (e *Experiment) Print(w *tabwriter.Writer, verbose bool) {
	if verbose {
		fmt.Fprintf(w, "exp\t")
	}
	if e.Name != nil {
		fmt.Fprintf(w, "%s", *e.Name)
	}
	fmt.Fprint(w, "\t")
	if verbose {
		if e.Start != nil {
			fmt.Fprintf(w, "%d", *e.Start)
		}
		fmt.Fprint(w, "\t")
		if e.End != nil {
			fmt.Fprintf(w, "%d", *e.End)
		}
		fmt.Fprint(w, "\t")
	}
	if e.NbEvents != nil {
		fmt.Fprintf(w, "%d", *e.NbEvents)
	}
	fmt.Fprint(w, "\t")
	if e.IndexingStatus != nil {
		fmt.Fprintf(w, "%s", *e.IndexingStatus)
	}
	fmt.Fprint(w, "\t")
	if e.UUID != nil {
		fmt.Fprintf(w, "%s", *e.UUID)
	}
	fmt.Fprint(w, "\t")
	if e.Traces != nil {
		traceUUIDs := make([]string, len(*e.Traces))
		for i, t := range *e.Traces {
			traceUUIDs[i] = t.UUID.String()
		}
		fmt.Fprintf(w, "%s", strings.Join(traceUUIDs, ","))
	}
	fmt.Fprint(w, "\t")
	fmt.Fprintln(w)
	if !verbose || e.Traces == nil {
		return
	}
	for _, t := range *e.Traces {
		fmt.Fprintf(w, "trc\t")
		t.Print(w, verbose)
	}
}

func (d *DataProvider) PrintHeader(w *tabwriter.Writer, verbose bool) {
	fmt.Fprintf(w, "Name\tType\tId\t")
	if verbose {
		fmt.Fprintf(w, "Description\t")
	}
	fmt.Fprintln(w)
}

func (d *DataProvider) Print(w *tabwriter.Writer, verbose bool) {
	if d.Name != nil {
		fmt.Fprintf(w, "%s", *d.Name)
	}
	fmt.Fprint(w, "\t")
	if d.Type != nil {
		fmt.Fprintf(w, "%s", *d.Type)
	}
	fmt.Fprint(w, "\t")
	if d.Id != nil {
		fmt.Fprintf(w, "%s", *d.Id)
	}
	fmt.Fprint(w, "\t")
	if verbose {
		if d.Description != nil {
			fmt.Fprintf(w, "%s", *d.Description)
		}
		fmt.Fprint(w, "\t")
	}
	fmt.Fprintln(w)
}

func (s *ServerStatus) PrintHeader(w *tabwriter.Writer, verbose bool) {
	fmt.Fprintln(w, "Status\t")
}

func (s *ServerStatus) Print(w *tabwriter.Writer, verbose bool) {
	if s.Status != nil {
		fmt.Fprintf(w, "%s", *s.Status)
	}
	fmt.Fprint(w, "\t")
	fmt.Fprintln(w)
}

func (r *XYTreeResponse) PrintHeader(w *tabwriter.Writer, verbose bool) {
	fmt.Fprintln(w, "Status\tMessage\t")
}

func (r *XYTreeResponse) Print(w *tabwriter.Writer, verbose bool) {
	if r.Status != nil {
		fmt.Fprintf(w, "%s", *r.Status)
	}
	fmt.Fprint(w, "\t")
	if r.StatusMessage != nil {
		fmt.Fprintf(w, "%s", *r.StatusMessage)
	}
	fmt.Fprint(w, "\t")
	fmt.Fprintln(w)

	if r.Model == nil {
		return
	}
	fmt.Fprintln(w) // spacing

	if verbose {
		fmt.Fprint(w, "Id\tParentId\t")
	}
	for _, h := range *r.Model.Headers {
		fmt.Fprintf(w, "%s\t", h.Name)
	}
	fmt.Fprintln(w)
	for _, e := range r.Model.Entries {
		if e.HasData != nil && !*e.HasData {
			continue
		}
		if verbose {
			fmt.Fprintf(w, "%d\t", e.Id)
			if e.ParentId != nil {
				fmt.Fprintf(w, "%d", *e.ParentId)
			}
			fmt.Fprint(w, "\t")
		}
		for _, l := range e.Labels {
			fmt.Fprintf(w, "%s\t", l)
		}
		fmt.Fprintln(w)
	}
}

func (r *TableColumnHeadersResponse) PrintHeader(w *tabwriter.Writer, verbose bool) {
	fmt.Fprintln(w, "Status\tMessage\t")
}

func (r *TableColumnHeadersResponse) Print(w *tabwriter.Writer, verbose bool) {
	if r.Status != nil {
		fmt.Fprintf(w, "%s", *r.Status)
	}
	fmt.Fprint(w, "\t")
	if r.StatusMessage != nil {
		fmt.Fprintf(w, "%s", *r.StatusMessage)
	}
	fmt.Fprint(w, "\t")
	fmt.Fprintln(w)

	if r.Model == nil {
		return
	}
	fmt.Fprintln(w)

	fmt.Fprintf(w, "Name\tId\tType\t")
	if verbose {
		fmt.Fprintf(w, "Description\t")
	}
	fmt.Fprintln(w)
	for _, h := range *r.Model {
		if h.Name != nil {
			fmt.Fprintf(w, "%s", *h.Name)
		}
		fmt.Fprint(w, "\t")
		if h.Id != nil {
			fmt.Fprintf(w, "%d", *h.Id)
		}
		fmt.Fprint(w, "\t")
		if h.Type != nil {
			fmt.Fprintf(w, "%s", *h.Type)
		}
		fmt.Fprint(w, "\t")
		if verbose {
			if h.Description != nil {
				fmt.Fprintf(w, "%s", *h.Description)
			}
			fmt.Fprint(w, "\t")
		}
		fmt.Fprintln(w)
	}
}

func (r *VirtualTableResponse) PrintHeader(w *tabwriter.Writer, verbose bool) {
	// fmt.Fprintln(w, "Status\tMessage\t")
}

func (r *VirtualTableResponse) Print(w *tabwriter.Writer, verbose bool) {
	if r.Model == nil {
		fmt.Fprintln(w, "Status\tMessage\t")
	} else {
		fmt.Fprintln(w, "Status\tMessage\tSize\tLowIndex\tColumnIds\t")
	}
	if r.Status != nil {
		fmt.Fprintf(w, "%s", *r.Status)
	}
	fmt.Fprint(w, "\t")
	if r.StatusMessage != nil {
		fmt.Fprintf(w, "%s", *r.StatusMessage)
	}
	fmt.Fprint(w, "\t")
	if r.Model == nil {
		fmt.Fprintln(w)
		return
	}
	if r.Model.Size != nil {
		fmt.Fprintf(w, "%d", *r.Model.Size)
	}
	fmt.Fprint(w, "\t")
	if r.Model.LowIndex != nil {
		fmt.Fprintf(w, "%d", *r.Model.LowIndex)
	}
	fmt.Fprint(w, "\t")
	if r.Model.ColumnIds != nil {
		fmt.Fprintf(w, "%v", *r.Model.ColumnIds)
	}
	fmt.Fprint(w, "\t")
	fmt.Fprintln(w)

	if r.Model.Lines == nil {
		return
	}
	fmt.Fprintln(w)

	fmt.Fprintf(w, "Index\t")
	if r.Model.ColumnIds != nil {
		for _, c := range *r.Model.ColumnIds {
			fmt.Fprintf(w, "%d\t", c)
		}
	}
	fmt.Fprintln(w)

	for _, l := range *r.Model.Lines {
		if l.Index != nil {
			fmt.Fprintf(w, "%d", *l.Index)
		}
		fmt.Fprint(w, "\t")
		if l.Cells == nil {
			continue
		}
		for _, c := range *l.Cells {
			if c.Content != nil {
				fmt.Fprintf(w, "%s", *c.Content)
			}
			fmt.Fprint(w, "\t")
		}
		fmt.Fprintln(w)
	}
}

func (r *XYResponse) PrintHeader(w *tabwriter.Writer, verbose bool) {
	fmt.Fprintln(w, "Status\tMessage\t")
}

func (r *XYResponse) Print(w *tabwriter.Writer, verbose bool) {
	if r.Status != nil {
		fmt.Fprintf(w, "%s", *r.Status)
	}
	fmt.Fprint(w, "\t")
	if r.StatusMessage != nil {
		fmt.Fprintf(w, "%s", *r.StatusMessage)
	}
	fmt.Fprint(w, "\t")
	fmt.Fprintln(w)

	if r.Model == nil {
		return
	}
	fmt.Fprintln(w)

	fmt.Fprintln(w, "Model Title:", r.Model.Title)
	for _, series := range r.Model.Series {
		fmt.Fprintln(w)
		fmt.Fprintf(w, "Series Name(Id): %s(%d)\n", series.SeriesName, series.SeriesId)
		fmt.Fprintln(w, "X Values\tY Values\t")
		for i := 0; i < len(series.XValues); i++ {
			fmt.Fprintf(w, "%d\t%f\t\n", series.XValues[i], series.YValues[i])
		}
	}
}

func (r *TimeGraphArrowsResponse) PrintHeader(w *tabwriter.Writer, verbose bool) {
	fmt.Fprintln(w, "Status\tMessage\t")
}

func (r *TimeGraphArrowsResponse) Print(w *tabwriter.Writer, verbose bool) {
	if r.Status != nil {
		fmt.Fprintf(w, "%s", *r.Status)
	}
	fmt.Fprint(w, "\t")
	if r.StatusMessage != nil {
		fmt.Fprintf(w, "%s", *r.StatusMessage)
	}
	fmt.Fprint(w, "\t")
	fmt.Fprintln(w)
	if r.Model == nil {
		return
	}
	fmt.Fprintln(w)

	fmt.Fprintln(w, "SourceId\tTargetId\tStart\tEnd\t")
	for _, a := range *r.Model {
		fmt.Fprintf(w, "%d\t%d\t%d\t%d\t\n", a.SourceId, a.TargetId, a.Start, a.End)
	}

}

func (r *TimeGraphStatesResponse) PrintHeader(w *tabwriter.Writer, verbose bool) {
	fmt.Fprintln(w, "Status\tMessage\t")
}

func (r *TimeGraphStatesResponse) Print(w *tabwriter.Writer, verbose bool) {
	fmt.Fprintf(w, "%s\t%s\t\n", *r.Status, *r.StatusMessage)
	if r.Model == nil || r.Model.Rows == nil {
		return
	}
	for _, row := range *r.Model.Rows {
		fmt.Fprintln(w)
		fmt.Fprintf(w, "EntryId: %d\n", row.EntryId)

		fmt.Fprintln(w, "Label\tStart\tEnd\t")
		for _, state := range row.States {
			if state.Label != nil {
				fmt.Fprintf(w, "%s", *state.Label)
			}
			fmt.Fprint(w, "\t")
			fmt.Fprintf(w, "%d\t%d\t\n", state.Start, state.End)
		}
	}
}

func (r *TimeGraphTooltipResponse) PrintHeader(w *tabwriter.Writer, verbose bool) {
	fmt.Fprintln(w, "Status\tMessage\t")
}

func (r *TimeGraphTooltipResponse) Print(w *tabwriter.Writer, verbose bool) {
	fmt.Fprintf(w, "%s\t%s\t\n", *r.Status, *r.StatusMessage)
	if r.Model == nil {
		return
	}
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Key\tValue\t")
	for _, tp := range *r.Model {
		if tp.Key != nil {
			fmt.Fprintf(w, "%s\t", *tp.Key)
		}
		fmt.Fprint(w, "\t")
		if tp.Value != nil {
			fmt.Fprintf(w, "%s\t", *tp.Value)
		}
		fmt.Fprintln(w)
	}
}

func (r *TimeGraphTreeResponse) PrintHeader(w *tabwriter.Writer, verbose bool) {
	fmt.Fprintln(w, "Status\tMessage\t")
}

func (r *TimeGraphTreeResponse) Print(w *tabwriter.Writer, verbose bool) {
	fmt.Fprintf(w, "%s\t%s\t\n", *r.Status, *r.StatusMessage)
	if r.Model == nil {
		return
	}
	fmt.Fprintln(w)

	if verbose {
		fmt.Fprint(w, "Id\tParendId\t")
	}
	fmt.Fprint(w, "Start\tEnd\t")
	if r.Model.Headers != nil {
		for _, h := range *r.Model.Headers {
			fmt.Fprintf(w, "%s\t", h.Name)
		}
	}
	fmt.Fprintln(w)

	for _, e := range r.Model.Entries {
		if e.HasData != nil && !*e.HasData {
			continue
		}
		if verbose {
			fmt.Fprintf(w, "%d\t", e.Id)
			if e.ParentId != nil {
				fmt.Fprintf(w, "%d", *e.ParentId)
			}
			fmt.Fprint(w, "\t")
		}
		if e.Start != nil {
			fmt.Fprintf(w, "%d", *e.Start)
		}
		fmt.Fprint(w, "\t")
		if e.End != nil {
			fmt.Fprintf(w, "%d", *e.End)
		}
		fmt.Fprint(w, "\t")
		for _, l := range e.Labels {
			fmt.Fprintf(w, "%s\t", l)
		}
		fmt.Fprintln(w)
	}
}
