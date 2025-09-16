package domain


type AlibabaIdleAffiliateUserActionLogQueryUserActionLogQueryParams struct {
    /*
        查询日期（格式：yyyymmdd），时间以创建结算单时间为准，支持30天内历史数据查询     */
    QueryUserDate  *string `json:"query_user_date,omitempty" `

    /*
        分页参数     */
    PageRequest  *AlibabaIdleAffiliateUserActionLogQueryPageRequest `json:"page_request,omitempty" `

}

func (s *AlibabaIdleAffiliateUserActionLogQueryUserActionLogQueryParams) SetQueryUserDate(v string) *AlibabaIdleAffiliateUserActionLogQueryUserActionLogQueryParams {
    s.QueryUserDate = &v
    return s
}
func (s *AlibabaIdleAffiliateUserActionLogQueryUserActionLogQueryParams) SetPageRequest(v AlibabaIdleAffiliateUserActionLogQueryPageRequest) *AlibabaIdleAffiliateUserActionLogQueryUserActionLogQueryParams {
    s.PageRequest = &v
    return s
}
