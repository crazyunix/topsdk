package domain


type AlibabaIdleAffiliateMaterialQueryMaterialDTO struct {
    /*
        物料类型 1商品 2权益 5活动页面     */
    MaterialType  *int64 `json:"material_type,omitempty" `

    /*
        商品     */
    ItemInfo  *AlibabaIdleAffiliateMaterialQueryItemInfo `json:"item_info,omitempty" `

    /*
        券信息     */
    CouponInfo  *AlibabaIdleAffiliateMaterialQueryCouponInfo `json:"coupon_info,omitempty" `

    /*
        营销活动页面信息     */
    ActivityPageInfo  *AlibabaIdleAffiliateMaterialQueryAffiliateActivityPage `json:"activity_page_info,omitempty" `

}

func (s *AlibabaIdleAffiliateMaterialQueryMaterialDTO) SetMaterialType(v int64) *AlibabaIdleAffiliateMaterialQueryMaterialDTO {
    s.MaterialType = &v
    return s
}
func (s *AlibabaIdleAffiliateMaterialQueryMaterialDTO) SetItemInfo(v AlibabaIdleAffiliateMaterialQueryItemInfo) *AlibabaIdleAffiliateMaterialQueryMaterialDTO {
    s.ItemInfo = &v
    return s
}
func (s *AlibabaIdleAffiliateMaterialQueryMaterialDTO) SetCouponInfo(v AlibabaIdleAffiliateMaterialQueryCouponInfo) *AlibabaIdleAffiliateMaterialQueryMaterialDTO {
    s.CouponInfo = &v
    return s
}
func (s *AlibabaIdleAffiliateMaterialQueryMaterialDTO) SetActivityPageInfo(v AlibabaIdleAffiliateMaterialQueryAffiliateActivityPage) *AlibabaIdleAffiliateMaterialQueryMaterialDTO {
    s.ActivityPageInfo = &v
    return s
}
