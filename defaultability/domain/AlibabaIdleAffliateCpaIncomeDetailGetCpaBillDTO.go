package domain

import (
	"github.com/crazyunix/topsdk/util"
)

type AlibabaIdleAffliateCpaIncomeDetailGetCpaBillDTO struct {
	/*
	   佣金     */
	Amount *string `json:"amount,omitempty" `

	/*
	   佣金计算公式     */
	AmountCalculationFormula *string `json:"amount_calculation_formula,omitempty" `

	/*
	   stateText     */
	StateText *string `json:"state_text,omitempty" `

	/*
	   createTime     */
	CreateTime *util.LocalTime `json:"create_time,omitempty" `

	/*
	   活动id     */
	CampaignId *string `json:"campaign_id,omitempty" `

	/*
	   佣金单id(加密的)     */
	BillId *string `json:"bill_id,omitempty" `

	/*
	   updateTime     */
	UpdateTime *util.LocalTime `json:"update_time,omitempty" `

	/*
	   佣金备注     */
	Comment *string `json:"comment,omitempty" `

	/*
	   0待发放 1待到账 2已到账 3已失效     */
	State *int64 `json:"state,omitempty" `

	/*
	   二级推广者id     */
	SubPublisherId *string `json:"sub_publisher_id,omitempty" `

	/*
	   被邀请者id     */
	InviteeId *string `json:"invitee_id,omitempty" `

	/*
	   拓展信息 可获取用户标签等信息     */
	ExtendInfo *string `json:"extend_info,omitempty" `
}

func (s *AlibabaIdleAffliateCpaIncomeDetailGetCpaBillDTO) SetAmount(v string) *AlibabaIdleAffliateCpaIncomeDetailGetCpaBillDTO {
	s.Amount = &v
	return s
}
func (s *AlibabaIdleAffliateCpaIncomeDetailGetCpaBillDTO) SetAmountCalculationFormula(v string) *AlibabaIdleAffliateCpaIncomeDetailGetCpaBillDTO {
	s.AmountCalculationFormula = &v
	return s
}
func (s *AlibabaIdleAffliateCpaIncomeDetailGetCpaBillDTO) SetStateText(v string) *AlibabaIdleAffliateCpaIncomeDetailGetCpaBillDTO {
	s.StateText = &v
	return s
}
func (s *AlibabaIdleAffliateCpaIncomeDetailGetCpaBillDTO) SetCreateTime(v util.LocalTime) *AlibabaIdleAffliateCpaIncomeDetailGetCpaBillDTO {
	s.CreateTime = &v
	return s
}
func (s *AlibabaIdleAffliateCpaIncomeDetailGetCpaBillDTO) SetCampaignId(v string) *AlibabaIdleAffliateCpaIncomeDetailGetCpaBillDTO {
	s.CampaignId = &v
	return s
}
func (s *AlibabaIdleAffliateCpaIncomeDetailGetCpaBillDTO) SetBillId(v string) *AlibabaIdleAffliateCpaIncomeDetailGetCpaBillDTO {
	s.BillId = &v
	return s
}
func (s *AlibabaIdleAffliateCpaIncomeDetailGetCpaBillDTO) SetUpdateTime(v util.LocalTime) *AlibabaIdleAffliateCpaIncomeDetailGetCpaBillDTO {
	s.UpdateTime = &v
	return s
}
func (s *AlibabaIdleAffliateCpaIncomeDetailGetCpaBillDTO) SetComment(v string) *AlibabaIdleAffliateCpaIncomeDetailGetCpaBillDTO {
	s.Comment = &v
	return s
}
func (s *AlibabaIdleAffliateCpaIncomeDetailGetCpaBillDTO) SetState(v int64) *AlibabaIdleAffliateCpaIncomeDetailGetCpaBillDTO {
	s.State = &v
	return s
}
func (s *AlibabaIdleAffliateCpaIncomeDetailGetCpaBillDTO) SetSubPublisherId(v string) *AlibabaIdleAffliateCpaIncomeDetailGetCpaBillDTO {
	s.SubPublisherId = &v
	return s
}
func (s *AlibabaIdleAffliateCpaIncomeDetailGetCpaBillDTO) SetInviteeId(v string) *AlibabaIdleAffliateCpaIncomeDetailGetCpaBillDTO {
	s.InviteeId = &v
	return s
}
func (s *AlibabaIdleAffliateCpaIncomeDetailGetCpaBillDTO) SetExtendInfo(v string) *AlibabaIdleAffliateCpaIncomeDetailGetCpaBillDTO {
	s.ExtendInfo = &v
	return s
}
