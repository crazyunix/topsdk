package domain


type AlibabaIdleAffiliateGeneralLinkConvertLinkConvertDTO struct {
    /*
        短链     */
    ShortUrl  *string `json:"short_url,omitempty" `

    /*
        口令     */
    ShortTpwd  *string `json:"short_tpwd,omitempty" `

    /*
        deeplink链接 有端唤端到目标页 无端需自行判断跳转     */
    Deeplink  *string `json:"deeplink,omitempty" `

}

func (s *AlibabaIdleAffiliateGeneralLinkConvertLinkConvertDTO) SetShortUrl(v string) *AlibabaIdleAffiliateGeneralLinkConvertLinkConvertDTO {
    s.ShortUrl = &v
    return s
}
func (s *AlibabaIdleAffiliateGeneralLinkConvertLinkConvertDTO) SetShortTpwd(v string) *AlibabaIdleAffiliateGeneralLinkConvertLinkConvertDTO {
    s.ShortTpwd = &v
    return s
}
func (s *AlibabaIdleAffiliateGeneralLinkConvertLinkConvertDTO) SetDeeplink(v string) *AlibabaIdleAffiliateGeneralLinkConvertLinkConvertDTO {
    s.Deeplink = &v
    return s
}
