package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BusinessSearch struct {
	BusinessDate   *time.Time `json:"businessDate" form:"businessDate"`
	BusinessStatus uint       `json:"businessStatus" form:"businessStatus"`
	request.PageInfo
}
