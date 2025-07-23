package biz

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz"
	bizReq "github.com/flipped-aurora/gin-vue-admin/server/model/biz/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BusinessApi struct{}

func (b *BusinessApi) CreateBusiness(c *gin.Context) {
	var business biz.Business
	err := c.ShouldBindJSON(&business)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 校验参数
	err = utils.Verify(business, utils.BusinessVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 查询营业日期重复
	count, err := businessService.HasDuplicateBusinessDate(business.BusinessDate)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.GVA_LOG.Error("查询失败", zap.Error(err))
		response.FailWithMessage("查询失败: "+err.Error(), c)
		return
	}

	// 营业日期已存在
	if count > 0 {
		global.GVA_LOG.Error("营业日期已存在")
		response.FailWithMessage("营业日期已存在", c)
		return
	}

	// 创建营业
	err = businessService.CreateBusiness(business)
	if err != nil {
		global.GVA_LOG.Error("创建失败", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}

	response.OkWithMessage("创建成功", c)
}

func (b *BusinessApi) DeleteBusiness(c *gin.Context) {
	var business biz.Business
	err := c.ShouldBindJSON(&business)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = utils.Verify(business.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = businessService.DeleteBusiness(business)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

func (b *BusinessApi) UpdateBusiness(c *gin.Context) {
	var business biz.Business
	err := c.ShouldBindJSON(&business)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = businessService.UpdateBusiness(business)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

func (b *BusinessApi) UpdateBusinessStatus(c *gin.Context) {
	var business biz.Business
	err := c.ShouldBindJSON(&business)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = businessService.UpdateBusinessStatus(business)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

func (b *BusinessApi) GetBusinessList(c *gin.Context) {
	var pageInfo bizReq.BusinessSearch
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

	businessList, total, err := businessService.GetBusinessList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     businessList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

func (b *BusinessApi) GetBusiness(c *gin.Context) {
	var business biz.Business
	err := c.ShouldBindQuery(&business)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = utils.Verify(business.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	data, err := businessService.GetBusiness(business.ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithDetailed(data, "获取成功", c)
}
