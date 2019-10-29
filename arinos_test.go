package arinos

import (
	"testing"
)

func TestNew(t *testing.T) {
	t.Log("Start")
	result := New(false)
	expect := &Arinos{
		LocalHost: false,
		Options: &Options{
			Port: 8000,
		},
	}
	if result.LocalHost != expect.LocalHost &&
		result.Options.Port != result.Options.Port {
		t.Error("\nResult： ", result, "\nExpect： ", expect)
	}
	t.Log("End")
}
