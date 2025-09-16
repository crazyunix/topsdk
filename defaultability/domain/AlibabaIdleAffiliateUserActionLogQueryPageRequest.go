package domain


type AlibabaIdleAffiliateUserActionLogQueryPageRequest struct {
    /*
        分页大小，不能超过20     */
    PageSize  *int64 `json:"page_size,omitempty" `

    /*
        页码     */
    PageNum  *int64 `json:"page_num,omitempty" `

}

func (s *AlibabaIdleAffiliateUserActionLogQueryPageRequest) SetPageSize(v int64) *AlibabaIdleAffiliateUserActionLogQueryPageRequest {
    s.PageSize = &v
    return s
}
func (s *AlibabaIdleAffiliateUserActionLogQueryPageRequest) SetPageNum(v int64) *AlibabaIdleAffiliateUserActionLogQueryPageRequest {
    s.PageNum = &v
    return s
}
