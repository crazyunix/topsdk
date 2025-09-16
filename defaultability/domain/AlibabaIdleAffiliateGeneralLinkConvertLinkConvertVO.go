package domain


type AlibabaIdleAffiliateGeneralLinkConvertLinkConvertVO struct {
    /*
        子推广者id     */
    SubPublisherId  *string `json:"sub_publisher_id,omitempty" `

    /*
        物料类型(1商品，2权益，3cpa活动，4闲鱼口令或短链，5营销活动页面)     */
    MaterialType  *int64 `json:"material_type,omitempty" `

    /*
        加密的商品id集合(可以从物料接口获取，测试时格式为array)     */
    ItemIds  *[]string `json:"item_ids,omitempty" `

    /*
        券id     */
    CouponId  *string `json:"coupon_id,omitempty" `

    /*
        cpa活动id，从联盟活动接口获取     */
    CampaignId  *string `json:"campaign_id,omitempty" `

    /*
        明文商品id集合(测试时格式为array)     */
    PlainItemIds  *[]string `json:"plain_item_ids,omitempty" `

    /*
        1.通过接口转链的短口令 2.从闲鱼帮卖广场分享出来的商品口令     */
    XyUrl  *[]int64 `json:"xy_url,omitempty" `

    /*
        营销活动页面id     */
    ActivityPageId  *string `json:"activity_page_id,omitempty" `

}

func (s *AlibabaIdleAffiliateGeneralLinkConvertLinkConvertVO) SetSubPublisherId(v string) *AlibabaIdleAffiliateGeneralLinkConvertLinkConvertVO {
    s.SubPublisherId = &v
    return s
}
func (s *AlibabaIdleAffiliateGeneralLinkConvertLinkConvertVO) SetMaterialType(v int64) *AlibabaIdleAffiliateGeneralLinkConvertLinkConvertVO {
    s.MaterialType = &v
    return s
}
func (s *AlibabaIdleAffiliateGeneralLinkConvertLinkConvertVO) SetItemIds(v []string) *AlibabaIdleAffiliateGeneralLinkConvertLinkConvertVO {
    s.ItemIds = &v
    return s
}
func (s *AlibabaIdleAffiliateGeneralLinkConvertLinkConvertVO) SetCouponId(v string) *AlibabaIdleAffiliateGeneralLinkConvertLinkConvertVO {
    s.CouponId = &v
    return s
}
func (s *AlibabaIdleAffiliateGeneralLinkConvertLinkConvertVO) SetCampaignId(v string) *AlibabaIdleAffiliateGeneralLinkConvertLinkConvertVO {
    s.CampaignId = &v
    return s
}
func (s *AlibabaIdleAffiliateGeneralLinkConvertLinkConvertVO) SetPlainItemIds(v []string) *AlibabaIdleAffiliateGeneralLinkConvertLinkConvertVO {
    s.PlainItemIds = &v
    return s
}
func (s *AlibabaIdleAffiliateGeneralLinkConvertLinkConvertVO) SetXyUrl(v []int64) *AlibabaIdleAffiliateGeneralLinkConvertLinkConvertVO {
    s.XyUrl = &v
    return s
}
func (s *AlibabaIdleAffiliateGeneralLinkConvertLinkConvertVO) SetActivityPageId(v string) *AlibabaIdleAffiliateGeneralLinkConvertLinkConvertVO {
    s.ActivityPageId = &v
    return s
}
