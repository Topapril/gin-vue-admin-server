package biz

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type Print struct {
	global.GVA_MODEL
	BusinessName        string     `json:"businessName"`
	GoodsName           string     `json:"goodsName"`
	GoodsQuantity       int        `json:"goodsQuantity"`
	MealPeriod          uint       `json:"mealPeriod"`
	WechatName          string     `json:"wechatName"`
	ConsumerName        string     `json:"consumerName"`
	ConsumerPhone       string     `json:"consumerPhone"`
	ConsumerAddress     string     `json:"consumerAddress"`
	OrderDate           *time.Time `json:"orderDate" gorm:"type:date"`
	BusinessQrCode      string     `json:"businessQrCode"`
	BusinessQrCodeTitle string     `json:"businessQrCodeTitle"`
	BusinessPhone       string     `json:"businessPhone"`
	PrintResult         string     `json:"printResult"`
	PrintStatus         uint       `json:"printStatus"`
	Remark              string     `json:"remark"`
}

func (Print) TableName() string {
	return "biz_print"
}
