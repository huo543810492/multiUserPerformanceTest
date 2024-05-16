package main

//
//import (
//	"bytes"
//	"encoding/json"
//	"fmt"
//	"io/ioutil"
//	"net/http"
//	"sync"
//	"time"
//)
//
//func sendRequest(url string, cookie string, wg *sync.WaitGroup) {
//	defer wg.Done() // 在协程完成时调用Done
//
//	// 假设这是你要发送的JSON数据
//	requestData := map[string]interface{}{
//		"symbol":      "BTCUSDT_CC",
//		"type":        1, //限价
//		"side":        1,
//		"price":       30000, //市价
//		"orderQty":    1,
//		"source":      1,
//		"totalAmount": "30000",
//		"brand":       "BSK",
//	}
//
//	// 将数据编码为JSON
//	data, err := json.Marshal(requestData)
//	if err != nil {
//		fmt.Printf("Error encoding request data: %v\n", err)
//		return
//	}
//
//	// 定义HTTP客户端
//	client := &http.Client{}
//
//	// 创建HTTP请求
//	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
//	if err != nil {
//		fmt.Printf("Error creating request: %v\n", err)
//		return
//	}
//
//	// 添加Cookie和Content-Type到请求头
//	req.Header.Add("uid", "838372881825415168")
//	req.Header.Add("Content-Type", "application/json")
//
//	// 发送请求并测量时间
//	startTime := time.Now()
//	response, err := client.Do(req)
//	duration := time.Since(startTime)
//
//	if err != nil {
//		fmt.Printf("Error making request: %v\n", err)
//		return
//	}
//	defer response.Body.Close()
//
//	// 读取响应内容
//	body, err := ioutil.ReadAll(response.Body)
//	if err != nil {
//		fmt.Printf("Error reading response body: %v\n", err)
//		return
//	}
//
//	//打印请求所用时间和响应内容
//	fmt.Printf("Request completed in %v\nResponse: %s\n", duration, body)
//}
//
//func main() {
//	url := "http://localhost:8080/spot/order/create?symbol=STRUSDT_CC&buy=true&type=0&srcCurrency=USDT&desCurrency=BTC&price=2.156&volume=5.1" // 替换成您要压测的URL
//	cookie := "TOKEN=f2978862-c696-4b5f-a754-a4b90ec0ff95; JSESSIONID=A54A0B87D10A27C001A0E4452626C264"                                        // 替换成您要添加的Cookie值
//
//	// 定义要发送的请求次数
//	requestCount := 100 // 发送100次请求
//	var wg sync.WaitGroup
//
//	// 创建并启动多个goroutines
//	for i := 0; i < requestCount; i++ {
//		wg.Add(1)
//		go sendRequest(url, cookie, &wg)
//	}
//
//	// 等待所有请求完成
//	wg.Wait()
//	fmt.Println("All requests completed")
//}
