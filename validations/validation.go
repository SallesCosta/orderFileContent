package validations

import (
	"fmt"
	"os"

	"github.com/sallescosta/readCSVfile/helper"
)

func ArgsValidation() {
	if len(os.Args) != 3 {
		fmt.Println("***FileName or NewFileName is undefined.***")
		return
	}
}

func FilterValidation(i string) bool {
	_, exists := helper.InputMap[i]
	return exists
}
