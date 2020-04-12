package main

import (
	"fmt"

	"github.com/ikngtty/golib/io"
)

func main() {
	sc := io.NewScanner()

	a := sc.ReadInt()
	b := sc.ReadInt()
	c := sc.ReadInt()
	fmt.Println(a + b + c) // 60

	d := sc.ReadInt64()
	e := sc.ReadInt64()
	fmt.Println(d + e) // 90

	x := sc.ReadString()
	y := sc.ReadString()
	fmt.Println(x) // abc
	fmt.Println(y) // def
}
