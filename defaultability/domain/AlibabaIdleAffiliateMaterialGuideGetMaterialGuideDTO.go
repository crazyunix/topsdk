package domain


type AlibabaIdleAffiliateMaterialGuideGetMaterialGuideDTO struct {
    /*
        商品导览     */
    ItemGuide  *AlibabaIdleAffiliateMaterialGuideGetMaterialGuide `json:"item_guide,omitempty" `

}

func (s *AlibabaIdleAffiliateMaterialGuideGetMaterialGuideDTO) SetItemGuide(v AlibabaIdleAffiliateMaterialGuideGetMaterialGuide) *AlibabaIdleAffiliateMaterialGuideGetMaterialGuideDTO {
    s.ItemGuide = &v
    return s
}
