package domain


type AlibabaIdleAffiliateCpaPeriodSummaryQueryIdleAffiliatePageResult struct {
    /*
        result     */
    Result  *[]AlibabaIdleAffiliateCpaPeriodSummaryQueryCpaPeriodBillDTO `json:"result,omitempty" `

    /*
        displayErrorMsg     */
    DisplayErrorMsg  *string `json:"display_error_msg,omitempty" `

    /*
        nextPage     */
    NextPage  *bool `json:"next_page,omitempty" `

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

func (s *AlibabaIdleAffiliateCpaPeriodSummaryQueryIdleAffiliatePageResult) SetResult(v []AlibabaIdleAffiliateCpaPeriodSummaryQueryCpaPeriodBillDTO) *AlibabaIdleAffiliateCpaPeriodSummaryQueryIdleAffiliatePageResult {
    s.Result = &v
    return s
}
func (s *AlibabaIdleAffiliateCpaPeriodSummaryQueryIdleAffiliatePageResult) SetDisplayErrorMsg(v string) *AlibabaIdleAffiliateCpaPeriodSummaryQueryIdleAffiliatePageResult {
    s.DisplayErrorMsg = &v
    return s
}
func (s *AlibabaIdleAffiliateCpaPeriodSummaryQueryIdleAffiliatePageResult) SetNextPage(v bool) *AlibabaIdleAffiliateCpaPeriodSummaryQueryIdleAffiliatePageResult {
    s.NextPage = &v
    return s
}
func (s *AlibabaIdleAffiliateCpaPeriodSummaryQueryIdleAffiliatePageResult) SetSuccess(v bool) *AlibabaIdleAffiliateCpaPeriodSummaryQueryIdleAffiliatePageResult {
    s.Success = &v
    return s
}
func (s *AlibabaIdleAffiliateCpaPeriodSummaryQueryIdleAffiliatePageResult) SetErrCode(v string) *AlibabaIdleAffiliateCpaPeriodSummaryQueryIdleAffiliatePageResult {
    s.ErrCode = &v
    return s
}
func (s *AlibabaIdleAffiliateCpaPeriodSummaryQueryIdleAffiliatePageResult) SetErrMsg(v string) *AlibabaIdleAffiliateCpaPeriodSummaryQueryIdleAffiliatePageResult {
    s.ErrMsg = &v
    return s
}
