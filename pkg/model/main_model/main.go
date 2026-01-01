package main_model

import "github.com/KhoalaS/guitar-girl-offline/pkg/model/common_model"

type GetFollowerQuestData struct {
	I_id                  int32   `thrift:",1,omitempty"`
	I_UnlockFollowerID    int32   `thrift:",2,omitempty"`
	I_UnlockFollowerLevel int32   `thrift:",3,omitempty"`
	S_ConditionType1      string  `thrift:",4,omitempty"`
	D_Condition1_1        float64 `thrift:",5,omitempty"`
	D_Condition1_2        float64 `thrift:",6,omitempty"`
	S_ConditionType2      string  `thrift:",7,omitempty"`
	D_Condition2_1        float64 `thrift:",8,omitempty"`
	D_Condition2_2        float64 `thrift:",9,omitempty"`
	S_ConditionType3      string  `thrift:",10,omitempty"`
	D_Condition3_1        float64 `thrift:",11,omitempty"`
	D_Condition3_2        float64 `thrift:",12,omitempty"`
	I_RewardGroup1        int32   `thrift:",13,omitempty"`
	I_RewardGroup2        int32   `thrift:",14,omitempty"`
	I_RewardGroup3        int32   `thrift:",15,omitempty"`
	S_Description_KO      string  `thrift:",16,omitempty"`
	S_Description_EN      string  `thrift:",17,omitempty"`
	S_Description_ZH_CHS  string  `thrift:",18,omitempty"`
	S_Description_ZH_CHT  string  `thrift:",19,omitempty"`
	S_Description_JA      string  `thrift:",20,omitempty"`
	S_Description_VI      string  `thrift:",21,omitempty"`
	S_Description_ES      string  `thrift:",22,omitempty"`
	S_Description_IT      string  `thrift:",23,omitempty"`
	S_Description_ID      string  `thrift:",24,omitempty"`
	S_Description_TH      string  `thrift:",25,omitempty"`
	S_Description_PT      string  `thrift:",26,omitempty"`
	S_Description_HI      string  `thrift:",27,omitempty"`
	B_IsActive            int16   `thrift:",28,omitempty"`
	I_Type                int32   `thrift:",29,omitempty"`
}

type GetLocalPushData struct {
	I_id                    int32  `thrift:",1,omitempty"`
	S_LocalPushRegisterType string `thrift:",2,omitempty"`
	I_TimeForSeconds        int64  `thrift:",3,omitempty"`
	B_IsMultipleJson        int16  `thrift:",4,omitempty"`
	S_JsonMessage_KO        string `thrift:",5,omitempty"`
	S_JsonMessage_EN        string `thrift:",6,omitempty"`
	S_JsonMessage_JA        string `thrift:",7,omitempty"`
	S_JsonMessage_ZH_CHS    string `thrift:",8,omitempty"`
	S_JsonMessage_ZH_CHT    string `thrift:",9,omitempty"`
	S_JsonMessage_VI        string `thrift:",10,omitempty"`
	S_JsonMessage_ES        string `thrift:",11,omitempty"`
	S_JsonMessage_IT        string `thrift:",12,omitempty"`
	S_JsonMessage_ID        string `thrift:",13,omitempty"`
	S_JsonMessage_TH        string `thrift:",14,omitempty"`
	S_JsonMessage_PT        string `thrift:",15,omitempty"`
	S_JsonMessage_HI        string `thrift:",16,omitempty"`
	B_IsActive              int16  `thrift:",17,omitempty"`
}

type NewStoreListReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        map[any]any                  `thrift:",4,omitempty"` // TODO
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type AvatarListDataInfo struct {
	Type string `thrift:",1,omitempty"`
}

type GetLocalPushReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        []any                        `thrift:",4,omitempty"` // TODO
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type GetServerTime struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        MainDataInfo           `thrift:",2,omitempty"`
	Sub_mode    string                 `thrift:",3,omitempty"`
	Common_data common_model.ParamData `thrift:",4,omitempty"`
}

type GetNoticeListRetDataInfo struct {
	Seq          int32  `thrift:",1,omitempty"`
	Notice_name  string `thrift:",2,omitempty"`
	Location_url string `thrift:",3,omitempty"`
	Img_url      string `thrift:",4,omitempty"`
}

type GetAchievementData struct {
	I_id                       int32  `thrift:",1,omitempty"`
	S_Title_KO                 string `thrift:",2,omitempty"`
	S_Title_EN                 string `thrift:",3,omitempty"`
	S_Title_JA                 string `thrift:",4,omitempty"`
	S_Title_ZH_CHS             string `thrift:",5,omitempty"`
	S_Title_ZH_CHT             string `thrift:",6,omitempty"`
	S_Title_VI                 string `thrift:",7,omitempty"`
	S_Title_ES                 string `thrift:",8,omitempty"`
	S_Title_IT                 string `thrift:",9,omitempty"`
	S_Title_ID                 string `thrift:",10,omitempty"`
	S_Title_TH                 string `thrift:",11,omitempty"`
	S_Title_PT                 string `thrift:",12,omitempty"`
	S_Title_HI                 string `thrift:",13,omitempty"`
	S_Description_KO           string `thrift:",14,omitempty"`
	S_Description_EN           string `thrift:",15,omitempty"`
	S_Description_JA           string `thrift:",16,omitempty"`
	S_Description_ZH_CHS       string `thrift:",17,omitempty"`
	S_Description_ZH_CHT       string `thrift:",18,omitempty"`
	S_Description_VI           string `thrift:",19,omitempty"`
	S_Description_ES           string `thrift:",20,omitempty"`
	S_Description_IT           string `thrift:",21,omitempty"`
	S_Description_ID           string `thrift:",22,omitempty"`
	S_Description_TH           string `thrift:",23,omitempty"`
	S_Description_PT           string `thrift:",24,omitempty"`
	S_Description_HI           string `thrift:",25,omitempty"`
	S_ResourceName             string `thrift:",26,omitempty"`
	S_AchievementConditionType string `thrift:",27,omitempty"`
	S_RewardType               string `thrift:",28,omitempty"`
	S_RewardDescription_KO     string `thrift:",29,omitempty"`
	S_RewardDescription_EN     string `thrift:",30,omitempty"`
	S_RewardDescription_JA     string `thrift:",31,omitempty"`
	S_RewardDescription_ZH_CHS string `thrift:",32,omitempty"`
	S_RewardDescription_ZH_CHT string `thrift:",33,omitempty"`
	S_RewardDescription_VI     string `thrift:",34,omitempty"`
	S_RewardDescription_ES     string `thrift:",35,omitempty"`
	S_RewardDescription_IT     string `thrift:",36,omitempty"`
	S_RewardDescription_ID     string `thrift:",37,omitempty"`
	S_RewardDescription_TH     string `thrift:",38,omitempty"`
	S_RewardDescription_PT     string `thrift:",39,omitempty"`
	S_RewardDescription_HI     string `thrift:",40,omitempty"`
	I_Reward_1                 int64  `thrift:",41,omitempty"`
	I_Reward_2                 int64  `thrift:",42,omitempty"`
	I_Reward_3                 int64  `thrift:",43,omitempty"`
	I_Reward_4                 int64  `thrift:",44,omitempty"`
	I_Reward_5                 int64  `thrift:",45,omitempty"`
	I_Reward_6                 int64  `thrift:",46,omitempty"`
	I_Reward_7                 int64  `thrift:",47,omitempty"`
	I_Reward_8                 int64  `thrift:",48,omitempty"`
	I_Reward_9                 int64  `thrift:",49,omitempty"`
	I_Reward_10                int64  `thrift:",50,omitempty"`
	I_Area                     int16  `thrift:",51,omitempty"`
	B_IsActive                 int16  `thrift:",52,omitempty"`
	I_Reward_11                int64  `thrift:",53,omitempty"`
	I_Reward_12                int64  `thrift:",54,omitempty"`
	I_MaxLevel                 int16  `thrift:",55,omitempty"`
	S_Condition_1              string `thrift:",56,omitempty"`
	S_Condition_2              string `thrift:",57,omitempty"`
	S_Condition_3              string `thrift:",58,omitempty"`
	S_Condition_4              string `thrift:",59,omitempty"`
	S_Condition_5              string `thrift:",60,omitempty"`
	S_Condition_6              string `thrift:",61,omitempty"`
	S_Condition_7              string `thrift:",62,omitempty"`
	S_Condition_8              string `thrift:",63,omitempty"`
	S_Condition_9              string `thrift:",64,omitempty"`
	S_Condition_10             string `thrift:",65,omitempty"`
	S_Condition_11             string `thrift:",66,omitempty"`
	S_Condition_12             string `thrift:",67,omitempty"`
	S_Condition_13             string `thrift:",68,omitempty"`
	S_Condition_14             string `thrift:",69,omitempty"`
	S_Condition_15             string `thrift:",70,omitempty"`
	S_Condition_16             string `thrift:",71,omitempty"`
	S_Condition_17             string `thrift:",72,omitempty"`
	S_Condition_18             string `thrift:",73,omitempty"`
	S_Condition_19             string `thrift:",74,omitempty"`
	S_Condition_20             string `thrift:",75,omitempty"`
	I_Reward_13                int64  `thrift:",76,omitempty"`
	I_Reward_14                int64  `thrift:",77,omitempty"`
	I_Reward_15                int64  `thrift:",78,omitempty"`
	I_Reward_16                int64  `thrift:",79,omitempty"`
	I_Reward_17                int64  `thrift:",80,omitempty"`
	I_Reward_18                int64  `thrift:",81,omitempty"`
	I_Reward_19                int64  `thrift:",82,omitempty"`
	I_Reward_20                int64  `thrift:",83,omitempty"`
}

