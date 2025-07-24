package biz

import (
	"errors"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz"
	bizReq "github.com/flipped-aurora/gin-vue-admin/server/model/biz/request"
)

type BusinessService struct{}

var BusinessServiceApp = new(BusinessService)

func (b *BusinessService) CreateBusiness(business biz.Business) (err error) {
	err = global.GVA_DB.Create(&business).Error
	return err
}

func (b *BusinessService) DeleteBusiness(business biz.Business) (err error) {
	err = global.GVA_DB.Delete(&business).Error
	return err
}

func (b *BusinessService) UpdateBusiness(business biz.Business) (err error) {
	err = global.GVA_DB.Model(&biz.Business{}).Where("id = ?", business.ID).Updates(&business).Error
	return err
}

func (b *BusinessService) UpdateBusinessStatus(business biz.Business) (err error) {
	err = global.GVA_DB.Model(&biz.Business{}).Where("id = ?", business.ID).Update("visible", business.Visible).Error
	return err
}

func (b *BusinessService) GetBusinessList(info bizReq.BusinessSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	// 创建db
	db := global.GVA_DB.Model(&biz.Business{})
	var business []biz.Business

	// 营业日期
	if info.BusinessDate != nil && !info.BusinessDate.IsZero() {
		// 前端传的是UTC时间，后端需转为本地时间
		businessDateStr := info.BusinessDate.In(time.Local).Format("2006-01-02")
		db = db.Where("DATE(business_date) = ?", businessDateStr)
	}

	// 营业状态
	if info.BusinessStatus != 0 {
		db = db.Where("business_status = ?", info.BusinessStatus)
	}

	db = db.Order("business_date desc")

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&business).Error
	return business, total, err
}

func (b *BusinessService) GetBusiness(id uint) (business biz.Business, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&business).Error
	return
}

func (b *BusinessService) HasDuplicateBusinessDate(businessDate *time.Time) (count int64, err error) {
	if businessDate == nil {
		return 0, errors.New("营业日期错误")
	}

	businessDateStr := businessDate.In(time.Local).Format("2006-01-02")
	err = global.GVA_DB.Model(&biz.Business{}).Where("DATE(business_date) = ?", businessDateStr).Count(&count).Error
	return count, err
}

func (b *BusinessService) FetchBusinessForToday() (business biz.Business, err error) {
	businessDate := time.Now().Format("2006-01-02")
	err = global.GVA_DB.Where("DATE(business_date) = ?", businessDate).First(&business).Error
	return
}
