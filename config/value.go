package config

import (
	"time"

	"cqsim-admin-core/config/reader"
)

type value struct{}

func newValue() reader.Value {
	return new(value)
}

func (v *value) Bool(_ bool) bool {
	return false
}

func (v *value) Int(_ int) int {
	return 0
}

func (v *value) String(_ string) string {
	return ""
}

func (v *value) Float64(_ float64) float64 {
	return 0.0
}

func (v *value) Duration(_ time.Duration) time.Duration {
	return time.Duration(0)
}

func (v *value) StringSlice(_ []string) []string {
	return nil
}

func (v *value) StringMap(_ map[string]string) map[string]string {
	return map[string]string{}
}

func (v *value) Scan(_ interface{}) error {
	return nil
}

func (v *value) Bytes() []byte {
	return nil
}
