package biz

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type Reservation struct {
	global.GVA_MODEL
	GoodsName         string     `json:"goodsName"`
	GoodsQuantity     int        `json:"goodsQuantity"`
	MealPeriod        uint       `json:"mealPeriod"`
	WechatName        string     `json:"wechatName"`
	ConsumerId        string     `json:"consumerId"`
	ConsumerName      string     `json:"consumerName"`
	ConsumerPhone     string     `json:"consumerPhone"`
	ConsumerAddress   string     `json:"consumerAddress"`
	ReservationDate   *time.Time `json:"reservationDate" gorm:"type:date"`
	ReservationStatus uint       `json:"reservationStatus"`
	Remark            string     `json:"remark"`
}

func (Reservation) TableName() string {
	return "biz_reservation"
}
