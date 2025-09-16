package domain


type AlibabaIdleAffiliateCpaIncomeDetailsQueryIdleAffiliatePageResult struct {
    /*
        result     */
    Result  *[]AlibabaIdleAffiliateCpaIncomeDetailsQueryCpaBillDTO `json:"result,omitempty" `

    /*
        错误信息描述     */
    DisplayErrorMsg  *string `json:"display_error_msg,omitempty" `

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

}

func (s *AlibabaIdleAffiliateCpaIncomeDetailsQueryIdleAffiliatePageResult) SetResult(v []AlibabaIdleAffiliateCpaIncomeDetailsQueryCpaBillDTO) *AlibabaIdleAffiliateCpaIncomeDetailsQueryIdleAffiliatePageResult {
    s.Result = &v
    return s
}
func (s *AlibabaIdleAffiliateCpaIncomeDetailsQueryIdleAffiliatePageResult) SetDisplayErrorMsg(v string) *AlibabaIdleAffiliateCpaIncomeDetailsQueryIdleAffiliatePageResult {
    s.DisplayErrorMsg = &v
    return s
}
func (s *AlibabaIdleAffiliateCpaIncomeDetailsQueryIdleAffiliatePageResult) SetNextPage(v bool) *AlibabaIdleAffiliateCpaIncomeDetailsQueryIdleAffiliatePageResult {
    s.NextPage = &v
    return s
}
func (s *AlibabaIdleAffiliateCpaIncomeDetailsQueryIdleAffiliatePageResult) SetSuccess(v bool) *AlibabaIdleAffiliateCpaIncomeDetailsQueryIdleAffiliatePageResult {
    s.Success = &v
    return s
}
func (s *AlibabaIdleAffiliateCpaIncomeDetailsQueryIdleAffiliatePageResult) SetErrCode(v string) *AlibabaIdleAffiliateCpaIncomeDetailsQueryIdleAffiliatePageResult {
    s.ErrCode = &v
    return s
}
func (s *AlibabaIdleAffiliateCpaIncomeDetailsQueryIdleAffiliatePageResult) SetErrMsg(v string) *AlibabaIdleAffiliateCpaIncomeDetailsQueryIdleAffiliatePageResult {
    s.ErrMsg = &v
    return s
}
