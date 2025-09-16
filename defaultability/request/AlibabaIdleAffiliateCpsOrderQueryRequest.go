package request

import (
	"github.com/crazyunix/topsdk/defaultability/domain"
	"github.com/crazyunix/topsdk/util"
)

type AlibabaIdleAffiliateCpsOrderQueryRequest struct {
	/*
	   订单查询入参     */
	TradeOrderVo *domain.AlibabaIdleAffiliateCpsOrderQueryTradeOrderVO `json:"trade_order_vo,omitempty" required:"false" `
}

func (s *AlibabaIdleAffiliateCpsOrderQueryRequest) SetTradeOrderVo(v domain.AlibabaIdleAffiliateCpsOrderQueryTradeOrderVO) *AlibabaIdleAffiliateCpsOrderQueryRequest {
	s.TradeOrderVo = &v
	return s
}

func (req *AlibabaIdleAffiliateCpsOrderQueryRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.TradeOrderVo != nil {
		paramMap["trade_order_vo"] = util.ConvertStruct(*req.TradeOrderVo)
	}
	return paramMap
}

func (req *AlibabaIdleAffiliateCpsOrderQueryRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
