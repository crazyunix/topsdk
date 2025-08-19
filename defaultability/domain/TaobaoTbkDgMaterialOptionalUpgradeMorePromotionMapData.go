package domain

type TaobaoTbkDgMaterialOptionalUpgradeMorePromotionMapData struct {
	/*
	   优惠名称，如“商品券”、“跨店满减”、“单品直降”等     */
	PromotionTitle *string `json:"promotion_title,omitempty" `

	/*
	   优惠利益点文案，如“1件7.92折”、“每200减20”等     */
	PromotionDesc *string `json:"promotion_desc,omitempty" `

	/*
	   优惠开始时间     */
	PromotionStartTime *string `json:"promotion_start_time,omitempty" `

	/*
	   优惠结束时间     */
	PromotionEndTime *string `json:"promotion_end_time,omitempty" `

	/*
	   优惠ID     */
	PromotionId *string `json:"promotion_id,omitempty" `

	/*
	   当天优惠总库存【指定优惠透出，不对外开放】     */
	PromotionTotalCount *int64 `json:"promotion_total_count,omitempty" `

	/*
	   优惠使用路径【指定优惠透出，不对外开放】     */
	PromotionUrl *string `json:"promotion_url,omitempty" `
}

func (s *TaobaoTbkDgMaterialOptionalUpgradeMorePromotionMapData) SetPromotionTitle(v string) *TaobaoTbkDgMaterialOptionalUpgradeMorePromotionMapData {
	s.PromotionTitle = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeMorePromotionMapData) SetPromotionDesc(v string) *TaobaoTbkDgMaterialOptionalUpgradeMorePromotionMapData {
	s.PromotionDesc = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeMorePromotionMapData) SetPromotionStartTime(v string) *TaobaoTbkDgMaterialOptionalUpgradeMorePromotionMapData {
	s.PromotionStartTime = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeMorePromotionMapData) SetPromotionEndTime(v string) *TaobaoTbkDgMaterialOptionalUpgradeMorePromotionMapData {
	s.PromotionEndTime = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeMorePromotionMapData) SetPromotionId(v string) *TaobaoTbkDgMaterialOptionalUpgradeMorePromotionMapData {
	s.PromotionId = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeMorePromotionMapData) SetPromotionTotalCount(v int64) *TaobaoTbkDgMaterialOptionalUpgradeMorePromotionMapData {
	s.PromotionTotalCount = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeMorePromotionMapData) SetPromotionUrl(v string) *TaobaoTbkDgMaterialOptionalUpgradeMorePromotionMapData {
	s.PromotionUrl = &v
	return s
}
