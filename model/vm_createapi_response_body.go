// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

type VmCreateapiResponseBodyProcedureCodeEnum string

// List of ProcedureCode
const (
	VmCreateapiResponseBody_PREPAID          VmCreateapiResponseBodyProcedureCodeEnum = "PREPAID"
	VmCreateapiResponseBody_INTERNALAPPROVAL VmCreateapiResponseBodyProcedureCodeEnum = "INTERNALAPPROVAL"
	VmCreateapiResponseBody_POSTPAID         VmCreateapiResponseBodyProcedureCodeEnum = "POSTPAID"
)

/*
type VmCreateapiResponseBody string

// List of OrderExtTypes
const (
	VmCreateapiResponseBody_VM                      VmCreateapiResponseBody = "vm"
	VmCreateapiResponseBody_PERFORMANCEOPTIMIZATION VmCreateapiResponseBody = "performanceOptimization"
	VmCreateapiResponseBody_HIGHPERFORMANCE         VmCreateapiResponseBody = "highPerformance"
	VmCreateapiResponseBody_CAPEBS                  VmCreateapiResponseBody = "capebs"
	VmCreateapiResponseBody_SSDEBS                  VmCreateapiResponseBody = "ssdebs"
	VmCreateapiResponseBody_SSD                     VmCreateapiResponseBody = "ssd"
	VmCreateapiResponseBody_BANDWIDTH               VmCreateapiResponseBody = "bandwidth"
	VmCreateapiResponseBody_IP                      VmCreateapiResponseBody = "ip"
)*/

type VmCreateapiResponseBodyOrderExtTypesEnum string

// List of OrderExtTypes
const (
	VmCreateapiResponseBodyVM                      VmCreateapiResponseBodyOrderExtTypesEnum = "vm"
	VmCreateapiResponseBodyPERFORMANCEOPTIMIZATION VmCreateapiResponseBodyOrderExtTypesEnum = "performanceOptimization"
	VmCreateapiResponseBodyHIGHPERFORMANCE         VmCreateapiResponseBodyOrderExtTypesEnum = "highPerformance"
	VmCreateapiResponseBodyCAPEBS                  VmCreateapiResponseBodyOrderExtTypesEnum = "capebs"
	VmCreateapiResponseBodySSDEBS                  VmCreateapiResponseBodyOrderExtTypesEnum = "ssdebs"
	VmCreateapiResponseBodySSD                     VmCreateapiResponseBodyOrderExtTypesEnum = "ssd"
	VmCreateapiResponseBodyBANDWIDTH               VmCreateapiResponseBodyOrderExtTypesEnum = "bandwidth"
	VmCreateapiResponseBodyIP                      VmCreateapiResponseBodyOrderExtTypesEnum = "ip"
)

type VmCreateapiResponseBody struct {

	// ?????????id
	OrderExts *[]string `json:"orderExts,omitempty"`

	// ??????id
	OrderId string `json:"orderId,omitempty"`

	// ????????????
	ProcedureCode VmCreateapiResponseBodyProcedureCodeEnum `json:"procedureCode,omitempty"`

	// ???????????????
	OrderExtTypes VmCreateapiResponseBodyOrderExtTypesEnum `json:"orderExtTypes,omitempty"`

	// ????????????
	PaymentInfo VmCreateapiResponsePaymentInfo `json:"paymentInfo,omitempty"`
}
