package request

import (
	"github.com/crazyunix/topsdk/defaultability/domain"
	"github.com/crazyunix/topsdk/util"
)

type AlibabaIdleAffiliateCpsIncomeDetailsQueryRequest struct {
	/*
	   入参     */
	IncomeDetailVo *domain.AlibabaIdleAffiliateCpsIncomeDetailsQueryIncomeDetailVO `json:"income_detail_vo,omitempty" required:"false" `
}

func (s *AlibabaIdleAffiliateCpsIncomeDetailsQueryRequest) SetIncomeDetailVo(v domain.AlibabaIdleAffiliateCpsIncomeDetailsQueryIncomeDetailVO) *AlibabaIdleAffiliateCpsIncomeDetailsQueryRequest {
	s.IncomeDetailVo = &v
	return s
}

func (req *AlibabaIdleAffiliateCpsIncomeDetailsQueryRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.IncomeDetailVo != nil {
		paramMap["income_detail_vo"] = util.ConvertStruct(*req.IncomeDetailVo)
	}
	return paramMap
}

func (req *AlibabaIdleAffiliateCpsIncomeDetailsQueryRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
