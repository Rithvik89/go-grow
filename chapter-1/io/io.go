package main

import (
	"fmt"
	"io"
	"os"
)

func io_read() {
	// Open the file
	file, err := os.Open("../pending.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// we can directly use 'file' as an io.Reader because os.File implements the
	// io.Reader interface, among others, because it has a Read method with the signature
	// required by io.Reader.

	// Read from the file.
	for {
		// Read data into a byte slice
		buffer := make([]byte, 50) // Example buffer size
		bytesRead, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			fmt.Println("Error reading from file:", err)
			return
		}
		fmt.Println(bytesRead)

		// Check for end of file
		if bytesRead == 0 {
			fmt.Println(bytesRead)
			break
		}

		// Process the read data (in this example, simply print it)
		fmt.Printf("%v\n", string(buffer[:bytesRead]))
	}
}
