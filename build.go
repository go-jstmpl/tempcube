package tempcube

import (
	"fmt"
	"os/exec"

	"github.com/mattn/go-shellwords"
	"github.com/pkg/errors"
)

func Command(cmd string) (string, error) {
	args, err := shellwords.Parse(cmd)
	if err != nil {
		return "", errors.Wrap(err, cmd)
	}
	out, err := exec.Command(args[0], args[1:]...).Output()
	if err != nil {
		return "", errors.Wrap(err, cmd)
	}
	return string(out), nil
}

func Build(p, d, sc string) error {
	var cmd string

	// go get -u
	cmd = fmt.Sprintf("go get -u github.com/minodisk/go-jstmpl/cmd/jstmpl")
	o, err := Command(cmd)
	if err != nil && o != "" {
		return errors.Wrapf(err, "Getting go-jstmpl, Error: %s, Out: %s", err, o)
	}

	// jstmpl
	cmd = fmt.Sprintf("jstmpl -s %s -t %s -o %s", sc, p, d)
	o, err = Command(cmd)
	if err != nil {
		return errors.Wrapf(err, "building code using go-jstempl, Error: %s, Out: %s", err, o)
	}

	fmt.Println("...Building Succeed")
	return nil
}
