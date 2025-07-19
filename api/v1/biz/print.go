package biz

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz"
	bizReq "github.com/flipped-aurora/gin-vue-admin/server/model/biz/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PrintApi struct{}

func (p *PrintApi) CreatePrint(c *gin.Context) {
	var print biz.Print
	err := c.ShouldBindJSON(&print)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = utils.Verify(print, utils.PrintVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = printService.CreatePrint(print)
	if err != nil {
		global.GVA_LOG.Error("创建失败", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}

	response.OkWithMessage("创建成功", c)
}

func (p *PrintApi) GetPrintList(c *gin.Context) {
	var pageInfo bizReq.PrintSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 分页参数校验
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 查询打印列表
	customerList, total, err := printService.GetPrintInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     customerList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

func (p *PrintApi) PrinterData(c *gin.Context) {
	ID := c.Query("ID")

	// 查询该打印
	printInfo, err := printService.GetPrintInfo(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	// 飞鹅打印调用
	printData, err := printService.FeiePrinter(printInfo)
	if err != nil {
		global.GVA_LOG.Error("飞鹅打印失败!", zap.Error(err))
		response.FailWithMessage("飞鹅打印失败:"+err.Error(), c)
		return
	}

	// 飞鹅打印校验
	printDataStr := fmt.Sprint(printData.Data)
	respData, err := printService.VerifyPrintState(printDataStr)
	if err != nil {
		global.GVA_LOG.Error("飞鹅打印校验错误!", zap.Error(err))
		response.FailWithMessage("飞鹅打印校验错误:"+err.Error(), c)
		return
	}

	if respData.Ret != 0 {
		global.GVA_LOG.Error("飞鹅打印校验失败!", zap.String("msg", respData.Msg))
		response.FailWithMessage("飞鹅打印校验失败:"+respData.Msg, c)
		return
	}

	response.OkWithMessage("打印成功", c)
}

func (p *PrintApi) StopPrint(c *gin.Context) {
	// 停止打印
	data, err := printService.StopFeiePrinter()
	if err != nil {
		global.GVA_LOG.Error("停止飞鹅打印失败", zap.Error(err))
		response.FailWithMessage("停止飞鹅打印失败:"+err.Error(), c)
		return
	}

	if data.Ret != 0 {
		global.GVA_LOG.Error("停止飞鹅打印失败", zap.String("msg", data.Msg))
		response.FailWithMessage("停止飞鹅打印失败:"+data.Msg, c)
		return
	}

	response.OkWithMessage("已清空队列停止打印", c)
}
