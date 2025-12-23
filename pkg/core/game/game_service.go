package game

import (
	"time"
)

type GameService interface {
	Init(params InitParameters) (InitServerUrls, *ServiceError)
	GetServerTime(params GetServerTimeParams) (BaseTimestamp, *ServiceError)
	UserLogin(params UserLoginParams) (*UserLoginResult, *ServiceError)
	GetUpdateTime(params GetUpdateTimeParams) UpdateTime
	DefaultSettingList(params DefaultSettingListParams) DefaultSettingList
}

type GameServiceImpl struct{}

func (service *GameServiceImpl) Init(params InitParameters) (InitServerUrls, *ServiceError) {
	return InitServerUrls{
		UnknownInt:     253,
		GameServerUrl:  "https://game.gtgl.pmang.cloud",
		AssetServerUrl: "https://dl.gtgl.pmang.cloud",
	}, nil
}

func (service *GameServiceImpl) GetServerTime(params GetServerTimeParams) (BaseTimestamp, *ServiceError) {
	return getBaseTimeStamp(time.Now(), params.TimeZone), nil
}

func (service *GameServiceImpl) UserLogin(params UserLoginParams) (*UserLoginResult, *ServiceError) {
	// TODO internal logic
	return nil, NewServiceError(
		998, "Sorry. Not a registered user.",
	)
}

func (service *GameServiceImpl) GetUpdateTime(params GetUpdateTimeParams) UpdateTime {
	// TODO internal logic
	return UpdateTime{
		UnixSeconds: 1761876067,
	}
}

func (service *GameServiceImpl) DefaultSettingList(params DefaultSettingListParams) DefaultSettingList {
	return DefaultSettingList{
		UnknownYesNo: "Y",
		Settings: []Pair{
			{
				Key:   "shop_call_wait",
				Value: "2",
			},
			{
				Key:   "ad_cancel_delay",
				Value: "5",
			},
			{
				Key:   "force_menu_restric",
				Value: "true",
			},
			{
				Key:   "max_bundle_texutre_loading",
				Value: "1",
			},
			{
				Key:   "weather_snow",
				Value: "true",
			},
			{
				Key:   "check_time_cheat",
				Value: "false",
			},
			{
				Key:   "ap_max",
				Value: "60",
			},
			{
				Key:   "ap_time",
				Value: "300",
			},
			{
				Key:   "ap_use",
				Value: "5",
			},
			{
				Key:   "character_exp",
				Value: "45",
			},
			{
				Key:   "ap_shop_id",
				Value: "31",
			},
			{
				Key:   "ap_ad_list_id",
				Value: "2",
			},
			{
				Key:   "ch_third_unlock_follower_id",
				Value: "5",
			},
			{
				Key:   "ch_third_unlock_follower_level",
				Value: "1",
			},
			{
				Key:   "ad_follower_profile_exp_persent",
				Value: "0.1",
			},
			{
				Key:   "follower_profile_bonus_percent",
				Value: "1",
			},
			{
				Key:   "spine_load_max_count",
				Value: "0",
			},
			{
				Key:   "log_flag",
				Value: "true",
			},
			{
				Key:   "weather_cherry_blossom",
				Value: "false",
			},
			{
				Key:   "active_goods_ticket_buff",
				Value: "true",
			},
			{
				Key:   "active_fan_art_costume",
				Value: "true",
			},
			{
				Key:   "pass_goods_sale_start_time",
				Value: "2022-02-19 00:00",
			},
			{
				Key:   "pass_goods_sale_end_time",
				Value: "2022-02-27 23:59",
			},
			{
				Key:   "attendacne_buff_value",
				Value: "0.01",
			},
			{
				Key:   "attendacne_buff_limit_day",
				Value: "10000",
			},
			{
				Key:   "log_flag_unfinished_purchase",
				Value: "true",
			},
			{
				Key:   "jp_laws_flag",
				Value: "true",
			},
		},
	}
}

type DefaultSettingList struct {
	UnknownYesNo string `thrift:",1"`
	Settings     []Pair `thrift:",2"`
}

type Pair struct {
	Key   string `thrift:",1"`
	Value string `thrift:",2"`
}

type DefaultSettingListParams struct {
	DeviceId   string `thrift:",1"`
	UnknownInt int32  `thrift:",2"`
}

type ServiceError struct {
	ErrorCode    int    `thrift:",1"`
	ErrorMessage string `thrift:",2"`
}

func NewServiceError(code int, message string) *ServiceError {
	return &ServiceError{
		ErrorCode:    code,
		ErrorMessage: message,
	}
}

type GetUpdateTimeParams struct {
	DeviceId string
}

type UpdateTime struct {
	UnixSeconds int64 `thrift:",1"`
}

type UserLoginResult struct {
	ResultCode    int    `thrift:",1"`
	ResultMessage string `thrift:",2"`
}

type UserLoginParams struct {
	UnknownInt    int32
	UnknownString string
	UserId        string
	DeviceId      string
}

type InitParameters struct {
	Version     int32
	Environment string
	UnknownFlag int16
	DeviceId    string
}

type InitServerUrls struct {
	UnknownInt     int    `thrift:",1"`
	GameServerUrl  string `thrift:",2"`
	AssetServerUrl string `thrift:",3"`
}

type GetServerTimeParams struct {
	DeviceId string
	TimeZone string
}
