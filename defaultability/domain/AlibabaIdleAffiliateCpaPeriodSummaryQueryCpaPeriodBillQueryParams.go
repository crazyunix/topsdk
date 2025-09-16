package domain


type AlibabaIdleAffiliateCpaPeriodSummaryQueryCpaPeriodBillQueryParams struct {
    /*
        分页参数     */
    PageRequest  *AlibabaIdleAffiliateCpaPeriodSummaryQueryPageRequest `json:"page_request,omitempty" `

}

func (s *AlibabaIdleAffiliateCpaPeriodSummaryQueryCpaPeriodBillQueryParams) SetPageRequest(v AlibabaIdleAffiliateCpaPeriodSummaryQueryPageRequest) *AlibabaIdleAffiliateCpaPeriodSummaryQueryCpaPeriodBillQueryParams {
    s.PageRequest = &v
    return s
}
