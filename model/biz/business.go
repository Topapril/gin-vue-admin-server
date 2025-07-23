package biz

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type Business struct {
	global.GVA_MODEL
	BusinessDate      *time.Time `json:"businessDate"`
	LunchName         string     `json:"lunchName"`
	LunchDescription  string     `json:"lunchDescription"`
	DinnerName        string     `json:"dinnerName"`
	DinnerDescription string     `json:"dinnerDescription"`
	FestivalGift      string     `json:"festivalGift"`
	BusinessStatus    uint       `json:"businessStatus"`
	Visible           uint       `json:"visible"`
	Remark            string     `json:"remark"`
}

func (Business) TableName() string {
	return "biz_business"
}
