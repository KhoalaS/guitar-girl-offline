package game

import (
	"encoding/json"
	"time"

	embeds "github.com/KhoalaS/guitar-girl-offline"
	"github.com/KhoalaS/guitar-girl-offline/pkg/model/common_model"
	"github.com/KhoalaS/guitar-girl-offline/pkg/model/main_model"
	"github.com/KhoalaS/guitar-girl-offline/pkg/model/store_model"
	"github.com/KhoalaS/guitar-girl-offline/pkg/model/user_model"
	"github.com/rs/zerolog/log"
)

type GameService interface {
	Init(params main_model.InitDataInfo) (main_model.InitRetDataInfo, common_model.ErrorRetCode)
	GetServerTime(params main_model.MainDataInfo) (main_model.GetServerTimeRetDataInfo, common_model.ErrorRetCode)
	UserLogin(params user_model.UserLoginDataInfo) (user_model.UserLoginRetDataInfo, common_model.ErrorRetCode)
	GetUpdateTime(params main_model.GetUpdateTimeDataInfo) main_model.GetUpdateTimeRetDataInfo
	DefaultSettingList(params main_model.DefaultSettingListDataInfo) main_model.DefaultSettingListRetDataInfo
	GetGameDataList(params main_model.GetGameDataListDataInfo) (main_model.GetGameDataListRetDataInfo, common_model.ErrorRetCode)
	UserSave(params user_model.UserSaveDataInfo) (user_model.UserSaveRetDataInfo, common_model.ErrorRetCode)
	UserJoin(params user_model.UserJoinDataInfo) (user_model.UserJoinRetDataInfo, common_model.ErrorRetCode)
	BuyVarietyStore(params store_model.BuyVarietyStoreDataInfo) (store_model.BuyVarietyStoreRetDataInfo, common_model.ErrorRetCode)
	BuyCheck(params store_model.BuyCheckDataInfo) (store_model.BuyCheckRetDataInfo, common_model.ErrorRetCode)
	UserLoad(params user_model.UserLoadDataInfo) (user_model.UserLoadRetDataInfo, common_model.ErrorRetCode)
	SetSubscribe(params user_model.SetSubscribeDataInfo) (user_model.SetSubscribeRetDataInfo, common_model.ErrorRetCode)
}

type GameServiceImpl struct {
	UserRepository
	UserAreaRepository
	UserAchievementRepository
	UserCharacterRepository
	UserCostumeRepository
	Timezone string
}

func (service *GameServiceImpl) UserSave(params user_model.UserSaveDataInfo) (user_model.UserSaveRetDataInfo, common_model.ErrorRetCode) {
	for _, areaData := range params.User_area_info {
		// TODO db transaction inside SetAreaData that accepts UserAreaInfo slice
		service.UserAreaRepository.SetAreaData(params.Uuid, areaData)
	}

	err := service.UserAchievementRepository.SetAchievements(params.Uuid, params.User_achievement)
	if err != nil {
		log.Err(err).Send()
		return user_model.UserSaveRetDataInfo{
				Status: "N",
			}, common_model.ErrorRetCode{
				Code:   102,
				Errmsg: "Error saving user achievements.",
			}
	}

	service.UserCharacterRepository.SetCharacter(params.Uuid, params.User_character)
	service.UserCostumeRepository.SetCostumes(params.Uuid, params.User_costume)

	return user_model.UserSaveRetDataInfo{
		Status: "Y",
	}, common_model.ErrorRetCode{}
}

func (service *GameServiceImpl) Init(params main_model.InitDataInfo) (main_model.InitRetDataInfo, common_model.ErrorRetCode) {
	return main_model.InitRetDataInfo{
		Idx:      253,
		Game_url: "https://game.gtgl.pmang.cloud",
		Cdn_url:  "https://dl.gtgl.pmang.cloud",
	}, common_model.ErrorRetCode{}
}

func (service *GameServiceImpl) GetServerTime(params main_model.MainDataInfo) (main_model.GetServerTimeRetDataInfo, common_model.ErrorRetCode) {
	return getBaseTimeStamp(time.Now(), service.Timezone), common_model.ErrorRetCode{}
}

