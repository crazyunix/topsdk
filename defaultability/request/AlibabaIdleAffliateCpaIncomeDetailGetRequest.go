package request

import (
	"github.com/crazyunix/topsdk/defaultability/domain"
	"github.com/crazyunix/topsdk/util"
)

type AlibabaIdleAffliateCpaIncomeDetailGetRequest struct {
	/*
	   入参     */
	CpaBillQueryParams *domain.AlibabaIdleAffliateCpaIncomeDetailGetCpaBillQueryParams `json:"cpa_bill_query_params,omitempty" required:"false" `
}

func (s *AlibabaIdleAffliateCpaIncomeDetailGetRequest) SetCpaBillQueryParams(v domain.AlibabaIdleAffliateCpaIncomeDetailGetCpaBillQueryParams) *AlibabaIdleAffliateCpaIncomeDetailGetRequest {
	s.CpaBillQueryParams = &v
	return s
}

func (req *AlibabaIdleAffliateCpaIncomeDetailGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.CpaBillQueryParams != nil {
		paramMap["cpa_bill_query_params"] = util.ConvertStruct(*req.CpaBillQueryParams)
	}
	return paramMap
}

func (req *AlibabaIdleAffliateCpaIncomeDetailGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
