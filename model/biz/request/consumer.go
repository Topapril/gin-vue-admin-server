package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ConsumerSearch struct {
	WechatName    string `json:"wechatName"`
	ConsumerName  string `json:"consumerName"`
	ConsumerPhone string `json:"consumerPhone"`
	request.PageInfo
}

type ConsumerRecordSearch struct {
	ConsumerId string `json:"consumerId"`
	request.PageInfo
}
