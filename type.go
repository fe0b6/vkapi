package vkapi

import (
	"encoding/json"
	"log"
	"regexp"
	"sync"
)

var (
	executeErrorSkipReg *regexp.Regexp
	httpErrorReg        *regexp.Regexp
	contMap             contextMap
	exited              bool
)

func init() {
	executeErrorSkipReg = regexp.MustCompile("server sent GOAWAY|User authorization failed|unexpected EOF|Database problems, try later|Internal Server Error|Bad Request|Gateway Timeout|Bad Gateway|could not check access_token now|connection reset by peer|Request Entity Too Large|response size is too big|context canceled")
	httpErrorReg = regexp.MustCompile("unexpected EOF|server sent GOAWAY|Bad Request|Internal Server Error|Request Entity Too Large|context canceled")

	contMap = contextMap{h: make(map[string]func())}
}

type contextMap struct {
	h map[string]func()
	sync.Mutex
}

// API - главный объект
type API struct {
	AccessToken    string
	retryCount     int
	httpRetryCount int
	ErrorToSkip    []string
	sync.Mutex
	ExecuteErrors []ExecuteErrors
	ExecuteCode   string
}

// TokenData - объект получения токена
type TokenData struct {
	ClientID     int
	ClientSecret string
	Code         string
	RedirectURI  string
}

// Response - объект ответа VK
type Response struct {
	Response      json.RawMessage `json:"response"`
	Error         ResponseError   `json:"error"`
	ExecuteErrors []ExecuteErrors `json:"execute_errors"`
}

