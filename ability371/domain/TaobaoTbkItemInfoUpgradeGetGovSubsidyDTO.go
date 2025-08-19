package domain

type TaobaoTbkItemInfoUpgradeGetGovSubsidyDTO struct {
	/*
	   国家补贴     */
	TagName *string `json:"tag_name,omitempty" `

	/*
	   国补信息     */
	StateSubsidyInfo *TaobaoTbkItemInfoUpgradeGetStateSubsidyInfo `json:"state_subsidy_info,omitempty" `
}

func (s *TaobaoTbkItemInfoUpgradeGetGovSubsidyDTO) SetTagName(v string) *TaobaoTbkItemInfoUpgradeGetGovSubsidyDTO {
	s.TagName = &v
	return s
}
func (s *TaobaoTbkItemInfoUpgradeGetGovSubsidyDTO) SetStateSubsidyInfo(v TaobaoTbkItemInfoUpgradeGetStateSubsidyInfo) *TaobaoTbkItemInfoUpgradeGetGovSubsidyDTO {
	s.StateSubsidyInfo = &v
	return s
}
