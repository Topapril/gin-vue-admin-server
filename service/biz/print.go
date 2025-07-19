package biz

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz"
	bizReq "github.com/flipped-aurora/gin-vue-admin/server/model/biz/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz/response"
)

type PrintService struct{}

var PrintServiceApp = new(PrintService)

func (p *PrintService) CreatePrint(print biz.Print) (err error) {
	err = global.GVA_DB.Create(&print).Error
	return err
}

func (p *PrintService) GetPrintInfoList(info bizReq.PrintSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	// 创建db
	db := global.GVA_DB.Model(&biz.Print{})
	var sysPrint []biz.Print

	// 微信昵称
	if info.WechatName != "" {
		db = db.Where("wechat_name LIKE ?", "%"+info.WechatName+"%")
	}

	// 消费者
	if info.ConsumerName != "" {
		db = db.Where("consumer_name LIKE ?", "%"+info.ConsumerName+"%")
	}

	// 订单日期
	if info.OrderDate != nil && !info.OrderDate.IsZero() {
		orderDateStr := info.OrderDate.Format("2006-01-02")
		db = db.Where("DATE(order_date) = ?", orderDateStr)
	}

	db = db.Order("created_at desc")

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&sysPrint).Error
	return sysPrint, total, err
}

func (p *PrintService) GetPrintInfo(ID string) (print biz.Print, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&print).Error
	return
}

func (p *PrintService) UpdatePrint(ID string, result string) (err error) {
	err = global.GVA_DB.Model(&biz.Print{}).Where("id = ?", ID).Update("print_result", result).Error
	return err
}

// 飞鹅打印
func (p *PrintService) FeiePrinter(info biz.Print) (response.FeieResponse, error) {
	var mealPeriodMap = map[uint]string{1: "午餐", 2: "晚餐"}
	var mealPeriodText string
	if text, ok := mealPeriodMap[info.MealPeriod]; ok {
		mealPeriodText = text
	}

	content := `<CB>` + info.BusinessName + `</CB>`
	content += `<BR>`
	content += `<BOLD><L>商品名称：` + info.GoodsName + `</L></BOLD>`
	content += `<BR>`
	content += `--------------------------------`
	content += `<BR>`
	content += `<BOLD><L>送餐时间：` + info.OrderDate.Format("2006-01-02") + `   ` + mealPeriodText + `</L></BOLD>`
	content += `<BR>`
	content += `<L>数量：` + strconv.Itoa(int(info.GoodsQuantity)) + `份</L>`
	content += `<BR>`
	content += `<L>备注：` + info.Remark + `</L>`
	content += `<BR>`
	content += `<B>` + info.ConsumerName + `</B>`
	content += `<BR>`
	content += `<B>` + info.ConsumerPhone + `</B>`
	content += `<BR>`
	content += `<B>` + info.ConsumerAddress + `</B>`
	content += `<BR>`
	content += `<BR>`
	content += `<QR>` + info.BusinessQrCode + `</QR>`
	content += `<C><L>` + info.BusinessQrCodeTitle + `</L></C>`
	content += `<C><L>有问题联系` + info.BusinessPhone + `</L></C>`

	itime := time.Now().Unix()
	stime := strconv.FormatInt(itime, 10)
	//生成签名
	sig := SHA1(global.GVA_CONFIG.Feie.User + global.GVA_CONFIG.Feie.UKey + stime)

	// 接口超时
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	postValues := url.Values{}
	//账号名
	postValues.Add("user", global.GVA_CONFIG.Feie.User)
	//当前时间的秒数，请求时间
	postValues.Add("stime", stime)
	//签名
	postValues.Add("sig", sig)
	//固定
	postValues.Add("apiname", "Open_printMsg")
	//打印机编号
	postValues.Add("sn", global.GVA_CONFIG.Feie.SN)
	//打印内容
	postValues.Add("content", content)
	//打印次数
	postValues.Add("times", "1")

	res, _ := client.PostForm(global.GVA_CONFIG.Feie.Url, postValues)
	body, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	var resp response.FeieResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return response.FeieResponse{}, err
	}

	return resp, nil
}

func (p *PrintService) VerifyPrintState(orderId string) (response.FeieResponse, error) {
	itime := time.Now().Unix()
	stime := strconv.FormatInt(itime, 10)
	//生成签名
	sig := SHA1(global.GVA_CONFIG.Feie.User + global.GVA_CONFIG.Feie.UKey + stime)

	client := http.Client{}
	postValues := url.Values{}
	//账号名
	postValues.Add("user", global.GVA_CONFIG.Feie.User)
	//当前时间的秒数，请求时间
	postValues.Add("stime", stime)
	//签名
	postValues.Add("sig", sig)
	//固定
	postValues.Add("apiname", "Open_queryOrderState")
	//订单ID由打印订单返回
	postValues.Add("orderid", orderId)

	res, _ := client.PostForm(global.GVA_CONFIG.Feie.Url, postValues)
	body, _ := ioutil.ReadAll(res.Body)

	res.Body.Close()

	var resp response.FeieResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return response.FeieResponse{}, err
	}

	return resp, nil
}

func (p *PrintService) StopFeiePrinter() (response.FeieResponse, error) {
	itime := time.Now().Unix()
	stime := strconv.FormatInt(itime, 10)

	//生成签名
	sig := SHA1(global.GVA_CONFIG.Feie.User + global.GVA_CONFIG.Feie.UKey + stime)

	client := http.Client{}
	postValues := url.Values{}

	//账号名
	postValues.Add("user", global.GVA_CONFIG.Feie.User)

	//当前时间的秒数，请求时间
	postValues.Add("stime", stime)

	//签名
	postValues.Add("sig", sig)

	//固定
	postValues.Add("apiname", "Open_delPrinterSqs")

	//SN
	postValues.Add("sn", global.GVA_CONFIG.Feie.SN)

	res, _ := client.PostForm(global.GVA_CONFIG.Feie.Url, postValues)
	body, _ := ioutil.ReadAll(res.Body)

	res.Body.Close()

	var resp response.FeieResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return response.FeieResponse{}, err
	}

	return resp, nil
}

// 加密
func SHA1(str string) string {
	s := sha1.Sum([]byte(str))
	strsha1 := hex.EncodeToString(s[:])
	return strsha1
}
