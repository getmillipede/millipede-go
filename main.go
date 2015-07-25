package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/getmillipede/millipede-go/millipede"
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("usage: millipede <size>")
		os.Exit(1)
	}

	size, err := strconv.ParseUint(args[0], 10, 0)
	if err != nil {
		panic(err)
	}

	milli := millipede.New(size).Reverse().Reverse()
	milli.Println()
}
