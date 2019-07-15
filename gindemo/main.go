package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//初始化引擎
	engine := gin.Default()
	//注册一个路由和处理函数
	engine.Any("/", WebRoot)
	//绑定端口，然后启动应用
	engine.Run(":9205")
}

// WebRoot 请求处理函数
// 所有本次请求的相关方法都在context中
// 输出响应 hello, world
func WebRoot(context *gin.Context) {
	context.String(http.StatusOK, "hello, world")
}
