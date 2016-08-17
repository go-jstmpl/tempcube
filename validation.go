package tempcube

import (
	"fmt"
	"os"
	"path"

	"github.com/pkg/errors"
)

func PathList(d, p string) []string {
	return []string{
		fmt.Sprintf("%s/%s.go", d, p),
		fmt.Sprintf("%s/%s_helper.go", d, p),
		fmt.Sprintf("%s/%s_test.go", d, p),
		fmt.Sprintf("%s/%s_helper_test.go", d, p),
	}
}

func Validate(p, sc string) error {
	fmt.Println("validate testing...")

	_, f := path.Split(p)
	pt := "test_" + f
	err := Build(p, pt, sc)
	if err != nil {
		return errors.Wrap(err, "generating code")
	}
	defer os.RemoveAll(pt)

	cmds := PathList(pt, f)

	// gofmt
	for _, cmd := range cmds {
		o, err := Command("go fmt " + cmd)
		if err != nil {
			return errors.Wrapf(err, "gofmt failed, Out: %s, Error", o)
		}

		if o != "" && o != cmd {
			return errors.Wrapf(err, "gofmt failed, Out: %s, Error", o)

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
			return errors.Wrapf(err, "go test failed, Out: %s, Error", o)
		}
	}

	fmt.Println("... testing OK")
	return nil
}