type GetCostumeData struct {
	I_id                      int32   `thrift:",1,omitempty"`
	S_Name_KO                 string  `thrift:",2,omitempty"`
	S_Name_EN                 string  `thrift:",3,omitempty"`
	S_Name_JA                 string  `thrift:",4,omitempty"`
	S_Name_ZH_CHS             string  `thrift:",5,omitempty"`
	S_Name_ZH_CHT             string  `thrift:",6,omitempty"`
	S_Name_VI                 string  `thrift:",7,omitempty"`
	S_Name_ES                 string  `thrift:",8,omitempty"`
	S_Name_IT                 string  `thrift:",9,omitempty"`
	S_Name_ID                 string  `thrift:",10,omitempty"`
	S_Name_TH                 string  `thrift:",11,omitempty"`
	S_Name_PT                 string  `thrift:",12,omitempty"`
	S_Name_HI                 string  `thrift:",13,omitempty"`
	S_Description_KO          string  `thrift:",14,omitempty"`
	S_Description_EN          string  `thrift:",15,omitempty"`
	S_Description_JA          string  `thrift:",16,omitempty"`
	S_Description_ZH_CHS      string  `thrift:",17,omitempty"`
	S_Description_ZH_CHT      string  `thrift:",18,omitempty"`
	S_Description_VI          string  `thrift:",19,omitempty"`
	S_Description_ES          string  `thrift:",20,omitempty"`
	S_Description_IT          string  `thrift:",21,omitempty"`
	S_Description_ID          string  `thrift:",22,omitempty"`
	S_Description_TH          string  `thrift:",23,omitempty"`
	S_Description_PT          string  `thrift:",24,omitempty"`
	S_Description_HI          string  `thrift:",25,omitempty"`
	S_ResourceName            string  `thrift:",26,omitempty"`
	S_CostumeName             string  `thrift:",27,omitempty"`
	I_MaxLevel                int64   `thrift:",28,omitempty"`
	S_GoodsType               string  `thrift:",29,omitempty"`
	S_Cost                    string  `thrift:",30,omitempty"`
	F_CostIncreaseRate        float64 `thrift:",31,omitempty"`
	S_Operation               string  `thrift:",32,omitempty"`
	D_Amount                  int64   `thrift:",33,omitempty"`
	F_AmountIncreaseRate      float64 `thrift:",34,omitempty"`
	S_BonusOperation          string  `thrift:",35,omitempty"`
	D_BonusAmount             int64   `thrift:",36,omitempty"`
	F_BonusAmountIncreaseRate float64 `thrift:",37,omitempty"`
	I_UnlockLevel             int64   `thrift:",38,omitempty"`
	I_RequirementFollowerId_1 int64   `thrift:",39,omitempty"`
	I_RequirementFollowerId_2 int64   `thrift:",40,omitempty"`
	I_RequirementPropId_1     int64   `thrift:",41,omitempty"`
	I_RequirementPropId_2     int64   `thrift:",42,omitempty"`
	I_RequirementPropCount    int64   `thrift:",43,omitempty"`
	I_SortOrder               int64   `thrift:",44,omitempty"`
	F_FanMultiply             float64 `thrift:",45,omitempty"`
	I_AcquisitionType         int64   `thrift:",46,omitempty"`
	I_AcquisitionId           int64   `thrift:",47,omitempty"`
	B_IsActive                int16   `thrift:",48,omitempty"`
	I_Area                    int16   `thrift:",49,omitempty"`
	I_DownloadType            int16   `thrift:",50,omitempty"`
	S_AltCostumeName          string  `thrift:",51,omitempty"`
	S_AltCostumeStartTime     string  `thrift:",52,omitempty"`
	S_AltCostumeEndTime       string  `thrift:",53,omitempty"`
}

type GetEventRewardListDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Event_idx   int32  `thrift:",5,omitempty"`
}

type GetGameDataListDataInfo struct {
	Game_type   string `thrift:",1,omitempty"`
	Device_uuid string `thrift:",2,omitempty"`
}

type GetServerTimeReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        GetServerTimeRetDataInfo     `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type MainDataInfo struct {
	U_seq        int32  `thrift:",1,omitempty"`
	U_id         string `thrift:",2,omitempty"`
	Uuid         string `thrift:",3,omitempty"`
	Device_uuid  string `thrift:",4,omitempty"`
	Country_code string `thrift:",5,omitempty"`
}

type InitReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        InitRetDataInfo              `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type MoreGamesReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        map[any]any                  `thrift:",4,omitempty"` // TODO
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type DefaultSettingDataList struct {
	Setting_key   string `thrift:",1,omitempty"`
	Setting_value string `thrift:",2,omitempty"`
}

type GetChThirdChapterData struct {
	I_id             int32  `thrift:",1,omitempty"`
	I_unlockLevel    int32  `thrift:",2,omitempty"`
	I_RewardGroupID1 int32  `thrift:",3,omitempty"`
	I_GoalStar1      int32  `thrift:",4,omitempty"`
	I_RewardGroupID2 int32  `thrift:",5,omitempty"`
	I_GoalStar2      int32  `thrift:",6,omitempty"`
	I_RewardGroupID3 int32  `thrift:",7,omitempty"`
	I_GoalStar3      int32  `thrift:",8,omitempty"`
	S_Name_KO        string `thrift:",9,omitempty"`
	S_Name_EN        string `thrift:",10,omitempty"`
	S_Name_JA        string `thrift:",11,omitempty"`
	S_Name_ZH_CHS    string `thrift:",12,omitempty"`
	S_Name_ZH_CHT    string `thrift:",13,omitempty"`
	S_Name_VI        string `thrift:",14,omitempty"`
	S_Name_ES        string `thrift:",15,omitempty"`
	S_Name_IT        string `thrift:",16,omitempty"`
	S_Name_ID        string `thrift:",17,omitempty"`
	S_Name_TH        string `thrift:",18,omitempty"`
	S_Name_PT        string `thrift:",19,omitempty"`
	S_Name_HI        string `thrift:",20,omitempty"`
	B_IsActive       int16  `thrift:",21,omitempty"`
}

type GetCollection struct {
	Call string                `thrift:",1,omitempty"`
	Data GetCollectionDataInfo `thrift:",2,omitempty"`
}

type GetLocalPush struct {
	Call string               `thrift:",1,omitempty"`
	Data GetLocalPushDataInfo `thrift:",2,omitempty"`
}

type GetSystemStringData struct {
	I_StringID int64  `thrift:",1,omitempty"`
	S_Category string `thrift:",2,omitempty"`
	S_UIType   string `thrift:",3,omitempty"`
	S_Key      string `thrift:",4,omitempty"`
	S_KO       string `thrift:",5,omitempty"`
	S_EN       string `thrift:",6,omitempty"`
	S_JA       string `thrift:",7,omitempty"`
	S_ZH_CHS   string `thrift:",8,omitempty"`
	S_ZH_CHT   string `thrift:",9,omitempty"`
	S_ID       string `thrift:",10,omitempty"`
	S_PT       string `thrift:",11,omitempty"`
	S_ES       string `thrift:",12,omitempty"`
	S_RU       string `thrift:",13,omitempty"`
	S_DE       string `thrift:",14,omitempty"`
	S_AR_XA    string `thrift:",15,omitempty"`
	S_VI       string `thrift:",16,omitempty"`
	S_IT       string `thrift:",17,omitempty"`
	S_TH       string `thrift:",18,omitempty"`
	S_HI       string `thrift:",19,omitempty"`
}

type GetUpdateTime struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        GetUpdateTimeDataInfo  `thrift:",2,omitempty"`
	Common_data common_model.ParamData `thrift:",3,omitempty"`
}

type DefaultSettingListReturn struct {
	Error       common_model.ErrorRetCode     `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet    `thrift:",2,omitempty"`
	Mode        string                        `thrift:",3,omitempty"`
	Call        string                        `thrift:",4,omitempty"`
	Data        DefaultSettingListRetDataInfo `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData  `thrift:",6,omitempty"`
}

type GetUpdateTimeReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        GetUpdateTimeRetDataInfo     `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type MainRetDataInfo struct {
	GameInfo MainGameInfo `thrift:",1,omitempty"`
	Features []any        `thrift:",2,omitempty"` // TODO
}

type TitleListRetDataInfo struct {
	Idx           int16  `thrift:",1,omitempty"`
	Resource_name string `thrift:",2,omitempty"`
	Down_status   int16  `thrift:",3,omitempty"`
	Buy_type      int16  `thrift:",4,omitempty"`
	Price         int32  `thrift:",5,omitempty"`
	Buy_type2     int16  `thrift:",6,omitempty"`
	Price2        int32  `thrift:",7,omitempty"`
	Value         int32  `thrift:",8,omitempty"`
	Event_point   int64  `thrift:",9,omitempty"`
	Desc_ko       string `thrift:",10,omitempty"`
	Desc_en       string `thrift:",11,omitempty"`
	Desc_ja       string `thrift:",12,omitempty"`
}

type GetDailyMissionData struct {
	I_id                        int32   `thrift:",1,omitempty"`
	S_Title_KO                  string  `thrift:",2,omitempty"`
	S_Title_EN                  string  `thrift:",3,omitempty"`
	S_Title_JA                  string  `thrift:",4,omitempty"`
	S_Title_ZH_CHS              string  `thrift:",5,omitempty"`
	S_Title_ZH_CHT              string  `thrift:",6,omitempty"`
	S_Title_VI                  string  `thrift:",7,omitempty"`
	S_Title_ES                  string  `thrift:",8,omitempty"`
	S_Title_IT                  string  `thrift:",9,omitempty"`
	S_Title_ID                  string  `thrift:",10,omitempty"`
	S_Title_TH                  string  `thrift:",11,omitempty"`
	S_Title_PT                  string  `thrift:",12,omitempty"`
	S_Title_HI                  string  `thrift:",13,omitempty"`
	S_Description_KO            string  `thrift:",14,omitempty"`
	S_Description_EN            string  `thrift:",15,omitempty"`
	S_Description_JA            string  `thrift:",16,omitempty"`
	S_Description_ZH_CHS        string  `thrift:",17,omitempty"`
	S_Description_ZH_CHT        string  `thrift:",18,omitempty"`
	S_Description_VI            string  `thrift:",19,omitempty"`
	S_Description_ES            string  `thrift:",20,omitempty"`
	S_Description_IT            string  `thrift:",21,omitempty"`
	S_Description_ID            string  `thrift:",22,omitempty"`
	S_Description_TH            string  `thrift:",23,omitempty"`
	S_Description_PT            string  `thrift:",24,omitempty"`
	S_Description_HI            string  `thrift:",25,omitempty"`
	S_ResourceName              string  `thrift:",26,omitempty"`
	S_DailyMissionConditionType string  `thrift:",27,omitempty"`
	D_Condition_1               float64 `thrift:",28,omitempty"`
	S_RewardType                string  `thrift:",29,omitempty"`
	S_RewardDescription_KO      string  `thrift:",30,omitempty"`
	S_RewardDescription_EN      string  `thrift:",31,omitempty"`
	S_RewardDescription_JA      string  `thrift:",32,omitempty"`
	S_RewardDescription_ZH_CHS  string  `thrift:",33,omitempty"`
	S_RewardDescription_ZH_CHT  string  `thrift:",34,omitempty"`
	S_RewardDescription_VI      string  `thrift:",35,omitempty"`
	S_RewardDescription_ES      string  `thrift:",36,omitempty"`
	S_RewardDescription_IT      string  `thrift:",37,omitempty"`
	S_RewardDescription_ID      string  `thrift:",38,omitempty"`
	S_RewardDescription_TH      string  `thrift:",39,omitempty"`
	S_RewardDescription_PT      string  `thrift:",40,omitempty"`
	S_RewardDescription_HI      string  `thrift:",41,omitempty"`
	I_Reward_1                  int32   `thrift:",42,omitempty"`
	I_Area                      int16   `thrift:",43,omitempty"`
	B_IsActive                  int16   `thrift:",44,omitempty"`
}

type GetFollowerGiftItemData struct {
	I_id           int32  `thrift:",1,omitempty"`
	I_GiftType     int32  `thrift:",2,omitempty"`
	D_Value        int64  `thrift:",3,omitempty"`
	I_Limit        int32  `thrift:",4,omitempty"`
	S_ResourceName string `thrift:",5,omitempty"`
	S_Name_KO      string `thrift:",6,omitempty"`
	S_Name_EN      string `thrift:",7,omitempty"`
	S_Name_JA      string `thrift:",8,omitempty"`
	S_Name_ZH_CHS  string `thrift:",9,omitempty"`
	S_Name_ZH_CHT  string `thrift:",10,omitempty"`
	S_Name_VI      string `thrift:",11,omitempty"`
	S_Name_ES      string `thrift:",12,omitempty"`
	S_Name_IT      string `thrift:",13,omitempty"`
	S_Name_ID      string `thrift:",14,omitempty"`
	S_Name_TH      string `thrift:",15,omitempty"`
	S_Name_PT      string `thrift:",16,omitempty"`
	S_Name_HI      string `thrift:",17,omitempty"`
	B_IsActive     int16  `thrift:",18,omitempty"`
}

type GetUpdateTimeDataInfo struct {
	Device_uuid string `thrift:",1,omitempty"`
}

type MoreGamesRetDataInfo struct {
	App_id            int32  `thrift:",1,omitempty"`
	Material_id       int32  `thrift:",2,omitempty"`
	Material_type     string `thrift:",3,omitempty"`
	Material_category string `thrift:",4,omitempty"`
	Material_name     string `thrift:",5,omitempty"`
	Material_link     string `thrift:",6,omitempty"`
	Link_text         string `thrift:",7,omitempty"`
	Link_url_aos      string `thrift:",8,omitempty"`
	Link_url_ios      string `thrift:",9,omitempty"`
	Thumbnail_url     string `thrift:",10,omitempty"`
	Icon_url          string `thrift:",11,omitempty"`
	Update_date       string `thrift:",12,omitempty"`
}

type AvatarList struct {
	Call string             `thrift:",1,omitempty"`
	Data AvatarListDataInfo `thrift:",2,omitempty"`
}

type GetEventRewardListReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        map[any]any                  `thrift:",5,omitempty"` // TODO
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type GetLocalPushDataInfo struct {
	Device_uuid string `thrift:",1,omitempty"`
}

type GetNoticeList struct {
	Call string                `thrift:",1,omitempty"`
	Data GetNoticeListDataInfo `thrift:",2,omitempty"`
}

type MainFeatures struct {
	Album_idx    int32   `thrift:",1,omitempty"`
	F_type       int16   `thrift:",2,omitempty"`
	Buy_type     int16   `thrift:",3,omitempty"`
	New_status   int16   `thrift:",4,omitempty"`
	Discount     int16   `thrift:",5,omitempty"`
	Price        float64 `thrift:",6,omitempty"`
	Banner_url   string  `thrift:",7,omitempty"`
	Location_url string  `thrift:",8,omitempty"`
}

type DefaultSettingList struct {
	Call        string                     `thrift:",1,omitempty"`
	Data        DefaultSettingListDataInfo `thrift:",2,omitempty"`
	Common_data common_model.ParamData     `thrift:",3,omitempty"`
}

type GetGameDataListReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        map[any]any                  `thrift:",5,omitempty"` // TODO
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type GetTicketCollectionData struct {
	I_id                 int32  `thrift:",1,omitempty"`
	S_TicketType         string `thrift:",2,omitempty"`
	S_Title_KO           string `thrift:",3,omitempty"`
	S_Title_EN           string `thrift:",4,omitempty"`
	S_Title_JA           string `thrift:",5,omitempty"`
	S_Title_ZH_CHS       string `thrift:",6,omitempty"`
	S_Title_ZH_CHT       string `thrift:",7,omitempty"`
	S_Title_VI           string `thrift:",8,omitempty"`
	S_Title_ES           string `thrift:",9,omitempty"`
	S_Title_IT           string `thrift:",10,omitempty"`
	S_Title_ID           string `thrift:",11,omitempty"`
	S_Title_TH           string `thrift:",12,omitempty"`
	S_Title_PT           string `thrift:",13,omitempty"`
	S_Title_HI           string `thrift:",14,omitempty"`
	S_Description_KO     string `thrift:",15,omitempty"`
	S_Description_EN     string `thrift:",16,omitempty"`
	S_Description_JA     string `thrift:",17,omitempty"`
	S_Description_ZH_CHS string `thrift:",18,omitempty"`
	S_Description_ZH_CHT string `thrift:",19,omitempty"`
	S_Description_VI     string `thrift:",20,omitempty"`
	S_Description_ES     string `thrift:",21,omitempty"`
	S_Description_IT     string `thrift:",22,omitempty"`
	S_Description_ID     string `thrift:",23,omitempty"`
	S_Description_TH     string `thrift:",24,omitempty"`
	S_Description_PT     string `thrift:",25,omitempty"`
	S_Description_HI     string `thrift:",26,omitempty"`
	S_ResourceName       string `thrift:",27,omitempty"`
	S_ResourceName_Big   string `thrift:",28,omitempty"`
	I_Season             int16  `thrift:",29,omitempty"`
	S_PassUIBundleName   string `thrift:",30,omitempty"`
	S_BackgroundColor    string `thrift:",31,omitempty"`
	S_LabelColor         string `thrift:",32,omitempty"`
	S_SeasonLabelColor   string `thrift:",33,omitempty"`
	S_SeasonBackColor    string `thrift:",34,omitempty"`
	I_SortOrder          int32  `thrift:",35,omitempty"`
	B_IsActive           int16  `thrift:",36,omitempty"`
}

type NewStoreListDataInfo struct {
	Os           int16  `thrift:",1,omitempty"`
	Type         string `thrift:",2,omitempty"`
	Country_code string `thrift:",3,omitempty"`
	U_seq        int32  `thrift:",4,omitempty"`
	U_id         string `thrift:",5,omitempty"`
	Uuid         string `thrift:",6,omitempty"`
	Device_uuid  string `thrift:",7,omitempty"`
}

type GetFanData struct {
	I_Area        int64 `thrift:",1,omitempty"`
	I_Grade       int64 `thrift:",2,omitempty"`
	I_FanCount    int64 `thrift:",3,omitempty"`
	I_BonusRate   int64 `thrift:",4,omitempty"`
	I_BonusRateUI int64 `thrift:",5,omitempty"`
}

type GetUnitData struct {
	I_id                      int32   `thrift:",1,omitempty"`
	S_Name_KO                 string  `thrift:",2,omitempty"`
	S_Name_EN                 string  `thrift:",3,omitempty"`
	S_Name_JA                 string  `thrift:",4,omitempty"`
	S_Name_ZH_CHS             string  `thrift:",5,omitempty"`
	S_Name_ZH_CHT             string  `thrift:",6,omitempty"`
	S_Name_VI                 string  `thrift:",7,omitempty"`
	S_Name_ES                 string  `thrift:",8,omitempty"`
	S_Name_IT                 string  `thrift:",9,omitempty"`
	S_Name_ID                 string  `thrift:",10,omitempty"`
	S_Name_TH                 string  `thrift:",11,omitempty"`
	S_Name_PT                 string  `thrift:",12,omitempty"`
	S_Name_HI                 string  `thrift:",13,omitempty"`
	S_Description_KO          string  `thrift:",14,omitempty"`
	S_Description_EN          string  `thrift:",15,omitempty"`
	S_Description_JA          string  `thrift:",16,omitempty"`
	S_Description_ZH_CHS      string  `thrift:",17,omitempty"`
	S_Description_ZH_CHT      string  `thrift:",18,omitempty"`
	S_Description_VI          string  `thrift:",19,omitempty"`
	S_Description_ES          string  `thrift:",20,omitempty"`
	S_Description_IT          string  `thrift:",21,omitempty"`
	S_Description_ID          string  `thrift:",22,omitempty"`
	S_Description_TH          string  `thrift:",23,omitempty"`
	S_Description_PT          string  `thrift:",24,omitempty"`
	S_Description_HI          string  `thrift:",25,omitempty"`
	S_ResourceName            string  `thrift:",26,omitempty"`
	I_MaxLevel                int64   `thrift:",27,omitempty"`
	S_GoodsType               string  `thrift:",28,omitempty"`
	I_Cost                    float64 `thrift:",29,omitempty"`
	F_CostIncreaseValue       float64 `thrift:",30,omitempty"`
	F_Value                   float64 `thrift:",31,omitempty"`
	F_IncreaseValue           float64 `thrift:",32,omitempty"`
	I_UnlockLevel             int64   `thrift:",33,omitempty"`
	I_UnlockCost              int64   `thrift:",34,omitempty"`
	I_RequirementUnitId_1     int64   `thrift:",35,omitempty"`
	I_RequirementUnitId_2     int64   `thrift:",36,omitempty"`
	I_RequirementFollowerId_1 int64   `thrift:",37,omitempty"`
	I_RequirementFollowerId_2 int64   `thrift:",38,omitempty"`
	I_RequirementPropId_1     int64   `thrift:",39,omitempty"`
	I_RequirementPropId_2     int64   `thrift:",40,omitempty"`
	I_RequirementPropCount    int64   `thrift:",41,omitempty"`
	I_Area                    int16   `thrift:",42,omitempty"`
	B_IsActive                int16   `thrift:",43,omitempty"`
}

type MainGameInfo struct {
	Album_list      []any                        `thrift:",1,omitempty"` // TODO
	Music_list      []any                        `thrift:",2,omitempty"` // TODO
	Album_language  map[any]any                  `thrift:",3,omitempty"` // TODO
	Music_language  map[any]any                  `thrift:",4,omitempty"` // TODO
	Tab             map[any]any                  `thrift:",5,omitempty"` // TODO
	Notice          []any                        `thrift:",6,omitempty"` // TODO
	Default_setting GetDefaultSettingRetDataInfo `thrift:",7,omitempty"`
}

type StoreListData struct {
	Store_idx int16 `thrift:",1,omitempty"`
	Music_idx int32 `thrift:",2,omitempty"`
}

type TitleList struct {
	Call string            `thrift:",1,omitempty"`
	Data TitleListDataInfo `thrift:",2,omitempty"`
}

type TitleListDataInfo struct {
	Type string `thrift:",1,omitempty"`
}

type GetAdListData struct {
	I_id              int32  `thrift:",1,omitempty"`
	I_Category        int32  `thrift:",2,omitempty"`
	S_ConditionType1  string `thrift:",3,omitempty"`
	S_ConditionValue1 string `thrift:",4,omitempty"`
	S_ConditionType2  string `thrift:",5,omitempty"`
	S_ConditionValue2 string `thrift:",6,omitempty"`
	S_ResourceName    string `thrift:",7,omitempty"`
	S_Name_KO         string `thrift:",8,omitempty"`
	S_Name_EN         string `thrift:",9,omitempty"`
	S_Name_JA         string `thrift:",10,omitempty"`
	S_Name_ZH_CHS     string `thrift:",11,omitempty"`
	S_Name_ZH_CHT     string `thrift:",12,omitempty"`
	S_Name_VI         string `thrift:",13,omitempty"`
	S_Name_ES         string `thrift:",14,omitempty"`
	S_Name_IT         string `thrift:",15,omitempty"`
	S_Name_ID         string `thrift:",16,omitempty"`
	S_Name_TH         string `thrift:",17,omitempty"`
	S_Name_PT         string `thrift:",18,omitempty"`
	S_Name_HI         string `thrift:",19,omitempty"`
	B_IsActive        int16  `thrift:",20,omitempty"`
	S_RewardType      string `thrift:",21,omitempty"`
	I_RewardGroup     int32  `thrift:",22,omitempty"`
	S_ExtraValue      string `thrift:",23,omitempty"`
	S_AdType          string `thrift:",24,omitempty"`
}

type GetChThirdScoreData struct {
	I_id             int32 `thrift:",1,omitempty"`
	I_Level          int32 `thrift:",2,omitempty"`
	I_CharacterScore int32 `thrift:",3,omitempty"`
	I_FollowerScore  int32 `thrift:",4,omitempty"`
	I_MusicScore     int32 `thrift:",5,omitempty"`
	B_IsActive       int16 `thrift:",6,omitempty"`
}

type GetCollectionReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        []any                        `thrift:",4,omitempty"` // TODO
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type GetCollectionItem struct {
	Special_music_idx int64 `thrift:",1,omitempty"`
	Idx               int64 `thrift:",2,omitempty"`
	Item_type         int16 `thrift:",3,omitempty"`
	Item_value        int64 `thrift:",4,omitempty"`
	Buy_cnt           int64 `thrift:",5,omitempty"`
}

type GetGameDataListRetDataInfo struct {
	Consume                        []any `thrift:",1,omitempty"`  // TODO
	Music                          []any `thrift:",2,omitempty"`  // TODO
	Costume                        []any `thrift:",3,omitempty"`  // TODO
	Prop                           []any `thrift:",4,omitempty"`  // TODO
	Follower                       []any `thrift:",5,omitempty"`  // TODO
	Buff                           []any `thrift:",6,omitempty"`  // TODO
	Unit                           []any `thrift:",7,omitempty"`  // TODO
	Skill                          []any `thrift:",8,omitempty"`  // TODO
	Achievement                    []any `thrift:",9,omitempty"`  // TODO
	Daily_mission                  []any `thrift:",10,omitempty"` // TODO
	Music_level                    []any `thrift:",11,omitempty"` // TODO
	Shop                           []any `thrift:",12,omitempty"` // TODO
	Reward_group                   []any `thrift:",13,omitempty"` // TODO
	Fan                            []any `thrift:",14,omitempty"` // TODO
	SystemNotification             []any `thrift:",15,omitempty"` // TODO
	SystemString                   []any `thrift:",16,omitempty"` // TODO
	SubscribeList                  []any `thrift:",17,omitempty"` // TODO
	SubscribePassReward            []any `thrift:",18,omitempty"` // TODO
	SubscribePass                  []any `thrift:",19,omitempty"` // TODO
	Guitar                         []any `thrift:",20,omitempty"` // TODO
	SubscribePassRewardInformation []any `thrift:",21,omitempty"` // TODO
	Ticketcollection               []any `thrift:",22,omitempty"` // TODO
	Character                      []any `thrift:",23,omitempty"` // TODO
	Samseckevent                   []any `thrift:",24,omitempty"` // TODO
	Localpush                      []any `thrift:",25,omitempty"` // TODO
	Followerquest                  []any `thrift:",26,omitempty"` // TODO
	Proplevel                      []any `thrift:",27,omitempty"` // TODO
	Follower_profile               []any `thrift:",28,omitempty"` // TODO
	Passgoodsshop                  []any `thrift:",29,omitempty"` // TODO
	Followergiftitem               []any `thrift:",30,omitempty"` // TODO
	Followerprofilelevel           []any `thrift:",31,omitempty"` // TODO
	Ad_list                        []any `thrift:",32,omitempty"` // TODO
	Chthird_stage                  []any `thrift:",33,omitempty"` // TODO
	Chthird_score                  []any `thrift:",34,omitempty"` // TODO
	Chthird_chapter                []any `thrift:",35,omitempty"` // TODO
	Percent                        []any `thrift:",36,omitempty"` // TODO
	Ad_level                       []any `thrift:",37,omitempty"` // TODO
	Event_list                     []any `thrift:",38,omitempty"` // TODO
	Select_reward                  []any `thrift:",39,omitempty"` // TODO
}

type GetSubscribePass struct {
	I_id                 int64  `thrift:",1,omitempty"`
	I_SubscribeID        int64  `thrift:",2,omitempty"`
	S_AcquireList        string `thrift:",3,omitempty"`
	I_PointPrice         int16  `thrift:",4,omitempty"`
	I_PaidPoint          int64  `thrift:",5,omitempty"`
	I_ActiveBuffID       int64  `thrift:",6,omitempty"`
	I_HiddenStep         int64  `thrift:",7,omitempty"`
	I_ADPoint            int64  `thrift:",8,omitempty"`
	I_ADCoolTime         int64  `thrift:",9,omitempty"`
	I_TicketCollectionId int64  `thrift:",10,omitempty"`
}

type InitDataInfo struct {
	Type        string `thrift:",1,omitempty"`
	Os          int16  `thrift:",2,omitempty"`
	Ver         int32  `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type TitleListReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        []any                        `thrift:",4,omitempty"` // TODO
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type GetMusicLevelData struct {
	I_Level                          int16   `thrift:",1,omitempty"`
	F_EncoreBonusAppearRate          float64 `thrift:",2,omitempty"`
	I_EncoreBonusAppearPercent       int64   `thrift:",3,omitempty"`
	F_EncoreBonusAppearCooltime_Sec  float64 `thrift:",4,omitempty"`
	F_EncoreBonusAppearCooltime_Hour float64 `thrift:",5,omitempty"`
	I_EncoreBonusGiftAmount          int16   `thrift:",6,omitempty"`
	I_EncoreFollowerProfileExp       int32   `thrift:",7,omitempty"`
	I_ChThirdCoolTime                int32   `thrift:",8,omitempty"`
}

