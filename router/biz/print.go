package biz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PrintRouter struct{}

func (p *PrintRouter) InitPrintRouter(Router *gin.RouterGroup) {
	printRouter := Router.Group("print").Use(middleware.OperationRecord())
	printRouterWithoutRecord := Router.Group("print")
	{
		// 创建打印
		printRouter.POST("create", printApi.CreatePrint)
		// 停止打印
		printRouter.POST("stop", printApi.StopPrint)
	}
	{
		// 获取客户列表
		printRouterWithoutRecord.GET("list", printApi.GetPrintList)
		// 打印
		printRouterWithoutRecord.GET("printer", printApi.PrinterData)
	}
}
