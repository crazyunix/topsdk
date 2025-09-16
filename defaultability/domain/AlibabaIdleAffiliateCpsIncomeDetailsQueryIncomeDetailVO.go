package domain


type AlibabaIdleAffiliateCpsIncomeDetailsQueryIncomeDetailVO struct {
    /*
        账期 yyyymm     */
    CreateMonth  *string `json:"create_month,omitempty" `

    /*
        0待发放 1待到账 2已到账 3已失效，不传默认查询所有状态     */
    BillState  *int64 `json:"bill_state,omitempty" `

    /*
        分页入参     */
    PageRequest  *AlibabaIdleAffiliateCpsIncomeDetailsQueryPageRequest `json:"page_request,omitempty" `

    /*
        佣金单查询-开始时间(以创单时间检索)     */
    StartCreateTimeStamp  *int64 `json:"start_create_time_stamp,omitempty" `

    /*
        佣金单查询-结束时间(以创单时间检索)     */
    EndCreateTimeStamp  *int64 `json:"end_create_time_stamp,omitempty" `

    /*
        佣金单查询-开始时间(以更新订单时间检索)     */
    StartUpdateTime  *int64 `json:"start_update_time,omitempty" `

    /*
        佣金单查询-结束时间(以更新订单时间检索)     */
    EndUpdateTime  *int64 `json:"end_update_time,omitempty" `

}

func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryIncomeDetailVO) SetCreateMonth(v string) *AlibabaIdleAffiliateCpsIncomeDetailsQueryIncomeDetailVO {
    s.CreateMonth = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryIncomeDetailVO) SetBillState(v int64) *AlibabaIdleAffiliateCpsIncomeDetailsQueryIncomeDetailVO {
    s.BillState = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryIncomeDetailVO) SetPageRequest(v AlibabaIdleAffiliateCpsIncomeDetailsQueryPageRequest) *AlibabaIdleAffiliateCpsIncomeDetailsQueryIncomeDetailVO {
    s.PageRequest = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryIncomeDetailVO) SetStartCreateTimeStamp(v int64) *AlibabaIdleAffiliateCpsIncomeDetailsQueryIncomeDetailVO {
    s.StartCreateTimeStamp = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryIncomeDetailVO) SetEndCreateTimeStamp(v int64) *AlibabaIdleAffiliateCpsIncomeDetailsQueryIncomeDetailVO {
    s.EndCreateTimeStamp = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryIncomeDetailVO) SetStartUpdateTime(v int64) *AlibabaIdleAffiliateCpsIncomeDetailsQueryIncomeDetailVO {
    s.StartUpdateTime = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryIncomeDetailVO) SetEndUpdateTime(v int64) *AlibabaIdleAffiliateCpsIncomeDetailsQueryIncomeDetailVO {
    s.EndUpdateTime = &v
    return s
}
