// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model


type VmChangeResponseStateEnum string

// List of State
const (
    VmChangeResponse_OK VmChangeResponseStateEnum = "OK"
    VmChangeResponse_ERROR VmChangeResponseStateEnum = "ERROR"
    VmChangeResponse_EXCEPTION VmChangeResponseStateEnum = "EXCEPTION"
    VmChangeResponse_ALARM VmChangeResponseStateEnum = "ALARM"
    VmChangeResponse_FORBIDDEN VmChangeResponseStateEnum = "FORBIDDEN"
)

type VmChangeResponse struct {

	// 每个请求的序列号
	RequestId string `json:"requestId,omitempty"`

	// 页面国际化错误提示
	ErrorMessage string `json:"errorMessage,omitempty"`

	// 统一错误码
	ErrorCode string `json:"errorCode,omitempty"`

	// 返回状态码
	State VmChangeResponseStateEnum `json:"state,omitempty"`

	// 请求成功时返回的数据
	Body VmChangeResponseBody `json:"body,omitempty"`

	// 统一错误码的自定义参数
	ErrorParams *[]string `json:"errorParams,omitempty"`
}
