package main

import (
	"fmt"
	"strings"

	"golang.org/x/exp/slices"
)

func isThisTrue(value bool) bool {
	result := value == true

	return result
}

func yeahNah(input string) bool {

	splitFields := strings.Fields(input)
	fmt.Println(slices.Contains(splitFields, "yeah"))
	fmt.Println(slices.Index(splitFields, "nah"))

	for i, v := range splitFields {
		fmt.Println("Fields are:", v)
		if strings.Contains(v, "yeah") && i == 0 {
			return true
		}
		return false
	}
	return false
}

func main() {
	/* fmt.Println("go" + "lang")
	fmt.Println("1+1= ", 1+1)
	fmt.Println("7.0/3.0= ", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	//comment?
	var a = "inital"
	println(a)
	var ab string = "123"
	println(ab)

	println(isThisTrue(false))
	println(isThisTrue(true)) */
	println(yeahNah("nah yeah nah"))
}
