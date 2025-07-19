package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type MealSearch struct {
	BusinessDate *time.Time `json:"businessDate" form:"businessDate"`
	request.PageInfo
}
