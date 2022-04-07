package main

import (
	//people "eirc_test/people"
	"encoding/json"
	"fmt"
	"log"
	//"encoding/json"
	api "eirc_test/api"//resty實做

	//"github.com/go-resty/resty/v2" // 安裝resty
)

/*
var x, y, z int = 1, 2, 3
var (
	a int    = 4
	b string = `hi`
	c int    = 5
)

type students struct {
	Name   string `json:"name"`
	Number string "json:number"
}
*/
// var ()裡面宣告這種資料型態 不用逗號 := 也是宣告但不能擺在func外面，可以宣告任何型態

func main() {
	//d := "wow"
	/*
	amy := students{
		Name:   "amy",
		Number: "C109118211", //用于生成json字符串
	}

	vin := students{}
	vin.Name = "vin"
	vin.Number = "C109118211" //指针變量

	class := []students{vin, amy}
	class2 := []students{}

	for index, value := range class {
		class2 := append(class2, value) //append有連接的意思
		fmt.Println(index, class2)
	}
	//序列化
	marshal, err := json.Marshal(class)
	if err != nil { //nil是空值的意思,Marshal失败时err!=nil
		fmt.Println(err)
	}
	//[]byte
	//沒有用到用_

	//反序列化
	class3 := []students{}
	err = json.Unmarshal(marshal, &class3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(class3)

	//一個{}代表一個struct，[]在前面是GO用法

	//fmt.Println("hello", x, y, z, a, b, c, d)

	var a = 7.98
	var p = &a
	var pp = &p
	fmt.Println("a = ", a)
	fmt.Println("address of a = ", &a)
	fmt.Println("p=",p)
	fmt.Println("address of p= ", &p)
	fmt.Println("pp = ",pp)
	// Dereferencing a pointer to pointer
	fmt.Println("*pp=",*pp)
	fmt.Println("**pp=",**pp)
	
	//使用本地package
	people.Class()
	//fmt.Println(people.Student1)
	people.ClassName("Hello")
	result, num := people.Vscode(3)
	fmt.Println(result, num)
	
	
	
	fmt.Println(api.CompanyLogin("445da1f7"))
	log.Println("hello log")
	
	client := resty.New()//創建一個resty client
	resp, err := client.R().//點要連接要最後,表示方法延續
	Get("https://www.google.com.tw/?hl=zh_TW")

  	if err != nil {
    	log.Fatal(err)//log也可print;fatal 會做以下兩點
	//1.打印輸出內容 2.退出應用程序
  	}

  	fmt.Println("Response Info:")//響應基本信息
  	fmt.Println("Status Code:", resp.StatusCode())//狀態碼
	fmt.Println("Proto:",resp.Proto())//協議
  	fmt.Println("Status:", resp.Status())//狀態碼和狀態信息
  	fmt.Println("Time:", resp.Time())//從發送請求到收到響應的時間
  	fmt.Println("Received At:", resp.ReceivedAt())//接收到的時刻
*/
//resty post
	job := api.POST{
		Name: "hello",
		CompanyHash: "17C99EADD5A77BD5117719B12478E7",
	}

	joblist :=[]api.POST{job}
	//序列化
	marshal, err :=json.Marshal(joblist)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(marshal))
	//marshal為[]byte	
	
	fmt.Println(api.POST_job(string(marshal)))
	fmt.Println(api.PUT_job())
	fmt.Println(api.Delete_job("3014"))
}

