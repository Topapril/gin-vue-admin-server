package biz

import (
	"errors"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz"
	bizReq "github.com/flipped-aurora/gin-vue-admin/server/model/biz/request"
)

type MealService struct{}

var MealServiceApp = new(MealService)

func (m *MealService) CreateMeal(meal biz.Meal) (err error) {
	err = global.GVA_DB.Create(&meal).Error
	return err
}

func (m *MealService) DeleteMeal(meal biz.Meal) (err error) {
	err = global.GVA_DB.Delete(&meal).Error
	return err
}

func (m *MealService) UpdateMeal(meal biz.Meal) (err error) {
	err = global.GVA_DB.Model(&biz.Meal{}).Where("id = ?", meal.ID).Updates(&meal).Error
	return err
}

func (m *MealService) UpdateStatus(meal biz.Meal) (err error) {
	err = global.GVA_DB.Model(&biz.Meal{}).Where("id = ?", meal.ID).Update("visible", meal.Visible).Error
	return err
}

func (m *MealService) GetMealList(info bizReq.MealSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	// 创建db
	db := global.GVA_DB.Model(&biz.Meal{})
	var meal []biz.Meal

	// 营业日期
	if info.BusinessDate != nil && !info.BusinessDate.IsZero() {
		businessDateStr := info.BusinessDate.Format("2006-01-02")
		db = db.Where("DATE(business_date) = ?", businessDateStr)
	}

	db = db.Order("created_at desc")

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&meal).Error
	return meal, total, err
}

func (m *MealService) GetMeal(id uint) (meal biz.Meal, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&meal).Error
	return
}

func (m *MealService) HasDuplicateBusinessDate(businessDate *time.Time) (count int64, err error) {
	if businessDate == nil {
		return 0, errors.New("营业日期错误")
	}

	err = global.GVA_DB.Model(&biz.Meal{}).Where("DATE(business_date) = ?", businessDate).Count(&count).Error
	return count, err
}

func (m *MealService) FetchMealForToday() (meal biz.Meal, err error) {
	businessDate := time.Now().Format("2006-01-02")
	err = global.GVA_DB.Where("DATE(business_date) = ?", businessDate).First(&meal).Error
	return
}
