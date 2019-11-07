package main

import (
	"fmt"
	"gitTest/calc"
)

func main() {
	fmt.Println("add called!")
	res := calc.Add(10,20)
	fmt.Println("Add(10,20): ", res)
}
