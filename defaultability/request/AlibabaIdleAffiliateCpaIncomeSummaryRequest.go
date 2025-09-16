package request

import (
	"github.com/crazyunix/topsdk/defaultability/domain"
	"github.com/crazyunix/topsdk/util"
)

type AlibabaIdleAffiliateCpaIncomeSummaryRequest struct {
	/*
	   cpa佣金汇总查询入参     */
	CpaSummaryQueryParams *domain.AlibabaIdleAffiliateCpaIncomeSummaryCpaSummaryQueryParams `json:"cpa_summary_query_params,omitempty" required:"false" `
}

func (s *AlibabaIdleAffiliateCpaIncomeSummaryRequest) SetCpaSummaryQueryParams(v domain.AlibabaIdleAffiliateCpaIncomeSummaryCpaSummaryQueryParams) *AlibabaIdleAffiliateCpaIncomeSummaryRequest {
	s.CpaSummaryQueryParams = &v
	return s
}

func (req *AlibabaIdleAffiliateCpaIncomeSummaryRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.CpaSummaryQueryParams != nil {
		paramMap["cpa_summary_query_params"] = util.ConvertStruct(*req.CpaSummaryQueryParams)
	}
	return paramMap
}

func (req *AlibabaIdleAffiliateCpaIncomeSummaryRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
