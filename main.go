package main

import (
	"fmt"
	"gitTest/calc"
)

func main() {
	fmt.Println("calc add called!")
	res := calc.Add(10,20)
	fmt.Println("Add(10,20): ", res)
	fmt.Println("calc sub called!")
	res = calc.Sub(30,20)
	fmt.Println("Sub(10,20): ", res)
}
