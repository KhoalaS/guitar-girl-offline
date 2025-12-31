package model

type Error struct {
	Is_success      bool
	Err_code        int32
	Err_message     string
	Is_debug        bool
	Is_achieve_noti bool
	Mission_notis   []any
}
type MoreGamesLog struct {
	Call string
	Data MoreGamesLogDataInfo
}
type MoreGamesLogDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	App_id      int32
}
type MoreGamesLogRetDataInfo struct {
	Ret string
}
type MoreGamesLogReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        MoreGamesLogRetDataInfo
	Maintenance MaintenanceData
}
type Rect struct {
	Left   float64
	Top    float64
	Width  float64
	Height float64
}
type Request struct {
	Call        string
	Data        InitDataInfo
	Common_data ParamData
}
type StageFollowerProfileScore struct {
	I_id        int32
	Score       int32
	Bonus_score int32
}
type UserAdLevel struct {
	I_id    int32
	I_Level int32
	I_EXP   int32
}
type UserAdList struct {
	I_id           int32
	I_Count        int32
	I_TotalCount   int32
	I_LastViewTime int32
}
type UserApData struct {
	I_Ap         int32
	I_FullApTime int32
	I_MaxAp      int32
}
type UserChThirdChapterReward struct {
	I_id              int32
	I_ReceivedReward1 int32
	I_ReceivedReward2 int32
	I_ReceivedReward3 int32
}
type UserChThirdStage struct {
	I_id         int32
	I_ChapterId  int32
	I_StageIndex int32
	I_Star       int32
}
type UserFollowerGiftItem struct {
	I_id    int32
	I_Value int32
}
type UserFollowerProfile struct {
	I_id       int32
	I_Level    int32
	D_Exp      int64
	I_AddCandy int32
}
type UserFollowerProfileReward struct {
	I_id          int32
	I_RewardLevel int32
}
type UserSelectReward struct {
	I_GroupId          int32
	I_RewardGroupId    int32
	I_AltRewardGroupId int32
}
type UserTutorial struct {
	I_id int32
}
type Vector3 struct {
	X float64
	Y float64
	Z float64
}
type AdViewLog struct {
	Call        string
	Data        AdViewLogDataInfo
	Common_data ParamData
}
type AdViewLogDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Ad_type     string
	Ad_name     string
}
type AdViewLogRetDataInfo struct {
	Status string
}
type AdViewLogReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        AdViewLogRetDataInfo
	Maintenance MaintenanceData
}
type AlbumListInfo struct {
	Album_idx  int64
	Country    string
	Date_issue string
	Cdn_dir    string
	Album_img  string
	Language   map[any]any
}
type AlbumListLanguage struct {
	Title        string
	Company      string
	Album_artist string
}
type AvatarList struct {
	Call string
	Data AvatarListDataInfo
}
type AvatarListDataInfo struct {
	Type string
}
type AvatarListRetDataInfo struct {
	Idx           int16
	Resource_name string
	Down_status   int16
	Buy_type      int16
	Price         int32
	Buy_type2     int16
	Price2        int32
	Value         int32
	Event_point   int64
}
type AvatarListReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        []any
	Maintenance MaintenanceData
}
type Bookmark_list struct {
	Music_idx int32
	Flag      int16
}
type BuyAvatar struct {
	Call string
	Data BuyAvatarDataInfo
}
type BuyAvatarDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Idx         int16
}
type BuyAvatarRetDataInfo struct {
	Avatar_idx int16
	U_gold     int64
	U_cp       int64
	U_mp       int64
}
type BuyAvatarReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        BuyAvatarRetDataInfo
	Maintenance MaintenanceData
}
type BuyCheck struct {
	Call string
	Data BuyCheckDataInfo
}
type BuyCheckDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Idx         int32
}
type BuyCheckRetDataInfo struct {
	Result int64
}
type BuyCheckReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        BuyCheckRetDataInfo
	Maintenance MaintenanceData
}
type BuyCollectionItem struct {
	Call string
	Data BuyCollectionItemDataInfo
}
type BuyCollectionItemDataInfo struct {
	U_seq             int32
	U_id              string
	Uuid              string
	Device_uuid       string
	Special_music_idx int32
	Idx               int32
}
type BuyCollectionItemRetDataInfo struct {
	Item_type    int16
	Item_value   int64
	Buy_datetime string
}
type BuyCollectionItemReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        BuyCollectionItemRetDataInfo
	Maintenance MaintenanceData
}
type BuyCollectionMusic struct {
	Call string
	Data BuyCollectionMusicDataInfo
}
type BuyCollectionMusicDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Idx         int32
}
type BuyCollectionMusicRetDataInfo struct {
	Music_idx    int64
	Buy_datetime string
}
type BuyCollectionMusicReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        BuyCollectionMusicRetDataInfo
	Maintenance MaintenanceData
}
type BuyContents struct {
	Call        string
	Data        BuyContentsDataInfo
	Sub_mode    string
	Common_data ParamData
}
type BuyContentsDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Type        string
	Idx         int32
}
type BuyContentsRetDataInfo struct {
	U_cp       int64
	U_candy    float64
	User_skill UserSkill
	User_unit  UserUnit
}
type BuyContentsReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        BuyContentsRetDataInfo
	Maintenance MaintenanceData
}
type BuyExtraMode struct {
	Call string
	Data BuyExtraModeDataInfo
}
type BuyExtraModeDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Music_idx   int32
	Key_flg     int64
}
type BuyExtraModeRetDataInfo struct {
	Music_idx int32
}
type BuyExtraModeReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        BuyExtraModeRetDataInfo
	Maintenance MaintenanceData
}
type BuyItem struct {
	Call string
	Data BuyItemDataInfo
}
type BuyItemDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Idx         int16
}
type BuyItemRetDataInfo struct {
	Idx      int32
	U_gold   int64
	U_cp     int64
	U_mp     int64
	U_credit int64
}
type BuyItemReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        BuyItemRetDataInfo
	Maintenance MaintenanceData
}
type BuyMusic struct {
	Call string
	Data BuyMusicDataInfo
}
type BuyMusicDataInfo struct {
	Type          int16
	Idx           int32
	U_seq         int32
	U_id          string
	Uuid          string
	Device_uuid   string
	Purchase_type int16
}
type BuyMusicRetDataInfo struct {
	U_gold int32
	U_cp   int32
	U_mp   int32
	Music  map[any]any
}
type BuyMusicReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        BuyMusicRetDataInfo
	Maintenance MaintenanceData
}
type BuyPackage struct {
	Call string
	Data BuyPackageDataInfo
}
type BuyPackageDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Store_idx   int64
	Product_id  int64
}
type BuyPackageMusic struct {
	Music_idx int64
}
type BuyPackageRetDataInfo struct {
	U_gold     int64
	U_cp       int64
	U_mp       int64
	U_credit   int64
	Music_data []any
}
type BuyPackageReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        BuyPackageRetDataInfo
	Maintenance MaintenanceData
}
type BuyPiece struct {
	Call string
	Data BuyPieceDataInfo
}
type BuyPieceDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Idx         int32
	Cnt         int64
}
type BuyPieceRetDataInfo struct {
	Idx    int32
	U_gold int64
	U_cp   int32
	U_mp   int32
}
type BuyPieceReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        BuyPieceRetDataInfo
	Maintenance MaintenanceData
}
type BuyShop struct {
	Call        string
	Data        BuyShopDataInfo
	Sub_mode    string
	Common_data ParamData
}
type BuyShopDataInfo struct {
	U_seq             int32
	U_id              string
	Uuid              string
	Device_uuid       string
	Idx               int32
	Developer_payload string
	Pay_id            string
	Purchase_token    string
	Purchase_id       string
	Os                int16
	App_id            string
	Member_id         string
}
type BuyShopRetDataInfo struct {
	U_cp          int64
	U_candy       float64
	Reward_data   []any
	User_ad_level UserAdLevel
}
type BuyShopReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        BuyShopRetDataInfo
	Maintenance MaintenanceData
}
type BuyUserTitle struct {
	Call string
	Data BuyUserTitleDataInfo
}
type BuyUserTitleDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Idx         int16
}
type BuyUserTitleRetDataInfo struct {
	Title_idx int16
	U_gold    int64
	U_cp      int64
	U_mp      int64
}
type BuyUserTitleReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        BuyUserTitleRetDataInfo
	Maintenance MaintenanceData
}
type BuyVarietyStore struct {
	Call        string
	Data        BuyVarietyStoreDataInfo
	Common_data ParamData
}
type BuyVarietyStoreDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Idx         int32
}
type BuyVarietyStoreRetDataInfo struct {
	U_cp         int64
	U_candy      float64
	Reward_type  int16
	Reward_id    int32
	Reward_value int64
	Status       string
}
type BuyVarietyStoreReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        BuyVarietyStoreRetDataInfo
	Maintenance MaintenanceData
}
type ChThirdStage struct {
	Call        string
	Data        ChThirdStageDataInfo
	Common_data ParamData
}
type ChThirdStageDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	I_id        int32
	Music_id    int32
	Profile_ids string
}
type ChThirdStageRetDataInfo struct {
	Star                        int32
	Character_score             int32
	Music_score                 int32
	Follower_profile_score      int32
	Bonus_score                 int32
	Total_score                 int64
	Clear                       int32
	User_ap                     UserApData
	User_ch_third_stage         UserChThirdStage
	User_music                  UserMusic
	Reward_data                 []any
	Bonus_follower_profile_ids  []any
	User_follower_profile       UserFollowerProfile
	Bonus_music_score           int32
	User_follower_profile_score []any
}
type ChThirdStageReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        ChThirdStageRetDataInfo
	Maintenance MaintenanceData
}
type ChThirdStreamStage struct {
	Call        string
	Data        ChThirdStreamStageDataInfo
	Common_data ParamData
}
type ChThirdStreamStageDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	I_id        int32
	Count       int32
}
type ChThirdStreamStageRetDataInfo struct {
	I_id                  int32
	User_ap               UserApData
	User_follower_profile UserFollowerProfile
	Reward_data           map[any]any
}
type ChThirdStreamStageReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        ChThirdStreamStageRetDataInfo
	Maintenance MaintenanceData
}
type ChallengeMissionList struct {
	Call string
	Data ChallengeMissionListDataInfo
}
type ChallengeMissionListDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Music_idx   int32
	Difficulty  int16
	Up          int32
	Down        int32
}
type ChallengeMissionListInfo struct {
	Music_idx    int32
	Difficulty   int16
	Mission_list []any
}
type ChallengeMissionListRetDataInfo struct {
	Music_idx    int32
	Difficulty   int16
	Mission_list map[any]any
}
type ChallengeMissionListReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        ChallengeMissionListRetDataInfo
	Maintenance MaintenanceData
}
type ChangeUser struct {
	Call        string
	Data        ChangeUserDataInfo
	Common_data ParamData
}
type ChangeUserDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Change_uuid string
}
type ChangeUserRetDataInfo struct {
	Result string
}
type ChangeUserReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        ChangeUserRetDataInfo
	Maintenance MaintenanceData
}
type ChannelList struct {
	Channel        int16
	Url            string
	Port           int64
	Flg            int16
	Maintenance_cn string
	Maintenance_en string
	Maintenance_jp string
}
type CheckBuyProduct struct {
	Call string
	Data CheckBuyProductDataInfo
}
type CheckBuyProductDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Product_id  string
}
type CheckBuyProductRetDataInfo struct {
	Is_owner int16
}
type CheckBuyProductReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        CheckBuyProductRetDataInfo
	Maintenance MaintenanceData
}
type CheckBuyShop struct {
	Call        string
	Data        CheckBuyShopDataInfo
	Common_data ParamData
}
type CheckBuyShopDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Shop_id     string
}
type CheckBuyShopRetDataInfo struct {
	Is_owner int16
}
type CheckBuyShopReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        CheckBuyShopRetDataInfo
	Maintenance MaintenanceData
}
type CheckPurchased struct {
	Call        string
	Data        CheckPurchasedDataInfo
	Common_data ParamData
}
type CheckPurchasedDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
}
type CheckPurchasedRetDataInfo struct {
	Is_purchased int16
}
type CheckPurchasedReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        CheckPurchasedRetDataInfo
	Maintenance MaintenanceData
}
type CheckUser struct {
	Call        string
	Data        CheckUserDataInfo
	Common_data ParamData
}
type CheckUserDataInfo struct {
	Uuid        string
	Device_uuid string
}
type CheckUserRetDataInfo struct {
	U_seq int32
	U_id  string
}
type CheckUserReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        CheckUserRetDataInfo
	Maintenance MaintenanceData
}
type ChoiceUserData struct {
	U_girl_level int64
	U_fans       int64
	U_fans_grade int16
	U_cp         int64
	U_candy      float64
	U_like       float64
	U_last_login string
}
type CollectionRetData struct {
	Music_idx int64
	Cnt       int64
}
type CompleteLog struct {
	Call        string
	Data        CompleteLogDataInfo
	Common_data ParamData
}
type CompleteLogDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Type        string
	Id          int32
	Value       string
}
type CompleteLogRetDataInfo struct {
	Status string
}
type CompleteLogReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        CompleteLogRetDataInfo
	Maintenance MaintenanceData
}
type DefaultSettingDataList struct {
	Setting_key   string
	Setting_value string
}
type DefaultSettingList struct {
	Call        string
	Data        DefaultSettingListDataInfo
	Common_data ParamData
}
type DefaultSettingListDataInfo struct {
	Device_uuid string
	Os          int16
}
type DefaultSettingListRetDataInfo struct {
	Status       string
	Setting_list []any
}
type DefaultSettingListReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        DefaultSettingListRetDataInfo
	Maintenance MaintenanceData
}
type DeletePost struct {
	Call        string
	Data        DeletePostDataInfo
	Common_data ParamData
}
type DeletePostDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Idx         int64
	Type        int16
}
type DeletePostRetDataInfo struct {
	Idx int64
}
type DeletePostReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        DeletePostRetDataInfo
	Maintenance MaintenanceData
}
type DeleteUser struct {
	Call string
	Data DeleteUserDataInfo
}
type DeleteUserDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
}
type DeleteUserRetDataInfo struct {
	Result string
}
type DeleteUserReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        DeleteUserRetDataInfo
	Maintenance MaintenanceData
}
type EndPlayMusic struct {
	Call string
	Data EndPlayMusicDataInfo
}
type EndPlayMusicDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Score       int32
	Medal       int16
	Over_flg    int16
}
type EndPlayMusicRetDataInfo struct {
	Code int16
}
type EndPlayMusicReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        EndPlayMusicRetDataInfo
	Maintenance MaintenanceData
}
type ErrorRetCode struct {
	Code   int32
	Errmsg string
}
type EventRankRewardListArr struct {
	Event_idx    int32
	Title        string
	Desc_kr      string
	Desc_en      string
	Desc_jp      string
	Reward_type  int16
	Reward_value int64
}
type EventRankRewardListData struct {
	Title        string
	Desc_cn      string
	Desc_en      string
	Desc_jp      string
	Reward_type  int16
	Reward_value int64
}
type EventRewardList struct {
	Call string
	Data EventRewardListDataInfo
}
type EventRewardListArr struct {
	Event_idx    int32
	Gain_point   int64
	Reward_type  int16
	Reward_value int64
}
type EventRewardListData struct {
	Gain_point   int64
	Reward_type  int16
	Reward_value int64
}
type EventRewardListDataInfo struct {
	Event_idx   int32
	Device_uuid string
}
type EventRewardListRetDataInfo struct {
	Event_idx        int32
	Reward_list      []any
	Rank_reward_list []any
}
type EventRewardListReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        EventRewardListRetDataInfo
	Maintenance MaintenanceData
}
type FacebookUserJoin struct {
	Call        string
	Data        FacebookUserJoinDataInfo
	Common_data ParamData
}
type FacebookUserJoinDataInfo struct {
	Access_token string
}
type GetAchievementData struct {
	I_id                       int32
	S_Title_KO                 string
	S_Title_EN                 string
	S_Title_JA                 string
	S_Title_ZH_CHS             string
	S_Title_ZH_CHT             string
	S_Title_VI                 string
	S_Title_ES                 string
	S_Title_IT                 string
	S_Title_ID                 string
	S_Title_TH                 string
	S_Title_PT                 string
	S_Title_HI                 string
	S_Description_KO           string
	S_Description_EN           string
	S_Description_JA           string
	S_Description_ZH_CHS       string
	S_Description_ZH_CHT       string
	S_Description_VI           string
	S_Description_ES           string
	S_Description_IT           string
	S_Description_ID           string
	S_Description_TH           string
	S_Description_PT           string
	S_Description_HI           string
	S_ResourceName             string
	S_AchievementConditionType string
	S_RewardType               string
	S_RewardDescription_KO     string
	S_RewardDescription_EN     string
	S_RewardDescription_JA     string
	S_RewardDescription_ZH_CHS string
	S_RewardDescription_ZH_CHT string
	S_RewardDescription_VI     string
	S_RewardDescription_ES     string
	S_RewardDescription_IT     string
	S_RewardDescription_ID     string
	S_RewardDescription_TH     string
	S_RewardDescription_PT     string
	S_RewardDescription_HI     string
	I_Reward_1                 int64
	I_Reward_2                 int64
	I_Reward_3                 int64
	I_Reward_4                 int64
	I_Reward_5                 int64
	I_Reward_6                 int64
	I_Reward_7                 int64
	I_Reward_8                 int64
	I_Reward_9                 int64
	I_Reward_10                int64
	I_Area                     int16
	B_IsActive                 int16
	I_Reward_11                int64
	I_Reward_12                int64
	I_MaxLevel                 int16
	S_Condition_1              string
	S_Condition_2              string
	S_Condition_3              string
	S_Condition_4              string
	S_Condition_5              string
	S_Condition_6              string
	S_Condition_7              string
	S_Condition_8              string
	S_Condition_9              string
	S_Condition_10             string
	S_Condition_11             string
	S_Condition_12             string
	S_Condition_13             string
	S_Condition_14             string
	S_Condition_15             string
	S_Condition_16             string
	S_Condition_17             string
	S_Condition_18             string
	S_Condition_19             string
	S_Condition_20             string
	I_Reward_13                int64
	I_Reward_14                int64
	I_Reward_15                int64
	I_Reward_16                int64
	I_Reward_17                int64
	I_Reward_18                int64
	I_Reward_19                int64
	I_Reward_20                int64
}
type GetAdLevelData struct {
	I_GroupID     int32
	I_Level       int32
	I_RequireEXP  int32
	I_RewardGroup int32
	B_IsActive    int16
	I_id          int32
}
type GetAdListData struct {
	I_id              int32
	I_Category        int32
	S_ConditionType1  string
	S_ConditionValue1 string
	S_ConditionType2  string
	S_ConditionValue2 string
	S_ResourceName    string
	S_Name_KO         string
	S_Name_EN         string
	S_Name_JA         string
	S_Name_ZH_CHS     string
	S_Name_ZH_CHT     string
	S_Name_VI         string
	S_Name_ES         string
	S_Name_IT         string
	S_Name_ID         string
	S_Name_TH         string
	S_Name_PT         string
	S_Name_HI         string
	B_IsActive        int16
	S_RewardType      string
	I_RewardGroup     int32
	S_ExtraValue      string
	S_AdType          string
}
type GetAlbum struct {
	Call string
	Data GetAlbumDataInfo
}
type GetAlbumDataInfo struct {
	Null string
}
type GetAlbumInfoLanguage struct {
	Seq          int32
	Album_title  string
	Company      string
	Album_artist string
}
type GetAlbumInfoList struct {
	Call string
	Data GetAlbumListDataInfo
}
type GetAlbumInfoListDataInfo struct {
}
type GetAlbumInfoListRetDataInfo struct {
	Language map[any]any
}
type GetAlbumInfoListReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        map[any]any
	Maintenance MaintenanceData
}
type GetAlbumList struct {
	Call string
	Data GetAlbumListDataInfo
}
type GetAlbumListDataInfo struct {
}
type GetAlbumListRetDataInfo struct {
	Album_idx  int32
	Country    string
	Date_issue string
	Cdn_dir    string
	Album_img  string
}
type GetAlbumListReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        []any
	Maintenance MaintenanceData
}
type GetAlbumRetDataInfo struct {
	Album     AlbumListInfo
	List_info []any
}
type GetAlbumReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        []any
	Maintenance MaintenanceData
}
type GetAllPoint struct {
	Call     string
	Data     GetAllPointDataInfo
	Sub_mode string
}
type GetAllPointDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
}
type GetAllPointRetDataInfo struct {
	U_gold           int64
	U_cp             int64
	U_mp             int64
	U_credit         int64
	Next_credit_time int64
	Max_credit_time  int64
}
type GetAllPointReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        GetAllPointRetDataInfo
	Maintenance MaintenanceData
}
type GetBuffData struct {
	I_id               int32
	S_ResourceName     string
	S_Title_KO         string
	S_Title_EN         string
	S_Title_JA         string
	S_Title_ZH_CHS     string
	S_Title_ZH_CHT     string
	S_Title_VI         string
	S_Title_ES         string
	S_Title_IT         string
	S_Title_ID         string
	S_Title_TH         string
	S_Title_PT         string
	S_Title_HI         string
	S_Buff_Desc_KO     string
	S_Buff_Desc_EN     string
	S_Buff_Desc_JA     string
	S_Buff_Desc_ZH_CHS string
	S_Buff_Desc_ZH_CHT string
	S_Buff_Desc_VI     string
	S_Buff_Desc_ES     string
	S_Buff_Desc_IT     string
	S_Buff_Desc_ID     string
	S_Buff_Desc_TH     string
	S_Buff_Desc_PT     string
	S_Buff_Desc_HI     string
	I_Type             int16
	I_Retention_Time   int64
}
type GetCache struct {
	Call string
	Data GetCacheDataInfo
}
type GetCacheDataInfo struct {
	Type int16
}
type GetCacheRetDataInfo struct {
	Uptime int32
}
type GetCacheReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        GetCacheRetDataInfo
	Maintenance MaintenanceData
}
type GetChThird struct {
	Call        string
	Data        GetChThirdDataInfo
	Common_data ParamData
}
type GetChThirdChapterData struct {
	I_id             int32
	I_unlockLevel    int32
	I_RewardGroupID1 int32
	I_GoalStar1      int32
	I_RewardGroupID2 int32
	I_GoalStar2      int32
	I_RewardGroupID3 int32
	I_GoalStar3      int32
	S_Name_KO        string
	S_Name_EN        string
	S_Name_JA        string
	S_Name_ZH_CHS    string
	S_Name_ZH_CHT    string
	S_Name_VI        string
	S_Name_ES        string
	S_Name_IT        string
	S_Name_ID        string
	S_Name_TH        string
	S_Name_PT        string
	S_Name_HI        string
	B_IsActive       int16
}
type GetChThirdDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
}
type GetChThirdRetDataInfo struct {
	U_seq                        int32
	User_ap                      UserApData
	User_ch_third_stage          []any
	User_ch_third_chapter_reward []any
}
type GetChThirdReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        GetChThirdRetDataInfo
	Maintenance MaintenanceData
}
type GetChThirdScoreData struct {
	I_id             int32
	I_Level          int32
	I_CharacterScore int32
	I_FollowerScore  int32
	I_MusicScore     int32
	B_IsActive       int16
}
type GetChThirdStageData struct {
	I_id                   int32
	I_Chapter              int32
	I_StageIndex           int32
	I_RewardGroupID        int32
	I_GoalScore1           int32
	I_GoalScore2           int32
	I_GoalScore3           int32
	S_BonusProfileID       string
	S_Name_KO              string
	S_Name_EN              string
	S_Name_JA              string
	S_Name_ZH_CHS          string
	S_Name_ZH_CHT          string
	S_Name_VI              string
	S_Name_ES              string
	S_Name_IT              string
	S_Name_ID              string
	S_Name_TH              string
	S_Name_PT              string
	S_Name_HI              string
	S_Description_KO       string
	S_Description_EN       string
	S_Description_JA       string
	S_Description_ZH_CHS   string
	S_Description_ZH_CHT   string
	S_Description_VI       string
	S_Description_ES       string
	S_Description_IT       string
	S_Description_ID       string
	S_Description_TH       string
	S_Description_PT       string
	S_Description_HI       string
	B_IsActive             int16
	S_StageType            string
	I_StoryID              int32
	I_BonusMusicID         int32
	S_PercentRewardGroupID string
}
type GetChThirdStarReward struct {
	Call        string
	Data        GetChThirdStarRewardDataInfo
	Common_data ParamData
}
type GetChThirdStarRewardDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	I_id        int32
	Reward_num  int32
}
type GetChThirdStarRewardRetDataInfo struct {
	I_id        int32
	Reward_num  int32
	User_ap     UserApData
	Reward_data []any
}
type GetChThirdStarRewardReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        GetChThirdStarRewardRetDataInfo
	Maintenance MaintenanceData
}
type GetCharacterData struct {
	I_id                      int32
	I_Area                    int16
	S_BundleName              string
	S_Name_KO                 string
	S_Name_EN                 string
	S_Name_JA                 string
	S_Name_ZH_CHS             string
	S_Name_ZH_CHT             string
	S_Name_VI                 string
	S_Name_ES                 string
	S_Name_IT                 string
	S_Name_ID                 string
	S_Name_TH                 string
	S_Name_PT                 string
	S_Name_HI                 string
	S_Description_KO          string
	S_Description_EN          string
	S_Description_JA          string
	S_Description_ZH_CHS      string
	S_Description_ZH_CHT      string
	S_Description_VI          string
	S_Description_ES          string
	S_Description_IT          string
	S_Description_ID          string
	S_Description_TH          string
	S_Description_PT          string
	S_Description_HI          string
	S_ResourceName            string
	I_MaxLevel                int64
	D_LevelRate               float64
	D_LateCost                float64
	D_Cost                    float64
	D_CostIncreaseValue       float64
	D_CostValueConstant       float64
	D_Amount                  float64
	D_AmountIncreaseValue     float64
	D_AmountValueConstant     float64
	D_BonusAmountIncreaseRate float64
	I_UnlockLevel             int64
	D_UnlockCost              float64
	I_RequirementFollowerId_1 int64
	I_RequirementFollowerId_2 int64
	I_RequirementPropId_1     int64
	I_RequirementPropId_2     int64
	I_RequirementPropCount    int64
	B_IsActive                int16
}
type GetChoiceUser struct {
	Call        string
	Data        GetChoiceUserDataInfo
	Common_data ParamData
}
type GetChoiceUserDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Choice_uuid string
}
type GetChoiceUserRetDataInfo struct {
	User        ChoiceUserData
	Choice_user ChoiceUserData
}
type GetChoiceUserReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        GetChoiceUserRetDataInfo
	Maintenance MaintenanceData
}
type GetCollection struct {
	Call string
	Data GetCollectionDataInfo
}
type GetCollectionDataInfo struct {
}
type GetCollectionItem struct {
	Special_music_idx int64
	Idx               int64
	Item_type         int16
	Item_value        int64
	Buy_cnt           int64
}
type GetCollectionRetDataInfo struct {
	Idx      int32
	Buy_type int16
	Price    int64
	Item     []any
}
type GetCollectionReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        []any
	Maintenance MaintenanceData
}
type GetConsumeData struct {
	I_id           int32
	S_ResourceName string
	S_Title_KO     string
	S_Title_EN     string
	S_Title_JA     string
	S_Title_ZH_CHS string
	S_Title_ZH_CHT string
	S_Title_VI     string
	S_Title_ES     string
	S_Title_IT     string
	S_Title_ID     string
	S_Title_TH     string
	S_Title_PT     string
	S_Title_HI     string
	S_Limit        string
	I_Type         int16
	S_ExtraValue   string
	I_Area         int16
	B_IsActive     int16
}
type GetCostumeData struct {
	I_id                      int32
	S_Name_KO                 string
	S_Name_EN                 string
	S_Name_JA                 string
	S_Name_ZH_CHS             string
	S_Name_ZH_CHT             string
	S_Name_VI                 string
	S_Name_ES                 string
	S_Name_IT                 string
	S_Name_ID                 string
	S_Name_TH                 string
	S_Name_PT                 string
	S_Name_HI                 string
	S_Description_KO          string
	S_Description_EN          string
	S_Description_JA          string
	S_Description_ZH_CHS      string
	S_Description_ZH_CHT      string
	S_Description_VI          string
	S_Description_ES          string
	S_Description_IT          string
	S_Description_ID          string
	S_Description_TH          string
	S_Description_PT          string
	S_Description_HI          string
	S_ResourceName            string
	S_CostumeName             string
	I_MaxLevel                int64
	S_GoodsType               string
	S_Cost                    string
	F_CostIncreaseRate        float64
	S_Operation               string
	D_Amount                  int64
	F_AmountIncreaseRate      float64
	S_BonusOperation          string
	D_BonusAmount             int64
	F_BonusAmountIncreaseRate float64
	I_UnlockLevel             int64
	I_RequirementFollowerId_1 int64
	I_RequirementFollowerId_2 int64
	I_RequirementPropId_1     int64
	I_RequirementPropId_2     int64
	I_RequirementPropCount    int64
	I_SortOrder               int64
	F_FanMultiply             float64
	I_AcquisitionType         int64
	I_AcquisitionId           int64
	B_IsActive                int16
	I_Area                    int16
	I_DownloadType            int16
	S_AltCostumeName          string
	S_AltCostumeStartTime     string
	S_AltCostumeEndTime       string
}
type GetCountryMyRank struct {
	Call string
	Data GetCountryMyRankDataInfo
}
type GetCountryMyRankDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Country     string
	Start_limit int32
	End_limit   int32
}
type GetCountryMyRankRet struct {
	G           int16
	S           int16
	B           int16
	Total_score int32
	U_avatar    int16
	U_id        string
	U_nick      string
	U_seq       int32
	U_country   string
	Rank        int32
}
type GetCountryMyRankRetDataInfo struct {
	Select_country string
	Start_limit    int16
	End_limit      int16
	My_rank        GetCountryMyRankRet
	Rank_list      []any
}
type GetCountryMyRankReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        GetCountryMyRankRetDataInfo
	Maintenance MaintenanceData
}
type GetCountryRank struct {
	Call string
	Data GetCountryRankDataInfo
}
type GetCountryRankDataInfo struct {
}
type GetCountryRankRetDataInfo struct {
	Rank        int32
	Country     string
	Gold        int32
	Silver      int32
	Bronze      int32
	Total_score int32
}
type GetCountryRankReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        []any
	Maintenance MaintenanceData
}
type GetCountryUserRank struct {
	G           int16
	S           int16
	B           int16
	Total_score int32
	U_avatar    int16
	U_nick      string
	U_seq       int32
}
type GetCountryUserRankRet struct {
	G           int16
	S           int16
	B           int16
	Total_score int32
	U_avatar    int16
	U_id        string
	U_nick      string
	U_seq       int32
}
type GetDailyMissionData struct {
	I_id                        int32
	S_Title_KO                  string
	S_Title_EN                  string
	S_Title_JA                  string
	S_Title_ZH_CHS              string
	S_Title_ZH_CHT              string
	S_Title_VI                  string
	S_Title_ES                  string
	S_Title_IT                  string
	S_Title_ID                  string
	S_Title_TH                  string
	S_Title_PT                  string
	S_Title_HI                  string
	S_Description_KO            string
	S_Description_EN            string
	S_Description_JA            string
	S_Description_ZH_CHS        string
	S_Description_ZH_CHT        string
	S_Description_VI            string
	S_Description_ES            string
	S_Description_IT            string
	S_Description_ID            string
	S_Description_TH            string
	S_Description_PT            string
	S_Description_HI            string
	S_ResourceName              string
	S_DailyMissionConditionType string
	D_Condition_1               float64
	S_RewardType                string
	S_RewardDescription_KO      string
	S_RewardDescription_EN      string
	S_RewardDescription_JA      string
	S_RewardDescription_ZH_CHS  string
	S_RewardDescription_ZH_CHT  string
	S_RewardDescription_VI      string
	S_RewardDescription_ES      string
	S_RewardDescription_IT      string
	S_RewardDescription_ID      string
	S_RewardDescription_TH      string
	S_RewardDescription_PT      string
	S_RewardDescription_HI      string
	I_Reward_1                  int32
	I_Area                      int16
	B_IsActive                  int16
}
type GetDefaultSettingRetDataInfo struct {
	Ad_count          int32
	Ad_probability    float64
	Challenge_mission float64
	Seed_value        int16
	Rank_up_limit     int16
	Rank_down_limit   int16
	Reward_rate       float64
	Tour_btn_name     string
	Event_btn_name    string
	Event_rank_url    string
	Tour_intro_movie  string
	Tour_bg_movie     string
	Event_bg_movie    string
}
type GetDiamondBonus struct {
	Call string
	Data GetDiamondBonusDataInfo
}
type GetDiamondBonusDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
}
type GetDiamondBonusRetDataInfo struct {
	Bonus1 int64
	Bonus2 int64
	Bonus3 int64
	Bonus4 int64
	Bonus5 int64
	Bonus6 int64
}
type GetDiamondBonusReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        GetDiamondBonusRetDataInfo
	Maintenance MaintenanceData
}
type GetEventListData struct {
	Idx        int32
	Event_type string
	Start_date string
	End_date   string
	B_IsActive int16
}
type GetEventMain struct {
	Call string
	Data GetEventMainDataInfo
}
type GetEventMainDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
}
type GetEventMainRetDataInfo struct {
	Idx         int32
	Title       string
	Img1        string
	Img2        string
	Movie       string
	Utc         int64
	Banner      string
	Desc_cn     string
	Desc_en     string
	Desc_jp     string
	Server_time int64
	Start_time  int64
	End_time    int64
	My_score    int64
	My_rank     int64
}
type GetEventMainReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        GetEventMainRetDataInfo
	Maintenance MaintenanceData
}
type GetEventRank struct {
	Call string
	Data GetEventRankDataInfo
}
type GetEventRankDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Event_idx   int32
	Start_rank  int64
	End_rank    int64
}
type GetEventRankRetDataInfo struct {
	U_seq     int64
	U_nick    string
	U_id      string
	U_avatar  int32
	U_country string
	U_title   int32
	Rank      int64
	Score     int64
}
type GetEventRankReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        []any
	Maintenance MaintenanceData
}
type GetEventRewardList struct {
	Call        string
	Data        GetEventRewardListDataInfo
	Sub_mode    string
	Common_data ParamData
}
type GetEventRewardListData struct {
	Idx                int64
	Event_name         string
	Event_type         string
	Reward_idx         int64
	Reward_num         int32
	Reward_type        int32
	Reward_id          int32
	Reward_value       int32
	Reward_flg         string
	Get_date           int32
	S_CustomIconType   string
	S_CustomIconSprite string
}
type GetEventRewardListDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Event_idx   int32
}
type GetEventRewardListRetDataInfo struct {
	Reward_list []any
	Group_idx   int32
}
type GetEventRewardListReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        map[any]any
	Maintenance MaintenanceData
}
type GetFanData struct {
	I_Area        int64
	I_Grade       int64
	I_FanCount    int64
	I_BonusRate   int64
	I_BonusRateUI int64
}
type GetFollowerData struct {
	I_id                      int32
	S_Name_KO                 string
	S_Name_EN                 string
	S_Name_JA                 string
	S_Name_ZH_CHS             string
	S_Name_ZH_CHT             string
	S_Name_VI                 string
	S_Name_ES                 string
	S_Name_IT                 string
	S_Name_ID                 string
	S_Name_TH                 string
	S_Name_PT                 string
	S_Name_HI                 string
	S_Description_KO          string
	S_Description_EN          string
	S_Description_JA          string
	S_Description_ZH_CHS      string
	S_Description_ZH_CHT      string
	S_Description_VI          string
	S_Description_ES          string
	S_Description_IT          string
	S_Description_ID          string
	S_Description_TH          string
	S_Description_PT          string
	S_Description_HI          string
	S_ResourceName            string
	I_MaxLevel                int64
	D_LevelRate               float64
	D_LateCost                int64
	D_Cost                    float64
	D_CostIncreaseValue       float64
	D_CostValueConstant       int64
	D_Amount                  float64
	D_AmountIncreaseValue     float64
	D_AmountValueConstant     int64
	D_BonusAmountIncreaseRate int64
	I_UnlockLevel             int64
	D_UnlockCost              int64
	I_RequirementFollowerId_1 int64
	I_RequirementFollowerId_2 int64
	I_RequirementPropId_1     int64
	I_RequirementPropId_2     int64
	I_RequirementPropCount    int64
	B_IsActive                int16
	I_Area                    int16
	I_DownloadType            int32
	I_ProfileID               int32
}
type GetFollowerGiftItemData struct {
	I_id           int32
	I_GiftType     int32
	D_Value        int64
	I_Limit        int32
	S_ResourceName string
	S_Name_KO      string
	S_Name_EN      string
	S_Name_JA      string
	S_Name_ZH_CHS  string
	S_Name_ZH_CHT  string
	S_Name_VI      string
	S_Name_ES      string
	S_Name_IT      string
	S_Name_ID      string
	S_Name_TH      string
	S_Name_PT      string
	S_Name_HI      string
	B_IsActive     int16
}
type GetFollowerProfileData struct {
	I_id              int32
	S_Name_KO         string
	S_Name_EN         string
	S_Name_JA         string
	S_Name_ZH_CHS     string
	S_Name_ZH_CHT     string
	S_Name_VI         string
	S_Name_ES         string
	S_Name_IT         string
	S_Name_ID         string
	S_Name_TH         string
	S_Name_PT         string
	S_Name_HI         string
	B_IsActive        int16
	S_GiftItemID      string
	B_IsVisible       int16
	S_BackgroundColor string
}
type GetFollowerProfileLevelData struct {
	I_id                int32
	I_ProfileID         int32
	I_Level             int32
	D_RequireEXP        int64
	I_RewardGroup       int32
	S_UnlockInformation string
	I_AddCandy          int32
	I_UnlockStoryID     int32
	B_IsActive          int16
}
type GetFollowerQuestData struct {
	I_id                  int32
	I_UnlockFollowerID    int32
	I_UnlockFollowerLevel int32
	S_ConditionType1      string
	D_Condition1_1        float64
	D_Condition1_2        float64
	S_ConditionType2      string
	D_Condition2_1        float64
	D_Condition2_2        float64
	S_ConditionType3      string
	D_Condition3_1        float64
	D_Condition3_2        float64
	I_RewardGroup1        int32
	I_RewardGroup2        int32
	I_RewardGroup3        int32
	S_Description_KO      string
	S_Description_EN      string
	S_Description_ZH_CHS  string
	S_Description_ZH_CHT  string
	S_Description_JA      string
	S_Description_VI      string
	S_Description_ES      string
	S_Description_IT      string
	S_Description_ID      string
	S_Description_TH      string
	S_Description_PT      string
	S_Description_HI      string
	B_IsActive            int16
	I_Type                int32
}
type GetGameDataList struct {
	Call        string
	Data        GetGameDataListDataInfo
	Sub_mode    string
	Common_data ParamData
}
type GetGameDataListDataInfo struct {
	Game_type   string
	Device_uuid string
}
type GetGameDataListRetDataInfo struct {
	Consume                        []any
	Music                          []any
	Costume                        []any
	Prop                           []any
	Follower                       []any
	Buff                           []any
	Unit                           []any
	Skill                          []any
	Achievement                    []any
	Daily_mission                  []any
	Music_level                    []any
	Shop                           []any
	Reward_group                   []any
	Fan                            []any
	SystemNotification             []any
	SystemString                   []any
	SubscribeList                  []any
	SubscribePassReward            []any
	SubscribePass                  []any
	Guitar                         []any
	SubscribePassRewardInformation []any
	Ticketcollection               []any
	Character                      []any
	Samseckevent                   []any
	Localpush                      []any
	Followerquest                  []any
	Proplevel                      []any
	Follower_profile               []any
	Passgoodsshop                  []any
	Followergiftitem               []any
	Followerprofilelevel           []any
	Ad_list                        []any
	Chthird_stage                  []any
	Chthird_score                  []any
	Chthird_chapter                []any
	Percent                        []any
	Ad_level                       []any
	Event_list                     []any
	Select_reward                  []any
}
type GetGameDataListReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        map[any]any
	Maintenance MaintenanceData
}
type GetGuitarData struct {
	I_id                        int32
	I_Area                      int16
	S_Title_KO                  string
	S_Title_EN                  string
	S_Title_JA                  string
	S_Title_ZH_CHS              string
	S_Title_ZH_CHT              string
	S_Title_VI                  string
	S_Title_ES                  string
	S_Title_IT                  string
	S_Title_ID                  string
	S_Title_TH                  string
	S_Title_PT                  string
	S_Title_HI                  string
	S_Description_KO            string
	S_Description_EN            string
	S_Description_JA            string
	S_Description_ZH_CHS        string
	S_Description_ZH_CHT        string
	S_Description_VI            string
	S_Description_ES            string
	S_Description_IT            string
	S_Description_ID            string
	S_Description_TH            string
	S_Description_PT            string
	S_Description_HI            string
	S_ResourceName              string
	S_GuitarName                string
	I_GuitarType                int16
	S_GoodsType                 string
	S_Cost                      string
	I_RequirementCharacterLevel int64
	D_IncrementValue            float64
	I_AcquisitionType           int64
	I_AcquisitionId             int64
	I_SortOrder                 int64
	B_IsActive                  int16
	I_DownloadType              int16
}
type GetLocalPush struct {
	Call string
	Data GetLocalPushDataInfo
}
type GetLocalPushData struct {
	I_id                    int32
	S_LocalPushRegisterType string
	I_TimeForSeconds        int64
	B_IsMultipleJson        int16
	S_JsonMessage_KO        string
	S_JsonMessage_EN        string
	S_JsonMessage_JA        string
	S_JsonMessage_ZH_CHS    string
	S_JsonMessage_ZH_CHT    string
	S_JsonMessage_VI        string
	S_JsonMessage_ES        string
	S_JsonMessage_IT        string
	S_JsonMessage_ID        string
	S_JsonMessage_TH        string
	S_JsonMessage_PT        string
	S_JsonMessage_HI        string
	B_IsActive              int16
}
type GetLocalPushDataInfo struct {
	Device_uuid string
}
type GetLocalPushRetDataInfo struct {
	I_LocalPushID           int32
	S_LocalPushRegisterType int16
	I_TimeForSeconds        int64
	S_JsonMessage_ko        string
	S_JsonMessage_en        string
	S_JsonMessage_ja        string
}
type GetLocalPushReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        []any
	Maintenance MaintenanceData
}
type GetMastersRank struct {
	Call string
	Data GetMastersRankDataInfo
}
type GetMastersRankDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Start_limit int32
	End_limit   int32
}
type GetMastersRankRet struct {
	All_per   int16
	All_com   int16
	Score     int32
	U_avatar  int16
	U_id      string
	U_nick    string
	U_country string
	Rank      int32
}
type GetMastersRankRetDataInfo struct {
	Start_limit int16
	End_limit   int16
	My_rank     GetMastersRankRet
	Rank_list   []any
}
type GetMastersRankReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        GetMastersRankRetDataInfo
	Maintenance MaintenanceData
}
type GetMusicData struct {
	I_id                      int32
	S_Title_KO                string
	S_Title_EN                string
	S_Title_JA                string
	S_Title_ZH_CHS            string
	S_Title_ZH_CHT            string
	S_Title_VI                string
	S_Title_ES                string
	S_Title_IT                string
	S_Title_ID                string
	S_Title_TH                string
	S_Title_PT                string
	S_Title_HI                string
	S_Description_1_KO        string
	S_Description_1_EN        string
	S_Description_1_JA        string
	S_Description_1_ZH_CHS    string
	S_Description_1_ZH_CHT    string
	S_Description_1_VI        string
	S_Description_1_ES        string
	S_Description_1_IT        string
	S_Description_1_ID        string
	S_Description_1_TH        string
	S_Description_1_PT        string
	S_Description_1_HI        string
	S_ResourceName            string
	S_MusicFileName           string
	I_MaxLevel                int64
	D_LevelRate               float64
	D_LateCost                int64
	D_Cost                    float64
	D_CostIncreaseValue       float64
	D_CostValueConstant       int64
	D_Amount                  float64
	D_AmountIncreaseValue     float64
	D_AmountValueConstant     int64
	D_BonusAmountIncreaseRate int64
	I_UnlockLevel             int64
	D_UnlockCost              int64
	I_RequirementFollowerId_1 int64
	I_RequirementFollowerId_2 int64
	I_RequirementMusicId_1    int64
	I_RequirementMusicId_2    int64
	I_RequirementMusicCount   int64
	S_GoodsType               string
	I_SortOrder               int64
	I_AcquisitionType         int64
	I_AcquisitionId           int64
	B_IsActive                int16
	I_Area                    int16
	I_DownloadType_img        int32
	I_DownloadType_file       int32
}
type GetMusicInfoLanguage struct {
	Seq         int32
	Music_idx   int64
	Music_title string
	Artist      string
	Songwriter  string
	Lyricist    string
	Description string
}
type GetMusicInfoList struct {
	Call string
	Data GetMusicInfoListDataInfo
}
type GetMusicInfoListDataInfo struct {
	Country_code string
}
type GetMusicInfoListRetDataInfo struct {
	Language map[any]any
}
type GetMusicInfoListReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        map[any]any
	Maintenance MaintenanceData
}
type GetMusicLevelData struct {
	I_Level                          int16
	F_EncoreBonusAppearRate          float64
	I_EncoreBonusAppearPercent       int64
	F_EncoreBonusAppearCooltime_Sec  float64
	F_EncoreBonusAppearCooltime_Hour float64
	I_EncoreBonusGiftAmount          int16
	I_EncoreFollowerProfileExp       int32
	I_ChThirdCoolTime                int32
}
type GetMusicList struct {
	Call string
	Data GetMusicListDataInfo
}
type GetMusicListDataInfo struct {
	Country_code string
}
type GetMusicListRetDataInfo struct {
	Music_idx              int32
	Album_idx              int32
	Genre_idx              int32
	Difficulty1            int16
	Difficulty2            int16
	Difficulty3            int16
	Difficulty4            int16
	Purchaselink_type      int16
	Credit_cnt             int16
	Buy_type               int16
	Free_type              int16
	Price                  float64
	Is_extra               int16
	Extra_buy_type1        int16
	Extra_price1           int64
	Extra_buy_type2        int16
	Extra_price2           int64
	Buy_ad                 int32
	Music_ver              int32
	Bpm                    float64
	Cdn_dir                string
	Track_type             int16
	File_type              int16
	New_status             int16
	Update_time            int32
	Point_average          float64
	Use_movie              int16
	Movie_timing           int32
	Music_product_aos      MusicProduct
	Music_product_ios      MusicProduct
	Music_product_aos_pack MusicProduct
	Music_product_ios_pack MusicProduct
	Bga_path               string
	Bg_opacity             int16
}
type GetMusicListReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        []any
	Maintenance MaintenanceData
}
type GetMusicRank struct {
	Call string
	Data GetMusicRankDataInfo
}
type GetMusicRankDataInfo struct {
	Music_idx int32
}
type GetMusicRankRetDataInfo struct {
	Id    string
	Score int32
}
type GetMusicRankReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        GetMusicRankRetDataInfo
	Maintenance MaintenanceData
}
type GetMusicReward struct {
	Call        string
	Data        GetMusicRewardDataInfo
	Common_data ParamData
}
type GetMusicRewardDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	I_ids       []any
	I_levels    []any
}
type GetMusicRewardRetDataInfo struct {
	Total_reward_value    int32
	Reward_music_id       []any
	Reward_value          []any
	User_follower_profile []any
	Error_data            map[any]any
}
type GetMusicRewardReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        GetMusicRewardRetDataInfo
	Maintenance MaintenanceData
}
type GetNoticeList struct {
	Call string
	Data GetNoticeListDataInfo
}
type GetNoticeListDataInfo struct {
	Device_uuid string
}
type GetNoticeListRetDataInfo struct {
	Seq          int32
	Notice_name  string
	Location_url string
	Img_url      string
}
type GetNoticeListReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        []any
	Maintenance MaintenanceData
}
type GetPassGoodsShopData struct {
	I_id            int32
	I_TicketID      int32
	I_GoodsTicketID int32
	I_ShopID        int32
	B_IsActive      int16
	I_SaleShopID    int32
}
type GetPercentData struct {
	I_GroupID        int32
	I_RewardType     int32
	I_RewardID       int32
	L_RewardQuantity int32
	B_IsActive       int16
	I_id             int32
}
type GetPost struct {
	Call        string
	Data        GetPostDataInfo
	Common_data ParamData
}
type GetPostDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
}
type GetPostList struct {
	Idx                 int64
	Notice_type         int16
	Title_ko            string
	Memo_ko             string
	Title_en            string
	Memo_en             string
	Title_jp            string
	Memo_jp             string
	Title_zh_chs        string
	Memo_zh_chs         string
	Title_zh_cht        string
	Memo_zh_cht         string
	Title_vi            string
	Memo_vi             string
	Title_es            string
	Memo_es             string
	Title_it            string
	Memo_it             string
	Title_id            string
	Memo_id             string
	Title_th            string
	Memo_th             string
	Title_pt            string
	Memo_pt             string
	Title_hi            string
	Memo_hi             string
	Have_reward         int16
	Status              int16
	Unlimit_flg         int16
	Flg                 int16
	Create_time         int64
	Del_time            int64
	Item_list           []any
	Url                 string
	Image_resource_name string
}
type GetPostRetDataInfo struct {
	Server_time int64
	Post_list   []any
}
type GetPostReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        GetPostRetDataInfo
	Maintenance MaintenanceData
}
type GetPostTime struct {
	Call string
	Data GetPostTimeDataInfo
}
type GetPostTimeDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
}
type GetPostTimeRetDataInfo struct {
	Time int64
}
type GetPostTimeReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        GetPostTimeRetDataInfo
	Maintenance MaintenanceData
}
type GetProfile struct {
	Call string
	Data GetProfileDataInfo
}
type GetProfileDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Get_u_seq   int64
}
type GetProfileMusic struct {
	Call string
	Data GetProfileMusicDataInfo
}
type GetProfileMusicDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Get_u_seq   int64
	Start_limit int64
	End_limit   int64
}
type GetProfileMusicRetDataInfo struct {
	Music_idx    int32
	Single_cnt   int64
	Multi_cnt    int64
	Master_cnt   int64
	Single_score int64
	Multi_score  int64
}
type GetProfileMusicReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        []any
	Maintenance MaintenanceData
}
type GetProfileRetDataInfo struct {
	U_avatar        int32
	U_id            string
	U_nick          string
	U_country       string
	U_title         int32
	Allper          int32
	Allcom          int32
	Gold            int32
	Silver          int32
	Bronze          int32
	Single_cnt      int64
	Multi_cnt       int64
	Master_cnt      int64
	Single_score    int64
	Multi_score     int64
	Local_rank      int64
	Individual_rank int64
}
type GetProfileReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        GetProfileRetDataInfo
	Maintenance MaintenanceData
}
type GetPropData struct {
	I_id                 int32
	S_Name_KO            string
	S_Name_EN            string
	S_Name_JA            string
	S_Name_ZH_CHS        string
	S_Name_ZH_CHT        string
	S_Name_VI            string
	S_Name_ES            string
	S_Name_IT            string
	S_Name_ID            string
	S_Name_TH            string
	S_Name_PT            string
	S_Name_HI            string
	S_Description_KO     string
	S_Description_EN     string
	S_Description_JA     string
	S_Description_ZH_CHS string
	S_Description_ZH_CHT string
	S_Description_VI     string
	S_Description_ES     string
	S_Description_IT     string
	S_Description_ID     string
	S_Description_TH     string
	S_Description_PT     string
	S_Description_HI     string
	S_ResourceName       string
	I_RequirementPropId  int64
	I_Area               int64
	B_IsActive           int64
	I_MaxLevel           int16
}
type GetPropLevelData struct {
	I_id         int32
	I_PropId     int64
	I_Level      int64
	D_Cost       float64
	I_PerksId    int64
	D_PerksValue float64
	B_IsActive   int16
}
type GetPurchaseDiamondLog struct {
	Call string
	Data GetPurchaseDiamondLogDataInfo
}
type GetPurchaseDiamondLogDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
}
type GetPurchaseDiamondLogRetDataInfo struct {
	Bonus1 int64
	Bonus2 int64
	Bonus3 int64
	Bonus4 int64
	Bonus5 int64
	Bonus6 int64
}
type GetPurchaseDiamondLogReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        GetPurchaseDiamondLogRetDataInfo
	Maintenance MaintenanceData
}
type GetRankMain struct {
	Call string
	Data GetRankMainDataInfo
}
type GetRankMainDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
}
type GetRankMainRetDataInfo struct {
	My_country_rank       GetCountryMyRankRet
	Total_country_my_rank GetTotalCountryRankRet
	Country_rank          []any
	Masters_rank          GetMastersRankRet
}
type GetRankMainReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        GetRankMainRetDataInfo
	Maintenance MaintenanceData
}
type GetReviewPoint struct {
	Call string
	Data GetAlbumDataInfo
}
type GetReviewPointDataInfo struct {
	Null string
}
type GetReviewPointRetDataInfo struct {
	Music_idx  int32
	Difficulty int16
	Point      float64
}
type GetReviewPointReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        []any
	Maintenance MaintenanceData
}
type GetRewardGroupData struct {
	I_id               int64
	I_Group            int64
	I_RewardType       int64
	I_RewardID         int64
	L_RewardQuantity   int64
	I_BuyFirstQuantity int64
}
type GetSamSeckEventData struct {
	I_id                 int32
	S_Description_KO     string
	S_Description_EN     string
	S_Description_JA     string
	S_Description_ZH_CHS string
	S_Description_ZH_CHT string
	S_Description_VI     string
	S_Description_ES     string
	S_Description_IT     string
	S_Description_ID     string
	S_Description_TH     string
	S_Description_PT     string
	S_Description_HI     string
	S_ConditionType      string
	D_Condition          float64
	I_MailRewardID       int64
	B_IsActive           int16
}
type GetSamSeckList struct {
	Call        string
	Data        GetSamSeckListDataInfo
	Common_data ParamData
}
type GetSamSeckListDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
}
type GetSamSeckListRetDataInfo struct {
	Event_type string
	RewardList []any
}
type GetSamSeckListReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        GetSamSeckListRetDataInfo
	Maintenance MaintenanceData
}
type GetSamSeckReward struct {
	Call        string
	Data        GetSamSeckRewardDataInfo
	Common_data ParamData
}
type GetSamSeckRewardDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	I_id        int32
}
type GetSamSeckRewardRetDataInfo struct {
	Step int16
}
type GetSamSeckRewardReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        GetSamSeckRewardRetDataInfo
	Maintenance MaintenanceData
}
type GetSelectRewardData struct {
	I_id               int32
	I_GroupId          int32
	I_RewardGroupId    int32
	I_AltRewardGroupId int32
	B_IsActive         int16
}
type GetServerList struct {
	Call string
	Data GetServerListDataInfo
}
type GetServerListDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Type        string
}
type GetServerListRetDataInfo struct {
	User_data    UserData
	Channel_list []any
}
type GetServerListReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        GetServerListRetDataInfo
	Maintenance MaintenanceData
}
type GetServerTime struct {
	Call        string
	Data        MainDataInfo
	Sub_mode    string
	Common_data ParamData
}
type GetServerTimeDataInfo struct {
	Device_uuid string
}
type GetServerTimeRetDataInfo struct {
	Time     int64
	Datetime int64
}
type GetServerTimeReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        GetServerTimeRetDataInfo
	Maintenance MaintenanceData
}
type GetShopData struct {
	I_id                    int64
	S_ProductID_ios         string
	S_ProductID_aos         string
	I_ShopCategory          int64
	I_ProductType           int64
	S_ResourceName          string
	S_Title_KO              string
	S_Title_EN              string
	S_Title_JA              string
	S_Title_ZH_CHS          string
	S_Title_ZH_CHT          string
	S_Title_VI              string
	S_Title_ES              string
	S_Title_IT              string
	S_Title_ID              string
	S_Title_TH              string
	S_Title_PT              string
	S_Title_HI              string
	S_Description_KO        string
	S_Description_EN        string
	S_Description_JA        string
	S_Description_ZH_CHS    string
	S_Description_ZH_CHT    string
	S_Description_VI        string
	S_Description_ES        string
	S_Description_IT        string
	S_Description_ID        string
	S_Description_TH        string
	S_Description_PT        string
	S_Description_HI        string
	I_RewardGroup           int64
	I_Tag                   int64
	I_SellType              int64
	I_SellValue             int64
	B_IsActive              int16
	I_Condition             int64
	I_ConditionValue        int64
	S_StorePrice_AOS        string
	S_StorePrice_IOS        string
	I_SortIndex             int16
	S_AreaList              string
	S_AltResourceName       string
	S_AltDescription_KO     string
	S_AltDescription_EN     string
	S_AltDescription_JA     string
	S_AltDescription_ZH_CHS string
	S_AltDescription_ZH_CHT string
	S_AltDescription_VI     string
	S_AltDescription_ES     string
	S_AltDescription_IT     string
	S_AltDescription_ID     string
	S_AltDescription_TH     string
	S_AltDescription_PT     string
	S_AltDescription_HI     string
	S_Kor_StorePrice_AOS    string
	S_Kor_StorePrice_IOS    string
	B_IsLimitTime           int16
	S_UIStartTime           string
	S_StartTime             string
	S_EndTime               string
}
type GetSkillData struct {
	I_id                      int32
	I_Area                    int16
	S_Name_KO                 string
	S_Name_EN                 string
	S_Name_JA                 string
	S_Name_ZH_CHS             string
	S_Name_ZH_CHT             string
	S_Name_VI                 string
	S_Name_ES                 string
	S_Name_IT                 string
	S_Name_ID                 string
	S_Name_TH                 string
	S_Name_PT                 string
	S_Name_HI                 string
	S_Description_KO          string
	S_Description_EN          string
	S_Description_JA          string
	S_Description_ZH_CHS      string
	S_Description_ZH_CHT      string
	S_Description_VI          string
	S_Description_ES          string
	S_Description_IT          string
	S_Description_ID          string
	S_Description_TH          string
	S_Description_PT          string
	S_Description_HI          string
	S_ResourceName            string
	I_MaxLevel                int64
	S_GoodsType               string
	I_Cost                    float64
	F_CostIncreaseValue       float64
	F_SkillValue              float64
	F_SkillIncreaseValue      float64
	F_SkillTime               float64
	F_SkillTimeIncreaseValue  float64
	F_Cooltime                float64
	I_UnlockLevel             int64
	I_UnlockCost              int64
	I_RequirementSkillId_1    int64
	I_RequirementSkillId_2    int64
	I_RequirementFollowerId_1 int64
	I_RequirementFollowerId_2 int64
	I_RequirementPropId_1     int64
	I_RequirementPropId_2     int64
	I_RequirementPropCount    int64
}
type GetSubscribeList struct {
	I_id              int64
	I_Area            int16
	S_Type            string
	I_TimeLimit       int16
	I_StartYear       int32
	I_RepeatMonth     int32
	I_MonthGroupIndex int32
	B_IsActive        int16
}
type GetSubscribePass struct {
	I_id                 int64
	I_SubscribeID        int64
	S_AcquireList        string
	I_PointPrice         int16
	I_PaidPoint          int64
	I_ActiveBuffID       int64
	I_HiddenStep         int64
	I_ADPoint            int64
	I_ADCoolTime         int64
	I_TicketCollectionId int64
}
type GetSubscribePassReward struct {
	I_id              int64
	I_Group           int64
	I_Step            int16
	I_Goal            int64
	I_FreeRewardGroup int64
	I_PaidRewardGroup int64
}
type GetSubscribePassRewardInformation struct {
	I_id                 int32
	S_ResourceName       string
	S_Description_KO     string
	S_Description_EN     string
	S_Description_JA     string
	S_Description_ZH_CHS string
	S_Description_ZH_CHT string
	S_Description_VI     string
	S_Description_ES     string
	S_Description_IT     string
	S_Description_ID     string
	S_Description_TH     string
	S_Description_PT     string
	S_Description_HI     string
	B_IsActive           int16
}
type GetSystemNotificationData struct {
	I_Id                 int64
	S_DescriptionType    string
	I_Type               string
	S_Factor             string
	S_Description_KO     string
	S_Description_EN     string
	S_Description_JA     string
	S_Description_ZH_CHS string
	S_Description_ZH_CHT string
	S_Description_VI     string
	S_Description_ES     string
	S_Description_IT     string
	S_Description_ID     string
	S_Description_TH     string
	S_Description_PT     string
	S_Description_HI     string
}
type GetSystemStringData struct {
	I_StringID int64
	S_Category string
	S_UIType   string
	S_Key      string
	S_KO       string
	S_EN       string
	S_JA       string
	S_ZH_CHS   string
	S_ZH_CHT   string
	S_ID       string
	S_PT       string
	S_ES       string
	S_RU       string
	S_DE       string
	S_AR_XA    string
	S_VI       string
	S_IT       string
	S_TH       string
	S_HI       string
}
type GetTabList struct {
	Call string
	Data InitDataInfo
}
type GetTabListDataInfo struct {
	Type        string
	Device_uuid string
}
type GetTabListRetDataInfo struct {
	Tab_list []any
}
type GetTabListReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        map[any]any
	Maintenance MaintenanceData
}
type GetTicketCollectionData struct {
	I_id                 int32
	S_TicketType         string
	S_Title_KO           string
	S_Title_EN           string
	S_Title_JA           string
	S_Title_ZH_CHS       string
	S_Title_ZH_CHT       string
	S_Title_VI           string
	S_Title_ES           string
	S_Title_IT           string
	S_Title_ID           string
	S_Title_TH           string
	S_Title_PT           string
	S_Title_HI           string
	S_Description_KO     string
	S_Description_EN     string
	S_Description_JA     string
	S_Description_ZH_CHS string
	S_Description_ZH_CHT string
	S_Description_VI     string
	S_Description_ES     string
	S_Description_IT     string
	S_Description_ID     string
	S_Description_TH     string
	S_Description_PT     string
	S_Description_HI     string
	S_ResourceName       string
	S_ResourceName_Big   string
	I_Season             int16
	S_PassUIBundleName   string
	S_BackgroundColor    string
	S_LabelColor         string
	S_SeasonLabelColor   string
	S_SeasonBackColor    string
	I_SortOrder          int32
	B_IsActive           int16
}
type GetTotalCountryRank struct {
	Call string
	Data GetTotalCountryRankDataInfo
}
type GetTotalCountryRankDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Start_limit int32
	End_limit   int32
}
type GetTotalCountryRankRet struct {
	G           int16
	S           int16
	B           int16
	Total_score int32
	U_avatar    int16
	U_id        string
	U_nick      string
	U_country   string
	Rank        int32
}
type GetTotalCountryRankRetDataInfo struct {
	Start_limit int16
	End_limit   int16
	My_rank     GetTotalCountryRankRet
	Rank_list   []any
}
type GetTotalCountryRankReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        GetTotalCountryRankRetDataInfo
	Maintenance MaintenanceData
}
type GetTotalCountryUserRankRet struct {
	G           int16
	S           int16
	B           int16
	Total_score int32
	U_avatar    int16
	U_id        string
	U_nick      string
	U_country   string
}
type GetTotalMastersRankRet struct {
	All_per   int16
	All_com   int16
	Score     int32
	U_avatar  int16
	U_id      string
	U_nick    string
	U_country string
}
type GetTotalMusicMyRank struct {
	Grade string
	Rank  int32
	Score int32
}
type GetTotalMusicRank struct {
	Call string
	Data GetTotalMusicRankDataInfo
}
type GetTotalMusicRankDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Music_idx   int32
	Start_limit int32
	End_limit   int32
}
type GetTotalMusicRankList struct {
	Rank    int32
	Id      string
	Nick    string
	Country string
	Score   int32
}
type GetTotalMusicRankList1 struct {
	Rank  int32
	Id    string
	Score int32
}
type GetTotalMusicRankList2 struct {
	Rank  int32
	Id    string
	Score int32
}
type GetTotalMusicRankList3 struct {
	Rank  int32
	Id    string
	Score int32
}
type GetTotalMusicRankRetData struct {
	Total_rank_list []any
	Rank_list1      []any
	Rank_list2      []any
	Rank_list3      []any
	My_rank_list    []any
}
type GetTotalMusicRankRetDataInfo struct {
	Total_rank_list []any
	Rank_list1      []any
	Rank_list2      []any
	Rank_list3      []any
	My_rank_list    []any
}
type GetTotalMusicRankReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        []any
	Maintenance MaintenanceData
}
type GetTourList struct {
	Call string
	Data GetTourListDataInfo
}
type GetTourListDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
}
type GetTourListRetDataInfo struct {
	Tour_idx int32
	Title    string
}
type GetTourListReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        []any
	Maintenance MaintenanceData
}
type GetTourRank struct {
	Call string
	Data GetTourRankDataInfo
}
type GetTourRankData struct {
	U_seq     int64
	U_id      string
	U_nick    string
	U_avatar  int32
	U_country string
	Rank      int64
	Score     int64
	Medal     []any
}
type GetTourRankDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Tour_idx    int32
	Track_idx   int32
	Mode        string
	Start_limit int64
	End_limit   int64
}
type GetTourRankRetDataInfo struct {
	My_rank GetTourRankData
	Rank    []any
}
type GetTourRankReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        GetTourRankRetDataInfo
	Maintenance MaintenanceData
}
type GetTransferId struct {
	Call string
	Data GetTransferIdDataInfo
}
type GetTransferIdDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
}
type GetTransferIdRetDataInfo struct {
	Transfer_id string
}
type GetTransferIdReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        GetTransferIdRetDataInfo
	Maintenance MaintenanceData
}
type GetUnitData struct {
	I_id                      int32
	S_Name_KO                 string
	S_Name_EN                 string
	S_Name_JA                 string
	S_Name_ZH_CHS             string
	S_Name_ZH_CHT             string
	S_Name_VI                 string
	S_Name_ES                 string
	S_Name_IT                 string
	S_Name_ID                 string
	S_Name_TH                 string
	S_Name_PT                 string
	S_Name_HI                 string
	S_Description_KO          string
	S_Description_EN          string
	S_Description_JA          string
	S_Description_ZH_CHS      string
	S_Description_ZH_CHT      string
	S_Description_VI          string
	S_Description_ES          string
	S_Description_IT          string
	S_Description_ID          string
	S_Description_TH          string
	S_Description_PT          string
	S_Description_HI          string
	S_ResourceName            string
	I_MaxLevel                int64
	S_GoodsType               string
	I_Cost                    float64
	F_CostIncreaseValue       float64
	F_Value                   float64
	F_IncreaseValue           float64
	I_UnlockLevel             int64
	I_UnlockCost              int64
	I_RequirementUnitId_1     int64
	I_RequirementUnitId_2     int64
	I_RequirementFollowerId_1 int64
	I_RequirementFollowerId_2 int64
	I_RequirementPropId_1     int64
	I_RequirementPropId_2     int64
	I_RequirementPropCount    int64
	I_Area                    int16
	B_IsActive                int16
}
type GetUpdateTime struct {
	Call        string
	Data        GetUpdateTimeDataInfo
	Common_data ParamData
}
type GetUpdateTimeDataInfo struct {
	Device_uuid string
}
type GetUpdateTimeRetDataInfo struct {
	Upd_time int64
}
type GetUpdateTimeReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        GetUpdateTimeRetDataInfo
	Maintenance MaintenanceData
}
type GetUserCollection struct {
	Call string
	Data GetUserCollectionDataInfo
}
type GetUserCollectionDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
}
type GetUserCollectionRetDataInfo struct {
	Music_idx   int64
	Value       int64
	Get_type    int16
	Get_time    string
	Create_time string
}
type GetUserCollectionReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        []any
	Maintenance MaintenanceData
}
type GetUserInfo struct {
	Call        string
	Data        GetUserInfoDataInfo
	Common_data ParamData
}
type GetUserInfoDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
}
type GetUserInfoRetDataInfo struct {
	U_cp    int64
	U_candy float64
	U_like  float64
	U_fans  int64
}
type GetUserInfoReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        GetUserInfoRetDataInfo
	Maintenance MaintenanceData
}
type GetUserItemInven struct {
	Call string
	Data GetUserItemInvenDataInfo
}
type GetUserItemInvenDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
}
type GetUserItemInvenRetDataInfo struct {
	Item_inven_list map[any]any
	Empty           string
}
type GetUserItemInvenReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        GetUserItemInvenRetDataInfo
	Maintenance MaintenanceData
}
type GetVarietyStore struct {
	Call        string
	Data        GetVarietyStoreDataInfo
	Common_data ParamData
}
type GetVarietyStoreDataInfo struct {
	Device_uuid string
}
type GetVarietyStoreRetDataInfo struct {
	I_id                 int32
	S_ResourceName       string
	S_Title_KO           string
	S_Title_EN           string
	S_Title_JA           string
	S_Title_ZH_CHS       string
	S_Title_ZH_CHT       string
	S_Title_VI           string
	S_Title_ES           string
	S_Title_IT           string
	S_Title_ID           string
	S_Title_TH           string
	S_Title_PT           string
	S_Title_HI           string
	S_Description_KO     string
	S_Description_EN     string
	S_Description_JA     string
	S_Description_ZH_CHS string
	S_Description_ZH_CHT string
	S_Description_VI     string
	S_Description_ES     string
	S_Description_IT     string
	S_Description_ID     string
	S_Description_TH     string
	S_Description_PT     string
	S_Description_HI     string
	I_Cost               int32
	I_RewardType         int32
	I_RewardId           int32
	I_RewardQuantity     int32
}
type GetVarietyStoreReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        []any
	Maintenance MaintenanceData
}
type HibernationLog struct {
	Call string
	Data HibernationLogDataInfo
}
type HibernationLogDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
}
type HibernationLogRetDataInfo struct {
	Action int32
}
type HibernationLogReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        HibernationLogRetDataInfo
	Maintenance MaintenanceData
}
type InitDataInfo struct {
	Type        string
	Os          int16
	Ver         int32
	Device_uuid string
}
type InitRetDataInfo struct {
	Idx      int16
	Game_url string
	Cdn_url  string
}
type InitReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        InitRetDataInfo
	Maintenance MaintenanceData
}
type ItemList struct {
	I_RewardType     int32
	I_RewardId       int32
	D_RewardQuantity float64
}
type ItemPurchase struct {
	Call     string
	Data     ItemPurchaseDataInfo
	Sub_mode string
}
type ItemPurchaseDataInfo struct {
	U_seq             int32
	U_id              string
	Uuid              string
	Device_uuid       string
	Developer_payload string
	Pay_id            string
	Purchase_token    string
	Purchase_id       string
	Os                int16
	Restore           int16
}
type ItemPurchaseRetDataInfo struct {
	Item_type   int16
	Item_value  int64
	Pay_id      string
	Product_id  string
	Purchase_id string
}
type ItemPurchaseReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        ItemPurchaseRetDataInfo
	Maintenance MaintenanceData
}
type ItemStoreList struct {
	Call        string
	Data        ItemStoreListDataInfo
	Common_data ParamData
}
type ItemStoreListDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Store_type  int16
}
type ItemStoreListRetDataInfo struct {
	Bonus             int64
	Idx               int64
	Pp_idx            int64
	Title_cn          string
	Title_en          string
	Title_jp          string
	Desc_cn           string
	Desc_en           string
	Desc_jp           string
	Img_url           string
	Buy_type          int16
	Item_type         int16
	Prev_price        float64
	Price             float64
	Str_prev_price    string
	Str_price         string
	Item_value        int64
	Sale_status       int16
	Aos_pp_product_id string
	Ios_pp_product_id string
}
type ItemStoreListReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        []any
	Maintenance MaintenanceData
}
type LastSaveTime struct {
	Call        string
	Data        LastSaveTimeDataInfo
	Common_data ParamData
}
type LastSaveTimeDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
}
type LastSaveTimeRetDataInfo struct {
	Last_save_time int64
	Device_uuid    string
}
type LastSaveTimeReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        LastSaveTimeRetDataInfo
	Maintenance MaintenanceData
}
type LogCashComplete struct {
	Call string
	Data LogCashCompleteDataInfo
}
type LogCashCompleteDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Shop_id     int32
	Value       string
}
type LogCashCompleteRetDataInfo struct {
	Status string
}
type LogCashCompleteReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        LogCashCompleteRetDataInfo
	Maintenance MaintenanceData
}
type Main struct {
	Call     string
	Data     MainDataInfo
	Sub_mode string
}
type MainDataInfo struct {
	U_seq        int32
	U_id         string
	Uuid         string
	Device_uuid  string
	Country_code string
}
type MainFeatures struct {
	Album_idx    int32
	F_type       int16
	Buy_type     int16
	New_status   int16
	Discount     int16
	Price        float64
	Banner_url   string
	Location_url string
}
type MainGameInfo struct {
	Album_list      []any
	Music_list      []any
	Album_language  map[any]any
	Music_language  map[any]any
	Tab             map[any]any
	Notice          []any
	Default_setting GetDefaultSettingRetDataInfo
}
type MainRetDataInfo struct {
	GameInfo MainGameInfo
	Features []any
}
type MainReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        MainRetDataInfo
	Maintenance MaintenanceData
}
type MaintenanceData struct {
	Code           int16
	Title          string
	Description    string
	Utc_time       int16
	Facebook_url   string
	Start_datetime string
	End_datetime   string
}
type MaintenanceList struct {
	Idx            int16
	Title          string
	Start_datetime int32
	End_datetime   int32
	View_flg       int16
}
type MoreGames struct {
	Call string
	Data MoreGamesDataInfo
}
type MoreGamesDataInfo struct {
	Locale     string
	Adset_type string
	Uuid       string
}
type MoreGamesRetDataInfo struct {
	App_id            int32
	Material_id       int32
	Material_type     string
	Material_category string
	Material_name     string
	Material_link     string
	Link_text         string
	Link_url_aos      string
	Link_url_ios      string
	Thumbnail_url     string
	Icon_url          string
	Update_date       string
}
type MoreGamesReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        map[any]any
	Maintenance MaintenanceData
}
type MusicData struct {
	Music_idx int32
	Score1    int32
	Score2    int32
	Score3    int32
}
type MusicListInfo struct {
	Music_idx   int64
	Genre_idx   int32
	Difficulty1 int16
	Difficulty2 int16
	Difficulty3 int16
	Difficulty4 int16
	Buy_type    int16
	Free_type   int16
	Price       float64
	Buy_ad      int32
	Music_ver   int32
	Bpm         float64
	Cdn_dir     string
	Track_type  int16
	File_type   int16
	New_status  int16
	Update_time int32
	Language    map[any]any
}
type MusicListLanguage struct {
	Title      string
	Artist     string
	Songwriter string
	Lyricist   string
}
type MusicPointReview struct {
	Call string
	Data MusicPointReviewDataInfo
}
type MusicPointReviewDataInfo struct {
	U_seq        int32
	U_id         string
	Uuid         string
	Device_uuid  string
	Review_point []any
}
type MusicPointReviewRetDataInfo struct {
	Music_idx  int32
	Point      float64
	Difficulty int16
}
type MusicPointReviewReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        []any
	Maintenance MaintenanceData
}
type MusicProduct struct {
	Product_id     int32
	Product_type   int16
	Product_name   string
	Store_name     string
	IapID          string
	Price          float64
	Prev_price     float64
	Price_str      string
	Prev_price_str string
	Currency       string
	Image          string
	Thumbnail      string
	Main_desc      string
}
type NewStoreList struct {
	Call string
	Data NewStoreListDataInfo
}
type NewStoreListData struct {
	Store_idx int16
	Music_idx int32
}
type NewStoreListDataInfo struct {
	Os           int16
	Type         string
	Country_code string
	U_seq        int32
	U_id         string
	Uuid         string
	Device_uuid  string
}
type NewStoreListRetDataInfo struct {
	Store_idx    int16
	Store_type   int16
	Buy_type     int16
	Prev_price   int64
	Price        int64
	Title        string
	Sub_title    string
	Img_url      string
	Store_list   map[any]any
	Product_data NewStoreProduct
	Is_sale      int16
}
type NewStoreListReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        map[any]any
	Maintenance MaintenanceData
}
type NewStoreProduct struct {
	Product_id     int32
	Product_type   int16
	Product_name   string
	Store_name     string
	IapID          string
	Price          float64
	Prev_price     float64
	Price_str      string
	Prev_price_str string
	Currency       string
	Image          string
	Thumbnail      string
	Main_desc      string
}
type PaidEventPoint struct {
	Call        string
	Data        PaidEventPointDataInfo
	Common_data ParamData
}
type PaidEventPointDataInfo struct {
	U_seq         int32
	U_id          string
	Uuid          string
	Device_uuid   string
	Type          int16
	I_SubscribeID int64
	I_Version     int32
}
type PaidEventPointRetDataInfo struct {
	U_cp          int64
	U_candy       float64
	I_SubscribeID int64
	I_Point       int64
	I_Version     int32
}
type PaidEventPointReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        PaidEventPointRetDataInfo
	Maintenance MaintenanceData
}
type ParamData struct {
	Client_ver int32
	Type       string
	Os         int16
}
type PlayCheck struct {
	Call string
	Data PlayCheckDataInfo
}
type PlayCheckDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Tour_idx    int32
	Track_idx   int32
	Music_idx   int32
}
type PlayCheckRetDataInfo struct {
	Basic       int32
	Pro         int32
	Legend      int32
	Extra       int32
	Server_time int64
}
type PlayCheckReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        PlayCheckRetDataInfo
	Maintenance MaintenanceData
}
type PlayMusic struct {
	Call string
	Data PlayMusicDataInfo
}
type PlayMusicDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Music_idx   int32
	Difficulty  int32
}
type PlayMusicRetDataInfo struct {
	Code      int16
	Album_idx int32
	Music_idx int32
	Score1    int32
	Score2    int32
	Score3    int32
	Grade1    int16
	Grade2    int16
	Grade3    int16
	Play_cnt1 int32
	Play_cnt2 int32
	Play_cnt3 int32
	Buy_type  int16
	Price     float64
}
type PlayMusicReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        PlayMusicRetDataInfo
	Maintenance MaintenanceData
}
type ProvidePost struct {
	Call        string
	Data        ProvidePostDataInfo
	Common_data ParamData
}
type ProvidePostDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Idx         int64
	Type        int16
}
type ProvidePostRetDataInfo struct {
	I_RewardType     int32
	I_RewardId       int32
	D_RewardQuantity float64
}
type ProvidePostReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        []any
	Maintenance MaintenanceData
}
type ReadPost struct {
	Call string
	Data ReadPostDataInfo
}
type ReadPostDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Idx         int64
}
type ReadPostRetDataInfo struct {
	Idx int64
}
type ReadPostReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        ReadPostRetDataInfo
	Maintenance MaintenanceData
}
type RetReward struct {
	Reward_type  int16
	Reward_id    int32
	Reward_value float64
}
type ReviewDataList struct {
	Music_idx  int32
	Point      float64
	Difficulty int16
}
type RewardList struct {
	I_id             int64
	I_RewardType     int32
	I_RewardId       int32
	D_RewardQuantity float64
}
type RewardListData struct {
	I_id             int64
	I_RewardType     int32
	I_RewardId       int32
	D_RewardQuantity float64
}
type SaveUserAchievement struct {
	I_id       int64
	D_Quantity float64
	S_Quantity string
}
type SaveUserAreaInfo struct {
	U_area_num          int16
	D_Like              float64
	I_UserFanCount      int64
	I_UserFanGrade      int16
	I_SelectedCostumeId int64
	I_SelectedMusicId   int64
	I_SelectedGuitarId  int64
	D_Candy             float64
	S_TutorialList      string
	S_Gp1               string
	S_Gp2               string
}
type SaveUserCharacter struct {
	I_id         int64
	I_Level      int64
	I_BonusLevel int64
}
type SaveUserCostume struct {
	I_id         int64
	I_Level      int64
	I_BonusLevel int64
}
type SaveUserDailyMission struct {
	I_id       int64
	D_Quantity int64
}
type SaveUserEventPoint struct {
	S_EventType string
	I_DataID    int64
	I_Point     int64
	I_Step      int64
	I_Version   int32
}
type SaveUserFollower struct {
	I_id         int64
	I_Level      int64
	I_BonusLevel int64
}
type SaveUserFollowerQuest struct {
	I_CurrentID       int32
	D_ConditionValue1 float64
	D_ConditionValue2 float64
	D_ConditionValue3 float64
}
type SaveUserGuitar struct {
	I_id         int64
	I_Level      int64
	I_BonusLevel int64
}
type SaveUserInfo struct {
	U_like                float64
	U_fans                int64
	U_fans_grade          int16
	U_selected_costume_id int64
	U_selected_music_id   int64
}
type SaveUserMessenger struct {
	I_MessengerChatRoomId int64
	I_LastConfirmIndex    int64
	S_UnlockGroupList     string
	L_UpdateTimeTicks     int64
}
type SaveUserMusic struct {
	I_id                    int64
	I_Level                 int64
	I_BonusLevel            int64
	B_EncoreBonusAppear     int64
	I_EncoreBonusFollowerId int64
}
type SaveUserSkill struct {
	I_id               int64
	B_Activate         int16
	L_ActivateOnTicks  int64
	L_ActivateOffTicks int64
}
type ScoreDataList struct {
	Music_idx     int32
	Score         int32
	Grade         int16
	Play_cnt      int32
	Difficulty    int16
	Multiple      float64
	Gp            int32
	Achievement   int16
	Play_type     int16
	Play_datetime string
	Mission_clear int16
	Combo_grade   int16
	Is_credit     int16
}
type ServerTimeRet struct {
	Time     int32
	Datetime int64
}
type SetAdReward struct {
	Call        string
	Data        SetAdRewardDataInfo
	Common_data ParamData
}
type SetAdRewardDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	I_id        int32
	Param1      string
	Param2      string
	Param3      string
}
type SetAdRewardRetDataInfo struct {
	I_id                  int32
	User_ad_list          UserAdList
	Reward_data           []any
	User_follower_profile UserFollowerProfile
}
type SetAdRewardReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        SetAdRewardRetDataInfo
	Maintenance MaintenanceData
}
type SetAttendance struct {
	Call        string
	Data        SetAttendanceDataInfo
	Common_data ParamData
}
type SetAttendanceDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Type        string
}
type SetAttendanceRetDataInfo struct {
	Status                          string
	User_follower_quest             UserFollowerQuest
	Attendance_count                int32
	Attendance_date                 int32
	Max_coutinuous_attendance_count int32
}
type SetAttendanceReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        SetAttendanceRetDataInfo
	Maintenance MaintenanceData
}
type SetBookmark struct {
	Call string
	Data SetBookmarkDataInfo
}
type SetBookmarkDataInfo struct {
	U_seq         int32
	U_id          string
	Uuid          string
	Device_uuid   string
	Bookmark_list []any
}
type SetBookmarkDataList struct {
	Music_idx int32
	Flag      int16
}
type SetBookmarkRetDataInfo struct {
	Music_idx int32
	Flag      int16
}
type SetBookmarkReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        []any
	Maintenance MaintenanceData
}
type SetEventReward struct {
	Call        string
	Data        SetEventRewardDataInfo
	Common_data ParamData
}
type SetEventRewardDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Event_idx   int32
}
type SetEventRewardItem struct {
	Item_type int32
	Item_id   int32
	Count     int64
}
type SetEventRewardRetDataInfo struct {
	U_cp         int64
	U_candy      float64
	U_like       float64
	U_fans       int64
	Reward_type  int16
	Reward_id    int32
	Reward_value int32
	Status       string
	Reward_data  []any
}
type SetEventRewardReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        SetEventRewardRetDataInfo
	Maintenance MaintenanceData
}
type SetFollowerProfileGift struct {
	Call        string
	Data        SetFollowerProfileGiftDataInfo
	Common_data ParamData
}
type SetFollowerProfileGiftDataInfo struct {
	U_seq          int32
	U_id           string
	Uuid           string
	Device_uuid    string
	Profile_id     int32
	Gift_id        int32
	Use_gitf_value int32
}
type SetFollowerProfileGiftRetDataInfo struct {
	I_gift_type            int32
	User_follower_giftitem UserFollowerGiftItem
	User_follower_profile  UserFollowerProfile
}
type SetFollowerProfileGiftReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        SetFollowerProfileGiftRetDataInfo
	Maintenance MaintenanceData
}
type SetFollowerQuestInfinite struct {
	Call        string
	Data        SetFollowerQuestInfiniteDataInfo
	Common_data ParamData
}
type SetFollowerQuestInfiniteDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
}
type SetFollowerQuestInfiniteRetDataInfo struct {
	U_seq               int32
	User_follower_quest UserFollowerQuest
}
type SetFollowerQuestInfiniteReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        SetFollowerQuestInfiniteRetDataInfo
	Maintenance MaintenanceData
}
type SetGameReward struct {
	Call        string
	Data        SetGameRewardDataInfo
	Common_data ParamData
}
type SetGameRewardDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Type        string
	Id          int32
	Level       int16
	Quantity    float64
	S_quantity  string
}
type SetGameRewardRetDataInfo struct {
	Type                  string
	Id                    int32
	Level                 int16
	Reward_type           string
	Reward_value          int64
	Status                string
	User_follower_profile UserFollowerProfile
}
type SetGameRewardReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        SetGameRewardRetDataInfo
	Maintenance MaintenanceData
}
type SetPassReward struct {
	Call        string
	Data        SetPassRewardDataInfo
	Common_data ParamData
}
type SetPassRewardDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Group       int32
	Step        int32
	Type        int16
	I_Version   int32
}
type SetPassRewardRetDataInfo struct {
	Subscribe_pass_reward UserSubscribePassReward
	Reward_data           []any
}
type SetPassRewardReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        SetPassRewardRetDataInfo
	Maintenance MaintenanceData
}
type SetReviewPopup struct {
	Call        string
	Data        SetReviewPopupDataInfo
	Common_data ParamData
}
type SetReviewPopupDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
}
type SetReviewPopupRetDataInfo struct {
	Status string
}
type SetReviewPopupReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        SetReviewPopupRetDataInfo
	Maintenance MaintenanceData
}
type SetSelectReward struct {
	Call        string
	Data        SetSelectRewardDataInfo
	Common_data ParamData
}
type SetSelectRewardDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	I_id        int32
}
type SetSelectRewardInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	I_id        int32
}
type SetSelectRewardRetDataInfo struct {
	U_seq                   int32
	User_select_reward_list []any
}
type SetSelectRewardReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        SetSelectRewardRetDataInfo
	Maintenance MaintenanceData
}
type SetSubscribe struct {
	Call        string
	Data        SetSubscribeDataInfo
	Common_data ParamData
}
type SetSubscribeDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	I_ids       []any
}
type SetSubscribeRetDataInfo struct {
	U_seq               int32
	User_subscribe_list []any
}
type SetSubscribeReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        SetSubscribeRetDataInfo
	Maintenance MaintenanceData
}
type SetTutorial struct {
	Call        string
	Data        SetTutorialDataInfo
	Common_data ParamData
}
type SetTutorialDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Step        string
}
type SetTutorialNew struct {
	Call        string
	Data        SetTutorialNewDataInfo
	Common_data ParamData
}
type SetTutorialNewDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	I_ids       []any
}
type SetTutorialNewRetDataInfo struct {
	U_seq    int32
	Tutorial []any
}
type SetTutorialNewReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        SetTutorialNewRetDataInfo
	Maintenance MaintenanceData
}
type SetTutorialRetDataInfo struct {
	Step string
}
type SetTutorialReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        SetTutorialRetDataInfo
	Maintenance MaintenanceData
}
type SetUserFollowerProfileReward struct {
	Call        string
	Data        SetUserFollowerProfileRewardDataInfo
	Common_data ParamData
}
type SetUserFollowerProfileRewardDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	I_id        int32
	S_level     string
}
type SetUserFollowerProfileRewardRetDataInfo struct {
	I_id                  int32
	I_level               int32
	Reward_data           []any
	User_follower_profile UserFollowerProfile
}
type SetUserFollowerProfileRewardReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        []any
	Maintenance MaintenanceData
}
type SetUserFollowerQuest struct {
	Call        string
	Data        SetUserFollowerQuestDataInfo
	Common_data ParamData
}
type SetUserFollowerQuestDataInfo struct {
	U_seq            int32
	U_id             string
	Uuid             string
	Device_uuid      string
	I_CurrentID      int32
	I_SubID          int32
	D_ConditionValue float64
}
type SetUserFollowerQuestRetDataInfo struct {
	I_CurrentID       int32
	I_CompleteID      int32
	D_ConditionValue1 float64
	D_ConditionValue2 float64
	D_ConditionValue3 float64
	I_RewardReceived1 int16
	I_RewardReceived2 int16
	I_RewardReceived3 int16
	Next_flg          int16
	Reward_data       []any
}
type SetUserFollowerQuestReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        SetUserFollowerQuestRetDataInfo
	Maintenance MaintenanceData
}
type SetUserGP struct {
	Call string
	Data SetUserGPDataInfo
}
type SetUserGPDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Gp          int32
}
type SetUserGPRetDataInfo struct {
	Gp int32
}
type SetUserGPReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        SetUserGPRetDataInfo
	Maintenance MaintenanceData
}
type StoreList struct {
	Call string
	Data StoreListDataInfo
}
type StoreListData struct {
	Store_idx int16
	Music_idx int32
}
type StoreListDataInfo struct {
	U_seq int32
	U_id  string
	Uuid  string
}
type StoreListRetDataInfo struct {
	Store_idx  int16
	Store_type int16
	Title      string
	Sub_title  string
	Img_url    string
	Store_list map[any]any
}
type StoreListReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        map[any]any
	Maintenance MaintenanceData
}
type StoreNotice struct {
	Call string
	Data StoreNoticeDataInfo
}
type StoreNoticeDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
}
type StoreNoticeRetDataInfo struct {
	Idx            int32
	Product_id_aos int64
	Product_id_ios int64
	Front_img_url  string
	Back_img_url   string
	Main_desc      string
	Sub_desc       string
}
type StoreNoticeReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        []any
	Maintenance MaintenanceData
}
type TabListData struct {
	Tap_idx   int16
	Music_idx int32
}
type TitleList struct {
	Call string
	Data TitleListDataInfo
}
type TitleListDataInfo struct {
	Type string
}
type TitleListRetDataInfo struct {
	Idx           int16
	Resource_name string
	Down_status   int16
	Buy_type      int16
	Price         int32
	Buy_type2     int16
	Price2        int32
	Value         int32
	Event_point   int64
	Desc_ko       string
	Desc_en       string
	Desc_ja       string
}
type TitleListReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        []any
	Maintenance MaintenanceData
}
type TotalTopRank struct {
}
type TourMain struct {
	Call string
	Data TourMainDataInfo
}
type TourMainDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
}
type TourMainRetDataInfo struct {
	Idx         int32
	Title       string
	Start_time  int64
	End_time    int64
	Server_time int64
	Mode        string
	Tour_track  []any
	Tour_rank   TourRank
	Track_rank  map[any]any
	Utc         int16
}
type TourMainReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        TourMainRetDataInfo
	Maintenance MaintenanceData
}
type TourMusicList struct {
	Music_idx int32
}
type TourMyRank struct {
	Rank  int64
	Score int64
	Medal []any
}
type TourRank struct {
	My_rank  TourMyRank
	Top_rank []any
}
type TourRankData struct {
	Track []any
	Total TotalTopRank
}
type TourTrackList struct {
	Idx         int32
	Tour_idx    int32
	My_rank     int64
	Music_idx   string
	Start_time  int32
	End_time    int32
	Open_flg    int16
	Difficulty1 int16
	Difficulty2 int16
	Difficulty3 int16
	Difficulty4 int16
}
type TrackMyRank struct {
	Score int64
	Medal int16
}
type TrackMyScore struct {
	Score int64
	Medal int16
}
type TrackRank struct {
	My_score map[any]any
}
type TrackTopRank struct {
	U_seq     int64
	U_id      string
	U_nick    string
	U_country string
	Rank      int64
	Score     int64
}
type TransferUser struct {
	Call string
	Data TransferUserDataInfo
}
type TransferUserDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Email       string
}
type TransferUserRetDataInfo struct {
	Transfer_id string
}
type TransferUserReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        TransferUserRetDataInfo
	Maintenance MaintenanceData
}
type UpdateAchievement struct {
	Call string
	Data UpdateAchievementDataInfo
}
type UpdateAchievementDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Achievement int16
}
type UpdateAchievementRetDataInfo struct {
	Achievement int16
}
type UpdateAchievementReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        UpdateAchievementRetDataInfo
	Maintenance MaintenanceData
}
type UpdateAvatar struct {
	Call string
	Data UpdateAvatarDataInfo
}
type UpdateAvatarDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	U_avatar    int16
}
type UpdateAvatarRetDataInfo struct {
	U_avatar int16
}
type UpdateAvatarReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        UpdateAvatarRetDataInfo
	Maintenance MaintenanceData
}
type UpdateNickName struct {
	Call        string
	Data        UpdateNickNameDataInfo
	Common_data ParamData
}
type UpdateNickNameDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Nickname    string
}
type UpdateNickNameRetDataInfo struct {
	Nickname string
}
type UpdateNickNameReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        UpdateNickNameRetDataInfo
	Maintenance MaintenanceData
}
type UpdateScore struct {
	Call string
	Data UpdateScoreDataInfo
}
type UpdateScoreDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Score_list  []any
}
type UpdateScoreDataList struct {
	Music_idx   int32
	Score       int32
	Grade       int16
	Play_cnt    int32
	Difficulty  int16
	Combo_grade int16
	Collection  []any
}
type UpdateScoreRetDataInfo struct {
	Update_list []any
	U_gold      int64
	U_ep        int64
	Event_flg   int16
}
type UpdateScoreReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        UpdateScoreRetDataInfo
	Maintenance MaintenanceData
}
type UpdateUserTitle struct {
	Call string
	Data UpdateUserTitleDataInfo
}
type UpdateUserTitleDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	U_title     int16
}
type UpdateUserTitleRetDataInfo struct {
	U_title int16
}
type UpdateUserTitleReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        UpdateUserTitleRetDataInfo
	Maintenance MaintenanceData
}
type UseCoupon struct {
	Call string
	Data UseCouponDataInfo
}
type UseCouponDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Send_ppid   string
	Coupon      string
}
type UseCouponRetDataInfo struct {
	Ret string
}
type UseCouponReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        UseCouponRetDataInfo
	Maintenance MaintenanceData
}
type UseCredit struct {
	Call string
	Data UseCreditDataInfo
}
type UseCreditDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Credit      int64
}
type UseCreditRetDataInfo struct {
	U_gold   int64
	U_cp     int64
	U_mp     int64
	U_credit int64
}
type UseCreditReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        UseCreditRetDataInfo
	Maintenance MaintenanceData
}
type UserAchievement struct {
	I_id       int64
	I_Level    int64
	D_Quantity float64
	S_Quantity string
}
type UserAreaData struct {
	U_area_num          int16
	D_Candy             float64
	D_Like              float64
	I_UserFanCount      int64
	I_UserFanGrade      int16
	I_SelectedCostumeId int64
	I_SelectedMusicId   int64
	I_SelectedGuitarId  int64
	S_TutorialList      string
	S_Gp1               string
	S_Gp2               string
}
type UserAvatarList struct {
	Call string
	Data UserAvatarListDataInfo
}
type UserAvatarListDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
}
type UserAvatarListRetDataInfo struct {
	U_seq  int32
	Avatar []any
}
type UserAvatarListReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        UserAvatarListRetDataInfo
	Maintenance MaintenanceData
}
type UserBuff struct {
	I_id         int64
	I_Level      int64
	I_ActiveTime int64
	I_EndTime    int64
}
type UserCandyShop struct {
	I_id              int64
	I_CurrentBuyCount int64
	I_TotalBuyCount   int64
	L_LastBuyTick     float64
	Upd_day           int64
}
type UserCharacter struct {
	I_id         int64
	I_Level      int64
	I_BonusLevel int64
}
type UserContentsData struct {
	User_achievement             []any
	User_buff                    []any
	User_candy_shop              []any
	User_character               []any
	User_costume                 []any
	User_daily_mission           []any
	User_follower                []any
	User_music                   []any
	User_prop                    []any
	User_unit                    []any
	User_skill                   []any
	User_shop                    []any
	User_messenger               []any
	User_guitar                  []any
	User_event_point             []any
	User_subscribe_list          []any
	User_subscribe_pass_reward   []any
	User_ticketcollection        []any
	User_follower_quest          []any
	User_follower_profile_reward []any
	User_follower_profile        []any
	User_follower_giftitem       []any
	User_ad_list                 []any
	User_chthird_stage           []any
	User_tutorial                []any
	User_ad_level                []any
	User_select_reward           []any
}
type UserCostume struct {
	I_id         int64
	I_Level      int64
	I_BonusLevel int64
}
type UserDailyMission struct {
	I_id       int64
	I_Level    int64
	D_Quantity int64
	Upd_date   string
}
type UserData struct {
	U_seq int64
	U_mp  int64
}
type UserDel struct {
	Call        string
	Data        UserDelDataInfo
	Common_data ParamData
}
type UserDelDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
}
type UserDelRetDataInfo struct {
	Result string
}
type UserDelReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        UserDelRetDataInfo
	Maintenance MaintenanceData
}
type UserEventPoint struct {
	S_EventType  string
	I_DataID     int64
	I_Point      int64
	I_Step       int64
	I_ADViewTime int64
	I_Version    int32
}
type UserFollower struct {
	I_id         int64
	I_Level      int64
	I_BonusLevel int64
}
type UserFollowerQuest struct {
	I_id              int64
	I_CurrentID       int64
	I_CompleteID      int64
	D_ConditionValue1 float64
	D_ConditionValue2 float64
	D_ConditionValue3 float64
	I_RewardReceived1 int16
	I_RewardReceived2 int16
	I_RewardReceived3 int16
	I_isInfinity      int16
}
type UserGuitar struct {
	I_id         int64
	I_Level      int64
	I_BonusLevel int64
}
type UserInfo struct {
	U_like                float64
	U_fans                int64
	U_fans_grade          int16
	U_selected_costume_id int64
	U_selected_music_id   int64
}
type UserItemInvenList struct {
	Item_list []any
}
type UserItemList struct {
	Item_type int32
	Item_id   int32
	Count     int64
}
type UserJoin struct {
	Call        string
	Data        UserJoinDataInfo
	Common_data ParamData
}
type UserJoinDataInfo struct {
	Os           int16
	Join_type    int16
	Uuid         string
	Device_uuid  string
	Device_token string
	Sns_id       string
	Access_token string
	Country      string
	State        string
}
type UserJoinRetDataInfo struct {
	U_seq int32
	U_id  string
}
type UserJoinReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        UserJoinRetDataInfo
	Maintenance MaintenanceData
}
type UserLevel struct {
	U_girl_level     int64
	U_skill_level    int64
	U_mate_level     int64
	U_follower_level int64
	U_play_level     int64
}
type UserLoad struct {
	Call        string
	Data        UserLoadDataInfo
	Common_data ParamData
}
type UserLoadDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Type        string
}
type UserLoadRetDataInfo struct {
	User          UserData
	Area_data     map[any]any
	User_contents UserContentsData
}
type UserLoadReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        UserLoadRetDataInfo
	Maintenance MaintenanceData
}
type UserLogin struct {
	Call        string
	Data        UserLoginDataInfo
	Common_data ParamData
}
type UserLoginDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	U_state     string
}
type UserLoginRetDataInfo struct {
	User          UserData
	Area_data     map[any]any
	User_contents UserContentsData
}
type UserLoginReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        UserLoginRetDataInfo
	Maintenance MaintenanceData
}
type UserMessenger struct {
	I_MessengerChatRoomId int64
	I_LastConfirmIndex    int64
	S_UnlockGroupList     string
	L_UpdateTimeTicks     int64
}
type UserMusic struct {
	I_id                      int64
	I_Level                   int64
	I_BonusLevel              int64
	B_EncoreBonusAppear       int64
	L_EncoreBonusActivateTime int64
	I_EncoreBonusFollowerId   int64
	I_ChThirdActiveTime       int64
}
type UserProp struct {
	I_id    int64
	I_Level int64
}
type UserPurchase struct {
	Call     string
	Data     UserPurchaseDataInfo
	Sub_mode string
}
type UserPurchaseDataInfo struct {
	U_seq             int32
	U_id              string
	Uuid              string
	Device_uuid       string
	Developer_payload string
	Pay_id            string
	Purchase_token    string
	Purchase_id       string
	Os                int16
	Restore           int16
}
type UserPurchaseErrorLog struct {
	Call string
	Data UserPurchaseErrorLogDataInfo
}
type UserPurchaseErrorLogDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
	Error_log   string
}
type UserPurchaseErrorLogRetDataInfo struct {
	Ret string
}
type UserPurchaseErrorLogReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        UserPurchaseErrorLogRetDataInfo
	Maintenance MaintenanceData
}
type UserPurchaseMusicIdx struct {
	Music_idx int32
}
type UserPurchaseRetDataInfo struct {
	Pay_id      string
	Product_id  string
	Purchase_id string
	Music_data  []any
}
type UserPurchaseReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        UserPurchaseRetDataInfo
	Maintenance MaintenanceData
}
type UserSave struct {
	Call        string
	Data        UserSaveDataInfo
	Common_data ParamData
}
type UserSaveData struct {
	U_candy          int64
	U_like           int64
	U_fans           int64
	U_girl_level     int64
	U_skill_level    int64
	U_mate_level     int64
	U_follower_level int64
	U_play_level     int64
}
type UserSaveDataInfo struct {
	U_seq               int32
	U_id                string
	Uuid                string
	Device_uuid         string
	Type                string
	User_info           []any
	User_area_info      []any
	User_achievement    []any
	User_character      []any
	User_costume        []any
	User_daily_mission  []any
	User_follower       []any
	User_music          []any
	User_skill          []any
	User_messenger      []any
	User_guitar         []any
	User_event_point    []any
	User_follower_quest []any
}
type UserSaveRetDataInfo struct {
	Status string
}
type UserSaveReturn struct {
	Error       ErrorRetCode
	Server_time ServerTimeRet
	Mode        string
	Call        string
	Data        UserSaveRetDataInfo
	Maintenance MaintenanceData
}
type UserScore struct {
	Music_idx          int32
	Score1             int64
	Score2             int64
	Score3             int64
	Score4             int64
	Grade1             int16
	Grade2             int16
	Grade3             int16
	Grade4             int16
	Play_cnt1          int64
	Play_cnt2          int64
	Play_cnt3          int64
	Play_cnt4          int64
	Combo_grade1       int16
	Combo_grade2       int16
	Combo_grade3       int16
	Combo_grade4       int16
	Multi_score1       int64
	Multi_score2       int64
	Multi_score3       int64
	Multi_score4       int64
	Multi_play_cnt1    int64
	Multi_play_cnt2    int64
	Multi_play_cnt3    int64
	Multi_play_cnt4    int64
	Multi_combo_grade1 int16
	Multi_combo_grade2 int16
	Multi_combo_grade3 int16
	Multi_combo_grade4 int16
	Purchase_type      int16
	Is_extra           int16
	Review_point       int16
	Rank_list          GetTotalMusicRankRetDataInfo
}
type UserShop struct {
	I_id           int64
	I_Count        int64
	I_TotalCount   int64
	I_PurchaseTime int64
	Upd_day        int64
}
type UserSkill struct {
	I_id               int64
	I_Level            int64
	B_Activate         int16
	L_ActivateOnTicks  int64
	L_ActivateOffTicks int64
}
type UserSubscribeList struct {
	I_SubscribeID int64
	I_ActiveTime  int64
	I_isActive    int64
}
type UserSubscribePassReward struct {
	I_SubscribeID int64
	I_Type        int64
	I_Step        int64
	I_UpdateTime  int64
	I_Version     int32
}
type UserTicketCollection struct {
	I_id int64
}
type UserTitleList struct {
	Call string
	Data UserTitleListDataInfo
}
type UserTitleListDataInfo struct {
	U_seq       int32
	U_id        string
	Uuid        string
	Device_uuid string
}
type UserTitleListRetDataInfo struct {
	U_seq int32
	Title []any
}
type UserTitleListReturn struct {
	Error       ErrorRetCode
	Mode        string
	Call        string
	Data        UserTitleListRetDataInfo
	Maintenance MaintenanceData
}
type UserUnit struct {
	I_id    int64
	I_Level int64
}
