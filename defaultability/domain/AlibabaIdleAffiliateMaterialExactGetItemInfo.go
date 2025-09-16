package domain


type AlibabaIdleAffiliateMaterialExactGetItemInfo struct {
    /*
        商品基础信息     */
    ItemBaseInfo  *AlibabaIdleAffiliateMaterialExactGetItemBaseInfo `json:"item_base_info,omitempty" `

    /*
        佣金信息     */
    ItemPromoteInfo  *AlibabaIdleAffiliateMaterialExactGetItemPromoteInfo `json:"item_promote_info,omitempty" `

}

func (s *AlibabaIdleAffiliateMaterialExactGetItemInfo) SetItemBaseInfo(v AlibabaIdleAffiliateMaterialExactGetItemBaseInfo) *AlibabaIdleAffiliateMaterialExactGetItemInfo {
    s.ItemBaseInfo = &v
    return s
}
func (s *AlibabaIdleAffiliateMaterialExactGetItemInfo) SetItemPromoteInfo(v AlibabaIdleAffiliateMaterialExactGetItemPromoteInfo) *AlibabaIdleAffiliateMaterialExactGetItemInfo {
    s.ItemPromoteInfo = &v
    return s
}
