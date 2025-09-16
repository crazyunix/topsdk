package request

import (
	"github.com/crazyunix/topsdk/defaultability/domain"
	"github.com/crazyunix/topsdk/util"
)

type AlibabaIdleAffiliateGeneralLinkConvertRequest struct {
	/*
	   转链入参     */
	LinkConvertVo *domain.AlibabaIdleAffiliateGeneralLinkConvertLinkConvertVO `json:"link_convert_vo" required:"true" `
}

func (s *AlibabaIdleAffiliateGeneralLinkConvertRequest) SetLinkConvertVo(v domain.AlibabaIdleAffiliateGeneralLinkConvertLinkConvertVO) *AlibabaIdleAffiliateGeneralLinkConvertRequest {
	s.LinkConvertVo = &v
	return s
}

func (req *AlibabaIdleAffiliateGeneralLinkConvertRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.LinkConvertVo != nil {
		paramMap["link_convert_vo"] = util.ConvertStruct(*req.LinkConvertVo)
	}
	return paramMap
}

func (req *AlibabaIdleAffiliateGeneralLinkConvertRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
