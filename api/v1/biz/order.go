package biz

import (
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz"
	bizReq "github.com/flipped-aurora/gin-vue-admin/server/model/biz/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type OrderApi struct{}

func (o *OrderApi) CreateOrder(c *gin.Context) {
	var order biz.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 校验订单参数
	if err := utils.Verify(order, utils.OrderVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 开启事务
	if err := global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 查询消费者数据
		consumerdata, err := consumerService.FindConsumerByConsumerID(tx, order.ConsumerId)
		if err != nil {
			return fmt.Errorf("查询消费者失败: %w", err)
		}

		// 消费者没有开启用餐状态
		if consumerdata.ConsumerStatus != 1 {
			return errors.New("消费者未开启用餐状态")
		}

		// 餐品次数
		if consumerdata.UsageCount >= int(consumerdata.TotalMealCount) {
			return errors.New("没有剩余餐次数")
		}

		// 次数是否满足当前订单
		remainingCount := int(consumerdata.TotalMealCount) - consumerdata.UsageCount
		if order.GoodsQuantity > remainingCount {
			return errors.New("餐次数不足")
		}

		// 创建订单
		if err = orderService.CreateOrder(tx, order); err != nil {
			return fmt.Errorf("创建订单失败: %w", err)
		}

		// 更新消费者使用次数
		if err := consumerService.AddConsumerUsageCount(tx, order.ConsumerId, order.GoodsQuantity); err != nil {
			return fmt.Errorf("更新使用次数失败: %w", err)
		}

		// 消费记录 TransactionType = 1 扣餐 2 = 加餐
		consumerRecord := biz.ConsumerRecord{
			WechatName:      order.WechatName,
			ConsumerName:    order.ConsumerName,
			ConsumerId:      order.ConsumerId,
			UsageCount:      order.GoodsQuantity,
			OrderDate:       order.OrderDate,
			TransactionType: 1,
			Description:     order.OrderDate.Format("2006-01-02") + `用餐`,
			Remark:          order.Remark,
		}

		if err = consumerService.CreateConsumerRecord(tx, consumerRecord); err != nil {
			return fmt.Errorf("创建消费记录失败: %w", err)
		}

		return nil
	}); err != nil {
		global.GVA_LOG.Error("创建订单失败!", zap.Error(err))
		response.FailWithMessage("创建订单失败："+err.Error(), c)
		return
	}

	response.OkWithMessage("创建成功", c)
}

