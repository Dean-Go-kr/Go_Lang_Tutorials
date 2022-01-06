package main

import (
	"fmt"
	"strings"
)

/*
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}
*/

//"naked" return
func lenAndUpper(name string) (lenght int, uppercase string) {
	// "defer" runs a code after the function finishes
	defer fmt.Println("I'm done")
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	totalLenght, upperName := lenAndUpper("nico")
	fmt.Println(totalLenght, upperName)

	repeatMe("nico", "deango", "hello")
}
