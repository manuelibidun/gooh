package main

import (
	"flag"
	"fmt"
)

func main() {
	var filename string
	flag.StringVar(&filename,
		"filename",
		"problems.csv",
		"Input file",
	)
	flag.StringVar(&filename,
		"f",
		"problems.csv",
		"Input file",
	)
	flag.Parse()

	fmt.Println(filename)
}
