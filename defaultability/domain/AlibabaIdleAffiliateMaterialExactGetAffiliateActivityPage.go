package domain

import (
	"github.com/crazyunix/topsdk/util"
)

type AlibabaIdleAffiliateMaterialExactGetAffiliateActivityPage struct {
	/*
	   营销活动id     */
	ActivityPageId *string `json:"activity_page_id,omitempty" `

	/*
	   营销活动名称     */
	Title *string `json:"title,omitempty" `

	/*
	   活动起始时间     */
	StartTime *util.LocalTime `json:"start_time,omitempty" `

	/*
	   活动结束时间     */
	EndTime *util.LocalTime `json:"end_time,omitempty" `

	/*
	   活动是否在APP端展示     */
	DisplayOnApp *bool `json:"display_on_app,omitempty" `

	/*
	   活动文案     */
	Comment *string `json:"comment,omitempty" `

	/*
	   活动图片URL     */
	ImgUrl *string `json:"img_url,omitempty" `

	/*
	   活动页面URL     */
	LandingUrl *string `json:"landing_url,omitempty" `

	/*
	   活动排序     */
	Priority *int64 `json:"priority,omitempty" `
}

func (s *AlibabaIdleAffiliateMaterialExactGetAffiliateActivityPage) SetActivityPageId(v string) *AlibabaIdleAffiliateMaterialExactGetAffiliateActivityPage {
	s.ActivityPageId = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialExactGetAffiliateActivityPage) SetTitle(v string) *AlibabaIdleAffiliateMaterialExactGetAffiliateActivityPage {
	s.Title = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialExactGetAffiliateActivityPage) SetStartTime(v util.LocalTime) *AlibabaIdleAffiliateMaterialExactGetAffiliateActivityPage {
	s.StartTime = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialExactGetAffiliateActivityPage) SetEndTime(v util.LocalTime) *AlibabaIdleAffiliateMaterialExactGetAffiliateActivityPage {
	s.EndTime = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialExactGetAffiliateActivityPage) SetDisplayOnApp(v bool) *AlibabaIdleAffiliateMaterialExactGetAffiliateActivityPage {
	s.DisplayOnApp = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialExactGetAffiliateActivityPage) SetComment(v string) *AlibabaIdleAffiliateMaterialExactGetAffiliateActivityPage {
	s.Comment = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialExactGetAffiliateActivityPage) SetImgUrl(v string) *AlibabaIdleAffiliateMaterialExactGetAffiliateActivityPage {
	s.ImgUrl = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialExactGetAffiliateActivityPage) SetLandingUrl(v string) *AlibabaIdleAffiliateMaterialExactGetAffiliateActivityPage {
	s.LandingUrl = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialExactGetAffiliateActivityPage) SetPriority(v int64) *AlibabaIdleAffiliateMaterialExactGetAffiliateActivityPage {
	s.Priority = &v
	return s
}
