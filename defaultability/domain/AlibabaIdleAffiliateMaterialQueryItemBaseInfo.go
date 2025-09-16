package domain

import (
	"github.com/crazyunix/topsdk/util"
)

type AlibabaIdleAffiliateMaterialQueryItemBaseInfo struct {
	/*
	   商品id     */
	ItemId *string `json:"item_id,omitempty" `

	/*
	   商品售价     */
	ReservePrice *string `json:"reserve_price,omitempty" `

	/*
	   商品原价     */
	OriginalPrice *string `json:"original_price,omitempty" `

	/*
	   商品标题     */
	ItemTitle *string `json:"item_title,omitempty" `

	/*
	   新旧程度 1~9     */
	StuffStatus *string `json:"stuff_status,omitempty" `

	/*
	   商品图片     */
	ImageUrls *[]string `json:"image_urls,omitempty" `

	/*
	   交易类型 拍卖:a 一口价:b     */
	AuctionType *string `json:"auction_type,omitempty" `

	/*
	   商品描述     */
	ItemDesc *string `json:"item_desc,omitempty" `

	/*
	   类目名     */
	CategoryName *string `json:"category_name,omitempty" `

	/*
	   商品状态     */
	Status *string `json:"status,omitempty" `

	/*
	   是否多库存商品 true是 false否     */
	SkuItem *bool `json:"sku_item,omitempty" `

	/*
	   商品状态描述     */
	StatusDesc *string `json:"status_desc,omitempty" `

	/*
	   2024-12-20 15:15:25     */
	CreateTime *util.LocalTime `json:"create_time,omitempty" `

	/*
	   2024-12-20 15:15:25     */
	UpdateTime *util.LocalTime `json:"update_time,omitempty" `

	/*
	   true     */
	IsFishShop *bool `json:"is_fish_shop,omitempty" `

	/*
	   5     */
	UserFishShopGrade *int64 `json:"user_fish_shop_grade,omitempty" `

	/*
	   5极好 4优秀 3一般     */
	SellerLevelCode *int64 `json:"seller_level_code,omitempty" `

	/*
	   true     */
	YhbItem *bool `json:"yhb_item,omitempty" `
}

func (s *AlibabaIdleAffiliateMaterialQueryItemBaseInfo) SetItemId(v string) *AlibabaIdleAffiliateMaterialQueryItemBaseInfo {
	s.ItemId = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialQueryItemBaseInfo) SetReservePrice(v string) *AlibabaIdleAffiliateMaterialQueryItemBaseInfo {
	s.ReservePrice = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialQueryItemBaseInfo) SetOriginalPrice(v string) *AlibabaIdleAffiliateMaterialQueryItemBaseInfo {
	s.OriginalPrice = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialQueryItemBaseInfo) SetItemTitle(v string) *AlibabaIdleAffiliateMaterialQueryItemBaseInfo {
	s.ItemTitle = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialQueryItemBaseInfo) SetStuffStatus(v string) *AlibabaIdleAffiliateMaterialQueryItemBaseInfo {
	s.StuffStatus = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialQueryItemBaseInfo) SetImageUrls(v []string) *AlibabaIdleAffiliateMaterialQueryItemBaseInfo {
	s.ImageUrls = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialQueryItemBaseInfo) SetAuctionType(v string) *AlibabaIdleAffiliateMaterialQueryItemBaseInfo {
	s.AuctionType = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialQueryItemBaseInfo) SetItemDesc(v string) *AlibabaIdleAffiliateMaterialQueryItemBaseInfo {
	s.ItemDesc = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialQueryItemBaseInfo) SetCategoryName(v string) *AlibabaIdleAffiliateMaterialQueryItemBaseInfo {
	s.CategoryName = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialQueryItemBaseInfo) SetStatus(v string) *AlibabaIdleAffiliateMaterialQueryItemBaseInfo {
	s.Status = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialQueryItemBaseInfo) SetSkuItem(v bool) *AlibabaIdleAffiliateMaterialQueryItemBaseInfo {
	s.SkuItem = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialQueryItemBaseInfo) SetStatusDesc(v string) *AlibabaIdleAffiliateMaterialQueryItemBaseInfo {
	s.StatusDesc = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialQueryItemBaseInfo) SetCreateTime(v util.LocalTime) *AlibabaIdleAffiliateMaterialQueryItemBaseInfo {
	s.CreateTime = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialQueryItemBaseInfo) SetUpdateTime(v util.LocalTime) *AlibabaIdleAffiliateMaterialQueryItemBaseInfo {
	s.UpdateTime = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialQueryItemBaseInfo) SetIsFishShop(v bool) *AlibabaIdleAffiliateMaterialQueryItemBaseInfo {
	s.IsFishShop = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialQueryItemBaseInfo) SetUserFishShopGrade(v int64) *AlibabaIdleAffiliateMaterialQueryItemBaseInfo {
	s.UserFishShopGrade = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialQueryItemBaseInfo) SetSellerLevelCode(v int64) *AlibabaIdleAffiliateMaterialQueryItemBaseInfo {
	s.SellerLevelCode = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialQueryItemBaseInfo) SetYhbItem(v bool) *AlibabaIdleAffiliateMaterialQueryItemBaseInfo {
	s.YhbItem = &v
	return s
}
