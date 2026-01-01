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
	UserSave(params user_model.UserSaveDataInfo) (user_model.UserSaveRetDataInfo, common_model.ErrorRetCode)
}

type GameServiceImpl struct {
	UserRepository
}

func (service *GameServiceImpl) UserSave(params user_model.UserSaveDataInfo) (user_model.UserSaveRetDataInfo, common_model.ErrorRetCode) {
	return user_model.UserSaveRetDataInfo{
		Status: "Y",
	}, common_model.ErrorRetCode{}
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

	if params.UnknownInt != 0 {
		return user_model.UserLoginRetDataInfo{
			User: user_model.UserData{
				U_seq:                   params.UnknownInt,
				U_id:                    params.UnknownString,
				U_name:                  "User",
				U_nick:                  "User",
				U_cp:                    999,
				U_candy:                 9999,
				U_like:                  0,
				U_fans:                  0,
				U_fans_grade:            0,
				U_selected_costume_id:   0,
				U_selected_music_id:     0,
				U_retain_continuous_tap: 0,
				U_join_type:             1,
				U_last_login:            "2025-12-31 23:41:08",
				U_last_communication:    "2025-12-31 23:41:08",
				U_save_date:             "1767192069",
				U_create_time:           "1765733352",
				U_tutorial_step:         0,
				U_review_popup:          "N",
				Device_uuid:             "1856b5b57bd634cce52c7451917e15d3",
				U_samseck_step:          3,
				U_free_cp:               999,
				U_charge_cp:             0,
			},
			Area_data: map[int32]user_model.UserAreaData{
				1: user_model.UserAreaData{
					U_area_num:          1,
					D_Like:              0,
					I_UserFanCount:      0,
					I_UserFanGrade:      7668,
					I_SelectedCostumeId: 2,
					I_SelectedMusicId:   13,
					I_SelectedGuitarId:  3,
					D_Candy:             10,
					S_TutorialList:      "",
					S_Gp1:               "16981935241350",
					S_Gp2:               "0",
				},
				2: user_model.UserAreaData{
					U_area_num:          2,
					D_Like:              0,
					I_UserFanCount:      0,
					I_UserFanGrade:      0,
					I_SelectedCostumeId: 1,
					I_SelectedMusicId:   201,
					I_SelectedGuitarId:  201,
					D_Candy:             201,
					S_TutorialList:      "",
					S_Gp1:               "0",
					S_Gp2:               "0",
				},
			},
			User_contents: user_model.UserContentsData{
				User_achievement: []user_model.UserAchievement{
					user_model.UserAchievement{
						I_id:       1,
						I_Level:    2,
						D_Quantity: 0,
						S_Quantity: "227",
					},
					user_model.UserAchievement{
						I_id:       2,
						I_Level:    1,
						D_Quantity: 0,
						S_Quantity: "187",
					},
					user_model.UserAchievement{
						I_id:       3,
						I_Level:    1,
						D_Quantity: 0,
						S_Quantity: "16981935241350",
					},
					user_model.UserAchievement{
						I_id:       4,
						I_Level:    2,
						D_Quantity: 0,
						S_Quantity: "7668",
					},
					user_model.UserAchievement{
						I_id:       5,
						I_Level:    2,
						D_Quantity: 0,
						S_Quantity: "8",
					},
					user_model.UserAchievement{
						I_id:       6,
						I_Level:    1,
						D_Quantity: 0,
						S_Quantity: "645",
					},
					user_model.UserAchievement{
						I_id:       7,
						I_Level:    1,
						D_Quantity: 0,
						S_Quantity: "2",
					},
					user_model.UserAchievement{
						I_id:       8,
						I_Level:    1,
						D_Quantity: 0,
						S_Quantity: "1",
					},
					user_model.UserAchievement{
						I_id:       9,
						I_Level:    1,
						D_Quantity: 0,
						S_Quantity: "1",
					},
					user_model.UserAchievement{
						I_id:       10,
						I_Level:    2,
						D_Quantity: 0,
						S_Quantity: "6",
					},
					user_model.UserAchievement{
						I_id:       201,
						I_Level:    1,
						D_Quantity: 0,
						S_Quantity: "0",
					},
					user_model.UserAchievement{
						I_id:       202,
						I_Level:    1,
						D_Quantity: 0,
						S_Quantity: "0",
					},
					user_model.UserAchievement{
						I_id:       203,
						I_Level:    1,
						D_Quantity: 0,
						S_Quantity: "0",
					},
					user_model.UserAchievement{
						I_id:       204,
						I_Level:    1,
						D_Quantity: 0,
						S_Quantity: "0",
					},
					user_model.UserAchievement{
						I_id:       205,
						I_Level:    1,
						D_Quantity: 0,
						S_Quantity: "0",
					},
					user_model.UserAchievement{
						I_id:       206,
						I_Level:    1,
						D_Quantity: 0,
						S_Quantity: "0",
					},
					user_model.UserAchievement{
						I_id:       207,
						I_Level:    1,
						D_Quantity: 0,
						S_Quantity: "0",
					},
				},
				User_buff: nil,
				User_candy_shop: []user_model.UserCandyShop{
					user_model.UserCandyShop{
						I_id:              1,
						I_CurrentBuyCount: 1,
						I_TotalBuyCount:   1,
						L_LastBuyTick:     1767022695,
						Upd_day:           20251230,
					},
					user_model.UserCandyShop{
						I_id:              2,
						I_CurrentBuyCount: 1,
						I_TotalBuyCount:   1,
						L_LastBuyTick:     1767114498,
						Upd_day:           20251231,
					},
				},
				User_character: []user_model.UserCharacter{
					user_model.UserCharacter{
						I_id:         1,
						I_Level:      227,
						I_BonusLevel: 9,
					},
					user_model.UserCharacter{
						I_id:         2,
						I_Level:      1,
						I_BonusLevel: 0,
					},
				},
				User_costume: []user_model.UserCostume{
					user_model.UserCostume{
						I_id:         1,
						I_Level:      1,
						I_BonusLevel: 0,
					},
					user_model.UserCostume{
						I_id:         2,
						I_Level:      1,
						I_BonusLevel: 0,
					},
					user_model.UserCostume{
						I_id:         3,
						I_Level:      1,
						I_BonusLevel: 0,
					},
					user_model.UserCostume{
						I_id:         9,
						I_Level:      1,
						I_BonusLevel: 0,
					},
					user_model.UserCostume{
						I_id:         13,
						I_Level:      1,
						I_BonusLevel: 0,
					},
					user_model.UserCostume{
						I_id:         14,
						I_Level:      1,
						I_BonusLevel: 0,
					},
				},
				User_daily_mission: []user_model.UserDailyMission{
					user_model.UserDailyMission{
						I_id:       1,
						I_Level:    1,
						D_Quantity: 1,
						Upd_date:   "2025-12-30 00:37:17",
					},
					user_model.UserDailyMission{
						I_id:       2,
						I_Level:    1,
						D_Quantity: 23,
						Upd_date:   "",
					},
					user_model.UserDailyMission{
						I_id:       3,
						I_Level:    1,
						D_Quantity: 0,
						Upd_date:   "",
					},
					user_model.UserDailyMission{
						I_id:       4,
						I_Level:    1,
						D_Quantity: 1,
						Upd_date:   "",
					},
					user_model.UserDailyMission{
						I_id:       5,
						I_Level:    1,
						D_Quantity: 902,
						Upd_date:   "",
					},
					user_model.UserDailyMission{
						I_id:       6,
						I_Level:    1,
						D_Quantity: 0,
						Upd_date:   "",
					},
				},
				User_follower: []user_model.UserFollower{
					user_model.UserFollower{
						I_id:         1,
						I_Level:      125,
						I_BonusLevel: 5,
					},
					user_model.UserFollower{
						I_id:         2,
						I_Level:      48,
						I_BonusLevel: 1,
					},
					user_model.UserFollower{
						I_id:         3,
						I_Level:      13,
						I_BonusLevel: 0,
					},
					user_model.UserFollower{
						I_id:         4,
						I_Level:      1,
						I_BonusLevel: 0,
					},
				},
				User_music: []user_model.UserMusic{
					user_model.UserMusic{
						I_id:                      1,
						I_Level:                   6,
						I_BonusLevel:              0,
						B_EncoreBonusAppear:       0,
						L_EncoreBonusActivateTime: 1767096954,
						I_EncoreBonusFollowerId:   0,
						I_ChThirdActiveTime:       0,
					},
					user_model.UserMusic{
						I_id:                      2,
						I_Level:                   1,
						I_BonusLevel:              0,
						B_EncoreBonusAppear:       0,
						L_EncoreBonusActivateTime: 1767097063,
						I_EncoreBonusFollowerId:   0,
						I_ChThirdActiveTime:       0,
					},
					user_model.UserMusic{
						I_id:                      3,
						I_Level:                   1,
						I_BonusLevel:              0,
						B_EncoreBonusAppear:       0,
						L_EncoreBonusActivateTime: 0,
						I_EncoreBonusFollowerId:   0,
						I_ChThirdActiveTime:       0,
					},
				},
				User_prop: []user_model.UserProp{
					user_model.UserProp{
						I_id:    1,
						I_Level: 1,
					},
					user_model.UserProp{
						I_id:    2,
						I_Level: 1,
					},
				},
				User_unit: []user_model.UserUnit{
					user_model.UserUnit{
						I_id:    1,
						I_Level: 1,
					},
				},
				User_skill: []user_model.UserSkill{
					user_model.UserSkill{
						I_id:               1,
						I_Level:            1,
						B_Activate:         0,
						L_ActivateOnTicks:  0,
						L_ActivateOffTicks: 0,
					},
					user_model.UserSkill{
						I_id:               4,
						I_Level:            1,
						B_Activate:         0,
						L_ActivateOnTicks:  0,
						L_ActivateOffTicks: 0,
					},
				},
				User_shop: []user_model.UserShop{
					user_model.UserShop{
						I_id:           1,
						I_Count:        1,
						I_TotalCount:   1,
						I_PurchaseTime: 1765734809,
						Upd_day:        20251215,
					},
					user_model.UserShop{
						I_id:           13,
						I_Count:        1,
						I_TotalCount:   1,
						I_PurchaseTime: 1767114545,
						Upd_day:        20251231,
					},
				},
				User_messenger: []user_model.UserMessenger{
					user_model.UserMessenger{
						I_MessengerChatRoomId: 1,
						I_LastConfirmIndex:    14,
						S_UnlockGroupList:     "1,3,4",
						L_UpdateTimeTicks:     639013642974540000,
					},
					user_model.UserMessenger{
						I_MessengerChatRoomId: 2,
						I_LastConfirmIndex:    22,
						S_UnlockGroupList:     "2",
						L_UpdateTimeTicks:     639013638752490000,
					},
					user_model.UserMessenger{
						I_MessengerChatRoomId: 3,
						I_LastConfirmIndex:    10,
						S_UnlockGroupList:     "6",
						L_UpdateTimeTicks:     639013641110540000,
					},
				},
				User_guitar: []user_model.UserGuitar{
					user_model.UserGuitar{
						I_id:         1,
						I_Level:      1,
						I_BonusLevel: 0,
					},
					user_model.UserGuitar{
						I_id:         2,
						I_Level:      1,
						I_BonusLevel: 0,
					},
					user_model.UserGuitar{
						I_id:         10,
						I_Level:      1,
						I_BonusLevel: 0,
					},
				},
				User_event_point: []user_model.UserEventPoint{
					user_model.UserEventPoint{
						S_EventType:  "Pass",
						I_DataID:     5,
						I_Point:      3715,
						I_Step:       1,
						I_ADViewTime: 0,
						I_Version:    5,
					},
				},
				User_subscribe_list: nil,
				User_subscribe_pass_reward: []user_model.UserSubscribePassReward{
					user_model.UserSubscribePassReward{
						I_SubscribeID: 5,
						I_Type:        0,
						I_Step:        0,
						I_UpdateTime:  1765734537,
						I_Version:     5,
					},
					user_model.UserSubscribePassReward{
						I_SubscribeID: 5,
						I_Type:        0,
						I_Step:        1,
						I_UpdateTime:  1767114615,
						I_Version:     5,
					},
				},
				User_ticketcollection: nil,
				User_follower_quest:   nil,
				User_follower_profile_reward: []user_model.UserFollowerProfileReward{
					user_model.UserFollowerProfileReward{
						I_id:          1,
						I_RewardLevel: 1,
					},
					user_model.UserFollowerProfileReward{
						I_id:          1,
						I_RewardLevel: 2,
					},
					user_model.UserFollowerProfileReward{
						I_id:          1,
						I_RewardLevel: 3,
					},
					user_model.UserFollowerProfileReward{
						I_id:          2,
						I_RewardLevel: 1,
					},
					user_model.UserFollowerProfileReward{
						I_id:          2,
						I_RewardLevel: 2,
					},
					user_model.UserFollowerProfileReward{
						I_id:          3,
						I_RewardLevel: 1,
					},
					user_model.UserFollowerProfileReward{
						I_id:          10000,
						I_RewardLevel: 1,
					},
				},
				User_follower_profile: []user_model.UserFollowerProfile{
					user_model.UserFollowerProfile{
						I_id:       1,
						I_Level:    3,
						D_Exp:      150,
						I_AddCandy: 0,
					},
					user_model.UserFollowerProfile{
						I_id:       2,
						I_Level:    2,
						D_Exp:      110,
						I_AddCandy: 0,
					},
					user_model.UserFollowerProfile{
						I_id:       3,
						I_Level:    2,
						D_Exp:      70,
						I_AddCandy: 0,
					},
				},
				User_follower_giftitem: []user_model.UserFollowerGiftItem{
					user_model.UserFollowerGiftItem{
						I_id:    1,
						I_Value: 141,
					},
				},
				User_ad_list:       nil,
				User_chthird_stage: nil,
				User_tutorial: []user_model.UserTutorial{
					user_model.UserTutorial{
						I_id: 1,
					},
					user_model.UserTutorial{
						I_id: 2,
					},
					user_model.UserTutorial{
						I_id: 3,
					},
					user_model.UserTutorial{
						I_id: 4,
					},
					user_model.UserTutorial{
						I_id: 5,
					},
					user_model.UserTutorial{
						I_id: 6,
					},
					user_model.UserTutorial{
						I_id: 7,
					},
				},
				User_ad_level: []user_model.UserAdLevel{
					user_model.UserAdLevel{
						I_id:    210010,
						I_Level: 1,
						I_EXP:   1,
					},
				},
				User_select_reward: nil,
			},
		}, common_model.ErrorRetCode{}
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
