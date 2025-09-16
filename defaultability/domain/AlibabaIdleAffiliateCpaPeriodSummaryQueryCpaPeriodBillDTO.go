package domain


type AlibabaIdleAffiliateCpaPeriodSummaryQueryCpaPeriodBillDTO struct {
    /*
        到账金额     */
    RealSettleCycleAmount  *string `json:"real_settle_cycle_amount,omitempty" `

    /*
        到账比例(企业身份关注)     */
    AccountedPercentText  *string `json:"accounted_percent_text,omitempty" `

    /*
        个税(个人身份关注)     */
    IndividualTaxAmountText  *string `json:"individual_tax_amount_text,omitempty" `

    /*
        奖励金额     */
    SettleCycleAmount  *string `json:"settle_cycle_amount,omitempty" `

    /*
        账期     */
    BillCycle  *string `json:"bill_cycle,omitempty" `

}

func (s *AlibabaIdleAffiliateCpaPeriodSummaryQueryCpaPeriodBillDTO) SetRealSettleCycleAmount(v string) *AlibabaIdleAffiliateCpaPeriodSummaryQueryCpaPeriodBillDTO {
    s.RealSettleCycleAmount = &v
    return s
}
func (s *AlibabaIdleAffiliateCpaPeriodSummaryQueryCpaPeriodBillDTO) SetAccountedPercentText(v string) *AlibabaIdleAffiliateCpaPeriodSummaryQueryCpaPeriodBillDTO {
    s.AccountedPercentText = &v
    return s
}
func (s *AlibabaIdleAffiliateCpaPeriodSummaryQueryCpaPeriodBillDTO) SetIndividualTaxAmountText(v string) *AlibabaIdleAffiliateCpaPeriodSummaryQueryCpaPeriodBillDTO {
    s.IndividualTaxAmountText = &v
    return s
}
func (s *AlibabaIdleAffiliateCpaPeriodSummaryQueryCpaPeriodBillDTO) SetSettleCycleAmount(v string) *AlibabaIdleAffiliateCpaPeriodSummaryQueryCpaPeriodBillDTO {
    s.SettleCycleAmount = &v
    return s
}
func (s *AlibabaIdleAffiliateCpaPeriodSummaryQueryCpaPeriodBillDTO) SetBillCycle(v string) *AlibabaIdleAffiliateCpaPeriodSummaryQueryCpaPeriodBillDTO {
    s.BillCycle = &v
    return s
}
