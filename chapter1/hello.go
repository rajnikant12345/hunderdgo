//this is a comment
//this is first line of every go parogam.
// main means this is the main package and our main function is present in this package
package main

// one go program can import other go packages
// here we are importing a different package named "fmt"
import (
	"fmt"
)

// this is main function
func main() {
	// and this is where we call a function present in "fmt" package
	fmt.Println("Hello, Gophers....")
}


//Key points:
/*
1. A Go programs always begins with packahe keyword.
2. main() is the entry point for a program , in Go . Just like C or C++ ot Java.
3. for running a go program
	a. Use "go run hello.go"
	b. -- There is another way, but we will discuss that in later part----
