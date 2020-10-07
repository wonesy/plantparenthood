package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateID(t *testing.T) {
	id := GenerateID()
	assert.NotEmpty(t, id)
}

func TestIDIsValid(t *testing.T) {
	goodID := GenerateID()
	badID := "badid"

	assert.True(t, IDIsValid(goodID))
	assert.False(t, IDIsValid(badID))
}
