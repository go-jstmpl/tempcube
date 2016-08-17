package tempcube

import (
	"testing"
)

// pass
func TestVaridatePass(t *testing.T) {
	path := "test/pass"
	scm := "test/schema.yml"

	err := Validate(path, scm)
	if err != nil {
		t.Errorf("Pass test failed: %s", err)
	}
}

// gofmt fail
func TestVaridateFailFmt(t *testing.T) {
	path := "test/invalid"
	scm := "test/schema.yml"

	err := Validate(path, scm)
	if err == nil {
		t.Error("Fail gofmt test failed")
	}

}

// gofmt pass and test fail
func TestVaridateFailTest(t *testing.T) {
	path := "test/invalid"
	scm := "test/schema.yml"

	err := Validate(path, scm)
	if err == nil {
		t.Error("Fail gotest test failed")
	}

}