type GetConsumeData struct {
	I_id           int32  `thrift:",1,omitempty"`
	S_ResourceName string `thrift:",2,omitempty"`
	S_Title_KO     string `thrift:",3,omitempty"`
	S_Title_EN     string `thrift:",4,omitempty"`
	S_Title_JA     string `thrift:",5,omitempty"`
	S_Title_ZH_CHS string `thrift:",6,omitempty"`
	S_Title_ZH_CHT string `thrift:",7,omitempty"`
	S_Title_VI     string `thrift:",8,omitempty"`
	S_Title_ES     string `thrift:",9,omitempty"`
	S_Title_IT     string `thrift:",10,omitempty"`
	S_Title_ID     string `thrift:",11,omitempty"`
	S_Title_TH     string `thrift:",12,omitempty"`
	S_Title_PT     string `thrift:",13,omitempty"`
	S_Title_HI     string `thrift:",14,omitempty"`
	S_Limit        string `thrift:",15,omitempty"`
	I_Type         int16  `thrift:",16,omitempty"`
	S_ExtraValue   string `thrift:",17,omitempty"`
	I_Area         int16  `thrift:",18,omitempty"`
	B_IsActive     int16  `thrift:",19,omitempty"`
}

type GetTabListReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        map[any]any                  `thrift:",4,omitempty"` // TODO
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type NewStoreProduct struct {
	Product_id     int32   `thrift:",1,omitempty"`
	Product_type   int16   `thrift:",2,omitempty"`
	Product_name   string  `thrift:",3,omitempty"`
	Store_name     string  `thrift:",4,omitempty"`
	IapID          string  `thrift:",5,omitempty"`
	Price          float64 `thrift:",6,omitempty"`
	Prev_price     float64 `thrift:",7,omitempty"`
	Price_str      string  `thrift:",8,omitempty"`
	Prev_price_str string  `thrift:",9,omitempty"`
	Currency       string  `thrift:",10,omitempty"`
	Image          string  `thrift:",11,omitempty"`
	Thumbnail      string  `thrift:",12,omitempty"`
	Main_desc      string  `thrift:",13,omitempty"`
}

