package biz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz"
	bizReq "github.com/flipped-aurora/gin-vue-admin/server/model/biz/request"
	"gorm.io/gorm"
)

type ConsumerService struct{}

var ConsumerServiceApp = new(ConsumerService)

func (c *ConsumerService) CreateConsumer(consumer biz.Consumer) (err error) {
	err = global.GVA_DB.Create(&consumer).Error
	return err
}

func (c *ConsumerService) DeleteConsumer(consumer biz.Consumer) (err error) {
	err = global.GVA_DB.Delete(&consumer).Error
	return err
}

func (c *ConsumerService) UpdateConsumer(consumer biz.Consumer) (err error) {
	err = global.GVA_DB.Model(&biz.Consumer{}).Where("id = ?", consumer.ID).Updates(&consumer).Error
	return err
}

func (c *ConsumerService) UpdateConsumerMealCount(id uint, totalMealCount int) (err error) {
	err = global.GVA_DB.Model(&biz.Consumer{}).Where("id = ?", id).Update("total_meal_count", totalMealCount).Error
	return err
}

func (c *ConsumerService) UpdateConsumerStatus(id uint, consumerStatus uint) (err error) {
	err = global.GVA_DB.Model(&biz.Consumer{}).Where("id = ?", id).Update("consumer_status", consumerStatus).Error
	return err
}

func (c *ConsumerService) AddConsumerUsageCount(db *gorm.DB, consumerId string, goodsQuantity int) (err error) {
	// 没有数量
	if goodsQuantity == 0 {
		return nil
	}

	err = db.Model(&biz.Consumer{}).Where("consumer_id = ?", consumerId).Update("usage_count", gorm.Expr("usage_count + ?", goodsQuantity)).Error
	return err
}

func (c *ConsumerService) SubConsumerUsageCount(db *gorm.DB, consumerId string, goodsQuantity int) (err error) {
	// 没有数量
	if goodsQuantity == 0 {
		return nil
	}

	err = db.Model(&biz.Consumer{}).Where("consumer_id = ? AND usage_count >= ?", consumerId, goodsQuantity).Update("usage_count", gorm.Expr("usage_count - ?", goodsQuantity)).Error
	return err
}

func (c *ConsumerService) GetConsumerList(info bizReq.ConsumerSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&biz.Consumer{})
	var consumer []biz.Consumer

	if info.WechatName != "" {
		db = db.Where("wechat_name LIKE ?", "%"+info.WechatName+"%")
	}

	if info.ConsumerName != "" {
		db = db.Where("consumer_name LIKE ?", "%"+info.ConsumerName+"%")
	}

	if info.ConsumerPhone != "" {
		db = db.Where("consumer_phone = ?", info.ConsumerPhone)
	}

	db = db.Order("created_at desc")

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&consumer).Error
	return consumer, total, err
}

func (c *ConsumerService) GetConsumer(id uint) (consumer biz.Consumer, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&consumer).Error
	return
}

func (c *ConsumerService) FindConsumerByConsumerID(db *gorm.DB, consumerId string) (consumer biz.Consumer, err error) {
	err = db.Where("consumer_id = ?", consumerId).First(&consumer).Error
	return
}

func (c *ConsumerService) GetConsumerRecordList(info bizReq.ConsumerRecordSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&biz.ConsumerRecord{})
	var consumerRecord []biz.ConsumerRecord

	if info.ConsumerId != "" {
		db = db.Where("consumer_id = ?", info.ConsumerId)
	}

	db = db.Order("created_at desc")

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&consumerRecord).Error
	return consumerRecord, total, err
}

func (c *ConsumerService) CreateConsumerRecord(db *gorm.DB, consumerRecord biz.ConsumerRecord) (err error) {
	err = db.Create(&consumerRecord).Error
	return err
}

func (c *ConsumerService) GetConsumerAvailableList() (list []biz.Consumer, err error) {
	db := global.GVA_DB.Model(&biz.Consumer{})

	db = db.Where("consumer_status = ? AND usage_count < total_meal_count", 1)

	db = db.Order("created_at desc")

	err = db.Find(&list).Error
	return list, err
}
