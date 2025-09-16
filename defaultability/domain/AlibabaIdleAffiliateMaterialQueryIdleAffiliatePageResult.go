package domain


type AlibabaIdleAffiliateMaterialQueryIdleAffiliatePageResult struct {
    /*
        result     */
    Result  *[]AlibabaIdleAffiliateMaterialQueryMaterialDTO `json:"result,omitempty" `

    /*
        错误信息描述     */
    DisplayErrorMsg  *string `json:"display_error_msg,omitempty" `

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

}

func (s *AlibabaIdleAffiliateMaterialQueryIdleAffiliatePageResult) SetResult(v []AlibabaIdleAffiliateMaterialQueryMaterialDTO) *AlibabaIdleAffiliateMaterialQueryIdleAffiliatePageResult {
    s.Result = &v
    return s
}
func (s *AlibabaIdleAffiliateMaterialQueryIdleAffiliatePageResult) SetDisplayErrorMsg(v string) *AlibabaIdleAffiliateMaterialQueryIdleAffiliatePageResult {
    s.DisplayErrorMsg = &v
    return s
}
func (s *AlibabaIdleAffiliateMaterialQueryIdleAffiliatePageResult) SetNextPage(v bool) *AlibabaIdleAffiliateMaterialQueryIdleAffiliatePageResult {
    s.NextPage = &v
    return s
}
func (s *AlibabaIdleAffiliateMaterialQueryIdleAffiliatePageResult) SetSuccess(v bool) *AlibabaIdleAffiliateMaterialQueryIdleAffiliatePageResult {
    s.Success = &v
    return s
}
func (s *AlibabaIdleAffiliateMaterialQueryIdleAffiliatePageResult) SetErrCode(v string) *AlibabaIdleAffiliateMaterialQueryIdleAffiliatePageResult {
    s.ErrCode = &v
    return s
}
func (s *AlibabaIdleAffiliateMaterialQueryIdleAffiliatePageResult) SetErrMsg(v string) *AlibabaIdleAffiliateMaterialQueryIdleAffiliatePageResult {
    s.ErrMsg = &v
    return s
}
