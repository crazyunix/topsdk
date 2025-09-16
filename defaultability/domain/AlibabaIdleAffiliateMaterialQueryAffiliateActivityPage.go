package domain

import (
	"github.com/crazyunix/topsdk/util"
)

type AlibabaIdleAffiliateMaterialQueryAffiliateActivityPage struct {
	/*
	   活动页面主图     */
	ImgUrl *string `json:"img_url,omitempty" `

	/*
	   活动页面是否在App端展示     */
	DisplayOnApp *bool `json:"display_on_app,omitempty" `

	/*
	   活动开始时间     */
	StartTime *util.LocalTime `json:"start_time,omitempty" `

	/*
	   活动结束时间     */
	EndTime *util.LocalTime `json:"end_time,omitempty" `

	/*
	   活动页面详情文案     */
	Comment *string `json:"comment,omitempty" `

	/*
	   活动页面ID     */
	ActivityPageId *string `json:"activity_page_id,omitempty" `

	/*
	   "SERVING"活动有效 “CLOSED”活动无效     */
	State *string `json:"state,omitempty" `

	/*
	   活动页面标题     */
	Title *string `json:"title,omitempty" `

	/*
	   活动页面优先级     */
	Priority *int64 `json:"priority,omitempty" `

	/*
	   活动页面链接     */
	LandingUrl *string `json:"landing_url,omitempty" `
}

func (s *AlibabaIdleAffiliateMaterialQueryAffiliateActivityPage) SetImgUrl(v string) *AlibabaIdleAffiliateMaterialQueryAffiliateActivityPage {
	s.ImgUrl = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialQueryAffiliateActivityPage) SetDisplayOnApp(v bool) *AlibabaIdleAffiliateMaterialQueryAffiliateActivityPage {
	s.DisplayOnApp = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialQueryAffiliateActivityPage) SetStartTime(v util.LocalTime) *AlibabaIdleAffiliateMaterialQueryAffiliateActivityPage {
	s.StartTime = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialQueryAffiliateActivityPage) SetEndTime(v util.LocalTime) *AlibabaIdleAffiliateMaterialQueryAffiliateActivityPage {
	s.EndTime = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialQueryAffiliateActivityPage) SetComment(v string) *AlibabaIdleAffiliateMaterialQueryAffiliateActivityPage {
	s.Comment = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialQueryAffiliateActivityPage) SetActivityPageId(v string) *AlibabaIdleAffiliateMaterialQueryAffiliateActivityPage {
	s.ActivityPageId = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialQueryAffiliateActivityPage) SetState(v string) *AlibabaIdleAffiliateMaterialQueryAffiliateActivityPage {
	s.State = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialQueryAffiliateActivityPage) SetTitle(v string) *AlibabaIdleAffiliateMaterialQueryAffiliateActivityPage {
	s.Title = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialQueryAffiliateActivityPage) SetPriority(v int64) *AlibabaIdleAffiliateMaterialQueryAffiliateActivityPage {
	s.Priority = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialQueryAffiliateActivityPage) SetLandingUrl(v string) *AlibabaIdleAffiliateMaterialQueryAffiliateActivityPage {
	s.LandingUrl = &v
	return s
}
