package domain


type AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryCommissionDetailDTO struct {
    /*
        billInfo     */
    BillInfo  *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryBillDTO `json:"bill_info,omitempty" `

    /*
        billId     */
    BillId  *string `json:"bill_id,omitempty" `

    /*
        tradeOrderDTO     */
    TradeOrderDTO  *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryTradeOrderDTO `json:"trade_order_d_t_o,omitempty" `

}

func (s *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryCommissionDetailDTO) SetBillInfo(v AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryBillDTO) *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryCommissionDetailDTO {
    s.BillInfo = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryCommissionDetailDTO) SetBillId(v string) *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryCommissionDetailDTO {
    s.BillId = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryCommissionDetailDTO) SetTradeOrderDTO(v AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryTradeOrderDTO) *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryCommissionDetailDTO {
    s.TradeOrderDTO = &v
    return s
}
