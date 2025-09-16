package request

import (
	"github.com/crazyunix/topsdk/defaultability/domain"
	"github.com/crazyunix/topsdk/util"
)

type AlibabaIdleAffiliateMaterialQueryRequest struct {
	/*
	   入参     */
	MaterialsQueryVo *domain.AlibabaIdleAffiliateMaterialQueryMaterialsQueryVO `json:"materials_query_vo" required:"true" `
}

func (s *AlibabaIdleAffiliateMaterialQueryRequest) SetMaterialsQueryVo(v domain.AlibabaIdleAffiliateMaterialQueryMaterialsQueryVO) *AlibabaIdleAffiliateMaterialQueryRequest {
	s.MaterialsQueryVo = &v
	return s
}

func (req *AlibabaIdleAffiliateMaterialQueryRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.MaterialsQueryVo != nil {
		paramMap["materials_query_vo"] = util.ConvertStruct(*req.MaterialsQueryVo)
	}
	return paramMap
}

func (req *AlibabaIdleAffiliateMaterialQueryRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
