package biz

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type ConsumerRecord struct {
	global.GVA_MODEL
	WechatName      string     `json:"wechatName"`
	ConsumerId      string     `json:"consumerId"`
	ConsumerName    string     `json:"consumerName"`
	MealPeriod      uint       `json:"mealPeriod"`
	UsageCount      int        `json:"usageCount"`
	OrderDate       *time.Time `json:"orderDate" gorm:"type:date"`
	TransactionType uint       `json:"transactionType"`
	Description     string     `json:"description"`
	Remark          string     `json:"remark"`
}

func (ConsumerRecord) TableName() string {
	return "biz_consumer_record"
}
