package core

import (
	"fmt"
	"time"
)

// Most likely meaning of abbreviations for internal fields
// _dt -> Date
// _cd -> Code ?
// fcm -> Firebase Cloud Messaging
// _yn -> Yes/No
// reg_ -> Registration

type AuthApi interface {
	Analytics() BaseResponse
	Register(fcmToken string, deviceId string, puid string, localeCd string) BaseResponse
	Login(osVer string, releaseYmd string, adNight YN, deviceCd string, privacy YN, locale string, mobSvc YN, version int, appKey string, ad YN, threadName int, appVer string, udid string, appId int, appSecret string, localCd string, jailbreak YN) BaseResponse
	ReferrerSave(referrer string, adId string, appKey string, installBeginTime int64, referrerClickTime int, deviceCd string, appId int, appSecret string, localCd string) BaseResponse
	GeoipCountryCodeInfo(deviceCd string, appId int, localCd string) BaseResponse
	Eula(resultScheme int, platformCd string, releaseYmd string, nationCd string) BaseResponse
}

type BaseResponse struct {
	value       any
	result_code string
	result_msg  string
}

type AuthApiImpl struct {
}

func (api AuthApiImpl) Analytics() BaseResponse {
	return BaseResponse{
		value:       true,
		result_code: "000",
		result_msg:  API_OK,
	}
}

func (api AuthApiImpl) Register(fcmToken string, deviceId string, puid string, localeCd string) BaseResponse {
	// TODO What is this "fcmheader" field? Firebase Cloud Messaging?
	// Is it relevant? Is the value the same for all requests or dependant on the parameters?
	return BaseResponse{
		value: RegisterResponseValue{
			Unsubscription: []any{},
			Subscription:   []any{},
			FCMHeader:      "CsYsPQf1UEU:APA91bE1dMmiUqnw5HACZ9GstHG8U_K-5sNxSgWWPbNZBOfP93v2M7PzjATrXetY_vnzRc5aFQAbS2TGpKDniifN5DDcfUlPG4MVWkhSHjUHS_X6ViwpImmU5BteBUZhjBAAAAq1Zi1q",
		},
	}
}

func (api AuthApiImpl) Login(osVer string, releaseYmd string, adNight YN, deviceCd string, privacy YN, locale string, mobSvc YN, version int, appKey string, ad YN, threadName int, appVer string, udid string, appId int, appSecret string, localCd string, jailbreak YN) BaseResponse {
	// we assume that the udid is the parameter which maps onto a memberId
	nowMillis := time.Now().UnixMilli()
	memberId := 429071553
	// this value is generated on the server and different for every session
	// there is no schema that the characters have to follow
	unknownOpt_1 := "618de92ee950717f49afcd82d359bc92b2c34a7f"
	serverLocale := "KR"
	accessToken := fmt.Sprintf("%d|%d|%s|%s|%s|%d", memberId, appId, deviceCd, serverLocale, unknownOpt_1, time.Now().UnixMilli())
	return BaseResponse{
		value: LoginResponseValue{
			AccessToken: accessToken,
			Member: LoginResponseValueMember{
				CurrentDate:         nowMillis,
				UpdateDate:          nowMillis,
				StatusCd:            nil,
				MemberId:            memberId,
				Nickname:            "User",
				ProfileImageUrl:     nil,
				Feeling:             nil,
				AdultAuth:           nil,
				AdultAuthDate:       nil,
				RecentLoginDate:     nil,
				RecentAppId:         appId,
				Email:               nil,
				Anonymous:           YES,
				FriendAcceptCd:      nil,
				RegPath:             nil,
				RecentAppTitle:      "Guitar Girl",
				LastMsgDate:         nil,
				NewMsg:              nil,
				ConflictMemberId:    0,
				RegIp:               "127.0.0.1",
				RegNation:           "US",
				IsGuestLogin:        true,
				ProviderDisplayName: "",
				Pushgroup:           nil,
				Locale:              nil,
			},
			ForceReceipt:     nil,
			ConflictMemberId: nil,
			IsGuestLogin:     true,
			OldMemberId:      nil,
			Jailbreak:        NO,
			FCMSendLang:      "EN",
			UnregStatus:      "NO",
			UnregRemainTime:  nil,
			CallTime:         0,
		},
		result_code: "000",
		result_msg:  API_OK,
	}
}

func (api AuthApiImpl) ReferrerSave(referrer string, adId string, appKey string, installBeginTime int64, referrerClickTime int, deviceCd string, appId int, appSecret string, localCd string) BaseResponse {
	return BaseResponse{
		value:       "OK",
		result_code: "000",
		result_msg:  API_OK,
	}
}

func (api AuthApiImpl) GeoipCountryCodeInfo(deviceCd string, appId int, localCd string) BaseResponse {
	return BaseResponse{
		value:       "US",
		result_code: "000",
		result_msg:  API_OK,
	}
}

func (api AuthApiImpl) Eula(resultScheme int, platformCd string, releaseYmd string, nationCd string) BaseResponse {
	return BaseResponse{
		value:       true,
		result_code: "000",
		result_msg:  API_OK,
	}
}

const (
	API_OK string = "API_OK"
)

type RegisterResponseValue struct {
	Unsubscription []any  `json:"unsubscription"`
	FCMHeader      string `json:"fcmheader"`
	Subscription   []any  `json:"subscription"`
}

type LoginResponseValue struct {
	AccessToken      string                   `json:"access_token"`
	Member           LoginResponseValueMember `json:"member"`
	ForceReceipt     any                      `json:"force_receipt"`
	ConflictMemberId any                      `json:"conflict_member_id"`
	IsGuestLogin     bool                     `json:"is_guest_login"`
	OldMemberId      any                      `json:"old_member_id"`
	Jailbreak        YN                       `json:"jailbreak_yn"`
	FCMSendLang      string                   `json:"fcm_send_lang"`
	UnregStatus      string                   `json:"unreg_status"`
	UnregRemainTime  any                      `json:"unreg_remain_time"`
	CallTime         int                      `json:"callTime"`
}

type LoginResponseValueMember struct {
	CurrentDate         int64   `json:"crt_dt"`
	UpdateDate          int64   `json:"upd_dt"`
	StatusCd            any     `json:"status_cd"`
	MemberId            int     `json:"member_id"`
	Nickname            string  `json:"nick_name"`
	ProfileImageUrl     *string `json:"profile_img_url"`
	Feeling             any     `json:"feeling"`
	AdultAuth           *YN     `json:"adult_auth_yn"`
	AdultAuthDate       *int64  `json:"adult_auth_dt"`
	RecentLoginDate     *int64  `json:"recent_login_dt"`
	RecentAppId         int     `json:"recent_app_id"`
	Email               *string `json:"email"`
	Anonymous           YN      `json:"anonymous_yn"`
	FriendAcceptCd      *string `json:"friend_accept_cd"`
	RegPath             any     `json:"reg_path"`
	RecentAppTitle      string  `json:"recent_app_title"`
	LastMsgDate         *int64  `json:"last_msg_dt"`
	NewMsg              *YN     `json:"new_msg_yn"`
	ConflictMemberId    int     `json:"conflict_member_id"`
	RegIp               string  `json:"reg_ip"`
	RegNation           string  `json:"reg_nation"`
	IsGuestLogin        bool    `json:"is_guest_login"`
	ProviderDisplayName string  `json:"provider_display_name"`
	Pushgroup           any     `json:"pushgroup"`
	Locale              any     `json:"locale"`
}

type YN string

const (
	YES YN = "Y"
	NO  YN = "N"
)
