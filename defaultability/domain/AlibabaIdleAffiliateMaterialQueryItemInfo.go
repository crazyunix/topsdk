package domain


type AlibabaIdleAffiliateMaterialQueryItemInfo struct {
    /*
        itemBaseInfo     */
    ItemBaseInfo  *AlibabaIdleAffiliateMaterialQueryItemBaseInfo `json:"item_base_info,omitempty" `

    /*
        itemPromoteInfo     */
    ItemPromoteInfo  *AlibabaIdleAffiliateMaterialQueryItemPromoteInfo `json:"item_promote_info,omitempty" `

}

func (s *AlibabaIdleAffiliateMaterialQueryItemInfo) SetItemBaseInfo(v AlibabaIdleAffiliateMaterialQueryItemBaseInfo) *AlibabaIdleAffiliateMaterialQueryItemInfo {
    s.ItemBaseInfo = &v
    return s
}
func (s *AlibabaIdleAffiliateMaterialQueryItemInfo) SetItemPromoteInfo(v AlibabaIdleAffiliateMaterialQueryItemPromoteInfo) *AlibabaIdleAffiliateMaterialQueryItemInfo {
    s.ItemPromoteInfo = &v
    return s
}
