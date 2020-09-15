package objectid

import (
	"testing"
)

func TestNew(t *testing.T) {
	objectId := New()
	t.Log(objectId)
	t.Log(objectId.Machine())
}

func TestEqual(t *testing.T) {
	objectId := New()
	compareObjectId, _ := Parse(objectId.String())
	if objectId != compareObjectId {
		t.Error("two instance is not equal.")
	}
}
