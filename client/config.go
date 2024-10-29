package client

import (
	"context"
)

type GlobalConfig struct {
	Verbose       bool
	Tspc          *ClientWithResponses
	OutputType    OutputType
	FailOnHTTPErr bool
}

type key int

var cfgKey key

// NewContext returns a new Context that carries value cfg
func NewContext(ctx context.Context, cfg *GlobalConfig) context.Context {
	return context.WithValue(ctx, cfgKey, cfg)
}

// FromContext returns the GlobalConfig value stored in ctx, if any.
func FromContext(ctx context.Context) (*GlobalConfig, bool) {
	cfg, ok := ctx.Value(cfgKey).(*GlobalConfig)
	return cfg, ok
}

// MustFromContext returns the GlobalConfig value stored in ctx
// or dies horribly.
func MustFromContext(ctx context.Context) *GlobalConfig {
	cfg, ok := FromContext(ctx)
	if !ok {
		panic("global config not found in context")
	}
	return cfg
}