// ExecuteErrors - объект ошибок execute
type ExecuteErrors struct {
	Method    string `json:"method"`
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

// ResponseError - объект ошибки выболнения запроса
type ResponseError struct {
	ErrorCode     int                 `json:"error_code"`
	ErrorMsg      string              `json:"error_msg"`
	RequestParams []map[string]string `json:"request_params"`
}

/*
	Получение токена
*/

// GetTokenAns - объект ответа при получении покена
type GetTokenAns struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	UserID           int    `json:"user_id"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

/*
	Users
*/

// UsersGetAns - объект ответа при запросе пользователей
type UsersGetAns struct {
	ID                     int                  `json:"id"`
	FirstName              string               `json:"first_name"`
	LastName               string               `json:"last_name"`
	About                  string               `json:"about"`
	Activities             string               `json:"activities"`
	Bdate                  string               `json:"bdate"`
	Blacklisted            int                  `json:"blacklisted"`
	BlacklistedByMe        int                  `json:"blacklisted_by_me"`
	Books                  string               `json:"books"`
	CanPost                int                  `json:"can_post"`
	CanSeeAllPosts         int                  `json:"can_see_all_posts"`
	CanSeeAudio            int                  `json:"can_see_audio"`
	CanSendFriendRequest   int                  `json:"can_send_friend_request"`
	CanWritePrivateMessage int                  `json:"can_write_private_message"`
	Career                 UserCareer           `json:"career"`
	City                   City                 `json:"city"`
	CommonCount            int                  `json:"common_count"`
	Skype                  string               `json:"skype"`
	Facebook               string               `json:"facebook"`
	Twitter                string               `json:"twitter"`
	Livejournal            string               `json:"livejournal"`
	Instagram              string               `json:"instagram"`
	Contacts               UserContacts         `json:"contacts"`
	Counters               UserCounters         `json:"counters"`
	Country                City                 `json:"country"`
	Domain                 string               `json:"domain"`
	Education              UserEducation        `json:"education"`
	FollowersCount         int                  `json:"followers_count"`
	FriendStatus           int                  `json:"friend_status"`
	Games                  string               `json:"games"`
	HasMobile              int                  `json:"has_mobile"`
	HasPhoto               int                  `json:"has_photo"`
	HomeTown               string               `json:"home_town"`
	Interests              string               `json:"interests"`
	IsFavorite             int                  `json:"is_favorite"`
	IsFriend               int                  `json:"is_friend"`
	IsHiddenFromFeed       int                  `json:"is_hidden_from_feed"`
	LastSeen               UserLastSeen         `json:"last_seen"`
	Lists                  string               `json:"lists"`
	MaidenName             string               `json:"maiden_name"`
	Military               UserMilitary         `json:"military"`
	Movies                 string               `json:"movies"`
	Music                  string               `json:"music"`
	Nickname               string               `json:"nickname"`
	Occupation             UserOccupation       `json:"occupation"`
	Online                 int                  `json:"online"`
	Personal               UserPersonal         `json:"personal"`
	Photo50                string               `json:"photo_50"`
	Photo100               string               `json:"photo_100"`
	Photo200Orig           string               `json:"photo_200_orig"`
	Photo200               string               `json:"photo_200"`
	Photo400Orig           string               `json:"photo_400_orig"`
	PhotoID                string               `json:"photo_id"`
	PhotoMax               string               `json:"photo_max"`
	PhotoMaxOrig           string               `json:"photo_max_orig"`
	Quotes                 string               `json:"quotes"`
	Relatives              []UserRelatives      `json:"relatives"`
	Relation               int                  `json:"relation"`
	RelationPartner        UsersRelationPartner `json:"relation_partner"`
	Schools                []UserSchools        `json:"schools"`
	ScreenName             string               `json:"screen_name"`
	Sex                    int                  `json:"sex"`
	Site                   string               `json:"site"`
	Status                 string               `json:"status"`
	Timezone               int                  `json:"timezone"`
	Trending               int                  `json:"trending"`
	Tv                     string               `json:"tv"`
	Universities           []UserUniversities   `json:"universities"`
	Verified               int                  `json:"verified"`
	Role                   string               `json:"role"`
}

// UsersRelationPartner - объект партнера
type UsersRelationPartner struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// UserCareer - Объект информации о карьере человека
type UserCareer struct {
	GroupID   int    `json:"group_id"`
	Company   string `json:"company"`
	CountryID int    `json:"country_id"`
	CityID    int    `json:"city_id"`
	CityName  string `json:"city_name"`
	From      int    `json:"from"`
	Until     int    `json:"until"`
	Position  string `json:"position"`
}

// UserUniversities - информация о университете
type UserUniversities struct {
	ID              int    `json:"id"`
	Country         int    `json:"country"`
	City            int    `json:"city"`
	Name            string `json:"name"`
	Faculty         int    `json:"faculty"`
	FacultyName     string `json:"faculty_name"`
	Chair           int    `json:"chair"`
	ChairName       string `json:"chair_name"`
	Graduation      int    `json:"graduation"`
	EducationForm   string `json:"education_form"`
	EducationStatus string `json:"education_status"`
}

// UserSchools - информация о школе веловека
type UserSchools struct {
	ID            int    `json:"id"`
	Country       int    `json:"country"`
	City          int    `json:"city"`
	Name          string `json:"name"`
	YearFrom      int    `json:"year_from"`
	YearTo        int    `json:"year_to"`
	YearGraduated int    `json:"year_graduated"`
	Class         string `json:"class"`
	Speciality    string `json:"speciality"`
	Type          int    `json:"type"`
	TypeStr       string `json:"type_str"`
}

// UserRelatives - объект с информацией о родвственнике человека
type UserRelatives struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// UserPersonal - Объект информации о жизненной позиции человека
type UserPersonal struct {
	Political  int      `json:"political"`
	Langs      []string `json:"langs"`
	Religion   string   `json:"religion"`
	InspiredBy string   `json:"inspired_by"`
	PeopleMain int      `json:"people_main"`
	LifeMain   int      `json:"life_main"`
	Smoking    int      `json:"smoking"`
	Alcohol    int      `json:"alcohol"`
}

// UserOccupation - Объект информации о текущем занятии человека
type UserOccupation struct {
	Type string `json:"type"`
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// UserMilitary - информация о военной службе
type UserMilitary struct {
	Unit      string `json:"unit"`
	UnitID    int    `json:"unit_id"`
	CountryID int    `json:"country_id"`
	From      int    `json:"from"`
	Until     int    `json:"until"`
}

// UserLastSeen - Информация о последнем посещении
type UserLastSeen struct {
	Time     int `json:"time"`
	Platform int `json:"platform"`
}

// UserEducation - объект информации об университете человека
type UserEducation struct {
	University     int    `json:"university"`
	UniversityName string `json:"university_name"`
	Faculty        int    `json:"faculty"`
	FacultyName    string `json:"faculty_name"`
	Graduation     int    `json:"graduation"`
}

// UserContacts - Объект контактов человека
type UserContacts struct {
	MobilePhone string `json:"mobile_phone"`
	HomePhone   string `json:"home_phone"`
}

// UserCounters - объект счетчиков человека
type UserCounters struct {
	Albums        int `json:"albums"`
	Videos        int `json:"videos"`
	Audios        int `json:"audios"`
	Photos        int `json:"photos"`
	Notes         int `json:"notes"`
	Friends       int `json:"friends"`
	Groups        int `json:"groups"`
	OnlineFriends int `json:"online_friends"`
	MutualFriends int `json:"mutual_friends"`
	UserVideos    int `json:"user_videos"`
	Followers     int `json:"followers"`
	Pages         int `json:"pages"`
}

// City - Объект с информацией о городе
type City struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// ScriptUsersMultiGetAns - объект списка информации о пользователях
type ScriptUsersMultiGetAns struct {
	RqData []map[string]interface{} `json:"rq_data"`
	Items  []UsersGetAns            `json:"items"`
}

// UsersGetSubscriptionsAns - объект списка подписок человека
type UsersGetSubscriptionsAns struct {
	Count  int            `json:"count"`
	Items  []GroupsGetAns `json:"items"`
	Users  CountInt       `json:"users"`
	Groups CountInt       `json:"groups"`
}

// CountInt - объект из списка идентификаторов и их кол-ва
type CountInt struct {
	Count int   `json:"count"`
	Items []int `json:"items"`
}

// MultiUsersGetSubscriptionsAns - объект списка подписок человеков
type MultiUsersGetSubscriptionsAns struct {
	RqData []map[string]interface{}   `json:"rq_data"`
	Items  []UsersGetSubscriptionsAns `json:"items"`
}

/*
	Friends
*/

// ScriptMultiFriendsGetAns - Объект списка друзей для нескольких человек
type ScriptMultiFriendsGetAns struct {
	RqData []map[string]interface{}    `json:"rq_data"`
	Items  []ScriptGroupsGetMembersAns `json:"items"`
}

/*
	Groups
*/

// GroupsGetAns - объект ответа при запросе групп
type GroupsGetAns struct {
	ID           int             `json:"id"`
	Name         string          `json:"name"`
	ScreenName   string          `json:"screen_name"`
	IsClosed     int             `json:"is_closed"`
	Deactivated  string          `json:"deactivated"`
	IsAdmin      int             `json:"is_admin"`
	AdminLevel   int             `json:"admin_level"`
	IsMember     int             `json:"is_member"`
	InvitedBy    int             `json:"invited_by"`
	Type         string          `json:"type"`
	Photo50      string          `json:"photo_50"`
	Photo100     string          `json:"photo_100"`
	Photo200     string          `json:"photo_200"`
	AgeLimits    int             `json:"age_limits "`
	Description  string          `json:"description"`
	MembersCount int             `json:"members_count"`
	Verified     int             `json:"verified"`
	Contacts     []GroupContacts `json:"contacts"`
}

// GroupContacts - объект контакта группы
type GroupContacts struct {
	UserID int    `json:"user_id"`
	Desc   string `json:"desc"`
	Phone  string `json:"phone"`
	Email  string `json:"email"`
}

// GroupsGetMembersAns - объект ответа при запросе подписчиков групп
type GroupsGetMembersAns struct {
	Count int             `json:"count"`
	Items json.RawMessage `json:"items"`
}

// ScriptGroupsGetMembersAns - объект ответа при подписчиков (execute)
type ScriptGroupsGetMembersAns struct {
	Count  int   `json:"count"`
	Offset int   `json:"offset"`
	Items  []int `json:"items"`
}

// GroupsGetTokenPermissionsAns - объект прав доступа сообщества
type GroupsGetTokenPermissionsAns struct {
	Mask     int                                    `json:"mask"`
	Settings []GroupsGetTokenPermissionsAnsSettings `json:"settings"`
}

// GroupsGetTokenPermissionsAnsSettings - подробные права сообщества
type GroupsGetTokenPermissionsAnsSettings struct {
	Setting int    `json:"setting"`
	Name    string `json:"name"`
}

/*
	Stats
*/

// StatsGetAns - объект ответа при запросе статистики группы
type StatsGetAns struct {
	GroupID          int             `json:"group_id"`
	Day              string          `json:"day"`
	Views            int             `json:"views"`
	Visitors         int             `json:"visitors"`
	Reach            int             `json:"reach"`
	ReachSubscribers int             `json:"reach_subscribers"`
	Subscribed       int             `json:"subscribed"`
	Unsubscribed     int             `json:"unsubscribed"`
	Sex              []StatsGetValue `json:"sex"`
	Age              []StatsGetValue `json:"age"`
	SexAge           []StatsGetValue `json:"sex_age"`
	Cities           []StatsGetValue `json:"cities"`
	Countries        []StatsGetValue `json:"countries"`
}

// StatsGetValue - объект статистики
type StatsGetValue struct {
	Visitors int         `json:"visitors"`
	Value    interface{} `json:"value"`
	Name     string      `json:"name"`
}

/*
	Wall
*/

// WallGetAns - объект ответа при запросе постов
type WallGetAns struct {
	Count int              `json:"count"`
	Items []WallGetByIDAns `json:"items"`
}

// MultiWallGetAns - объект списка постов нескольких сообществ или людей
type MultiWallGetAns struct {
	RqData []map[string]interface{} `json:"rq_data"`
	Items  []WallGetAns             `json:"items"`
}

// WallGetByIDAns - обект постов
type WallGetByIDAns struct {
	ID           int              `json:"id"`
	OwnerID      int              `json:"owner_id"`
	FromID       int              `json:"from_id"`
	CreatedBy    int              `json:"created_by"`
	Date         int              `json:"date"`
	Text         string           `json:"text"`
	ReplyOwnerID int              `json:"reply_owner_id"`
	ReplyPostID  int              `json:"reply_post_id"`
	FriendsOnly  int              `json:"friends_only"`
	Comments     CommentData      `json:"comments"`
	Likes        LikeData         `json:"likes"`
	Reposts      LikeData         `json:"reposts"`
	Views        LikeData         `json:"views"`
	PostType     string           `json:"post_type"`
	Attachments  []Attachments    `json:"attachments"`
	SignerID     int              `json:"signer_id"`
	CopyHistory  []WallGetByIDAns `json:"copy_history"`
	IsPinned     int              `json:"is_pinned"`
	MarkedAsAds  int              `json:"marked_as_ads"`
	PostponedID  int              `json:"postponed_id"`
}

// Attachments - объект аттача
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

// AttachmentsPrettyCards - объект карточек аттача
type AttachmentsPrettyCards struct {
	Cards []AttachmentsPrettyCardsCards `json:"cards"`
}

// AttachmentsPrettyCardsCards - объект карточек аттача
type AttachmentsPrettyCardsCards struct {
	CardID  string `json:"card_id"`
	LinkURL string `json:"link_url"`
	Title   string `json:"title"`
}

// GetPrettyCards - Преобрахуем данные карточек в объекты
func (a *Attachments) GetPrettyCards() (t AttachmentsPrettyCards) {
	err := json.Unmarshal(*a.PrettyCards, &t)
	if err != nil {
		log.Println("[error]", err)
		return
	}

	return
}

// CommentData - объект комментариев
type CommentData struct {
	Count         int  `json:"count"`
	CanPost       int  `json:"can_post"`
	GroupsCanPost bool `json:"groups_can_post"`
}

// LikeData - объект лайков
type LikeData struct {
	Count int `json:"count"`
}

// WallGetCommentsAns - объект комментариев
type WallGetCommentsAns struct {
	Count  int                   `json:"count"`
	Offset int                   `json:"offset"`
	Items  []WallGetCommentsItem `json:"items"`
}

// MultiWallGetCommentsAns - объект комментариев для выборки из нескольких сообществ
type MultiWallGetCommentsAns struct {
	RqData []map[string]interface{} `json:"rq_data"`
	Items  []WallGetCommentsAns     `json:"items"`
}

// WallGetCommentsItem - объект комментария
type WallGetCommentsItem struct {
	ID             int           `json:"id"`
	FromID         int           `json:"from_id"`
	Date           int           `json:"date"`
	Text           string        `json:"text"`
	ReplyToUser    int           `json:"reply_to_user"`
	ReplyToComment int           `json:"reply_to_comment"`
	Attachments    []Attachments `json:"attachments"`
	Likes          LikeData      `json:"likes"`
	PhotoID        int           `json:"photo_id"`
	PhotoOwnerID   int           `json:"photo_owner_id"`
	VideoID        int           `json:"video_id"`
	VideoOwnerID   int           `json:"video_owner_id"`
	PostID         int           `json:"post_id"`
	PostOwnerID    int           `json:"post_owner_id"`
	MarketOwnerID  int           `json:"market_owner_id"`
	ItemID         int           `json:"item_id"`
}

/*
	Likes
*/

// LikesGetListAns - объект лайков
type LikesGetListAns struct {
	Count  int   `json:"count"`
	Offset int   `json:"offset"`
	Items  []int `json:"items"`
}

// MultiLikesGetListAns - объект лайков для нескольких объектов
type MultiLikesGetListAns struct {
	RqData []map[string]interface{} `json:"rq_data"`
	Items  []LikesGetListAns        `json:"items"`
}

/*
	Utils
*/

// UtilsGetShortLinkAns - объект ответа при запросе короткой ссылки
type UtilsGetShortLinkAns struct {
	ShortURL  string `json:"short_url"`
	URL       string `json:"url"`
	Key       string `json:"key"`
	AccessKey string `json:"access_key"`
}

// UtilsGetLinkStatsAns - объект ответа при запросе статистики короткой ссылки
type UtilsGetLinkStatsAns struct {
	Key   string                   `json:"key"`
	Stats []UtilsGetLinkStatsStats `json:"stats"`
}

// UtilsGetLinkStatsStats - объект статистики короткой ссылки
type UtilsGetLinkStatsStats struct {
	Timestamp int         `json:"timestamp"`
	Views     int         `json:"views"`
	SexAge    []SexAge    `json:"sex_age"`
	Countries []Countries `json:"countries"`
	Cities    []Cities    `json:"cities"`
}

// UtilsResolveScreenNameAns - объект ответа при запросе резольвинка короткого имени
type UtilsResolveScreenNameAns struct {
	Type     string `json:"type"`
	ObjectID int    `json:"object_id"`
}

/*
	Board
*/

// BoardGetTopicsAns - объект списка обсуждений
type BoardGetTopicsAns struct {
	Count  int                  `json:"count"`
	Offset int                  `json:"offset"`
	Items  []BoardGetTopicsItem `json:"items"`
}

// MultiBoardGetTopicsAns - объект списка обсуждений для нескольких групп
type MultiBoardGetTopicsAns struct {
	Items  []BoardGetTopicsAns      `json:"items"`
	RqData []map[string]interface{} `json:"rq_data"`
}

// BoardGetTopicsItem - объект обсуждения
type BoardGetTopicsItem struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Created      int    `json:"created"`
	CreatedBy    int    `json:"created_by"`
	Updated      int    `json:"updated"`
	UpdatedBy    int    `json:"updated_by"`
	IsClosed     int    `json:"is_closed"`
	IsFixed      int    `json:"is_fixed"`
	Comments     int    `json:"comments"`
	FirstComment int    `json:"first_comment"`
	LastComment  int    `json:"last_comment"`
	TopicID      int    `json:"topic_id"`
	TopicOwnerID int    `json:"topic_owner_id"`
}

// BoardGetCommentsAns - объект списка комментариев обсуждения
type BoardGetCommentsAns struct {
	Count  int                   `json:"count"`
	Offset int                   `json:"offset"`
	Items  []WallGetCommentsItem `json:"items"`
}

// MultiBoardGetCommentsAns - объект списка комментариев нескольких обсуждения
type MultiBoardGetCommentsAns struct {
	RqData []map[string]interface{} `json:"rq_data"`
	Items  []BoardGetCommentsAns    `json:"items"`
}

/*
	Photos
*/

// PhotosGetAlbumsAns - объект списка альбомов
type PhotosGetAlbumsAns struct {
	Count int                   `json:"count"`
	Items []PhotosGetAlbumsItem `json:"items"`
}

// MultiPhotosGetAlbumsAns - объект списка альбомов
type MultiPhotosGetAlbumsAns struct {
	RqData []map[string]interface{} `json:"rq_data"`
	Items  []PhotosGetAlbumsAns     `json:"items"`
}

// PhotosGetAlbumsItem - объект альбома
type PhotosGetAlbumsItem struct {
	ID                 int    `json:"id"`
	ThumbID            int    `json:"thumb_id"`
	OwnerID            int    `json:"owner_id"`
	Title              string `json:"title"`
	Description        string `json:"description"`
	Created            int    `json:"created"`
	Updated            int    `json:"updated"`
	Size               int    `json:"size"`
	CanUpload          int    `json:"can_upload"`
	UploadByAdminsOnly int    `json:"UploadByAdminsOnly"`
	CommentsDisabled   int    `json:"comments_disabled"`
}

// PhotosGetAns - объект списка фотографий
type PhotosGetAns struct {
	Count  int             `json:"count"`
	Offset int             `json:"offset"`
	Items  []PhotosGetItem `json:"items"`
}

// MultiPhotosGetAns - объект списка фотографий из разных альбомов
type MultiPhotosGetAns struct {
	RqData []map[string]interface{} `json:"rq_data"`
	Items  []PhotosGetAns           `json:"items"`
}

// PhotosGetItem - объект фотографии
type PhotosGetItem struct {
	ID        int         `json:"id"`
	AlbumID   int         `json:"album_id"`
	OwnerID   int         `json:"owner_id"`
	UserID    int         `json:"user_id"`
	Photo75   string      `json:"photo_75"`
	Photo130  string      `json:"photo_130"`
	Photo604  string      `json:"photo_604"`
	Photo807  string      `json:"photo_807"`
	Photo1280 string      `json:"photo_1280"`
	Photo2560 string      `json:"photo_2560"`
	Text      string      `json:"text"`
	Date      int         `json:"date"`
	Width     int         `json:"width"`
	Height    int         `json:"height"`
	PostID    int         `json:"post_id"`
	Likes     LikeData    `json:"likes"`
	Reposts   LikeData    `json:"reposts"`
	Comments  CommentData `json:"comments"`
}

// PhotosGetCommentsAns - объект списка комментариев фото
type PhotosGetCommentsAns struct {
	Count  int                   `json:"count"`
	Offset int                   `json:"offset"`
	Items  []WallGetCommentsItem `json:"items"`
}

// MultiPhotosGetCommentsAns - объект списка комментариев разных фото
type MultiPhotosGetCommentsAns struct {
	RqData []map[string]interface{} `json:"rq_data"`
	Items  []PhotosGetCommentsAns   `json:"items"`
}

/*
	Video
*/

// VideoGetAns - объект списка видео
type VideoGetAns struct {
	Count  int            `json:"count"`
	Offset int            `json:"offset"`
	Items  []VideoGetItem `json:"items"`
}

// MultiVideoGetAns - объект списка видео нескольких сообществ
type MultiVideoGetAns struct {
	Items  []VideoGetAns            `json:"items"`
	RqData []map[string]interface{} `json:"rq_data"`
}

// VideoGetItem - объект видео
type VideoGetItem struct {
	ID         int      `json:"id"`
	OwnerID    int      `json:"owner_id"`
	Title      string   `json:"title"`
	Duration   int      `json:"duration"`
	Date       int      `json:"date"`
	Comments   int      `json:"comments"`
	Views      int      `json:"views"`
	Likes      LikeData `json:"likes"`
	Reposts    LikeData `json:"reposts"`
	Platform   string   `json:"platform"`
	Player     string   `json:"player"`
	AddingDate int      `json:"adding_date"`
}

// VideoGetCommentsAns - объект списка комментариев
type VideoGetCommentsAns struct {
	Count  int                   `json:"count"`
	Offset int                   `json:"offset"`
	Items  []WallGetCommentsItem `json:"items"`
}

// MultiVideoGetCommentsAns - объект списка комментариев нескольких видео
type MultiVideoGetCommentsAns struct {
	Items  []VideoGetCommentsAns    `json:"items"`
	RqData []map[string]interface{} `json:"rq_data"`
}

/*
	Message
*/

// MessagesGetAns - объект сообщений
type MessagesGetAns struct {
	ID          int               `json:"id"`
	UserID      int               `json:"user_id"`
	FromID      int               `json:"from_id"`
	Date        int               `json:"date"`
	ReadState   int               `json:"read_state"`
	Out         int               `json:"out"`
	Title       string            `json:"title"`
	Body        string            `json:"body"`
	Geo         MessagesGetAnsGeo `json:"geo"`
	Attachments []Attachments     `json:"attachments"`
	Emoji       int               `json:"emoji"`
	Important   int               `json:"important"`
	Deleted     int               `json:"deleted"`
	RandomID    int               `json:"random_id"`
}

// MessagesGetAnsGeo - объект места в сообщении
type MessagesGetAnsGeo struct {
	Type        string                 `json:"type"`
	Coordinates string                 `json:"coordinates"`
	Place       MessagesGetAnsGeoPlace `json:"place"`
}

// MessagesGetAnsGeoPlace - объект описание места
type MessagesGetAnsGeoPlace struct {
	ID        int     `json:"id"`
	Title     string  `json:"title"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Created   int     `json:"created"`
	Icon      string  `json:"icon"`
	Country   string  `json:"country"`
	City      string  `json:"city"`
}

