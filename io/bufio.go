package main

import (
	"bufio"
	"fmt"
	"os"
)

func bufio_read() {
	// Open the file
	file, err := os.Open("../pending.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	//  With bufio, you have control over the buffer size, allowing you to tune performance based on your specific requirements.
	//  You can adjust the buffer size according to the expected data size and the efficiency of I/O operations.
	//  No need to mention the hardcoded buffer size.

	// Also with bufio  bufio provides additional methods and abstractions that make working with buffered I/O easier and more convenient.
	// For example, it provides a Scanner type for scanning input and breaking it into lines or tokens, as well as methods like ReadLine,
	// ReadString, and ReadBytes for reading data in various formats.

	reader := bufio.NewReader(file)
	// Read from the reader.
	for {

		line, err := reader.ReadString('\n')

		if err != nil && err.Error() == "EOF" {
			fmt.Println("Error reading from file:", err)
			return
		}
		fmt.Println(line)
	}
}
