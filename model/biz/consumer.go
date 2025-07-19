package biz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/google/uuid"
)

type Consumer struct {
	global.GVA_MODEL
	ConsumerId      uuid.UUID `json:"consumerId"`
	WechatName      string    `json:"wechatName"`
	ConsumerName    string    `json:"consumerName"`
	ConsumerPhone   string    `json:"consumerPhone"`
	ConsumerAddress string    `json:"consumerAddress"`
	DeliveryFee     float64   `json:"deliveryFee"`
	TotalMealCount  int       `json:"totalMealCount"`
	UsageCount      int       `json:"usageCount"`
	ConsumerChannel uint      `json:"consumerChannel"`
	ConsumerStatus  uint      `json:"consumerStatus"`
	Remark          string    `json:"remark"`
}

func (Consumer) TableName() string {
	return "biz_consumer"
}
