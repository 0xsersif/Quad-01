func QuadB(x, y int) {
	//check if both values are valid
	if x <= 0 || y <= 0 {
		return
	}

	//print the upper horizontal line
	for i := 0; i < x; i++ {
		if i == 0 {
			fmt.Printf("/")
		} else if i == x-1 {
			fmt.Printf("\\")
		} else {
			fmt.Printf("*")
		}
	}
	
	

	fmt.Printf("\n")

	//print the vertical line
	for j := 0; j < y-2; j++ {
		for l := 0; l < x; l++ {
			if l == 0 || l == x-1 {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}

	//print the lower horizontal line
	if y >= 2 {

		for k := 0; k < x; k++ {
			if k == 0 {
				fmt.Printf("\\")
			} else if k == x-1 {
				fmt.Printf("/")
			} else {
				fmt.Printf("*")
			}
		}
	}
	fmt.Printf("\n")
}