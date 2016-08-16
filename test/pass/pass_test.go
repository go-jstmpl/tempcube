package pass

import "testing"

func TestPass(t *testing.T) {
	s := Pass()
	if s != "ok" {
		t.Fatalf("expect ok, but %s", s)
	}
}
