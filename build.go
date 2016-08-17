package tempcube

import (
	"fmt"
	"os/exec"
	"sync"

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
	var once = new(sync.Once)

	once.Do(func() {
		// go get -u
		cmd = fmt.Sprintf("go get -u github.com/minodisk/go-jstmpl/cmd/jstmpl")
		o, err := Command(cmd)
		if err != nil && o != "" {
			fmt.Println(errors.Wrapf(err, "Getting go-jstmpl, Out: %s", o))
		}
	})
	// jstmpl
	cmd = fmt.Sprintf("jstmpl -s %s -t %s -o %s", sc, p, d)
	o, err := Command(cmd)
	if err != nil {
		return errors.Wrapf(err, "building code using go-jstempl, Out: %s", o)
	}

	fmt.Println("... Build succesd")
	return nil
}