// MessagesIsMessagesFromGroupAllowedAns - объект проверки разрешена ли отправка сообщений
type MessagesIsMessagesFromGroupAllowedAns struct {
	IsAllowed int `json:"is_allowed"`
}

/*
	Market
*/

// ScriptMultiMarketGetAns - Список объектов товаров
type ScriptMultiMarketGetAns struct {
	Items  []MarketGetAns           `json:"items"`
	RqData []map[string]interface{} `json:"rq_data"`
}

// MarketGetAns - объетк списка товаров
type MarketGetAns struct {
	Count  int                `json:"count"`
	Offset int                `json:"offset"`
	Items  []MarketGetByIDAns `json:"items"`
}

// MarketGetByIDAns - объект товара
type MarketGetByIDAns struct {
	ID           int             `json:"id"`
	OwnerID      int             `json:"owner_id"`
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	Price        MarketPrice     `json:"price"`
	Category     MarketCategory  `json:"category"`
	Date         int             `json:"date"`
	ThumbPhoto   string          `json:"thumb_photo"`
	Availability int             `json:"availability"`
	AlbumsIDs    []int           `json:"albums_ids"`
	Photos       []PhotosGetItem `json:"photos"`
	CanComment   int             `json:"can_comment"`
	CanRepost    int             `json:"can_repost"`
	Likes        LikeData        `json:"likes"`
	ViewsCount   int             `json:"views_count"`
}

