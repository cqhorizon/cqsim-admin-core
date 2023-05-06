package config

type Resource struct {
	Mongo        string
	License      string
	Dist         string
	Algorithm    string
	AllowDiskUse bool
}

var ResourceConfig = new(Resource)
