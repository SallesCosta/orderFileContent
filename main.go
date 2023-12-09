package main

import (
	"fmt"
	"os"

	"github.com/sallescosta/readCSVfile/helper"
	"github.com/sallescosta/readCSVfile/validations"
)

func main() {
	var filter string

	validations.ArgsValidation()

	fileName := os.Args[1]
	newFileName := os.Args[2]

	fmt.Println("What is the criterion for filtering? (Name, Age, or Points)")
	fmt.Scan(&filter)

	if !validations.FilterValidation(filter){
		fmt.Println("***You must choose Name, Age, or Points.***")
		return
	}

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}

	newFile, err := os.Create(newFileName)
	if err != nil {
		fmt.Println("erro ao criar o arquivo.", err)
	}

	defer file.Close()
	defer newFile.Close()


	data := helper.ScanFile(file)

	helper.CreateFileOrdernated(data, newFile, filter, newFileName)

	fmt.Printf("Created new file %s filtered by %s\n", newFileName, filter)
}
