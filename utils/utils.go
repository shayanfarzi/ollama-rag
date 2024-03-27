package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUserInput(promptString string) (string, error) {
	fmt.Print(promptString, ": ")
	var Input string
	reader := bufio.NewReader(os.Stdin)

	Input, _ = reader.ReadString('\n')

	Input = strings.TrimSuffix(Input, "\n")
	Input = strings.TrimSuffix(Input, "\r")

	return Input, nil
}

// for books in https://www.gutenberg.org/ with PlainText .txt format
func ClearTexts(dirFile string) (string, error) {
	file, err := os.ReadFile(dirFile)
	if err != nil {
		return "", err
	}
	stringFile := string(file)

	stringFile = strings.ReplaceAll(stringFile, "\n", " ")
	stringFile = strings.ReplaceAll(stringFile, "\r", " ")
	stringFile = strings.ReplaceAll(stringFile, "       ", " ")
	stringFile = strings.ReplaceAll(stringFile, "    ", " ")

	return stringFile, nil
}
