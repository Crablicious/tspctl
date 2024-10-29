package cmdutil

import (
	"errors"
	"strconv"
	"strings"

	"github.com/Crablicious/tspctl/client"
)

type TimeRange client.TimeRange

func (r *TimeRange) String() string {
	if r == nil {
		return ""
	}
	return strconv.FormatInt(int64(r.Start), 10) + "," + strconv.FormatInt(int64(r.End), 10)
}

func (r *TimeRange) Set(value string) error {
	startEnd := strings.Split(value, ",")
	if len(startEnd) != 2 {
		return errors.New("a timerange should consist of two comma separated values <start>,<end>")
	}
	start, err := strconv.ParseInt(startEnd[0], 10, 0)
	if err != nil {
		return err
	}
	end, err := strconv.ParseInt(startEnd[1], 10, 0)
	if err != nil {
		return err
	}
	r.Start = start
	r.End = end
	return nil
}

func (r *TimeRange) Type() string {
	return "timerange"
}
