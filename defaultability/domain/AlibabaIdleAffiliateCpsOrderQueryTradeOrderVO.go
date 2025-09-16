package domain


type AlibabaIdleAffiliateCpsOrderQueryTradeOrderVO struct {
    /*
        分页参数     */
    PageRequest  *AlibabaIdleAffiliateCpsOrderQueryPageRequest `json:"page_request,omitempty" `

}

func (s *AlibabaIdleAffiliateCpsOrderQueryTradeOrderVO) SetPageRequest(v AlibabaIdleAffiliateCpsOrderQueryPageRequest) *AlibabaIdleAffiliateCpsOrderQueryTradeOrderVO {
    s.PageRequest = &v
    return s
}
