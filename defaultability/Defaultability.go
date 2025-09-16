package defaultability

import (
	"errors"
	"log"

	"github.com/crazyunix/topsdk"
	"github.com/crazyunix/topsdk/defaultability/request"
	"github.com/crazyunix/topsdk/defaultability/response"
	util "github.com/crazyunix/topsdk/util"
)

type Defaultability struct {
	Client *topsdk.TopClient
}

func NewDefaultability(client *topsdk.TopClient) *Defaultability {
	return &Defaultability{client}
}

/*
淘宝客-推广者-万能转链
*/
func (ability *Defaultability) TaobaoTbkDgGeneralLinkConvert(req *request.TaobaoTbkDgGeneralLinkConvertRequest) (*response.TaobaoTbkDgGeneralLinkConvertResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tbk.dg.general.link.convert", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTbkDgGeneralLinkConvertResponse{}
	if err != nil {
		log.Println("taobaoTbkDgGeneralLinkConvert error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
关键词过滤匹配
*/
func (ability *Defaultability) TaobaoKfcKeywordSearch(req *request.TaobaoKfcKeywordSearchRequest, session string) (*response.TaobaoKfcKeywordSearchResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.kfc.keyword.search", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoKfcKeywordSearchResponse{}
	if err != nil {
		log.Println("taobaoKfcKeywordSearch error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
淘宝客-推广者-物料id列表查询
*/
func (ability *Defaultability) TaobaoTbkOptimusTouMaterialIdsGet(req *request.TaobaoTbkOptimusTouMaterialIdsGetRequest) (*response.TaobaoTbkOptimusTouMaterialIdsGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tbk.optimus.tou.material.ids.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTbkOptimusTouMaterialIdsGetResponse{}
	if err != nil {
		log.Println("taobaoTbkOptimusTouMaterialIdsGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
淘宝客-推广者-物料搜索升级版
*/
func (ability *Defaultability) TaobaoTbkDgMaterialOptionalUpgrade(req *request.TaobaoTbkDgMaterialOptionalUpgradeRequest) (*response.TaobaoTbkDgMaterialOptionalUpgradeResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tbk.dg.material.optional.upgrade", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTbkDgMaterialOptionalUpgradeResponse{}
	if err != nil {
		log.Println("taobaoTbkDgMaterialOptionalUpgrade error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
淘宝客-推广者-物料精选升级版
*/
func (ability *Defaultability) TaobaoTbkDgMaterialRecommend(req *request.TaobaoTbkDgMaterialRecommendRequest) (*response.TaobaoTbkDgMaterialRecommendResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tbk.dg.material.recommend", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTbkDgMaterialRecommendResponse{}
	if err != nil {
		log.Println("taobaoTbkDgMaterialRecommend error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
闲鱼联盟-cps佣金明细单条查询
*/
func (ability *Defaultability) AlibabaIdleAffiliateCpsIncomeDetailsSignleQuery(req *request.AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryRequest, session string) (*response.AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("alibaba.idle.affiliate.cps.income.details.signle.query", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.AlibabaIdleAffiliateCpsIncomeDetailsSignleQueryResponse{}
	if err != nil {
		log.Println("alibabaIdleAffiliateCpsIncomeDetailsSignleQuery error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
闲鱼联盟-cpa佣金明细查询-精准查询
*/
func (ability *Defaultability) AlibabaIdleAffliateCpaIncomeDetailGet(req *request.AlibabaIdleAffliateCpaIncomeDetailGetRequest, session string) (*response.AlibabaIdleAffliateCpaIncomeDetailGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("alibaba.idle.affliate.cpa.income.detail.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.AlibabaIdleAffliateCpaIncomeDetailGetResponse{}
	if err != nil {
		log.Println("alibabaIdleAffliateCpaIncomeDetailGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
闲鱼联盟-cpa佣金账期汇总-查询
*/
func (ability *Defaultability) AlibabaIdleAffiliateCpaPeriodSummaryQuery(req *request.AlibabaIdleAffiliateCpaPeriodSummaryQueryRequest, session string) (*response.AlibabaIdleAffiliateCpaPeriodSummaryQueryResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("alibaba.idle.affiliate.cpa.period.summary.query", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.AlibabaIdleAffiliateCpaPeriodSummaryQueryResponse{}
	if err != nil {
		log.Println("alibabaIdleAffiliateCpaPeriodSummaryQuery error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
闲鱼联盟-cpa佣金明细查询
*/
func (ability *Defaultability) AlibabaIdleAffiliateCpaIncomeDetailsQuery(req *request.AlibabaIdleAffiliateCpaIncomeDetailsQueryRequest, session string) (*response.AlibabaIdleAffiliateCpaIncomeDetailsQueryResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("alibaba.idle.affiliate.cpa.income.details.query", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.AlibabaIdleAffiliateCpaIncomeDetailsQueryResponse{}
	if err != nil {
		log.Println("alibabaIdleAffiliateCpaIncomeDetailsQuery error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
闲鱼联盟-活动获取
*/
func (ability *Defaultability) AlibabaIdleAffiliateCampaignGet(req *request.AlibabaIdleAffiliateCampaignGetRequest, session string) (*response.AlibabaIdleAffiliateCampaignGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("alibaba.idle.affiliate.campaign.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.AlibabaIdleAffiliateCampaignGetResponse{}
	if err != nil {
		log.Println("alibabaIdleAffiliateCampaignGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
闲鱼联盟-cpa收入汇总
*/
func (ability *Defaultability) AlibabaIdleAffiliateCpaIncomeSummary(req *request.AlibabaIdleAffiliateCpaIncomeSummaryRequest, session string) (*response.AlibabaIdleAffiliateCpaIncomeSummaryResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("alibaba.idle.affiliate.cpa.income.summary", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.AlibabaIdleAffiliateCpaIncomeSummaryResponse{}
	if err != nil {
		log.Println("alibabaIdleAffiliateCpaIncomeSummary error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
闲鱼联盟-用户行为明细查询
*/
func (ability *Defaultability) AlibabaIdleAffiliateUserActionLogQuery(req *request.AlibabaIdleAffiliateUserActionLogQueryRequest, session string) (*response.AlibabaIdleAffiliateUserActionLogQueryResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("alibaba.idle.affiliate.user.action.log.query", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.AlibabaIdleAffiliateUserActionLogQueryResponse{}
	if err != nil {
		log.Println("alibabaIdleAffiliateUserActionLogQuery error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
获取商家所在分组及其已授权(广播)消息topics
*/
func (ability *Defaultability) TaobaoTmcUserGet(req *request.TaobaoTmcUserGetRequest) (*response.TaobaoTmcUserGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tmc.user.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTmcUserGetResponse{}
	if err != nil {
		log.Println("taobaoTmcUserGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
发布单条消息
*/
func (ability *Defaultability) TaobaoTmcMessageProduce(req *request.TaobaoTmcMessageProduceRequest, session string) (*response.TaobaoTmcMessageProduceResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.tmc.message.produce", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoTmcMessageProduceResponse{}
	if err != nil {
		log.Println("taobaoTmcMessageProduce error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
取消用户的消息服务
*/
func (ability *Defaultability) TaobaoTmcUserCancel(req *request.TaobaoTmcUserCancelRequest) (*response.TaobaoTmcUserCancelResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tmc.user.cancel", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTmcUserCancelResponse{}
	if err != nil {
		log.Println("taobaoTmcUserCancel error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
为已授权的用户开通消息服务
*/
func (ability *Defaultability) TaobaoTmcUserPermit(req *request.TaobaoTmcUserPermitRequest, session string) (*response.TaobaoTmcUserPermitResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.tmc.user.permit", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoTmcUserPermitResponse{}
	if err != nil {
		log.Println("taobaoTmcUserPermit error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
闲鱼联盟-物料指南
*/
func (ability *Defaultability) AlibabaIdleAffiliateMaterialGuideGet(req *request.AlibabaIdleAffiliateMaterialGuideGetRequest) (*response.AlibabaIdleAffiliateMaterialGuideGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.idle.affiliate.material.guide.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaIdleAffiliateMaterialGuideGetResponse{}
	if err != nil {
		log.Println("alibabaIdleAffiliateMaterialGuideGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
闲鱼联盟-推广转链
*/
func (ability *Defaultability) AlibabaIdleAffiliateGeneralLinkConvert(req *request.AlibabaIdleAffiliateGeneralLinkConvertRequest, session string) (*response.AlibabaIdleAffiliateGeneralLinkConvertResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("alibaba.idle.affiliate.general.link.convert", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.AlibabaIdleAffiliateGeneralLinkConvertResponse{}
	if err != nil {
		log.Println("alibabaIdleAffiliateGeneralLinkConvert error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
闲鱼联盟-cps订单查询
*/
func (ability *Defaultability) AlibabaIdleAffiliateCpsOrderQuery(req *request.AlibabaIdleAffiliateCpsOrderQueryRequest, session string) (*response.AlibabaIdleAffiliateCpsOrderQueryResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("alibaba.idle.affiliate.cps.order.query", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.AlibabaIdleAffiliateCpsOrderQueryResponse{}
	if err != nil {
		log.Println("alibabaIdleAffiliateCpsOrderQuery error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
闲鱼联盟-物料查询
*/
func (ability *Defaultability) AlibabaIdleAffiliateMaterialQuery(req *request.AlibabaIdleAffiliateMaterialQueryRequest, session string) (*response.AlibabaIdleAffiliateMaterialQueryResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("alibaba.idle.affiliate.material.query", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.AlibabaIdleAffiliateMaterialQueryResponse{}
	if err != nil {
		log.Println("alibabaIdleAffiliateMaterialQuery error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
闲鱼联盟-物料精确获取
*/
func (ability *Defaultability) AlibabaIdleAffiliateMaterialExactGet(req *request.AlibabaIdleAffiliateMaterialExactGetRequest, session string) (*response.AlibabaIdleAffiliateMaterialExactGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("alibaba.idle.affiliate.material.exact.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.AlibabaIdleAffiliateMaterialExactGetResponse{}
	if err != nil {
		log.Println("alibabaIdleAffiliateMaterialExactGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
闲鱼联盟-cps收入汇总
*/
func (ability *Defaultability) AlibabaIdleAffiliateIncomeSummary(req *request.AlibabaIdleAffiliateIncomeSummaryRequest, session string) (*response.AlibabaIdleAffiliateIncomeSummaryResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("alibaba.idle.affiliate.income.summary", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.AlibabaIdleAffiliateIncomeSummaryResponse{}
	if err != nil {
		log.Println("alibabaIdleAffiliateIncomeSummary error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
闲鱼联盟-cps佣金明细查询
*/
func (ability *Defaultability) AlibabaIdleAffiliateCpsIncomeDetailsQuery(req *request.AlibabaIdleAffiliateCpsIncomeDetailsQueryRequest, session string) (*response.AlibabaIdleAffiliateCpsIncomeDetailsQueryResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("alibaba.idle.affiliate.cps.income.details.query", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.AlibabaIdleAffiliateCpsIncomeDetailsQueryResponse{}
	if err != nil {
		log.Println("alibabaIdleAffiliateCpsIncomeDetailsQuery error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
