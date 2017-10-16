package vkapi

import (
	"encoding/json"
	"sync"
)

type Api struct {
	AccessToken string
	retryCount  int
	sync.Mutex
}

type TokenData struct {
	ClientId     int
	ClientSecret string
	Code         string
	RedirectUri  string
}

type Response struct {
	Response json.RawMessage `json:"response"`
	Error    ResponseError   `json:"error"`
}

type ResponseError struct {
	ErrorCode     int                 `json:"error_code"`
	ErrorMsg      string              `json:"error_msg"`
	RequestParams []map[string]string `json:"request_params"`
}

/*
	Получение токена
*/

type GetTokenAns struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	UserId           int    `json:"user_id"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

/*
	Users
*/

type UsersGetAns struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Photo100  string `json:"photo_100"`
	Site      string `json:"site"`
	Sex       int    `json:"sex"`
	Status    string `json:"status"`
}

/*
	Wall
*/

type WallGetByIdAns struct {
	Id          int           `json:"id"`
	FromId      int           `json:"from_id"`
	OwnerId     int           `json:"owner_id"`
	Date        int           `json:"date"`
	Text        string        `json:"text"`
	Attachments []Attachments `json:"attachments"`
	Comments    LikeData      `json:"comments"`
	Likes       LikeData      `json:"likes"`
	Reposts     LikeData      `json:"reposts"`
	Views       LikeData      `json:"views"`
}

type Attachments struct {
	Type        string           `json:"type"`
	Photo       *json.RawMessage `json:"photo"`
	Audio       *json.RawMessage `json:"audio"`
	Video       *json.RawMessage `json:"video"`
	Poll        *json.RawMessage `json:"poll"`
	Page        *json.RawMessage `json:"page"`
	Album       *json.RawMessage `json:"album"`
	Link        *json.RawMessage `json:"link"`
	Doc         *json.RawMessage `json:"doc"`
	Note        *json.RawMessage `json:"note"`
	Sticker     *json.RawMessage `json:"sticker"`
	PrettyCards *json.RawMessage `json:"pretty_cards"`
}

type AttachmentsPrettyCards struct {
	Cards []AttachmentsPrettyCardsCards `json:"cards"`
}

type AttachmentsPrettyCardsCards struct {
	CardId  string `json:"card_id"`
	LinkUrl string `json:"link_url"`
	Title   string `json:"title"`
}

type LikeData struct {
	Count int `json:"count"`
}

/*
	Utils
*/
type UtilsGetShortLinkAns struct {
	ShortUrl  string `json:"short_url"`
	Url       string `json:"url"`
	Key       string `json:"key"`
	AccessKey string `json:"access_key"`
}

type UtilsGetLinkStatsAns struct {
	Key   string                   `json:"key"`
	Stats []UtilsGetLinkStatsStats `json:"stats"`
}

type UtilsGetLinkStatsStats struct {
	Timestamp int         `json:"timestamp"`
	Views     int         `json:"views"`
	SexAge    []SexAge    `json:"sex_age"`
	Countries []Countries `json:"countries"`
	Cities    []Cities    `json:"cities"`
}

/*
	Ads
*/

type AdsGetCampaignsAns struct {
	Id   int    `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

type AdsGetAdsLayoutAns struct {
	Id         string `json:"id"`
	CampaignId int    `json:"campaign_id"`
	Title      string `json:"title"`
	LinkUrl    string `json:"link_url"`
}

type AdsGetStatisticsAns struct {
	Id       int                        `json:"id"`
	Type     string                     `json:"type"`
	StatsBug []json.RawMessage          `json:"stats"`
	Stats    []AdsGetStatisticsAnsStats `json:"-"`
}

type AdsGetStatisticsAnsStats struct {
	Day         string `json:"day"`
	Spent       string `json:"spent"`
	Impressions int    `json:"impressions"`
	Clicks      int    `json:"clicks"`
	Reach       int    `json:"reach"`
}

type AdsGetStatisticsAnsStatsBug struct {
	Day         string `json:"day"`
	Spent       string `json:"spent"`
	Impressions string `json:"impressions"`
	Clicks      int    `json:"clicks"`
	Reach       int    `json:"reach"`
}

/*
	Other
*/

type SexAge struct {
	AgeRange string `json:"age_range"`
	Female   int    `json:"female"`
	Male     int    `json:"male"`
}

type Countries struct {
	CountryId int `json:"country_id"`
	Views     int `json:"views"`
}

type Cities struct {
	CityId int `json:"city_id"`
	Views  int `json:"views"`
}
