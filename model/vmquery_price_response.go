// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model


type VmqueryPriceResponseStateEnum string

// List of State
const (
    VmqueryPriceResponse_OK VmqueryPriceResponseStateEnum = "OK"
    VmqueryPriceResponse_ERROR VmqueryPriceResponseStateEnum = "ERROR"
    VmqueryPriceResponse_EXCEPTION VmqueryPriceResponseStateEnum = "EXCEPTION"
    VmqueryPriceResponse_ALARM VmqueryPriceResponseStateEnum = "ALARM"
    VmqueryPriceResponse_FORBIDDEN VmqueryPriceResponseStateEnum = "FORBIDDEN"
)

type VmqueryPriceResponse struct {

	// 每个请求的序列号
	RequestId string `json:"requestId,omitempty"`

	// 页面国际化错误提示
	ErrorMessage string `json:"errorMessage,omitempty"`

	// 统一错误码
	ErrorCode string `json:"errorCode,omitempty"`

	// 返回状态码
	State VmqueryPriceResponseStateEnum `json:"state,omitempty"`

	// 请求成功时返回的数据
	Body VmqueryPriceResponseBody `json:"body,omitempty"`

	// 统一错误码的自定义参数
	ErrorParams *[]string `json:"errorParams,omitempty"`
}
