package domain


type AlibabaIdleAffiliateCpaIncomeSummaryIdleAffiliateCommonResult struct {
    /*
        result     */
    Result  *AlibabaIdleAffiliateCpaIncomeSummaryCpaIncomeSummaryDTO `json:"result,omitempty" `

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

func (s *AlibabaIdleAffiliateCpaIncomeSummaryIdleAffiliateCommonResult) SetResult(v AlibabaIdleAffiliateCpaIncomeSummaryCpaIncomeSummaryDTO) *AlibabaIdleAffiliateCpaIncomeSummaryIdleAffiliateCommonResult {
    s.Result = &v
    return s
}
func (s *AlibabaIdleAffiliateCpaIncomeSummaryIdleAffiliateCommonResult) SetSuccess(v bool) *AlibabaIdleAffiliateCpaIncomeSummaryIdleAffiliateCommonResult {
    s.Success = &v
    return s
}
func (s *AlibabaIdleAffiliateCpaIncomeSummaryIdleAffiliateCommonResult) SetDisplayErrorMsg(v string) *AlibabaIdleAffiliateCpaIncomeSummaryIdleAffiliateCommonResult {
    s.DisplayErrorMsg = &v
    return s
}
func (s *AlibabaIdleAffiliateCpaIncomeSummaryIdleAffiliateCommonResult) SetErrCode(v string) *AlibabaIdleAffiliateCpaIncomeSummaryIdleAffiliateCommonResult {
    s.ErrCode = &v
    return s
}
func (s *AlibabaIdleAffiliateCpaIncomeSummaryIdleAffiliateCommonResult) SetErrMsg(v string) *AlibabaIdleAffiliateCpaIncomeSummaryIdleAffiliateCommonResult {
    s.ErrMsg = &v
    return s
}
