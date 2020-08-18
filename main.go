package main

import (
	"fmt"

	"github.com/ikngtty/go-contestlib/io"
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

	f := sc.ReadFloat()
	g := sc.ReadFloat()
	fmt.Println(f + g) // (around) 15.6

	x := sc.ReadString()
	y := sc.ReadString()
	fmt.Println(x) // abc
	fmt.Println(y) // def

	s := sc.ReadString()
	fmt.Println(len(s)) // 100,000
}
