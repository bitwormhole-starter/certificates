package unit4libcerts

import "github.com/bitwormhole/starter/application"

// ExportConfig ...
func ExportConfig(cb application.ConfigBuilder) error {
	return autoGenConfig(cb)
}
