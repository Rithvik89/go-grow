package main

import (
	"bufio"
	"fmt"
	"os"
)

func dup2() {
	files := os.Args[1:]
	counter := make(map[string]int)

	if len(files) == 0 {
		std_counter(os.Stdin, &counter)
	} else {
		for _, file := range files {
			file, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Err while opening file: %v \n", err)
			}
			std_counter(file, &counter)
		}
	}
}

// Stdin, Stdout, and Stderr are open Files pointing to the standard input,
// standard output, and standard error file descriptors.
// So we passed an argumet with object os.file

func std_counter(rd *os.File, counter *map[string]int) {
	reader := bufio.NewReader(rd)

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err.Error() == "EOF" {
			fmt.Fprintf(os.Stderr, "Error occured: %v \n", err)
			break
		}

		(*counter)[line]++
	}
	for line, count := range *counter {
		if count > 1 {
			fmt.Println(count, " ", line)
		}
	}
}

// Yes, you can create a function that takes a pointer to an io.Reader as an argument.
//Since os.File implements the io.Reader interface, you can pass an os.File object to such a function.
