// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
)

type VmGetServerDetailQuery struct {
	position.Query
	// detail
	Detail bool `json:"detail,omitempty"`
}
