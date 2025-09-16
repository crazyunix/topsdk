package domain


type AlibabaIdleAffiliateMaterialQueryPageRequest struct {
    /*
        最大不超过20页 defalutValue:10    */
    PageSize  *int64 `json:"page_size,omitempty" `

    /*
        页码 defalutValue:1    */
    PageNum  *int64 `json:"page_num,omitempty" `

}

func (s *AlibabaIdleAffiliateMaterialQueryPageRequest) SetPageSize(v int64) *AlibabaIdleAffiliateMaterialQueryPageRequest {
    s.PageSize = &v
    return s
}
func (s *AlibabaIdleAffiliateMaterialQueryPageRequest) SetPageNum(v int64) *AlibabaIdleAffiliateMaterialQueryPageRequest {
    s.PageNum = &v
    return s
}
