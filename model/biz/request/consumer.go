package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ConsumerSearch struct {
	WechatName    string `json:"wechatName" form:"wechatName"`
	ConsumerName  string `json:"consumerName" form:"consumerName"`
	ConsumerPhone string `json:"consumerPhone" form:"consumerPhone"`
	request.PageInfo
}

type ConsumerRecordSearch struct {
	ConsumerId string `json:"consumerId" form:"consumerId"`
	request.PageInfo
}
