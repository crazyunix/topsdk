package request

import (
	"github.com/crazyunix/topsdk/defaultability/domain"
	"github.com/crazyunix/topsdk/util"
)

type AlibabaIdleAffiliateUserActionLogQueryRequest struct {
	/*
	   行为查询参数     */
	UserActionLogQueryParams *domain.AlibabaIdleAffiliateUserActionLogQueryUserActionLogQueryParams `json:"user_action_log_query_params" required:"true" `
}

func (s *AlibabaIdleAffiliateUserActionLogQueryRequest) SetUserActionLogQueryParams(v domain.AlibabaIdleAffiliateUserActionLogQueryUserActionLogQueryParams) *AlibabaIdleAffiliateUserActionLogQueryRequest {
	s.UserActionLogQueryParams = &v
	return s
}

func (req *AlibabaIdleAffiliateUserActionLogQueryRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.UserActionLogQueryParams != nil {
		paramMap["user_action_log_query_params"] = util.ConvertStruct(*req.UserActionLogQueryParams)
	}
	return paramMap
}

func (req *AlibabaIdleAffiliateUserActionLogQueryRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