func (service *GameServiceImpl) UserLogin(params user_model.UserLoginDataInfo) (user_model.UserLoginRetDataInfo, common_model.ErrorRetCode) {
	// TODO internal logic
	userData, err := service.UserRepository.GetUserByMemberId(params.Uuid)
	if err != nil {
		return user_model.UserLoginRetDataInfo{}, common_model.ErrorRetCode{
			Code:   998,
			Errmsg: "Sorry. Not a registered user.",
		}
	}

	areaData, _ := service.GetAllAreaData(params.Uuid)
	areaDataMap := map[int32]user_model.UserAreaData{}
	for _, area := range areaData {
		areaDataMap[int32(area.U_area_num)] = area
	}

	achievements, _ := service.UserAchievementRepository.GetAchievements(params.Uuid)
	characters, _ := service.UserCharacterRepository.GetCharacter(params.Uuid)
	costumes, _ := service.UserCostumeRepository.GetCostumes(params.Uuid)

	if params.U_seq != 0 {
		return user_model.UserLoginRetDataInfo{
			User:      userData,
			Area_data: areaDataMap,
			User_contents: user_model.UserContentsData{
				User_achievement: achievements,
				User_buff:        nil,
				User_candy_shop: []user_model.UserCandyShop{
					{
						I_id:              1,
						I_CurrentBuyCount: 1,
						I_TotalBuyCount:   1,
						L_LastBuyTick:     1767022695,
						Upd_day:           20251230,
					},
					{
						I_id:              2,
						I_CurrentBuyCount: 1,
						I_TotalBuyCount:   1,
						L_LastBuyTick:     1767114498,
						Upd_day:           20251231,
					},
				},
				User_character: characters,
				User_costume:   costumes,
				User_daily_mission: []user_model.UserDailyMission{
					{
						I_id:       1,
						I_Level:    1,
						D_Quantity: 1,
						Upd_date:   "2025-12-30 00:37:17",
					},
					{
						I_id:       2,
						I_Level:    1,
						D_Quantity: 23,
						Upd_date:   "",
					},
					{
						I_id:       3,
						I_Level:    1,
						D_Quantity: 0,
						Upd_date:   "",
					},
					{
						I_id:       4,
						I_Level:    1,
						D_Quantity: 1,
						Upd_date:   "",
					},
					{
						I_id:       5,
						I_Level:    1,
						D_Quantity: 902,
						Upd_date:   "",
					},
					{
						I_id:       6,
						I_Level:    1,
						D_Quantity: 0,
						Upd_date:   "",
					},
				},
				User_follower: []user_model.UserFollower{
					{
						I_id:         1,
						I_Level:      125,
						I_BonusLevel: 5,
					},
					{
						I_id:         2,
						I_Level:      48,
						I_BonusLevel: 1,
					},
					{
						I_id:         3,
						I_Level:      13,
						I_BonusLevel: 0,
					},
					{
						I_id:         4,
						I_Level:      1,
						I_BonusLevel: 0,
					},
				},
				User_music: []user_model.UserMusic{
					{
						I_id:                      1,
						I_Level:                   6,
						I_BonusLevel:              0,
						B_EncoreBonusAppear:       0,
						L_EncoreBonusActivateTime: 1767096954,
						I_EncoreBonusFollowerId:   0,
						I_ChThirdActiveTime:       0,
					},
					{
						I_id:                      2,
						I_Level:                   1,
						I_BonusLevel:              0,
						B_EncoreBonusAppear:       0,
						L_EncoreBonusActivateTime: 1767097063,
						I_EncoreBonusFollowerId:   0,
						I_ChThirdActiveTime:       0,
					},
					{
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
					{
						I_id:           1,
						I_Count:        1,
						I_TotalCount:   1,
						I_PurchaseTime: 1765734809,
						Upd_day:        20251215,
					},
					{
						I_id:           13,
						I_Count:        1,
						I_TotalCount:   1,
						I_PurchaseTime: 1767114545,
						Upd_day:        20251231,
					},
					{
						I_id:           3013,
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
					{
						I_SubscribeID: 5,
						I_Type:        0,
						I_Step:        0,
						I_UpdateTime:  1765734537,
						I_Version:     5,
					},
					{
						I_SubscribeID: 5,
						I_Type:        0,
						I_Step:        1,
						I_UpdateTime:  1767114615,
						I_Version:     5,
					},
				},
				User_ticketcollection: []user_model.UserTicketCollection{
					{
						I_id: 1,
					},
					{
						I_id: 2,
					},
					{
						I_id: 3,
					},
					{
						I_id: 4,
					},
					{
						I_id: 5,
					},
					{
						I_id: 6,
					},
					{
						I_id: 7,
					},
					{
						I_id: 8,
					},
					{
						I_id: 9,
					},
					{
						I_id: 10,
					},
					{
						I_id: 11,
					},
					{
						I_id: 12,
					},
					{
						I_id: 13,
					},
				},
				User_follower_quest: nil,
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
		User: user_model.UserData{
			U_seq: userData.U_seq,
			U_id:  userData.U_id,
		},
	}, common_model.ErrorRetCode{}

}

func (service *GameServiceImpl) GetUpdateTime(params main_model.GetUpdateTimeDataInfo) main_model.GetUpdateTimeRetDataInfo {
	// TODO internal logic
	return main_model.GetUpdateTimeRetDataInfo{
		Upd_time: 1761876067,
	}
}

func (service *GameServiceImpl) DefaultSettingList(params main_model.DefaultSettingListDataInfo) main_model.DefaultSettingListRetDataInfo {
	return main_model.DefaultSettingListRetDataInfo{
		Status: "Y",
		Setting_list: []main_model.DefaultSettingDataList{
			{
				Setting_key:   "shop_call_wait",
				Setting_value: "2",
			},
			{
				Setting_key:   "ad_cancel_delay",
				Setting_value: "5",
			},
			{
				Setting_key:   "force_menu_restric",
				Setting_value: "true",
			},
			{
				Setting_key:   "max_bundle_texutre_loading",
				Setting_value: "1",
			},
			{
				Setting_key:   "weather_snow",
				Setting_value: "true",
			},
			{
				Setting_key:   "check_time_cheat",
				Setting_value: "false",
			},
			{
				Setting_key:   "ap_max",
				Setting_value: "60",
			},
			{
				Setting_key:   "ap_time",
				Setting_value: "300",
			},
			{
				Setting_key:   "ap_use",
				Setting_value: "5",
			},
			{
				Setting_key:   "character_exp",
				Setting_value: "45",
			},
			{
				Setting_key:   "ap_shop_id",
				Setting_value: "31",
			},
			{
				Setting_key:   "ap_ad_list_id",
				Setting_value: "2",
			},
			{
				Setting_key:   "ch_third_unlock_follower_id",
				Setting_value: "5",
			},
			{
				Setting_key:   "ch_third_unlock_follower_level",
				Setting_value: "1",
			},
			{
				Setting_key:   "ad_follower_profile_exp_persent",
				Setting_value: "0.1",
			},
			{
				Setting_key:   "follower_profile_bonus_percent",
				Setting_value: "1",
			},
			{
				Setting_key:   "spine_load_max_count",
				Setting_value: "0",
			},
			{
				Setting_key:   "log_flag",
				Setting_value: "true",
			},
			{
				Setting_key:   "weather_cherry_blossom",
				Setting_value: "false",
			},
			{
				Setting_key:   "active_goods_ticket_buff",
				Setting_value: "true",
			},
			{
				Setting_key:   "active_fan_art_costume",
				Setting_value: "true",
			},
			{
				Setting_key:   "pass_goods_sale_start_time",
				Setting_value: "2022-02-19 00:00",
			},
			{
				Setting_key:   "pass_goods_sale_end_time",
				Setting_value: "2999-02-27 23:59",
			},
			{
				Setting_key:   "attendacne_buff_value",
				Setting_value: "0.01",
			},
			{
				Setting_key:   "attendacne_buff_limit_day",
				Setting_value: "10000",
			},
			{
				Setting_key:   "log_flag_unfinished_purchase",
				Setting_value: "true",
			},
			{
				Setting_key:   "jp_laws_flag",
				Setting_value: "true",
			},
		},
	}
}

func (service *GameServiceImpl) GetGameDataList(params main_model.GetGameDataListDataInfo) (main_model.GetGameDataListRetDataInfo, common_model.ErrorRetCode) {
	var data main_model.GetGameDataListRetDataInfo

	err := json.Unmarshal(embeds.GameData, &data)
	if err != nil {
		return main_model.GetGameDataListRetDataInfo{}, common_model.ErrorRetCode{Code: 1, Errmsg: err.Error()}
	}

	return data, common_model.ErrorRetCode{}
}

func (service *GameServiceImpl) UserJoin(params user_model.UserJoinDataInfo) (user_model.UserJoinRetDataInfo, common_model.ErrorRetCode) {
	// TODO put everything inside a transaction

	user, err := service.UserRepository.CreateUser(params.Uuid, params.Device_uuid)
	if err != nil {
		return user_model.UserJoinRetDataInfo{}, common_model.ErrorRetCode{
			Code:   800,
			Errmsg: "Error creating user",
		}
	}

	service.UserAreaRepository.SetAreaData(params.Uuid, user_model.SaveUserAreaInfo{
		U_area_num:          1,
		D_Candy:             0,
		D_Like:              0,
		I_UserFanCount:      0,
		I_UserFanGrade:      1,
		I_SelectedCostumeId: 1,
		I_SelectedMusicId:   1,
		I_SelectedGuitarId:  1,
		S_TutorialList:      "",
		S_Gp1:               "",
		S_Gp2:               "",
	})

	service.UserAreaRepository.SetAreaData(params.Uuid, user_model.SaveUserAreaInfo{
		U_area_num:          2,
		D_Candy:             0,
		D_Like:              0,
		I_UserFanCount:      0,
		I_UserFanGrade:      1,
		I_SelectedCostumeId: 1,
		I_SelectedMusicId:   1,
		I_SelectedGuitarId:  1,
		S_TutorialList:      "",
		S_Gp1:               "",
		S_Gp2:               "",
	})

	service.UserCharacterRepository.SetCharacter(params.Uuid, []user_model.SaveUserCharacter{
		{
			I_id:         1,
			I_Level:      1,
			I_BonusLevel: 0,
		},
	})

	defaultAchievements := []user_model.SaveUserAchievement{}
	for i := 1; i <= 10; i++ {
		defaultAchievements = append(defaultAchievements, user_model.SaveUserAchievement{
			I_id:       int64(i),
			D_Quantity: 1,
			S_Quantity: "",
		})
	}

	service.UserAchievementRepository.SetAchievements(params.Uuid, defaultAchievements)

	defaultCostumes := []user_model.SaveUserCostume{
		{
			I_id:         1,
			I_Level:      1,
			I_BonusLevel: 0,
		},
	}
	service.UserCostumeRepository.SetCostumes(params.Uuid, defaultCostumes)

	return user_model.UserJoinRetDataInfo{
		U_seq: user.U_seq,
		U_id:  user.U_id,
	}, common_model.ErrorRetCode{}
}

func (service *GameServiceImpl) BuyVarietyStore(params store_model.BuyVarietyStoreDataInfo) (store_model.BuyVarietyStoreRetDataInfo, common_model.ErrorRetCode) {
	// TODO internal logic, storing purchase
	user, _ := service.UserRepository.GetUserByMemberId(params.Uuid)

	return store_model.BuyVarietyStoreRetDataInfo{
		U_cp:         user.U_cp,
		U_candy:      user.U_candy,
		Reward_type:  4,
		Reward_id:    params.Idx,
		Reward_value: 1,
		Status:       "Y",
	}, common_model.ErrorRetCode{}
}

func (service *GameServiceImpl) BuyCheck(params store_model.BuyCheckDataInfo) (store_model.BuyCheckRetDataInfo, common_model.ErrorRetCode) {
	return store_model.BuyCheckRetDataInfo{
		Result: "Y",
	}, common_model.ErrorRetCode{}
}

func (service *GameServiceImpl) UserLoad(params user_model.UserLoadDataInfo) (user_model.UserLoadRetDataInfo, common_model.ErrorRetCode) {
	// TODO
	if params.Type == "shop" {
		return user_model.UserLoadRetDataInfo{
			User_contents: user_model.UserContentsData{
				User_shop: []user_model.UserShop{
					{
						I_id:           1,
						I_Count:        1,
						I_TotalCount:   1,
						I_PurchaseTime: 1765734809,
						Upd_day:        20251215,
					},
					{
						I_id:           13,
						I_Count:        1,
						I_TotalCount:   1,
						I_PurchaseTime: 1767114545,
						Upd_day:        20251231,
					},
					{
						I_id:           3012,
						I_Count:        1,
						I_TotalCount:   1,
						I_PurchaseTime: 1767114545,
						Upd_day:        20251231,
					},
					{
						I_id:           3013,
						I_Count:        1,
						I_TotalCount:   1,
						I_PurchaseTime: 1767114545,
						Upd_day:        20251231,
					},
				},
			},
		}, common_model.ErrorRetCode{}
	}

	return user_model.UserLoadRetDataInfo{}, common_model.ErrorRetCode{
		Code:   999,
		Errmsg: "Unimplemented",
	}
}

func (service *GameServiceImpl) SetSubscribe(params user_model.SetSubscribeDataInfo) (user_model.SetSubscribeRetDataInfo, common_model.ErrorRetCode) {
	// TODO db persistence

	subscriptions := []user_model.UserSubscribeList{}

	for id := range 100 {
		subscriptions = append(subscriptions, user_model.UserSubscribeList{
			I_SubscribeID: int64(id + 1),
			I_ActiveTime:  time.Now().Unix(),
			I_isActive:    1,
		})
	}

	return user_model.SetSubscribeRetDataInfo{
		U_seq:               params.U_seq,
		User_subscribe_list: subscriptions,
	}, common_model.ErrorRetCode{}
}
