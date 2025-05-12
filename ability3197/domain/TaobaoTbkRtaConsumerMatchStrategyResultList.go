package domain


type TaobaoTbkRtaConsumerMatchStrategyResultList struct {
    /*
        策略ID     */
    StrategyId  *string `json:"strategy_id,omitempty" `

    /*
        状态：1-符合活动要求 ，3-用户不匹配活动，4-系统异常，6-策略ID不存在，7-策略ID无效     */
    Status  *string `json:"status,omitempty" `

}

func (s *TaobaoTbkRtaConsumerMatchStrategyResultList) SetStrategyId(v string) *TaobaoTbkRtaConsumerMatchStrategyResultList {
    s.StrategyId = &v
    return s
}
func (s *TaobaoTbkRtaConsumerMatchStrategyResultList) SetStatus(v string) *TaobaoTbkRtaConsumerMatchStrategyResultList {
    s.Status = &v
    return s
}
