// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model


type VmRebootResponseStateEnum string

// List of State
const (
    VmRebootResponse_OK VmRebootResponseStateEnum = "OK"
    VmRebootResponse_ERROR VmRebootResponseStateEnum = "ERROR"
    VmRebootResponse_EXCEPTION VmRebootResponseStateEnum = "EXCEPTION"
    VmRebootResponse_ALARM VmRebootResponseStateEnum = "ALARM"
    VmRebootResponse_FORBIDDEN VmRebootResponseStateEnum = "FORBIDDEN"
)

type VmRebootResponse struct {

	// 每个请求的序列号
	RequestId string `json:"requestId,omitempty"`

	// 页面国际化错误提示
	ErrorMessage string `json:"errorMessage,omitempty"`

	// 统一错误码
	ErrorCode string `json:"errorCode,omitempty"`

	// 返回状态码
	State VmRebootResponseStateEnum `json:"state,omitempty"`

	// 请求成功时返回的数据
	Body interface{} `json:"body,omitempty"`

	// 统一错误码的自定义参数
	ErrorParams *[]string `json:"errorParams,omitempty"`
}
