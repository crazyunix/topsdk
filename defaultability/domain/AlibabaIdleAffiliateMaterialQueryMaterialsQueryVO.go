package domain


type AlibabaIdleAffiliateMaterialQueryMaterialsQueryVO struct {
    /*
        1商品 2权益 5活动页面 defalutValue:1    */
    MaterialType  *int64 `json:"material_type,omitempty" `

    /*
        分页参数     */
    PageRequest  *AlibabaIdleAffiliateMaterialQueryPageRequest `json:"page_request,omitempty" `

    /*
        商品导览入参     */
    ItemGuideVO  *AlibabaIdleAffiliateMaterialQueryItemGuideVO `json:"item_guide_v_o,omitempty" `

}

func (s *AlibabaIdleAffiliateMaterialQueryMaterialsQueryVO) SetMaterialType(v int64) *AlibabaIdleAffiliateMaterialQueryMaterialsQueryVO {
    s.MaterialType = &v
    return s
}
func (s *AlibabaIdleAffiliateMaterialQueryMaterialsQueryVO) SetPageRequest(v AlibabaIdleAffiliateMaterialQueryPageRequest) *AlibabaIdleAffiliateMaterialQueryMaterialsQueryVO {
    s.PageRequest = &v
    return s
}
func (s *AlibabaIdleAffiliateMaterialQueryMaterialsQueryVO) SetItemGuideVO(v AlibabaIdleAffiliateMaterialQueryItemGuideVO) *AlibabaIdleAffiliateMaterialQueryMaterialsQueryVO {
    s.ItemGuideVO = &v
    return s
}
