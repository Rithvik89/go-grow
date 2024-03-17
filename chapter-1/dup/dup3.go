package main

import (
	"fmt"
	"os"
	"strings"
)

func dup3() {
	counter := make(map[string]int)
	for _, file := range os.Args[1:] {
		bytes, err := os.ReadFile(file)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while reading file: %v \n", err)
		}
		for _, line := range strings.Split(string(bytes), "\n") {
			counter[line]++
		}
	}
	for line, count := range counter {
		if count > 1 {
			fmt.Println(count, " ", line)
		}
	}
}
