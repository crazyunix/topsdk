package domain

import (
	"github.com/crazyunix/topsdk/util"
)

type AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryTradeOrderDTO struct {
	/*
	   商品图     */
	ItemPicUrl *string `json:"item_pic_url,omitempty" `

	/*
	   商品id     */
	ItemId *string `json:"item_id,omitempty" `

	/*
	   付款时间     */
	PayTime *util.LocalTime `json:"pay_time,omitempty" `

	/*
	   商品标题     */
	ItemTitle *string `json:"item_title,omitempty" `

	/*
	   订单id     */
	OrderId *string `json:"order_id,omitempty" `

	/*
	   订单状态描述     */
	StateDesc *string `json:"state_desc,omitempty" `

	/*
	   确收金额     */
	PartConfirmFee *string `json:"part_confirm_fee,omitempty" `

	/*
	   实付金额     */
	ActualPaidFee *string `json:"actual_paid_fee,omitempty" `

	/*
	   创建时间     */
	GmtCreate *util.LocalTime `json:"gmt_create,omitempty" `

	/*
	   订单状态     */
	StateCode *int64 `json:"state_code,omitempty" `

	/*
	   结束时间     */
	EndTime *util.LocalTime `json:"end_time,omitempty" `

	/*
	   红包优惠抵扣金额     */
	CouponDiscountFee *string `json:"coupon_discount_fee,omitempty" `
}

func (s *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryTradeOrderDTO) SetItemPicUrl(v string) *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryTradeOrderDTO {
	s.ItemPicUrl = &v
	return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryTradeOrderDTO) SetItemId(v string) *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryTradeOrderDTO {
	s.ItemId = &v
	return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryTradeOrderDTO) SetPayTime(v util.LocalTime) *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryTradeOrderDTO {
	s.PayTime = &v
	return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryTradeOrderDTO) SetItemTitle(v string) *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryTradeOrderDTO {
	s.ItemTitle = &v
	return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryTradeOrderDTO) SetOrderId(v string) *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryTradeOrderDTO {
	s.OrderId = &v
	return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryTradeOrderDTO) SetStateDesc(v string) *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryTradeOrderDTO {
	s.StateDesc = &v
	return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryTradeOrderDTO) SetPartConfirmFee(v string) *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryTradeOrderDTO {
	s.PartConfirmFee = &v
	return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryTradeOrderDTO) SetActualPaidFee(v string) *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryTradeOrderDTO {
	s.ActualPaidFee = &v
	return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryTradeOrderDTO) SetGmtCreate(v util.LocalTime) *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryTradeOrderDTO {
	s.GmtCreate = &v
	return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryTradeOrderDTO) SetStateCode(v int64) *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryTradeOrderDTO {
	s.StateCode = &v
	return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryTradeOrderDTO) SetEndTime(v util.LocalTime) *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryTradeOrderDTO {
	s.EndTime = &v
	return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryTradeOrderDTO) SetCouponDiscountFee(v string) *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryTradeOrderDTO {
	s.CouponDiscountFee = &v
	return s
}
