package main

import (
	"fmt"
	"time"
)

// time包提供了时间的显示和测量用的函数
func main() {
	now := time.Now() //获取当前时间
	fmt.Printf("current time:%v\n", now)

	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	timestamp := timestampTest()
	timestampTest2(timestamp)
	timesActionTest()
	tickDemo()
	formatDemo()
}

// 基于时间对象获取时间戳
func timestampTest() (timestamp int64) {
	now := time.Now()            //获取当前时间
	timestamp1 := now.Unix()     //时间戳
	timestamp2 := now.UnixNano() //纳秒时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)
	return timestamp1
}

// 使用time.Unix()函数可以将时间戳转为时间格式
func timestampTest2(timestamp int64) {
	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
	fmt.Println(timeObj)
	year := timeObj.Year()     //年
	month := timeObj.Month()   //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //小时
	minute := timeObj.Minute() //分钟
	second := timeObj.Second() //秒
	fmt.Printf("时间戳转为时间格式: %d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

// 时间操作
func timesActionTest() {
	now := time.Now()
	// 当前时间加1小时后的时间
	later := now.Add(time.Hour)
	fmt.Println("时间加1小时: ", later)

	//求两个时间之间的差值
	fmt.Println("时间差值: ", later.Sub(now))

	//判断两个时间是否相同
	later2 := now.Add(time.Hour)
	fmt.Println("时间是否相同: ", later.Equal(now), later.Equal(later2))

	//如果t代表的时间点在u之前
	fmt.Println("Before: ", later.Before(now), now.Before(later))

	//如果t代表的时间点在u之后
	fmt.Println("After: ", later.After(now), now.After(later))
}

func tickDemo() {
	// ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器
	// for i := range ticker {
	// 	fmt.Println(i) //每秒都会执行的任务
	// }

	newTicker := time.NewTicker(time.Second)
	defer newTicker.Stop()
	go func(t *time.Ticker) {
		for {
			// 每1秒中从chan t.C 中读取一次
			<-t.C
			fmt.Println("Ticker:", time.Now().Format("2006-01-02 15:04:05"))
		}
	}(newTicker)
	time.Sleep(6 * time.Second)
	fmt.Println("定时任务完成")
}

// 时间格式化 Go的诞生时间2006年1月2号15点04分（记忆口诀为2006 1 2 3 4）
func formatDemo() {
	now := time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}
