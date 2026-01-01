package user_model

import "github.com/KhoalaS/guitar-girl-offline/pkg/model/common_model"

type CheckBuyProductReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        CheckBuyProductRetDataInfo   `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type DeleteUserRetDataInfo struct {
	Result string `thrift:",1,omitempty"`
}

type UserDailyMission struct {
	I_id       int64  `thrift:",1,omitempty"`
	I_Level    int64  `thrift:",2,omitempty"`
	D_Quantity int64  `thrift:",3,omitempty"`
	Upd_date   string `thrift:",4,omitempty"`
}

type UserPurchase struct {
	Call     string               `thrift:",1,omitempty"`
	Data     UserPurchaseDataInfo `thrift:",2,omitempty"`
	Sub_mode string               `thrift:",3,omitempty"`
}

type UserSaveData struct {
	U_candy          int64 `thrift:",1,omitempty"`
	U_like           int64 `thrift:",2,omitempty"`
	U_fans           int64 `thrift:",3,omitempty"`
	U_girl_level     int64 `thrift:",4,omitempty"`
	U_skill_level    int64 `thrift:",5,omitempty"`
	U_mate_level     int64 `thrift:",6,omitempty"`
	U_follower_level int64 `thrift:",7,omitempty"`
	U_play_level     int64 `thrift:",8,omitempty"`
}

type DeleteUser struct {
	Call string             `thrift:",1,omitempty"`
	Data DeleteUserDataInfo `thrift:",2,omitempty"`
}

type MoreGamesLogDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	App_id      int32  `thrift:",5,omitempty"`
}

type BuyAvatarDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Idx         int16  `thrift:",5,omitempty"`
}

type BuyUserTitle struct {
	Call string               `thrift:",1,omitempty"`
	Data BuyUserTitleDataInfo `thrift:",2,omitempty"`
}

type GetPurchaseDiamondLogReturn struct {
	Error       common_model.ErrorRetCode        `thrift:",1,omitempty"`
	Mode        string                           `thrift:",2,omitempty"`
	Call        string                           `thrift:",3,omitempty"`
	Data        GetPurchaseDiamondLogRetDataInfo `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData     `thrift:",5,omitempty"`
}

type GetUserCollectionDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type PaidEventPointRetDataInfo struct {
	U_cp          int64   `thrift:",1,omitempty"`
	U_candy       float64 `thrift:",2,omitempty"`
	I_SubscribeID int64   `thrift:",3,omitempty"`
	I_Point       int64   `thrift:",4,omitempty"`
	I_Version     int32   `thrift:",5,omitempty"`
}

type SetGameReward struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        SetGameRewardDataInfo  `thrift:",2,omitempty"`
	Common_data common_model.ParamData `thrift:",3,omitempty"`
}

type PlayMusicReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        PlayMusicRetDataInfo         `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type SaveUserMessenger struct {
	I_MessengerChatRoomId int64  `thrift:",1,omitempty"`
	I_LastConfirmIndex    int64  `thrift:",2,omitempty"`
	S_UnlockGroupList     string `thrift:",3,omitempty"`
	L_UpdateTimeTicks     int64  `thrift:",4,omitempty"`
}

type FacebookUserJoin struct {
	Call        string                   `thrift:",1,omitempty"`
	Data        FacebookUserJoinDataInfo `thrift:",2,omitempty"`
	Common_data common_model.ParamData   `thrift:",3,omitempty"`
}

type BuyCollectionMusicReturn struct {
	Error       common_model.ErrorRetCode     `thrift:",1,omitempty"`
	Mode        string                        `thrift:",2,omitempty"`
	Call        string                        `thrift:",3,omitempty"`
	Data        BuyCollectionMusicRetDataInfo `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData  `thrift:",5,omitempty"`
}

type ChallengeMissionListDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Music_idx   int32  `thrift:",5,omitempty"`
	Difficulty  int16  `thrift:",6,omitempty"`
	Up          int32  `thrift:",7,omitempty"`
	Down        int32  `thrift:",8,omitempty"`
}

type TransferUser struct {
	Call string               `thrift:",1,omitempty"`
	Data TransferUserDataInfo `thrift:",2,omitempty"`
}

type UserSubscribeList struct {
	I_SubscribeID int64 `thrift:",1,omitempty"`
	I_ActiveTime  int64 `thrift:",2,omitempty"`
	I_isActive    int64 `thrift:",3,omitempty"`
}

type ChThirdStreamStageRetDataInfo struct {
	I_id                  int32               `thrift:",1,omitempty"`
	User_ap               UserApData          `thrift:",2,omitempty"`
	User_follower_profile UserFollowerProfile `thrift:",3,omitempty"`
	Reward_data           map[any]any         `thrift:",4,omitempty"` // TODO
}

type ChallengeMissionListRetDataInfo struct {
	Music_idx    int32       `thrift:",1,omitempty"`
	Difficulty   int16       `thrift:",2,omitempty"`
	Mission_list map[any]any `thrift:",3,omitempty"` // TODO
}

type GetUserItemInven struct {
	Call string                   `thrift:",1,omitempty"`
	Data GetUserItemInvenDataInfo `thrift:",2,omitempty"`
}

type ItemPurchaseRetDataInfo struct {
	Item_type   int16  `thrift:",1,omitempty"`
	Item_value  int64  `thrift:",2,omitempty"`
	Pay_id      string `thrift:",3,omitempty"`
	Product_id  string `thrift:",4,omitempty"`
	Purchase_id string `thrift:",5,omitempty"`
}

type UserInfo struct {
	U_like                float64 `thrift:",1,omitempty"`
	U_fans                int64   `thrift:",2,omitempty"`
	U_fans_grade          int16   `thrift:",3,omitempty"`
	U_selected_costume_id int64   `thrift:",4,omitempty"`
	U_selected_music_id   int64   `thrift:",5,omitempty"`
}

type BuyCollectionMusicDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Idx         int32  `thrift:",5,omitempty"`
}

type BuyVarietyStore struct {
	Call string                  `thrift:",1,omitempty"`
	Data BuyVarietyStoreDataInfo `thrift:",2,omitempty"`
}

type GetChThirdReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        GetChThirdRetDataInfo        `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type GetDiamondBonusReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        GetDiamondBonusRetDataInfo   `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type UserPurchaseErrorLogRetDataInfo struct {
	Ret string `thrift:",1,omitempty"`
}

type BuyVarietyStoreDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Idx         int32  `thrift:",5,omitempty"`
}

type CheckBuyShopReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        CheckBuyShopRetDataInfo      `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type GetAllPointDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type SetAdRewardReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        SetAdRewardRetDataInfo       `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type UpdateAchievementReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        UpdateAchievementRetDataInfo `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type UserBuff struct {
	I_id         int64 `thrift:",1,omitempty"`
	I_Level      int64 `thrift:",2,omitempty"`
	I_ActiveTime int64 `thrift:",3,omitempty"`
	I_EndTime    int64 `thrift:",4,omitempty"`
}

type UserItemList struct {
	Item_type int32 `thrift:",1,omitempty"`
	Item_id   int32 `thrift:",2,omitempty"`
	Count     int64 `thrift:",3,omitempty"`
}

type BuyCheckDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Idx         int32  `thrift:",5,omitempty"`
}

type BuyExtraModeRetDataInfo struct {
	Music_idx int32 `thrift:",1,omitempty"`
}

type MusicPointReviewRetDataInfo struct {
	Music_idx  int32   `thrift:",1,omitempty"`
	Point      float64 `thrift:",2,omitempty"`
	Difficulty int16   `thrift:",3,omitempty"`
}

type UseCreditRetDataInfo struct {
	U_gold   int64 `thrift:",1,omitempty"`
	U_cp     int64 `thrift:",2,omitempty"`
	U_mp     int64 `thrift:",3,omitempty"`
	U_credit int64 `thrift:",4,omitempty"`
}

type UserAdList struct {
	I_id           int32 `thrift:",1,omitempty"`
	I_Count        int32 `thrift:",2,omitempty"`
	I_TotalCount   int32 `thrift:",3,omitempty"`
	I_LastViewTime int32 `thrift:",4,omitempty"`
}

type GetProfileMusicReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        []any                        `thrift:",4,omitempty"` // TODO
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type MusicPointReview struct {
	Call string                   `thrift:",1,omitempty"`
	Data MusicPointReviewDataInfo `thrift:",2,omitempty"`
}

type UserTitleListRetDataInfo struct {
	U_seq int32 `thrift:",1,omitempty"`
	Title []any `thrift:",2,omitempty"` // TODO
}

type BuyExtraMode struct {
	Call string               `thrift:",1,omitempty"`
	Data BuyExtraModeDataInfo `thrift:",2,omitempty"`
}

type BuyPackageDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Store_idx   int64  `thrift:",5,omitempty"`
	Product_id  int64  `thrift:",6,omitempty"`
}

type ChThirdStreamStage struct {
	Call        string                     `thrift:",1,omitempty"`
	Data        ChThirdStreamStageDataInfo `thrift:",2,omitempty"`
	Common_data common_model.ParamData     `thrift:",3,omitempty"`
}

type SaveUserGuitar struct {
	I_id         int64 `thrift:",1,omitempty"`
	I_Level      int64 `thrift:",2,omitempty"`
	I_BonusLevel int64 `thrift:",3,omitempty"`
}

type SetTutorialNewRetDataInfo struct {
	U_seq    int32 `thrift:",1,omitempty"`
	Tutorial []any `thrift:",2,omitempty"` // TODO
}

type SetUserFollowerProfileRewardRetDataInfo struct {
	I_id                  int32               `thrift:",1,omitempty"`
	I_level               int32               `thrift:",2,omitempty"`
	Reward_data           []any               `thrift:",3,omitempty"` // TODO
	User_follower_profile UserFollowerProfile `thrift:",4,omitempty"`
}

type SetUserFollowerQuestRetDataInfo struct {
	I_CurrentID       int32   `thrift:",1,omitempty"`
	I_CompleteID      int32   `thrift:",2,omitempty"`
	D_ConditionValue1 float64 `thrift:",3,omitempty"`
	D_ConditionValue2 float64 `thrift:",4,omitempty"`
	D_ConditionValue3 float64 `thrift:",5,omitempty"`
	I_RewardReceived1 int16   `thrift:",6,omitempty"`
	I_RewardReceived2 int16   `thrift:",7,omitempty"`
	I_RewardReceived3 int16   `thrift:",8,omitempty"`
	Next_flg          int16   `thrift:",9,omitempty"`
	Reward_data       []any   `thrift:",10,omitempty"` // TODO
}

type UserLevel struct {
	U_girl_level     int64 `thrift:",1,omitempty"`
	U_skill_level    int64 `thrift:",2,omitempty"`
	U_mate_level     int64 `thrift:",3,omitempty"`
	U_follower_level int64 `thrift:",4,omitempty"`
	U_play_level     int64 `thrift:",5,omitempty"`
}

type CheckBuyProduct struct {
	Call string                  `thrift:",1,omitempty"`
	Data CheckBuyProductDataInfo `thrift:",2,omitempty"`
}

type GetProfileRetDataInfo struct {
	U_avatar        int32  `thrift:",1,omitempty"`
	U_id            string `thrift:",2,omitempty"`
	U_nick          string `thrift:",3,omitempty"`
	U_country       string `thrift:",4,omitempty"`
	U_title         int32  `thrift:",5,omitempty"`
	Allper          int32  `thrift:",6,omitempty"`
	Allcom          int32  `thrift:",7,omitempty"`
	Gold            int32  `thrift:",8,omitempty"`
	Silver          int32  `thrift:",9,omitempty"`
	Bronze          int32  `thrift:",10,omitempty"`
	Single_cnt      int64  `thrift:",11,omitempty"`
	Multi_cnt       int64  `thrift:",12,omitempty"`
	Master_cnt      int64  `thrift:",13,omitempty"`
	Single_score    int64  `thrift:",14,omitempty"`
	Multi_score     int64  `thrift:",15,omitempty"`
	Local_rank      int64  `thrift:",16,omitempty"`
	Individual_rank int64  `thrift:",17,omitempty"`
}

type GetUserItemInvenRetDataInfo struct {
	Item_inven_list map[any]any `thrift:",1,omitempty"` // TODO
	Empty           string      `thrift:",2,omitempty"`
}

type UserCharacter struct {
	I_id         int64 `thrift:",1,omitempty"`
	I_Level      int64 `thrift:",2,omitempty"`
	I_BonusLevel int64 `thrift:",3,omitempty"`
}

type UserItemInvenList struct {
	Item_list []any `thrift:",1,omitempty"` // TODO
}

type UserLoad struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        UserLoadDataInfo       `thrift:",2,omitempty"`
	Common_data common_model.ParamData `thrift:",3,omitempty"`
}

type UserMessenger struct {
	I_MessengerChatRoomId int64  `thrift:",1,omitempty"`
	I_LastConfirmIndex    int64  `thrift:",2,omitempty"`
	S_UnlockGroupList     string `thrift:",3,omitempty"`
	L_UpdateTimeTicks     int64  `thrift:",4,omitempty"`
}

type GetChoiceUserRetDataInfo struct {
	User        ChoiceUserData `thrift:",1,omitempty"`
	Choice_user ChoiceUserData `thrift:",2,omitempty"`
}

type SetAttendanceReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        SetAttendanceRetDataInfo     `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type SetFollowerQuestInfiniteRetDataInfo struct {
	U_seq               int32             `thrift:",1,omitempty"`
	User_follower_quest UserFollowerQuest `thrift:",2,omitempty"`
}

type GetTransferIdReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        GetTransferIdRetDataInfo     `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type SaveUserSkill struct {
	I_id               int64 `thrift:",1,omitempty"`
	B_Activate         int16 `thrift:",2,omitempty"`
	L_ActivateOnTicks  int64 `thrift:",3,omitempty"`
	L_ActivateOffTicks int64 `thrift:",4,omitempty"`
}

type UpdateAvatarReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        UpdateAvatarRetDataInfo      `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type UpdateUserTitleReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        UpdateUserTitleRetDataInfo   `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type UserFollower struct {
	I_id         int64 `thrift:",1,omitempty"`
	I_Level      int64 `thrift:",2,omitempty"`
	I_BonusLevel int64 `thrift:",3,omitempty"`
}

type GetUserInfoReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        GetUserInfoRetDataInfo       `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type PaidEventPoint struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        PaidEventPointDataInfo `thrift:",2,omitempty"`
	Common_data common_model.ParamData `thrift:",3,omitempty"`
}

type SetAdRewardDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	I_id        int32  `thrift:",5,omitempty"`
	Param1      string `thrift:",6,omitempty"`
	Param2      string `thrift:",7,omitempty"`
	Param3      string `thrift:",8,omitempty"`
}

type SetAttendanceRetDataInfo struct {
	Status                          string            `thrift:",1,omitempty"`
	User_follower_quest             UserFollowerQuest `thrift:",2,omitempty"`
	Attendance_count                int32             `thrift:",3,omitempty"`
	Attendance_date                 int32             `thrift:",4,omitempty"`
	Max_coutinuous_attendance_count int32             `thrift:",5,omitempty"`
}

type UpdateAchievement struct {
	Call string                    `thrift:",1,omitempty"`
	Data UpdateAchievementDataInfo `thrift:",2,omitempty"`
}

type UserShop struct {
	I_id           int64 `thrift:",1,omitempty"`
	I_Count        int64 `thrift:",2,omitempty"`
	I_TotalCount   int64 `thrift:",3,omitempty"`
	I_PurchaseTime int64 `thrift:",4,omitempty"`
	Upd_day        int64 `thrift:",5,omitempty"`
}

type UserTitleListDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type UserFollowerGiftItem struct {
	I_id    int32 `thrift:",1,omitempty"`
	I_Value int32 `thrift:",2,omitempty"`
}

type AdViewLog struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        AdViewLogDataInfo      `thrift:",2,omitempty"`
	Common_data common_model.ParamData `thrift:",3,omitempty"`
}

type BuyItemRetDataInfo struct {
	Idx      int32 `thrift:",1,omitempty"`
	U_gold   int64 `thrift:",2,omitempty"`
	U_cp     int64 `thrift:",3,omitempty"`
	U_mp     int64 `thrift:",4,omitempty"`
	U_credit int64 `thrift:",5,omitempty"`
}

type CheckBuyShopDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Shop_id     string `thrift:",5,omitempty"`
}

type GetUserInfo struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        GetUserInfoDataInfo    `thrift:",2,omitempty"`
	Common_data common_model.ParamData `thrift:",3,omitempty"`
}

type SetBookmarkReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        []any                        `thrift:",5,omitempty"` // TODO
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type UpdateAvatarDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	U_avatar    int16  `thrift:",5,omitempty"`
}

type UserMusic struct {
	I_id                      int64 `thrift:",1,omitempty"`
	I_Level                   int64 `thrift:",2,omitempty"`
	I_BonusLevel              int64 `thrift:",3,omitempty"`
	B_EncoreBonusAppear       int64 `thrift:",4,omitempty"`
	L_EncoreBonusActivateTime int64 `thrift:",5,omitempty"`
	I_EncoreBonusFollowerId   int64 `thrift:",6,omitempty"`
	I_ChThirdActiveTime       int64 `thrift:",7,omitempty"`
}

type GetProfileMusicRetDataInfo struct {
	Music_idx    int32 `thrift:",1,omitempty"`
	Single_cnt   int64 `thrift:",2,omitempty"`
	Multi_cnt    int64 `thrift:",3,omitempty"`
	Master_cnt   int64 `thrift:",4,omitempty"`
	Single_score int64 `thrift:",5,omitempty"`
	Multi_score  int64 `thrift:",6,omitempty"`
}

type RetReward struct {
	Reward_type  int16   `thrift:",1,omitempty"`
	Reward_id    int32   `thrift:",2,omitempty"`
	Reward_value float64 `thrift:",3,omitempty"`
}

type SetPassRewardReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        SetPassRewardRetDataInfo     `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type UserAvatarListReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        UserAvatarListRetDataInfo    `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type GetAllPointRetDataInfo struct {
	U_gold           int64 `thrift:",1,omitempty"`
	U_cp             int64 `thrift:",2,omitempty"`
	U_mp             int64 `thrift:",3,omitempty"`
	U_credit         int64 `thrift:",4,omitempty"`
	Next_credit_time int64 `thrift:",5,omitempty"`
	Max_credit_time  int64 `thrift:",6,omitempty"`
}

type GetChoiceUserReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        GetChoiceUserRetDataInfo     `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type GetUserCollectionRetDataInfo struct {
	Music_idx   int64  `thrift:",1,omitempty"`
	Value       int64  `thrift:",2,omitempty"`
	Get_type    int16  `thrift:",3,omitempty"`
	Get_time    string `thrift:",4,omitempty"`
	Create_time string `thrift:",5,omitempty"`
}

type GetUserItemInvenDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type SetBookmarkRetDataInfo struct {
	Music_idx int32 `thrift:",1,omitempty"`
	Flag      int16 `thrift:",2,omitempty"`
}

type UpdateNickNameRetDataInfo struct {
	Nickname string `thrift:",1,omitempty"`
}

type UserUnit struct {
	I_id    int64 `thrift:",1,omitempty"`
	I_Level int64 `thrift:",2,omitempty"`
}

type BuyExtraModeDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Music_idx   int32  `thrift:",5,omitempty"`
	Key_flg     int64  `thrift:",6,omitempty"`
}

type ChangeUserDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Change_uuid string `thrift:",5,omitempty"`
}

type GetAllPoint struct {
	Call     string              `thrift:",1,omitempty"`
	Data     GetAllPointDataInfo `thrift:",2,omitempty"`
	Sub_mode string              `thrift:",3,omitempty"`
}

type GetUserItemInvenReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        GetUserItemInvenRetDataInfo  `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type SetBookmark struct {
	Call string              `thrift:",1,omitempty"`
	Data SetBookmarkDataInfo `thrift:",2,omitempty"`
}

