package tempcube

import (
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	p := "sample_project"
	err := Init(p)
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(p)

	for _, f := range PathList(p, p) {
		_, err = os.Stat(f)
		if err != nil {
			t.Error(err)
		}
	}
}
