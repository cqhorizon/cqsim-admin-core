package reader

import (
	"github.com/cqhorizon/cqsim-admin-core/config/encoder"
	"github.com/cqhorizon/cqsim-admin-core/config/encoder/json"
	"github.com/cqhorizon/cqsim-admin-core/config/encoder/yaml"
)

type Options struct {
	Encoding map[string]encoder.Encoder
}

type Option func(o *Options)

func NewOptions(opts ...Option) Options {
	options := Options{
		Encoding: map[string]encoder.Encoder{
			"json": json.NewEncoder(),
			"yml":  yaml.NewEncoder(),
		},
	}
	for _, o := range opts {
		o(&options)
	}
	return options
}
