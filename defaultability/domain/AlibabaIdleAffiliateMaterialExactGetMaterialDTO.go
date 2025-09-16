package domain


type AlibabaIdleAffiliateMaterialExactGetMaterialDTO struct {
    /*
        物料类型 1商品 2权益 5营销活动     */
    MaterialType  *int64 `json:"material_type,omitempty" `

    /*
        商品信息     */
    ItemInfo  *AlibabaIdleAffiliateMaterialExactGetItemInfo `json:"item_info,omitempty" `

    /*
        券信息     */
    CouponInfo  *AlibabaIdleAffiliateMaterialExactGetCouponInfo `json:"coupon_info,omitempty" `

    /*
        营销活动信息     */
    ActivityPageInfo  *AlibabaIdleAffiliateMaterialExactGetAffiliateActivityPage `json:"activity_page_info,omitempty" `

}

func (s *AlibabaIdleAffiliateMaterialExactGetMaterialDTO) SetMaterialType(v int64) *AlibabaIdleAffiliateMaterialExactGetMaterialDTO {
    s.MaterialType = &v
    return s
}
func (s *AlibabaIdleAffiliateMaterialExactGetMaterialDTO) SetItemInfo(v AlibabaIdleAffiliateMaterialExactGetItemInfo) *AlibabaIdleAffiliateMaterialExactGetMaterialDTO {
    s.ItemInfo = &v
    return s
}
func (s *AlibabaIdleAffiliateMaterialExactGetMaterialDTO) SetCouponInfo(v AlibabaIdleAffiliateMaterialExactGetCouponInfo) *AlibabaIdleAffiliateMaterialExactGetMaterialDTO {
    s.CouponInfo = &v
    return s
}
func (s *AlibabaIdleAffiliateMaterialExactGetMaterialDTO) SetActivityPageInfo(v AlibabaIdleAffiliateMaterialExactGetAffiliateActivityPage) *AlibabaIdleAffiliateMaterialExactGetMaterialDTO {
    s.ActivityPageInfo = &v
    return s
}
