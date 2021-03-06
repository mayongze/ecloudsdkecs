// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"

/*type VmlistServerPortsDetailRespQuery string

// List of EcStatus
const (
    VmlistServerPortsDetailRespQuery_ACTIVE VmlistServerPortsDetailRespQuery = "ACTIVE"
    VmlistServerPortsDetailRespQuery_BUILD VmlistServerPortsDetailRespQuery = "BUILD"
    VmlistServerPortsDetailRespQuery_REBUILD VmlistServerPortsDetailRespQuery = "REBUILD"
    VmlistServerPortsDetailRespQuery_RESCUE VmlistServerPortsDetailRespQuery = "RESCUE"
    VmlistServerPortsDetailRespQuery_SUSPENDED VmlistServerPortsDetailRespQuery = "SUSPENDED"
    VmlistServerPortsDetailRespQuery_PAUSED VmlistServerPortsDetailRespQuery = "PAUSED"
    VmlistServerPortsDetailRespQuery_RESIZE VmlistServerPortsDetailRespQuery = "RESIZE"
    VmlistServerPortsDetailRespQuery_VERIFY_RESIZE VmlistServerPortsDetailRespQuery = "VERIFY_RESIZE"
    VmlistServerPortsDetailRespQuery_REVERT_RESIZE VmlistServerPortsDetailRespQuery = "REVERT_RESIZE"
    VmlistServerPortsDetailRespQuery_PASSWORD VmlistServerPortsDetailRespQuery = "PASSWORD"
    VmlistServerPortsDetailRespQuery_REBOOT VmlistServerPortsDetailRespQuery = "REBOOT"
    VmlistServerPortsDetailRespQuery_HARD_REBOOT VmlistServerPortsDetailRespQuery = "HARD_REBOOT"
    VmlistServerPortsDetailRespQuery_DELETED VmlistServerPortsDetailRespQuery = "DELETED"
    VmlistServerPortsDetailRespQuery_UNKNOWN VmlistServerPortsDetailRespQuery = "UNKNOWN"
    VmlistServerPortsDetailRespQuery_ERROR VmlistServerPortsDetailRespQuery = "ERROR"
    VmlistServerPortsDetailRespQuery_STOPPED VmlistServerPortsDetailRespQuery = "STOPPED"
    VmlistServerPortsDetailRespQuery_SHUTOFF VmlistServerPortsDetailRespQuery = "SHUTOFF"
    VmlistServerPortsDetailRespQuery_MIGRATING VmlistServerPortsDetailRespQuery = "MIGRATING"
    VmlistServerPortsDetailRespQuery_SHELVED VmlistServerPortsDetailRespQuery = "SHELVED"
    VmlistServerPortsDetailRespQuery_SHELVED_OFFLOADED VmlistServerPortsDetailRespQuery = "SHELVED_OFFLOADED"
    VmlistServerPortsDetailRespQuery_SOFT_DELETED VmlistServerPortsDetailRespQuery = "SOFT_DELETED"
    VmlistServerPortsDetailRespQuery_UNRECOGNIZED VmlistServerPortsDetailRespQuery = "UNRECOGNIZED"
)*/
type VmlistServerPortsDetailRespQueryEcStatusEnum string

