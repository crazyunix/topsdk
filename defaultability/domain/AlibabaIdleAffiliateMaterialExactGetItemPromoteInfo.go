package domain


type AlibabaIdleAffiliateMaterialExactGetItemPromoteInfo struct {
    /*
        佣金比例     */
    CommissionRate  *string `json:"commission_rate,omitempty" `

    /*
        预估佣金     */
    EstimatedCommission  *string `json:"estimated_commission,omitempty" `

}

func (s *AlibabaIdleAffiliateMaterialExactGetItemPromoteInfo) SetCommissionRate(v string) *AlibabaIdleAffiliateMaterialExactGetItemPromoteInfo {
    s.CommissionRate = &v
    return s
}
func (s *AlibabaIdleAffiliateMaterialExactGetItemPromoteInfo) SetEstimatedCommission(v string) *AlibabaIdleAffiliateMaterialExactGetItemPromoteInfo {
    s.EstimatedCommission = &v
    return s
}
