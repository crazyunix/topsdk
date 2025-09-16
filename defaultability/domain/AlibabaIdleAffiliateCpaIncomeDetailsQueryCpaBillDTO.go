package domain

import "github.com/crazyunix/topsdk/util"

type AlibabaIdleAffiliateCpaIncomeDetailsQueryCpaBillDTO struct {
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
	   二级推广者id 服务商转链时传入，明细会透传     */
	SubPublisherId *string `json:"sub_publisher_id,omitempty" `

	/*
	   被邀请者id     */
	InviteeId *string `json:"invitee_id,omitempty" `

	/*
	   拓展信息，U1：新用户；U2：30天未登录；U3：当日未登录；trade_order_id：订单激励金所关联的订单号     */
	ExtendInfo *string `json:"extend_info,omitempty" `
}

func (s *AlibabaIdleAffiliateCpaIncomeDetailsQueryCpaBillDTO) SetAmount(v string) *AlibabaIdleAffiliateCpaIncomeDetailsQueryCpaBillDTO {
	s.Amount = &v
	return s
}
func (s *AlibabaIdleAffiliateCpaIncomeDetailsQueryCpaBillDTO) SetAmountCalculationFormula(v string) *AlibabaIdleAffiliateCpaIncomeDetailsQueryCpaBillDTO {
	s.AmountCalculationFormula = &v
	return s
}
func (s *AlibabaIdleAffiliateCpaIncomeDetailsQueryCpaBillDTO) SetStateText(v string) *AlibabaIdleAffiliateCpaIncomeDetailsQueryCpaBillDTO {
	s.StateText = &v
	return s
}
func (s *AlibabaIdleAffiliateCpaIncomeDetailsQueryCpaBillDTO) SetCreateTime(v util.LocalTime) *AlibabaIdleAffiliateCpaIncomeDetailsQueryCpaBillDTO {
	s.CreateTime = &v
	return s
}
func (s *AlibabaIdleAffiliateCpaIncomeDetailsQueryCpaBillDTO) SetCampaignId(v string) *AlibabaIdleAffiliateCpaIncomeDetailsQueryCpaBillDTO {
	s.CampaignId = &v
	return s
}
func (s *AlibabaIdleAffiliateCpaIncomeDetailsQueryCpaBillDTO) SetBillId(v string) *AlibabaIdleAffiliateCpaIncomeDetailsQueryCpaBillDTO {
	s.BillId = &v
	return s
}
func (s *AlibabaIdleAffiliateCpaIncomeDetailsQueryCpaBillDTO) SetUpdateTime(v util.LocalTime) *AlibabaIdleAffiliateCpaIncomeDetailsQueryCpaBillDTO {
	s.UpdateTime = &v
	return s
}
func (s *AlibabaIdleAffiliateCpaIncomeDetailsQueryCpaBillDTO) SetComment(v string) *AlibabaIdleAffiliateCpaIncomeDetailsQueryCpaBillDTO {
	s.Comment = &v
	return s
}
func (s *AlibabaIdleAffiliateCpaIncomeDetailsQueryCpaBillDTO) SetState(v int64) *AlibabaIdleAffiliateCpaIncomeDetailsQueryCpaBillDTO {
	s.State = &v
	return s
}
func (s *AlibabaIdleAffiliateCpaIncomeDetailsQueryCpaBillDTO) SetSubPublisherId(v string) *AlibabaIdleAffiliateCpaIncomeDetailsQueryCpaBillDTO {
	s.SubPublisherId = &v
	return s
}
func (s *AlibabaIdleAffiliateCpaIncomeDetailsQueryCpaBillDTO) SetInviteeId(v string) *AlibabaIdleAffiliateCpaIncomeDetailsQueryCpaBillDTO {
	s.InviteeId = &v
	return s
}
func (s *AlibabaIdleAffiliateCpaIncomeDetailsQueryCpaBillDTO) SetExtendInfo(v string) *AlibabaIdleAffiliateCpaIncomeDetailsQueryCpaBillDTO {
	s.ExtendInfo = &v
	return s
}
