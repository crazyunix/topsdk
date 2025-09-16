package domain


type AlibabaIdleAffiliateCampaignGetIdleAffiliateCommonResult struct {
    /*
        result     */
    Result  *[]AlibabaIdleAffiliateCampaignGetOpenAffiliateCampaign `json:"result,omitempty" `

    /*
        success     */
    Success  *bool `json:"success,omitempty" `

    /*
        错误信息描述     */
    DisplayErrorMsg  *string `json:"display_error_msg,omitempty" `

    /*
        错误码     */
    ErrCode  *string `json:"err_code,omitempty" `

    /*
        错误信息     */
    ErrMsg  *string `json:"err_msg,omitempty" `

}

func (s *AlibabaIdleAffiliateCampaignGetIdleAffiliateCommonResult) SetResult(v []AlibabaIdleAffiliateCampaignGetOpenAffiliateCampaign) *AlibabaIdleAffiliateCampaignGetIdleAffiliateCommonResult {
    s.Result = &v
    return s
}
func (s *AlibabaIdleAffiliateCampaignGetIdleAffiliateCommonResult) SetSuccess(v bool) *AlibabaIdleAffiliateCampaignGetIdleAffiliateCommonResult {
    s.Success = &v
    return s
}
func (s *AlibabaIdleAffiliateCampaignGetIdleAffiliateCommonResult) SetDisplayErrorMsg(v string) *AlibabaIdleAffiliateCampaignGetIdleAffiliateCommonResult {
    s.DisplayErrorMsg = &v
    return s
}
func (s *AlibabaIdleAffiliateCampaignGetIdleAffiliateCommonResult) SetErrCode(v string) *AlibabaIdleAffiliateCampaignGetIdleAffiliateCommonResult {
    s.ErrCode = &v
    return s
}
func (s *AlibabaIdleAffiliateCampaignGetIdleAffiliateCommonResult) SetErrMsg(v string) *AlibabaIdleAffiliateCampaignGetIdleAffiliateCommonResult {
    s.ErrMsg = &v
    return s
}
