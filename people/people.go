package people

import "fmt"

var (
	Student1 = "vin"
	Student2 = "amy" //🈯從別的檔案移到另一個檔案(不管是宣告還是func)要用大寫
)

func Class() {
	fmt.Println(Student1, Student2)
}
func ClassName(name string) {
	fmt.Println("Hello")
}
func init() {
	fmt.Println("init") //先執行
}
func Vscode(num int) (string,int){
	if num>0{
		return "大於0",num
	}
	return "不大於0",num
}
