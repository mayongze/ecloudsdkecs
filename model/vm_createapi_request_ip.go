// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model


type VmCreateapiRequestIpIpTypeEnum string

// List of IpType
const (
    VmCreateapiRequestIp_MOBILE VmCreateapiRequestIpIpTypeEnum = "MOBILE"
    VmCreateapiRequestIp_MULTI_LINE VmCreateapiRequestIpIpTypeEnum = "MULTI_LINE"
)

type VmCreateapiRequestIp struct {

	// 公网IP类型,MOBILE:弹性公网IP;MULTI_LINE:优享版弹性公网IP
	IpType VmCreateapiRequestIpIpTypeEnum `json:"ipType,omitempty"`
}