// MarketCategory - объект категории
type MarketCategory struct {
	ID      int                   `json:"id"`
	Name    string                `json:"name"`
	Section MarketCategorySection `json:"section"`
}

// MarketCategorySection - объект секции категории товара
type MarketCategorySection struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// MarketPrice - объект цены товара
type MarketPrice struct {
	Amount   string              `json:"amount"`
	Currency MarketPriceCurrency `json:"currency"`
	Text     string              `json:"text"`
}

// MarketPriceCurrency - объект валюты товара
type MarketPriceCurrency struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// MultiMarketGetCommentsAns - объект списка коментов товара
type MultiMarketGetCommentsAns struct {
	Items  []WallGetCommentsAns     `json:"items"`
	RqData []map[string]interface{} `json:"rq_data"`
}

/*
	Ads
*/

// AdsGetAccountsAns - объект информации об аккаунте
type AdsGetAccountsAns struct {
	AccountID     int    `json:"account_id"`
	AccountType   string `json:"account_type"`
	AccountStatus int    `json:"account_status"`
	AccountName   string `json:"account_name"`
	AccessRole    string `json:"access_role"`
}

// AdsСreateTargetGroupAns - Ответ на созданную аудиторию
type AdsСreateTargetGroupAns struct {
	ID int `json:"id"`
}

