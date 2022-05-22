// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
)

type NmBindServerTagsBody struct {
	position.Body
	// 云主机id
	ServerId string `json:"serverId,omitempty"`

	// 标签
	Tags *[]NmBindServerTagsRequestTags `json:"tags,omitempty"`
}
