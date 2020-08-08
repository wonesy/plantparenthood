package util

import (
	"github.com/segmentio/ksuid"
)

// GenerateID creates a new KSUID ID
func GenerateID() string {
	return ksuid.New().String()
}
