package main

import (
	"errors"
	"fmt"
)

func fmtErrorF() error {
	return fmt.Errorf("Error: %s", "fmtErrorF-example")
}

func errorsNew() error {
	return errors.New(fmt.Sprintf("Error: %s", "errorsNew-example"))
}

func main() {
	fmt.Printf("fmt.Errorf: %#v\n", fmtErrorF())
	fmt.Printf("errors.New: %#v\n", errorsNew())
}
