package main

// go build - builds an executeable command that can be run on any other system without even setting up go environment
// after building use --> ./FileName <-- to directly run the code without using gu run command
// Recompile everytime before running the code
// go build && ./FileName

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
}
