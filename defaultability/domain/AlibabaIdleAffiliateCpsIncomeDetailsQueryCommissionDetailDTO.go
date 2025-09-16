package domain


type AlibabaIdleAffiliateCpsIncomeDetailsQueryCommissionDetailDTO struct {
    /*
        billInfo     */
    BillInfo  *AlibabaIdleAffiliateCpsIncomeDetailsQueryBillDTO `json:"bill_info,omitempty" `

    /*
        billId     */
    BillId  *string `json:"bill_id,omitempty" `

    /*
        tradeOrderDTO     */
    TradeOrderDTO  *AlibabaIdleAffiliateCpsIncomeDetailsQueryTradeOrderDTO `json:"trade_order_d_t_o,omitempty" `

    /*
        未加密订单号     */
    PlainBillId  *string `json:"plain_bill_id,omitempty" `

    /*
        订单跳转链接     */
    Deeplink  *string `json:"deeplink,omitempty" `

}

func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryCommissionDetailDTO) SetBillInfo(v AlibabaIdleAffiliateCpsIncomeDetailsQueryBillDTO) *AlibabaIdleAffiliateCpsIncomeDetailsQueryCommissionDetailDTO {
    s.BillInfo = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryCommissionDetailDTO) SetBillId(v string) *AlibabaIdleAffiliateCpsIncomeDetailsQueryCommissionDetailDTO {
    s.BillId = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryCommissionDetailDTO) SetTradeOrderDTO(v AlibabaIdleAffiliateCpsIncomeDetailsQueryTradeOrderDTO) *AlibabaIdleAffiliateCpsIncomeDetailsQueryCommissionDetailDTO {
    s.TradeOrderDTO = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryCommissionDetailDTO) SetPlainBillId(v string) *AlibabaIdleAffiliateCpsIncomeDetailsQueryCommissionDetailDTO {
    s.PlainBillId = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryCommissionDetailDTO) SetDeeplink(v string) *AlibabaIdleAffiliateCpsIncomeDetailsQueryCommissionDetailDTO {
    s.Deeplink = &v
    return s
}
