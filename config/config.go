// Package config is an interface for dynamic configuration.
package config

import (
	"context"

	"github.com/cqhorizon/cqsim-admin-core/config/loader"
	"github.com/cqhorizon/cqsim-admin-core/config/reader"
	"github.com/cqhorizon/cqsim-admin-core/config/source"
)

// Config is an interface abstraction for dynamic configuration
type Config interface {
	// Values provide the reader.Values interface
	reader.Values
	// Init the config
	Init(opts ...Option) error
	// Options in the config
	Options() Options
	// Close Stop the config loader/watcher
	Close() error
	// Load config sources
	Load(source ...source.Source) error
	// Sync Force a source changeset sync
	Sync() error
	// Watch a value for changes
	Watch(path ...string) (Watcher, error)
}

// Watcher is the config watcher
type Watcher interface {
	Next() (reader.Value, error)
	Stop() error
}

// Entity 配置实体
type Entity interface {
	OnChange()
}

// Options 配置的参数
type Options struct {
	Loader loader.Loader
	Reader reader.Reader
	Source []source.Source

	// for alternative data
	Context context.Context

	Entity Entity
}

// Option 调用类型
type Option func(o *Options)

// NewConfig returns new config
func NewConfig(opts ...Option) (Config, error) {
	return newConfig(opts...)
}
