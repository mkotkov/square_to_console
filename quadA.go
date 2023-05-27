package piscine

import "fmt"
import "strings"

func QuadA(x,y int) {

	if x > 0 && y > 0 {

		if x == 1 {
			fmt.Println ("o")
		} else {
			fmt.Println ("o" + strings.Repeat("-", (x-2)) + "o")
		}
		for i := 0 ; i < y-2 ; i++ {
			fmt.Println ("|" + strings.Repeat(" ", (x-2)) + "|")
		}
		}
	
	}
		

func main() {
	QuadA(1,1)
}
