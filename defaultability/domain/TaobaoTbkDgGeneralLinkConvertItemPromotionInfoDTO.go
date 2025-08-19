package domain

type TaobaoTbkDgGeneralLinkConvertItemPromotionInfoDTO struct {
	/*
	   商品收入比率(%)：商品佣金比率+补贴比率。同物料类接口income_rate     */
	CommissionRate *string `json:"commission_rate,omitempty" `

	/*
	   佣金类型。MKT表示营销计划，SP表示定向计划，COMMON表示通用计划，ZX表示自选计划     */
	CommissionType *string `json:"commission_type,omitempty" `

	/*
	   商品到手价。单位分     */
	PromotionPrice *int64 `json:"promotion_price,omitempty" `

	/*
	   多件价需拍件数     */
	MultipleItemsPricesCount *int64 `json:"multiple_items_prices_count,omitempty" `
}

func (s *TaobaoTbkDgGeneralLinkConvertItemPromotionInfoDTO) SetCommissionRate(v string) *TaobaoTbkDgGeneralLinkConvertItemPromotionInfoDTO {
	s.CommissionRate = &v
	return s
}
func (s *TaobaoTbkDgGeneralLinkConvertItemPromotionInfoDTO) SetCommissionType(v string) *TaobaoTbkDgGeneralLinkConvertItemPromotionInfoDTO {
	s.CommissionType = &v
	return s
}
func (s *TaobaoTbkDgGeneralLinkConvertItemPromotionInfoDTO) SetPromotionPrice(v int64) *TaobaoTbkDgGeneralLinkConvertItemPromotionInfoDTO {
	s.PromotionPrice = &v
	return s
}
func (s *TaobaoTbkDgGeneralLinkConvertItemPromotionInfoDTO) SetMultipleItemsPricesCount(v int64) *TaobaoTbkDgGeneralLinkConvertItemPromotionInfoDTO {
	s.MultipleItemsPricesCount = &v
	return s
}
