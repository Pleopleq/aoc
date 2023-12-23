package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var filePath = "/home/pleopleq/Documents/Advent_of_code_2023/day_01/go/actual_file.txt"

func main() {
	var total int64
	var unparsedNums [][]string

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	for _, line := range bytes.Split(fileContent, []byte("\n")) {
		lineLenght := len(line) - 1
		var lineIndexFound int
		var temp []string

		for i := 0; i < lineLenght; i++ {
			if isStringNumber(line[i]) {
				temp = append(temp, string(line[i]))
				lineIndexFound = i + 1
				break
			}
		}

		if lineIndexFound == lineLenght-1 {
			lineIndexFound = 0
			break
		}

		for i := lineLenght; i > lineIndexFound; i-- {
			if isStringNumber(line[i]) {
				temp = append(temp, string(line[i]))
				lineIndexFound = 0
				break
			}
		}

		unparsedNums = append(unparsedNums, temp)
	}

	for _, numStr := range unparsedNums {
		if len(numStr) == 1 {
			numStr = append(numStr, numStr[0])
		}

		joinedNumStr := strings.Join(numStr, "")
		parsedNums, err := strconv.ParseInt(joinedNumStr, 10, 64)
		fmt.Println(parsedNums)
		if err != nil {
			fmt.Println(err)
		}

		total = total + parsedNums
	}

	fmt.Println(total)
}

func isStringNumber(input byte) bool {
	str := string(input)
	_, err := strconv.Atoi(str)
	return err == nil
}
