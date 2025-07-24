package biz

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz"
	bizReq "github.com/flipped-aurora/gin-vue-admin/server/model/biz/request"
	"gorm.io/gorm"
)

type ReservationService struct{}

var ReservationServiceApp = new(ReservationService)

func (r *ReservationService) CreateReservation(db *gorm.DB, reservation biz.Reservation) (err error) {
	err = db.Create(&reservation).Error
	return err
}

func (r *ReservationService) UpdateReservation(reservation biz.Reservation) (err error) {
	err = global.GVA_DB.Model(&biz.Reservation{}).Where("id = ?", reservation.ID).Select("meal_period", "consumer_address", "delivery_fee", "remark").Updates(&reservation).Error
	return err
}

func (r *ReservationService) GetReservationList(info bizReq.ReservationSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	// 创建db
	db := global.GVA_DB.Model(&biz.Reservation{})
	var reservation []biz.Reservation

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

	// 预约日期
	if info.ReservationDate != nil && !info.ReservationDate.IsZero() {
		// 前端传的是UTC时间，后端需转为本地时间
		reservationDateStr := info.ReservationDate.In(time.Local).Format("2006-01-02")
		db = db.Where("DATE(reservation_date) = ?", reservationDateStr)
	}

	// 预约状态
	if info.ReservationStatus != 0 {
		db = db.Where("reservation_status = ?", info.ReservationStatus)
	}

	db = db.Order("created_at desc")

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&reservation).Error
	return reservation, total, err
}

func (r *ReservationService) GetReservation(id uint) (reservation biz.Reservation, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&reservation).Error
	return
}
