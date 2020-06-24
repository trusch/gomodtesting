package gomodtesting

import "fmt"

func Print(args ...interface{}) {
	fmt.Print("v2: ")
	fmt.Println(args...)
}
