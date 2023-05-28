package main

import "fmt"

func main() {
	numPro := 'b'

	switch numPro {
	case 'a':
		QuadA(5, 3)
	case 'b':
		QuadB(5, 3)
	case 'c':
		QuadC(5, 3)
	case 'd':
		QuadD(5, 3)
	case 'e':
		QuadE(5, 3)
	}
}

func QuadA(x, y int) {
	if x > 0 && y > 0 {
		for i := 0; i < y; i++ {
			for j := 0; j < x; j++ {
				if i == 0 || i == y-1 {
					if j == 0 || j == x-1 {
						fmt.Print("o")
					} else {
						fmt.Print("-")
					}
				} else {
					if j == 0 || j == x-1 {
						fmt.Print("|")
					} else {
						fmt.Print(" ")
					}
				}
			}
			fmt.Println()
		}
	}
}

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

func QuadC(x, y int) {
	if x > 0 && y > 0 {
		for i := 0; i < y; i++ {
			for j := 0; j < x; j++ {
				if i == 0 && j == 0 {
					fmt.Print("A")
				} else if i == 0 && j == x-1 {
					fmt.Print("A")
				} else if j == 0 && i == y-1 {
					fmt.Print("C")
				} else if j == x-1 && i == y-1 {
					fmt.Print("C")
				} else if i == 0 || i == y-1 {
					fmt.Print("B")
				} else if j == 0 || j == x-1 {
					fmt.Print("B")
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

func QuadD(x, y int) {
	if x > 0 && y > 0 {
		for i := 0; i < y; i++ {
			for j := 0; j < x; j++ {
				if i == 0 && (j == 0 || j == x-1) {
					fmt.Print("A")
				} else if i == y-1 && (j == 0 || j == x-1) {
					fmt.Print("A")
				} else if i == 0 || i == y-1 || j == 0 || j == x-1 {
					fmt.Print("B")
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

func QuadE(x, y int) {
	if x <= 0 || y <= 0 {
		return // if x or y is negative or zero, do not print anything
	}

	for iy := 0; iy < y; iy++ { // y-axis
		if iy > 0 {
			fmt.Print("\n") // skip newline for the first row
		}

		for ix := 0; ix < x; ix++ { // x-axis
			if (iy == 0 && ix == 0) || (iy == y-1 && ix == x-1) && y != 1 && x != 1 { // upper-left and lower-right corners
				fmt.Print("A") // corners
			} else if (iy == 0 && ix == x-1) || (iy == y-1 && ix == 0) { // upper-right and lower-left corners
				fmt.Print("C") // corners
			} else if (iy == 0 || iy == y-1 && (ix != 0 || ix != x-1)) || (ix == 0 || ix == x-1 && (iy != 0 || iy != y-1)) { // top horizontal line and left vertical line
				fmt.Print("B")
			} else {
				fmt.Print(" ") // inner area
			}
		}
	}

	fmt.Print("\n") // newline at the end
}
