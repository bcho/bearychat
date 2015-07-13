package bearychat

import (
	"testing"
)

func TestClientSetHook(t *testing.T) {
	c := NewClient("")

	expectedHook := "foobar"
	c.SetHook(expectedHook)
	if c.Hook != expectedHook {
		t.Errorf("expected hook %s got %s", expectedHook, c.Hook)
	}
}
