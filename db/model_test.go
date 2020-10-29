package db

import (
	"testing"
)

func TestNewModel(t *testing.T) {
	res := NewModel(nil)

	if res == nil {
		t.Error("Error creating new model")
	}
}