// AdsGetTargetingStatsAns - объект информации о размере аудитории
type AdsGetTargetingStatsAns struct {
	AudienceCount  int    `json:"audience_count"`
	RecommendedCPC string `json:"recommended_cpc"`
	RecommendedCPM string `json:"recommended_cpm"`
}

// AdsGetTargetingStatsCriteria - объект критериев настроек таргета
type AdsGetTargetingStatsCriteria struct {
	Sex                  int    `json:"sex"`
	AgeFrom              int    `json:"age_from"`
	AgeTo                int    `json:"age_to"`
	Birthday             int    `json:"birthday"`
	Country              int    `json:"country"`
	Cities               string `json:"cities"`
	CitiesNot            string `json:"cities_not"`
	GeoNear              string `json:"geo_near"`
	GeoPointType         string `json:"geo_point_type"`
	Statuses             string `json:"statuses"`
	Groups               string `json:"groups"`
	GroupsNot            string `json:"groups_not"`
	Apps                 string `json:"apps"`
	AppsNot              string `json:"apps_not"`
	Districts            string `json:"districts"`
	Stations             string `json:"stations"`
	Streets              string `json:"streets"`
	Schools              string `json:"schools"`
	Positions            string `json:"positions"`
	Religions            string `json:"religions"`
	InterestCategories   string `json:"interest_categories"`
	Interests            string `json:"interests"`
	UserDevices          string `json:"user_devices"`
	UserOS               string `json:"user_os"`
	UserBrowsers         string `json:"user_browsers"`
	RetargetingGroups    string `json:"retargeting_groups"`
	RetargetingGroupsNot string `json:"retargeting_groups_not"`
	Paying               int    `json:"paying"`
	Travellers           int    `json:"travellers"`
	SchoolFrom           int    `json:"school_from"`
	SchoolTo             int    `json:"school_to"`
	UniFrom              int    `json:"uni_from"`
	UniTo                int    `json:"uni_to"`
}

