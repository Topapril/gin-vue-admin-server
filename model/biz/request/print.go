package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type PrintSearch struct {
	WechatName   string     `json:"wechatName"`
	ConsumerName string     `json:"consumerName"`
	OrderDate    *time.Time `json:"orderDate"`
	request.PageInfo
}
