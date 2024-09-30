package env

import (
	"fmt"
	"os"
)

func Must(v string) string {
	res, exists := os.LookupEnv(v)

	if !exists {
		fmt.Fprintf(os.Stderr, "%s not set!\n", v)
		os.Exit(1)
	}

	return res
}
