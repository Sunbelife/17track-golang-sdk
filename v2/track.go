package track17

import (
	"net/http"

	"github.com/Sunbelife/17track-golang-sdk/v2"
	modelsv2 "github.com/Sunbelife/17track-golang-sdk/v2/models/v2"
)

// Client 17track客户端
type Client struct {
	C *track17.Client
}

// NewClient 新客户端
func NewClient(secret string, configures ...track17.Configurer) (*Client, error) {
	client, err := track17.NewClient(secret, configures...)
	if err != nil {
		return &Client{}, err
	}
	client.WithApiVersion("/track/v2")
	return &Client{
		C: client,
	}, nil
}

// Registe 创建物流追踪
func (c Client) Registe(in []*modelsv2.TrackReq) (*modelsv2.TrackResp, *track17.Error) {
	var resp modelsv2.TrackResp
	err := c.C.Do(http.MethodPost, "/register", in, &resp)
	return &resp, err
}

// GetInfo 获取物流详情
func (c Client) GetInfo(in []*modelsv2.TrackReq) (*modelsv2.TrackResp, *track17.Error) {
	var resp modelsv2.TrackResp
	err := c.C.Do(http.MethodPost, "/gettrackinfo", in, &resp)
	return &resp, err
}

// Delete 删除物流追踪
func (c Client) Delete(in []*modelsv2.TrackReq) (*modelsv2.TrackResp, *track17.Error) {
	var resp modelsv2.TrackResp
	err := c.C.Do(http.MethodPost, "/deletetrack", in, &resp)
	return &resp, err
}

// Stop 停止追踪
func (c Client) Stop(in []*modelsv2.TrackReq) (*modelsv2.TrackResp, *track17.Error) {
	var resp modelsv2.TrackResp
	err := c.C.Do(http.MethodPost, "/stoptrack", in, &resp)
	return &resp, err
}

// Retrack 重启追踪
func (c Client) Retrack(in []*modelsv2.TrackReq) (*modelsv2.TrackResp, *track17.Error) {
	var resp modelsv2.TrackResp
	err := c.C.Do(http.MethodPost, "/retrack", in, &resp)
	return &resp, err
}

// Push 重新推送
func (c Client) Push(in []*modelsv2.TrackReq) (*modelsv2.TrackResp, *track17.Error) {
	var resp modelsv2.TrackResp
	err := c.C.Do(http.MethodPost, "/push", in, &resp)
	return &resp, err
}
