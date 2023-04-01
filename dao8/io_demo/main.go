package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var fileUrl string = "./dao8/io_demo/io_demo.txt"

// os.Stdin：标准输入的文件实例，类型为*File
// os.Stdout：标准输出的文件实例，类型为*File
// os.Stderr：标准错误输出的文件实例，类型为*File
func main() {
	createFile()
	readFileTest()
	copyFileTest()

	wr()
	re()

	wrIoutil()
	reIoutil()
}

func createFile() {
	// 新建文件
	file, err := os.Create(fileUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	file.WriteString("go\n")
	file.Write([]byte("hello\n"))
}

// 读文件 file.Read()和file.ReadAt()，读到文件末尾会返回io.EOF的错误
func readFileTest() {
	// 打开文件
	file, err := os.Open(fileUrl)
	if err != nil {
		fmt.Println("open file err :", err)
		return
	}
	defer file.Close()
	// 定义接收文件读取的字节数组
	var buf [128]byte
	var content []byte
	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			// 读取结束
			break
		}
		if err != nil {
			fmt.Println("read file err ", err)
			return
		}
		content = append(content, buf[:n]...)
	}
	fmt.Println(string(content))
}

// 拷贝文件
func copyFileTest() {
	// 打开源文件
	srcFile, err := os.Open(fileUrl)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 创建新文件
	dstFile, err2 := os.Create("./dao8/io_demo/copy.txt")
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	// 缓冲读取
	buf := make([]byte, 1024)
	for {
		// 从源文件读数据
		n, err := srcFile.Read(buf)
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		//写出去
		dstFile.Write(buf[:n])
	}
	srcFile.Close()
	dstFile.Close()
}

// bufio包实现了带缓冲区的读写，是对文件读写的封装
// bufio缓冲写数据

// os.O_WRONLY	只写
// os.O_CREATE	创建文件
// os.O_RDONLY	只读
// os.O_RDWR	读写
// os.O_TRUNC	清空
// os.O_APPEND	追加

func wr() {
	// 参数2：打开模式，所有模式d都在上面
	// 参数3是权限控制
	// w写 r读 x执行   w  2   r  4   x  1
	file, err := os.OpenFile("./dao8/io_demo/bufio.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	defer file.Close()
	// 获取writer对象
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello\n")
	}
	// 刷新缓冲区，强制写出
	writer.Flush()
}

func re() {
	file, err := os.Open(fileUrl)
	if err != nil {
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
		fmt.Println(string(line))
	}

}

// 工具包写文件
// 工具包读取文件
func wrIoutil() {
	err := os.WriteFile("./dao8/io_demo/ioutil.txt", []byte("https://www.topgoer.com/%E5%B8%B8%E7%94%A8%E6%A0%87%E5%87%86%E5%BA%93/IO%E6%93%8D%E4%BD%9C.html"), 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func reIoutil() {
	content, err := os.ReadFile("./dao8/io_demo/ioutil.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
}
