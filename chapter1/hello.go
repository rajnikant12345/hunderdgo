//this is a comment
//this is first line of every go parogam.
// main means this is the main package and our main function is present in this package
package main

// one go program can import other go packages
// here we are importing a dofferent package named "fmt"
import (
	"fmt"
)

// this is main function
func main() {
	// and this is where we call a function present in "fmt" package
	fmt.Println("Hello, playground")
}
