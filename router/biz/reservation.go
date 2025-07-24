package biz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ReservationRouter struct{}

func (r *ReservationRouter) InitReservationRouter(Router *gin.RouterGroup) {
	reservationRouter := Router.Group("reservation").Use(middleware.OperationRecord())
	reservationRouterWithoutRecord := Router.Group("reservation")
	{
		// 更新预约
		reservationRouter.PUT("update", reservationApi.UpdateReservation)
	}
	{
		// 获取预约列表
		reservationRouterWithoutRecord.GET("list", reservationApi.GetReservationList)
	}
}
