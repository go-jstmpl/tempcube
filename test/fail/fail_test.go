package fail

import "testing"

func TestPass(t *testing.T) {
	s := Pass()
	if s != "ok" {
		t.Fatalf("expect not ok, but %s", s)
	}
}
