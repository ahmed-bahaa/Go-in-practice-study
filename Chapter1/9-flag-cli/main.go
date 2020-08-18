package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "world", "this option -name- has default value world")
var italian bool

func init() {
	flag.BoolVar(&italian, "italian", false, "Use italian language")
	flag.BoolVar(&italian, "i", false, "Use italian language short name")
}

func main() {
	flag.Parse()
	if italian == true {
		fmt.Printf("Ciao %s!\n", *name)
	} else {
		fmt.Printf("Hello %s!\n", *name)
	}
}
