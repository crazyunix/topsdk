package response

import (
	"github.com/crazyunix/topsdk/defaultability/domain"
)

type AlibabaIdleAffiliateMaterialExactGetResponse struct {

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
	Result domain.AlibabaIdleAffiliateMaterialExactGetIdleAffiliateCommonResult `json:"result,omitempty" `
}
