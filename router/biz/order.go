package biz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OrderRouter struct{}

func (o *OrderRouter) InitOrderRouter(Router *gin.RouterGroup) {
	orderRouter := Router.Group("order").Use(middleware.OperationRecord())
	orderRouterWithoutRecord := Router.Group("order")
	{
		// 创建订单
		orderRouter.POST("create", orderApi.CreateOrder)
		// 撤销订单
		orderRouter.PUT("revoke", orderApi.RevokeOrder)
		// 更新订单
		orderRouter.PUT("update", orderApi.UpdateOrder)
		// 打印订单
		orderRouter.POST("printer", orderApi.PrintOrder)
	}
	{
		// 获取订单列表
		orderRouterWithoutRecord.GET("list", orderApi.GetOrderList)
		// 获取订单
		orderRouterWithoutRecord.GET("info", orderApi.GetOrder)
		// 打印日订单
		orderRouterWithoutRecord.GET("day/printer", orderApi.PrintDayOrders)
		// 生成订单
		orderRouterWithoutRecord.GET("generate", orderApi.GenerateOrder)
	}
}
