package domain


type AlibabaIdleAffiliateGeneralLinkConvertIdleAffiliateCommonResult struct {
    /*
        result     */
    Result  *AlibabaIdleAffiliateGeneralLinkConvertLinkConvertDTO `json:"result,omitempty" `

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
        错误信息     */
    DisplayErrorMsg  *string `json:"display_error_msg,omitempty" `

}

func (s *AlibabaIdleAffiliateGeneralLinkConvertIdleAffiliateCommonResult) SetResult(v AlibabaIdleAffiliateGeneralLinkConvertLinkConvertDTO) *AlibabaIdleAffiliateGeneralLinkConvertIdleAffiliateCommonResult {
    s.Result = &v
    return s
}
func (s *AlibabaIdleAffiliateGeneralLinkConvertIdleAffiliateCommonResult) SetSuccess(v bool) *AlibabaIdleAffiliateGeneralLinkConvertIdleAffiliateCommonResult {
    s.Success = &v
    return s
}
func (s *AlibabaIdleAffiliateGeneralLinkConvertIdleAffiliateCommonResult) SetErrCode(v string) *AlibabaIdleAffiliateGeneralLinkConvertIdleAffiliateCommonResult {
    s.ErrCode = &v
    return s
}
func (s *AlibabaIdleAffiliateGeneralLinkConvertIdleAffiliateCommonResult) SetErrMsg(v string) *AlibabaIdleAffiliateGeneralLinkConvertIdleAffiliateCommonResult {
    s.ErrMsg = &v
    return s
}
func (s *AlibabaIdleAffiliateGeneralLinkConvertIdleAffiliateCommonResult) SetDisplayErrorMsg(v string) *AlibabaIdleAffiliateGeneralLinkConvertIdleAffiliateCommonResult {
    s.DisplayErrorMsg = &v
    return s
}
