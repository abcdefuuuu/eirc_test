package localpackage

import (
	"fmt"
)

var (
	TeamManber1 = "林姿穎" //🈯從別的檔案移到另一個檔案(不管是宣告還是func)要用大寫
	TeamManber2 = "林姿一"
	TeamManber3 = "0-1"
)

func Room() {
	fmt.Println(TeamManber1, TeamManber2, TeamManber3)
}
func init() {
	fmt.Println("我會先執行")
}
func RoomCode(Code string) {
	fmt.Println("c109118211")
}
func Number(num int) (int, string) {
	if num > 10 {
		return num, " 大於10"
	}
	return num, " 小於10"
}
