package db

import (
	"fmt"
)

// NoSuchEntity error for when DB cannot find anything
type NoSuchEntity struct {
	Type string
}

func (e *NoSuchEntity) Error() string {
	if e.Type == "" {
		return fmt.Sprintf("No such entity")
	}
	return fmt.Sprintf("No such entity of type %s", e.Type)
}
