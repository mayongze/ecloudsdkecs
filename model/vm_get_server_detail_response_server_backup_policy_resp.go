// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

/*type VmGetServerDetailResponseServerBackupPolicyResp string

// List of BackupInterval
const (
    VmGetServerDetailResponseServerBackupPolicyResp_SUN VmGetServerDetailResponseServerBackupPolicyResp = "SUN"
    VmGetServerDetailResponseServerBackupPolicyResp_MON VmGetServerDetailResponseServerBackupPolicyResp = "MON"
    VmGetServerDetailResponseServerBackupPolicyResp_TUE VmGetServerDetailResponseServerBackupPolicyResp = "TUE"
    VmGetServerDetailResponseServerBackupPolicyResp_WED VmGetServerDetailResponseServerBackupPolicyResp = "WED"
    VmGetServerDetailResponseServerBackupPolicyResp_THU VmGetServerDetailResponseServerBackupPolicyResp = "THU"
    VmGetServerDetailResponseServerBackupPolicyResp_FRI VmGetServerDetailResponseServerBackupPolicyResp = "FRI"
    VmGetServerDetailResponseServerBackupPolicyResp_SAT VmGetServerDetailResponseServerBackupPolicyResp = "SAT"
)*/
type VmGetServerDetailResponseServerBackupPolicyRespBackupIntervalEnum string

// List of BackupInterval
const (
	VmGetServerDetailResponseServerBackupPolicyRespSUN VmGetServerDetailResponseServerBackupPolicyRespBackupIntervalEnum = "SUN"
	VmGetServerDetailResponseServerBackupPolicyRespMON VmGetServerDetailResponseServerBackupPolicyRespBackupIntervalEnum = "MON"
	VmGetServerDetailResponseServerBackupPolicyRespTUE VmGetServerDetailResponseServerBackupPolicyRespBackupIntervalEnum = "TUE"
	VmGetServerDetailResponseServerBackupPolicyRespWED VmGetServerDetailResponseServerBackupPolicyRespBackupIntervalEnum = "WED"
	VmGetServerDetailResponseServerBackupPolicyRespTHU VmGetServerDetailResponseServerBackupPolicyRespBackupIntervalEnum = "THU"
	VmGetServerDetailResponseServerBackupPolicyRespFRI VmGetServerDetailResponseServerBackupPolicyRespBackupIntervalEnum = "FRI"
	VmGetServerDetailResponseServerBackupPolicyRespSAT VmGetServerDetailResponseServerBackupPolicyRespBackupIntervalEnum = "SAT"
)

type VmGetServerDetailResponseServerBackupPolicyRespResourceTypeEnum string

// List of ResourceType
const (
	VmGetServerDetailResponseServerBackupPolicyResp_SERVER VmGetServerDetailResponseServerBackupPolicyRespResourceTypeEnum = "SERVER"
	VmGetServerDetailResponseServerBackupPolicyResp_VOLUME VmGetServerDetailResponseServerBackupPolicyRespResourceTypeEnum = "VOLUME"
)

type VmGetServerDetailResponseServerBackupPolicyResp struct {

	// ??????????????????
	PreservedNum int32 `json:"preservedNum,omitempty"`

	// ???????????????
	Deleted bool `json:"deleted,omitempty"`

	// ??????id
	PolicyId string `json:"policyId,omitempty"`

	// ?????????
	CreatedBy string `json:"createdBy,omitempty"`

	// ?????????????????? ?????? SUN,?????? MON,?????? TUE,?????? WED,?????? THU,?????? FRI,?????? SAT
	BackupInterval VmGetServerDetailResponseServerBackupPolicyRespBackupIntervalEnum `json:"backupInterval,omitempty"`

	// ????????????
	CustomerId string `json:"customerId,omitempty"`

	// ??????name
	Name string `json:"name,omitempty"`

	// ????????????
	CreatedTime string `json:"createdTime,omitempty"`

	// ??????????????????,??????0~23??????
	StartTime *[]string `json:"startTime,omitempty"`

	// ????????????
	ResourceType VmGetServerDetailResponseServerBackupPolicyRespResourceTypeEnum `json:"resourceType,omitempty"`
}
