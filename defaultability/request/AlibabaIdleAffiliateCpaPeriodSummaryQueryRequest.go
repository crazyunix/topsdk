package request

import (
	"github.com/crazyunix/topsdk/defaultability/domain"
	"github.com/crazyunix/topsdk/util"
)

type AlibabaIdleAffiliateCpaPeriodSummaryQueryRequest struct {
	/*
	   请求参数     */
	CpaPeriodBillQueryParams *domain.AlibabaIdleAffiliateCpaPeriodSummaryQueryCpaPeriodBillQueryParams `json:"cpa_period_bill_query_params,omitempty" required:"false" `
}

func (s *AlibabaIdleAffiliateCpaPeriodSummaryQueryRequest) SetCpaPeriodBillQueryParams(v domain.AlibabaIdleAffiliateCpaPeriodSummaryQueryCpaPeriodBillQueryParams) *AlibabaIdleAffiliateCpaPeriodSummaryQueryRequest {
	s.CpaPeriodBillQueryParams = &v
	return s
}

func (req *AlibabaIdleAffiliateCpaPeriodSummaryQueryRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.CpaPeriodBillQueryParams != nil {
		paramMap["cpa_period_bill_query_params"] = util.ConvertStruct(*req.CpaPeriodBillQueryParams)
	}
	return paramMap
}

func (req *AlibabaIdleAffiliateCpaPeriodSummaryQueryRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
