package domain

type TaobaoTbkDgMaterialRecommendStateSubsidyInfo struct {
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

func (s *TaobaoTbkDgMaterialRecommendStateSubsidyInfo) SetMaxRebate(v int64) *TaobaoTbkDgMaterialRecommendStateSubsidyInfo {
	s.MaxRebate = &v
	return s
}
func (s *TaobaoTbkDgMaterialRecommendStateSubsidyInfo) SetMinRebate(v int64) *TaobaoTbkDgMaterialRecommendStateSubsidyInfo {
	s.MinRebate = &v
	return s
}
func (s *TaobaoTbkDgMaterialRecommendStateSubsidyInfo) SetMaxDiscount(v string) *TaobaoTbkDgMaterialRecommendStateSubsidyInfo {
	s.MaxDiscount = &v
	return s
}
func (s *TaobaoTbkDgMaterialRecommendStateSubsidyInfo) SetMinDiscount(v string) *TaobaoTbkDgMaterialRecommendStateSubsidyInfo {
	s.MinDiscount = &v
	return s
}
func (s *TaobaoTbkDgMaterialRecommendStateSubsidyInfo) SetProvinceList(v []string) *TaobaoTbkDgMaterialRecommendStateSubsidyInfo {
	s.ProvinceList = &v
	return s
}
