package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	/* "golang.org/x/exp/slices" */)

func isThisTrue(value bool) bool {
	result := value == true

	return result
}

func yeahNah(input string) string {

	splitFields := strings.Fields(input)
	/* fmt.Println(slices.Contains(splitFields, "yeah"))
	fmt.Println(slices.Index(splitFields, "nah")) */

	lastIndex := len(splitFields) - 1

	if splitFields[lastIndex] == "yeah" {
		return "this means yes"
	}
	if splitFields[lastIndex] == "nah" {
		return "this means no"
	}

	/* 	if splitFields[lastIndex] != "yeah" && splitFields[lastIndex] != "nah" {
		println()
	} */
	return "who knows"
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	char, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	println(yeahNah(char))
}
