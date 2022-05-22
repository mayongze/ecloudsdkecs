// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

type VmChangeServerResponseBodyProcedureCodeEnum string

// List of ProcedureCode
const (
	VmChangeServerResponseBody_PREPAID          VmChangeServerResponseBodyProcedureCodeEnum = "PREPAID"
	VmChangeServerResponseBody_INTERNALAPPROVAL VmChangeServerResponseBodyProcedureCodeEnum = "INTERNALAPPROVAL"
	VmChangeServerResponseBody_POSTPAID         VmChangeServerResponseBodyProcedureCodeEnum = "POSTPAID"
)

/*type VmChangeServerResponseBody string

// List of OrderExtTypes
const (
    VmChangeServerResponseBody_VM VmChangeServerResponseBody = "vm"
    VmChangeServerResponseBody_PERFORMANCEOPTIMIZATION VmChangeServerResponseBody = "performanceOptimization"
    VmChangeServerResponseBody_HIGHPERFORMANCE VmChangeServerResponseBody = "highPerformance"
    VmChangeServerResponseBody_CAPEBS VmChangeServerResponseBody = "capebs"
    VmChangeServerResponseBody_SSDEBS VmChangeServerResponseBody = "ssdebs"
    VmChangeServerResponseBody_SSD VmChangeServerResponseBody = "ssd"
    VmChangeServerResponseBody_BANDWIDTH VmChangeServerResponseBody = "bandwidth"
    VmChangeServerResponseBody_IP VmChangeServerResponseBody = "ip"
)*/
type VmChangeServerResponseBodyOrderExtTypesEnum string

// List of OrderExtTypes
const (
	VmChangeServerResponseBodyVM                      VmChangeServerResponseBodyOrderExtTypesEnum = "vm"
	VmChangeServerResponseBodyPERFORMANCEOPTIMIZATION VmChangeServerResponseBodyOrderExtTypesEnum = "performanceOptimization"
	VmChangeServerResponseBodyHIGHPERFORMANCE         VmChangeServerResponseBodyOrderExtTypesEnum = "highPerformance"
	VmChangeServerResponseBodyCAPEBS                  VmChangeServerResponseBodyOrderExtTypesEnum = "capebs"
	VmChangeServerResponseBodySSDEBS                  VmChangeServerResponseBodyOrderExtTypesEnum = "ssdebs"
	VmChangeServerResponseBodySSD                     VmChangeServerResponseBodyOrderExtTypesEnum = "ssd"
	VmChangeServerResponseBodyBANDWIDTH               VmChangeServerResponseBodyOrderExtTypesEnum = "bandwidth"
	VmChangeServerResponseBodyIP                      VmChangeServerResponseBodyOrderExtTypesEnum = "ip"
)

type VmChangeServerResponseBody struct {

	// 订单项id
	OrderExts *[]string `json:"orderExts,omitempty"`

	// 订单id
	OrderId string `json:"orderId,omitempty"`

	// 付费类型
	ProcedureCode VmChangeServerResponseBodyProcedureCodeEnum `json:"procedureCode,omitempty"`

	// 订单项类型
	OrderExtTypes VmChangeServerResponseBodyOrderExtTypesEnum `json:"orderExtTypes,omitempty"`

	// 支付信息
	PaymentInfo VmChangeServerResponsePaymentInfo `json:"paymentInfo,omitempty"`
}
