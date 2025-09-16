package domain

import (
        "github.com/crazyunix/topsdk/util"
    )

type AlibabaIdleAffiliateCpsIncomeDetailsQueryTradeOrderDTO struct {
    /*
        商品图     */
    ItemPicUrl  *string `json:"item_pic_url,omitempty" `

    /*
        商品id     */
    ItemId  *string `json:"item_id,omitempty" `

    /*
        付款时间     */
    PayTime  *util.LocalTime `json:"pay_time,omitempty" `

    /*
        商品标题     */
    ItemTitle  *string `json:"item_title,omitempty" `

    /*
        订单id     */
    OrderId  *string `json:"order_id,omitempty" `

    /*
        订单状态描述     */
    StateDesc  *string `json:"state_desc,omitempty" `

    /*
        确收金额     */
    PartConfirmFee  *string `json:"part_confirm_fee,omitempty" `

    /*
        实付金额     */
    ActualPaidFee  *string `json:"actual_paid_fee,omitempty" `

    /*
        创建时间     */
    GmtCreate  *util.LocalTime `json:"gmt_create,omitempty" `

    /*
        订单状态     */
    StateCode  *int64 `json:"state_code,omitempty" `

    /*
        订单结束时间     */
    EndTime  *util.LocalTime `json:"end_time,omitempty" `

    /*
        红包优惠抵扣金额     */
    CouponDiscountFee  *string `json:"coupon_discount_fee,omitempty" `

}

func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryTradeOrderDTO) SetItemPicUrl(v string) *AlibabaIdleAffiliateCpsIncomeDetailsQueryTradeOrderDTO {
    s.ItemPicUrl = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryTradeOrderDTO) SetItemId(v string) *AlibabaIdleAffiliateCpsIncomeDetailsQueryTradeOrderDTO {
    s.ItemId = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryTradeOrderDTO) SetPayTime(v util.LocalTime) *AlibabaIdleAffiliateCpsIncomeDetailsQueryTradeOrderDTO {
    s.PayTime = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryTradeOrderDTO) SetItemTitle(v string) *AlibabaIdleAffiliateCpsIncomeDetailsQueryTradeOrderDTO {
    s.ItemTitle = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryTradeOrderDTO) SetOrderId(v string) *AlibabaIdleAffiliateCpsIncomeDetailsQueryTradeOrderDTO {
    s.OrderId = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryTradeOrderDTO) SetStateDesc(v string) *AlibabaIdleAffiliateCpsIncomeDetailsQueryTradeOrderDTO {
    s.StateDesc = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryTradeOrderDTO) SetPartConfirmFee(v string) *AlibabaIdleAffiliateCpsIncomeDetailsQueryTradeOrderDTO {
    s.PartConfirmFee = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryTradeOrderDTO) SetActualPaidFee(v string) *AlibabaIdleAffiliateCpsIncomeDetailsQueryTradeOrderDTO {
    s.ActualPaidFee = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryTradeOrderDTO) SetGmtCreate(v util.LocalTime) *AlibabaIdleAffiliateCpsIncomeDetailsQueryTradeOrderDTO {
    s.GmtCreate = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryTradeOrderDTO) SetStateCode(v int64) *AlibabaIdleAffiliateCpsIncomeDetailsQueryTradeOrderDTO {
    s.StateCode = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryTradeOrderDTO) SetEndTime(v util.LocalTime) *AlibabaIdleAffiliateCpsIncomeDetailsQueryTradeOrderDTO {
    s.EndTime = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryTradeOrderDTO) SetCouponDiscountFee(v string) *AlibabaIdleAffiliateCpsIncomeDetailsQueryTradeOrderDTO {
    s.CouponDiscountFee = &v
    return s
}
