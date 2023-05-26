package file

import (
	"github.com/cqhorizon/cqsim-admin-core/config/encoder"
	"strings"
)

func format(p string, e encoder.Encoder) string {
	parts := strings.Split(p, ".")
	if len(parts) > 1 {
		return parts[len(parts)-1]
	}
	return e.String()
}
