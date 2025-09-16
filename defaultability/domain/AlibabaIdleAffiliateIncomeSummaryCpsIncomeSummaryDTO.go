package domain


type AlibabaIdleAffiliateIncomeSummaryCpsIncomeSummaryDTO struct {
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

func (s *AlibabaIdleAffiliateIncomeSummaryCpsIncomeSummaryDTO) SetAccumulatedAmountText(v string) *AlibabaIdleAffiliateIncomeSummaryCpsIncomeSummaryDTO {
    s.AccumulatedAmountText = &v
    return s
}
func (s *AlibabaIdleAffiliateIncomeSummaryCpsIncomeSummaryDTO) SetUnaccountedAmountText(v string) *AlibabaIdleAffiliateIncomeSummaryCpsIncomeSummaryDTO {
    s.UnaccountedAmountText = &v
    return s
}
func (s *AlibabaIdleAffiliateIncomeSummaryCpsIncomeSummaryDTO) SetBalanceText(v string) *AlibabaIdleAffiliateIncomeSummaryCpsIncomeSummaryDTO {
    s.BalanceText = &v
    return s
}
