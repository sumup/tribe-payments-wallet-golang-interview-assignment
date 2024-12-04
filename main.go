package main

import (
	"fmt"
	stdOs "os"

	"github.com/sumup-oss/go-pkgs/os"

	"github.com/sumup/payments-bank-account-golang-assignment/internal/cmd"
)

func main() {
	osExecutor := &os.RealOsExecutor{}

	err := cmd.NewRootCmd(osExecutor).Execute()
	if err != nil {
		_, _ = fmt.Fprintf(osExecutor.Stderr(), "unsuccessful command, err: %+v", err)

		stdOs.Exit(1)
	}
}
