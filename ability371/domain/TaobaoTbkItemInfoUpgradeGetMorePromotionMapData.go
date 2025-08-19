package domain

type TaobaoTbkItemInfoUpgradeGetMorePromotionMapData struct {
	/*
	   预热优惠名称，如“商品券”、“跨店满减”、“单品直降”等     */
	PromotionTitle *string `json:"promotion_title,omitempty" `

	/*
	   预热优惠利益点文案，如“1件7.92折”、“每200减20”等     */
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

func (s *TaobaoTbkItemInfoUpgradeGetMorePromotionMapData) SetPromotionTitle(v string) *TaobaoTbkItemInfoUpgradeGetMorePromotionMapData {
	s.PromotionTitle = &v
	return s
}
func (s *TaobaoTbkItemInfoUpgradeGetMorePromotionMapData) SetPromotionDesc(v string) *TaobaoTbkItemInfoUpgradeGetMorePromotionMapData {
	s.PromotionDesc = &v
	return s
}
func (s *TaobaoTbkItemInfoUpgradeGetMorePromotionMapData) SetPromotionStartTime(v string) *TaobaoTbkItemInfoUpgradeGetMorePromotionMapData {
	s.PromotionStartTime = &v
	return s
}
func (s *TaobaoTbkItemInfoUpgradeGetMorePromotionMapData) SetPromotionEndTime(v string) *TaobaoTbkItemInfoUpgradeGetMorePromotionMapData {
	s.PromotionEndTime = &v
	return s
}
func (s *TaobaoTbkItemInfoUpgradeGetMorePromotionMapData) SetPromotionId(v string) *TaobaoTbkItemInfoUpgradeGetMorePromotionMapData {
	s.PromotionId = &v
	return s
}
func (s *TaobaoTbkItemInfoUpgradeGetMorePromotionMapData) SetPromotionTotalCount(v int64) *TaobaoTbkItemInfoUpgradeGetMorePromotionMapData {
	s.PromotionTotalCount = &v
	return s
}
func (s *TaobaoTbkItemInfoUpgradeGetMorePromotionMapData) SetPromotionUrl(v string) *TaobaoTbkItemInfoUpgradeGetMorePromotionMapData {
	s.PromotionUrl = &v
	return s
}
