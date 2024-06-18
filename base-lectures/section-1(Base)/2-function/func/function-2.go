package main

import "fmt"

func main() {
	fmt.Println(findMinNum(2, 5, 7, -55, 33, -2, 44, -77, 4))
	fmt.Println(findMaxNum(2, 3, 4343, 3423423, 2323, 234234))
}

func findMaxNum(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	maxNum := numbers[0]

	for _, i := range numbers {
		if i > maxNum {
			maxNum = i
		}
	}

	return maxNum
}

func findMinNum(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	minNum := numbers[0]

	for _, i := range numbers {
		if i < minNum {
			minNum = i
		}
	}

	return minNum
}
