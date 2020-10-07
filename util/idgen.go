package util

import (
	"github.com/segmentio/ksuid"
)

// GenerateID creates a new KSUID ID
func GenerateID() string {
	return ksuid.New().String()
}

// IDIsValid checks if a provided ID is a valid ksuid
func IDIsValid(id string) bool {
	_, err := ksuid.Parse(id)
	return err == nil
}
