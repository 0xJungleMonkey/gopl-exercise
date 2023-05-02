// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	//requirement: 打印每个参数的索引和值，每个一行。
	for index, arg := range os.Args[1:] {
		fmt.Println(index, arg)
	}
	
}

//!-
