package gomodtesting

import "fmt"

func Print(args ...interface{}) {
	fmt.Print("v3: ")
	fmt.Println(args...)
}
