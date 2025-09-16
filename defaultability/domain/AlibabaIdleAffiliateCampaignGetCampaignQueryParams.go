package domain


type AlibabaIdleAffiliateCampaignGetCampaignQueryParams struct {
    /*
        活动类型 CPA / CPS     */
    CampaignType  *string `json:"campaign_type,omitempty" `

}

func (s *AlibabaIdleAffiliateCampaignGetCampaignQueryParams) SetCampaignType(v string) *AlibabaIdleAffiliateCampaignGetCampaignQueryParams {
    s.CampaignType = &v
    return s
}
