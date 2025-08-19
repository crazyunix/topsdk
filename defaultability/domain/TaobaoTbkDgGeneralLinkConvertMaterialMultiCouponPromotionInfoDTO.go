package domain

type TaobaoTbkDgGeneralLinkConvertMaterialMultiCouponPromotionInfoDTO struct {
	/*
	   优惠名称     */
	CouponTitle *string `json:"coupon_title,omitempty" `

	/*
	   优惠券信息，满XX元减XX元，满x件减y元     */
	CouponDesc *string `json:"coupon_desc,omitempty" `

	/*
	   优惠金额（元）     */
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

func (s *TaobaoTbkDgGeneralLinkConvertMaterialMultiCouponPromotionInfoDTO) SetCouponTitle(v string) *TaobaoTbkDgGeneralLinkConvertMaterialMultiCouponPromotionInfoDTO {
	s.CouponTitle = &v
	return s
}
func (s *TaobaoTbkDgGeneralLinkConvertMaterialMultiCouponPromotionInfoDTO) SetCouponDesc(v string) *TaobaoTbkDgGeneralLinkConvertMaterialMultiCouponPromotionInfoDTO {
	s.CouponDesc = &v
	return s
}
func (s *TaobaoTbkDgGeneralLinkConvertMaterialMultiCouponPromotionInfoDTO) SetCoupoonFee(v string) *TaobaoTbkDgGeneralLinkConvertMaterialMultiCouponPromotionInfoDTO {
	s.CoupoonFee = &v
	return s
}
func (s *TaobaoTbkDgGeneralLinkConvertMaterialMultiCouponPromotionInfoDTO) SetCouponStartTime(v string) *TaobaoTbkDgGeneralLinkConvertMaterialMultiCouponPromotionInfoDTO {
	s.CouponStartTime = &v
	return s
}
func (s *TaobaoTbkDgGeneralLinkConvertMaterialMultiCouponPromotionInfoDTO) SetCouponEndTime(v string) *TaobaoTbkDgGeneralLinkConvertMaterialMultiCouponPromotionInfoDTO {
	s.CouponEndTime = &v
	return s
}
func (s *TaobaoTbkDgGeneralLinkConvertMaterialMultiCouponPromotionInfoDTO) SetActivityId(v string) *TaobaoTbkDgGeneralLinkConvertMaterialMultiCouponPromotionInfoDTO {
	s.ActivityId = &v
	return s
}
func (s *TaobaoTbkDgGeneralLinkConvertMaterialMultiCouponPromotionInfoDTO) SetCouponRemainCount(v int64) *TaobaoTbkDgGeneralLinkConvertMaterialMultiCouponPromotionInfoDTO {
	s.CouponRemainCount = &v
	return s
}
