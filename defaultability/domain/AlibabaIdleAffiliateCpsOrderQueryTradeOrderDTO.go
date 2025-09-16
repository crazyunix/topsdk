package domain

import (
	"github.com/crazyunix/topsdk/util"
)

type AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO struct {
	/*
	   退款金额     */
	RefundFee *string `json:"refund_fee,omitempty" `

	/*
	   商品图     */
	ItemPicUrl *string `json:"item_pic_url,omitempty" `

	/*
	   订单id     */
	OrderId *string `json:"order_id,omitempty" `

	/*
	   商品标题     */
	ItemTitle *string `json:"item_title,omitempty" `

	/*
	   付款时间     */
	PayTime *util.LocalTime `json:"pay_time,omitempty" `

	/*
	   订单状态描述     */
	StateDesc *string `json:"state_desc,omitempty" `

	/*
	   确收金额     */
	PartConfirmFee *string `json:"part_confirm_fee,omitempty" `

	/*
	   实际付款金额     */
	ActualPaidFee *string `json:"actual_paid_fee,omitempty" `

	/*
	   创建时间     */
	GmtCreate *util.LocalTime `json:"gmt_create,omitempty" `

	/*
	   商品id     */
	ItemId *string `json:"item_id,omitempty" `

	/*
	   预估佣金     */
	EstimateCommission *string `json:"estimate_commission,omitempty" `

	/*
	   预估费率     */
	EstimateCommissionRate *string `json:"estimate_commission_rate,omitempty" `

	/*
	   订单状态码     */
	StateCode *int64 `json:"state_code,omitempty" `

	/*
	   商品价格     */
	ItemPrice *string `json:"item_price,omitempty" `

	/*
	   订单结束时间     */
	EndTime *util.LocalTime `json:"end_time,omitempty" `

	/*
	   itemImageUrl     */
	ItemImageUrl *string `json:"item_image_url,omitempty" `

	/*
	   子推广者id     */
	SubPublisherId *string `json:"sub_publisher_id,omitempty" `

	/*
	   红包优惠抵扣金额     */
	CouponDiscountFee *string `json:"coupon_discount_fee,omitempty" `
}

func (s *AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO) SetRefundFee(v string) *AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO {
	s.RefundFee = &v
	return s
}
func (s *AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO) SetItemPicUrl(v string) *AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO {
	s.ItemPicUrl = &v
	return s
}
func (s *AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO) SetOrderId(v string) *AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO {
	s.OrderId = &v
	return s
}
func (s *AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO) SetItemTitle(v string) *AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO {
	s.ItemTitle = &v
	return s
}
func (s *AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO) SetPayTime(v util.LocalTime) *AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO {
	s.PayTime = &v
	return s
}
func (s *AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO) SetStateDesc(v string) *AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO {
	s.StateDesc = &v
	return s
}
func (s *AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO) SetPartConfirmFee(v string) *AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO {
	s.PartConfirmFee = &v
	return s
}
func (s *AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO) SetActualPaidFee(v string) *AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO {
	s.ActualPaidFee = &v
	return s
}
func (s *AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO) SetGmtCreate(v util.LocalTime) *AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO {
	s.GmtCreate = &v
	return s
}
func (s *AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO) SetItemId(v string) *AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO {
	s.ItemId = &v
	return s
}
func (s *AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO) SetEstimateCommission(v string) *AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO {
	s.EstimateCommission = &v
	return s
}
func (s *AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO) SetEstimateCommissionRate(v string) *AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO {
	s.EstimateCommissionRate = &v
	return s
}
func (s *AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO) SetStateCode(v int64) *AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO {
	s.StateCode = &v
	return s
}
func (s *AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO) SetItemPrice(v string) *AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO {
	s.ItemPrice = &v
	return s
}
func (s *AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO) SetEndTime(v util.LocalTime) *AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO {
	s.EndTime = &v
	return s
}
func (s *AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO) SetItemImageUrl(v string) *AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO {
	s.ItemImageUrl = &v
	return s
}
func (s *AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO) SetSubPublisherId(v string) *AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO {
	s.SubPublisherId = &v
	return s
}
func (s *AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO) SetCouponDiscountFee(v string) *AlibabaIdleAffiliateCpsOrderQueryTradeOrderDTO {
	s.CouponDiscountFee = &v
	return s
}
