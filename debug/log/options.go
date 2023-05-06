package log

import "time"

// Option used by the logger
type Option func(*Options)

// Options are logger options
type Options struct {
	// Name of the log
	Name string
	// Size is the size of ring buffer
	Size int
	// Format specifies the output format
	Format FormatFunc
}

// Name of the log

// Size sets the size of the ring buffer

// DefaultOptions returns default options

// ReadOptions for querying the logs
type ReadOptions struct {
	// Since what time in past to return the logs
	Since time.Time
	// Count specifies number of logs to return
	Count int
	// Stream requests continuous log stream
	Stream bool
}

// ReadOption used for reading the logs
type ReadOption func(*ReadOptions)

// Since sets the time since which to return the log records

// Count sets the number of log records to return
