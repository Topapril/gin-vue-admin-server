package biz

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz"
	bizReq "github.com/flipped-aurora/gin-vue-admin/server/model/biz/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type MealApi struct{}

func (m *MealApi) CreateMeal(c *gin.Context) {
	var meal biz.Meal
	err := c.ShouldBindJSON(&meal)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 校验参数
	err = utils.Verify(meal, utils.MealVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 查询营业日期重复
	_, err = mealService.GetMealByDate(meal.BusinessDate)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.GVA_LOG.Error("查询失败", zap.Error(err))
		response.FailWithMessage("查询失败: "+err.Error(), c)
		return
	}

	// 日期存在
	if err == nil {
		global.GVA_LOG.Error("营业日期已存在")
		response.FailWithMessage("营业日期已存在", c)
		return
	}

	// 创建餐品
	err = mealService.CreateMeal(meal)
	if err != nil {
		global.GVA_LOG.Error("创建失败", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}

	response.OkWithMessage("创建成功", c)
}

func (m *MealApi) DeleteMeal(c *gin.Context) {
	var meal biz.Meal
	err := c.ShouldBindJSON(&meal)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = utils.Verify(meal.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = mealService.DeleteMeal(meal)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

func (m *MealApi) UpdateMeal(c *gin.Context) {
	var meal biz.Meal
	err := c.ShouldBindJSON(&meal)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = mealService.UpdateMeal(meal)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

func (m *MealApi) UpdateMealStatus(c *gin.Context) {
	var meal biz.Meal
	err := c.ShouldBindJSON(&meal)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = mealService.UpdateStatus(meal)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

func (m *MealApi) GetMealList(c *gin.Context) {
	var pageInfo bizReq.MealSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	mealList, total, err := mealService.GetMealList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     mealList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

func (m *MealApi) GetMeal(c *gin.Context) {
	var meal biz.Meal
	err := c.ShouldBindQuery(&meal)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = utils.Verify(meal.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	data, err := mealService.GetMeal(meal.ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithDetailed(data, "获取成功", c)
}
