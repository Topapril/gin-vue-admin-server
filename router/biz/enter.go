package biz

import (
	api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
)

type RouterGroup struct {
	ConsumerRouter
	BusinessRouter
	OrderRouter
	PrintRouter
	ReservationRouter
}

var (
	consumerApi    = api.ApiGroupApp.BizApiGroup.ConsumerApi
	orderApi       = api.ApiGroupApp.BizApiGroup.OrderApi
	businessApi    = api.ApiGroupApp.BizApiGroup.BusinessApi
	printApi       = api.ApiGroupApp.BizApiGroup.PrintApi
	reservationApi = api.ApiGroupApp.BizApiGroup.ReservationApi
)
