// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
)

type VmRebuildBody struct {
	position.Body
	// 镜像id
	ImageId string `json:"imageId,omitempty"`

	// 自定义参数
	UserData string `json:"userData,omitempty"`

	// 云主机密码
	AdminPass string `json:"adminPass,omitempty"`

	// 云主机id
	ServerId string `json:"serverId,omitempty"`
}
