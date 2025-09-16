package domain


type AlibabaIdleAffiliateCpaIncomeDetailsQueryCpaBillQueryParams struct {
    /*
        0待发放 1待到账 2已到账 3已失效     */
    BillState  *int64 `json:"bill_state,omitempty" `

    /*
        分页入参     */
    PageRequest  *AlibabaIdleAffiliateCpaIncomeDetailsQueryPageRequest `json:"page_request,omitempty" `

    /*
        查询开始时间     */
    StartTime  *int64 `json:"start_time,omitempty" `

    /*
        查询结束时间     */
    EndTime  *int64 `json:"end_time,omitempty" `

}

func (s *AlibabaIdleAffiliateCpaIncomeDetailsQueryCpaBillQueryParams) SetBillState(v int64) *AlibabaIdleAffiliateCpaIncomeDetailsQueryCpaBillQueryParams {
    s.BillState = &v
    return s
}
func (s *AlibabaIdleAffiliateCpaIncomeDetailsQueryCpaBillQueryParams) SetPageRequest(v AlibabaIdleAffiliateCpaIncomeDetailsQueryPageRequest) *AlibabaIdleAffiliateCpaIncomeDetailsQueryCpaBillQueryParams {
    s.PageRequest = &v
    return s
}
func (s *AlibabaIdleAffiliateCpaIncomeDetailsQueryCpaBillQueryParams) SetStartTime(v int64) *AlibabaIdleAffiliateCpaIncomeDetailsQueryCpaBillQueryParams {
    s.StartTime = &v
    return s
}
func (s *AlibabaIdleAffiliateCpaIncomeDetailsQueryCpaBillQueryParams) SetEndTime(v int64) *AlibabaIdleAffiliateCpaIncomeDetailsQueryCpaBillQueryParams {
    s.EndTime = &v
    return s
}
