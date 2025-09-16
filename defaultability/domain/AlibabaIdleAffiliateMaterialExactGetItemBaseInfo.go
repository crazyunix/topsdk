package domain

import (
	"github.com/crazyunix/topsdk/util"
)

type AlibabaIdleAffiliateMaterialExactGetItemBaseInfo struct {
	/*
	   商品id     */
	ItemId *string `json:"item_id,omitempty" `

	/*
	   售价     */
	ReservePrice *string `json:"reserve_price,omitempty" `

	/*
	   库存     */
	Quantity *int64 `json:"quantity,omitempty" `

	/*
	   原价     */
	OriginalPrice *string `json:"original_price,omitempty" `

	/*
	   标题     */
	ItemTitle *string `json:"item_title,omitempty" `

	/*
	   新旧程度 1~9     */
	StuffStatus *string `json:"stuff_status,omitempty" `

	/*
	   商品图     */
	ImageUrls *[]string `json:"image_urls,omitempty" `

	/*
	   交易类型  拍卖:a 一口价:b     */
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

func (s *AlibabaIdleAffiliateMaterialExactGetItemBaseInfo) SetItemId(v string) *AlibabaIdleAffiliateMaterialExactGetItemBaseInfo {
	s.ItemId = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialExactGetItemBaseInfo) SetReservePrice(v string) *AlibabaIdleAffiliateMaterialExactGetItemBaseInfo {
	s.ReservePrice = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialExactGetItemBaseInfo) SetQuantity(v int64) *AlibabaIdleAffiliateMaterialExactGetItemBaseInfo {
	s.Quantity = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialExactGetItemBaseInfo) SetOriginalPrice(v string) *AlibabaIdleAffiliateMaterialExactGetItemBaseInfo {
	s.OriginalPrice = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialExactGetItemBaseInfo) SetItemTitle(v string) *AlibabaIdleAffiliateMaterialExactGetItemBaseInfo {
	s.ItemTitle = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialExactGetItemBaseInfo) SetStuffStatus(v string) *AlibabaIdleAffiliateMaterialExactGetItemBaseInfo {
	s.StuffStatus = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialExactGetItemBaseInfo) SetImageUrls(v []string) *AlibabaIdleAffiliateMaterialExactGetItemBaseInfo {
	s.ImageUrls = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialExactGetItemBaseInfo) SetAuctionType(v string) *AlibabaIdleAffiliateMaterialExactGetItemBaseInfo {
	s.AuctionType = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialExactGetItemBaseInfo) SetItemDesc(v string) *AlibabaIdleAffiliateMaterialExactGetItemBaseInfo {
	s.ItemDesc = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialExactGetItemBaseInfo) SetCategoryName(v string) *AlibabaIdleAffiliateMaterialExactGetItemBaseInfo {
	s.CategoryName = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialExactGetItemBaseInfo) SetStatus(v string) *AlibabaIdleAffiliateMaterialExactGetItemBaseInfo {
	s.Status = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialExactGetItemBaseInfo) SetStatusDesc(v string) *AlibabaIdleAffiliateMaterialExactGetItemBaseInfo {
	s.StatusDesc = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialExactGetItemBaseInfo) SetCreateTime(v util.LocalTime) *AlibabaIdleAffiliateMaterialExactGetItemBaseInfo {
	s.CreateTime = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialExactGetItemBaseInfo) SetUpdateTime(v util.LocalTime) *AlibabaIdleAffiliateMaterialExactGetItemBaseInfo {
	s.UpdateTime = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialExactGetItemBaseInfo) SetIsFishShop(v bool) *AlibabaIdleAffiliateMaterialExactGetItemBaseInfo {
	s.IsFishShop = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialExactGetItemBaseInfo) SetUserFishShopGrade(v int64) *AlibabaIdleAffiliateMaterialExactGetItemBaseInfo {
	s.UserFishShopGrade = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialExactGetItemBaseInfo) SetSellerLevelCode(v int64) *AlibabaIdleAffiliateMaterialExactGetItemBaseInfo {
	s.SellerLevelCode = &v
	return s
}
func (s *AlibabaIdleAffiliateMaterialExactGetItemBaseInfo) SetYhbItem(v bool) *AlibabaIdleAffiliateMaterialExactGetItemBaseInfo {
	s.YhbItem = &v
	return s
}
