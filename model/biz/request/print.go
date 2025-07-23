package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type PrintSearch struct {
	WechatName   string     `json:"wechatName" form:"wechatName"`
	ConsumerName string     `json:"consumerName" form:"consumerName"`
	OrderDate    *time.Time `json:"orderDate" form:"orderDate"`
	request.PageInfo
}