// List of EcStatus
const (
	VmlistServerPortsDetailRespQueryACTIVE            VmlistServerPortsDetailRespQueryEcStatusEnum = "ACTIVE"
	VmlistServerPortsDetailRespQueryBUILD             VmlistServerPortsDetailRespQueryEcStatusEnum = "BUILD"
	VmlistServerPortsDetailRespQueryREBUILD           VmlistServerPortsDetailRespQueryEcStatusEnum = "REBUILD"
	VmlistServerPortsDetailRespQueryRESCUE            VmlistServerPortsDetailRespQueryEcStatusEnum = "RESCUE"
	VmlistServerPortsDetailRespQuerySUSPENDED         VmlistServerPortsDetailRespQueryEcStatusEnum = "SUSPENDED"
	VmlistServerPortsDetailRespQueryPAUSED            VmlistServerPortsDetailRespQueryEcStatusEnum = "PAUSED"
	VmlistServerPortsDetailRespQueryRESIZE            VmlistServerPortsDetailRespQueryEcStatusEnum = "RESIZE"
	VmlistServerPortsDetailRespQueryVERIFY_RESIZE     VmlistServerPortsDetailRespQueryEcStatusEnum = "VERIFY_RESIZE"
	VmlistServerPortsDetailRespQueryREVERT_RESIZE     VmlistServerPortsDetailRespQueryEcStatusEnum = "REVERT_RESIZE"
	VmlistServerPortsDetailRespQueryPASSWORD          VmlistServerPortsDetailRespQueryEcStatusEnum = "PASSWORD"
	VmlistServerPortsDetailRespQueryREBOOT            VmlistServerPortsDetailRespQueryEcStatusEnum = "REBOOT"
	VmlistServerPortsDetailRespQueryHARD_REBOOT       VmlistServerPortsDetailRespQueryEcStatusEnum = "HARD_REBOOT"
	VmlistServerPortsDetailRespQueryDELETED           VmlistServerPortsDetailRespQueryEcStatusEnum = "DELETED"
	VmlistServerPortsDetailRespQueryUNKNOWN           VmlistServerPortsDetailRespQueryEcStatusEnum = "UNKNOWN"
	VmlistServerPortsDetailRespQueryERROR             VmlistServerPortsDetailRespQueryEcStatusEnum = "ERROR"
	VmlistServerPortsDetailRespQuerySTOPPED           VmlistServerPortsDetailRespQueryEcStatusEnum = "STOPPED"
	VmlistServerPortsDetailRespQuerySHUTOFF           VmlistServerPortsDetailRespQueryEcStatusEnum = "SHUTOFF"
	VmlistServerPortsDetailRespQueryMIGRATING         VmlistServerPortsDetailRespQueryEcStatusEnum = "MIGRATING"
	VmlistServerPortsDetailRespQuerySHELVED           VmlistServerPortsDetailRespQueryEcStatusEnum = "SHELVED"
	VmlistServerPortsDetailRespQuerySHELVED_OFFLOADED VmlistServerPortsDetailRespQueryEcStatusEnum = "SHELVED_OFFLOADED"
	VmlistServerPortsDetailRespQuerySOFT_DELETED      VmlistServerPortsDetailRespQueryEcStatusEnum = "SOFT_DELETED"
	VmlistServerPortsDetailRespQueryUNRECOGNIZED      VmlistServerPortsDetailRespQueryEcStatusEnum = "UNRECOGNIZED"
)

/*type VmlistServerPortsDetailRespQuery string

// List of ServerTypes
const (
    VmlistServerPortsDetailRespQuery_VM VmlistServerPortsDetailRespQuery = "VM"
    VmlistServerPortsDetailRespQuery_IRONIC VmlistServerPortsDetailRespQuery = "IRONIC"
    VmlistServerPortsDetailRespQuery_EBM VmlistServerPortsDetailRespQuery = "EBM"
)*/
type VmlistServerPortsDetailRespQueryServerTypesEnum string

// List of ServerTypes
const (
	VmlistServerPortsDetailRespQueryVM     VmlistServerPortsDetailRespQueryServerTypesEnum = "VM"
	VmlistServerPortsDetailRespQueryIRONIC VmlistServerPortsDetailRespQueryServerTypesEnum = "IRONIC"
	VmlistServerPortsDetailRespQueryEBM    VmlistServerPortsDetailRespQueryServerTypesEnum = "EBM"
)

/*type VmlistServerPortsDetailRespQuery string

// List of ImageOsTypes
const (
    VmlistServerPortsDetailRespQuery_LINUX VmlistServerPortsDetailRespQuery = "linux"
    VmlistServerPortsDetailRespQuery_WINDOWS VmlistServerPortsDetailRespQuery = "windows"
    VmlistServerPortsDetailRespQuery_OTHER VmlistServerPortsDetailRespQuery = "other"
)*/

type VmlistServerPortsDetailRespQueryImageOsTypesEnum string

// List of ImageOsTypes
const (
	VmlistServerPortsDetailRespQueryLINUX   VmlistServerPortsDetailRespQueryImageOsTypesEnum = "linux"
	VmlistServerPortsDetailRespQueryWINDOWS VmlistServerPortsDetailRespQueryImageOsTypesEnum = "windows"
	VmlistServerPortsDetailRespQueryOTHER   VmlistServerPortsDetailRespQueryImageOsTypesEnum = "other"
)

