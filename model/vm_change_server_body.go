// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
)

type VmChangeServerBody struct {
	position.Body
	// 计费方式变更时长(月/年)
	Duration int32 `json:"duration,omitempty"`

	// 计费方式
	BillingType string `json:"billingType,omitempty"`

	// 云主机id
	ServerId string `json:"serverId,omitempty"`
}
