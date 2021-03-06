// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

type VmChangeResponseBodyProcedureCodeEnum string

// List of ProcedureCode
const (
	VmChangeResponseBody_PREPAID          VmChangeResponseBodyProcedureCodeEnum = "PREPAID"
	VmChangeResponseBody_INTERNALAPPROVAL VmChangeResponseBodyProcedureCodeEnum = "INTERNALAPPROVAL"
	VmChangeResponseBody_POSTPAID         VmChangeResponseBodyProcedureCodeEnum = "POSTPAID"
)

/*type VmChangeResponseBody string

// List of OrderExtTypes
const (
	VmChangeResponseBody_VM                      VmChangeResponseBody = "vm"
	VmChangeResponseBody_PERFORMANCEOPTIMIZATION VmChangeResponseBody = "performanceOptimization"
	VmChangeResponseBody_HIGHPERFORMANCE         VmChangeResponseBody = "highPerformance"
	VmChangeResponseBody_CAPEBS                  VmChangeResponseBody = "capebs"
	VmChangeResponseBody_SSDEBS                  VmChangeResponseBody = "ssdebs"
	VmChangeResponseBody_SSD                     VmChangeResponseBody = "ssd"
	VmChangeResponseBody_BANDWIDTH               VmChangeResponseBody = "bandwidth"
	VmChangeResponseBody_IP                      VmChangeResponseBody = "ip"
)*/

type VmChangeResponseBodyOrderExtTypesEnum string

// List of OrderExtTypes
const (
	VmChangeResponseBodyVM                      VmChangeResponseBodyOrderExtTypesEnum = "vm"
	VmChangeResponseBodyPERFORMANCEOPTIMIZATION VmChangeResponseBodyOrderExtTypesEnum = "performanceOptimization"
	VmChangeResponseBodyHIGHPERFORMANCE         VmChangeResponseBodyOrderExtTypesEnum = "highPerformance"
	VmChangeResponseBodyCAPEBS                  VmChangeResponseBodyOrderExtTypesEnum = "capebs"
	VmChangeResponseBodySSDEBS                  VmChangeResponseBodyOrderExtTypesEnum = "ssdebs"
	VmChangeResponseBodySSD                     VmChangeResponseBodyOrderExtTypesEnum = "ssd"
	VmChangeResponseBodyBANDWIDTH               VmChangeResponseBodyOrderExtTypesEnum = "bandwidth"
	VmChangeResponseBodyIP                      VmChangeResponseBodyOrderExtTypesEnum = "ip"
)

type VmChangeResponseBody struct {

	// ?????????id
	OrderExts *[]string `json:"orderExts,omitempty"`

	// ??????id
	OrderId string `json:"orderId,omitempty"`

	// ????????????
	ProcedureCode VmChangeResponseBodyProcedureCodeEnum `json:"procedureCode,omitempty"`

	// ???????????????
	OrderExtTypes VmChangeResponseBodyOrderExtTypesEnum `json:"orderExtTypes,omitempty"`

	// ????????????
	PaymentInfo VmChangeResponsePaymentInfo `json:"paymentInfo,omitempty"`
}
