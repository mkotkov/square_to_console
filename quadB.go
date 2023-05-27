package piscine

import "fmt"

func QuadB(x, y int) {
	if x > 0 && y > 0 {
		for i := 0; i < y; i++ {
			for j := 0; j < x; j++ {
				if i == 0 && j == 0 {
					fmt.Print("/")
				} else if i == 0 && j == x-1 {
					fmt.Print("\\")
				} else if j == 0 && i == y-1 {
					fmt.Print("\\")
				} else if j == x-1 && i == y-1 {
					fmt.Print("/")
				} else if i == 0 || i == y-1 {
					fmt.Print("*")
				} else if j == 0 || j == x-1 {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
	} else {
		fmt.Println("The dimensions must be positive.")
	}
}
