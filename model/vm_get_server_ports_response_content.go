// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model



type VmGetServerPortsResponseContent struct {

	// mac地址
	MacAddress string `json:"macAddress,omitempty"`

	// 网卡名称
	Name string `json:"name,omitempty"`

	// 公网IP的带宽,单位:Mb
	BandwidthSize int32 `json:"bandwidthSize,omitempty"`

	// 内网IP信息列表
	FixedIpResps *[]VmGetServerPortsResponseFixedIpResps `json:"fixedIpResps,omitempty"`

	// 创建时间
	CreatedTime string `json:"createdTime,omitempty"`

	// 公网IP
	PublicIp string `json:"publicIp,omitempty"`

	// 网卡id
	Id string `json:"id,omitempty"`

	// 子网名称
	SubnetName string `json:"subnetName,omitempty"`

	// 安全组列表
	SecurityGroupResps *[]VmGetServerPortsResponseSecurityGroupResps `json:"securityGroupResps,omitempty"`
}
