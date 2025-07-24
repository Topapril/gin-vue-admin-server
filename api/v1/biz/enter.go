package biz

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	ConsumerApi
	BusinessApi
	OrderApi
	PrintApi
	ReservationApi
}

var (
	consumerService    = service.ServiceGroupApp.BizServiceGroup.ConsumerService
	orderService       = service.ServiceGroupApp.BizServiceGroup.OrderService
	businessService    = service.ServiceGroupApp.BizServiceGroup.BusinessService
	printService       = service.ServiceGroupApp.BizServiceGroup.PrintService
	reservationService = service.ServiceGroupApp.BizServiceGroup.ReservationService
)
