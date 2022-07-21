package main

import "fmt"

func main() {
	// Тут демонстрация работы функции
	ww := []string{"aba", "abc", "aba"}
	fmt.Printf("%v\n", isSorted(ww))
}

func isSorted(ww []string) bool {
	if ww == nil {
		return false
	}
	for i := 0; i < len(ww)-1; i++ {
		if ww[i] > ww[i+1] {
			return false
		}
	}
	return true
}
