// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package ecloudsdkecs

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/config"
	"gitlab.ecloud.com/ecloud/ecloudsdkecs/model"
)

type Client struct {
	APIClient   *ecloudsdkcore.APIClient
	config      *config.Config
	httpRequest *ecloudsdkcore.HttpRequest
}

func NewClient(config *config.Config) *Client {
	client := &Client{}
	client.config = config
	apiClient := ecloudsdkcore.NewAPIClient()
	httpRequest := ecloudsdkcore.NewDefaultHttpRequest()
	httpRequest.Product = product
	httpRequest.Version = version
	httpRequest.SdkVersion = sdkVersion
	client.httpRequest = httpRequest
	client.APIClient = apiClient
	return client
}

func NewClientByCustomized(config *config.Config, httpRequest *ecloudsdkcore.HttpRequest) *Client {
	client := &Client{}
	client.config = config
	apiClient := ecloudsdkcore.NewAPIClient()
	httpRequest.Product = product
	httpRequest.Version = version
	httpRequest.SdkVersion = sdkVersion
	client.httpRequest = httpRequest
	client.APIClient = apiClient
	return client
}

const (
	product    string = "ECS"
	version    string = "v1"
	sdkVersion string = "1.0.0"
)

// VmCreateKeypair
func (c *Client) VmCreateKeypair(request *model.VmCreateKeypairRequest) (*model.VmCreateKeypairResponse, error) {
	c.httpRequest.Action = "vmCreateKeypair"
	c.httpRequest.Body = request
	var returnValue = &model.VmCreateKeypairResponse{}
	if _, err := c.APIClient.Excute(c.httpRequest, c.config, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// VmListKeypair
func (c *Client) VmListKeypair(request *model.VmListKeypairRequest) (*model.VmListKeypairResponse, error) {
	c.httpRequest.Action = "VmListKeypair"
	c.httpRequest.Body = request
	var returnValue = &model.VmListKeypairResponse{}
	if _, err := c.APIClient.Excute(c.httpRequest, c.config, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// VmDeletekeypair
func (c *Client) VmDeletekeypair(request *model.VmDeletekeypairRequest) (*model.VmDeletekeypairResponse, error) {
	c.httpRequest.Action = "vmDeletekeypair"
	c.httpRequest.Body = request
	var returnValue = &model.VmDeletekeypairResponse{}
	if _, err := c.APIClient.Excute(c.httpRequest, c.config, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// VmGetServerDetail
func (c *Client) VmGetServerDetail(request *model.VmGetServerDetailRequest) (*model.VmGetServerDetailResponse, error) {
	c.httpRequest.Action = "VmGetServerDetail"
	c.httpRequest.Body = request
	var returnValue = &model.VmGetServerDetailResponse{}
	if _, err := c.APIClient.Excute(c.httpRequest, c.config, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// VmGetKeypairDetail
func (c *Client) VmGetKeypairDetail(request *model.VmGetKeypairDetailRequest) (*model.VmGetKeypairDetailResponse, error) {
	c.httpRequest.Action = "vmGetKeypairDetail"
	c.httpRequest.Body = request
	var returnValue = &model.VmGetKeypairDetailResponse{}
	if _, err := c.APIClient.Excute(c.httpRequest, c.config, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// VmGetServerTags
func (c *Client) VmGetServerTags(request *model.VmGetServerTagsRequest) (*model.VmGetServerTagsResponse, error) {
	c.httpRequest.Action = "VmGetServerTags"
	c.httpRequest.Body = request
	var returnValue = &model.VmGetServerTagsResponse{}
	if _, err := c.APIClient.Excute(c.httpRequest, c.config, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// VmBindServerKeypair
func (c *Client) VmBindServerKeypair(request *model.VmBindServerKeypairRequest) (*model.VmBindServerKeypairResponse, error) {
	c.httpRequest.Action = "vmBindServerKeypair"
	c.httpRequest.Body = request
	var returnValue = &model.VmBindServerKeypairResponse{}
	if _, err := c.APIClient.Excute(c.httpRequest, c.config, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// VmUpdate
func (c *Client) VmUpdate(request *model.VmUpdateRequest) (*model.VmUpdateResponse, error) {
	c.httpRequest.Action = "vmUpdate"
	c.httpRequest.Body = request
	var returnValue = &model.VmUpdateResponse{}
	if _, err := c.APIClient.Excute(c.httpRequest, c.config, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// VmGetServerChangeable
func (c *Client) VmGetServerChangeable(request *model.VmGetServerChangeableRequest) (*model.VmGetServerChangeableResponse, error) {
	c.httpRequest.Action = "vmGetServerChangeable"
	c.httpRequest.Body = request
	var returnValue = &model.VmGetServerChangeableResponse{}
	if _, err := c.APIClient.Excute(c.httpRequest, c.config, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// NmBindServerTags
func (c *Client) NmBindServerTags(request *model.NmBindServerTagsRequest) (*model.NmBindServerTagsResponse, error) {
	c.httpRequest.Action = "nmBindServerTags"
	c.httpRequest.Body = request
	var returnValue = &model.NmBindServerTagsResponse{}
	if _, err := c.APIClient.Excute(c.httpRequest, c.config, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// VmupdateName
func (c *Client) VmupdateName(request *model.VmupdateNameRequest) (*model.VmupdateNameResponse, error) {
	c.httpRequest.Action = "vmupdateName"
	c.httpRequest.Body = request
	var returnValue = &model.VmupdateNameResponse{}
	if _, err := c.APIClient.Excute(c.httpRequest, c.config, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// VmUpdatePassword
func (c *Client) VmUpdatePassword(request *model.VmUpdatePasswordRequest) (*model.VmUpdatePasswordResponse, error) {
	c.httpRequest.Action = "vmUpdatePassword"
	c.httpRequest.Body = request
	var returnValue = &model.VmUpdatePasswordResponse{}
	if _, err := c.APIClient.Excute(c.httpRequest, c.config, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// VmUnbindServerTags
func (c *Client) VmUnbindServerTags(request *model.VmUnbindServerTagsRequest) (*model.VmUnbindServerTagsResponse, error) {
	c.httpRequest.Action = "vmUnbindServerTags"
	c.httpRequest.Body = request
	var returnValue = &model.VmUnbindServerTagsResponse{}
	if _, err := c.APIClient.Excute(c.httpRequest, c.config, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// VmsetServerAutoRelease
func (c *Client) VmsetServerAutoRelease(request *model.VmsetServerAutoReleaseRequest) (*model.VmsetServerAutoReleaseResponse, error) {
	c.httpRequest.Action = "VmsetServerAutoRelease"
	c.httpRequest.Body = request
	var returnValue = &model.VmsetServerAutoReleaseResponse{}
	if _, err := c.APIClient.Excute(c.httpRequest, c.config, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// VMsetServerAutoRenew
func (c *Client) VMsetServerAutoRenew(request *model.VMsetServerAutoRenewRequest) (*model.VMsetServerAutoRenewResponse, error) {
	c.httpRequest.Action = "VMsetServerAutoRenew"
	c.httpRequest.Body = request
	var returnValue = &model.VMsetServerAutoRenewResponse{}
	if _, err := c.APIClient.Excute(c.httpRequest, c.config, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// VmRebuild
func (c *Client) VmRebuild(request *model.VmRebuildRequest) (*model.VmRebuildResponse, error) {
	c.httpRequest.Action = "vmRebuild"
	c.httpRequest.Body = request
	var returnValue = &model.VmRebuildResponse{}
	if _, err := c.APIClient.Excute(c.httpRequest, c.config, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// VmGetServerPorts
func (c *Client) VmGetServerPorts(request *model.VmGetServerPortsRequest) (*model.VmGetServerPortsResponse, error) {
	c.httpRequest.Action = "vmGetServerPorts"
	c.httpRequest.Body = request
	var returnValue = &model.VmGetServerPortsResponse{}
	if _, err := c.APIClient.Excute(c.httpRequest, c.config, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// VmGetRebuildImage
func (c *Client) VmGetRebuildImage(request *model.VmGetRebuildImageRequest) (*model.VmGetRebuildImageResponse, error) {
	c.httpRequest.Action = "vmGetRebuildImage"
	c.httpRequest.Body = request
	var returnValue = &model.VmGetRebuildImageResponse{}
	if _, err := c.APIClient.Excute(c.httpRequest, c.config, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// VmChangeServer
func (c *Client) VmChangeServer(request *model.VmChangeServerRequest) (*model.VmChangeServerResponse, error) {
	c.httpRequest.Action = "VmChangeServer"
	c.httpRequest.Body = request
	var returnValue = &model.VmChangeServerResponse{}
	if _, err := c.APIClient.Excute(c.httpRequest, c.config, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// VmReboot
func (c *Client) VmReboot(request *model.VmRebootRequest) (*model.VmRebootResponse, error) {
	c.httpRequest.Action = "vmReboot"
	c.httpRequest.Body = request
	var returnValue = &model.VmRebootResponse{}
	if _, err := c.APIClient.Excute(c.httpRequest, c.config, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// VmCreateapi
func (c *Client) VmCreateapi(request *model.VmCreateapiRequest) (*model.VmCreateapiResponse, error) {
	c.httpRequest.Action = "VmCreateapi"
	c.httpRequest.Body = request
	var returnValue = &model.VmCreateapiResponse{}
	if _, err := c.APIClient.Excute(c.httpRequest, c.config, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// VmqueryPrice
func (c *Client) VmqueryPrice(request *model.VmqueryPriceRequest) (*model.VmqueryPriceResponse, error) {
	c.httpRequest.Action = "VmqueryPrice"
	c.httpRequest.Body = request
	var returnValue = &model.VmqueryPriceResponse{}
	if _, err := c.APIClient.Excute(c.httpRequest, c.config, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// VmStart
func (c *Client) VmStart(request *model.VmStartRequest) (*model.VmStartResponse, error) {
	c.httpRequest.Action = "vmStart"
	c.httpRequest.Body = request
	var returnValue = &model.VmStartResponse{}
	if _, err := c.APIClient.Excute(c.httpRequest, c.config, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// VmStop
func (c *Client) VmStop(request *model.VmStopRequest) (*model.VmStopResponse, error) {
	c.httpRequest.Action = "vmStop"
	c.httpRequest.Body = request
	var returnValue = &model.VmStopResponse{}
	if _, err := c.APIClient.Excute(c.httpRequest, c.config, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// VmRenew
func (c *Client) VmRenew(request *model.VmRenewRequest) (*model.VmRenewResponse, error) {
	c.httpRequest.Action = "VmRenew"
	c.httpRequest.Body = request
	var returnValue = &model.VmRenewResponse{}
	if _, err := c.APIClient.Excute(c.httpRequest, c.config, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// VmChange
func (c *Client) VmChange(request *model.VmChangeRequest) (*model.VmChangeResponse, error) {
	c.httpRequest.Action = "VmChange"
	c.httpRequest.Body = request
	var returnValue = &model.VmChangeResponse{}
	if _, err := c.APIClient.Excute(c.httpRequest, c.config, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// VmGetVnc
func (c *Client) VmGetVnc(request *model.VmGetVncRequest) (*model.VmGetVncResponse, error) {
	c.httpRequest.Action = "vmGetVnc"
	c.httpRequest.Body = request
	var returnValue = &model.VmGetVncResponse{}
	if _, err := c.APIClient.Excute(c.httpRequest, c.config, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// VmDis
func (c *Client) VmDis(request *model.VmDisRequest) (*model.VmDisResponse, error) {
	c.httpRequest.Action = "VmDis"
	c.httpRequest.Body = request
	var returnValue = &model.VmDisResponse{}
	if _, err := c.APIClient.Excute(c.httpRequest, c.config, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// VmlistServerPortsDetailResp
func (c *Client) VmlistServerPortsDetailResp(request *model.VmlistServerPortsDetailRespRequest) (*model.VmlistServerPortsDetailRespResponse, error) {
	c.httpRequest.Action = "VmlistServerPortsDetailResp"
	c.httpRequest.Body = request
	var returnValue = &model.VmlistServerPortsDetailRespResponse{}
	if _, err := c.APIClient.Excute(c.httpRequest, c.config, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// VmDelete
func (c *Client) VmDelete(request *model.VmDeleteRequest) (*model.VmDeleteResponse, error) {
	c.httpRequest.Action = "vmDelete"
	c.httpRequest.Body = request
	var returnValue = &model.VmDeleteResponse{}
	if _, err := c.APIClient.Excute(c.httpRequest, c.config, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}
