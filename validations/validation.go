package validations

import (
	"os"
)

func ArgsValidation() bool {
	if len(os.Args) == 3 {
		return true
	}
	return false
}

func FilterValidation(i string, firstLine []string) bool {
	for _, item := range firstLine {
		if item == i {
			return true
		}
	}
	return false
}
