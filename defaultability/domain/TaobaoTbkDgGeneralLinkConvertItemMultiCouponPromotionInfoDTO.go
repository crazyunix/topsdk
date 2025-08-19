package domain

type TaobaoTbkDgGeneralLinkConvertItemMultiCouponPromotionInfoDTO struct {
	/*
	   优惠名称     */
	CouponTitle *string `json:"coupon_title,omitempty" `

	/*
	   优惠券信息，满XX元减XX元，满x件减y元     */
	CouponDesc *string `json:"coupon_desc,omitempty" `

	/*
	   优惠金额。单位元     */
	CoupoonFee *string `json:"coupoon_fee,omitempty" `

	/*
	   优惠开始时间     */
	CouponStartTime *string `json:"coupon_start_time,omitempty" `

	/*
	   优惠结束时间     */
	CouponEndTime *string `json:"coupon_end_time,omitempty" `

	/*
	   优惠ID     */
	ActivityId *string `json:"activity_id,omitempty" `

	/*
	   优惠剩余数量     */
	CouponRemainCount *int64 `json:"coupon_remain_count,omitempty" `
}

func (s *TaobaoTbkDgGeneralLinkConvertItemMultiCouponPromotionInfoDTO) SetCouponTitle(v string) *TaobaoTbkDgGeneralLinkConvertItemMultiCouponPromotionInfoDTO {
	s.CouponTitle = &v
	return s
}
func (s *TaobaoTbkDgGeneralLinkConvertItemMultiCouponPromotionInfoDTO) SetCouponDesc(v string) *TaobaoTbkDgGeneralLinkConvertItemMultiCouponPromotionInfoDTO {
	s.CouponDesc = &v
	return s
}
func (s *TaobaoTbkDgGeneralLinkConvertItemMultiCouponPromotionInfoDTO) SetCoupoonFee(v string) *TaobaoTbkDgGeneralLinkConvertItemMultiCouponPromotionInfoDTO {
	s.CoupoonFee = &v
	return s
}
func (s *TaobaoTbkDgGeneralLinkConvertItemMultiCouponPromotionInfoDTO) SetCouponStartTime(v string) *TaobaoTbkDgGeneralLinkConvertItemMultiCouponPromotionInfoDTO {
	s.CouponStartTime = &v
	return s
}
func (s *TaobaoTbkDgGeneralLinkConvertItemMultiCouponPromotionInfoDTO) SetCouponEndTime(v string) *TaobaoTbkDgGeneralLinkConvertItemMultiCouponPromotionInfoDTO {
	s.CouponEndTime = &v
	return s
}
func (s *TaobaoTbkDgGeneralLinkConvertItemMultiCouponPromotionInfoDTO) SetActivityId(v string) *TaobaoTbkDgGeneralLinkConvertItemMultiCouponPromotionInfoDTO {
	s.ActivityId = &v
	return s
}
func (s *TaobaoTbkDgGeneralLinkConvertItemMultiCouponPromotionInfoDTO) SetCouponRemainCount(v int64) *TaobaoTbkDgGeneralLinkConvertItemMultiCouponPromotionInfoDTO {
	s.CouponRemainCount = &v
	return s
}
