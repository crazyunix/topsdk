package domain


type AlibabaIdleAffiliateCpsIncomeDetailsQueryPageRequest struct {
    /*
        每页多少 最多20     */
    PageSize  *int64 `json:"page_size,omitempty" `

    /*
        页码     */
    PageNum  *int64 `json:"page_num,omitempty" `

}

func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryPageRequest) SetPageSize(v int64) *AlibabaIdleAffiliateCpsIncomeDetailsQueryPageRequest {
    s.PageSize = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryPageRequest) SetPageNum(v int64) *AlibabaIdleAffiliateCpsIncomeDetailsQueryPageRequest {
    s.PageNum = &v
    return s
}
