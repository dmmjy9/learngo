package main

import (
	"fmt"
	"io"
)

type Duck interface {
	Run()
	Walk()
}

type Chicken struct {
	Age	int
}

func (Chicken) Run() {
	fmt.Println("chicken run")
}

func (Chicken) Walk() {
	fmt.Println("chicken walk")
}

func (Chicken) String() string {
	return "this is a chicken"
}

func (Chicken) Read(p []byte) (n int, err error) {
	return 0, nil
}

func Goooo(d Duck) {
	d.Run()
	d.Walk()
}

func Toooo(r io.Reader) {
	fmt.Println("chicken reader")
}
