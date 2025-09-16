package domain


type AlibabaIdleAffiliateCpsOrderQueryIdleAffiliatePageResult struct {
    /*
        result     */
    Result  *[]AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO `json:"result,omitempty" `

    /*
        nextPage     */
    NextPage  *bool `json:"next_page,omitempty" `

    /*
        是否成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        错误码     */
    ErrCode  *string `json:"err_code,omitempty" `

    /*
        错误信息     */
    ErrMsg  *string `json:"err_msg,omitempty" `

    /*
        错误信息描述     */
    DisplayErrorMsg  *string `json:"display_error_msg,omitempty" `

}

func (s *AlibabaIdleAffiliateCpsOrderQueryIdleAffiliatePageResult) SetResult(v []AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO) *AlibabaIdleAffiliateCpsOrderQueryIdleAffiliatePageResult {
    s.Result = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsOrderQueryIdleAffiliatePageResult) SetNextPage(v bool) *AlibabaIdleAffiliateCpsOrderQueryIdleAffiliatePageResult {
    s.NextPage = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsOrderQueryIdleAffiliatePageResult) SetSuccess(v bool) *AlibabaIdleAffiliateCpsOrderQueryIdleAffiliatePageResult {
    s.Success = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsOrderQueryIdleAffiliatePageResult) SetErrCode(v string) *AlibabaIdleAffiliateCpsOrderQueryIdleAffiliatePageResult {
    s.ErrCode = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsOrderQueryIdleAffiliatePageResult) SetErrMsg(v string) *AlibabaIdleAffiliateCpsOrderQueryIdleAffiliatePageResult {
    s.ErrMsg = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsOrderQueryIdleAffiliatePageResult) SetDisplayErrorMsg(v string) *AlibabaIdleAffiliateCpsOrderQueryIdleAffiliatePageResult {
    s.DisplayErrorMsg = &v
    return s
}
