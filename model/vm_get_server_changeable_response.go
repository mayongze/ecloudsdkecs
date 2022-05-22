// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model


type VmGetServerChangeableResponseStateEnum string

// List of State
const (
    VmGetServerChangeableResponse_OK VmGetServerChangeableResponseStateEnum = "OK"
    VmGetServerChangeableResponse_ERROR VmGetServerChangeableResponseStateEnum = "ERROR"
    VmGetServerChangeableResponse_EXCEPTION VmGetServerChangeableResponseStateEnum = "EXCEPTION"
    VmGetServerChangeableResponse_ALARM VmGetServerChangeableResponseStateEnum = "ALARM"
    VmGetServerChangeableResponse_FORBIDDEN VmGetServerChangeableResponseStateEnum = "FORBIDDEN"
)

type VmGetServerChangeableResponse struct {

	// 每个请求的序列号
	RequestId string `json:"requestId,omitempty"`

	// 页面国际化错误提示
	ErrorMessage string `json:"errorMessage,omitempty"`

	// 统一错误码
	ErrorCode string `json:"errorCode,omitempty"`

	// 返回状态码
	State VmGetServerChangeableResponseStateEnum `json:"state,omitempty"`

	// 请求成功时返回的数据
	Body *[]VmGetServerChangeableResponseBody `json:"body,omitempty"`

	// 统一错误码的自定义参数
	ErrorParams *[]string `json:"errorParams,omitempty"`
}
