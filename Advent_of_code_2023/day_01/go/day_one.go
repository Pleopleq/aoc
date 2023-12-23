package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

var filePath = "D:\\Documentos\\adventofcode\\Advent_of_code_2023\\day_01\\go\\test_file.txt"

func main() {
	var total int

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	for _, line := range bytes.Split(fileContent, []byte("\n")) {
		lineLenght := len(line) - 1
		var temp int

		for i := 0; i < lineLenght; i++ {
			if isStringNumber(line[i]) {
				temp = int(line[i]) * 10

				break
			}
		}

		for i := lineLenght; i >= 0; i-- {
			if isStringNumber(line[i]) {
				temp += int(line[i])
				break
			}
		}

		total += temp
	}

	fmt.Println(total)
}

func isStringNumber(input byte) bool {
	str := string(input)
	_, err := strconv.Atoi(str)
	return err == nil
}
