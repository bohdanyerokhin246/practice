package main

import (
	"fmt"
	"strings"
)

func main() {
	string := "Hello NIX Solutions"
	averageWordLength(string)
}
func averageWordLength(str string) {
	numberOfSpaces := 0
	numberOfSpaces = strings.Count(str, " ")
	fmt.Println(numberOfSpaces)
	fmt.Println(len(str) / (numberOfSpaces + 1))
}
