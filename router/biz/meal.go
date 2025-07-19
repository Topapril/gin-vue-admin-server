package biz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MealRouter struct{}

func (m *MealRouter) InitMealRouter(Router *gin.RouterGroup) {
	mealRouter := Router.Group("meal").Use(middleware.OperationRecord())
	mealRouterWithoutRecord := Router.Group("meal")
	{
		// 创建餐品
		mealRouter.POST("create", mealApi.CreateMeal)
		// 删除餐品
		mealRouter.DELETE("delete", mealApi.DeleteMeal)
		// 更新餐品
		mealRouter.PUT("update", mealApi.UpdateMeal)
		// 更新显示状态
		mealRouter.PUT("status", mealApi.UpdateMealStatus)
	}
	{
		// 获取餐品列表
		mealRouterWithoutRecord.GET("list", mealApi.GetMealList)
		// 获取餐品
		mealRouterWithoutRecord.GET("info", mealApi.GetMeal)
	}
}
