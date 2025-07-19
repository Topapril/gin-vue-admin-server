package biz

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	ConsumerApi
	MealApi
	OrderApi
	PrintApi
}

var (
	consumerService = service.ServiceGroupApp.BizServiceGroup.ConsumerService
	orderService    = service.ServiceGroupApp.BizServiceGroup.OrderService
	mealService     = service.ServiceGroupApp.BizServiceGroup.MealService
	printService    = service.ServiceGroupApp.BizServiceGroup.PrintService
)
