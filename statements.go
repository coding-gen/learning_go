package main
import "fmt"
import "time"

func main() {

	for i := 0 ; i < 10; i++ {
		fmt.Println(i)
		if i %2 == 0{
			fmt.Println("Even.")
		} else {
			fmt.Println("Odd.")
		}
	}
	fmt.Println(time.Now())
	for i:= 10 ; i < 100; i += 2 {
		fmt.Println(i)
		switch true {
		case i%7 == 0: fmt.Println("Divisible by 7") ; i++
		case i%5 == 0: fmt.Println("Divisible by 5")
		case i%2 == 0: fmt.Println("Divisible by 2") ; i++
		default: fmt.Println("Weird number")
		}
	}
}
