package domain


type AlibabaIdleAffiliateCampaignGetOpenAffiliateCampaign struct {
    /*
        佣金规则图     */
    CommissionImgInfo  *AlibabaIdleAffiliateCampaignGetImgInfo `json:"commission_img_info,omitempty" `

    /*
        佣金title     */
    AmountTitle  *string `json:"amount_title,omitempty" `

    /*
        活动id     */
    CampaignId  *string `json:"campaign_id,omitempty" `

    /*
        主标题图片     */
    TitleImgInfo  *AlibabaIdleAffiliateCampaignGetImgInfo `json:"title_img_info,omitempty" `

    /*
        活动开始时间     */
    StartTime  *int64 `json:"start_time,omitempty" `

    /*
        活动结束时间     */
    EndTime  *int64 `json:"end_time,omitempty" `

    /*
        标题     */
    Title  *string `json:"title,omitempty" `

    /*
        commissionInfoList     */
    CommissionInfoList  *[]AlibabaIdleAffiliateCampaignGetCampaignCommissionInfo `json:"commission_info_list,omitempty" `

    /*
        活动是否在线     */
    Serving  *bool `json:"serving,omitempty" `

    /*
        活动类型 CPA / CPS     */
    CampaignType  *string `json:"campaign_type,omitempty" `

}

func (s *AlibabaIdleAffiliateCampaignGetOpenAffiliateCampaign) SetCommissionImgInfo(v AlibabaIdleAffiliateCampaignGetImgInfo) *AlibabaIdleAffiliateCampaignGetOpenAffiliateCampaign {
    s.CommissionImgInfo = &v
    return s
}
func (s *AlibabaIdleAffiliateCampaignGetOpenAffiliateCampaign) SetAmountTitle(v string) *AlibabaIdleAffiliateCampaignGetOpenAffiliateCampaign {
    s.AmountTitle = &v
    return s
}
func (s *AlibabaIdleAffiliateCampaignGetOpenAffiliateCampaign) SetCampaignId(v string) *AlibabaIdleAffiliateCampaignGetOpenAffiliateCampaign {
    s.CampaignId = &v
    return s
}
func (s *AlibabaIdleAffiliateCampaignGetOpenAffiliateCampaign) SetTitleImgInfo(v AlibabaIdleAffiliateCampaignGetImgInfo) *AlibabaIdleAffiliateCampaignGetOpenAffiliateCampaign {
    s.TitleImgInfo = &v
    return s
}
func (s *AlibabaIdleAffiliateCampaignGetOpenAffiliateCampaign) SetStartTime(v int64) *AlibabaIdleAffiliateCampaignGetOpenAffiliateCampaign {
    s.StartTime = &v
    return s
}
func (s *AlibabaIdleAffiliateCampaignGetOpenAffiliateCampaign) SetEndTime(v int64) *AlibabaIdleAffiliateCampaignGetOpenAffiliateCampaign {
    s.EndTime = &v
    return s
}
func (s *AlibabaIdleAffiliateCampaignGetOpenAffiliateCampaign) SetTitle(v string) *AlibabaIdleAffiliateCampaignGetOpenAffiliateCampaign {
    s.Title = &v
    return s
}
func (s *AlibabaIdleAffiliateCampaignGetOpenAffiliateCampaign) SetCommissionInfoList(v []AlibabaIdleAffiliateCampaignGetCampaignCommissionInfo) *AlibabaIdleAffiliateCampaignGetOpenAffiliateCampaign {
    s.CommissionInfoList = &v
    return s
}
func (s *AlibabaIdleAffiliateCampaignGetOpenAffiliateCampaign) SetServing(v bool) *AlibabaIdleAffiliateCampaignGetOpenAffiliateCampaign {
    s.Serving = &v
    return s
}
func (s *AlibabaIdleAffiliateCampaignGetOpenAffiliateCampaign) SetCampaignType(v string) *AlibabaIdleAffiliateCampaignGetOpenAffiliateCampaign {
    s.CampaignType = &v
    return s
}
