// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

type VmRenewResponseBodyProcedureCodeEnum string

// List of ProcedureCode
const (
	VmRenewResponseBody_PREPAID          VmRenewResponseBodyProcedureCodeEnum = "PREPAID"
	VmRenewResponseBody_INTERNALAPPROVAL VmRenewResponseBodyProcedureCodeEnum = "INTERNALAPPROVAL"
	VmRenewResponseBody_POSTPAID         VmRenewResponseBodyProcedureCodeEnum = "POSTPAID"
)

/*type VmRenewResponseBody string

// List of OrderExtTypes
const (
    VmRenewResponseBody_VM VmRenewResponseBody = "vm"
    VmRenewResponseBody_PERFORMANCEOPTIMIZATION VmRenewResponseBody = "performanceOptimization"
    VmRenewResponseBody_HIGHPERFORMANCE VmRenewResponseBody = "highPerformance"
    VmRenewResponseBody_CAPEBS VmRenewResponseBody = "capebs"
    VmRenewResponseBody_SSDEBS VmRenewResponseBody = "ssdebs"
    VmRenewResponseBody_SSD VmRenewResponseBody = "ssd"
    VmRenewResponseBody_BANDWIDTH VmRenewResponseBody = "bandwidth"
    VmRenewResponseBody_IP VmRenewResponseBody = "ip"
)*/
type VmRenewResponseBodyOrderExtTypesEnum string

// List of OrderExtTypes
const (
	VmRenewResponseBodyVM                      VmRenewResponseBodyOrderExtTypesEnum = "vm"
	VmRenewResponseBodyPERFORMANCEOPTIMIZATION VmRenewResponseBodyOrderExtTypesEnum = "performanceOptimization"
	VmRenewResponseBodyHIGHPERFORMANCE         VmRenewResponseBodyOrderExtTypesEnum = "highPerformance"
	VmRenewResponseBodyCAPEBS                  VmRenewResponseBodyOrderExtTypesEnum = "capebs"
	VmRenewResponseBodySSDEBS                  VmRenewResponseBodyOrderExtTypesEnum = "ssdebs"
	VmRenewResponseBodySSD                     VmRenewResponseBodyOrderExtTypesEnum = "ssd"
	VmRenewResponseBodyBANDWIDTH               VmRenewResponseBodyOrderExtTypesEnum = "bandwidth"
	VmRenewResponseBodyIP                      VmRenewResponseBodyOrderExtTypesEnum = "ip"
)

type VmRenewResponseBody struct {

	// 订单项id
	OrderExts *[]string `json:"orderExts,omitempty"`

	// 订单id
	OrderId string `json:"orderId,omitempty"`

	// 付费类型
	ProcedureCode VmRenewResponseBodyProcedureCodeEnum `json:"procedureCode,omitempty"`

	// 订单项类型
	OrderExtTypes VmRenewResponseBodyOrderExtTypesEnum `json:"orderExtTypes,omitempty"`

	// 支付信息
	PaymentInfo VmRenewResponsePaymentInfo `json:"paymentInfo,omitempty"`
}
