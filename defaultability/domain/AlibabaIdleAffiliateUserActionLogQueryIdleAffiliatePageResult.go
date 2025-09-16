package domain


type AlibabaIdleAffiliateUserActionLogQueryIdleAffiliatePageResult struct {
    /*
        用户行为     */
    Result  *[]AlibabaIdleAffiliateUserActionLogQueryUserActionLogDTO `json:"result,omitempty" `

    /*
        nextpage     */
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

func (s *AlibabaIdleAffiliateUserActionLogQueryIdleAffiliatePageResult) SetResult(v []AlibabaIdleAffiliateUserActionLogQueryUserActionLogDTO) *AlibabaIdleAffiliateUserActionLogQueryIdleAffiliatePageResult {
    s.Result = &v
    return s
}
func (s *AlibabaIdleAffiliateUserActionLogQueryIdleAffiliatePageResult) SetNextPage(v bool) *AlibabaIdleAffiliateUserActionLogQueryIdleAffiliatePageResult {
    s.NextPage = &v
    return s
}
func (s *AlibabaIdleAffiliateUserActionLogQueryIdleAffiliatePageResult) SetSuccess(v bool) *AlibabaIdleAffiliateUserActionLogQueryIdleAffiliatePageResult {
    s.Success = &v
    return s
}
func (s *AlibabaIdleAffiliateUserActionLogQueryIdleAffiliatePageResult) SetErrCode(v string) *AlibabaIdleAffiliateUserActionLogQueryIdleAffiliatePageResult {
    s.ErrCode = &v
    return s
}
func (s *AlibabaIdleAffiliateUserActionLogQueryIdleAffiliatePageResult) SetErrMsg(v string) *AlibabaIdleAffiliateUserActionLogQueryIdleAffiliatePageResult {
    s.ErrMsg = &v
    return s
}
func (s *AlibabaIdleAffiliateUserActionLogQueryIdleAffiliatePageResult) SetDisplayErrorMsg(v string) *AlibabaIdleAffiliateUserActionLogQueryIdleAffiliatePageResult {
    s.DisplayErrorMsg = &v
    return s
}
