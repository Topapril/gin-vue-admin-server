package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ReservationSearch struct {
	WechatName        string     `json:"wechatName" form:"wechatName"`
	ConsumerName      string     `json:"consumerName" form:"consumerName"`
	ConsumerPhone     string     `json:"consumerPhone" form:"consumerPhone"`
	ReservationDate   *time.Time `json:"reservationDate" form:"reservationDate"`
	ReservationStatus uint       `json:"reservationStatus" form:"reservationStatus"`
	request.PageInfo
}
