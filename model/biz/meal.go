package biz

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type Meal struct {
	global.GVA_MODEL
	BusinessDate      *time.Time `json:"businessDate" time_format:"2006-01-02" gorm:"type:date"`
	LunchName         string     `json:"lunchName"`
	LunchDescription  string     `json:"lunchDescription"`
	DinnerName        string     `json:"dinnerName"`
	DinnerDescription string     `json:"dinnerDescription"`
	FestivalGift      string     `json:"festivalGift"`
	BusinessStatus    uint       `json:"businessStatus"`
	Visible           uint       `json:"visible"`
	Remark            string     `json:"remark"`
}

func (Meal) TableName() string {
	return "biz_meal"
}
