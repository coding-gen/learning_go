package main
import "fmt"

func main() {
	var x string 
	x = "first"
	fmt.Println(x)

        var y string = "second"
        fmt.Println(x + y)

        z := "third"
        fmt.Println(z)
        fmt.Println(x == z)
        fmt.Println(x + y == "firstsecond")

	// Camelcase compound var name
        dogName := "Rover"
        fmt.Println(dogName + " where's your ball?!")

	var (
		a = "a"
		b = "b"
		c = "c"
	)
	fmt.Println(a + b + c )
	var (
		v = 5
		vi = 6
	)
	fmt.Println(v + vi)
}
