// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model


type VmGetServerPortsResponseFixedIpRespsIpVersionEnum string

// List of IpVersion
const (
    VmGetServerPortsResponseFixedIpResps__4 VmGetServerPortsResponseFixedIpRespsIpVersionEnum = "4"
    VmGetServerPortsResponseFixedIpResps__6 VmGetServerPortsResponseFixedIpRespsIpVersionEnum = "6"
)

type VmGetServerPortsResponseFixedIpResps struct {

	// ip版本：v4/v6
	IpVersion VmGetServerPortsResponseFixedIpRespsIpVersionEnum `json:"ipVersion,omitempty"`

	// 内网ip地址
	IpAddress string `json:"ipAddress,omitempty"`
}
