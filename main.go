package main

import (
	"fmt"
	stdOs "os"

	"github.com/sumup-oss/go-pkgs/os"

	"tribe-payments-wallet-golang-interview-assignment/internal/cmd"
)

func main() {
	osExecutor := &os.RealOsExecutor{}

	err := cmd.NewRootCmd(osExecutor).Execute()
	if err != nil {
		_, _ = fmt.Fprintf(osExecutor.Stderr(), "unsuccessful command, err: %+v", err)

		stdOs.Exit(1)
	}
}