/*type VmlistServerPortsDetailRespQuery string

// List of ProductTypes
const (
    VmlistServerPortsDetailRespQuery_NORMAL VmlistServerPortsDetailRespQuery = "NORMAL"
    VmlistServerPortsDetailRespQuery_DE_CLOUD VmlistServerPortsDetailRespQuery = "DE_CLOUD"
    VmlistServerPortsDetailRespQuery_AUTOSCALING VmlistServerPortsDetailRespQuery = "AUTOSCALING"
    VmlistServerPortsDetailRespQuery_VO VmlistServerPortsDetailRespQuery = "VO"
    VmlistServerPortsDetailRespQuery_CDN VmlistServerPortsDetailRespQuery = "CDN"
    VmlistServerPortsDetailRespQuery_PAAS_MASTER VmlistServerPortsDetailRespQuery = "PAAS_MASTER"
    VmlistServerPortsDetailRespQuery_PAAS_SLAVE VmlistServerPortsDetailRespQuery = "PAAS_SLAVE"
    VmlistServerPortsDetailRespQuery_VCPE VmlistServerPortsDetailRespQuery = "VCPE"
    VmlistServerPortsDetailRespQuery_EMR VmlistServerPortsDetailRespQuery = "EMR"
    VmlistServerPortsDetailRespQuery_LOGAUDIT VmlistServerPortsDetailRespQuery = "LOGAUDIT"
    VmlistServerPortsDetailRespQuery_MSE VmlistServerPortsDetailRespQuery = "MSE"
)*/
type VmlistServerPortsDetailRespQueryProductTypesEnum string

// List of ProductTypes
const (
	VmlistServerPortsDetailRespQueryNORMAL      VmlistServerPortsDetailRespQueryProductTypesEnum = "NORMAL"
	VmlistServerPortsDetailRespQueryDE_CLOUD    VmlistServerPortsDetailRespQueryProductTypesEnum = "DE_CLOUD"
	VmlistServerPortsDetailRespQueryAUTOSCALING VmlistServerPortsDetailRespQueryProductTypesEnum = "AUTOSCALING"
	VmlistServerPortsDetailRespQueryVO          VmlistServerPortsDetailRespQueryProductTypesEnum = "VO"
	VmlistServerPortsDetailRespQueryCDN         VmlistServerPortsDetailRespQueryProductTypesEnum = "CDN"
	VmlistServerPortsDetailRespQueryPAAS_MASTER VmlistServerPortsDetailRespQueryProductTypesEnum = "PAAS_MASTER"
	VmlistServerPortsDetailRespQueryPAAS_SLAVE  VmlistServerPortsDetailRespQueryProductTypesEnum = "PAAS_SLAVE"
	VmlistServerPortsDetailRespQueryVCPE        VmlistServerPortsDetailRespQueryProductTypesEnum = "VCPE"
	VmlistServerPortsDetailRespQueryEMR         VmlistServerPortsDetailRespQueryProductTypesEnum = "EMR"
	VmlistServerPortsDetailRespQueryLOGAUDIT    VmlistServerPortsDetailRespQueryProductTypesEnum = "LOGAUDIT"
	VmlistServerPortsDetailRespQueryMSE         VmlistServerPortsDetailRespQueryProductTypesEnum = "MSE"
)

type VmlistServerPortsDetailRespQueryOpStatusEnum string