func (o *OrderApi) DeleteOrder(c *gin.Context) {
	var order biz.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 校验订单ID
	err = utils.Verify(order.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 开启事务
	if err := global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 查询订单信息
		orderData, err := orderService.GetOrder(tx, order.ID)
		if err != nil {
			return fmt.Errorf("查询订单失败: %w", err)
		}

		// 更新订单状态 2 = 无效
		if err = orderService.UpdateOrderStatus(tx, order.ID); err != nil {
			return fmt.Errorf("更新订单状态失败: %w", err)
		}

		// 删除订单
		if err = orderService.DeleteOrder(tx, order); err != nil {
			return fmt.Errorf("删除订单失败!: %w", err)
		}

		// 更新消费者使用次数, true = 加餐
		if err := consumerService.SubConsumerUsageCount(tx, orderData.ConsumerId, orderData.GoodsQuantity); err != nil {
			return fmt.Errorf("更新使用次数失败: %w", err)
		}

		// 消费记录 TransactionType = 1 扣餐 2 = 加餐
		consumerRecord := biz.ConsumerRecord{
			WechatName:      orderData.WechatName,
			ConsumerName:    orderData.ConsumerName,
			ConsumerId:      orderData.ConsumerId,
			UsageCount:      orderData.GoodsQuantity,
			OrderDate:       orderData.OrderDate,
			TransactionType: 2,
			Description:     orderData.OrderDate.Format("2006-01-02") + `订单作废`,
			Remark:          orderData.Remark,
		}

		if err = consumerService.CreateConsumerRecord(tx, consumerRecord); err != nil {
			return fmt.Errorf("创建消费记录失败: %w", err)
		}

		return nil
	}); err != nil {
		global.GVA_LOG.Error("删除订单失败!", zap.Error(err))
		response.FailWithMessage("删除订单失败："+err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

func (o *OrderApi) UpdateOrder(c *gin.Context) {
	var order biz.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = orderService.UpdateOrder(order)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

func (o *OrderApi) GetOrderList(c *gin.Context) {
	var pageInfo bizReq.OrderSearch
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

	orderList, total, err := orderService.GetOrderList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     orderList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

func (o *OrderApi) GetOrder(c *gin.Context) {
	var order biz.Order
	err := c.ShouldBindQuery(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = utils.Verify(order.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	data, err := orderService.GetOrder(global.GVA_DB, order.ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithDetailed(data, "获取成功", c)
}

func (o *OrderApi) PrintOrder(c *gin.Context) {
	var order biz.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 查询订单信息
	orderData, err := orderService.GetOrder(global.GVA_DB, order.ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	// 插入打印管理表
	newPrintData := biz.Print{
		GoodsName:           orderData.GoodsName,
		GoodsQuantity:       orderData.GoodsQuantity,
		MealPeriod:          orderData.MealPeriod,
		OrderDate:           orderData.OrderDate,
		ConsumerName:        orderData.ConsumerName,
		ConsumerPhone:       orderData.ConsumerPhone,
		ConsumerAddress:     orderData.ConsumerAddress,
		BusinessName:        global.GVA_CONFIG.Feie.BusinessName,
		BusinessQrCode:      global.GVA_CONFIG.Feie.BusinessQrCode,
		BusinessQrCodeTitle: global.GVA_CONFIG.Feie.BusinessQrCodeTitle,
		BusinessPhone:       global.GVA_CONFIG.Feie.BusinessPhone,
		Remark:              orderData.Remark,
	}

	if err := printService.CreatePrint(newPrintData); err != nil {
		global.GVA_LOG.Error("创建失败", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}

	// 调用飞鹅打印
	printData, err := printService.FeiePrinter(newPrintData)
	if err != nil {
		global.GVA_LOG.Error("飞鹅打印失败", zap.Error(err))
		return
	}

	// 校验飞鹅打印
	printDataStr := fmt.Sprint(printData.Data)
	data, err := printService.VerifyPrintState(printDataStr)
	if err != nil {
		global.GVA_LOG.Error("飞鹅打印校验失败", zap.Error(err))
		return
	}

	if data.Ret != 0 {
		global.GVA_LOG.Error("飞鹅打印校验失败", zap.String("msg", data.Msg))
		response.FailWithMessage("飞鹅打印校验失败:"+data.Msg, c)
		return
	}

	response.OkWithMessage("打印成功", c)
}

func (o *OrderApi) PrintDayOrders(c *gin.Context) {
	var info bizReq.OrderPrintSearch
	err := c.ShouldBindQuery(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 日期校验
	if info.OrderDate == nil || info.OrderDate.IsZero() {
		global.GVA_LOG.Error("订单日期非法!", zap.Error(err))
		response.FailWithMessage("订单日期非法:", c)
		return
	}

	// 查询指定天订单
	printList, err := orderService.GetOrdersByDate(info)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	for _, printData := range printList.([]biz.Order) {
		newPrintData := biz.Print{
			GoodsName:           printData.GoodsName,
			GoodsQuantity:       printData.GoodsQuantity,
			MealPeriod:          printData.MealPeriod,
			OrderDate:           printData.OrderDate,
			ConsumerName:        printData.ConsumerName,
			ConsumerPhone:       printData.ConsumerPhone,
			ConsumerAddress:     printData.ConsumerAddress,
			BusinessName:        global.GVA_CONFIG.Feie.BusinessName,
			BusinessQrCode:      global.GVA_CONFIG.Feie.BusinessQrCode,
			BusinessQrCodeTitle: global.GVA_CONFIG.Feie.BusinessQrCodeTitle,
			BusinessPhone:       global.GVA_CONFIG.Feie.BusinessPhone,
			Remark:              printData.Remark,
		}

		if err := printService.CreatePrint(newPrintData); err != nil {
			global.GVA_LOG.Error("创建失败", zap.Error(err))
			response.FailWithMessage("创建失败", c)
			return
		}

		// 调用飞鹅打印
		printData, err := printService.FeiePrinter(newPrintData)
		if err != nil {
			global.GVA_LOG.Error("飞鹅打印失败", zap.Error(err))
			return
		}

		// 校验飞鹅打印
		printDataStr := fmt.Sprint(printData.Data)
		data, err := printService.VerifyPrintState(printDataStr)
		if err != nil {
			global.GVA_LOG.Error("飞鹅打印校验失败", zap.Error(err))
			return
		}

		if data.Ret != 0 {
			global.GVA_LOG.Error("飞鹅打印校验失败", zap.String("msg", data.Msg))
			response.FailWithMessage("飞鹅打印校验失败:"+data.Msg, c)
			return
		}
	}

	response.OkWithMessage("打印成功", c)
}

func (o *OrderApi) GenerateOrder(c *gin.Context) {
	// 获取营业状态
	mealData, err := mealService.FetchMealForToday()
	if err != nil {
		global.GVA_LOG.Error("查询失败", zap.Error(err))
		response.FailWithMessage("查询失败: "+err.Error(), c)
		return
	}

	// 没有营业日
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.FailWithMessage("找不到营业日数据", c)
		return
	}

	// 是否营业
	if mealData.BusinessStatus == 2 {
		response.FailWithMessage("今日不营业", c)
		return
	}

	// 查询当天是否生成过订单
	count, err := orderService.IsOrderCreatedToday()
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	if count > 0 {
		response.FailWithMessage("已生成过订单", c)
		return
	}

	// 查询状态正常并且总餐数大于已使用餐数的消费者
	consumerData, err := consumerService.GetConsumerAvailableList()
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	for _, consumer := range consumerData {
		// 开启事务
		if err := global.GVA_DB.Transaction(func(tx *gorm.DB) error {
			// 创建订单
			orderData := biz.Order{
				GoodsName:       mealData.LunchName,
				GoodsQuantity:   1,
				MealPeriod:      1,
				WechatName:      consumer.WechatName,
				ConsumerId:      consumer.ConsumerId.String(),
				ConsumerName:    consumer.ConsumerName,
				ConsumerPhone:   consumer.ConsumerPhone,
				ConsumerAddress: consumer.ConsumerAddress,
				DeliveryFee:     consumer.DeliveryFee,
				OrderDate:       mealData.BusinessDate,
				OrderType:       1,
				OrderStatus:     1,
				Remark:          consumer.Remark,
			}

			if err = orderService.CreateOrder(tx, orderData); err != nil {
				return fmt.Errorf("创建订单失败: %w", err)
			}

			// 更新消费者使用次数
			if err := consumerService.AddConsumerUsageCount(tx, consumer.ConsumerId.String(), 1); err != nil {
				return fmt.Errorf("更新使用次数失败: %w", err)
			}

			// 消费记录 TransactionType = 1 扣餐 2 = 加餐
			consumerRecord := biz.ConsumerRecord{
				WechatName:      consumer.WechatName,
				ConsumerName:    consumer.ConsumerName,
				ConsumerId:      consumer.ConsumerId.String(),
				UsageCount:      1,
				OrderDate:       mealData.BusinessDate,
				TransactionType: 1,
				Description:     mealData.BusinessDate.Format("2006-01-02") + `用餐`,
				Remark:          consumer.Remark,
			}

			if err = consumerService.CreateConsumerRecord(tx, consumerRecord); err != nil {
				return fmt.Errorf("创建消费记录失败: %w", err)
			}

			return nil
		}); err != nil {
			global.GVA_LOG.Error("创建订单失败!", zap.Error(err))
			response.FailWithMessage("创建订单失败："+err.Error(), c)
			return
		}
	}

	response.OkWithMessage("订单已生成", c)
}
