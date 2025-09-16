package domain


type AlibabaIdleAffiliateMaterialGuideGetMaterialGuide struct {
    /*
        商品tab名list     */
    TabNameList  *[]string `json:"tab_name_list,omitempty" `

}

func (s *AlibabaIdleAffiliateMaterialGuideGetMaterialGuide) SetTabNameList(v []string) *AlibabaIdleAffiliateMaterialGuideGetMaterialGuide {
    s.TabNameList = &v
    return s
}
