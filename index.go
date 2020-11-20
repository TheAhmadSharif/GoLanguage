package main
import (
	"fmt"
	"time"
)


func addition (x int, y int) int {
	return x + y;
}
func main() {
	var text string
		text = "James Bond"
	fmt.Println(text) // ln creates new line.
	fmt.Println("Hello World")

	fmt.Println(time.Now().Second())
	fmt.Println("Addition output:", addition(2,3))

	
}
