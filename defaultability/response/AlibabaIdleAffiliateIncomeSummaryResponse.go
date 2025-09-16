package response

import (
	"github.com/crazyunix/topsdk/defaultability/domain"
)

type AlibabaIdleAffiliateIncomeSummaryResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   result
	*/
	Result domain.AlibabaIdleAffiliateIncomeSummaryIdleAffiliateCommonResult `json:"result,omitempty" `
}
