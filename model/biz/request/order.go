package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type OrderSearch struct {
	WechatName    string     `json:"wechatName"`
	ConsumerName  string     `json:"consumerName"`
	ConsumerPhone string     `json:"consumerPhone"`
	OrderDate     *time.Time `json:"orderDate"`
	request.PageInfo
}

type OrderPrintSearch struct {
	OrderDate *time.Time `json:"orderDate"`
}
