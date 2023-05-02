// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.
//!+

// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	// 短变量声明创建bufio.Scanner类型的变量input.该变量从程序的标准输入中读取内容。每次调用input.Text()得到。Scan 函数在读到一行时返回true，不再有输入时返回false
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		// set up the end point for the user input
		// or client can do ctrl + D is EOF end of input.
		// if input.Text() == "exit" {
		// 	break
		// }
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

//!-
