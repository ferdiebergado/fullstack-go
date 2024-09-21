package env

import (
	"fmt"
	"os"
)

func GetEnv(v string) string {
	res, exists := os.LookupEnv(v)

	if !exists {
		fmt.Fprintln(os.Stderr, v+" not set!")
		os.Exit(1)
	}

	return res
}
