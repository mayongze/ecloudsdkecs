// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model


type VmlistServerPortsDetailRespResponseFixedIpDetailRespsBandwidthTypeEnum string

// List of BandwidthType
const (
    VmlistServerPortsDetailRespResponseFixedIpDetailResps_SHARED VmlistServerPortsDetailRespResponseFixedIpDetailRespsBandwidthTypeEnum = "shared"
    VmlistServerPortsDetailRespResponseFixedIpDetailResps_EXCLUSIVE VmlistServerPortsDetailRespResponseFixedIpDetailRespsBandwidthTypeEnum = "exclusive"
    VmlistServerPortsDetailRespResponseFixedIpDetailResps_IPV6QOS VmlistServerPortsDetailRespResponseFixedIpDetailRespsBandwidthTypeEnum = "ipv6Qos"
    VmlistServerPortsDetailRespResponseFixedIpDetailResps_IPV6SHARED VmlistServerPortsDetailRespResponseFixedIpDetailRespsBandwidthTypeEnum = "ipv6Shared"
)
type VmlistServerPortsDetailRespResponseFixedIpDetailRespsIpVersionEnum string

// List of IpVersion
const (
    VmlistServerPortsDetailRespResponseFixedIpDetailResps__4 VmlistServerPortsDetailRespResponseFixedIpDetailRespsIpVersionEnum = "4"
    VmlistServerPortsDetailRespResponseFixedIpDetailResps__6 VmlistServerPortsDetailRespResponseFixedIpDetailRespsIpVersionEnum = "6"
)

type VmlistServerPortsDetailRespResponseFixedIpDetailResps struct {

	// 带宽类型
	BandwidthType VmlistServerPortsDetailRespResponseFixedIpDetailRespsBandwidthTypeEnum `json:"bandwidthType,omitempty"`

	// 子网Id
	SubnetId string `json:"subnetId,omitempty"`

	// ipVsersion，v4或v6
	IpVersion VmlistServerPortsDetailRespResponseFixedIpDetailRespsIpVersionEnum `json:"ipVersion,omitempty"`

	// 公网IP 带宽Mb
	BandwidthSize int32 `json:"bandwidthSize,omitempty"`

	// 内网ip
	IpAddress string `json:"ipAddress,omitempty"`

	// IPv6带宽值Mb
	Ipv6BandwidthSize int32 `json:"ipv6BandwidthSize,omitempty"`

	// 公网IP
	PublicIP string `json:"publicIP,omitempty"`

	// 子网Name
	SubnetName string `json:"subnetName,omitempty"`
}
