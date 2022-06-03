package main

import "github.com/goldenBill/douyin-fighting/initialize"

func main() {
	initialize.Global() //初始化全局变量
	initialize.MySQL()  //初始化 MySQL 连接
	initialize.Redis()
	initialize.Router() //初始化 GinRouter
}
