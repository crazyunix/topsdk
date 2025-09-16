package domain


type AlibabaIdleAffiliateMaterialExactGetMaterialsExactQueryVO struct {
    /*
        1 商品 2 权益 5 营销活动 defalutValue:1    */
    MaterialType  *int64 `json:"material_type,omitempty" `

    /*
        商品id     */
    ItemIds  *[]string `json:"item_ids,omitempty" `

    /*
        商品明文id     */
    PlainItemIds  *[]string `json:"plain_item_ids,omitempty" `

    /*
        营销活动id     */
    ActivityPageIds  *[]string `json:"activity_page_ids,omitempty" `

}

func (s *AlibabaIdleAffiliateMaterialExactGetMaterialsExactQueryVO) SetMaterialType(v int64) *AlibabaIdleAffiliateMaterialExactGetMaterialsExactQueryVO {
    s.MaterialType = &v
    return s
}
func (s *AlibabaIdleAffiliateMaterialExactGetMaterialsExactQueryVO) SetItemIds(v []string) *AlibabaIdleAffiliateMaterialExactGetMaterialsExactQueryVO {
    s.ItemIds = &v
    return s
}
func (s *AlibabaIdleAffiliateMaterialExactGetMaterialsExactQueryVO) SetPlainItemIds(v []string) *AlibabaIdleAffiliateMaterialExactGetMaterialsExactQueryVO {
    s.PlainItemIds = &v
    return s
}
func (s *AlibabaIdleAffiliateMaterialExactGetMaterialsExactQueryVO) SetActivityPageIds(v []string) *AlibabaIdleAffiliateMaterialExactGetMaterialsExactQueryVO {
    s.ActivityPageIds = &v
    return s
}
