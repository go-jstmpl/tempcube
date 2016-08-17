package tempcube

import (
	"os"
	"testing"
)

func TestBuild(t *testing.T) {
	path := "test/pass"
	dest := "test_build"
	scm := "test/schema.yml"

	err := Build(path, dest, scm)
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dest)

	for _, f := range PathList(dest, "pass") {
		_, err := os.Stat(f)
		if err != nil {
			t.Errorf("%s file not exist \n", f)
		}
	}
}
