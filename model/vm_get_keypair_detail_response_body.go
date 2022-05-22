// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model



type VmGetKeypairDetailResponseBody struct {

	// 密钥对修改时间
	ModifiedTime string `json:"modifiedTime,omitempty"`

	// 密钥对私钥
	PrivateKey string `json:"privateKey,omitempty"`

	// 密钥对名称
	Name string `json:"name,omitempty"`

	// 密钥对创建时间
	CreatedTime string `json:"createdTime,omitempty"`

	// 密钥对id
	Id string `json:"id,omitempty"`

	// 公钥
	PublicKey string `json:"publicKey,omitempty"`

	// 密钥对可用区
	Region string `json:"region,omitempty"`

	// 密钥类型
	Type string `json:"type,omitempty"`

	// 密钥对指纹
	FingerPrint string `json:"fingerPrint,omitempty"`

	// 密钥对状态
	Status string `json:"status,omitempty"`
}
