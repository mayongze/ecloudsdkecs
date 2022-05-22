// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model



type VmDisResponsePaymentInfo struct {

	// 产品名称
	ProductName string `json:"ProductName,omitempty"`

	// 交易代码
	ActivityCode string `json:"ActivityCode,omitempty"`

	// 商户自定义参数
	CustomParam string `json:"CustomParam,omitempty"`

	// 订单号
	OrderNo string `json:"OrderNo,omitempty"`

	// 产品编号
	ProductID string `json:"ProductID,omitempty"`

	// 支付成功跳转返回页面URL
	ReturnURL string `json:"ReturnURL,omitempty"`

	// 请求方渠道编码
	ReqChannel string `json:"ReqChannel,omitempty"`

	// 中国移动用户号码
	IDValue string `json:"IDValue,omitempty"`

	// 中国移动用户标识类型
	IDType string `json:"IDType,omitempty"`

	// 微信公众号ID/应用ID
	WeiXinAppId string `json:"WeiXinAppId,omitempty"`

	// 商品展示网址
	ProductURL string `json:"ProductURL,omitempty"`

	// 支付链接
	PayLink string `json:"payLink,omitempty"`

	// 业务类型
	BusiType string `json:"BusiType,omitempty"`

	// 用户支付金额
	Payment string `json:"Payment,omitempty"`

	// 证书标识串
	CerID string `json:"CerID,omitempty"`

	// 客户编码
	CustomerNumber string `json:"CustomerNumber,omitempty"`

	// 归属省份
	HomeProv string `json:"HomeProv,omitempty"`

	// 订单总金额
	OrderMoney string `json:"OrderMoney,omitempty"`

	// 订单结果通知URL
	NotifyURL string `json:"NotifyURL,omitempty"`

	// 产品描述
	ProductDesc string `json:"ProductDesc,omitempty"`

	// 移动云全网渠道编码
	BusinessChannels string `json:"BusinessChannels,omitempty"`

	// 充值金额
	ChargeMoney string `json:"ChargeMoney,omitempty"`

	// 签名值
	SignValue string `json:"SignValue,omitempty"`
}
