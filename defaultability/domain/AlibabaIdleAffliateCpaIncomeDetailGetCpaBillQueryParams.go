package domain


type AlibabaIdleAffliateCpaIncomeDetailGetCpaBillQueryParams struct {
    /*
        佣金单id集合     */
    BillIdList  *[]string `json:"bill_id_list,omitempty" `

}

func (s *AlibabaIdleAffliateCpaIncomeDetailGetCpaBillQueryParams) SetBillIdList(v []string) *AlibabaIdleAffliateCpaIncomeDetailGetCpaBillQueryParams {
    s.BillIdList = &v
    return s
}
