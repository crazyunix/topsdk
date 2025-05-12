package domain


type TaobaoTbkRtaConsumerMatchData struct {
    /*
        返回结果列表     */
    ResultList  *[]TaobaoTbkRtaConsumerMatchResultlist `json:"result_list,omitempty" `

    /*
        策略ID的匹配结果，仅在入参strategy_id_list字段非空时返回     */
    StrategyResultList  *[]TaobaoTbkRtaConsumerMatchStrategyResultList `json:"strategy_result_list,omitempty" `

}

func (s *TaobaoTbkRtaConsumerMatchData) SetResultList(v []TaobaoTbkRtaConsumerMatchResultlist) *TaobaoTbkRtaConsumerMatchData {
    s.ResultList = &v
    return s
}
func (s *TaobaoTbkRtaConsumerMatchData) SetStrategyResultList(v []TaobaoTbkRtaConsumerMatchStrategyResultList) *TaobaoTbkRtaConsumerMatchData {
    s.StrategyResultList = &v
    return s
}
