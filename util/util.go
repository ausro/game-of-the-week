package util

import (
	"fmt"
	"os"
)

func Must[T any](obj T, err error) T {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	return obj
}
