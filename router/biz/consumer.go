package biz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ConsumerRouter struct{}

func (c *ConsumerRouter) InitConsumerRouter(Router *gin.RouterGroup) {
	consumerRouter := Router.Group("consumer").Use(middleware.OperationRecord())
	consumerRouterWithoutRecord := Router.Group("consumer")
	{
		// 创建消费者
		consumerRouter.POST("create", consumerApi.CreateConsumer)
		// 删除消费者
		consumerRouter.DELETE("delete", consumerApi.DeleteConsumer)
		// 更新消费者
		consumerRouter.PUT("update", consumerApi.UpdateConsumer)
		// 更新消费者餐次数
		consumerRouter.PUT("meal/count", consumerApi.UpdateConsumerMealCount)
		// 更新消费者状态
		consumerRouter.PUT("status", consumerApi.UpdateConsumerStatus)
	}
	{
		// 获取消费者列表
		consumerRouterWithoutRecord.GET("list", consumerApi.GetConsumerList)
		// 获取消费者
		consumerRouterWithoutRecord.GET("info", consumerApi.GetConsumer)
		// 获取消费者记录
		consumerRouterWithoutRecord.GET("record", consumerApi.GetConsumerRecordList)
	}
}
