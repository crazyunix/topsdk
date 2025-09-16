package domain


type AlibabaIdleAffiliateIncomeSummaryIdleAffiliateCommonResult struct {
    /*
        result     */
    Result  *AlibabaIdleAffiliateIncomeSummaryCpsIncomeSummaryDTO `json:"result,omitempty" `

    /*
        displayErrorMsg     */
    DisplayErrorMsg  *string `json:"display_error_msg,omitempty" `

    /*
        success     */
    Success  *bool `json:"success,omitempty" `

    /*
        errCode     */
    ErrCode  *string `json:"err_code,omitempty" `

    /*
        errMsg     */
    ErrMsg  *string `json:"err_msg,omitempty" `

}

func (s *AlibabaIdleAffiliateIncomeSummaryIdleAffiliateCommonResult) SetResult(v AlibabaIdleAffiliateIncomeSummaryCpsIncomeSummaryDTO) *AlibabaIdleAffiliateIncomeSummaryIdleAffiliateCommonResult {
    s.Result = &v
    return s
}
func (s *AlibabaIdleAffiliateIncomeSummaryIdleAffiliateCommonResult) SetDisplayErrorMsg(v string) *AlibabaIdleAffiliateIncomeSummaryIdleAffiliateCommonResult {
    s.DisplayErrorMsg = &v
    return s
}
func (s *AlibabaIdleAffiliateIncomeSummaryIdleAffiliateCommonResult) SetSuccess(v bool) *AlibabaIdleAffiliateIncomeSummaryIdleAffiliateCommonResult {
    s.Success = &v
    return s
}
func (s *AlibabaIdleAffiliateIncomeSummaryIdleAffiliateCommonResult) SetErrCode(v string) *AlibabaIdleAffiliateIncomeSummaryIdleAffiliateCommonResult {
    s.ErrCode = &v
    return s
}
func (s *AlibabaIdleAffiliateIncomeSummaryIdleAffiliateCommonResult) SetErrMsg(v string) *AlibabaIdleAffiliateIncomeSummaryIdleAffiliateCommonResult {
    s.ErrMsg = &v
    return s
}
