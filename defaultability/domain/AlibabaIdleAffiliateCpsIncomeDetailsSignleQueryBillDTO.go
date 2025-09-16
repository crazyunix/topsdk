package domain

import (
        "github.com/crazyunix/topsdk/util"
    )

type AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryBillDTO struct {
    /*
        佣金状态描述     */
    StateText  *string `json:"state_text,omitempty" `

    /*
        结算佣金     */
    SettleAmountText  *string `json:"settle_amount_text,omitempty" `

    /*
        预估佣金     */
    AssessAmountText  *string `json:"assess_amount_text,omitempty" `

    /*
        二级推广者id     */
    SubPublisherId  *string `json:"sub_publisher_id,omitempty" `

    /*
        0待发放 1待到账 2已到账 3已失效     */
    State  *int64 `json:"state,omitempty" `

    /*
        记账时间     */
    AccountingTime  *util.LocalTime `json:"accounting_time,omitempty" `

    /*
        销账时间     */
    AccountingWriteOffTime  *util.LocalTime `json:"accounting_write_off_time,omitempty" `

}

func (s *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryBillDTO) SetStateText(v string) *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryBillDTO {
    s.StateText = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryBillDTO) SetSettleAmountText(v string) *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryBillDTO {
    s.SettleAmountText = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryBillDTO) SetAssessAmountText(v string) *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryBillDTO {
    s.AssessAmountText = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryBillDTO) SetSubPublisherId(v string) *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryBillDTO {
    s.SubPublisherId = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryBillDTO) SetState(v int64) *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryBillDTO {
    s.State = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryBillDTO) SetAccountingTime(v util.LocalTime) *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryBillDTO {
    s.AccountingTime = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryBillDTO) SetAccountingWriteOffTime(v util.LocalTime) *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryBillDTO {
    s.AccountingWriteOffTime = &v
    return s
}
