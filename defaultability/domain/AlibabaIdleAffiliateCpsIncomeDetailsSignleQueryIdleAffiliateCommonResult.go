package domain


type AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryIdleAffiliateCommonResult struct {
    /*
        result     */
    Result  *[]AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryCommissionDetailDTO `json:"result,omitempty" `

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

func (s *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryIdleAffiliateCommonResult) SetResult(v []AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryCommissionDetailDTO) *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryIdleAffiliateCommonResult {
    s.Result = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryIdleAffiliateCommonResult) SetSuccess(v bool) *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryIdleAffiliateCommonResult {
    s.Success = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryIdleAffiliateCommonResult) SetDisplayErrorMsg(v string) *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryIdleAffiliateCommonResult {
    s.DisplayErrorMsg = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryIdleAffiliateCommonResult) SetErrCode(v string) *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryIdleAffiliateCommonResult {
    s.ErrCode = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryIdleAffiliateCommonResult) SetErrMsg(v string) *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryIdleAffiliateCommonResult {
    s.ErrMsg = &v
    return s
}
