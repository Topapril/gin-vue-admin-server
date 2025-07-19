package biz

import (
	api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
)

type RouterGroup struct {
	ConsumerRouter
	MealRouter
	OrderRouter
	PrintRouter
}

var (
	consumerApi = api.ApiGroupApp.BizApiGroup.ConsumerApi
	orderApi    = api.ApiGroupApp.BizApiGroup.OrderApi
	mealApi     = api.ApiGroupApp.BizApiGroup.MealApi
	printApi    = api.ApiGroupApp.BizApiGroup.PrintApi
)
