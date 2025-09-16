package domain


type AlibabaIdleAffiliateMaterialGuideGetIdleAffiliateCommonResult struct {
    /*
        materialGuideDTO     */
    Result  *AlibabaIdleAffiliateMaterialGuideGetMaterialGuideDTO `json:"result,omitempty" `

    /*
        本次请求结果成功与否     */
    Success  *bool `json:"success,omitempty" `

    /*
        错误码     */
    ErrCode  *string `json:"err_code,omitempty" `

    /*
        错误信息     */
    ErrMsg  *string `json:"err_msg,omitempty" `

    /*
        错误信息     */
    DisplayErrorMsg  *string `json:"display_error_msg,omitempty" `

}

func (s *AlibabaIdleAffiliateMaterialGuideGetIdleAffiliateCommonResult) SetResult(v AlibabaIdleAffiliateMaterialGuideGetMaterialGuideDTO) *AlibabaIdleAffiliateMaterialGuideGetIdleAffiliateCommonResult {
    s.Result = &v
    return s
}
func (s *AlibabaIdleAffiliateMaterialGuideGetIdleAffiliateCommonResult) SetSuccess(v bool) *AlibabaIdleAffiliateMaterialGuideGetIdleAffiliateCommonResult {
    s.Success = &v
    return s
}
func (s *AlibabaIdleAffiliateMaterialGuideGetIdleAffiliateCommonResult) SetErrCode(v string) *AlibabaIdleAffiliateMaterialGuideGetIdleAffiliateCommonResult {
    s.ErrCode = &v
    return s
}
func (s *AlibabaIdleAffiliateMaterialGuideGetIdleAffiliateCommonResult) SetErrMsg(v string) *AlibabaIdleAffiliateMaterialGuideGetIdleAffiliateCommonResult {
    s.ErrMsg = &v
    return s
}
func (s *AlibabaIdleAffiliateMaterialGuideGetIdleAffiliateCommonResult) SetDisplayErrorMsg(v string) *AlibabaIdleAffiliateMaterialGuideGetIdleAffiliateCommonResult {
    s.DisplayErrorMsg = &v
    return s
}
