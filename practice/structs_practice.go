package main

import (
	"fmt"
	"strconv"
)

type Foo struct {
	a int
	b string
}

// embedding

type Bar struct {
	c string
	Foo
}

func main() {

	f := Foo{a: 1, b: "krishna"}

	b := Bar{c: "tripathi", Foo: f}

	fmt.Printf("Foo=> A: %d, B: %s\n", f.a, f.b)
	fmt.Printf("Bar=> A: %d, B: %s, c: %s\n", b.a, b.b, b.c)

	fmt.Println(strconv.ParseInt("11", 2, 8))

}
