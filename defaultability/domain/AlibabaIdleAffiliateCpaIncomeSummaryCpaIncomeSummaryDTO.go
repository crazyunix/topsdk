package domain


type AlibabaIdleAffiliateCpaIncomeSummaryCpaIncomeSummaryDTO struct {
    /*
        累积金额(元)     */
    AccumulatedAmountText  *string `json:"accumulated_amount_text,omitempty" `

    /*
        交易中的金额(元)     */
    UnaccountedAmountText  *string `json:"unaccounted_amount_text,omitempty" `

    /*
        可提现金额(元)     */
    BalanceText  *string `json:"balance_text,omitempty" `

}

func (s *AlibabaIdleAffiliateCpaIncomeSummaryCpaIncomeSummaryDTO) SetAccumulatedAmountText(v string) *AlibabaIdleAffiliateCpaIncomeSummaryCpaIncomeSummaryDTO {
    s.AccumulatedAmountText = &v
    return s
}
func (s *AlibabaIdleAffiliateCpaIncomeSummaryCpaIncomeSummaryDTO) SetUnaccountedAmountText(v string) *AlibabaIdleAffiliateCpaIncomeSummaryCpaIncomeSummaryDTO {
    s.UnaccountedAmountText = &v
    return s
}
func (s *AlibabaIdleAffiliateCpaIncomeSummaryCpaIncomeSummaryDTO) SetBalanceText(v string) *AlibabaIdleAffiliateCpaIncomeSummaryCpaIncomeSummaryDTO {
    s.BalanceText = &v
    return s
}
