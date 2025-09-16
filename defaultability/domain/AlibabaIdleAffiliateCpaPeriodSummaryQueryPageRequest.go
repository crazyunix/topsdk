package domain


type AlibabaIdleAffiliateCpaPeriodSummaryQueryPageRequest struct {
    /*
        每页多少 最多20     */
    PageSize  *int64 `json:"page_size,omitempty" `

    /*
        页码     */
    PageNum  *int64 `json:"page_num,omitempty" `

}

func (s *AlibabaIdleAffiliateCpaPeriodSummaryQueryPageRequest) SetPageSize(v int64) *AlibabaIdleAffiliateCpaPeriodSummaryQueryPageRequest {
    s.PageSize = &v
    return s
}
func (s *AlibabaIdleAffiliateCpaPeriodSummaryQueryPageRequest) SetPageNum(v int64) *AlibabaIdleAffiliateCpaPeriodSummaryQueryPageRequest {
    s.PageNum = &v
    return s
}
