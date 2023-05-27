package piscine

import "fmt"

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
