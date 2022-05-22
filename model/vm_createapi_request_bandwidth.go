// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model


type VmCreateapiRequestBandwidthChargeModeEnum string

// List of ChargeMode
const (
    VmCreateapiRequestBandwidth_BANDWIDTHCHARGE VmCreateapiRequestBandwidthChargeModeEnum = "bandwidthCharge"
    VmCreateapiRequestBandwidth_TRAFFICCHARGE VmCreateapiRequestBandwidthChargeModeEnum = "trafficCharge"
)

type VmCreateapiRequestBandwidth struct {

	// 公网IP带宽计费方式,bandwidthCharge:按带宽计费;trafficCharge:按流量计费
	ChargeMode VmCreateapiRequestBandwidthChargeModeEnum `json:"chargeMode,omitempty"`

	// 公网IP带宽大小
	BandwidthSize int32 `json:"bandwidthSize,omitempty"`
}
