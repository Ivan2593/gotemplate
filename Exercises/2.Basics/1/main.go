package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(doubleDetector(nums))
}

func doubleDetector(nums []int) bool {
	numsMap := make(map[int]int, 0)
	for _, value := range nums {
		numsMap[value]++
		if numsMap[value] > 1 {
			return true
		}
	}
	return false
}
