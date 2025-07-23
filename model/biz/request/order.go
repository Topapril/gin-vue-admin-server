package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type OrderSearch struct {
	WechatName    string     `json:"wechatName" form:"wechatName"`
	ConsumerName  string     `json:"consumerName" form:"consumerName"`
	ConsumerPhone string     `json:"consumerPhone" form:"consumerPhone"`
	OrderDate     *time.Time `json:"orderDate" form:"orderDate"`
	request.PageInfo
}

type OrderPrintSearch struct {
	OrderDate *time.Time `json:"orderDate" form:"orderDate"`
}
