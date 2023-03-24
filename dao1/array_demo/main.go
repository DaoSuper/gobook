// 1. 数组：是同一种数据类型的固定长度的序列。
// 2. 数组定义：var a [len]int，比如：var a [5]int，数组长度必须是常量，且是类型的组成部分。一旦定义，长度不能变。
// 3. 长度是数组类型的一部分，因此，var a[5] int和var a[10]int是不同的类型。
// 4. 数组可以通过下标进行访问，下标是从0开始，最后一个元素下标是：len-1
// for i := 0; i < len(a); i++ {
// }
// for index, v := range a {
// }
// 5. 访问越界，如果下标在数组合法范围之外，则触发访问越界，会panic
// 6. 数组是值类型，赋值和传参会复制整个数组，而不是指针。因此改变副本的值，不会改变本身的值。
// 7.支持 "=="、"!=" 操作符，因为内存总是被初始化过的。
// 8.指针数组 [n]*T，数组指针 *[n]T。
package main

import (
	"fmt"
)

var arr0 [5]int = [5]int{1, 2, 3}
var arr1 = [...]string{"a", "b", "c", "d"}
var str = [5]string{3: "hello world", 4: "tom"}

// 多维数组
var muArr0 [5][3]int
var muArr1 [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}

func main() {
	a := [3]int{1, 2}               // 未初始化元素值为 0
	b := [...]string{"a", "b", "c"} // 通过初始化值确定数组长度
	c := [5]int{2: 100, 4: 200}     // 使用索引号初始化元素。
	d := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 10}, // 可省略元素类型。
		{"user2", 20}, // 别忘了最后一行的逗号。
	}
	fmt.Println("数组", arr0, arr1, str)
	fmt.Println("数组", a, b, c, d)

	//多维数组
	muA := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	muB := [...][2]int{{1, 1}, {2, 2}, {3, 3}} // 第 2 纬度不能用 "..."。
	fmt.Println("多维数组", muArr0, muArr1)
	fmt.Println("多维数组", muA, muB)

	// len 和 cap 都返回数组长度 (元素数量)
	println("数组长度", len(a), cap(a))

	muForDemo()

	copyArr()

	//求数组所有元素之和
	test := [5]int{1, 3, 5, 8, 7}
	sum := sumArr(test)
	fmt.Printf("数组：%v sum=%d\n", test, sum)

	totalIndex(test, 8)
}

// 遍历
func muForDemo() {
	var f [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}
	for k1, v1 := range f {
		for k2, v2 := range v1 {
			fmt.Printf("(%d,%d)=%d ", k1, k2, v2)
		}
		fmt.Println()
	}
}

// 数组拷贝和传参
func copyArr() {
	var arr1 [5]int
	printArr(&arr1)
	fmt.Println(arr1)
	arr2 := [...]int{2, 4, 6, 8, 10}
	printArr(&arr2)
	fmt.Println(arr2)
}

func printArr(arr *[5]int) {
	arr[0] = 10
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

// 求数组所有元素之和
func sumArr(a [5]int) int {
	var sum int = 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum
}

// 找出数组中和为给定值的两个元素的下标，例如数组[1,3,5,8,7]，
// 找出两个元素之和等于8的下标分别是（0，4）和（1，2）
// 求元素和，是给定的值
func totalIndex(arr [5]int, target int) {
	fmt.Printf("数组中和为%d的两个元素的下标 %v\n", target, arr)
	for i := 0; i < len(arr); i++ {
		other := target - arr[i]
		for j := i + 1; j < len(arr); j++ {
			if arr[j] == other {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
}
