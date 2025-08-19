package domain

type TaobaoTbkDgMaterialOptionalUpgradeGovSubsidyDTO struct {
	/*
	   国家补贴     */
	TagName *string `json:"tag_name,omitempty" `

	/*
	   国补信息     */
	StateSubsidyInfo *TaobaoTbkDgMaterialOptionalUpgradeStateSubsidyInfo `json:"state_subsidy_info,omitempty" `
}

func (s *TaobaoTbkDgMaterialOptionalUpgradeGovSubsidyDTO) SetTagName(v string) *TaobaoTbkDgMaterialOptionalUpgradeGovSubsidyDTO {
	s.TagName = &v
	return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeGovSubsidyDTO) SetStateSubsidyInfo(v TaobaoTbkDgMaterialOptionalUpgradeStateSubsidyInfo) *TaobaoTbkDgMaterialOptionalUpgradeGovSubsidyDTO {
	s.StateSubsidyInfo = &v
	return s
}
