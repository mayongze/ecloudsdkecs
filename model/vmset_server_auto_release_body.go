// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
)

type VmsetServerAutoReleaseBody struct {
	position.Body
	// 自动释放时间,如需设置自动释放时间则传yyyy-MM-dd HH:mm:ss格式，如取消自动释放则不传该参数
	AutoReleaseTime string `json:"autoReleaseTime,omitempty"`

	// 云主机id
	ServerId string `json:"serverId,omitempty"`
}
