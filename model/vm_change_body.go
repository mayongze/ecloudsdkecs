// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
)

type VmChangeBody struct {
	position.Body
	// cpu
	Cpu int32 `json:"cpu,omitempty"`

	// 预付费下单成功跳转地址，前端填写
	ReturnUrl string `json:"returnUrl,omitempty"`

	// 云主机id
	ServerId string `json:"serverId,omitempty"`

	// 内存
	Ram int32 `json:"ram,omitempty"`
}
