package domain


type AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryIncomeDetailVO struct {
    /*
        佣金明细id     */
    BillIdList  *[]string `json:"bill_id_list,omitempty" `

    /*
        明文订单号     */
    PlainBillId  *string `json:"plain_bill_id,omitempty" `

    /*
        明文订单号列表     */
    PlainBillIdList  *[]string `json:"plain_bill_id_list,omitempty" `

}

func (s *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryIncomeDetailVO) SetBillIdList(v []string) *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryIncomeDetailVO {
    s.BillIdList = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryIncomeDetailVO) SetPlainBillId(v string) *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryIncomeDetailVO {
    s.PlainBillId = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryIncomeDetailVO) SetPlainBillIdList(v []string) *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryIncomeDetailVO {
    s.PlainBillIdList = &v
    return s
}
