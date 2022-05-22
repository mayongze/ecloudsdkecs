// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

type VmDisResponseBodyProcedureCodeEnum string

// List of ProcedureCode
const (
	VmDisResponseBody_PREPAID          VmDisResponseBodyProcedureCodeEnum = "PREPAID"
	VmDisResponseBody_INTERNALAPPROVAL VmDisResponseBodyProcedureCodeEnum = "INTERNALAPPROVAL"
	VmDisResponseBody_POSTPAID         VmDisResponseBodyProcedureCodeEnum = "POSTPAID"
)

/*type VmDisResponseBody string

// List of OrderExtTypes
const (
	VmDisResponseBody_VM                      VmDisResponseBody = "vm"
	VmDisResponseBody_PERFORMANCEOPTIMIZATION VmDisResponseBody = "performanceOptimization"
	VmDisResponseBody_HIGHPERFORMANCE         VmDisResponseBody = "highPerformance"
	VmDisResponseBody_CAPEBS                  VmDisResponseBody = "capebs"
	VmDisResponseBody_SSDEBS                  VmDisResponseBody = "ssdebs"
	VmDisResponseBody_SSD                     VmDisResponseBody = "ssd"
	VmDisResponseBody_BANDWIDTH               VmDisResponseBody = "bandwidth"
	VmDisResponseBody_IP                      VmDisResponseBody = "ip"
)*/

type VmDisResponseBodyOrderExtTypesEnum string

// List of OrderExtTypes
const (
	VmDisResponseBodyVM                      VmDisResponseBodyOrderExtTypesEnum = "vm"
	VmDisResponseBodyPERFORMANCEOPTIMIZATION VmDisResponseBodyOrderExtTypesEnum = "performanceOptimization"
	VmDisResponseBodyHIGHPERFORMANCE         VmDisResponseBodyOrderExtTypesEnum = "highPerformance"
	VmDisResponseBodyCAPEBS                  VmDisResponseBodyOrderExtTypesEnum = "capebs"
	VmDisResponseBodySSDEBS                  VmDisResponseBodyOrderExtTypesEnum = "ssdebs"
	VmDisResponseBodySSD                     VmDisResponseBodyOrderExtTypesEnum = "ssd"
	VmDisResponseBodyBANDWIDTH               VmDisResponseBodyOrderExtTypesEnum = "bandwidth"
	VmDisResponseBodyIP                      VmDisResponseBodyOrderExtTypesEnum = "ip"
)

type VmDisResponseBody struct {

	// 订单项id
	OrderExts *[]string `json:"orderExts,omitempty"`

	// 订单id
	OrderId string `json:"orderId,omitempty"`

	// 付费类型
	ProcedureCode VmDisResponseBodyProcedureCodeEnum `json:"procedureCode,omitempty"`

	// 订单项类型
	OrderExtTypes VmDisResponseBodyOrderExtTypesEnum `json:"orderExtTypes,omitempty"`

	// 支付信息
	PaymentInfo VmDisResponsePaymentInfo `json:"paymentInfo,omitempty"`
}
