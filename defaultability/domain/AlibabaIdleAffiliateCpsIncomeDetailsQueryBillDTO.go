package domain

import (
        "github.com/crazyunix/topsdk/util"
    )

type AlibabaIdleAffiliateCpsIncomeDetailsQueryBillDTO struct {
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

func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryBillDTO) SetStateText(v string) *AlibabaIdleAffiliateCpsIncomeDetailsQueryBillDTO {
    s.StateText = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryBillDTO) SetSettleAmountText(v string) *AlibabaIdleAffiliateCpsIncomeDetailsQueryBillDTO {
    s.SettleAmountText = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryBillDTO) SetAssessAmountText(v string) *AlibabaIdleAffiliateCpsIncomeDetailsQueryBillDTO {
    s.AssessAmountText = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryBillDTO) SetSubPublisherId(v string) *AlibabaIdleAffiliateCpsIncomeDetailsQueryBillDTO {
    s.SubPublisherId = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryBillDTO) SetState(v int64) *AlibabaIdleAffiliateCpsIncomeDetailsQueryBillDTO {
    s.State = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryBillDTO) SetAccountingTime(v util.LocalTime) *AlibabaIdleAffiliateCpsIncomeDetailsQueryBillDTO {
    s.AccountingTime = &v
    return s
}
func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryBillDTO) SetAccountingWriteOffTime(v util.LocalTime) *AlibabaIdleAffiliateCpsIncomeDetailsQueryBillDTO {
    s.AccountingWriteOffTime = &v
    return s
}
