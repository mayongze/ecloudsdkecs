// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model


type VmRenewResponseStateEnum string

// List of State
const (
    VmRenewResponse_OK VmRenewResponseStateEnum = "OK"
    VmRenewResponse_ERROR VmRenewResponseStateEnum = "ERROR"
    VmRenewResponse_EXCEPTION VmRenewResponseStateEnum = "EXCEPTION"
    VmRenewResponse_ALARM VmRenewResponseStateEnum = "ALARM"
    VmRenewResponse_FORBIDDEN VmRenewResponseStateEnum = "FORBIDDEN"
)

type VmRenewResponse struct {

	// 每个请求的序列号
	RequestId string `json:"requestId,omitempty"`

	// 页面国际化错误提示
	ErrorMessage string `json:"errorMessage,omitempty"`

	// 统一错误码
	ErrorCode string `json:"errorCode,omitempty"`

	// 返回状态码
	State VmRenewResponseStateEnum `json:"state,omitempty"`

	// 请求成功时返回的数据
	Body VmRenewResponseBody `json:"body,omitempty"`

	// 统一错误码的自定义参数
	ErrorParams *[]string `json:"errorParams,omitempty"`
}
