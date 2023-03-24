package main

import (
	"fmt"

	"gobook.com/hello/ch2/tempconv"
)

func main() {
	x := 1
	p := &x
	y := x
	fmt.Printf("start p=%d x=%d y=%d\n", *p, x, y)
	*p = 2
	y = 3
	fmt.Printf("end p=%d x=%d y=%d\n", *p, x, y)

	n := new(int)
	fmt.Println(*n)
	*n = 2
	fmt.Println(*n)

	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC)
	fmt.Println(tempconv.CToF(tempconv.BoilingC))

	t := 32
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)

	fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
}
