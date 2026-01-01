package game

import (
	"encoding/json"
	"time"

	embeds "github.com/KhoalaS/guitar-girl-offline"
	"github.com/KhoalaS/guitar-girl-offline/pkg/model/common_model"
	"github.com/KhoalaS/guitar-girl-offline/pkg/model/main_model"
	"github.com/KhoalaS/guitar-girl-offline/pkg/model/user_model"
)

type GameService interface {
	Init(params InitParameters) (main_model.InitRetDataInfo, common_model.ErrorRetCode)
	GetServerTime(params GetServerTimeParams) (BaseTimestamp, common_model.ErrorRetCode)
	UserLogin(params UserLoginParams) (user_model.UserLoginRetDataInfo, common_model.ErrorRetCode)
	GetUpdateTime(params GetUpdateTimeParams) main_model.GetUpdateTimeRetDataInfo
	DefaultSettingList(params DefaultSettingListParams) DefaultSettingList
	GetGameDataList(params GetGameDataListParams) (map[string]any, common_model.ErrorRetCode)
}

type GameServiceImpl struct {
	UserRepository
}

func (service *GameServiceImpl) Init(params InitParameters) (main_model.InitRetDataInfo, common_model.ErrorRetCode) {
	return main_model.InitRetDataInfo{
		Idx:      253,
		Game_url: "https://game.gtgl.pmang.cloud",
		Cdn_url:  "https://dl.gtgl.pmang.cloud",
	}, common_model.ErrorRetCode{}
}

func (service *GameServiceImpl) GetServerTime(params GetServerTimeParams) (BaseTimestamp, common_model.ErrorRetCode) {
	return getBaseTimeStamp(time.Now(), params.TimeZone), common_model.ErrorRetCode{}
}

func (service *GameServiceImpl) UserLogin(params UserLoginParams) (user_model.UserLoginRetDataInfo, common_model.ErrorRetCode) {
	// TODO internal logic
	userData, err := service.UserRepository.GetUser(params.UserId)
	if err != nil {
		return user_model.UserLoginRetDataInfo{}, common_model.ErrorRetCode{
			Code:   998,
			Errmsg: "Sorry. Not a registered user.",
		}
	}

	return user_model.UserLoginRetDataInfo{
		User: userData,
	}, common_model.ErrorRetCode{}

}

func (service *GameServiceImpl) GetUpdateTime(params GetUpdateTimeParams) main_model.GetUpdateTimeRetDataInfo {
	// TODO internal logic
	return main_model.GetUpdateTimeRetDataInfo{
		Upd_time: 1761876067,
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

func (service *GameServiceImpl) GetGameDataList(params GetGameDataListParams) (map[string]any, common_model.ErrorRetCode) {
	var data map[string]any

	err := json.Unmarshal(embeds.GameData, &data)
	if err != nil {
		return nil, common_model.ErrorRetCode{Code: 1, Errmsg: err.Error()}
	}

	return data, common_model.ErrorRetCode{}
}

type GetGameDataListParams struct {
	Type     string `thrift:",1"`
	DeviceId string `thrift:",2"`
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
	UnknownInt int16  `thrift:",2"`
}

type GetUpdateTimeParams struct {
	DeviceId string
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
