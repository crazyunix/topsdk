package domain


type AlibabaIdleAffiliateMaterialQueryItemPromoteInfo struct {
    /*
        佣金比例     */
    CommissionRate  *string `json:"commission_rate,omitempty" `

    /*
        预估佣金     */
    EstimatedCommission  *string `json:"estimated_commission,omitempty" `

}

func (s *AlibabaIdleAffiliateMaterialQueryItemPromoteInfo) SetCommissionRate(v string) *AlibabaIdleAffiliateMaterialQueryItemPromoteInfo {
    s.CommissionRate = &v
    return s
}
func (s *AlibabaIdleAffiliateMaterialQueryItemPromoteInfo) SetEstimatedCommission(v string) *AlibabaIdleAffiliateMaterialQueryItemPromoteInfo {
    s.EstimatedCommission = &v
    return s
}
