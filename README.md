# tspctl

A trace-server-protocol client in Go.

Not fully tested. Not complete.

## Building

```
go build
```

## Updating the openapi specification

Check out a new version of the [trace-server-protocol
submodule](client/trace-server-protocol).

Generate new client-code:
```
go generate -C client
```

Fix any compatibility issues or add functionality for the updates.

Commit the new version of the trace-server-protocol submodule as well
as the newly generated client-code.

## Required kludge

The reference implementation as of writing is not one-to-one to the
openapi specification. The trace-server-protocol submodule is
therefore a fork with a compatibility patch on top.

- "parameters" not in request body for putTrace
- "parameters" not in request body for postExperiment
- putTrace has error responses listed as being content-type
  "application/json" but the server returns strings without quotation
  marks or even HTML making the response invalid JSON.
  - Changed all error responses to text/plain on the assumption that
    all of them exhibits this behavior.
- What is the parameter for TreeParameters? requested\_timerange or
  requested\_times? None seem to do anything. YAML says
  requested\_times and example says requested\_timerange.
- OutputElementStyle's values gets generated as
  `*map[string]map[string]interface{}` but the response JSON is:
  ```json
  "values": {
    "series-type": "scatter"
  }
  ```
  It's defined in openapi spec as:
  ```yaml
  values:
    type: object
    additionalProperties:
      type: object
  ```
  additionalProperties is supposed to be an object but is the string
  "scatter".
- Trace-server produces yValues like 0.0 for XYModel which Go's JSON
  parser does not like since it cannot unmarshal it into an int64.
  - The trace-server is actually marshalling doubles into yValues.
    `ISeriesModel:getData` returns a list of doubles. Should yValues
    actually be doubles?
- TreeParameters for time graph tree is documented as requested_times,
  but example says requested\_timerange.

## TODO

- Uplift to a newer protocol.
- A convenience subcommand, e.g. to open a list of trace paths in an
  experiment.
- Nicer virtual table output
