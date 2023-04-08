package main

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

//https://github.com/vv314/quotes
// 扇贝单词
// API：https://apiv3.shanbay.com/weapps/dailyquote/quote
// 请求方法：GET
// 调性：励志，鸡汤
// 类型：图文

type Quote struct {
	Name string      `json:"name"`
	Form string      `json:"form"`
	Data []QuoteInfo `json:"data"`
}

type QuoteInfo struct {
	Id            string   `json:"id"`
	Content       string   `json:"content"`
	Author        string   `json:"author"`
	AssignDate    string   `json:"assign_date"`
	OriginImgUrls []string `json:"origin_img_urls"`
	Translation   string   `json:"translation"`
}

func main() {
	// now := time.Now()
	// start := time.Date(2016, time.Month(10), 1, 0, 0, 0, 0, now.Location())
	// fmt.Print("later", start.Format(time.DateOnly))
	// for i := 1; i < 40; i++ {
	// 	end := start.AddDate(0, 0, i).Format(time.DateOnly)
	// 	fmt.Printf("later %s \n", end)
	// }
	// for start.Format(time.DateOnly) != now.Format(time.DateOnly) {
	// 	start = start.AddDate(0,0,1)
	// 	resp, err := http.Get(fmt.Sprintf("https://apiv3.shanbay.com/weapps/dailyquote/quote?date=%s",start.Format(time.DateOnly)))
	// 	HandleError(err, "http.Get url")
	// 	defer resp.Body.Close()
	// 	pageBytes, err := io.ReadAll(resp.Body)
	// 	HandleError(err, "io.ReadAll")
	// 	// 字节转字符串
	// 	pageStr := string(pageBytes)
	// }
	filePtr, err := os.Create("./reptile/shanbay_quote/quote.json")
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	defer filePtr.Close()

	q := &Quote{
		Name: "每日一句",
		Form: "扇贝",
		Data: make([]QuoteInfo, 0),
	}

	info := getInfo()
	q.Data = append(q.Data, info)

	data, err := json.Marshal(q)
	HandleError(err, "json Marshal")
	_, err = filePtr.Write(data)
	HandleError(err, "filePtr Write")
}

func getInfo() QuoteInfo {
	resp, err := http.Get("https://apiv3.shanbay.com/weapps/dailyquote/quote")
	HandleError(err, "http.Get url")
	defer resp.Body.Close()
	pageBytes, err := io.ReadAll(resp.Body)
	HandleError(err, "io.ReadAll")
	var info QuoteInfo
	var v interface{}
	data := v.(map[string]interface{})
	json.Unmarshal(pageBytes, &data)
	info.Id = data["id"].(string)
	info.Content = data["content"].(string)
	info.Author = data["author"].(string)
	info.AssignDate = data["assign_date"].(string)
	info.Translation = data["translation"].(string)
	imgs := data["translation"].([]string)
	fmt.Println(imgs)
	return info
}

// 处理异常
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}
