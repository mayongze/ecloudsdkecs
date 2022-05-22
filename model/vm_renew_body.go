// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
)

type VmRenewBody struct {
	position.Body
	// 续订时长(月/年)
	Duration int32 `json:"duration,omitempty"`

	// 云主机id
	ServerId string `json:"serverId,omitempty"`
}
