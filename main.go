package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"go-gin-example/pkg/setting"
	"go-gin-example/routers"
	"log"
	"syscall"
)

func main() {
	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

	// endless.NewServer 返回一个初始化的 endlessServer 对象
	server := endless.NewServer(endPoint, routers.InitRouter())
	// 在BeforeBegin 时输出当前进程的 pid ，调用 ListenAndServe 将实际“启动”服 务
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}

	//r := routers.InitRouter()
	//
	//s := &http.Server{
	//	Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
	//	Handler:        r,
	//	ReadTimeout:    setting.ReadTimeout,
	//	WriteTimeout:   setting.WriteTimeout,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//s.ListenAndServe()
}
