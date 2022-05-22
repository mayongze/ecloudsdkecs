// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model


type VmGetVncResponseStateEnum string

// List of State
const (
    VmGetVncResponse_OK VmGetVncResponseStateEnum = "OK"
    VmGetVncResponse_ERROR VmGetVncResponseStateEnum = "ERROR"
    VmGetVncResponse_EXCEPTION VmGetVncResponseStateEnum = "EXCEPTION"
    VmGetVncResponse_ALARM VmGetVncResponseStateEnum = "ALARM"
    VmGetVncResponse_FORBIDDEN VmGetVncResponseStateEnum = "FORBIDDEN"
)
type VmGetVncResponseBodyEnum string

// List of Body
const (
    VmGetVncResponse_ALLOW VmGetVncResponseBodyEnum = "ALLOW"
    VmGetVncResponse_TEMP_DENY VmGetVncResponseBodyEnum = "TEMP_DENY"
    VmGetVncResponse_FOREVER_DENY VmGetVncResponseBodyEnum = "FOREVER_DENY"
)

type VmGetVncResponse struct {

	// 每个请求的序列号
	RequestId string `json:"requestId,omitempty"`

	// 页面国际化错误提示
	ErrorMessage string `json:"errorMessage,omitempty"`

	// 统一错误码
	ErrorCode string `json:"errorCode,omitempty"`

	// 返回状态码
	State VmGetVncResponseStateEnum `json:"state,omitempty"`

	// 请求成功时返回的数据
	Body VmGetVncResponseBodyEnum `json:"body,omitempty"`

	// 统一错误码的自定义参数
	ErrorParams *[]string `json:"errorParams,omitempty"`
}
