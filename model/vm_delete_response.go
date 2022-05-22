// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model


type VmDeleteResponseStateEnum string

// List of State
const (
    VmDeleteResponse_OK VmDeleteResponseStateEnum = "OK"
    VmDeleteResponse_ERROR VmDeleteResponseStateEnum = "ERROR"
    VmDeleteResponse_EXCEPTION VmDeleteResponseStateEnum = "EXCEPTION"
    VmDeleteResponse_ALARM VmDeleteResponseStateEnum = "ALARM"
    VmDeleteResponse_FORBIDDEN VmDeleteResponseStateEnum = "FORBIDDEN"
)
type VmDeleteResponseBodyEnum string

// List of Body
const (
    VmDeleteResponse_ALLOW VmDeleteResponseBodyEnum = "ALLOW"
    VmDeleteResponse_TEMP_DENY VmDeleteResponseBodyEnum = "TEMP_DENY"
    VmDeleteResponse_FOREVER_DENY VmDeleteResponseBodyEnum = "FOREVER_DENY"
)

type VmDeleteResponse struct {

	// 每个请求的序列号
	RequestId string `json:"requestId,omitempty"`

	// 页面国际化错误提示
	ErrorMessage string `json:"errorMessage,omitempty"`

	// 统一错误码
	ErrorCode string `json:"errorCode,omitempty"`

	// 返回状态码
	State VmDeleteResponseStateEnum `json:"state,omitempty"`

	// 请求成功时返回的数据
	Body VmDeleteResponseBodyEnum `json:"body,omitempty"`

	// 统一错误码的自定义参数
	ErrorParams *[]string `json:"errorParams,omitempty"`
}
