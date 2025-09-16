package domain


type AlibabaIdleAffiliateCampaignGetCampaignCommissionInfo struct {
    /*
        佣金     */
    CommissionAmount  *string `json:"commission_amount,omitempty" `

    /*
        佣金描述     */
    CommissionDesc  *string `json:"commission_desc,omitempty" `

}

func (s *AlibabaIdleAffiliateCampaignGetCampaignCommissionInfo) SetCommissionAmount(v string) *AlibabaIdleAffiliateCampaignGetCampaignCommissionInfo {
    s.CommissionAmount = &v
    return s
}
func (s *AlibabaIdleAffiliateCampaignGetCampaignCommissionInfo) SetCommissionDesc(v string) *AlibabaIdleAffiliateCampaignGetCampaignCommissionInfo {
    s.CommissionDesc = &v
    return s
}
