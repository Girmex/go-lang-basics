package main

import "fmt"

func Pointers() {
    x := 42
    p := &x        // p is a pointer to x

    fmt.Println("x =", x)
    fmt.Println("p =", p)       // address of x
    fmt.Println("*p =", *p)     // value at address p (which is x)

    *p = 100        // change the value at the memory address
    fmt.Println("x after *p = 100:", x)
}
