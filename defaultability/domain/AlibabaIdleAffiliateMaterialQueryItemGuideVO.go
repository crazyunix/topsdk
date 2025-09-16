package domain


type AlibabaIdleAffiliateMaterialQueryItemGuideVO struct {
    /*
        【推广赚钱】页面tab名     */
    TabName  *string `json:"tab_name,omitempty" `

    /*
        查询范围 in1day in3days in7days in14days     */
    ItemPublisherTime  *string `json:"item_publisher_time,omitempty" `

    /*
        卖家信用 excellent 极好 good 优秀     */
    SellerCreditLevel  *string `json:"seller_credit_level,omitempty" `

    /*
        筛选鱼小铺等级5以上     */
    FilterLevel5Yxp  *bool `json:"filter_level_5_yxp,omitempty" `

    /*
        筛选验货宝商品     */
    FilterYhb  *bool `json:"filter_yhb,omitempty" `

    /*
        卖家会员名     */
    SellerName  *string `json:"seller_name,omitempty" `

    /*
        搜索关键词     */
    Keyword  *string `json:"keyword,omitempty" `

}

func (s *AlibabaIdleAffiliateMaterialQueryItemGuideVO) SetTabName(v string) *AlibabaIdleAffiliateMaterialQueryItemGuideVO {
    s.TabName = &v
    return s
}
func (s *AlibabaIdleAffiliateMaterialQueryItemGuideVO) SetItemPublisherTime(v string) *AlibabaIdleAffiliateMaterialQueryItemGuideVO {
    s.ItemPublisherTime = &v
    return s
}
func (s *AlibabaIdleAffiliateMaterialQueryItemGuideVO) SetSellerCreditLevel(v string) *AlibabaIdleAffiliateMaterialQueryItemGuideVO {
    s.SellerCreditLevel = &v
    return s
}
func (s *AlibabaIdleAffiliateMaterialQueryItemGuideVO) SetFilterLevel5Yxp(v bool) *AlibabaIdleAffiliateMaterialQueryItemGuideVO {
    s.FilterLevel5Yxp = &v
    return s
}
func (s *AlibabaIdleAffiliateMaterialQueryItemGuideVO) SetFilterYhb(v bool) *AlibabaIdleAffiliateMaterialQueryItemGuideVO {
    s.FilterYhb = &v
    return s
}
func (s *AlibabaIdleAffiliateMaterialQueryItemGuideVO) SetSellerName(v string) *AlibabaIdleAffiliateMaterialQueryItemGuideVO {
    s.SellerName = &v
    return s
}
func (s *AlibabaIdleAffiliateMaterialQueryItemGuideVO) SetKeyword(v string) *AlibabaIdleAffiliateMaterialQueryItemGuideVO {
    s.Keyword = &v
    return s
}
