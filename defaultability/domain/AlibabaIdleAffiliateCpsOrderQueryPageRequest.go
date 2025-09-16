package domain


type AlibabaIdleAffiliateCpsOrderQueryPageRequest struct {
    /*
        分页大小 defalutValue:1    */
    PageSize  *int64 `json:"page_size,omitempty" `

    /*
        页码 defalutValue:10    */
    PageNum  *int64 `json:"page_num,omitempty" `

}

func (s *AlibabaIdleAffiliateCpsOrderQueryPageRequest) SetPageSize(v int64) *AlibabaIdleAffiliateCpsOrderQueryPageRequest {
    s.PageSize = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsOrderQueryPageRequest) SetPageNum(v int64) *AlibabaIdleAffiliateCpsOrderQueryPageRequest {
    s.PageNum = &v
    return s
}
