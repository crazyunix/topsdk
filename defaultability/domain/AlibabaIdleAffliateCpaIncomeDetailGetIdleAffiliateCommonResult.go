package domain


type AlibabaIdleAffliateCpaIncomeDetailGetIdleAffiliateCommonResult struct {
    /*
        result     */
    Result  *[]AlibabaIdleAffliateCpaIncomeDetailGetCpaBillDTO `json:"result,omitempty" `

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

func (s *AlibabaIdleAffliateCpaIncomeDetailGetIdleAffiliateCommonResult) SetResult(v []AlibabaIdleAffliateCpaIncomeDetailGetCpaBillDTO) *AlibabaIdleAffliateCpaIncomeDetailGetIdleAffiliateCommonResult {
    s.Result = &v
    return s
}
func (s *AlibabaIdleAffliateCpaIncomeDetailGetIdleAffiliateCommonResult) SetDisplayErrorMsg(v string) *AlibabaIdleAffliateCpaIncomeDetailGetIdleAffiliateCommonResult {
    s.DisplayErrorMsg = &v
    return s
}
func (s *AlibabaIdleAffliateCpaIncomeDetailGetIdleAffiliateCommonResult) SetSuccess(v bool) *AlibabaIdleAffliateCpaIncomeDetailGetIdleAffiliateCommonResult {
    s.Success = &v
    return s
}
func (s *AlibabaIdleAffliateCpaIncomeDetailGetIdleAffiliateCommonResult) SetErrCode(v string) *AlibabaIdleAffliateCpaIncomeDetailGetIdleAffiliateCommonResult {
    s.ErrCode = &v
    return s
}
func (s *AlibabaIdleAffliateCpaIncomeDetailGetIdleAffiliateCommonResult) SetErrMsg(v string) *AlibabaIdleAffliateCpaIncomeDetailGetIdleAffiliateCommonResult {
    s.ErrMsg = &v
    return s
}
