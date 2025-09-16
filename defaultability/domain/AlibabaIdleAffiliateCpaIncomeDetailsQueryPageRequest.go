package domain


type AlibabaIdleAffiliateCpaIncomeDetailsQueryPageRequest struct {
    /*
        每页多少 最多20     */
    PageSize  *int64 `json:"page_size,omitempty" `

    /*
        页码     */
    PageNum  *int64 `json:"page_num,omitempty" `

}

func (s *AlibabaIdleAffiliateCpaIncomeDetailsQueryPageRequest) SetPageSize(v int64) *AlibabaIdleAffiliateCpaIncomeDetailsQueryPageRequest {
    s.PageSize = &v
    return s
}
func (s *AlibabaIdleAffiliateCpaIncomeDetailsQueryPageRequest) SetPageNum(v int64) *AlibabaIdleAffiliateCpaIncomeDetailsQueryPageRequest {
    s.PageNum = &v
    return s
}
