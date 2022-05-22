// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model


type VmGetServerPortsResponseStateEnum string

// List of State
const (
    VmGetServerPortsResponse_OK VmGetServerPortsResponseStateEnum = "OK"
    VmGetServerPortsResponse_ERROR VmGetServerPortsResponseStateEnum = "ERROR"
    VmGetServerPortsResponse_EXCEPTION VmGetServerPortsResponseStateEnum = "EXCEPTION"
    VmGetServerPortsResponse_ALARM VmGetServerPortsResponseStateEnum = "ALARM"
    VmGetServerPortsResponse_FORBIDDEN VmGetServerPortsResponseStateEnum = "FORBIDDEN"
)

type VmGetServerPortsResponse struct {

	// 每个请求的序列号
	RequestId string `json:"requestId,omitempty"`

	// 页面国际化错误提示
	ErrorMessage string `json:"errorMessage,omitempty"`

	// 统一错误码
	ErrorCode string `json:"errorCode,omitempty"`

	// 返回状态码
	State VmGetServerPortsResponseStateEnum `json:"state,omitempty"`

	// 请求成功时返回的数据
	Body VmGetServerPortsResponseBody `json:"body,omitempty"`

	// 统一错误码的自定义参数
	ErrorParams *[]string `json:"errorParams,omitempty"`
}
