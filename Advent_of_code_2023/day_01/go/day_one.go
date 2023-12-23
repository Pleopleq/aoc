package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

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
				integer, err := strconv.Atoi(string(line[i]))
				if err != nil {
					fmt.Println(err)
				}

				temp = integer * 10
				break
			}
		}

		for i := lineLenght; i >= 0; i-- {
			if isStringNumber(line[i]) {
				integer, err := strconv.Atoi(string(line[i]))
				if err != nil {
					fmt.Println(err)
				}
				temp += integer
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
