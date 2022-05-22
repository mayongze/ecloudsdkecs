// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model


type VmDisResponseStateEnum string

// List of State
const (
    VmDisResponse_OK VmDisResponseStateEnum = "OK"
    VmDisResponse_ERROR VmDisResponseStateEnum = "ERROR"
    VmDisResponse_EXCEPTION VmDisResponseStateEnum = "EXCEPTION"
    VmDisResponse_ALARM VmDisResponseStateEnum = "ALARM"
    VmDisResponse_FORBIDDEN VmDisResponseStateEnum = "FORBIDDEN"
)

type VmDisResponse struct {

	// 每个请求的序列号
	RequestId string `json:"requestId,omitempty"`

	// 页面国际化错误提示
	ErrorMessage string `json:"errorMessage,omitempty"`

	// 统一错误码
	ErrorCode string `json:"errorCode,omitempty"`

	// 返回状态码
	State VmDisResponseStateEnum `json:"state,omitempty"`

	// 请求成功时返回的数据
	Body VmDisResponseBody `json:"body,omitempty"`

	// 统一错误码的自定义参数
	ErrorParams *[]string `json:"errorParams,omitempty"`
}
