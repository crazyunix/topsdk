package domain


type AlibabaIdleAffiliateCpsIncomeDetailsQueryIdleAffiliatePageResult struct {
    /*
        result     */
    Result  *[]AlibabaIdleAffiliateCpsIncomeDetailsQueryCommissionDetailDTO `json:"result,omitempty" `

    /*
        nextPage     */
    NextPage  *bool `json:"next_page,omitempty" `

    /*
        success     */
    Success  *bool `json:"success,omitempty" `

    /*
        错误码     */
    ErrCode  *string `json:"err_code,omitempty" `

    /*
        错误信息     */
    ErrMsg  *string `json:"err_msg,omitempty" `

    /*
        错误信息提示     */
    DisplayErrorMsg  *string `json:"display_error_msg,omitempty" `

}

func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryIdleAffiliatePageResult) SetResult(v []AlibabaIdleAffiliateCpsIncomeDetailsQueryCommissionDetailDTO) *AlibabaIdleAffiliateCpsIncomeDetailsQueryIdleAffiliatePageResult {
    s.Result = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryIdleAffiliatePageResult) SetNextPage(v bool) *AlibabaIdleAffiliateCpsIncomeDetailsQueryIdleAffiliatePageResult {
    s.NextPage = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryIdleAffiliatePageResult) SetSuccess(v bool) *AlibabaIdleAffiliateCpsIncomeDetailsQueryIdleAffiliatePageResult {
    s.Success = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryIdleAffiliatePageResult) SetErrCode(v string) *AlibabaIdleAffiliateCpsIncomeDetailsQueryIdleAffiliatePageResult {
    s.ErrCode = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryIdleAffiliatePageResult) SetErrMsg(v string) *AlibabaIdleAffiliateCpsIncomeDetailsQueryIdleAffiliatePageResult {
    s.ErrMsg = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryIdleAffiliatePageResult) SetDisplayErrorMsg(v string) *AlibabaIdleAffiliateCpsIncomeDetailsQueryIdleAffiliatePageResult {
    s.DisplayErrorMsg = &v
    return s
}
