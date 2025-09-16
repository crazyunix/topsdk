package domain


type AlibabaIdleAffiliateMaterialExactGetCouponInfo struct {
    /*
        券id     */
    CouponId  *string `json:"coupon_id,omitempty" `

    /*
        券名称     */
    CouponName  *string `json:"coupon_name,omitempty" `

}

func (s *AlibabaIdleAffiliateMaterialExactGetCouponInfo) SetCouponId(v string) *AlibabaIdleAffiliateMaterialExactGetCouponInfo {
    s.CouponId = &v
    return s
}
func (s *AlibabaIdleAffiliateMaterialExactGetCouponInfo) SetCouponName(v string) *AlibabaIdleAffiliateMaterialExactGetCouponInfo {
    s.CouponName = &v
    return s
}
