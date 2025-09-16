package domain


type AlibabaIdleAffiliateMaterialExactGetIdleAffiliateCommonResult struct {
    /*
        result     */
    Result  *[]AlibabaIdleAffiliateMaterialExactGetMaterialDTO `json:"result,omitempty" `

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

func (s *AlibabaIdleAffiliateMaterialExactGetIdleAffiliateCommonResult) SetResult(v []AlibabaIdleAffiliateMaterialExactGetMaterialDTO) *AlibabaIdleAffiliateMaterialExactGetIdleAffiliateCommonResult {
    s.Result = &v
    return s
}
func (s *AlibabaIdleAffiliateMaterialExactGetIdleAffiliateCommonResult) SetSuccess(v bool) *AlibabaIdleAffiliateMaterialExactGetIdleAffiliateCommonResult {
    s.Success = &v
    return s
}
func (s *AlibabaIdleAffiliateMaterialExactGetIdleAffiliateCommonResult) SetErrCode(v string) *AlibabaIdleAffiliateMaterialExactGetIdleAffiliateCommonResult {
    s.ErrCode = &v
    return s
}
func (s *AlibabaIdleAffiliateMaterialExactGetIdleAffiliateCommonResult) SetErrMsg(v string) *AlibabaIdleAffiliateMaterialExactGetIdleAffiliateCommonResult {
    s.ErrMsg = &v
    return s
}
