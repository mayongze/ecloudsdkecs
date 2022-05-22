// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model


type VmGetRebuildImageResponseStateEnum string

// List of State
const (
    VmGetRebuildImageResponse_OK VmGetRebuildImageResponseStateEnum = "OK"
    VmGetRebuildImageResponse_ERROR VmGetRebuildImageResponseStateEnum = "ERROR"
    VmGetRebuildImageResponse_EXCEPTION VmGetRebuildImageResponseStateEnum = "EXCEPTION"
    VmGetRebuildImageResponse_ALARM VmGetRebuildImageResponseStateEnum = "ALARM"
    VmGetRebuildImageResponse_FORBIDDEN VmGetRebuildImageResponseStateEnum = "FORBIDDEN"
)

type VmGetRebuildImageResponse struct {

	// 每个请求的序列号
	RequestId string `json:"requestId,omitempty"`

	// 页面国际化错误提示
	ErrorMessage string `json:"errorMessage,omitempty"`

	// 统一错误码
	ErrorCode string `json:"errorCode,omitempty"`

	// 返回状态码
	State VmGetRebuildImageResponseStateEnum `json:"state,omitempty"`

	// 请求成功时返回的数据
	Body *[]VmGetRebuildImageResponseBody `json:"body,omitempty"`

	// 统一错误码的自定义参数
	ErrorParams *[]string `json:"errorParams,omitempty"`
}
