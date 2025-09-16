package request

import (
	"github.com/crazyunix/topsdk/defaultability/domain"
	"github.com/crazyunix/topsdk/util"
)

type AlibabaIdleAffiliateCpaIncomeDetailsQueryRequest struct {
	/*
	   入参     */
	CpaBillQueryParams *domain.AlibabaIdleAffiliateCpaIncomeDetailsQueryCpaBillQueryParams `json:"cpa_bill_query_params,omitempty" required:"false" `
}

func (s *AlibabaIdleAffiliateCpaIncomeDetailsQueryRequest) SetCpaBillQueryParams(v domain.AlibabaIdleAffiliateCpaIncomeDetailsQueryCpaBillQueryParams) *AlibabaIdleAffiliateCpaIncomeDetailsQueryRequest {
	s.CpaBillQueryParams = &v
	return s
}

func (req *AlibabaIdleAffiliateCpaIncomeDetailsQueryRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.CpaBillQueryParams != nil {
		paramMap["cpa_bill_query_params"] = util.ConvertStruct(*req.CpaBillQueryParams)
	}
	return paramMap
}

func (req *AlibabaIdleAffiliateCpaIncomeDetailsQueryRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
