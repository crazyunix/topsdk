package domain

type TaobaoTbkItemInfoUpgradeGetStateSubsidyInfo struct {
	/*
	   最高优惠比例（%），入参具体IP时，仅展示该IP最高优惠比例。     */
	MaxRebate *int64 `json:"max_rebate,omitempty" `

	/*
	   最低优惠比例（%），入参具体IP时，仅展示该IP最低优惠比例。     */
	MinRebate *int64 `json:"min_rebate,omitempty" `

	/*
	   最高优惠金额（元，如15代表15 元），入参具体IP时，仅展示该IP最高优惠金额     */
	MaxDiscount *string `json:"max_discount,omitempty" `

	/*
	   最低优惠金额（元，如15代表15 元），入参具体IP时，仅展示该IP最低优惠金额     */
	MinDiscount *string `json:"min_discount,omitempty" `

	/*
	   国补生效区域（省份）。不入参IP时展示各可用省份；入参IP时，全国可用商品展示各可用省份，区域可用商品仅展示IP对应省份。     */
	ProvinceList *[]string `json:"province_list,omitempty" `
}

func (s *TaobaoTbkItemInfoUpgradeGetStateSubsidyInfo) SetMaxRebate(v int64) *TaobaoTbkItemInfoUpgradeGetStateSubsidyInfo {
	s.MaxRebate = &v
	return s
}
func (s *TaobaoTbkItemInfoUpgradeGetStateSubsidyInfo) SetMinRebate(v int64) *TaobaoTbkItemInfoUpgradeGetStateSubsidyInfo {
	s.MinRebate = &v
	return s
}
func (s *TaobaoTbkItemInfoUpgradeGetStateSubsidyInfo) SetMaxDiscount(v string) *TaobaoTbkItemInfoUpgradeGetStateSubsidyInfo {
	s.MaxDiscount = &v
	return s
}
func (s *TaobaoTbkItemInfoUpgradeGetStateSubsidyInfo) SetMinDiscount(v string) *TaobaoTbkItemInfoUpgradeGetStateSubsidyInfo {
	s.MinDiscount = &v
	return s
}
func (s *TaobaoTbkItemInfoUpgradeGetStateSubsidyInfo) SetProvinceList(v []string) *TaobaoTbkItemInfoUpgradeGetStateSubsidyInfo {
	s.ProvinceList = &v
	return s
}
