package biz

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz"
	bizReq "github.com/flipped-aurora/gin-vue-admin/server/model/biz/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type ConsumerApi struct{}

func (co *ConsumerApi) CreateConsumer(c *gin.Context) {
	var consumer biz.Consumer
	err := c.ShouldBindJSON(&consumer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	consumer.ConsumerId = uuid.New()

	err = consumerService.CreateConsumer(consumer)
	if err != nil {
		global.GVA_LOG.Error("创建失败", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}

	response.OkWithMessage("创建成功", c)
}

func (co *ConsumerApi) DeleteConsumer(c *gin.Context) {
	var consumer biz.Consumer
	err := c.ShouldBindJSON(&consumer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = utils.Verify(consumer.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = consumerService.DeleteConsumer(consumer)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (co *ConsumerApi) UpdateConsumer(c *gin.Context) {
	var consumer biz.Consumer

	err := c.ShouldBindJSON(&consumer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = consumerService.UpdateConsumer(consumer)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (co *ConsumerApi) UpdateConsumerMealCount(c *gin.Context) {
	var consumer biz.Consumer
	err := c.ShouldBindJSON(&consumer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 校验消费者ID
	err = utils.Verify(consumer.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 校验总参数
	err = utils.Verify(consumer, utils.ConsumerMealCountVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 查询消费者信息
	consumerInfo, err := consumerService.GetConsumer(consumer.ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	// 计算餐总次数
	totalMealCount := consumerInfo.TotalMealCount + consumer.TotalMealCount

	// 更新餐总次数
	err = consumerService.UpdateConsumerMealCount(consumer.ID, totalMealCount)
	if err != nil {
		global.GVA_LOG.Error("增加餐品次数失败!", zap.Error(err))
		response.FailWithMessage("增加餐品次数失败:"+err.Error(), c)
		return
	}

	orderDate, err := time.ParseInLocation("2006-01-02", time.Now().Format("2006-01-02"), time.Local)
	if err != nil {
		global.GVA_LOG.Error("时间转换失败!", zap.Error(err))
	}

	consumerRecord := biz.ConsumerRecord{
		WechatName:      consumerInfo.WechatName,
		ConsumerName:    consumerInfo.ConsumerName,
		ConsumerId:      consumerInfo.ConsumerId.String(),
		UsageCount:      consumer.TotalMealCount,
		TransactionType: 2,
		OrderDate:       &orderDate,
		Description:     "加餐",
		Remark:          consumer.Remark,
	}

	// 记录加餐、扣餐
	err = consumerService.CreateConsumerRecord(global.GVA_DB, consumerRecord)
	if err != nil {
		global.GVA_LOG.Error("创建失败", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}

	response.OkWithMessage("已增加", c)
}

func (co *ConsumerApi) UpdateConsumerStatus(c *gin.Context) {
	var consumer biz.Consumer
	err := c.ShouldBindJSON(&consumer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = utils.Verify(consumer, utils.ConsumerStatusVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = consumerService.UpdateConsumerStatus(consumer.ID, consumer.ConsumerStatus)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (co *ConsumerApi) GetConsumerList(c *gin.Context) {
	var pageInfo bizReq.ConsumerSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	consumerList, total, err := consumerService.GetConsumerList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     consumerList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

func (co *ConsumerApi) GetConsumer(c *gin.Context) {
	var consumer biz.Consumer
	err := c.ShouldBindQuery(&consumer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = utils.Verify(consumer.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	data, err := consumerService.GetConsumer(consumer.ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(data, "获取成功", c)
}

func (co *ConsumerApi) GetConsumerRecordList(c *gin.Context) {
	var pageInfo bizReq.ConsumerRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 分页校验
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 消费者ID校验
	err = utils.Verify(pageInfo, utils.ConsumerIdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	consumerRecordList, total, err := consumerService.GetConsumerRecordList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     consumerRecordList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}
