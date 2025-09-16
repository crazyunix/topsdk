package request

import (
	"github.com/crazyunix/topsdk/defaultability/domain"
	"github.com/crazyunix/topsdk/util"
)

type AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryRequest struct {
	/*
	   入参     */
	IncomeDetailVo *domain.AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryIncomeDetailVO `json:"income_detail_vo,omitempty" required:"false" `
}

func (s *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryRequest) SetIncomeDetailVo(v domain.AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryIncomeDetailVO) *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryRequest {
	s.IncomeDetailVo = &v
	return s
}

func (req *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.IncomeDetailVo != nil {
		paramMap["income_detail_vo"] = util.ConvertStruct(*req.IncomeDetailVo)
	}
	return paramMap
}

func (req *AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
