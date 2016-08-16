package tempcube

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func Test(p, sc string) error {
	fmt.Println("validate testing...")

	pt := "test_" + p
	err := Build(p, pt, sc)
	if err != nil {
		return errors.Wrap(err, "generating code")
	}
	defer os.RemoveAll(pt)

	cmds := []string{
		fmt.Sprintf("%s/%s.go", pt, p),
		fmt.Sprintf("%s/%s_helper.go", pt, p),
		fmt.Sprintf("%s/%s_test.go", pt, p),
		fmt.Sprintf("%s/%s_helper_test.go", pt, p),
	}

	// gofmt
	for _, cmd := range cmds {
		o, err := Command("go fmt" + cmd)
		if err != nil && o != "" {
			return errors.Wrapf(err, "gofmt failed, Out: %s", o)
		}
	}

	fmt.Println("... validation test OK")
	fmt.Println("testing ...")

	// go test
	cmd := fmt.Sprintf("cd %s && go test", pt)
	o, err := Command(cmd)
	if err != nil && o[0:4] != "PASS" {
		// there is no test
		if o[0:7] != "testing" {
			return errors.Wrapf(err, "go test failed, Out: %s", o)
		}
	}

	fmt.Println("... testing OK")
	return nil
}
