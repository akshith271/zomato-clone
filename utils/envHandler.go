package utils

import (
	"fmt"
	"os"
)

func CheckEnvError(envErr error) error {
	if envErr != nil {
		fmt.Printf("Could not load .env file")
		os.Exit(1)
	}
	return nil
}
