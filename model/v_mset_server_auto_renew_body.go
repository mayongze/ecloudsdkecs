// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
)

type VMsetServerAutoRenewBody struct {
	position.Body
	// 设置/取消自动续订
	AutoRenew bool `json:"autoRenew,omitempty"`

	// 云主机id
	ServerId string `json:"serverId,omitempty"`
}
