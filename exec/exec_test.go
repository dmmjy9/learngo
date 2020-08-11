package main

import (
	"fmt"
	"testing"
)

func TestExec(t *testing.T) {
	c := Chicken{
		Age: 20,
	}
	Goooo(c)

	fmt.Println(c)
	Toooo(c)
}
