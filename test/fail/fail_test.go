package fail

import "testing"

func TestPass(t *testing.T) {
	s := Fail()
	if s == "ok" {
		t.Fatalf("expect not ok, but %s", s)
	}
}
