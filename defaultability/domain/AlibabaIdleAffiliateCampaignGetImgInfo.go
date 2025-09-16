package domain


type AlibabaIdleAffiliateCampaignGetImgInfo struct {
    /*
        width     */
    Width  *int64 `json:"width,omitempty" `

    /*
        url     */
    Url  *string `json:"url,omitempty" `

    /*
        height     */
    Height  *int64 `json:"height,omitempty" `

}

func (s *AlibabaIdleAffiliateCampaignGetImgInfo) SetWidth(v int64) *AlibabaIdleAffiliateCampaignGetImgInfo {
    s.Width = &v
    return s
}
func (s *AlibabaIdleAffiliateCampaignGetImgInfo) SetUrl(v string) *AlibabaIdleAffiliateCampaignGetImgInfo {
    s.Url = &v
    return s
}
func (s *AlibabaIdleAffiliateCampaignGetImgInfo) SetHeight(v int64) *AlibabaIdleAffiliateCampaignGetImgInfo {
    s.Height = &v
    return s
}
