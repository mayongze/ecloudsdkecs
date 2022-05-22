// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model


type VmListKeypairResponseStateEnum string

// List of State
const (
    VmListKeypairResponse_OK VmListKeypairResponseStateEnum = "OK"
    VmListKeypairResponse_ERROR VmListKeypairResponseStateEnum = "ERROR"
    VmListKeypairResponse_EXCEPTION VmListKeypairResponseStateEnum = "EXCEPTION"
    VmListKeypairResponse_ALARM VmListKeypairResponseStateEnum = "ALARM"
    VmListKeypairResponse_FORBIDDEN VmListKeypairResponseStateEnum = "FORBIDDEN"
)

type VmListKeypairResponse struct {

	// 每个请求的序列号
	RequestId string `json:"requestId,omitempty"`

	// 页面国际化错误提示
	ErrorMessage string `json:"errorMessage,omitempty"`

	// 统一错误码
	ErrorCode string `json:"errorCode,omitempty"`

	// 返回状态码
	State VmListKeypairResponseStateEnum `json:"state,omitempty"`

	// 请求成功时返回的数据
	Body VmListKeypairResponseBody `json:"body,omitempty"`

	// 统一错误码的自定义参数
	ErrorParams *[]string `json:"errorParams,omitempty"`
}
