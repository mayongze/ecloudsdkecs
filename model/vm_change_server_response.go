// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model


type VmChangeServerResponseStateEnum string

// List of State
const (
    VmChangeServerResponse_OK VmChangeServerResponseStateEnum = "OK"
    VmChangeServerResponse_ERROR VmChangeServerResponseStateEnum = "ERROR"
    VmChangeServerResponse_EXCEPTION VmChangeServerResponseStateEnum = "EXCEPTION"
    VmChangeServerResponse_ALARM VmChangeServerResponseStateEnum = "ALARM"
    VmChangeServerResponse_FORBIDDEN VmChangeServerResponseStateEnum = "FORBIDDEN"
)

type VmChangeServerResponse struct {

	// 每个请求的序列号
	RequestId string `json:"requestId,omitempty"`

	// 页面国际化错误提示
	ErrorMessage string `json:"errorMessage,omitempty"`

	// 统一错误码
	ErrorCode string `json:"errorCode,omitempty"`

	// 返回状态码
	State VmChangeServerResponseStateEnum `json:"state,omitempty"`

	// 请求成功时返回的数据
	Body VmChangeServerResponseBody `json:"body,omitempty"`

	// 统一错误码的自定义参数
	ErrorParams *[]string `json:"errorParams,omitempty"`
}
