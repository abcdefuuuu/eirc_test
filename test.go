package main

import (
	//"fmt"
	local "eirc_test/local_package"
)

func main(){
	//使用本地package
	local.Room()//🈯從別的檔案移到另一個檔案(不管是宣告還是func)要用大寫
	local.RoomCode("印出學號")
	num,result := local.Number(3)
	println(num,result)
}
