package main

import (
	"fmt"
	"os"

	"github.com/sallescosta/readCSVfile/helper"
	"github.com/sallescosta/readCSVfile/validations"
)

func main() {
	var filter string

	if !validations.ArgsValidation() {
		fmt.Println("***FileName or NewFileName is undefined.***")
		return
	}

	fileName := os.Args[1]
	newFileName := os.Args[2]

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
	firstLine := data[0]

	fmt.Printf("What is the criterion for filtering? %v", firstLine)
	fmt.Scan(&filter)

	if !validations.FilterValidation(filter, firstLine){
		fmt.Printf("***You must choose one of %v ***", firstLine)
		return
	}

	helper.CreateFileOrdernated(data, newFile, filter, newFileName)

	fmt.Printf("Created new file %s filtered by %s\n", newFileName, filter)
}