type GetMusicData struct {
	I_id                      int32   `thrift:",1,omitempty"`
	S_Title_KO                string  `thrift:",2,omitempty"`
	S_Title_EN                string  `thrift:",3,omitempty"`
	S_Title_JA                string  `thrift:",4,omitempty"`
	S_Title_ZH_CHS            string  `thrift:",5,omitempty"`
	S_Title_ZH_CHT            string  `thrift:",6,omitempty"`
	S_Title_VI                string  `thrift:",7,omitempty"`
	S_Title_ES                string  `thrift:",8,omitempty"`
	S_Title_IT                string  `thrift:",9,omitempty"`
	S_Title_ID                string  `thrift:",10,omitempty"`
	S_Title_TH                string  `thrift:",11,omitempty"`
	S_Title_PT                string  `thrift:",12,omitempty"`
	S_Title_HI                string  `thrift:",13,omitempty"`
	S_Description_1_KO        string  `thrift:",14,omitempty"`
	S_Description_1_EN        string  `thrift:",15,omitempty"`
	S_Description_1_JA        string  `thrift:",16,omitempty"`
	S_Description_1_ZH_CHS    string  `thrift:",17,omitempty"`
	S_Description_1_ZH_CHT    string  `thrift:",18,omitempty"`
	S_Description_1_VI        string  `thrift:",19,omitempty"`
	S_Description_1_ES        string  `thrift:",20,omitempty"`
	S_Description_1_IT        string  `thrift:",21,omitempty"`
	S_Description_1_ID        string  `thrift:",22,omitempty"`
	S_Description_1_TH        string  `thrift:",23,omitempty"`
	S_Description_1_PT        string  `thrift:",24,omitempty"`
	S_Description_1_HI        string  `thrift:",25,omitempty"`
	S_ResourceName            string  `thrift:",26,omitempty"`
	S_MusicFileName           string  `thrift:",27,omitempty"`
	I_MaxLevel                int64   `thrift:",28,omitempty"`
	D_LevelRate               float64 `thrift:",29,omitempty"`
	D_LateCost                int64   `thrift:",30,omitempty"`
	D_Cost                    float64 `thrift:",31,omitempty"`
	D_CostIncreaseValue       float64 `thrift:",32,omitempty"`
	D_CostValueConstant       int64   `thrift:",33,omitempty"`
	D_Amount                  float64 `thrift:",34,omitempty"`
	D_AmountIncreaseValue     float64 `thrift:",35,omitempty"`
	D_AmountValueConstant     int64   `thrift:",36,omitempty"`
	D_BonusAmountIncreaseRate int64   `thrift:",37,omitempty"`
	I_UnlockLevel             int64   `thrift:",38,omitempty"`
	D_UnlockCost              int64   `thrift:",39,omitempty"`
	I_RequirementFollowerId_1 int64   `thrift:",40,omitempty"`
	I_RequirementFollowerId_2 int64   `thrift:",41,omitempty"`
	I_RequirementMusicId_1    int64   `thrift:",42,omitempty"`
	I_RequirementMusicId_2    int64   `thrift:",43,omitempty"`
	I_RequirementMusicCount   int64   `thrift:",44,omitempty"`
	S_GoodsType               string  `thrift:",45,omitempty"`
	I_SortOrder               int64   `thrift:",46,omitempty"`
	I_AcquisitionType         int64   `thrift:",47,omitempty"`
	I_AcquisitionId           int64   `thrift:",48,omitempty"`
	B_IsActive                int16   `thrift:",49,omitempty"`
	I_Area                    int16   `thrift:",50,omitempty"`
	I_DownloadType_img        int32   `thrift:",51,omitempty"`
	I_DownloadType_file       int32   `thrift:",52,omitempty"`
}

type GetSamSeckEventData struct {
	I_id                 int32   `thrift:",1,omitempty"`
	S_Description_KO     string  `thrift:",2,omitempty"`
	S_Description_EN     string  `thrift:",3,omitempty"`
	S_Description_JA     string  `thrift:",4,omitempty"`
	S_Description_ZH_CHS string  `thrift:",5,omitempty"`
	S_Description_ZH_CHT string  `thrift:",6,omitempty"`
	S_Description_VI     string  `thrift:",7,omitempty"`
	S_Description_ES     string  `thrift:",8,omitempty"`
	S_Description_IT     string  `thrift:",9,omitempty"`
	S_Description_ID     string  `thrift:",10,omitempty"`
	S_Description_TH     string  `thrift:",11,omitempty"`
	S_Description_PT     string  `thrift:",12,omitempty"`
	S_Description_HI     string  `thrift:",13,omitempty"`
	S_ConditionType      string  `thrift:",14,omitempty"`
	D_Condition          float64 `thrift:",15,omitempty"`
	I_MailRewardID       int64   `thrift:",16,omitempty"`
	B_IsActive           int16   `thrift:",17,omitempty"`
}

type GetTabListRetDataInfo struct {
	Tab_list []any `thrift:",1,omitempty"` // TODO
}

type MaintenanceList struct {
	Idx            int16  `thrift:",1,omitempty"`
	Title          string `thrift:",2,omitempty"`
	Start_datetime int32  `thrift:",3,omitempty"`
	End_datetime   int32  `thrift:",4,omitempty"`
	View_flg       int16  `thrift:",5,omitempty"`
}

type TabListData struct {
	Tap_idx   int16 `thrift:",1,omitempty"`
	Music_idx int32 `thrift:",2,omitempty"`
}

type Request struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        InitDataInfo           `thrift:",2,omitempty"`
	Common_data common_model.ParamData `thrift:",3,omitempty"`
}

type GetCharacterData struct {
	I_id                      int32   `thrift:",1,omitempty"`
	I_Area                    int16   `thrift:",2,omitempty"`
	S_BundleName              string  `thrift:",3,omitempty"`
	S_Name_KO                 string  `thrift:",4,omitempty"`
	S_Name_EN                 string  `thrift:",5,omitempty"`
	S_Name_JA                 string  `thrift:",6,omitempty"`
	S_Name_ZH_CHS             string  `thrift:",7,omitempty"`
	S_Name_ZH_CHT             string  `thrift:",8,omitempty"`
	S_Name_VI                 string  `thrift:",9,omitempty"`
	S_Name_ES                 string  `thrift:",10,omitempty"`
	S_Name_IT                 string  `thrift:",11,omitempty"`
	S_Name_ID                 string  `thrift:",12,omitempty"`
	S_Name_TH                 string  `thrift:",13,omitempty"`
	S_Name_PT                 string  `thrift:",14,omitempty"`
	S_Name_HI                 string  `thrift:",15,omitempty"`
	S_Description_KO          string  `thrift:",16,omitempty"`
	S_Description_EN          string  `thrift:",17,omitempty"`
	S_Description_JA          string  `thrift:",18,omitempty"`
	S_Description_ZH_CHS      string  `thrift:",19,omitempty"`
	S_Description_ZH_CHT      string  `thrift:",20,omitempty"`
	S_Description_VI          string  `thrift:",21,omitempty"`
	S_Description_ES          string  `thrift:",22,omitempty"`
	S_Description_IT          string  `thrift:",23,omitempty"`
	S_Description_ID          string  `thrift:",24,omitempty"`
	S_Description_TH          string  `thrift:",25,omitempty"`
	S_Description_PT          string  `thrift:",26,omitempty"`
	S_Description_HI          string  `thrift:",27,omitempty"`
	S_ResourceName            string  `thrift:",28,omitempty"`
	I_MaxLevel                int64   `thrift:",29,omitempty"`
	D_LevelRate               float64 `thrift:",30,omitempty"`
	D_LateCost                float64 `thrift:",31,omitempty"`
	D_Cost                    float64 `thrift:",32,omitempty"`
	D_CostIncreaseValue       float64 `thrift:",33,omitempty"`
	D_CostValueConstant       float64 `thrift:",34,omitempty"`
	D_Amount                  float64 `thrift:",35,omitempty"`
	D_AmountIncreaseValue     float64 `thrift:",36,omitempty"`
	D_AmountValueConstant     float64 `thrift:",37,omitempty"`
	D_BonusAmountIncreaseRate float64 `thrift:",38,omitempty"`
	I_UnlockLevel             int64   `thrift:",39,omitempty"`
	D_UnlockCost              float64 `thrift:",40,omitempty"`
	I_RequirementFollowerId_1 int64   `thrift:",41,omitempty"`
	I_RequirementFollowerId_2 int64   `thrift:",42,omitempty"`
	I_RequirementPropId_1     int64   `thrift:",43,omitempty"`
	I_RequirementPropId_2     int64   `thrift:",44,omitempty"`
	I_RequirementPropCount    int64   `thrift:",45,omitempty"`
	B_IsActive                int16   `thrift:",46,omitempty"`
}

type GetSelectRewardData struct {
	I_id               int32 `thrift:",1,omitempty"`
	I_GroupId          int32 `thrift:",2,omitempty"`
	I_RewardGroupId    int32 `thrift:",3,omitempty"`
	I_AltRewardGroupId int32 `thrift:",4,omitempty"`
	B_IsActive         int16 `thrift:",5,omitempty"`
}

type GetTabList struct {
	Call string       `thrift:",1,omitempty"`
	Data InitDataInfo `thrift:",2,omitempty"`
}

type StoreListRetDataInfo struct {
	Store_idx  int16       `thrift:",1,omitempty"`
	Store_type int16       `thrift:",2,omitempty"`
	Title      string      `thrift:",3,omitempty"`
	Sub_title  string      `thrift:",4,omitempty"`
	Img_url    string      `thrift:",5,omitempty"`
	Store_list map[any]any `thrift:",6,omitempty"` // TODO
}

type GetFollowerData struct {
	I_id                      int32   `thrift:",1,omitempty"`
	S_Name_KO                 string  `thrift:",2,omitempty"`
	S_Name_EN                 string  `thrift:",3,omitempty"`
	S_Name_JA                 string  `thrift:",4,omitempty"`
	S_Name_ZH_CHS             string  `thrift:",5,omitempty"`
	S_Name_ZH_CHT             string  `thrift:",6,omitempty"`
	S_Name_VI                 string  `thrift:",7,omitempty"`
	S_Name_ES                 string  `thrift:",8,omitempty"`
	S_Name_IT                 string  `thrift:",9,omitempty"`
	S_Name_ID                 string  `thrift:",10,omitempty"`
	S_Name_TH                 string  `thrift:",11,omitempty"`
	S_Name_PT                 string  `thrift:",12,omitempty"`
	S_Name_HI                 string  `thrift:",13,omitempty"`
	S_Description_KO          string  `thrift:",14,omitempty"`
	S_Description_EN          string  `thrift:",15,omitempty"`
	S_Description_JA          string  `thrift:",16,omitempty"`
	S_Description_ZH_CHS      string  `thrift:",17,omitempty"`
	S_Description_ZH_CHT      string  `thrift:",18,omitempty"`
	S_Description_VI          string  `thrift:",19,omitempty"`
	S_Description_ES          string  `thrift:",20,omitempty"`
	S_Description_IT          string  `thrift:",21,omitempty"`
	S_Description_ID          string  `thrift:",22,omitempty"`
	S_Description_TH          string  `thrift:",23,omitempty"`
	S_Description_PT          string  `thrift:",24,omitempty"`
	S_Description_HI          string  `thrift:",25,omitempty"`
	S_ResourceName            string  `thrift:",26,omitempty"`
	I_MaxLevel                int64   `thrift:",27,omitempty"`
	D_LevelRate               float64 `thrift:",28,omitempty"`
	D_LateCost                int64   `thrift:",29,omitempty"`
	D_Cost                    float64 `thrift:",30,omitempty"`
	D_CostIncreaseValue       float64 `thrift:",31,omitempty"`
	D_CostValueConstant       int64   `thrift:",32,omitempty"`
	D_Amount                  float64 `thrift:",33,omitempty"`
	D_AmountIncreaseValue     float64 `thrift:",34,omitempty"`
	D_AmountValueConstant     int64   `thrift:",35,omitempty"`
	D_BonusAmountIncreaseRate int64   `thrift:",36,omitempty"`
	I_UnlockLevel             int64   `thrift:",37,omitempty"`
	D_UnlockCost              int64   `thrift:",38,omitempty"`
	I_RequirementFollowerId_1 int64   `thrift:",39,omitempty"`
	I_RequirementFollowerId_2 int64   `thrift:",40,omitempty"`
	I_RequirementPropId_1     int64   `thrift:",41,omitempty"`
	I_RequirementPropId_2     int64   `thrift:",42,omitempty"`
	I_RequirementPropCount    int64   `thrift:",43,omitempty"`
	B_IsActive                int16   `thrift:",44,omitempty"`
	I_Area                    int16   `thrift:",45,omitempty"`
	I_DownloadType            int32   `thrift:",46,omitempty"`
	I_ProfileID               int32   `thrift:",47,omitempty"`
}

