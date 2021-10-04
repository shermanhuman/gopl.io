// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 4.
//!+

// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s string = ""
	for i := 1; i < len(os.Args); i++ {
		s += strconv.Itoa(i) + " " + os.Args[i] + "\n"
	}
	fmt.Println(s)
}

//!-
