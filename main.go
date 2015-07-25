package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func millipede(size uint64) string {
	paddingOffsets := []string{"  ", " ", "", " ", "  ", "   ", "    ", "    ", "   "}

	bodyLines := []string{"    ╚⊙ ⊙╝"}
	var x uint64
	for x = 0; x < size; x++ {
		line := paddingOffsets[x%9] + "╚═(███)═╝"
		bodyLines = append(bodyLines, line)
	}

	return strings.Join(bodyLines, "\n")
}

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

	fmt.Println(millipede(size))
}