// List of OpStatus
const (
	VmlistServerPortsDetailRespQuery_CREATING             VmlistServerPortsDetailRespQueryOpStatusEnum = "CREATING"
	VmlistServerPortsDetailRespQuery_OK                   VmlistServerPortsDetailRespQueryOpStatusEnum = "OK"
	VmlistServerPortsDetailRespQuery_REBOOTING            VmlistServerPortsDetailRespQueryOpStatusEnum = "REBOOTING"
	VmlistServerPortsDetailRespQuery_FROZEN               VmlistServerPortsDetailRespQueryOpStatusEnum = "FROZEN"
	VmlistServerPortsDetailRespQuery_BACKUP_REBUILDING    VmlistServerPortsDetailRespQueryOpStatusEnum = "BACKUP_REBUILDING"
	VmlistServerPortsDetailRespQuery_REBUILDING           VmlistServerPortsDetailRespQueryOpStatusEnum = "REBUILDING"
	VmlistServerPortsDetailRespQuery_RESIZING             VmlistServerPortsDetailRespQueryOpStatusEnum = "RESIZING"
	VmlistServerPortsDetailRespQuery_SHELVING             VmlistServerPortsDetailRespQueryOpStatusEnum = "SHELVING"
	VmlistServerPortsDetailRespQuery_UNSHELVING           VmlistServerPortsDetailRespQueryOpStatusEnum = "UNSHELVING"
	VmlistServerPortsDetailRespQuery_STARTING             VmlistServerPortsDetailRespQueryOpStatusEnum = "STARTING"
	VmlistServerPortsDetailRespQuery_STOPING              VmlistServerPortsDetailRespQueryOpStatusEnum = "STOPING"
	VmlistServerPortsDetailRespQuery_PASSWORD_UPDATING    VmlistServerPortsDetailRespQueryOpStatusEnum = "PASSWORD_UPDATING"
	VmlistServerPortsDetailRespQuery_BACKUPING            VmlistServerPortsDetailRespQueryOpStatusEnum = "BACKUPING"
	VmlistServerPortsDetailRespQuery_IMAGE_CREATING       VmlistServerPortsDetailRespQueryOpStatusEnum = "IMAGE_CREATING"
	VmlistServerPortsDetailRespQuery_UNRECOGNIZED         VmlistServerPortsDetailRespQueryOpStatusEnum = "UNRECOGNIZED"
	VmlistServerPortsDetailRespQuery_BOOT_VOLUME_RESIZING VmlistServerPortsDetailRespQueryOpStatusEnum = "BOOT_VOLUME_RESIZING"
	VmlistServerPortsDetailRespQuery_DELETING             VmlistServerPortsDetailRespQueryOpStatusEnum = "DELETING"
	VmlistServerPortsDetailRespQuery_SECOND_DELETING      VmlistServerPortsDetailRespQueryOpStatusEnum = "SECOND_DELETING"
	VmlistServerPortsDetailRespQuery_MIGRATING            VmlistServerPortsDetailRespQueryOpStatusEnum = "MIGRATING"
)

type VmlistServerPortsDetailRespQuery struct {
	position.Query
	// ?????????????????????
	EcStatus VmlistServerPortsDetailRespQueryEcStatusEnum `json:"ecStatus,omitempty"`

	// ?????????????????????
	Visible bool `json:"visible,omitempty"`

	// ???????????????
	ServerTypes VmlistServerPortsDetailRespQueryServerTypesEnum `json:"serverTypes,omitempty"`

	// ??????ip??????
	PrivateIp string `json:"privateIp,omitempty"`

	// ???????????????????????????
	ImageOsTypes VmlistServerPortsDetailRespQueryImageOsTypesEnum `json:"imageOsTypes,omitempty"`

	// ?????????????????????
	KeyName string `json:"keyName,omitempty"`

	// ???????????????
	ServerName string `json:"serverName,omitempty"`

	// cpu??????
	Cpu int32 `json:"cpu,omitempty"`

	// ?????????????????????????????????
	ProductTypes VmlistServerPortsDetailRespQueryProductTypesEnum `json:"productTypes,omitempty"`

	// ??????ip??????
	PublicIp string `json:"publicIp,omitempty"`

	// ???????????????????????????
	QueryWordName string `json:"queryWordName,omitempty"`

	// ?????????id
	ServerId string `json:"serverId,omitempty"`

	// ?????????id
	SecurityGroupId string `json:"securityGroupId,omitempty"`

	// ???????????????
	Disk int32 `json:"disk,omitempty"`

	// ????????????????????????
	Size int32 `json:"size,omitempty"`

	// ????????????????????????????????????????????????
	BackupPolicyNeeded bool `json:"backupPolicyNeeded,omitempty"`

	// ?????????????????????id
	NetworkId string `json:"networkId,omitempty"`

	// ??????????????????
	Page int32 `json:"page,omitempty"`

	// ?????????op?????????
	OpStatus VmlistServerPortsDetailRespQueryOpStatusEnum `json:"opStatus,omitempty"`

	// ????????????
	Ram int32 `json:"ram,omitempty"`
}
