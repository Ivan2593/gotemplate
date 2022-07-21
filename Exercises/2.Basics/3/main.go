package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(stringStat("съешь ещё этих мягких французских булок, да выпей чаю"))
}

func stringStat(word string) string {
	symbolsMap := make(map[rune]int, 0)
	var keys []rune
	word = strings.ToLower(word)
	str := ""
	for _, value := range word {
		if string(value) != " " {
			symbolsMap[value]++
			if symbolsMap[value] == 1 {
				keys = append(keys, value)
			}
		}
	}
	for _, value := range keys {
		str = str + string(value) + " - " + strconv.Itoa(symbolsMap[value]) + "\n"
	}
	str = str[0 : len(str)-1]
	return str
}
