package utils

import "log"

func CheckError(err error) error {
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}