type UserApData struct {
	I_Ap         int32 `thrift:",1,omitempty"`
	I_FullApTime int32 `thrift:",2,omitempty"`
	I_MaxAp      int32 `thrift:",3,omitempty"`
}

type UserChThirdChapterReward struct {
	I_id              int32 `thrift:",1,omitempty"`
	I_ReceivedReward1 int32 `thrift:",2,omitempty"`
	I_ReceivedReward2 int32 `thrift:",3,omitempty"`
	I_ReceivedReward3 int32 `thrift:",4,omitempty"`
}

type BuyCollectionMusic struct {
	Call string                     `thrift:",1,omitempty"`
	Data BuyCollectionMusicDataInfo `thrift:",2,omitempty"`
}

type BuyExtraModeReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        BuyExtraModeRetDataInfo      `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type ChThirdStageDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	I_id        int32  `thrift:",5,omitempty"`
	Music_id    int32  `thrift:",6,omitempty"`
	Profile_ids string `thrift:",7,omitempty"`
}

type GetMusicRewardReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        GetMusicRewardRetDataInfo    `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type PlayMusic struct {
	Call string            `thrift:",1,omitempty"`
	Data PlayMusicDataInfo `thrift:",2,omitempty"`
}

type SetSubscribeRetDataInfo struct {
	U_seq               int32 `thrift:",1,omitempty"`
	User_subscribe_list []any `thrift:",2,omitempty"` // TODO
}

type FacebookUserJoinDataInfo struct {
	Access_token string `thrift:",1,omitempty"`
}

type AdViewLogReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        AdViewLogRetDataInfo         `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type CollectionRetData struct {
	Music_idx int64 `thrift:",1,omitempty"`
	Cnt       int64 `thrift:",2,omitempty"`
}

type SetReviewPopupDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type SetTutorialReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        SetTutorialRetDataInfo       `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type UserDel struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        UserDelDataInfo        `thrift:",2,omitempty"`
	Common_data common_model.ParamData `thrift:",3,omitempty"`
}

type UserPurchaseErrorLog struct {
	Call string                       `thrift:",1,omitempty"`
	Data UserPurchaseErrorLogDataInfo `thrift:",2,omitempty"`
}

type UserTitleList struct {
	Call string                `thrift:",1,omitempty"`
	Data UserTitleListDataInfo `thrift:",2,omitempty"`
}

type BuyPackageMusic struct {
	Music_idx int64 `thrift:",1,omitempty"`
}

type CheckUserReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        CheckUserRetDataInfo         `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type GetChThirdStarRewardReturn struct {
	Error       common_model.ErrorRetCode       `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet      `thrift:",2,omitempty"`
	Mode        string                          `thrift:",3,omitempty"`
	Call        string                          `thrift:",4,omitempty"`
	Data        GetChThirdStarRewardRetDataInfo `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData    `thrift:",6,omitempty"`
}

type SaveUserEventPoint struct {
	S_EventType string `thrift:",1,omitempty"`
	I_DataID    int64  `thrift:",2,omitempty"`
	I_Point     int64  `thrift:",3,omitempty"`
	I_Step      int64  `thrift:",4,omitempty"`
	I_Version   int32  `thrift:",5,omitempty"`
}

type SetPassReward struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        SetPassRewardDataInfo  `thrift:",2,omitempty"`
	Common_data common_model.ParamData `thrift:",3,omitempty"`
}

type UpdateAvatar struct {
	Call string               `thrift:",1,omitempty"`
	Data UpdateAvatarDataInfo `thrift:",2,omitempty"`
}

type UserData struct {
	U_seq                   int32   `thrift:",1,omitempty"`
	U_id                    string  `thrift:",2,omitempty"`
	U_name                  string  `thrift:",3,omitempty"`
	U_nick                  string  `thrift:",4,omitempty"`
	U_cp                    int64   `thrift:",5,omitempty"`
	U_candy                 float64 `thrift:",6,omitempty"`
	U_like                  float64 `thrift:",7,omitempty"`
	U_fans                  int64   `thrift:",8,omitempty"`
	U_fans_grade            int16   `thrift:",9,omitempty"`
	U_selected_costume_id   int64   `thrift:",10,omitempty"`
	U_selected_music_id     int64   `thrift:",11,omitempty"`
	U_retain_continuous_tap int16   `thrift:",12,omitempty"`
	U_join_type             int16   `thrift:",13,omitempty"`
	U_last_login            string  `thrift:",14,omitempty"`
	U_last_communication    string  `thrift:",15,omitempty"`
	U_save_date             string  `thrift:",16,omitempty"`
	U_create_time           string  `thrift:",17,omitempty"`
	U_tutorial_step         int16   `thrift:",18,omitempty"`
	U_review_popup          string  `thrift:",19,omitempty"`
	Device_uuid             string  `thrift:",20,omitempty"`
	U_samseck_step          int16   `thrift:",21,omitempty"`
	U_free_cp               int64   `thrift:",22,omitempty"`
	U_charge_cp             int64   `thrift:",23,omitempty"`
}

type SetFollowerProfileGift struct {
	Call        string                         `thrift:",1,omitempty"`
	Data        SetFollowerProfileGiftDataInfo `thrift:",2,omitempty"`
	Common_data common_model.ParamData         `thrift:",3,omitempty"`
}

type UserPurchaseErrorLogReturn struct {
	Error       common_model.ErrorRetCode       `thrift:",1,omitempty"`
	Mode        string                          `thrift:",2,omitempty"`
	Call        string                          `thrift:",3,omitempty"`
	Data        UserPurchaseErrorLogRetDataInfo `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData    `thrift:",5,omitempty"`
}

type ChThirdStageReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        ChThirdStageRetDataInfo      `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type ChallengeMissionList struct {
	Call string                       `thrift:",1,omitempty"`
	Data ChallengeMissionListDataInfo `thrift:",2,omitempty"`
}

type SaveUserAreaInfo struct {
	U_area_num          int16   `thrift:",1,omitempty"`
	D_Like              float64 `thrift:",2,omitempty"`
	I_UserFanCount      int64   `thrift:",3,omitempty"`
	I_UserFanGrade      int16   `thrift:",4,omitempty"`
	I_SelectedCostumeId int64   `thrift:",5,omitempty"`
	I_SelectedMusicId   int64   `thrift:",6,omitempty"`
	I_SelectedGuitarId  int64   `thrift:",7,omitempty"`
	D_Candy             float64 `thrift:",8,omitempty"`
	S_TutorialList      string  `thrift:",9,omitempty"`
	S_Gp1               string  `thrift:",10,omitempty"`
	S_Gp2               string  `thrift:",11,omitempty"`
}

type SetEventReward struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        SetEventRewardDataInfo `thrift:",2,omitempty"`
	Common_data common_model.ParamData `thrift:",3,omitempty"`
}

type TransferUserDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Email       string `thrift:",5,omitempty"`
}

type UpdateAvatarRetDataInfo struct {
	U_avatar int16 `thrift:",1,omitempty"`
}

type UseCoupon struct {
	Call string            `thrift:",1,omitempty"`
	Data UseCouponDataInfo `thrift:",2,omitempty"`
}

type UseCouponReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        UseCouponRetDataInfo         `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type UserAdLevel struct {
	I_id    int32 `thrift:",1,omitempty"`
	I_Level int32 `thrift:",2,omitempty"`
	I_EXP   int32 `thrift:",3,omitempty"`
}

type BuyPiece struct {
	Call string           `thrift:",1,omitempty"`
	Data BuyPieceDataInfo `thrift:",2,omitempty"`
}

type DeleteUserReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        DeleteUserRetDataInfo        `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type HibernationLogDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type LogCashCompleteDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Shop_id     int32  `thrift:",5,omitempty"`
	Value       string `thrift:",6,omitempty"`
}

type LogCashCompleteRetDataInfo struct {
	Status string `thrift:",1,omitempty"`
}

type SaveUserDailyMission struct {
	I_id       int64 `thrift:",1,omitempty"`
	D_Quantity int64 `thrift:",2,omitempty"`
}

type UserSave struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        UserSaveDataInfo       `thrift:",2,omitempty"`
	Common_data common_model.ParamData `thrift:",3,omitempty"`
}

type AdViewLogDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Ad_type     string `thrift:",5,omitempty"`
	Ad_name     string `thrift:",6,omitempty"`
}

type LastSaveTime struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        LastSaveTimeDataInfo   `thrift:",2,omitempty"`
	Common_data common_model.ParamData `thrift:",3,omitempty"`
}

type UpdateUserTitleRetDataInfo struct {
	U_title int16 `thrift:",1,omitempty"`
}

type UserSaveRetDataInfo struct {
	Status string `thrift:",1,omitempty"`
}

type BuyPackageRetDataInfo struct {
	U_gold     int64 `thrift:",1,omitempty"`
	U_cp       int64 `thrift:",2,omitempty"`
	U_mp       int64 `thrift:",3,omitempty"`
	U_credit   int64 `thrift:",4,omitempty"`
	Music_data []any `thrift:",5,omitempty"` // TODO
}

type BuyVarietyStoreRetDataInfo struct {
	U_cp         int64  `thrift:",1,omitempty"`
	U_candy      int64  `thrift:",2,omitempty"`
	Reward_type  int16  `thrift:",3,omitempty"`
	Reward_id    int32  `thrift:",4,omitempty"`
	Reward_value int32  `thrift:",5,omitempty"`
	Status       string `thrift:",6,omitempty"`
}

type ChallengeMissionListInfo struct {
	Music_idx    int32 `thrift:",1,omitempty"`
	Difficulty   int16 `thrift:",2,omitempty"`
	Mission_list []any `thrift:",3,omitempty"` // TODO
}

type GetChThirdRetDataInfo struct {
	U_seq                        int32      `thrift:",1,omitempty"`
	User_ap                      UserApData `thrift:",2,omitempty"`
	User_ch_third_stage          []any      `thrift:",3,omitempty"` // TODO
	User_ch_third_chapter_reward []any      `thrift:",4,omitempty"` // TODO
}

type GetDiamondBonusRetDataInfo struct {
	Bonus1 int64 `thrift:",1,omitempty"`
	Bonus2 int64 `thrift:",2,omitempty"`
	Bonus3 int64 `thrift:",3,omitempty"`
	Bonus4 int64 `thrift:",4,omitempty"`
	Bonus5 int64 `thrift:",5,omitempty"`
	Bonus6 int64 `thrift:",6,omitempty"`
}

type GetMusicRewardRetDataInfo struct {
	Total_reward_value    int32       `thrift:",1,omitempty"`
	Reward_music_id       []any       `thrift:",2,omitempty"` // TODO
	Reward_value          []any       `thrift:",3,omitempty"` // TODO
	User_follower_profile []any       `thrift:",4,omitempty"` // TODO
	Error_data            map[any]any `thrift:",5,omitempty"` // TODO
}

type GetTotalMusicRankList struct {
	Rank     int32  `thrift:",1,omitempty"`
	Id       string `thrift:",2,omitempty"`
	Nick     string `thrift:",3,omitempty"`
	Country  string `thrift:",4,omitempty"`
	Score    int32  `thrift:",5,omitempty"`
	U_avatar int16  `thrift:",6,omitempty"`
}

type UpdateScoreDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Score_list  []any  `thrift:",5,omitempty"` // TODO
}

type GetChThirdDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type GetUserCollection struct {
	Call string                    `thrift:",1,omitempty"`
	Data GetUserCollectionDataInfo `thrift:",2,omitempty"`
}

type PlayMusicRetDataInfo struct {
	Code      int16   `thrift:",1,omitempty"`
	Album_idx int32   `thrift:",2,omitempty"`
	Music_idx int32   `thrift:",3,omitempty"`
	Score1    int32   `thrift:",4,omitempty"`
	Score2    int32   `thrift:",5,omitempty"`
	Score3    int32   `thrift:",6,omitempty"`
	Grade1    int16   `thrift:",7,omitempty"`
	Grade2    int16   `thrift:",8,omitempty"`
	Grade3    int16   `thrift:",9,omitempty"`
	Play_cnt1 int32   `thrift:",10,omitempty"`
	Play_cnt2 int32   `thrift:",11,omitempty"`
	Play_cnt3 int32   `thrift:",12,omitempty"`
	Buy_type  int16   `thrift:",13,omitempty"`
	Price     float64 `thrift:",14,omitempty"`
}

type SetEventRewardReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        SetEventRewardRetDataInfo    `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type HibernationLogReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        HibernationLogRetDataInfo    `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type ItemPurchase struct {
	Call     string               `thrift:",1,omitempty"`
	Data     ItemPurchaseDataInfo `thrift:",2,omitempty"`
	Sub_mode string               `thrift:",3,omitempty"`
}

type SetFollowerQuestInfinite struct {
	Call        string                           `thrift:",1,omitempty"`
	Data        SetFollowerQuestInfiniteDataInfo `thrift:",2,omitempty"`
	Common_data common_model.ParamData           `thrift:",3,omitempty"`
}

type SetUserFollowerProfileReward struct {
	Call        string                               `thrift:",1,omitempty"`
	Data        SetUserFollowerProfileRewardDataInfo `thrift:",2,omitempty"`
	Common_data common_model.ParamData               `thrift:",3,omitempty"`
}

type MoreGamesLogRetDataInfo struct {
	Ret string `thrift:",1,omitempty"`
}

type DeleteUserDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type GetDiamondBonus struct {
	Call string                  `thrift:",1,omitempty"`
	Data GetDiamondBonusDataInfo `thrift:",2,omitempty"`
}

type HibernationLogRetDataInfo struct {
	Action int32 `thrift:",1,omitempty"`
}

type ReviewDataList struct {
	Music_idx  int32   `thrift:",1,omitempty"`
	Point      float64 `thrift:",2,omitempty"`
	Difficulty int16   `thrift:",3,omitempty"`
}

type SetBookmarkDataInfo struct {
	U_seq         int32  `thrift:",1,omitempty"`
	U_id          string `thrift:",2,omitempty"`
	Uuid          string `thrift:",3,omitempty"`
	Device_uuid   string `thrift:",4,omitempty"`
	Bookmark_list []any  `thrift:",5,omitempty"` // TODO
}

type TransferUserReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        TransferUserRetDataInfo      `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type UpdateNickNameDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Nickname    string `thrift:",5,omitempty"`
}

type GetChoiceUserDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Choice_uuid string `thrift:",5,omitempty"`
}

type GetUserCollectionReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        []any                        `thrift:",4,omitempty"` // TODO
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type SaveUserCharacter struct {
	I_id         int64 `thrift:",1,omitempty"`
	I_Level      int64 `thrift:",2,omitempty"`
	I_BonusLevel int64 `thrift:",3,omitempty"`
}

type SetTutorialNewDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	I_ids       []any  `thrift:",5,omitempty"` // TODO
}

type GetChThird struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        GetChThirdDataInfo     `thrift:",2,omitempty"`
	Common_data common_model.ParamData `thrift:",3,omitempty"`
}

type MusicPointReviewReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        []any                        `thrift:",4,omitempty"` // TODO
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type SetAttendance struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        SetAttendanceDataInfo  `thrift:",2,omitempty"`
	Common_data common_model.ParamData `thrift:",3,omitempty"`
}

type UserAvatarListRetDataInfo struct {
	U_seq  int32 `thrift:",1,omitempty"`
	Avatar []any `thrift:",2,omitempty"` // TODO
}

type UserLoginDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	U_state     string `thrift:",5,omitempty"`
}

type AdViewLogRetDataInfo struct {
	Status string `thrift:",1,omitempty"`
}

type BuyCollectionItemRetDataInfo struct {
	Item_type    int16  `thrift:",1,omitempty"`
	Item_value   int64  `thrift:",2,omitempty"`
	Buy_datetime string `thrift:",3,omitempty"`
}

type ChThirdStageRetDataInfo struct {
	Star                        int32               `thrift:",1,omitempty"`
	Character_score             int32               `thrift:",2,omitempty"`
	Music_score                 int32               `thrift:",3,omitempty"`
	Follower_profile_score      int32               `thrift:",4,omitempty"`
	Bonus_score                 int32               `thrift:",5,omitempty"`
	Total_score                 int64               `thrift:",6,omitempty"`
	Clear                       int32               `thrift:",7,omitempty"`
	User_ap                     UserApData          `thrift:",8,omitempty"`
	User_ch_third_stage         UserChThirdStage    `thrift:",9,omitempty"`
	User_music                  UserMusic           `thrift:",10,omitempty"`
	Reward_data                 []any               `thrift:",11,omitempty"` // TODO
	Bonus_follower_profile_ids  []any               `thrift:",12,omitempty"` // TODO
	User_follower_profile       UserFollowerProfile `thrift:",13,omitempty"`
	Bonus_music_score           int32               `thrift:",14,omitempty"`
	User_follower_profile_score []any               `thrift:",15,omitempty"` // TODO
}

type CheckPurchasedDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type GetChThirdStarReward struct {
	Call        string                       `thrift:",1,omitempty"`
	Data        GetChThirdStarRewardDataInfo `thrift:",2,omitempty"`
	Common_data common_model.ParamData       `thrift:",3,omitempty"`
}

type SetReviewPopup struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        SetReviewPopupDataInfo `thrift:",2,omitempty"`
	Common_data common_model.ParamData `thrift:",3,omitempty"`
}

type SetTutorialDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Step        string `thrift:",5,omitempty"`
}

type UserEventPoint struct {
	S_EventType  string `thrift:",1,omitempty"`
	I_DataID     int64  `thrift:",2,omitempty"`
	I_Point      int64  `thrift:",3,omitempty"`
	I_Step       int64  `thrift:",4,omitempty"`
	I_ADViewTime int64  `thrift:",5,omitempty"`
	I_Version    int32  `thrift:",6,omitempty"`
}

type BuyItemDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Idx         int16  `thrift:",5,omitempty"`
}

type SetGameRewardReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        SetGameRewardRetDataInfo     `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type UserLoginReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        UserLoginRetDataInfo         `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type BuyAvatarReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        BuyAvatarRetDataInfo         `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type GetDiamondBonusDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type SaveUserFollower struct {
	I_id         int64 `thrift:",1,omitempty"`
	I_Level      int64 `thrift:",2,omitempty"`
	I_BonusLevel int64 `thrift:",3,omitempty"`
}

type SetTutorial struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        SetTutorialDataInfo    `thrift:",2,omitempty"`
	Common_data common_model.ParamData `thrift:",3,omitempty"`
}

type UpdateNickName struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        UpdateNickNameDataInfo `thrift:",2,omitempty"`
	Common_data common_model.ParamData `thrift:",3,omitempty"`
}

type GetMusicRewardDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	I_ids       []any  `thrift:",5,omitempty"` // TODO
	I_levels    []any  `thrift:",6,omitempty"` // TODO
}

type PaidEventPointReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        PaidEventPointRetDataInfo    `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type UserScore struct {
	Music_idx          int32                        `thrift:",1,omitempty"`
	Score1             int64                        `thrift:",2,omitempty"`
	Score2             int64                        `thrift:",3,omitempty"`
	Score3             int64                        `thrift:",4,omitempty"`
	Score4             int64                        `thrift:",5,omitempty"`
	Grade1             int16                        `thrift:",6,omitempty"`
	Grade2             int16                        `thrift:",7,omitempty"`
	Grade3             int16                        `thrift:",8,omitempty"`
	Grade4             int16                        `thrift:",9,omitempty"`
	Play_cnt1          int64                        `thrift:",10,omitempty"`
	Play_cnt2          int64                        `thrift:",11,omitempty"`
	Play_cnt3          int64                        `thrift:",12,omitempty"`
	Play_cnt4          int64                        `thrift:",13,omitempty"`
	Combo_grade1       int16                        `thrift:",14,omitempty"`
	Combo_grade2       int16                        `thrift:",15,omitempty"`
	Combo_grade3       int16                        `thrift:",16,omitempty"`
	Combo_grade4       int16                        `thrift:",17,omitempty"`
	Multi_score1       int64                        `thrift:",18,omitempty"`
	Multi_score2       int64                        `thrift:",19,omitempty"`
	Multi_score3       int64                        `thrift:",20,omitempty"`
	Multi_score4       int64                        `thrift:",21,omitempty"`
	Multi_play_cnt1    int64                        `thrift:",22,omitempty"`
	Multi_play_cnt2    int64                        `thrift:",23,omitempty"`
	Multi_play_cnt3    int64                        `thrift:",24,omitempty"`
	Multi_play_cnt4    int64                        `thrift:",25,omitempty"`
	Multi_combo_grade1 int16                        `thrift:",26,omitempty"`
	Multi_combo_grade2 int16                        `thrift:",27,omitempty"`
	Multi_combo_grade3 int16                        `thrift:",28,omitempty"`
	Multi_combo_grade4 int16                        `thrift:",29,omitempty"`
	Purchase_type      int16                        `thrift:",30,omitempty"`
	Is_extra           int16                        `thrift:",31,omitempty"`
	Review_point       int16                        `thrift:",32,omitempty"`
	Rank_list          GetTotalMusicRankRetDataInfo `thrift:",33,omitempty"`
}

type BuyPieceDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Idx         int32  `thrift:",5,omitempty"`
	Cnt         int64  `thrift:",6,omitempty"`
}

type GetMusicReward struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        GetMusicRewardDataInfo `thrift:",2,omitempty"`
	Common_data common_model.ParamData `thrift:",3,omitempty"`
}

type GetUserInfoRetDataInfo struct {
	U_cp    int64   `thrift:",1,omitempty"`
	U_candy float64 `thrift:",2,omitempty"`
	U_like  float64 `thrift:",3,omitempty"`
	U_fans  int64   `thrift:",4,omitempty"`
}

type ItemPurchaseReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        ItemPurchaseRetDataInfo      `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type SetSelectRewardReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        SetSelectRewardRetDataInfo   `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type SetTutorialNew struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        SetTutorialNewDataInfo `thrift:",2,omitempty"`
	Common_data common_model.ParamData `thrift:",3,omitempty"`
}

type UpdateAchievementRetDataInfo struct {
	Achievement int16 `thrift:",1,omitempty"`
}

type UserDelReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        UserDelRetDataInfo           `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type BuyCheckReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        BuyCheckRetDataInfo          `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type UserDelRetDataInfo struct {
	Result string `thrift:",1,omitempty"`
}

type UserJoinDataInfo struct {
	Os           int16  `thrift:",1,omitempty"`
	Join_type    int16  `thrift:",2,omitempty"`
	Uuid         string `thrift:",3,omitempty"`
	Device_uuid  string `thrift:",4,omitempty"`
	Device_token string `thrift:",5,omitempty"`
	Sns_id       string `thrift:",6,omitempty"`
	Access_token string `thrift:",7,omitempty"`
	Country      string `thrift:",8,omitempty"`
	State        string `thrift:",9,omitempty"`
}

type UserPurchaseMusicIdx struct {
	Music_idx int32 `thrift:",1,omitempty"`
}

type UserPurchaseReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        UserPurchaseRetDataInfo      `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type UserSubscribePassReward struct {
	I_SubscribeID int64 `thrift:",1,omitempty"`
	I_Type        int64 `thrift:",2,omitempty"`
	I_Step        int64 `thrift:",3,omitempty"`
	I_UpdateTime  int64 `thrift:",4,omitempty"`
	I_Version     int32 `thrift:",5,omitempty"`
}

type BuyCollectionItemReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        BuyCollectionItemRetDataInfo `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type CheckBuyShop struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        CheckBuyShopDataInfo   `thrift:",2,omitempty"`
	Common_data common_model.ParamData `thrift:",3,omitempty"`
}

type UserLoadDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Type        string `thrift:",5,omitempty"`
}

type BuyPieceReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        BuyPieceRetDataInfo          `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type ChoiceUserData struct {
	U_girl_level int64   `thrift:",1,omitempty"`
	U_fans       int64   `thrift:",2,omitempty"`
	U_fans_grade int16   `thrift:",3,omitempty"`
	U_cp         int64   `thrift:",4,omitempty"`
	U_candy      float64 `thrift:",5,omitempty"`
	U_like       float64 `thrift:",6,omitempty"`
	U_last_login string  `thrift:",7,omitempty"`
}

type CompleteLog struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        CompleteLogDataInfo    `thrift:",2,omitempty"`
	Common_data common_model.ParamData `thrift:",3,omitempty"`
}

type LastSaveTimeRetDataInfo struct {
	Last_save_time int64  `thrift:",1,omitempty"`
	Device_uuid    string `thrift:",2,omitempty"`
}

type UpdateScore struct {
	Call string              `thrift:",1,omitempty"`
	Data UpdateScoreDataInfo `thrift:",2,omitempty"`
}

type UserJoin struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        UserJoinDataInfo       `thrift:",2,omitempty"`
	Common_data common_model.ParamData `thrift:",3,omitempty"`
}

type BuyCollectionItemDataInfo struct {
	U_seq             int32  `thrift:",1,omitempty"`
	U_id              string `thrift:",2,omitempty"`
	Uuid              string `thrift:",3,omitempty"`
	Device_uuid       string `thrift:",4,omitempty"`
	Special_music_idx int32  `thrift:",5,omitempty"`
	Idx               int32  `thrift:",6,omitempty"`
}

type ChThirdStage struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        ChThirdStageDataInfo   `thrift:",2,omitempty"`
	Common_data common_model.ParamData `thrift:",3,omitempty"`
}

type GetChThirdStarRewardDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	I_id        int32  `thrift:",5,omitempty"`
	Reward_num  int32  `thrift:",6,omitempty"`
}

type GetUserInfoDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type UpdateScoreDataList struct {
	Music_idx   int32 `thrift:",1,omitempty"`
	Score       int32 `thrift:",2,omitempty"`
	Grade       int16 `thrift:",3,omitempty"`
	Play_cnt    int32 `thrift:",4,omitempty"`
	Difficulty  int16 `thrift:",5,omitempty"`
	Combo_grade int16 `thrift:",6,omitempty"`
	Collection  []any `thrift:",7,omitempty"` // TODO
}

type UpdateScoreRetDataInfo struct {
	Update_list []any `thrift:",1,omitempty"` // TODO
	U_gold      int64 `thrift:",2,omitempty"`
	U_ep        int64 `thrift:",3,omitempty"`
	Event_flg   int16 `thrift:",4,omitempty"`
}

type PaidEventPointDataInfo struct {
	U_seq         int32  `thrift:",1,omitempty"`
	U_id          string `thrift:",2,omitempty"`
	Uuid          string `thrift:",3,omitempty"`
	Device_uuid   string `thrift:",4,omitempty"`
	Type          int16  `thrift:",5,omitempty"`
	I_SubscribeID int64  `thrift:",6,omitempty"`
	I_Version     int32  `thrift:",7,omitempty"`
}

type SaveUserAchievement struct {
	I_id       int64   `thrift:",1,omitempty"`
	D_Quantity float64 `thrift:",2,omitempty"`
	S_Quantity string  `thrift:",3,omitempty"`
}

type ScoreDataList struct {
	Music_idx     int32   `thrift:",1,omitempty"`
	Score         int32   `thrift:",2,omitempty"`
	Grade         int16   `thrift:",3,omitempty"`
	Play_cnt      int32   `thrift:",4,omitempty"`
	Difficulty    int16   `thrift:",5,omitempty"`
	Multiple      float64 `thrift:",6,omitempty"`
	Gp            int32   `thrift:",7,omitempty"`
	Achievement   int16   `thrift:",8,omitempty"`
	Play_type     int16   `thrift:",9,omitempty"`
	Play_datetime string  `thrift:",10,omitempty"`
	Mission_clear int16   `thrift:",11,omitempty"`
	Combo_grade   int16   `thrift:",12,omitempty"`
	Is_credit     int16   `thrift:",13,omitempty"`
}

type SetFollowerQuestInfiniteReturn struct {
	Error       common_model.ErrorRetCode           `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet          `thrift:",2,omitempty"`
	Mode        string                              `thrift:",3,omitempty"`
	Call        string                              `thrift:",4,omitempty"`
	Data        SetFollowerQuestInfiniteRetDataInfo `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData        `thrift:",6,omitempty"`
}

type SetSelectRewardDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	I_id        int32  `thrift:",5,omitempty"`
}

type SetSubscribe struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        SetSubscribeDataInfo   `thrift:",2,omitempty"`
	Common_data common_model.ParamData `thrift:",3,omitempty"`
}

type TransferUserRetDataInfo struct {
	Transfer_id string `thrift:",1,omitempty"`
}

type UserJoinRetDataInfo struct {
	U_seq int32  `thrift:",1,omitempty"`
	U_id  string `thrift:",2,omitempty"`
}

type UserPurchaseDataInfo struct {
	U_seq             int32  `thrift:",1,omitempty"`
	U_id              string `thrift:",2,omitempty"`
	Uuid              string `thrift:",3,omitempty"`
	Device_uuid       string `thrift:",4,omitempty"`
	Developer_payload string `thrift:",5,omitempty"`
	Pay_id            string `thrift:",6,omitempty"`
	Purchase_token    string `thrift:",7,omitempty"`
	Purchase_id       string `thrift:",8,omitempty"`
	Os                int16  `thrift:",9,omitempty"`
	Restore           int16  `thrift:",10,omitempty"`
}

type UserTitleListReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        UserTitleListRetDataInfo     `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type UserChThirdStage struct {
	I_id         int32 `thrift:",1,omitempty"`
	I_ChapterId  int32 `thrift:",2,omitempty"`
	I_StageIndex int32 `thrift:",3,omitempty"`
	I_Star       int32 `thrift:",4,omitempty"`
}

