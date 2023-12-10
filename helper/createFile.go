package helper

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

func ScanFile(file io.Reader) [][]string {
	var data [][]string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		data = append(data, fields)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return nil
	}

	return data
}

func CreateFileOrdernated(data [][]string, newFile io.Writer, filter, newFileName string) {
	filterCritery := findIndex(filter, data[0])

	sort.SliceStable(data[1:], func(i, j int) bool {
		return data[i+1][filterCritery] < data[j+1][filterCritery]
	})

	for _, row := range data {
		fmt.Fprintln(newFile, row)
	}

}

func findIndex(option string, slice []string) int {
	for i, item := range slice {
		if strings.EqualFold(option, item) {
			return i
		}
	}
	return -1
}
