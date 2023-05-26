package memory

import (
	"github.com/cqhorizon/cqsim-admin-core/config/loader"
	"github.com/cqhorizon/cqsim-admin-core/config/reader"
)

// WithSource appends a source to list of sources

// WithReader sets the config reader
func WithReader(r reader.Reader) loader.Option {
	return func(o *loader.Options) {
		o.Reader = r
	}
}
