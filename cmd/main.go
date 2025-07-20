package main

import (
	"fmt"
	"go-store/conf"
	"go-store/pkg/mytools"
	"go-store/routes"
)

func main() {
	fmt.Println("项目start")
	conf.Init() // 初始化项目配置
	mytools.Init()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
