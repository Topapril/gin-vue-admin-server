package biz

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type Order struct {
	global.GVA_MODEL
	GoodsName       string     `json:"goodsName"`
	GoodsQuantity   int        `json:"goodsQuantity"`
	MealPeriod      uint       `json:"mealPeriod"`
	WechatName      string     `json:"wechatName"`
	ConsumerId      string     `json:"consumerId"`
	ConsumerName    string     `json:"consumerName"`
	ConsumerPhone   string     `json:"consumerPhone"`
	ConsumerAddress string     `json:"consumerAddress"`
	DeliveryFee     float64    `json:"deliveryFee"`
	OrderDate       *time.Time `json:"orderDate" gorm:"type:date"`
	OrderSource     uint       `json:"orderSource"`
	OrderStatus     uint       `json:"orderStatus"`
	Remark          string     `json:"remark"`
}

func (Order) TableName() string {
	return "biz_order"
}