// AdsGetSuggestionsAns - объект подсказки
type AdsGetSuggestionsAns struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Parent string `json:"parent"`
}

// AdsGetCampaignsAns - объект ответа при запросе кампаний
type AdsGetCampaignsAns struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

// AdsGetAdsLayoutAns - объект ответа при запросе вида объявления
type AdsGetAdsLayoutAns struct {
	ID         string `json:"id"`
	CampaignID int    `json:"campaign_id"`
	Title      string `json:"title"`
	LinkURL    string `json:"link_url"`
}

// AdsGetStatisticsAns - объект ответа при запросе статистики
type AdsGetStatisticsAns struct {
	ID       int                        `json:"id"`
	Type     string                     `json:"type"`
	StatsBug []json.RawMessage          `json:"stats"`
	Stats    []AdsGetStatisticsAnsStats `json:"-"`
}

// AdsGetStatisticsAnsStats - объект статистики
type AdsGetStatisticsAnsStats struct {
	Day         string `json:"day"`
	Spent       string `json:"spent"`
	Impressions int    `json:"impressions"`
	Clicks      int    `json:"clicks"`
	Reach       int    `json:"reach"`
}

// AdsGetStatisticsAnsStatsBug - объект ответа при запросе статистики (если VK криво типы переменных сформировал)
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

// SexAge - объект пола/возраста
type SexAge struct {
	AgeRange string `json:"age_range"`
	Female   int    `json:"female"`
	Male     int    `json:"male"`
}

// Countries - объект статистики по странам
type Countries struct {
	CountryID int `json:"country_id"`
	Views     int `json:"views"`
}

// Cities - объект статистики по городам
type Cities struct {
	CityID int `json:"city_id"`
	Views  int `json:"views"`
}
