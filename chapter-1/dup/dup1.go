package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// mark true for the example which you want to be executed ...
	//TODO: Fix this how to run different examples ???
	dup3()
	dup2()
	dup()
}

func dup() {
	reader := bufio.NewReader(os.Stdin)

	counter := make(map[string]int)

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err.Error() == "EOF" {
			fmt.Fprintf(os.Stderr, "Error occured: %v \n", err)
			break
		}

		counter[line]++
	}
	for line, count := range counter {
		if count > 1 {
			fmt.Println(count, " ", line)
		}
	}
}