type GetPercentData struct {
	I_GroupID        int32 `thrift:",1,omitempty"`
	I_RewardType     int32 `thrift:",2,omitempty"`
	I_RewardID       int32 `thrift:",3,omitempty"`
	L_RewardQuantity int32 `thrift:",4,omitempty"`
	B_IsActive       int16 `thrift:",5,omitempty"`
	I_id             int32 `thrift:",6,omitempty"`
}

type GetRewardGroupData struct {
	I_id               int64 `thrift:",1,omitempty"`
	I_Group            int64 `thrift:",2,omitempty"`
	I_RewardType       int64 `thrift:",3,omitempty"`
	I_RewardID         int64 `thrift:",4,omitempty"`
	L_RewardQuantity   int64 `thrift:",5,omitempty"`
	I_BuyFirstQuantity int64 `thrift:",6,omitempty"`
}

type GetGameDataList struct {
	Call        string                  `thrift:",1,omitempty"`
	Data        GetGameDataListDataInfo `thrift:",2,omitempty"`
	Sub_mode    string                  `thrift:",3,omitempty"`
	Common_data common_model.ParamData  `thrift:",4,omitempty"`
}

type InitRetDataInfo struct {
	Idx      int16  `thrift:",1,omitempty"`
	Game_url string `thrift:",2,omitempty"`
	Cdn_url  string `thrift:",3,omitempty"`
}

type GetAdLevelData struct {
	I_GroupID     int32 `thrift:",1,omitempty"`
	I_Level       int32 `thrift:",2,omitempty"`
	I_RequireEXP  int32 `thrift:",3,omitempty"`
	I_RewardGroup int32 `thrift:",4,omitempty"`
	B_IsActive    int16 `thrift:",5,omitempty"`
	I_id          int32 `thrift:",6,omitempty"`
}

type GetBuffData struct {
	I_id               int32  `thrift:",1,omitempty"`
	S_ResourceName     string `thrift:",2,omitempty"`
	S_Title_KO         string `thrift:",3,omitempty"`
	S_Title_EN         string `thrift:",4,omitempty"`
	S_Title_JA         string `thrift:",5,omitempty"`
	S_Title_ZH_CHS     string `thrift:",6,omitempty"`
	S_Title_ZH_CHT     string `thrift:",7,omitempty"`
	S_Title_VI         string `thrift:",8,omitempty"`
	S_Title_ES         string `thrift:",9,omitempty"`
	S_Title_IT         string `thrift:",10,omitempty"`
	S_Title_ID         string `thrift:",11,omitempty"`
	S_Title_TH         string `thrift:",12,omitempty"`
	S_Title_PT         string `thrift:",13,omitempty"`
	S_Title_HI         string `thrift:",14,omitempty"`
	S_Buff_Desc_KO     string `thrift:",15,omitempty"`
	S_Buff_Desc_EN     string `thrift:",16,omitempty"`
	S_Buff_Desc_JA     string `thrift:",17,omitempty"`
	S_Buff_Desc_ZH_CHS string `thrift:",18,omitempty"`
	S_Buff_Desc_ZH_CHT string `thrift:",19,omitempty"`
	S_Buff_Desc_VI     string `thrift:",20,omitempty"`
	S_Buff_Desc_ES     string `thrift:",21,omitempty"`
	S_Buff_Desc_IT     string `thrift:",22,omitempty"`
	S_Buff_Desc_ID     string `thrift:",23,omitempty"`
	S_Buff_Desc_TH     string `thrift:",24,omitempty"`
	S_Buff_Desc_PT     string `thrift:",25,omitempty"`
	S_Buff_Desc_HI     string `thrift:",26,omitempty"`
	I_Type             int16  `thrift:",27,omitempty"`
	I_Retention_Time   int64  `thrift:",28,omitempty"`
}

type GetEventRewardListRetDataInfo struct {
	Reward_list []any `thrift:",1,omitempty"` // TODO
	Group_idx   int32 `thrift:",2,omitempty"`
}

type GetSubscribeList struct {
	I_id              int64  `thrift:",1,omitempty"`
	I_Area            int16  `thrift:",2,omitempty"`
	S_Type            string `thrift:",3,omitempty"`
	I_TimeLimit       int16  `thrift:",4,omitempty"`
	I_StartYear       int32  `thrift:",5,omitempty"`
	I_RepeatMonth     int32  `thrift:",6,omitempty"`
	I_MonthGroupIndex int32  `thrift:",7,omitempty"`
	B_IsActive        int16  `thrift:",8,omitempty"`
}

type MainReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        MainRetDataInfo              `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type MoreGames struct {
	Call string            `thrift:",1,omitempty"`
	Data MoreGamesDataInfo `thrift:",2,omitempty"`
}

type GetGuitarData struct {
	I_id                        int32   `thrift:",1,omitempty"`
	I_Area                      int16   `thrift:",2,omitempty"`
	S_Title_KO                  string  `thrift:",3,omitempty"`
	S_Title_EN                  string  `thrift:",4,omitempty"`
	S_Title_JA                  string  `thrift:",5,omitempty"`
	S_Title_ZH_CHS              string  `thrift:",6,omitempty"`
	S_Title_ZH_CHT              string  `thrift:",7,omitempty"`
	S_Title_VI                  string  `thrift:",8,omitempty"`
	S_Title_ES                  string  `thrift:",9,omitempty"`
	S_Title_IT                  string  `thrift:",10,omitempty"`
	S_Title_ID                  string  `thrift:",11,omitempty"`
	S_Title_TH                  string  `thrift:",12,omitempty"`
	S_Title_PT                  string  `thrift:",13,omitempty"`
	S_Title_HI                  string  `thrift:",14,omitempty"`
	S_Description_KO            string  `thrift:",15,omitempty"`
	S_Description_EN            string  `thrift:",16,omitempty"`
	S_Description_JA            string  `thrift:",17,omitempty"`
	S_Description_ZH_CHS        string  `thrift:",18,omitempty"`
	S_Description_ZH_CHT        string  `thrift:",19,omitempty"`
	S_Description_VI            string  `thrift:",20,omitempty"`
	S_Description_ES            string  `thrift:",21,omitempty"`
	S_Description_IT            string  `thrift:",22,omitempty"`
	S_Description_ID            string  `thrift:",23,omitempty"`
	S_Description_TH            string  `thrift:",24,omitempty"`
	S_Description_PT            string  `thrift:",25,omitempty"`
	S_Description_HI            string  `thrift:",26,omitempty"`
	S_ResourceName              string  `thrift:",27,omitempty"`
	S_GuitarName                string  `thrift:",28,omitempty"`
	I_GuitarType                int16   `thrift:",29,omitempty"`
	S_GoodsType                 string  `thrift:",30,omitempty"`
	S_Cost                      string  `thrift:",31,omitempty"`
	I_RequirementCharacterLevel int64   `thrift:",32,omitempty"`
	D_IncrementValue            float64 `thrift:",33,omitempty"`
	I_AcquisitionType           int64   `thrift:",34,omitempty"`
	I_AcquisitionId             int64   `thrift:",35,omitempty"`
	I_SortOrder                 int64   `thrift:",36,omitempty"`
	B_IsActive                  int16   `thrift:",37,omitempty"`
	I_DownloadType              int16   `thrift:",38,omitempty"`
}

type GetNoticeListReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        []any                        `thrift:",5,omitempty"` // TODO
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type GetPassGoodsShopData struct {
	I_id            int32 `thrift:",1,omitempty"`
	I_TicketID      int32 `thrift:",2,omitempty"`
	I_GoodsTicketID int32 `thrift:",3,omitempty"`
	I_ShopID        int32 `thrift:",4,omitempty"`
	B_IsActive      int16 `thrift:",5,omitempty"`
	I_SaleShopID    int32 `thrift:",6,omitempty"`
}

type GetServerTimeRetDataInfo struct {
	Time     int64 `thrift:",1,omitempty"`
	Datetime int64 `thrift:",2,omitempty"`
}

type Main struct {
	Call     string       `thrift:",1,omitempty"`
	Data     MainDataInfo `thrift:",2,omitempty"`
	Sub_mode string       `thrift:",3,omitempty"`
}

type NewStoreListData struct {
	Store_idx int16 `thrift:",1,omitempty"`
	Music_idx int32 `thrift:",2,omitempty"`
}

type GetChThirdStageData struct {
	I_id                   int32  `thrift:",1,omitempty"`
	I_Chapter              int32  `thrift:",2,omitempty"`
	I_StageIndex           int32  `thrift:",3,omitempty"`
	I_RewardGroupID        int32  `thrift:",4,omitempty"`
	I_GoalScore1           int32  `thrift:",5,omitempty"`
	I_GoalScore2           int32  `thrift:",6,omitempty"`
	I_GoalScore3           int32  `thrift:",7,omitempty"`
	S_BonusProfileID       string `thrift:",8,omitempty"`
	S_Name_KO              string `thrift:",9,omitempty"`
	S_Name_EN              string `thrift:",10,omitempty"`
	S_Name_JA              string `thrift:",11,omitempty"`
	S_Name_ZH_CHS          string `thrift:",12,omitempty"`
	S_Name_ZH_CHT          string `thrift:",13,omitempty"`
	S_Name_VI              string `thrift:",14,omitempty"`
	S_Name_ES              string `thrift:",15,omitempty"`
	S_Name_IT              string `thrift:",16,omitempty"`
	S_Name_ID              string `thrift:",17,omitempty"`
	S_Name_TH              string `thrift:",18,omitempty"`
	S_Name_PT              string `thrift:",19,omitempty"`
	S_Name_HI              string `thrift:",20,omitempty"`
	S_Description_KO       string `thrift:",21,omitempty"`
	S_Description_EN       string `thrift:",22,omitempty"`
	S_Description_JA       string `thrift:",23,omitempty"`
	S_Description_ZH_CHS   string `thrift:",24,omitempty"`
	S_Description_ZH_CHT   string `thrift:",25,omitempty"`
	S_Description_VI       string `thrift:",26,omitempty"`
	S_Description_ES       string `thrift:",27,omitempty"`
	S_Description_IT       string `thrift:",28,omitempty"`
	S_Description_ID       string `thrift:",29,omitempty"`
	S_Description_TH       string `thrift:",30,omitempty"`
	S_Description_PT       string `thrift:",31,omitempty"`
	S_Description_HI       string `thrift:",32,omitempty"`
	B_IsActive             int16  `thrift:",33,omitempty"`
	S_StageType            string `thrift:",34,omitempty"`
	I_StoryID              int32  `thrift:",35,omitempty"`
	I_BonusMusicID         int32  `thrift:",36,omitempty"`
	S_PercentRewardGroupID string `thrift:",37,omitempty"`
}

type GetFollowerProfileLevelData struct {
	I_id                int32  `thrift:",1,omitempty"`
	I_ProfileID         int32  `thrift:",2,omitempty"`
	I_Level             int32  `thrift:",3,omitempty"`
	D_RequireEXP        int64  `thrift:",4,omitempty"`
	I_RewardGroup       int32  `thrift:",5,omitempty"`
	S_UnlockInformation string `thrift:",6,omitempty"`
	I_AddCandy          int32  `thrift:",7,omitempty"`
	I_UnlockStoryID     int32  `thrift:",8,omitempty"`
	B_IsActive          int16  `thrift:",9,omitempty"`
}

type GetSubscribePassReward struct {
	I_id              int64 `thrift:",1,omitempty"`
	I_Group           int64 `thrift:",2,omitempty"`
	I_Step            int16 `thrift:",3,omitempty"`
	I_Goal            int64 `thrift:",4,omitempty"`
	I_FreeRewardGroup int64 `thrift:",5,omitempty"`
	I_PaidRewardGroup int64 `thrift:",6,omitempty"`
}

type AvatarListRetDataInfo struct {
	Idx           int16  `thrift:",1,omitempty"`
	Resource_name string `thrift:",2,omitempty"`
	Down_status   int16  `thrift:",3,omitempty"`
	Buy_type      int16  `thrift:",4,omitempty"`
	Price         int32  `thrift:",5,omitempty"`
	Buy_type2     int16  `thrift:",6,omitempty"`
	Price2        int32  `thrift:",7,omitempty"`
	Value         int32  `thrift:",8,omitempty"`
	Event_point   int64  `thrift:",9,omitempty"`
}

type GetEventListData struct {
	Idx        int32  `thrift:",1,omitempty"`
	Event_type string `thrift:",2,omitempty"`
	Start_date string `thrift:",3,omitempty"`
	End_date   string `thrift:",4,omitempty"`
	B_IsActive int16  `thrift:",5,omitempty"`
}

type GetShopData struct {
	I_id                    int64  `thrift:",1,omitempty"`
	S_ProductID_ios         string `thrift:",2,omitempty"`
	S_ProductID_aos         string `thrift:",3,omitempty"`
	I_ShopCategory          int64  `thrift:",4,omitempty"`
	I_ProductType           int64  `thrift:",5,omitempty"`
	S_ResourceName          string `thrift:",6,omitempty"`
	S_Title_KO              string `thrift:",7,omitempty"`
	S_Title_EN              string `thrift:",8,omitempty"`
	S_Title_JA              string `thrift:",9,omitempty"`
	S_Title_ZH_CHS          string `thrift:",10,omitempty"`
	S_Title_ZH_CHT          string `thrift:",11,omitempty"`
	S_Title_VI              string `thrift:",12,omitempty"`
	S_Title_ES              string `thrift:",13,omitempty"`
	S_Title_IT              string `thrift:",14,omitempty"`
	S_Title_ID              string `thrift:",15,omitempty"`
	S_Title_TH              string `thrift:",16,omitempty"`
	S_Title_PT              string `thrift:",17,omitempty"`
	S_Title_HI              string `thrift:",18,omitempty"`
	S_Description_KO        string `thrift:",19,omitempty"`
	S_Description_EN        string `thrift:",20,omitempty"`
	S_Description_JA        string `thrift:",21,omitempty"`
	S_Description_ZH_CHS    string `thrift:",22,omitempty"`
	S_Description_ZH_CHT    string `thrift:",23,omitempty"`
	S_Description_VI        string `thrift:",24,omitempty"`
	S_Description_ES        string `thrift:",25,omitempty"`
	S_Description_IT        string `thrift:",26,omitempty"`
	S_Description_ID        string `thrift:",27,omitempty"`
	S_Description_TH        string `thrift:",28,omitempty"`
	S_Description_PT        string `thrift:",29,omitempty"`
	S_Description_HI        string `thrift:",30,omitempty"`
	I_RewardGroup           int64  `thrift:",31,omitempty"`
	I_Tag                   int64  `thrift:",32,omitempty"`
	I_SellType              int64  `thrift:",33,omitempty"`
	I_SellValue             int64  `thrift:",34,omitempty"`
	B_IsActive              int16  `thrift:",35,omitempty"`
	I_Condition             int64  `thrift:",36,omitempty"`
	I_ConditionValue        int64  `thrift:",37,omitempty"`
	S_StorePrice_AOS        string `thrift:",38,omitempty"`
	S_StorePrice_IOS        string `thrift:",39,omitempty"`
	I_SortIndex             int16  `thrift:",40,omitempty"`
	S_AreaList              string `thrift:",41,omitempty"`
	S_AltResourceName       string `thrift:",42,omitempty"`
	S_AltDescription_KO     string `thrift:",43,omitempty"`
	S_AltDescription_EN     string `thrift:",44,omitempty"`
	S_AltDescription_JA     string `thrift:",45,omitempty"`
	S_AltDescription_ZH_CHS string `thrift:",46,omitempty"`
	S_AltDescription_ZH_CHT string `thrift:",47,omitempty"`
	S_AltDescription_VI     string `thrift:",48,omitempty"`
	S_AltDescription_ES     string `thrift:",49,omitempty"`
	S_AltDescription_IT     string `thrift:",50,omitempty"`
	S_AltDescription_ID     string `thrift:",51,omitempty"`
	S_AltDescription_TH     string `thrift:",52,omitempty"`
	S_AltDescription_PT     string `thrift:",53,omitempty"`
	S_AltDescription_HI     string `thrift:",54,omitempty"`
	S_Kor_StorePrice_AOS    string `thrift:",55,omitempty"`
	S_Kor_StorePrice_IOS    string `thrift:",56,omitempty"`
	B_IsLimitTime           int16  `thrift:",57,omitempty"`
	S_UIStartTime           string `thrift:",58,omitempty"`
	S_StartTime             string `thrift:",59,omitempty"`
	S_EndTime               string `thrift:",60,omitempty"`
}

type NewStoreListRetDataInfo struct {
	Store_idx    int16           `thrift:",1,omitempty"`
	Store_type   int16           `thrift:",2,omitempty"`
	Buy_type     int16           `thrift:",3,omitempty"`
	Prev_price   int64           `thrift:",4,omitempty"`
	Price        int64           `thrift:",5,omitempty"`
	Title        string          `thrift:",6,omitempty"`
	Sub_title    string          `thrift:",7,omitempty"`
	Img_url      string          `thrift:",8,omitempty"`
	Store_list   map[any]any     `thrift:",9,omitempty"` // TODO
	Product_data NewStoreProduct `thrift:",10,omitempty"`
	Is_sale      int16           `thrift:",11,omitempty"`
}

type GetEventRewardListData struct {
	Idx                int64  `thrift:",1,omitempty"`
	Event_name         string `thrift:",2,omitempty"`
	Event_type         string `thrift:",3,omitempty"`
	Reward_idx         int64  `thrift:",4,omitempty"`
	Reward_num         int32  `thrift:",5,omitempty"`
	Reward_type        int32  `thrift:",6,omitempty"`
	Reward_id          int32  `thrift:",7,omitempty"`
	Reward_value       int32  `thrift:",8,omitempty"`
	Reward_flg         string `thrift:",9,omitempty"`
	Get_date           int32  `thrift:",10,omitempty"`
	S_CustomIconType   string `thrift:",11,omitempty"`
	S_CustomIconSprite string `thrift:",12,omitempty"`
}

type GetSubscribePassRewardInformation struct {
	I_id                 int32  `thrift:",1,omitempty"`
	S_ResourceName       string `thrift:",2,omitempty"`
	S_Description_KO     string `thrift:",3,omitempty"`
	S_Description_EN     string `thrift:",4,omitempty"`
	S_Description_JA     string `thrift:",5,omitempty"`
	S_Description_ZH_CHS string `thrift:",6,omitempty"`
	S_Description_ZH_CHT string `thrift:",7,omitempty"`
	S_Description_VI     string `thrift:",8,omitempty"`
	S_Description_ES     string `thrift:",9,omitempty"`
	S_Description_IT     string `thrift:",10,omitempty"`
	S_Description_ID     string `thrift:",11,omitempty"`
	S_Description_TH     string `thrift:",12,omitempty"`
	S_Description_PT     string `thrift:",13,omitempty"`
	S_Description_HI     string `thrift:",14,omitempty"`
	B_IsActive           int16  `thrift:",15,omitempty"`
}

type GetUpdateTimeRetDataInfo struct {
	Upd_time int64 `thrift:",1,omitempty"`
}

type GetCollectionDataInfo struct {
}

type GetDefaultSettingRetDataInfo struct {
	Ad_count          int32   `thrift:",1,omitempty"`
	Ad_probability    float64 `thrift:",2,omitempty"`
	Challenge_mission float64 `thrift:",3,omitempty"`
	Seed_value        int16   `thrift:",4,omitempty"`
	Rank_up_limit     int16   `thrift:",5,omitempty"`
	Rank_down_limit   int16   `thrift:",6,omitempty"`
	Reward_rate       float64 `thrift:",7,omitempty"`
	Tour_btn_name     string  `thrift:",8,omitempty"`
	Event_btn_name    string  `thrift:",9,omitempty"`
	Event_rank_url    string  `thrift:",10,omitempty"`
	Tour_intro_movie  string  `thrift:",11,omitempty"`
	Tour_bg_movie     string  `thrift:",12,omitempty"`
	Event_bg_movie    string  `thrift:",13,omitempty"`
}

type GetNoticeListDataInfo struct {
	Device_uuid string `thrift:",1,omitempty"`
}

type GetServerTimeDataInfo struct {
	Device_uuid string `thrift:",1,omitempty"`
}

type GetEventRewardList struct {
	Call        string                     `thrift:",1,omitempty"`
	Data        GetEventRewardListDataInfo `thrift:",2,omitempty"`
	Sub_mode    string                     `thrift:",3,omitempty"`
	Common_data common_model.ParamData     `thrift:",4,omitempty"`
}

type GetSystemNotificationData struct {
	I_Id                 int64  `thrift:",1,omitempty"`
	S_DescriptionType    string `thrift:",2,omitempty"`
	I_Type               string `thrift:",3,omitempty"`
	S_Factor             string `thrift:",4,omitempty"`
	S_Description_KO     string `thrift:",5,omitempty"`
	S_Description_EN     string `thrift:",6,omitempty"`
	S_Description_JA     string `thrift:",7,omitempty"`
	S_Description_ZH_CHS string `thrift:",8,omitempty"`
	S_Description_ZH_CHT string `thrift:",9,omitempty"`
	S_Description_VI     string `thrift:",10,omitempty"`
	S_Description_ES     string `thrift:",11,omitempty"`
	S_Description_IT     string `thrift:",12,omitempty"`
	S_Description_ID     string `thrift:",13,omitempty"`
	S_Description_TH     string `thrift:",14,omitempty"`
	S_Description_PT     string `thrift:",15,omitempty"`
	S_Description_HI     string `thrift:",16,omitempty"`
}

type NewStoreList struct {
	Call string               `thrift:",1,omitempty"`
	Data NewStoreListDataInfo `thrift:",2,omitempty"`
}

type StoreList struct {
	Call string            `thrift:",1,omitempty"`
	Data StoreListDataInfo `thrift:",2,omitempty"`
}

type StoreListReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        map[any]any                  `thrift:",4,omitempty"` // TODO
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type GetLocalPushRetDataInfo struct {
	I_LocalPushID           int32  `thrift:",1,omitempty"`
	S_LocalPushRegisterType int16  `thrift:",2,omitempty"`
	I_TimeForSeconds        int64  `thrift:",3,omitempty"`
	S_JsonMessage_ko        string `thrift:",4,omitempty"`
	S_JsonMessage_en        string `thrift:",5,omitempty"`
	S_JsonMessage_ja        string `thrift:",6,omitempty"`
}

type GetTabListDataInfo struct {
	Type        string `thrift:",1,omitempty"`
	Device_uuid string `thrift:",2,omitempty"`
}

type DefaultSettingListDataInfo struct {
	Device_uuid string `thrift:",1,omitempty"`
	Os          int16  `thrift:",2,omitempty"`
}

type StoreListDataInfo struct {
	Type string `thrift:",1,omitempty"`
}

type GetCollectionRetDataInfo struct {
	Idx      int32 `thrift:",1,omitempty"`
	Buy_type int16 `thrift:",2,omitempty"`
	Price    int64 `thrift:",3,omitempty"`
	Item     []any `thrift:",4,omitempty"` // TODO
}

type GetFollowerProfileData struct {
	I_id              int32  `thrift:",1,omitempty"`
	S_Name_KO         string `thrift:",2,omitempty"`
	S_Name_EN         string `thrift:",3,omitempty"`
	S_Name_JA         string `thrift:",4,omitempty"`
	S_Name_ZH_CHS     string `thrift:",5,omitempty"`
	S_Name_ZH_CHT     string `thrift:",6,omitempty"`
	S_Name_VI         string `thrift:",7,omitempty"`
	S_Name_ES         string `thrift:",8,omitempty"`
	S_Name_IT         string `thrift:",9,omitempty"`
	S_Name_ID         string `thrift:",10,omitempty"`
	S_Name_TH         string `thrift:",11,omitempty"`
	S_Name_PT         string `thrift:",12,omitempty"`
	S_Name_HI         string `thrift:",13,omitempty"`
	B_IsActive        int16  `thrift:",14,omitempty"`
	S_GiftItemID      string `thrift:",15,omitempty"`
	B_IsVisible       int16  `thrift:",16,omitempty"`
	S_BackgroundColor string `thrift:",17,omitempty"`
}

type GetPropData struct {
	I_id                 int32  `thrift:",1,omitempty"`
	S_Name_KO            string `thrift:",2,omitempty"`
	S_Name_EN            string `thrift:",3,omitempty"`
	S_Name_JA            string `thrift:",4,omitempty"`
	S_Name_ZH_CHS        string `thrift:",5,omitempty"`
	S_Name_ZH_CHT        string `thrift:",6,omitempty"`
	S_Name_VI            string `thrift:",7,omitempty"`
	S_Name_ES            string `thrift:",8,omitempty"`
	S_Name_IT            string `thrift:",9,omitempty"`
	S_Name_ID            string `thrift:",10,omitempty"`
	S_Name_TH            string `thrift:",11,omitempty"`
	S_Name_PT            string `thrift:",12,omitempty"`
	S_Name_HI            string `thrift:",13,omitempty"`
	S_Description_KO     string `thrift:",14,omitempty"`
	S_Description_EN     string `thrift:",15,omitempty"`
	S_Description_JA     string `thrift:",16,omitempty"`
	S_Description_ZH_CHS string `thrift:",17,omitempty"`
	S_Description_ZH_CHT string `thrift:",18,omitempty"`
	S_Description_VI     string `thrift:",19,omitempty"`
	S_Description_ES     string `thrift:",20,omitempty"`
	S_Description_IT     string `thrift:",21,omitempty"`
	S_Description_ID     string `thrift:",22,omitempty"`
	S_Description_TH     string `thrift:",23,omitempty"`
	S_Description_PT     string `thrift:",24,omitempty"`
	S_Description_HI     string `thrift:",25,omitempty"`
	S_ResourceName       string `thrift:",26,omitempty"`
	I_RequirementPropId  int64  `thrift:",27,omitempty"`
	I_Area               int64  `thrift:",28,omitempty"`
	B_IsActive           int64  `thrift:",29,omitempty"`
	I_MaxLevel           int16  `thrift:",30,omitempty"`
}

type AvatarListReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        []any                        `thrift:",4,omitempty"` // TODO
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type DefaultSettingListRetDataInfo struct {
	Status       string                   `thrift:",1,omitempty"`
	Setting_list []DefaultSettingDataList `thrift:",2,omitempty"` // TODO
}

type GetPropLevelData struct {
	I_id         int32   `thrift:",1,omitempty"`
	I_PropId     int64   `thrift:",2,omitempty"`
	I_Level      int64   `thrift:",3,omitempty"`
	D_Cost       float64 `thrift:",4,omitempty"`
	I_PerksId    int64   `thrift:",5,omitempty"`
	D_PerksValue float64 `thrift:",6,omitempty"`
	B_IsActive   int16   `thrift:",7,omitempty"`
}

type MoreGamesDataInfo struct {
	Locale     string `thrift:",1,omitempty"`
	Adset_type string `thrift:",2,omitempty"`
	Uuid       string `thrift:",3,omitempty"`
}

type GetSkillData struct {
	I_id                      int32   `thrift:",1,omitempty"`
	I_Area                    int16   `thrift:",2,omitempty"`
	S_Name_KO                 string  `thrift:",3,omitempty"`
	S_Name_EN                 string  `thrift:",4,omitempty"`
	S_Name_JA                 string  `thrift:",5,omitempty"`
	S_Name_ZH_CHS             string  `thrift:",6,omitempty"`
	S_Name_ZH_CHT             string  `thrift:",7,omitempty"`
	S_Name_VI                 string  `thrift:",8,omitempty"`
	S_Name_ES                 string  `thrift:",9,omitempty"`
	S_Name_IT                 string  `thrift:",10,omitempty"`
	S_Name_ID                 string  `thrift:",11,omitempty"`
	S_Name_TH                 string  `thrift:",12,omitempty"`
	S_Name_PT                 string  `thrift:",13,omitempty"`
	S_Name_HI                 string  `thrift:",14,omitempty"`
	S_Description_KO          string  `thrift:",15,omitempty"`
	S_Description_EN          string  `thrift:",16,omitempty"`
	S_Description_JA          string  `thrift:",17,omitempty"`
	S_Description_ZH_CHS      string  `thrift:",18,omitempty"`
	S_Description_ZH_CHT      string  `thrift:",19,omitempty"`
	S_Description_VI          string  `thrift:",20,omitempty"`
	S_Description_ES          string  `thrift:",21,omitempty"`
	S_Description_IT          string  `thrift:",22,omitempty"`
	S_Description_ID          string  `thrift:",23,omitempty"`
	S_Description_TH          string  `thrift:",24,omitempty"`
	S_Description_PT          string  `thrift:",25,omitempty"`
	S_Description_HI          string  `thrift:",26,omitempty"`
	S_ResourceName            string  `thrift:",27,omitempty"`
	I_MaxLevel                int64   `thrift:",28,omitempty"`
	S_GoodsType               string  `thrift:",29,omitempty"`
	I_Cost                    float64 `thrift:",30,omitempty"`
	F_CostIncreaseValue       float64 `thrift:",31,omitempty"`
	F_SkillValue              float64 `thrift:",32,omitempty"`
	F_SkillIncreaseValue      float64 `thrift:",33,omitempty"`
	F_SkillTime               float64 `thrift:",34,omitempty"`
	F_SkillTimeIncreaseValue  float64 `thrift:",35,omitempty"`
	F_Cooltime                float64 `thrift:",36,omitempty"`
	I_UnlockLevel             int64   `thrift:",37,omitempty"`
	I_UnlockCost              int64   `thrift:",38,omitempty"`
	I_RequirementSkillId_1    int64   `thrift:",39,omitempty"`
	I_RequirementSkillId_2    int64   `thrift:",40,omitempty"`
	I_RequirementFollowerId_1 int64   `thrift:",41,omitempty"`
	I_RequirementFollowerId_2 int64   `thrift:",42,omitempty"`
	I_RequirementPropId_1     int64   `thrift:",43,omitempty"`
	I_RequirementPropId_2     int64   `thrift:",44,omitempty"`
	I_RequirementPropCount    int64   `thrift:",45,omitempty"`
}
