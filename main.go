package main

// go build - builds an executeable command that can be run on any other system without even setting up go environment
// after building use --> ./FileName <-- to directly run the code without using gu run command
// Recompile everytime before running the code
// go build && ./FileName

///////////////// IMPORTS ///////////
// 1) Copy Path name
// 2) use the package.Function(parameter)
// 3) go get path

import (
	"fmt"

	"github.com/RitaHC/mystrings"
)

func main() {
	fmt.Println(mystrings.Reverse("Hello World"))
	fmt.Println("------------------------------")
	// fmt.Println()
}
