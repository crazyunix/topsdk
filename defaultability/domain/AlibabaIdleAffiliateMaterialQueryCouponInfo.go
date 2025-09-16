package domain


type AlibabaIdleAffiliateMaterialQueryCouponInfo struct {
    /*
        券id     */
    CouponId  *string `json:"coupon_id,omitempty" `

    /*
        券名称     */
    CouponName  *string `json:"coupon_name,omitempty" `

}

func (s *AlibabaIdleAffiliateMaterialQueryCouponInfo) SetCouponId(v string) *AlibabaIdleAffiliateMaterialQueryCouponInfo {
    s.CouponId = &v
    return s
}
func (s *AlibabaIdleAffiliateMaterialQueryCouponInfo) SetCouponName(v string) *AlibabaIdleAffiliateMaterialQueryCouponInfo {
    s.CouponName = &v
    return s
}
