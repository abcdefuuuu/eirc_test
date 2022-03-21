//❗常錯 🈯:觀念 ⭐:重要觀念 🉐程式寫法
package main //🈯宣告封包的名稱，程式可以被執行一定要用main，程式如果寫給別人用就不一定要用main
import (
	"encoding/json"
	"fmt" //🈯載入內建的fmt封包，fmt用來做基本輸出輸入
)

type people struct { //json 是一種很好擴充的資料格式🉐type 結構名稱 struct{}
	Id   int    `json:id` //🈯key值要大寫才會被輸出🉐Id int...;但若是需要小寫或其他的名稱,則使用struct tag🉐Id int `json:id`
	Name string `json:name`
}

type employees struct {
	id   int
	name string
	age  int
}

type Point struct { //struct 結構  🉐type 結構名稱 struct{欄位名稱 資料型態.....}
	m int
	o int
}

//定義函式...⭐🉐func 函式名稱(參數名稱 資料型態,資料型態,....) 回傳直資料型態{函式內部的程式碼}  return 回傳值,須符合定義中的資料型態
func txt() (int, string) {
	return 5, "test"
}

//定義函式...⭐func 函式名稱(參數名稱 資料型態,資料型態,....){函式內部的程式碼};⭐呼叫函式...函式名稱(參數名稱 資料型態)
func show(msg string) {
	if msg == "Hello" {
		return
	}
	fmt.Println(msg) //❗裡面記得要print,因為之後只會呼叫副程式
}

func add(max int) {
	var up int = 0
	var h int
	for h = 0; h <= max; h++ {
		up += h
	}
	fmt.Println("使用副程式,從1加到", h-1, "是 ", up)
}

func main() { //🈯建立main函式，{是程式的進入點
	/* */ //🈯多行註解

	fmt.Println("字符a= ", 'a', 3.1415) //🈯 字符 rune 字元對應數字 a=97 b=98;🈯float64是資料型態

	var test bool = true //go的宣告變數方法....🉐var 變數名稱 資料型態 = 適當的資料
	fmt.Println(test)    //⭐go變數宣告了不能不用(一定要print出來)

	var y int
	var x int
	fmt.Println("輸入y x")
	fmt.Scanln(&y, &x)           //🉐fmt.Scanln(&變數名稱,&變數名稱,.....)  🈯&變數名稱 = 取得變數的指標(Pointer)
	fmt.Println("y * x = ", y*x) // = var result int = x * y + fmt.Println(result)

	//🈯單元運算 ++,-- ; 🈯邏輯運算 !,&&(and),||(or)

	var money int
	fmt.Println("請輸入你存的錢")
	fmt.Scanln(&money)
	if money < 10000 {
		fmt.Println(`ok`)
	} else {
		fmt.Println(`no ok`) //流程控制 if用法
	}

	var result int = 0
	var n int = 1
	for n <= 50 {
		result += n
		n++
	}
	fmt.Println(result) //流程控制 for用法

	var answer int = 0
	var a int
	for a = 0; a < 3; a++ {
		answer += a
	}
	fmt.Println(answer) //流程控制 for用法

	var b int = 0
	for b < 5 {
		if b == 5 {
			break
		}
		fmt.Println(b)
		b++ //break用法
	}

	var c int
	for c = 0; c < 5; c++ {
		if c%2 == 0 {
			continue
		}
		fmt.Println(c)
	}

	var sum int
	fmt.Println("累積數字,打0才會停止加總")
	for true {
		var d int
		fmt.Scanln(&d) //無限迴圈
		if d == 0 {
			break
		}
		sum += d
	}
	fmt.Println(sum)

	show("哈囉")
	show("你好") //⭐呼叫函式...函式名稱(參數名稱 資料型態)

	add(10)

	var g int
	var i string
	g, i = txt()
	fmt.Println(g, i)

	var z int = 3
	var zPtr *int = &z // 🈯指標變數-存放到指標變數.... *資料型態=&變數名稱
	fmt.Println(zPtr)  //🈯指標變數:指標變數名稱=會印出記憶體存放位址
	fmt.Println(*zPtr) //🈯反解指標變數:*指標變數名稱=會印出資料

	var p1 Point = Point{3, 4} //🉐1.結構名稱{欄位資料,欄位資料}🉐2.結構名稱{欄位名稱:資料,欄位名稱:資料,...}
	p1.m = 6
	fmt.Println(p1)

	fmt.Println("請輸入4人分數")
	var classmate [4]int
	for index := 0; index < len(classmate); index++ {
		fmt.Scanln(&classmate[index])
	}

	//var classmate [3]int = [3]int{60,90,100}
	var frog int
	for index := 0; index < len(classmate); index++ {
		//fmt.Println(classmate[index])
		frog += classmate[index] //簡單陣列
	}
	fmt.Println("全部加起來= ", frog)

	class := [2]employees{
		{1, "john", 33},
		{2, "mary", 23}, //struct + 陣列
	}
	fmt.Println(class)

	data := people{
		Id:   1,
		Name: "林姿穎", //用于生成json字符串
	}
	mate := people{}
	mate.Id = 2
	mate.Name = "林姿一" //指针變量

	puppy := []people{mate, data} //🉐名稱 := []結構名稱{欄位資料,欄位資料,...}
	puppy2 := []people{}
	for _, value := range puppy { //第二個變數索引值_,第二個變數是元素值value
		puppy2 = append(puppy2, value) //append有連接的意思
		fmt.Println(puppy2)
	}

	//序列化
	js, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(js))
	//反序列化
	var r people
	err = json.Unmarshal(js, &r)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
}
