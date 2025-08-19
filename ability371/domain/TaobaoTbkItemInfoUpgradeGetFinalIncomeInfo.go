package domain

type TaobaoTbkItemInfoUpgradeGetFinalIncomeInfo struct {
	/*
	   商品佣金比率     */
	CommissionRate *string `json:"commission_rate,omitempty" `

	/*
	   商品佣金金额     */
	CommissionAmount *string `json:"commission_amount,omitempty" `

	/*
	   补贴比率     */
	SubsidyRate *string `json:"subsidy_rate,omitempty" `

	/*
	   补贴金额     */
	SubsidyAmount *string `json:"subsidy_amount,omitempty" `

	/*
	   补贴上限；仅在单笔订单命中补贴上限时返回结果否则出参为空     */
	SubsidyUpperLimit *string `json:"subsidy_upper_limit,omitempty" `

	/*
	   补贴类型     */
	SubsidyType *string `json:"subsidy_type,omitempty" `
}

func (s *TaobaoTbkItemInfoUpgradeGetFinalIncomeInfo) SetCommissionRate(v string) *TaobaoTbkItemInfoUpgradeGetFinalIncomeInfo {
	s.CommissionRate = &v
	return s
}
func (s *TaobaoTbkItemInfoUpgradeGetFinalIncomeInfo) SetCommissionAmount(v string) *TaobaoTbkItemInfoUpgradeGetFinalIncomeInfo {
	s.CommissionAmount = &v
	return s
}
func (s *TaobaoTbkItemInfoUpgradeGetFinalIncomeInfo) SetSubsidyRate(v string) *TaobaoTbkItemInfoUpgradeGetFinalIncomeInfo {
	s.SubsidyRate = &v
	return s
}
func (s *TaobaoTbkItemInfoUpgradeGetFinalIncomeInfo) SetSubsidyAmount(v string) *TaobaoTbkItemInfoUpgradeGetFinalIncomeInfo {
	s.SubsidyAmount = &v
	return s
}
func (s *TaobaoTbkItemInfoUpgradeGetFinalIncomeInfo) SetSubsidyUpperLimit(v string) *TaobaoTbkItemInfoUpgradeGetFinalIncomeInfo {
	s.SubsidyUpperLimit = &v
	return s
}
func (s *TaobaoTbkItemInfoUpgradeGetFinalIncomeInfo) SetSubsidyType(v string) *TaobaoTbkItemInfoUpgradeGetFinalIncomeInfo {
	s.SubsidyType = &v
	return s
}
