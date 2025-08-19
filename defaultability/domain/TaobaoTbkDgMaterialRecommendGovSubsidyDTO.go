package domain

type TaobaoTbkDgMaterialRecommendGovSubsidyDTO struct {
	/*
	   国家补贴     */
	TagName *string `json:"tag_name,omitempty" `

	/*
	   国补信息     */
	StateSubsidyInfo *TaobaoTbkDgMaterialRecommendStateSubsidyInfo `json:"state_subsidy_info,omitempty" `
}

func (s *TaobaoTbkDgMaterialRecommendGovSubsidyDTO) SetTagName(v string) *TaobaoTbkDgMaterialRecommendGovSubsidyDTO {
	s.TagName = &v
	return s
}
func (s *TaobaoTbkDgMaterialRecommendGovSubsidyDTO) SetStateSubsidyInfo(v TaobaoTbkDgMaterialRecommendStateSubsidyInfo) *TaobaoTbkDgMaterialRecommendGovSubsidyDTO {
	s.StateSubsidyInfo = &v
	return s
}
