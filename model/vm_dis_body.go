// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
)

type VmDisBody struct {
	position.Body
	// 云主机系统盘大小
	Size int32 `json:"size,omitempty"`

	// 预付费下单成功跳转地址，前端填写
	ReturnUrl string `json:"returnUrl,omitempty"`

	// 云主机id
	ServerId string `json:"serverId,omitempty"`
}
