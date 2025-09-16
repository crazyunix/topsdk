package request

import (
	"github.com/crazyunix/topsdk/defaultability/domain"

	"github.com/crazyunix/topsdk/util"
)

type AlibabaIdleAffiliateCampaignGetRequest struct {
	/*
	   活动查询入参body     */
	CampaignQueryParams *domain.AlibabaIdleAffiliateCampaignGetCampaignQueryParams `json:"campaign_query_params" required:"true" `
}

func (s *AlibabaIdleAffiliateCampaignGetRequest) SetCampaignQueryParams(v domain.AlibabaIdleAffiliateCampaignGetCampaignQueryParams) *AlibabaIdleAffiliateCampaignGetRequest {
	s.CampaignQueryParams = &v
	return s
}

func (req *AlibabaIdleAffiliateCampaignGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.CampaignQueryParams != nil {
		paramMap["campaign_query_params"] = util.ConvertStruct(*req.CampaignQueryParams)
	}
	return paramMap
}

func (req *AlibabaIdleAffiliateCampaignGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
