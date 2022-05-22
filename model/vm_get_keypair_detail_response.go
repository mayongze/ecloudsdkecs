// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model


type VmGetKeypairDetailResponseStateEnum string

// List of State
const (
    VmGetKeypairDetailResponse_OK VmGetKeypairDetailResponseStateEnum = "OK"
    VmGetKeypairDetailResponse_ERROR VmGetKeypairDetailResponseStateEnum = "ERROR"
    VmGetKeypairDetailResponse_EXCEPTION VmGetKeypairDetailResponseStateEnum = "EXCEPTION"
    VmGetKeypairDetailResponse_ALARM VmGetKeypairDetailResponseStateEnum = "ALARM"
    VmGetKeypairDetailResponse_FORBIDDEN VmGetKeypairDetailResponseStateEnum = "FORBIDDEN"
)

type VmGetKeypairDetailResponse struct {

	// 每个请求的序列号
	RequestId string `json:"requestId,omitempty"`

	// 页面国际化错误提示
	ErrorMessage string `json:"errorMessage,omitempty"`

	// 统一错误码
	ErrorCode string `json:"errorCode,omitempty"`

	// 返回状态码
	State VmGetKeypairDetailResponseStateEnum `json:"state,omitempty"`

	// 请求成功时返回的数据
	Body VmGetKeypairDetailResponseBody `json:"body,omitempty"`

	// 统一错误码的自定义参数
	ErrorParams *[]string `json:"errorParams,omitempty"`
}
