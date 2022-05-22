// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model


type VmGetServerTagsResponseStateEnum string

// List of State
const (
    VmGetServerTagsResponse_OK VmGetServerTagsResponseStateEnum = "OK"
    VmGetServerTagsResponse_ERROR VmGetServerTagsResponseStateEnum = "ERROR"
    VmGetServerTagsResponse_EXCEPTION VmGetServerTagsResponseStateEnum = "EXCEPTION"
    VmGetServerTagsResponse_ALARM VmGetServerTagsResponseStateEnum = "ALARM"
    VmGetServerTagsResponse_FORBIDDEN VmGetServerTagsResponseStateEnum = "FORBIDDEN"
)

type VmGetServerTagsResponse struct {

	// 每个请求的序列号
	RequestId string `json:"requestId,omitempty"`

	// 页面国际化错误提示
	ErrorMessage string `json:"errorMessage,omitempty"`

	// 统一错误码
	ErrorCode string `json:"errorCode,omitempty"`

	// 返回状态码
	State VmGetServerTagsResponseStateEnum `json:"state,omitempty"`

	// 请求成功时返回的数据
	Body *[]VmGetServerTagsResponseBody `json:"body,omitempty"`

	// 统一错误码的自定义参数
	ErrorParams *[]string `json:"errorParams,omitempty"`
}
