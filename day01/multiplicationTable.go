package main

import "fmt"

func main() {
	/*
		乘法口诀
		1 * 1 = 1
		2 * 1 = 2 	2 * 2 = 4
		3 * 1 = 3	3 * 2 = 6	3 * 3 = 9
		...
	*/
	for i := 1; i < 10; i++ {
		// fmt.Println(i)
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %2d\t", i, j, i*j)
			if i == j {
				fmt.Println()
			}
		}
	}
}
