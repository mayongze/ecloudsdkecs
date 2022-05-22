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

	// 备份保留数量
	PreservedNum int32 `json:"preservedNum,omitempty"`

	// 是否被删除
	Deleted bool `json:"deleted,omitempty"`

	// 策略id
	PolicyId string `json:"policyId,omitempty"`

	// 创建者
	CreatedBy string `json:"createdBy,omitempty"`

	// 备份重复周期 周日 SUN,周一 MON,周二 TUE,周三 WED,周四 THU,周五 FRI,周六 SAT
	BackupInterval VmGetServerDetailResponseServerBackupPolicyRespBackupIntervalEnum `json:"backupInterval,omitempty"`

	// 所属客户
	CustomerId string `json:"customerId,omitempty"`

	// 策略name
	Name string `json:"name,omitempty"`

	// 创建时间
	CreatedTime string `json:"createdTime,omitempty"`

	// 备份创建时间,每天0~23点整
	StartTime *[]string `json:"startTime,omitempty"`

	// 资源类型
	ResourceType VmGetServerDetailResponseServerBackupPolicyRespResourceTypeEnum `json:"resourceType,omitempty"`
}
