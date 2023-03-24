package main

import "fmt"

func main() {

	//start 结果是溢出
	//超出的高位的bit位部分将被丢弃。如果原始的数值是有符号类型，
	// 而且最左边的bit位是1的话，那么最终结果可能是负的
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)

	var i int8 = 127
	fmt.Println(i, i+1, i*i)
	//end

	var apples int32 = 1
	var apples2 int16 = 1
	var oranges int16 = 2
	var compote int = int(apples) + int(oranges)

	fmt.Println(compote, &apples, &apples2, oranges)

	//float32类型的累计计算误差很容易扩散
	var f float32 = 16777216  // 1 << 24
	var f2 float64 = 16777216 // 1 << 24
	fmt.Println(f == f+1, f2 == f2+1)
}
