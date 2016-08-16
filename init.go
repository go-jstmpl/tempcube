package tempcube

import (
	"bytes"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func InitFile(p string, b bytes.Buffer) error {
	_, err := os.Stat(p)
	if err != nil {
		f, err := os.Create(p)
		if err != nil {
			return errors.Wrapf(err, "failed to create file %s", p)
		}

		_, err = f.WriteString(b.String())
		if err != nil {
			return errors.Wrapf(err, "failed to write file %s", p)
		}
	}
	return nil
}

func Init(p string) error {
	_, err := os.Stat(p)

	if err != nil {
		err = os.Mkdir(p, 0777)
		if err != nil {
			return errors.Wrapf(err, "failed to mkdir %s", p)
		}
	}

	fmt.Println("create starting ...")

	var tmpl bytes.Buffer
	tmpl.WriteString("package ")
	tmpl.WriteString(p)

	// sample/sample.go

	pt := fmt.Sprintf("%s/%s.go", p, p)
	err = InitFile(pt, tmpl)
	if err != nil {
		return errors.Wrapf(err, "failed to create file %s", pt)
	}
	fmt.Printf("--->  create %s\n", pt)

	// sample/sample_helper.go
	pt = fmt.Sprintf("%s/%s_helper.go", p, p)
	err = InitFile(pt, tmpl)
	if err != nil {
		return errors.Wrapf(err, "failed to create file %s", pt)
	}
	fmt.Printf("--->  create %s\n", pt)

	tmpl.WriteString("\n")
	tmpl.WriteString("import \"testing\"")

	// sample/sample_test.go
	pt = fmt.Sprintf("%s/%s_test.go", p, p)
	err = InitFile(pt, tmpl)
	if err != nil {
		return errors.Wrapf(err, "failed to create file %s", pt)
	}
	fmt.Printf("--->  create %s\n", pt)

	// sample/sample_helper_test.go
	pt = fmt.Sprintf("%s/%s_helper_test.go", p, p)
	err = InitFile(pt, tmpl)
	if err != nil {
		return errors.Wrapf(err, "failed to create file %s", pt)
	}
	fmt.Printf("--->  create %s\n", pt)

	fmt.Println("Finishing creating initial files.")
	return nil
}
