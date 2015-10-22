/*
 * In The Name Of God
 * ========================================
 * [] File Name : fibonacci.go
 *
 * [] Creation Date : 28-05-2015
 *
 * [] Created By : Parham Alvani (parham.alvani@gmail.com)
 * =======================================
 */
/*
 * Copyright (c) 2015 Parham Alvani.
 */
package main

import (
	"fmt"
)

func Fibonacci(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	fmt.Println(Fibonacci(8))
}