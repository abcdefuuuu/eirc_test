package main

import (
	"encoding/json"
	"fmt"
)

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

// var ()裡面宣告這種資料型態 不用逗號 := 也是宣告但不能擺在func外面，可以宣告任何型態

func main() {
	//d := "wow"
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
}
