package tempcube

import (
	"fmt"
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	p := "sample_project"
	err := Init(p)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(p)

	// sample_project case init files
	fs := []string{
		fmt.Sprintf("%s/%s.go", p, p),
		fmt.Sprintf("%s/%s_helper.go", p, p),
		fmt.Sprintf("%s/%s_test.go", p, p),
		fmt.Sprintf("%s/%s_helper_test.go", p, p),
	}

	for _, f := range fs {
		_, err = os.Stat(f)
		if err != nil {
			t.Error(err)
		}
	}
}
