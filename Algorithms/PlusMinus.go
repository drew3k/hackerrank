package Algorithms

import "fmt"

func plusMinus(arr []int32) {
	sum := float64(len(arr))
	var pos, neg, zero float64

	for _, num := range arr {
		if num > 0 {
			pos++
		} else if num == 0 {
			zero++
		} else {
			neg++
		}
	}
	posC := pos / sum
	negC := neg / sum
	zeroC := zero / sum

	fmt.Printf("%.6f\n", posC)
	fmt.Printf("%.6f\n", negC)
	fmt.Printf("%.6f\n", zeroC)
}
