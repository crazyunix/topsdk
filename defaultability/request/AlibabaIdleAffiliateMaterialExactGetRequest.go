package request

import (
	"github.com/crazyunix/topsdk/defaultability/domain"
	"github.com/crazyunix/topsdk/util"
)

type AlibabaIdleAffiliateMaterialExactGetRequest struct {
	/*
	   入参     */
	MaterialsExactQueryVo *domain.AlibabaIdleAffiliateMaterialExactGetMaterialsExactQueryVO `json:"materials_exact_query_vo,omitempty" required:"false" `
}

func (s *AlibabaIdleAffiliateMaterialExactGetRequest) SetMaterialsExactQueryVo(v domain.AlibabaIdleAffiliateMaterialExactGetMaterialsExactQueryVO) *AlibabaIdleAffiliateMaterialExactGetRequest {
	s.MaterialsExactQueryVo = &v
	return s
}

func (req *AlibabaIdleAffiliateMaterialExactGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.MaterialsExactQueryVo != nil {
		paramMap["materials_exact_query_vo"] = util.ConvertStruct(*req.MaterialsExactQueryVo)
	}
	return paramMap
}

func (req *AlibabaIdleAffiliateMaterialExactGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
