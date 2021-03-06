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

	// ?????????id
	OrderExts *[]string `json:"orderExts,omitempty"`

	// ??????id
	OrderId string `json:"orderId,omitempty"`

	// ????????????
	ProcedureCode VmChangeServerResponseBodyProcedureCodeEnum `json:"procedureCode,omitempty"`

	// ???????????????
	OrderExtTypes VmChangeServerResponseBodyOrderExtTypesEnum `json:"orderExtTypes,omitempty"`

	// ????????????
	PaymentInfo VmChangeServerResponsePaymentInfo `json:"paymentInfo,omitempty"`
}