type CheckUser struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        CheckUserDataInfo      `thrift:",2,omitempty"`
	Common_data common_model.ParamData `thrift:",3,omitempty"`
}

type MoreGamesLog struct {
	Call string               `thrift:",1,omitempty"`
	Data MoreGamesLogDataInfo `thrift:",2,omitempty"`
}

type Bookmark_list struct {
	Music_idx int32 `thrift:",1,omitempty"`
	Flag      int16 `thrift:",2,omitempty"`
}

type BuyUserTitleDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Idx         int16  `thrift:",5,omitempty"`
}

type GetAllPointReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        GetAllPointRetDataInfo       `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type SetTutorialRetDataInfo struct {
	Step string `thrift:",1,omitempty"`
}

type UpdateNickNameReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        UpdateNickNameRetDataInfo    `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type UpdateScoreReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        UpdateScoreRetDataInfo       `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type UserGuitar struct {
	I_id         int64 `thrift:",1,omitempty"`
	I_Level      int64 `thrift:",2,omitempty"`
	I_BonusLevel int64 `thrift:",3,omitempty"`
}

type StageFollowerProfileScore struct {
	I_id        int32 `thrift:",1,omitempty"`
	Score       int32 `thrift:",2,omitempty"`
	Bonus_score int32 `thrift:",3,omitempty"`
}

type BuyMusicRetDataInfo struct {
	U_gold int32       `thrift:",1,omitempty"`
	U_cp   int32       `thrift:",2,omitempty"`
	U_mp   int32       `thrift:",3,omitempty"`
	Music  map[any]any `thrift:",4,omitempty"` // TODO
}

type GetProfileReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        GetProfileRetDataInfo        `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type HibernationLog struct {
	Call string                 `thrift:",1,omitempty"`
	Data HibernationLogDataInfo `thrift:",2,omitempty"`
}

type MusicData struct {
	Music_idx int32 `thrift:",1,omitempty"`
	Score1    int32 `thrift:",2,omitempty"`
	Score2    int32 `thrift:",3,omitempty"`
	Score3    int32 `thrift:",4,omitempty"`
}

type BuyMusic struct {
	Call string           `thrift:",1,omitempty"`
	Data BuyMusicDataInfo `thrift:",2,omitempty"`
}

type GetTransferId struct {
	Call string                `thrift:",1,omitempty"`
	Data GetTransferIdDataInfo `thrift:",2,omitempty"`
}

type SetSelectReward struct {
	Call        string                  `thrift:",1,omitempty"`
	Data        SetSelectRewardDataInfo `thrift:",2,omitempty"`
	Common_data common_model.ParamData  `thrift:",3,omitempty"`
}

type SetSubscribeReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        SetSubscribeRetDataInfo      `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type UserProp struct {
	I_id    int64 `thrift:",1,omitempty"`
	I_Level int64 `thrift:",2,omitempty"`
}

type UserFollowerProfile struct {
	I_id       int32 `thrift:",1,omitempty"`
	I_Level    int32 `thrift:",2,omitempty"`
	D_Exp      int64 `thrift:",3,omitempty"`
	I_AddCandy int32 `thrift:",4,omitempty"`
}

type CheckBuyProductRetDataInfo struct {
	Is_owner int16 `thrift:",1,omitempty"`
}

type CheckBuyShopRetDataInfo struct {
	Is_owner int16 `thrift:",1,omitempty"`
}

type ItemPurchaseDataInfo struct {
	U_seq             int32  `thrift:",1,omitempty"`
	U_id              string `thrift:",2,omitempty"`
	Uuid              string `thrift:",3,omitempty"`
	Device_uuid       string `thrift:",4,omitempty"`
	Developer_payload string `thrift:",5,omitempty"`
	Pay_id            string `thrift:",6,omitempty"`
	Purchase_token    string `thrift:",7,omitempty"`
	Purchase_id       string `thrift:",8,omitempty"`
	Os                int16  `thrift:",9,omitempty"`
	Restore           int16  `thrift:",10,omitempty"`
}

type UseCouponRetDataInfo struct {
	Ret string `thrift:",1,omitempty"`
}

type UserAvatarListDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type UserLoadReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        UserLoadRetDataInfo          `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type UserSkill struct {
	I_id               int64 `thrift:",1,omitempty"`
	I_Level            int64 `thrift:",2,omitempty"`
	B_Activate         int16 `thrift:",3,omitempty"`
	L_ActivateOnTicks  int64 `thrift:",4,omitempty"`
	L_ActivateOffTicks int64 `thrift:",5,omitempty"`
}

type BuyItemReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        BuyItemRetDataInfo           `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type LogCashCompleteReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        LogCashCompleteRetDataInfo   `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type SaveUserInfo struct {
	U_like                float64 `thrift:",1,omitempty"`
	U_fans                int64   `thrift:",2,omitempty"`
	U_fans_grade          int16   `thrift:",3,omitempty"`
	U_selected_costume_id int64   `thrift:",4,omitempty"`
	U_selected_music_id   int64   `thrift:",5,omitempty"`
}

type SetPassRewardRetDataInfo struct {
	Subscribe_pass_reward UserSubscribePassReward `thrift:",1,omitempty"`
	Reward_data           []any                   `thrift:",2,omitempty"` // TODO
}

type SetReviewPopupRetDataInfo struct {
	Status string `thrift:",1,omitempty"`
}

type UserPurchaseRetDataInfo struct {
	Pay_id      string `thrift:",1,omitempty"`
	Product_id  string `thrift:",2,omitempty"`
	Purchase_id string `thrift:",3,omitempty"`
	Music_data  []any  `thrift:",4,omitempty"` // TODO
}

type BuyCheck struct {
	Call string           `thrift:",1,omitempty"`
	Data BuyCheckDataInfo `thrift:",2,omitempty"`
}

type GetProfile struct {
	Call string             `thrift:",1,omitempty"`
	Data GetProfileDataInfo `thrift:",2,omitempty"`
}

type GetTransferIdRetDataInfo struct {
	Transfer_id string `thrift:",1,omitempty"`
}

type LastSaveTimeReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        LastSaveTimeRetDataInfo      `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type SetUserGPDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Gp          int32  `thrift:",5,omitempty"`
}

type BuyUserTitleRetDataInfo struct {
	Title_idx int16 `thrift:",1,omitempty"`
	U_gold    int64 `thrift:",2,omitempty"`
	U_cp      int64 `thrift:",3,omitempty"`
	U_mp      int64 `thrift:",4,omitempty"`
}

type ChangeUser struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        ChangeUserDataInfo     `thrift:",2,omitempty"`
	Common_data common_model.ParamData `thrift:",3,omitempty"`
}

type CompleteLogReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        CompleteLogRetDataInfo       `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type CheckPurchased struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        CheckPurchasedDataInfo `thrift:",2,omitempty"`
	Common_data common_model.ParamData `thrift:",3,omitempty"`
}

type SetBookmarkDataList struct {
	Music_idx int32 `thrift:",1,omitempty"`
	Flag      int16 `thrift:",2,omitempty"`
}

type SetUserFollowerQuestReturn struct {
	Error       common_model.ErrorRetCode       `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet      `thrift:",2,omitempty"`
	Mode        string                          `thrift:",3,omitempty"`
	Call        string                          `thrift:",4,omitempty"`
	Data        SetUserFollowerQuestRetDataInfo `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData    `thrift:",6,omitempty"`
}

type UpdateUserTitle struct {
	Call string                  `thrift:",1,omitempty"`
	Data UpdateUserTitleDataInfo `thrift:",2,omitempty"`
}

type UserTutorial struct {
	I_id int32 `thrift:",1,omitempty"`
}

type BuyVarietyStoreReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        BuyVarietyStoreRetDataInfo   `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type SaveUserCostume struct {
	I_id         int64 `thrift:",1,omitempty"`
	I_Level      int64 `thrift:",2,omitempty"`
	I_BonusLevel int64 `thrift:",3,omitempty"`
}

type SetFollowerQuestInfiniteDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type SetUserFollowerProfileRewardDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	I_id        int32  `thrift:",5,omitempty"`
	S_level     string `thrift:",6,omitempty"`
}

type SetUserFollowerQuestDataInfo struct {
	U_seq            int32   `thrift:",1,omitempty"`
	U_id             string  `thrift:",2,omitempty"`
	Uuid             string  `thrift:",3,omitempty"`
	Device_uuid      string  `thrift:",4,omitempty"`
	I_CurrentID      int32   `thrift:",5,omitempty"`
	I_SubID          int32   `thrift:",6,omitempty"`
	D_ConditionValue float64 `thrift:",7,omitempty"`
}

type UserAreaData struct {
	U_area_num          int16   `thrift:",1,omitempty"`
	D_Candy             float64 `thrift:",2,omitempty"`
	D_Like              float64 `thrift:",3,omitempty"`
	I_UserFanCount      int64   `thrift:",4,omitempty"`
	I_UserFanGrade      int16   `thrift:",5,omitempty"`
	I_SelectedCostumeId int64   `thrift:",6,omitempty"`
	I_SelectedMusicId   int64   `thrift:",7,omitempty"`
	I_SelectedGuitarId  int64   `thrift:",8,omitempty"`
	S_TutorialList      string  `thrift:",9,omitempty"`
	S_Gp1               string  `thrift:",10,omitempty"`
	S_Gp2               string  `thrift:",11,omitempty"`
}

type BuyCollectionMusicRetDataInfo struct {
	Music_idx    int64  `thrift:",1,omitempty"`
	Buy_datetime string `thrift:",2,omitempty"`
}

type CheckBuyProductDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Product_id  string `thrift:",5,omitempty"`
}

type PlayMusicDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Music_idx   int32  `thrift:",5,omitempty"`
}

type SetAdRewardRetDataInfo struct {
	I_id                  int32               `thrift:",1,omitempty"`
	User_ad_list          UserAdList          `thrift:",2,omitempty"`
	Reward_data           []any               `thrift:",3,omitempty"` // TODO
	User_follower_profile UserFollowerProfile `thrift:",4,omitempty"`
}

type SetSelectRewardRetDataInfo struct {
	U_seq                   int32 `thrift:",1,omitempty"`
	User_select_reward_list []any `thrift:",2,omitempty"` // TODO
}

type SetUserGPRetDataInfo struct {
	Gp int32 `thrift:",1,omitempty"`
}

type UserTicketCollection struct {
	I_id int64 `thrift:",1,omitempty"`
}

type BuyCheckRetDataInfo struct {
	Result int64 `thrift:",1,omitempty"`
}

type BuyPackage struct {
	Call string             `thrift:",1,omitempty"`
	Data BuyPackageDataInfo `thrift:",2,omitempty"`
}

type CompleteLogRetDataInfo struct {
	Status string `thrift:",1,omitempty"`
}

type SaveUserMusic struct {
	I_id                    int64 `thrift:",1,omitempty"`
	I_Level                 int64 `thrift:",2,omitempty"`
	I_BonusLevel            int64 `thrift:",3,omitempty"`
	B_EncoreBonusAppear     int64 `thrift:",4,omitempty"`
	I_EncoreBonusFollowerId int64 `thrift:",5,omitempty"`
}

type SetUserGP struct {
	Call string            `thrift:",1,omitempty"`
	Data SetUserGPDataInfo `thrift:",2,omitempty"`
}

type UseCreditReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        UseCreditRetDataInfo         `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type UserDelDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type MoreGamesLogReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        MoreGamesLogRetDataInfo      `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type BuyMusicDataInfo struct {
	Type          int16  `thrift:",1,omitempty"`
	Idx           int32  `thrift:",2,omitempty"`
	U_seq         int32  `thrift:",3,omitempty"`
	U_id          string `thrift:",4,omitempty"`
	Uuid          string `thrift:",5,omitempty"`
	Device_uuid   string `thrift:",6,omitempty"`
	Purchase_type int16  `thrift:",7,omitempty"`
}

type MusicPointReviewDataInfo struct {
	U_seq        int32  `thrift:",1,omitempty"`
	U_id         string `thrift:",2,omitempty"`
	Uuid         string `thrift:",3,omitempty"`
	Device_uuid  string `thrift:",4,omitempty"`
	Review_point []any  `thrift:",5,omitempty"` // TODO
}

type SetEventRewardItem struct {
	Item_type int32 `thrift:",1,omitempty"`
	Item_id   int32 `thrift:",2,omitempty"`
	Count     int64 `thrift:",3,omitempty"`
}

type UserFollowerProfileReward struct {
	I_id          int32 `thrift:",1,omitempty"`
	I_RewardLevel int32 `thrift:",2,omitempty"`
}

type ChallengeMissionListReturn struct {
	Error       common_model.ErrorRetCode       `thrift:",1,omitempty"`
	Mode        string                          `thrift:",2,omitempty"`
	Call        string                          `thrift:",3,omitempty"`
	Data        ChallengeMissionListRetDataInfo `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData    `thrift:",5,omitempty"`
}

type ChangeUserRetDataInfo struct {
	Result string `thrift:",1,omitempty"`
}

type CheckUserDataInfo struct {
	Uuid        string `thrift:",1,omitempty"`
	Device_uuid string `thrift:",2,omitempty"`
}

type GetChoiceUser struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        GetChoiceUserDataInfo  `thrift:",2,omitempty"`
	Common_data common_model.ParamData `thrift:",3,omitempty"`
}

type GetTransferIdDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type LastSaveTimeDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type SetUserGPReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        SetUserGPRetDataInfo         `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type UserSelectReward struct {
	I_GroupId          int32 `thrift:",1,omitempty"`
	I_RewardGroupId    int32 `thrift:",2,omitempty"`
	I_AltRewardGroupId int32 `thrift:",3,omitempty"`
}

type BuyPackageReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        BuyPackageRetDataInfo        `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type CheckPurchasedReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        CheckPurchasedRetDataInfo    `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type CompleteLogDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Type        string `thrift:",5,omitempty"`
	Id          int32  `thrift:",6,omitempty"`
	Value       string `thrift:",7,omitempty"`
}

type GetPurchaseDiamondLog struct {
	Call string                        `thrift:",1,omitempty"`
	Data GetPurchaseDiamondLogDataInfo `thrift:",2,omitempty"`
}

type SetSelectRewardInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	I_id        int32  `thrift:",5,omitempty"`
}

type UpdateUserTitleDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	U_title     int16  `thrift:",5,omitempty"`
}

type UseCouponDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Send_ppid   string `thrift:",5,omitempty"`
	Coupon      string `thrift:",6,omitempty"`
}

type BuyAvatar struct {
	Call string            `thrift:",1,omitempty"`
	Data BuyAvatarDataInfo `thrift:",2,omitempty"`
}

type BuyAvatarRetDataInfo struct {
	Avatar_idx int16 `thrift:",1,omitempty"`
	U_gold     int64 `thrift:",2,omitempty"`
	U_cp       int64 `thrift:",3,omitempty"`
	U_mp       int64 `thrift:",4,omitempty"`
}

type BuyMusicReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        BuyMusicRetDataInfo          `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type GetPurchaseDiamondLogDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type SetEventRewardDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Event_idx   int32  `thrift:",5,omitempty"`
}

type SetEventRewardRetDataInfo struct {
	U_cp         int64   `thrift:",1,omitempty"`
	U_candy      float64 `thrift:",2,omitempty"`
	U_like       float64 `thrift:",3,omitempty"`
	U_fans       int64   `thrift:",4,omitempty"`
	Reward_type  int16   `thrift:",5,omitempty"`
	Reward_id    int32   `thrift:",6,omitempty"`
	Reward_value int32   `thrift:",7,omitempty"`
	Status       string  `thrift:",8,omitempty"`
	Reward_data  []any   `thrift:",9,omitempty"` // TODO
}

type SetGameRewardDataInfo struct {
	U_seq       int32   `thrift:",1,omitempty"`
	U_id        string  `thrift:",2,omitempty"`
	Uuid        string  `thrift:",3,omitempty"`
	Device_uuid string  `thrift:",4,omitempty"`
	Type        string  `thrift:",5,omitempty"`
	Id          int32   `thrift:",6,omitempty"`
	Level       int16   `thrift:",7,omitempty"`
	Quantity    float64 `thrift:",8,omitempty"`
	S_quantity  string  `thrift:",9,omitempty"`
}

type SetSubscribeDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	I_ids       []any  `thrift:",5,omitempty"` // TODO
}

type BuyItem struct {
	Call string          `thrift:",1,omitempty"`
	Data BuyItemDataInfo `thrift:",2,omitempty"`
}

type GetProfileMusicDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Get_u_seq   int64  `thrift:",5,omitempty"`
	Start_limit int64  `thrift:",6,omitempty"`
	End_limit   int64  `thrift:",7,omitempty"`
}

type SetTutorialNewReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        SetTutorialNewRetDataInfo    `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type SetUserFollowerQuest struct {
	Call        string                       `thrift:",1,omitempty"`
	Data        SetUserFollowerQuestDataInfo `thrift:",2,omitempty"`
	Common_data common_model.ParamData       `thrift:",3,omitempty"`
}

type UserContentsData struct {
	User_achievement             []any `thrift:",1,omitempty"`  // TODO
	User_buff                    []any `thrift:",2,omitempty"`  // TODO
	User_candy_shop              []any `thrift:",3,omitempty"`  // TODO
	User_character               []any `thrift:",4,omitempty"`  // TODO
	User_costume                 []any `thrift:",5,omitempty"`  // TODO
	User_daily_mission           []any `thrift:",6,omitempty"`  // TODO
	User_follower                []any `thrift:",7,omitempty"`  // TODO
	User_music                   []any `thrift:",8,omitempty"`  // TODO
	User_prop                    []any `thrift:",9,omitempty"`  // TODO
	User_unit                    []any `thrift:",10,omitempty"` // TODO
	User_skill                   []any `thrift:",11,omitempty"` // TODO
	User_shop                    []any `thrift:",12,omitempty"` // TODO
	User_messenger               []any `thrift:",13,omitempty"` // TODO
	User_guitar                  []any `thrift:",14,omitempty"` // TODO
	User_event_point             []any `thrift:",15,omitempty"` // TODO
	User_subscribe_list          []any `thrift:",16,omitempty"` // TODO
	User_subscribe_pass_reward   []any `thrift:",17,omitempty"` // TODO
	User_ticketcollection        []any `thrift:",18,omitempty"` // TODO
	User_follower_quest          []any `thrift:",19,omitempty"` // TODO
	User_follower_profile_reward []any `thrift:",20,omitempty"` // TODO
	User_follower_profile        []any `thrift:",21,omitempty"` // TODO
	User_follower_giftitem       []any `thrift:",22,omitempty"` // TODO
	User_ad_list                 []any `thrift:",23,omitempty"` // TODO
	User_chthird_stage           []any `thrift:",24,omitempty"` // TODO
	User_tutorial                []any `thrift:",25,omitempty"` // TODO
	User_ad_level                []any `thrift:",26,omitempty"` // TODO
	User_select_reward           []any `thrift:",27,omitempty"` // TODO
}

type UserFollowerQuest struct {
	I_id              int64   `thrift:",1,omitempty"`
	I_CurrentID       int64   `thrift:",2,omitempty"`
	I_CompleteID      int64   `thrift:",3,omitempty"`
	D_ConditionValue1 float64 `thrift:",4,omitempty"`
	D_ConditionValue2 float64 `thrift:",5,omitempty"`
	D_ConditionValue3 float64 `thrift:",6,omitempty"`
	I_RewardReceived1 int16   `thrift:",7,omitempty"`
	I_RewardReceived2 int16   `thrift:",8,omitempty"`
	I_RewardReceived3 int16   `thrift:",9,omitempty"`
	I_isInfinity      int16   `thrift:",10,omitempty"`
}

type UserJoinReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        UserJoinRetDataInfo          `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type UserLoginRetDataInfo struct {
	User          UserData         `thrift:",1,omitempty"`
	Area_data     map[any]any      `thrift:",2,omitempty"` // TODO
	User_contents UserContentsData `thrift:",3,omitempty"`
}

type BuyUserTitleReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        BuyUserTitleRetDataInfo      `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type GetTotalMusicMyRank struct {
	Grade string `thrift:",1,omitempty"`
	Rank  int32  `thrift:",2,omitempty"`
	Score int32  `thrift:",3,omitempty"`
}

type SetAttendanceDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Type        string `thrift:",5,omitempty"`
}

type SetFollowerProfileGiftRetDataInfo struct {
	I_gift_type            int32                `thrift:",1,omitempty"`
	User_follower_giftitem UserFollowerGiftItem `thrift:",2,omitempty"`
	User_follower_profile  UserFollowerProfile  `thrift:",3,omitempty"`
}

type SetReviewPopupReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        SetReviewPopupRetDataInfo    `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type UserAvatarList struct {
	Call string                 `thrift:",1,omitempty"`
	Data UserAvatarListDataInfo `thrift:",2,omitempty"`
}

type UserCostume struct {
	I_id         int64 `thrift:",1,omitempty"`
	I_Level      int64 `thrift:",2,omitempty"`
	I_BonusLevel int64 `thrift:",3,omitempty"`
}

type UserLogin struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        UserLoginDataInfo      `thrift:",2,omitempty"`
	Common_data common_model.ParamData `thrift:",3,omitempty"`
}

type SetFollowerProfileGiftDataInfo struct {
	U_seq          int32  `thrift:",1,omitempty"`
	U_id           string `thrift:",2,omitempty"`
	Uuid           string `thrift:",3,omitempty"`
	Device_uuid    string `thrift:",4,omitempty"`
	Profile_id     int32  `thrift:",5,omitempty"`
	Gift_id        int32  `thrift:",6,omitempty"`
	Use_gitf_value int32  `thrift:",7,omitempty"`
}

type SetGameRewardRetDataInfo struct {
	Type                  string              `thrift:",1,omitempty"`
	Id                    int32               `thrift:",2,omitempty"`
	Level                 int16               `thrift:",3,omitempty"`
	Reward_type           string              `thrift:",4,omitempty"`
	Reward_value          int64               `thrift:",5,omitempty"`
	Status                string              `thrift:",6,omitempty"`
	User_follower_profile UserFollowerProfile `thrift:",7,omitempty"`
}

type UserCandyShop struct {
	I_id              int64   `thrift:",1,omitempty"`
	I_CurrentBuyCount int64   `thrift:",2,omitempty"`
	I_TotalBuyCount   int64   `thrift:",3,omitempty"`
	L_LastBuyTick     float64 `thrift:",4,omitempty"`
	Upd_day           int64   `thrift:",5,omitempty"`
}

type UserSaveDataInfo struct {
	U_seq               int32  `thrift:",1,omitempty"`
	U_id                string `thrift:",2,omitempty"`
	Uuid                string `thrift:",3,omitempty"`
	Device_uuid         string `thrift:",4,omitempty"`
	Type                string `thrift:",5,omitempty"`
	User_info           []any  `thrift:",6,omitempty"`  // TODO
	User_area_info      []any  `thrift:",7,omitempty"`  // TODO
	User_achievement    []any  `thrift:",8,omitempty"`  // TODO
	User_character      []any  `thrift:",9,omitempty"`  // TODO
	User_costume        []any  `thrift:",10,omitempty"` // TODO
	User_daily_mission  []any  `thrift:",11,omitempty"` // TODO
	User_follower       []any  `thrift:",12,omitempty"` // TODO
	User_music          []any  `thrift:",13,omitempty"` // TODO
	User_skill          []any  `thrift:",14,omitempty"` // TODO
	User_messenger      []any  `thrift:",15,omitempty"` // TODO
	User_guitar         []any  `thrift:",16,omitempty"` // TODO
	User_event_point    []any  `thrift:",17,omitempty"` // TODO
	User_follower_quest []any  `thrift:",18,omitempty"` // TODO
}

type ChThirdStreamStageDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	I_id        int32  `thrift:",5,omitempty"`
	Count       int32  `thrift:",6,omitempty"`
}

type ChThirdStreamStageReturn struct {
	Error       common_model.ErrorRetCode     `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet    `thrift:",2,omitempty"`
	Mode        string                        `thrift:",3,omitempty"`
	Call        string                        `thrift:",4,omitempty"`
	Data        ChThirdStreamStageRetDataInfo `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData  `thrift:",6,omitempty"`
}

type CheckPurchasedRetDataInfo struct {
	Is_purchased int16 `thrift:",1,omitempty"`
}

type GetProfileMusic struct {
	Call string                  `thrift:",1,omitempty"`
	Data GetProfileMusicDataInfo `thrift:",2,omitempty"`
}

type GetTotalMusicRankRetDataInfo struct {
	Total_rank_list []any `thrift:",1,omitempty"` // TODO
	Rank_list1      []any `thrift:",2,omitempty"` // TODO
	Rank_list2      []any `thrift:",3,omitempty"` // TODO
	Rank_list3      []any `thrift:",4,omitempty"` // TODO
	My_rank_list    []any `thrift:",5,omitempty"` // TODO
}

type LogCashComplete struct {
	Call string                  `thrift:",1,omitempty"`
	Data LogCashCompleteDataInfo `thrift:",2,omitempty"`
}

type UpdateAchievementDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Achievement int16  `thrift:",5,omitempty"`
}

type UseCredit struct {
	Call string            `thrift:",1,omitempty"`
	Data UseCreditDataInfo `thrift:",2,omitempty"`
}

type GetChThirdStarRewardRetDataInfo struct {
	I_id        int32      `thrift:",1,omitempty"`
	Reward_num  int32      `thrift:",2,omitempty"`
	User_ap     UserApData `thrift:",3,omitempty"`
	Reward_data []any      `thrift:",4,omitempty"` // TODO
}

type SetAdReward struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        SetAdRewardDataInfo    `thrift:",2,omitempty"`
	Common_data common_model.ParamData `thrift:",3,omitempty"`
}

type SetPassRewardDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Group       int32  `thrift:",5,omitempty"`
	Step        int32  `thrift:",6,omitempty"`
	Type        int16  `thrift:",7,omitempty"`
	I_Version   int32  `thrift:",8,omitempty"`
}

type SetUserFollowerProfileRewardReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        []any                        `thrift:",5,omitempty"` // TODO
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type UseCreditDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Credit      int64  `thrift:",5,omitempty"`
}

type UserAchievement struct {
	I_id       int64   `thrift:",1,omitempty"`
	I_Level    int64   `thrift:",2,omitempty"`
	D_Quantity float64 `thrift:",3,omitempty"`
	S_Quantity string  `thrift:",4,omitempty"`
}

type UserSaveReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        UserSaveRetDataInfo          `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type BuyPieceRetDataInfo struct {
	Idx    int32 `thrift:",1,omitempty"`
	U_gold int64 `thrift:",2,omitempty"`
	U_cp   int32 `thrift:",3,omitempty"`
	U_mp   int32 `thrift:",4,omitempty"`
}

type CheckUserRetDataInfo struct {
	U_seq int32  `thrift:",1,omitempty"`
	U_id  string `thrift:",2,omitempty"`
}

type GetProfileDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Get_u_seq   int64  `thrift:",5,omitempty"`
}

type SetFollowerProfileGiftReturn struct {
	Error       common_model.ErrorRetCode         `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet        `thrift:",2,omitempty"`
	Mode        string                            `thrift:",3,omitempty"`
	Call        string                            `thrift:",4,omitempty"`
	Data        SetFollowerProfileGiftRetDataInfo `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData      `thrift:",6,omitempty"`
}

type BuyCollectionItem struct {
	Call string                    `thrift:",1,omitempty"`
	Data BuyCollectionItemDataInfo `thrift:",2,omitempty"`
}

type ChangeUserReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        ChangeUserRetDataInfo        `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type GetPurchaseDiamondLogRetDataInfo struct {
	Bonus1 int64 `thrift:",1,omitempty"`
	Bonus2 int64 `thrift:",2,omitempty"`
	Bonus3 int64 `thrift:",3,omitempty"`
	Bonus4 int64 `thrift:",4,omitempty"`
	Bonus5 int64 `thrift:",5,omitempty"`
	Bonus6 int64 `thrift:",6,omitempty"`
}

type SaveUserFollowerQuest struct {
	I_CurrentID       int32   `thrift:",1,omitempty"`
	D_ConditionValue1 float64 `thrift:",2,omitempty"`
	D_ConditionValue2 float64 `thrift:",3,omitempty"`
	D_ConditionValue3 float64 `thrift:",4,omitempty"`
}

type UserLoadRetDataInfo struct {
	User          UserData         `thrift:",1,omitempty"`
	Area_data     map[any]any      `thrift:",2,omitempty"` // TODO
	User_contents UserContentsData `thrift:",3,omitempty"`
}

type UserPurchaseErrorLogDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Error_log   string `thrift:",5,omitempty"`
}
