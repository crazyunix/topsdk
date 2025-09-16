package domain


type AlibabaIdleAffiliateUserActionLogQueryUserActionLogDTO struct {
    /*
        加密用户id     */
    UserId  *string `json:"user_id,omitempty" `

    /*
        用户是否次留（用户在T日和T+1日均有访问行为，T为查询日期）     */
    DayOneUserRetention  *bool `json:"day_one_user_retention,omitempty" `

    /*
        用户是否为卡券新用户     */
    CardCouponNewUser  *bool `json:"card_coupon_new_user,omitempty" `

    /*
        子推广者id     */
    SubPublisherId  *string `json:"sub_publisher_id,omitempty" `

}

func (s *AlibabaIdleAffiliateUserActionLogQueryUserActionLogDTO) SetUserId(v string) *AlibabaIdleAffiliateUserActionLogQueryUserActionLogDTO {
    s.UserId = &v
    return s
}
func (s *AlibabaIdleAffiliateUserActionLogQueryUserActionLogDTO) SetDayOneUserRetention(v bool) *AlibabaIdleAffiliateUserActionLogQueryUserActionLogDTO {
    s.DayOneUserRetention = &v
    return s
}
func (s *AlibabaIdleAffiliateUserActionLogQueryUserActionLogDTO) SetCardCouponNewUser(v bool) *AlibabaIdleAffiliateUserActionLogQueryUserActionLogDTO {
    s.CardCouponNewUser = &v
    return s
}
func (s *AlibabaIdleAffiliateUserActionLogQueryUserActionLogDTO) SetSubPublisherId(v string) *AlibabaIdleAffiliateUserActionLogQueryUserActionLogDTO {
    s.SubPublisherId = &v
    return s
}
