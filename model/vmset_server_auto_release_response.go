// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model


type VmsetServerAutoReleaseResponseStateEnum string

// List of State
const (
    VmsetServerAutoReleaseResponse_OK VmsetServerAutoReleaseResponseStateEnum = "OK"
    VmsetServerAutoReleaseResponse_ERROR VmsetServerAutoReleaseResponseStateEnum = "ERROR"
    VmsetServerAutoReleaseResponse_EXCEPTION VmsetServerAutoReleaseResponseStateEnum = "EXCEPTION"
    VmsetServerAutoReleaseResponse_ALARM VmsetServerAutoReleaseResponseStateEnum = "ALARM"
    VmsetServerAutoReleaseResponse_FORBIDDEN VmsetServerAutoReleaseResponseStateEnum = "FORBIDDEN"
)

type VmsetServerAutoReleaseResponse struct {

	// 每个请求的序列号
	RequestId string `json:"requestId,omitempty"`

	// 页面国际化错误提示
	ErrorMessage string `json:"errorMessage,omitempty"`

	// 统一错误码
	ErrorCode string `json:"errorCode,omitempty"`

	// 返回状态码
	State VmsetServerAutoReleaseResponseStateEnum `json:"state,omitempty"`

	// 请求成功时返回的数据
	Body interface{} `json:"body,omitempty"`

	// 统一错误码的自定义参数
	ErrorParams *[]string `json:"errorParams,omitempty"`
}
