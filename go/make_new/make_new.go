/*
 * In The Name Of God
 * ========================================
 * [] File Name : make_new.go
 *
 * [] Creation Date : 16-11-2015
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

type parham struct {
	age int
}

func main() {
	parhamNew := new(parham)
	parhamNew.age = 22
	parhamMake := make([]parham, 10)

	fmt.Println(parhamNew)
	fmt.Println(parhamMake)
}
