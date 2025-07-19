package biz

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz"
	bizReq "github.com/flipped-aurora/gin-vue-admin/server/model/biz/request"
	"gorm.io/gorm"
)

type OrderService struct{}

var OrderServiceApp = new(OrderService)

func (o *OrderService) CreateOrder(db *gorm.DB, order biz.Order) (err error) {
	err = db.Create(&order).Error
	return err
}

func (o *OrderService) DeleteOrder(db *gorm.DB, order biz.Order) (err error) {
	err = db.Delete(&order).Error
	return err
}

func (o *OrderService) UpdateOrder(order biz.Order) (err error) {
	err = global.GVA_DB.Model(&biz.Order{}).Where("id = ?", order.ID).Select("meal_period", "consumer_address", "delivery_fee", "remark").Updates(&order).Error
	return err
}

func (o *OrderService) UpdateOrderStatus(db *gorm.DB, id uint) (err error) {
	err = db.Model(&biz.Order{}).Where("id = ?", id).Update("order_status", 2).Error
	return err
}

func (o *OrderService) GetOrderList(info bizReq.OrderSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	// 创建db
	db := global.GVA_DB.Model(&biz.Order{})
	var order []biz.Order

	// 微信昵称
	if info.WechatName != "" {
		db = db.Where("wechat_name LIKE ?", "%"+info.WechatName+"%")
	}

	// 消费者
	if info.ConsumerName != "" {
		db = db.Where("consumer_name LIKE ?", "%"+info.ConsumerName+"%")
	}

	// 手机号
	if info.ConsumerPhone != "" {
		db = db.Where("consumer_phone = ?", info.ConsumerPhone)
	}

	// 订单日期
	if info.OrderDate != nil && !info.OrderDate.IsZero() {
		orderDateStr := info.OrderDate.In(time.Local).Format("2006-01-02")
		db = db.Where("DATE(order_date) = ?", orderDateStr)
	}

	db = db.Order("created_at desc")

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&order).Error
	return order, total, err
}

func (c *OrderService) GetOrder(db *gorm.DB, id uint) (order biz.Order, err error) {
	err = db.Where("id = ?", id).First(&order).Error
	return
}

func (o *OrderService) GetOrdersByDate(info bizReq.OrderPrintSearch) (list interface{}, err error) {
	db := global.GVA_DB.Model(&biz.Order{})

	orderDateStr := info.OrderDate.In(time.Local).Format("2006-01-02")
	db = db.Where("DATE(order_date) = ?", orderDateStr)

	db = db.Order("created_at desc")

	var order []biz.Order

	err = db.Find(&order).Error
	return order, err
}
