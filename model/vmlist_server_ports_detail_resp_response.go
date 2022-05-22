// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model


type VmlistServerPortsDetailRespResponseStateEnum string

// List of State
const (
    VmlistServerPortsDetailRespResponse_OK VmlistServerPortsDetailRespResponseStateEnum = "OK"
    VmlistServerPortsDetailRespResponse_ERROR VmlistServerPortsDetailRespResponseStateEnum = "ERROR"
    VmlistServerPortsDetailRespResponse_EXCEPTION VmlistServerPortsDetailRespResponseStateEnum = "EXCEPTION"
    VmlistServerPortsDetailRespResponse_ALARM VmlistServerPortsDetailRespResponseStateEnum = "ALARM"
    VmlistServerPortsDetailRespResponse_FORBIDDEN VmlistServerPortsDetailRespResponseStateEnum = "FORBIDDEN"
)

type VmlistServerPortsDetailRespResponse struct {

	// 每个请求的序列号
	RequestId string `json:"requestId,omitempty"`

	// 页面国际化错误提示
	ErrorMessage string `json:"errorMessage,omitempty"`

	// 统一错误码
	ErrorCode string `json:"errorCode,omitempty"`

	// 返回状态码
	State VmlistServerPortsDetailRespResponseStateEnum `json:"state,omitempty"`

	// 请求成功时返回的数据
	Body VmlistServerPortsDetailRespResponseBody `json:"body,omitempty"`

	// 统一错误码的自定义参数
	ErrorParams *[]string `json:"errorParams,omitempty"`
}
