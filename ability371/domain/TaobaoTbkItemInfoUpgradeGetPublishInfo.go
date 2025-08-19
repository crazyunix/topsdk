package domain

type TaobaoTbkItemInfoUpgradeGetPublishInfo struct {
	/*
	   商品信息-收入比率(商品佣金比率+补贴比率)。15.5表示15.5%     */
	IncomeRate *string `json:"income_rate,omitempty" `

	/*
	   前N件佣金信息-前N件佣金生效或预热时透出以下字段     */
	TopnInfo *TaobaoTbkItemInfoUpgradeGetTopNInfoDTO `json:"topn_info,omitempty" `

	/*
	   单品淘礼金今日剩余可创建个数     */
	TljRemainNum *int64 `json:"tlj_remain_num,omitempty" `

	/*
	   两小时推广销量。 非实时，约半小时更新     */
	TwoHourPromotionSales *int64 `json:"two_hour_promotion_sales,omitempty" `

	/*
	   当天推广销量。 非实时，约1小时更新     */
	DailyPromotionSales *int64 `json:"daily_promotion_sales,omitempty" `

	/*
	   额外奖励活动类型，如果一个商品有多个奖励类型，返回结果使用空格分割，0=预售单单奖励，1=618超级U选单单补     */
	CpaRewardType *string `json:"cpa_reward_type,omitempty" `

	/*
	   额外奖励活动金额，活动奖励金额的类型与cpa_reward_type字段对应，如果一个商品有多个奖励类型，返回结果使用空格分割     */
	CpaRewardAmount *string `json:"cpa_reward_amount,omitempty" `

	/*
	   商品佣金信息     */
	IncomeInfo *TaobaoTbkItemInfoUpgradeGetFinalIncomeInfo `json:"income_info,omitempty" `
}

func (s *TaobaoTbkItemInfoUpgradeGetPublishInfo) SetIncomeRate(v string) *TaobaoTbkItemInfoUpgradeGetPublishInfo {
	s.IncomeRate = &v
	return s
}
func (s *TaobaoTbkItemInfoUpgradeGetPublishInfo) SetTopnInfo(v TaobaoTbkItemInfoUpgradeGetTopNInfoDTO) *TaobaoTbkItemInfoUpgradeGetPublishInfo {
	s.TopnInfo = &v
	return s
}
func (s *TaobaoTbkItemInfoUpgradeGetPublishInfo) SetTljRemainNum(v int64) *TaobaoTbkItemInfoUpgradeGetPublishInfo {
	s.TljRemainNum = &v
	return s
}
func (s *TaobaoTbkItemInfoUpgradeGetPublishInfo) SetTwoHourPromotionSales(v int64) *TaobaoTbkItemInfoUpgradeGetPublishInfo {
	s.TwoHourPromotionSales = &v
	return s
}
func (s *TaobaoTbkItemInfoUpgradeGetPublishInfo) SetDailyPromotionSales(v int64) *TaobaoTbkItemInfoUpgradeGetPublishInfo {
	s.DailyPromotionSales = &v
	return s
}
func (s *TaobaoTbkItemInfoUpgradeGetPublishInfo) SetCpaRewardType(v string) *TaobaoTbkItemInfoUpgradeGetPublishInfo {
	s.CpaRewardType = &v
	return s
}
func (s *TaobaoTbkItemInfoUpgradeGetPublishInfo) SetCpaRewardAmount(v string) *TaobaoTbkItemInfoUpgradeGetPublishInfo {
	s.CpaRewardAmount = &v
	return s
}
func (s *TaobaoTbkItemInfoUpgradeGetPublishInfo) SetIncomeInfo(v TaobaoTbkItemInfoUpgradeGetFinalIncomeInfo) *TaobaoTbkItemInfoUpgradeGetPublishInfo {
	s.IncomeInfo = &v
	return s
}
