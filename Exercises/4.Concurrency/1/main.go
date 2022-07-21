package main

import (
	"fmt"
)

const panicMessage = "panic happened" // вставьте меня в сообщение паники

func main() {
	fmt.Println(foo())
}

func foo() (err error) {
	defer func() {
		e := recover()
		if e != nil {
			fmt.Println(err)
			err = fmt.Errorf("error with %v", e)
			return
		}
	}()
	panic(panicMessage)
	return nil
}
