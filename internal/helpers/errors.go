package helpers

import (
	"fmt"
	"io"
	"os"
)

func AssertNoErrors(err error) bool {
	return err != nil
}

func FatalIfErr(err error, out io.Writer) {
	if !AssertNoErrors(err) {
		fmt.Fprintf(out, "%s", err.Error())
		os.Exit(1)
	}
}
