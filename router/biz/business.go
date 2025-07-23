package biz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BusinessRouter struct{}

func (b *BusinessRouter) InitBusinessRouter(Router *gin.RouterGroup) {
	businessRouter := Router.Group("business").Use(middleware.OperationRecord())
	businessRouterWithoutRecord := Router.Group("business")
	{
		// 创建营业
		businessRouter.POST("create", businessApi.CreateBusiness)
		// 删除营业
		businessRouter.DELETE("delete", businessApi.DeleteBusiness)
		// 更新营业
		businessRouter.PUT("update", businessApi.UpdateBusiness)
		// 更新营业状态
		businessRouter.PUT("status", businessApi.UpdateBusinessStatus)
	}
	{
		// 获取营业列表
		businessRouterWithoutRecord.GET("list", businessApi.GetBusinessList)
		// 获取营业
		businessRouterWithoutRecord.GET("info", businessApi.GetBusiness)
	}
}
