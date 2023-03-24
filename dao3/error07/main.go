// Go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理。

// panic：
// 1、内置函数
// 2、假如函数F中书写了panic语句，会终止其后要执行的代码，在panic所在函数F内如果存在要执行的defer函数列表，按照defer的逆序执行
// 3、返回函数F的调用者G，在G中，调用函数F语句之后的代码不会执行，假如函数G中存在要执行的defer函数列表，按照defer的逆序执行
// 4、直到goroutine整个退出，并报告错误

// recover：
// 1、内置函数
// 2、用来控制一个goroutine的panicking行为，捕获panic，从而影响应用的行为
// 3、一般的调用建议
// 	a). 在defer函数中，通过recever来终止一个goroutine的panicking过程，从而恢复正常代码的执行
// 	b). 可以获取通过panic传递的error

// 1.利用recover处理panic指令，defer 必须放在 panic 之前定义，另外 recover 只有在 defer 调用的函数中才有效。否则当panic时，recover无法捕获到panic，无法防止panic扩散。
// 2.recover 处理异常后，逻辑并不会恢复到 panic 那个点去，函数跑到 defer 之后的那个点。
// 3.多个 defer 会形成 defer 栈，后定义的 defer 语句会被最先调用。

// 导致关键流程出现不可修复性错误的使用 panic，其他使用 error。
package main

import "fmt"

func main() {
	test()

	deferError()

	try(func() {
		panic("try test panic")
	}, func(err interface{}) {
		fmt.Println(err)
	})
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			println(err.(string))
		}
	}()

	panic("panic error!")
}

// 延迟调用中引发的错误，可被后续延迟调用捕获，但仅最后一个错误可被捕获。
// 捕获函数 recover 只有在延迟调用内直接调用才会终止错误，否则总是返回 nil
func deferError() {
	defer func() {
		fmt.Println(recover())
	}()

	defer func() {
		panic("defer panic")
	}()

	panic("test panic")
}

// 实现类似 try catch 的异常处理
func try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}
