package main

import (
	"context"
	"fmt"
)

func main() {

	context.Context()
	fmt.Println(deferFuncReturn())
}

func deferFuncReturn() (result int) {
	i := 1
	defer func() {
		result++
	}()

	return i
}
