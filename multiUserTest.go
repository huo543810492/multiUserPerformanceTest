package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

// 定义用户结构体
type User struct {
	Cookie string // 用户的Cookie
}

// 发送请求的函数
func sendRequest(url string, user User, wg *sync.WaitGroup) {
	defer wg.Done() // 在协程完成时调用Done

	// 定义HTTP客户端
	client := &http.Client{}

	// 创建HTTP请求
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	// 添加Cookie到请求头
	req.Header.Add("Cookie", user.Cookie)
	req.Header.Add("Content-Type", "application/json")

	// 发送请求并测量时间
	startTime := time.Now()
	response, err := client.Do(req)
	duration := time.Since(startTime)

	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return
	}
	defer response.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	// 打印请求所用时间和响应内容
	fmt.Printf("Request completed in %v\nResponse: %s\n", duration, body)
}

func main() {
	url := "http://test-ssg.bskbacekfjd.com/api/order/open?identity=identity=LNH29klnV09II9Au&commodity=STRUSDT_CC&buy=true&type=0&srcCurrency=USDT&desCurrency=STRK&price=1.910&totalAmount=11&volume=5.6&platform=pc&chargeEagle=0" // 替换成您要压测的URL
	cookie1 := "G_ENABLED_IDPS=google; _ga_TB7FT7C911=GS1.1.1711096194.199.1.1711096353.60.0.0; _ga=GA1.2.680037802.1677118638; TOKEN=3cdf5ba4-e838-4fce-b914-343c06870b1a; JSESSIONID=9626FA033B4370D979155E8DE988C90F"              // 替换成您要添加的Cookie值
	cookie2 := "JSESSIONID=6557A3D15D118AB42965440707251D22; TOKEN=bd67c90d-07a6-4f8b-9fa6-d033ffca6d3d"
	cookie3 := "TOKEN=fe2ac812-911f-4b0c-8ca5-0645e4859107; JSESSIONID=49A08E98D94990B6FAA86722F5650B90"
	cookie4 := "TOKEN=bd67c90d-07a6-4f8b-9fa6-d033ffca6d3d; JSESSIONID=1F4E695E7C574C80FB9F2D1695D80F11"
	cookie5 := "TOKEN=ca874cfc-4890-4a78-b9b0-beeb5c1e04b5; JSESSIONID=3AD17119BC8D333F86357ABCE6DCDF70"
	requestCount := 100 // 发送100次请求
	userCount := 5      // 同时发送请求的用户数

	// 创建用户切片
	users := make([]User, userCount)
	users[0] = User{Cookie: cookie1}
	users[1] = User{Cookie: cookie2}
	users[2] = User{Cookie: cookie3}
	users[3] = User{Cookie: cookie4}
	users[4] = User{Cookie: cookie5}

	// 获取当前时间
	startTime := time.Now()
	fmt.Println("开始时间:", startTime)

	// 使用WaitGroup等待所有goroutine完成
	var wg sync.WaitGroup

	// 每个用户发送请求
	for _, user := range users {
		for i := 0; i < requestCount; i++ {
			wg.Add(1)                      // 移动到这里，为每个请求添加计数
			go sendRequest(url, user, &wg) // 注意：这里应该使用goroutine启动sendRequest
		}
	}

	// 等待所有请求完成
	wg.Wait()

	// 获取当前时间
	endTime := time.Now()
	fmt.Println("结束时间:", endTime)
}
