package main

import (
	"errors"
	"fmt"
)

func findLeft(array []uint) (error, uint) {
	for i, v := range array {
		if i == len(array)-1 {
			break
		}
		if array[i+1] < v {
			if i > 0 {
				return nil, uint(i+1)
			}
			return nil, uint(i)
		}
	}
	return errors.New("Nothing"), 0
}

func findUnsortedSubarray(array []uint) (left uint, right uint) {
	err, left := findLeft(array)
	if err != nil {
		fmt.Println("Can't get LEFT")
		return 0, 0
	}
	if left == 0 {
		for i, v := range array {
			if i == len(array)-1 {
				break
			}
			if array[i+1] > v {
				return left, uint(i)
			}
		}
	}
	return left, uint(len(array)-1)
}

func main() {
	var arrayEnd = []uint{9, 10, 11, 12, 13, 10, 5, 1, 2}
	var arrayFirst = []uint{10, 5, 4, 2, 1, 9, 10, 11, 12, 13}
	fmt.Println(findUnsortedSubarray(arrayEnd))
	fmt.Println(findUnsortedSubarray(arrayFirst))
}