package main

import (
	"fmt"
	"os"
	"strings"
)

func arg() {
	fmt.Printf("args: %s\n", strings.Join(os.Args[1:], " "))

}

// func index_and_arg() {
// 	for i, arg := range os.Args[1:] {
// 		fmt.Println(i, arg)
// 	}
// }
