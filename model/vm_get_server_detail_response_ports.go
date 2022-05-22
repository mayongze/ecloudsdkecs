// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model



type VmGetServerDetailResponsePorts struct {

	// mac地址
	MacAddress string `json:"macAddress,omitempty"`

	// 公网ip列表
	IpId *[]string `json:"ipId,omitempty"`

	// 内网ip列表
	PrivateIp *[]string `json:"privateIp,omitempty"`

	// 网卡名称
	Name string `json:"name,omitempty"`

	// 公网ip列表
	PublicIp *[]string `json:"publicIp,omitempty"`

	// 网卡id
	Id string `json:"id,omitempty"`
}
