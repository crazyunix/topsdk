package domain

type TaobaoTbkDgGeneralLinkConvertMaterialPromotionInfoDTO struct {
	/*
	   商品收入比率(%)：商品佣金比率+补贴比率。同物料类接口income_rate     */
	CommissionRate *string `json:"commission_rate,omitempty" `

	/*
	   佣金类型。MKT表示营销计划，SP表示定向计划，COMMON表示通用计划，ZX表示自选计划     */
	CommissionType *string `json:"commission_type,omitempty" `

	/*
	   多件价需拍件数     */
	MultipleItemsPricesCount *int64 `json:"multiple_items_prices_count,omitempty" `
}

func (s *TaobaoTbkDgGeneralLinkConvertMaterialPromotionInfoDTO) SetCommissionRate(v string) *TaobaoTbkDgGeneralLinkConvertMaterialPromotionInfoDTO {
	s.CommissionRate = &v
	return s
}
func (s *TaobaoTbkDgGeneralLinkConvertMaterialPromotionInfoDTO) SetCommissionType(v string) *TaobaoTbkDgGeneralLinkConvertMaterialPromotionInfoDTO {
	s.CommissionType = &v
	return s
}
func (s *TaobaoTbkDgGeneralLinkConvertMaterialPromotionInfoDTO) SetMultipleItemsPricesCount(v int64) *TaobaoTbkDgGeneralLinkConvertMaterialPromotionInfoDTO {
	s.MultipleItemsPricesCount = &v
	return s
}
