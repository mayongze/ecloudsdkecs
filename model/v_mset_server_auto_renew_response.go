// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model


type VMsetServerAutoRenewResponseStateEnum string

// List of State
const (
    VMsetServerAutoRenewResponse_OK VMsetServerAutoRenewResponseStateEnum = "OK"
    VMsetServerAutoRenewResponse_ERROR VMsetServerAutoRenewResponseStateEnum = "ERROR"
    VMsetServerAutoRenewResponse_EXCEPTION VMsetServerAutoRenewResponseStateEnum = "EXCEPTION"
    VMsetServerAutoRenewResponse_ALARM VMsetServerAutoRenewResponseStateEnum = "ALARM"
    VMsetServerAutoRenewResponse_FORBIDDEN VMsetServerAutoRenewResponseStateEnum = "FORBIDDEN"
)

type VMsetServerAutoRenewResponse struct {

	// 每个请求的序列号
	RequestId string `json:"requestId,omitempty"`

	// 页面国际化错误提示
	ErrorMessage string `json:"errorMessage,omitempty"`

	// 统一错误码
	ErrorCode string `json:"errorCode,omitempty"`

	// 返回状态码
	State VMsetServerAutoRenewResponseStateEnum `json:"state,omitempty"`

	// 请求成功时返回的数据
	Body interface{} `json:"body,omitempty"`

	// 统一错误码的自定义参数
	ErrorParams *[]string `json:"errorParams,omitempty"`
}
