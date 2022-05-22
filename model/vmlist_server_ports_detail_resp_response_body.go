// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model



type VmlistServerPortsDetailRespResponseBody struct {

	// 满足查询条件的总个数
	Total int `json:"total,omitempty"`

	// 分页查询返回的具体结构体
	Content *[]VmlistServerPortsDetailRespResponseContent `json:"content,omitempty"`
}
