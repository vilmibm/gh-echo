package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	fmt.Println(strings.Join(os.Args[1:len(os.Args)], " "))
}
