// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model



type VmlistServerPortsDetailRespResponsePortDetail struct {

	// vpc名称
	VpcName string `json:"vpcName,omitempty"`

	// 公网IP v6地址
	FipV6Address string `json:"fipV6Address,omitempty"`

	// routerId
	RouterId string `json:"routerId,omitempty"`

	// 内网ip
	PrivateIp string `json:"privateIp,omitempty"`

	// fixedIpResps
	FixedIpDetailResps *[]VmlistServerPortsDetailRespResponseFixedIpDetailResps `json:"fixedIpDetailResps,omitempty"`

	// 带宽大小
	FipBandwidthSize int32 `json:"fipBandwidthSize,omitempty"`

	// portName
	PortName string `json:"portName,omitempty"`

	// portId
	PortId string `json:"portId,omitempty"`

	// 子网名称
	SubnetName string `json:"subnetName,omitempty"`

	// 公网IP地址
	FipAddress string `json:"fipAddress,omitempty"`
}
