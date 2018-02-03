package main
import "fmt"

func main() {
	var x [10]int
	for i := 0 ; i < len(x) ; i++ {
		fmt.Println(x[i])
	}
        for i := 0 ; i < len(x) ; i++ {
		x[i] = i
        }
        for i := 0 ; i < len(x) ; i++ {
                fmt.Println(x[i])
        }
}
