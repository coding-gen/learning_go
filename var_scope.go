package main
import "fmt"

/// Must be declarative statement outside of function. Like this:
var x string = "Hello World"

func main() {
	x = "old socks"
	fmt.Println(x)
}
