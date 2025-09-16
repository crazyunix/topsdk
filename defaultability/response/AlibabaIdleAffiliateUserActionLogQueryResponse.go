package response

import (
	"github.com/crazyunix/topsdk/defaultability/domain"
)

type AlibabaIdleAffiliateUserActionLogQueryResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   分页查询结果
	*/
	Result domain.AlibabaIdleAffiliateUserActionLogQueryIdleAffiliatePageResult `json:"result,omitempty" `
}
