package model

type Error struct {
	Is_success      bool   `thrift:",1,omitempty"`
	Err_code        int32  `thrift:",2,omitempty"`
	Err_message     string `thrift:",3,omitempty"`
	Is_debug        bool   `thrift:",4,omitempty"`
	Is_achieve_noti bool   `thrift:",5,omitempty"`
	Mission_notis   []any  `thrift:",6,omitempty"` // TODO
}

type MoreGamesLog struct {
	Call string               `thrift:",1,omitempty"`
	Data MoreGamesLogDataInfo `thrift:",2,omitempty"`
}

type MoreGamesLogDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	App_id      int32  `thrift:",5,omitempty"`
}

type MoreGamesLogRetDataInfo struct {
	Ret string `thrift:",1,omitempty"`
}

type MoreGamesLogReturn struct {
	Error       ErrorRetCode            `thrift:",1,omitempty"`
	Server_time ServerTimeRet           `thrift:",2,omitempty"`
	Mode        string                  `thrift:",3,omitempty"`
	Call        string                  `thrift:",4,omitempty"`
	Data        MoreGamesLogRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData         `thrift:",6,omitempty"`
}

type Rect struct {
	Left   float64 `thrift:",1,omitempty"`
	Top    float64 `thrift:",2,omitempty"`
	Width  float64 `thrift:",3,omitempty"`
	Height float64 `thrift:",4,omitempty"`
}

type Request struct {
	Call        string       `thrift:",1,omitempty"`
	Data        InitDataInfo `thrift:",2,omitempty"`
	Common_data ParamData    `thrift:",3,omitempty"`
}

type StageFollowerProfileScore struct {
	I_id        int32 `thrift:",1,omitempty"`
	Score       int32 `thrift:",2,omitempty"`
	Bonus_score int32 `thrift:",3,omitempty"`
}

type UserAdLevel struct {
	I_id    int32 `thrift:",1,omitempty"`
	I_Level int32 `thrift:",2,omitempty"`
	I_EXP   int32 `thrift:",3,omitempty"`
}

type UserAdList struct {
	I_id           int32 `thrift:",1,omitempty"`
	I_Count        int32 `thrift:",2,omitempty"`
	I_TotalCount   int32 `thrift:",3,omitempty"`
	I_LastViewTime int32 `thrift:",4,omitempty"`
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

type UserChThirdStage struct {
	I_id         int32 `thrift:",1,omitempty"`
	I_ChapterId  int32 `thrift:",2,omitempty"`
	I_StageIndex int32 `thrift:",3,omitempty"`
	I_Star       int32 `thrift:",4,omitempty"`
}

type UserFollowerGiftItem struct {
	I_id    int32 `thrift:",1,omitempty"`
	I_Value int32 `thrift:",2,omitempty"`
}

type UserFollowerProfile struct {
	I_id       int32 `thrift:",1,omitempty"`
	I_Level    int32 `thrift:",2,omitempty"`
	D_Exp      int64 `thrift:",3,omitempty"`
	I_AddCandy int32 `thrift:",4,omitempty"`
}

type UserFollowerProfileReward struct {
	I_id          int32 `thrift:",1,omitempty"`
	I_RewardLevel int32 `thrift:",2,omitempty"`
}

type UserSelectReward struct {
	I_GroupId          int32 `thrift:",1,omitempty"`
	I_RewardGroupId    int32 `thrift:",2,omitempty"`
	I_AltRewardGroupId int32 `thrift:",3,omitempty"`
}

type UserTutorial struct {
	I_id int32 `thrift:",1,omitempty"`
}

type Vector3 struct {
	X float64 `thrift:",1,omitempty"`
	Y float64 `thrift:",2,omitempty"`
	Z float64 `thrift:",3,omitempty"`
}

type AdViewLog struct {
	Call        string            `thrift:",1,omitempty"`
	Data        AdViewLogDataInfo `thrift:",2,omitempty"`
	Common_data ParamData         `thrift:",3,omitempty"`
}

type AdViewLogDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Ad_type     string `thrift:",5,omitempty"`
	Ad_name     string `thrift:",6,omitempty"`
}

type AdViewLogRetDataInfo struct {
	Status string `thrift:",1,omitempty"`
}

type AdViewLogReturn struct {
	Error       ErrorRetCode         `thrift:",1,omitempty"`
	Server_time ServerTimeRet        `thrift:",2,omitempty"`
	Mode        string               `thrift:",3,omitempty"`
	Call        string               `thrift:",4,omitempty"`
	Data        AdViewLogRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData      `thrift:",6,omitempty"`
}

type AlbumListInfo struct {
	Album_idx  int64       `thrift:",1,omitempty"`
	Country    string      `thrift:",2,omitempty"`
	Date_issue string      `thrift:",3,omitempty"`
	Cdn_dir    string      `thrift:",4,omitempty"`
	Album_img  string      `thrift:",5,omitempty"`
	Language   map[any]any `thrift:",6,omitempty"` // TODO
}

type AlbumListLanguage struct {
	Title        string `thrift:",1,omitempty"`
	Company      string `thrift:",2,omitempty"`
	Album_artist string `thrift:",3,omitempty"`
}

type AvatarList struct {
	Call string             `thrift:",1,omitempty"`
	Data AvatarListDataInfo `thrift:",2,omitempty"`
}

type AvatarListDataInfo struct {
	Type string `thrift:",1,omitempty"`
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

type AvatarListReturn struct {
	Error       ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string          `thrift:",2,omitempty"`
	Call        string          `thrift:",3,omitempty"`
	Data        []any           `thrift:",4,omitempty"` // TODO
	Maintenance MaintenanceData `thrift:",5,omitempty"`
}

type Bookmark_list struct {
	Music_idx int32 `thrift:",1,omitempty"`
	Flag      int16 `thrift:",2,omitempty"`
}

type BuyAvatar struct {
	Call string            `thrift:",1,omitempty"`
	Data BuyAvatarDataInfo `thrift:",2,omitempty"`
}

type BuyAvatarDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Idx         int16  `thrift:",5,omitempty"`
}

type BuyAvatarRetDataInfo struct {
	Avatar_idx int16 `thrift:",1,omitempty"`
	U_gold     int64 `thrift:",2,omitempty"`
	U_cp       int64 `thrift:",3,omitempty"`
	U_mp       int64 `thrift:",4,omitempty"`
}

type BuyAvatarReturn struct {
	Error       ErrorRetCode         `thrift:",1,omitempty"`
	Mode        string               `thrift:",2,omitempty"`
	Call        string               `thrift:",3,omitempty"`
	Data        BuyAvatarRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData      `thrift:",5,omitempty"`
}

type BuyCheck struct {
	Call string           `thrift:",1,omitempty"`
	Data BuyCheckDataInfo `thrift:",2,omitempty"`
}

type BuyCheckDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Idx         int32  `thrift:",5,omitempty"`
}

type BuyCheckRetDataInfo struct {
	Result int64 `thrift:",1,omitempty"`
}

type BuyCheckReturn struct {
	Error       ErrorRetCode        `thrift:",1,omitempty"`
	Server_time ServerTimeRet       `thrift:",2,omitempty"`
	Mode        string              `thrift:",3,omitempty"`
	Call        string              `thrift:",4,omitempty"`
	Data        BuyCheckRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData     `thrift:",6,omitempty"`
}

type BuyCollectionItem struct {
	Call string                    `thrift:",1,omitempty"`
	Data BuyCollectionItemDataInfo `thrift:",2,omitempty"`
}

type BuyCollectionItemDataInfo struct {
	U_seq             int32  `thrift:",1,omitempty"`
	U_id              string `thrift:",2,omitempty"`
	Uuid              string `thrift:",3,omitempty"`
	Device_uuid       string `thrift:",4,omitempty"`
	Special_music_idx int32  `thrift:",5,omitempty"`
	Idx               int32  `thrift:",6,omitempty"`
}

type BuyCollectionItemRetDataInfo struct {
	Item_type    int16  `thrift:",1,omitempty"`
	Item_value   int64  `thrift:",2,omitempty"`
	Buy_datetime string `thrift:",3,omitempty"`
}

type BuyCollectionItemReturn struct {
	Error       ErrorRetCode                 `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        BuyCollectionItemRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData              `thrift:",5,omitempty"`
}

type BuyCollectionMusic struct {
	Call string                     `thrift:",1,omitempty"`
	Data BuyCollectionMusicDataInfo `thrift:",2,omitempty"`
}

type BuyCollectionMusicDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Idx         int32  `thrift:",5,omitempty"`
}

type BuyCollectionMusicRetDataInfo struct {
	Music_idx    int64  `thrift:",1,omitempty"`
	Buy_datetime string `thrift:",2,omitempty"`
}

type BuyCollectionMusicReturn struct {
	Error       ErrorRetCode                  `thrift:",1,omitempty"`
	Mode        string                        `thrift:",2,omitempty"`
	Call        string                        `thrift:",3,omitempty"`
	Data        BuyCollectionMusicRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData               `thrift:",5,omitempty"`
}

type BuyContents struct {
	Call        string              `thrift:",1,omitempty"`
	Data        BuyContentsDataInfo `thrift:",2,omitempty"`
	Sub_mode    string              `thrift:",3,omitempty"`
	Common_data ParamData           `thrift:",4,omitempty"`
}

type BuyContentsDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Type        string `thrift:",5,omitempty"`
	Idx         int32  `thrift:",6,omitempty"`
}

type BuyContentsRetDataInfo struct {
	U_cp       int64     `thrift:",1,omitempty"`
	U_candy    float64   `thrift:",2,omitempty"`
	User_skill UserSkill `thrift:",3,omitempty"`
	User_unit  UserUnit  `thrift:",4,omitempty"`
}

type BuyContentsReturn struct {
	Error       ErrorRetCode           `thrift:",1,omitempty"`
	Server_time ServerTimeRet          `thrift:",2,omitempty"`
	Mode        string                 `thrift:",3,omitempty"`
	Call        string                 `thrift:",4,omitempty"`
	Data        BuyContentsRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData        `thrift:",6,omitempty"`
}

type BuyExtraMode struct {
	Call string               `thrift:",1,omitempty"`
	Data BuyExtraModeDataInfo `thrift:",2,omitempty"`
}

type BuyExtraModeDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Music_idx   int32  `thrift:",5,omitempty"`
	Key_flg     int64  `thrift:",6,omitempty"`
}

type BuyExtraModeRetDataInfo struct {
	Music_idx int32 `thrift:",1,omitempty"`
}

type BuyExtraModeReturn struct {
	Error       ErrorRetCode            `thrift:",1,omitempty"`
	Mode        string                  `thrift:",2,omitempty"`
	Call        string                  `thrift:",3,omitempty"`
	Data        BuyExtraModeRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData         `thrift:",5,omitempty"`
}

type BuyItem struct {
	Call string          `thrift:",1,omitempty"`
	Data BuyItemDataInfo `thrift:",2,omitempty"`
}

type BuyItemDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Idx         int16  `thrift:",5,omitempty"`
}

type BuyItemRetDataInfo struct {
	Idx      int32 `thrift:",1,omitempty"`
	U_gold   int64 `thrift:",2,omitempty"`
	U_cp     int64 `thrift:",3,omitempty"`
	U_mp     int64 `thrift:",4,omitempty"`
	U_credit int64 `thrift:",5,omitempty"`
}

type BuyItemReturn struct {
	Error       ErrorRetCode       `thrift:",1,omitempty"`
	Mode        string             `thrift:",2,omitempty"`
	Call        string             `thrift:",3,omitempty"`
	Data        BuyItemRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData    `thrift:",5,omitempty"`
}

type BuyMusic struct {
	Call string           `thrift:",1,omitempty"`
	Data BuyMusicDataInfo `thrift:",2,omitempty"`
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

type BuyMusicRetDataInfo struct {
	U_gold int32       `thrift:",1,omitempty"`
	U_cp   int32       `thrift:",2,omitempty"`
	U_mp   int32       `thrift:",3,omitempty"`
	Music  map[any]any `thrift:",4,omitempty"` // TODO
}

type BuyMusicReturn struct {
	Error       ErrorRetCode        `thrift:",1,omitempty"`
	Server_time ServerTimeRet       `thrift:",2,omitempty"`
	Mode        string              `thrift:",3,omitempty"`
	Call        string              `thrift:",4,omitempty"`
	Data        BuyMusicRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData     `thrift:",6,omitempty"`
}

type BuyPackage struct {
	Call string             `thrift:",1,omitempty"`
	Data BuyPackageDataInfo `thrift:",2,omitempty"`
}

type BuyPackageDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Store_idx   int64  `thrift:",5,omitempty"`
	Product_id  int64  `thrift:",6,omitempty"`
}

type BuyPackageMusic struct {
	Music_idx int64 `thrift:",1,omitempty"`
}

type BuyPackageRetDataInfo struct {
	U_gold     int64 `thrift:",1,omitempty"`
	U_cp       int64 `thrift:",2,omitempty"`
	U_mp       int64 `thrift:",3,omitempty"`
	U_credit   int64 `thrift:",4,omitempty"`
	Music_data []any `thrift:",5,omitempty"` // TODO
}

type BuyPackageReturn struct {
	Error       ErrorRetCode          `thrift:",1,omitempty"`
	Mode        string                `thrift:",2,omitempty"`
	Call        string                `thrift:",3,omitempty"`
	Data        BuyPackageRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData       `thrift:",5,omitempty"`
}

type BuyPiece struct {
	Call string           `thrift:",1,omitempty"`
	Data BuyPieceDataInfo `thrift:",2,omitempty"`
}

type BuyPieceDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Idx         int32  `thrift:",5,omitempty"`
	Cnt         int64  `thrift:",6,omitempty"`
}

type BuyPieceRetDataInfo struct {
	Idx    int32 `thrift:",1,omitempty"`
	U_gold int64 `thrift:",2,omitempty"`
	U_cp   int32 `thrift:",3,omitempty"`
	U_mp   int32 `thrift:",4,omitempty"`
}

type BuyPieceReturn struct {
	Error       ErrorRetCode        `thrift:",1,omitempty"`
	Mode        string              `thrift:",2,omitempty"`
	Call        string              `thrift:",3,omitempty"`
	Data        BuyPieceRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData     `thrift:",5,omitempty"`
}

type BuyShop struct {
	Call        string          `thrift:",1,omitempty"`
	Data        BuyShopDataInfo `thrift:",2,omitempty"`
	Sub_mode    string          `thrift:",3,omitempty"`
	Common_data ParamData       `thrift:",4,omitempty"`
}

type BuyShopDataInfo struct {
	U_seq             int32  `thrift:",1,omitempty"`
	U_id              string `thrift:",2,omitempty"`
	Uuid              string `thrift:",3,omitempty"`
	Device_uuid       string `thrift:",4,omitempty"`
	Idx               int32  `thrift:",5,omitempty"`
	Developer_payload string `thrift:",6,omitempty"`
	Pay_id            string `thrift:",7,omitempty"`
	Purchase_token    string `thrift:",8,omitempty"`
	Purchase_id       string `thrift:",9,omitempty"`
	Os                int16  `thrift:",10,omitempty"`
	App_id            string `thrift:",11,omitempty"`
	Member_id         string `thrift:",12,omitempty"`
}

type BuyShopRetDataInfo struct {
	U_cp          int64       `thrift:",1,omitempty"`
	U_candy       float64     `thrift:",2,omitempty"`
	Reward_data   []any       `thrift:",3,omitempty"` // TODO
	User_ad_level UserAdLevel `thrift:",4,omitempty"`
}

type BuyShopReturn struct {
	Error       ErrorRetCode       `thrift:",1,omitempty"`
	Server_time ServerTimeRet      `thrift:",2,omitempty"`
	Mode        string             `thrift:",3,omitempty"`
	Call        string             `thrift:",4,omitempty"`
	Data        BuyShopRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData    `thrift:",6,omitempty"`
}

type BuyUserTitle struct {
	Call string               `thrift:",1,omitempty"`
	Data BuyUserTitleDataInfo `thrift:",2,omitempty"`
}

type BuyUserTitleDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Idx         int16  `thrift:",5,omitempty"`
}

type BuyUserTitleRetDataInfo struct {
	Title_idx int16 `thrift:",1,omitempty"`
	U_gold    int64 `thrift:",2,omitempty"`
	U_cp      int64 `thrift:",3,omitempty"`
	U_mp      int64 `thrift:",4,omitempty"`
}

type BuyUserTitleReturn struct {
	Error       ErrorRetCode            `thrift:",1,omitempty"`
	Mode        string                  `thrift:",2,omitempty"`
	Call        string                  `thrift:",3,omitempty"`
	Data        BuyUserTitleRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData         `thrift:",5,omitempty"`
}

type BuyVarietyStore struct {
	Call        string                  `thrift:",1,omitempty"`
	Data        BuyVarietyStoreDataInfo `thrift:",2,omitempty"`
	Common_data ParamData               `thrift:",3,omitempty"`
}

type BuyVarietyStoreDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Idx         int32  `thrift:",5,omitempty"`
}

type BuyVarietyStoreRetDataInfo struct {
	U_cp         int64   `thrift:",1,omitempty"`
	U_candy      float64 `thrift:",2,omitempty"`
	Reward_type  int16   `thrift:",3,omitempty"`
	Reward_id    int32   `thrift:",4,omitempty"`
	Reward_value int64   `thrift:",5,omitempty"`
	Status       string  `thrift:",6,omitempty"`
}

type BuyVarietyStoreReturn struct {
	Error       ErrorRetCode               `thrift:",1,omitempty"`
	Server_time ServerTimeRet              `thrift:",2,omitempty"`
	Mode        string                     `thrift:",3,omitempty"`
	Call        string                     `thrift:",4,omitempty"`
	Data        BuyVarietyStoreRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData            `thrift:",6,omitempty"`
}

type ChThirdStage struct {
	Call        string               `thrift:",1,omitempty"`
	Data        ChThirdStageDataInfo `thrift:",2,omitempty"`
	Common_data ParamData            `thrift:",3,omitempty"`
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

type ChThirdStageReturn struct {
	Error       ErrorRetCode            `thrift:",1,omitempty"`
	Server_time ServerTimeRet           `thrift:",2,omitempty"`
	Mode        string                  `thrift:",3,omitempty"`
	Call        string                  `thrift:",4,omitempty"`
	Data        ChThirdStageRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData         `thrift:",6,omitempty"`
}

type ChThirdStreamStage struct {
	Call        string                     `thrift:",1,omitempty"`
	Data        ChThirdStreamStageDataInfo `thrift:",2,omitempty"`
	Common_data ParamData                  `thrift:",3,omitempty"`
}

type ChThirdStreamStageDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	I_id        int32  `thrift:",5,omitempty"`
	Count       int32  `thrift:",6,omitempty"`
}

type ChThirdStreamStageRetDataInfo struct {
	I_id                  int32               `thrift:",1,omitempty"`
	User_ap               UserApData          `thrift:",2,omitempty"`
	User_follower_profile UserFollowerProfile `thrift:",3,omitempty"`
	Reward_data           map[any]any         `thrift:",4,omitempty"` // TODO
}

type ChThirdStreamStageReturn struct {
	Error       ErrorRetCode                  `thrift:",1,omitempty"`
	Server_time ServerTimeRet                 `thrift:",2,omitempty"`
	Mode        string                        `thrift:",3,omitempty"`
	Call        string                        `thrift:",4,omitempty"`
	Data        ChThirdStreamStageRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData               `thrift:",6,omitempty"`
}

type ChallengeMissionList struct {
	Call string                       `thrift:",1,omitempty"`
	Data ChallengeMissionListDataInfo `thrift:",2,omitempty"`
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

type ChallengeMissionListInfo struct {
	Music_idx    int32 `thrift:",1,omitempty"`
	Difficulty   int16 `thrift:",2,omitempty"`
	Mission_list []any `thrift:",3,omitempty"` // TODO
}

type ChallengeMissionListRetDataInfo struct {
	Music_idx    int32       `thrift:",1,omitempty"`
	Difficulty   int16       `thrift:",2,omitempty"`
	Mission_list map[any]any `thrift:",3,omitempty"` // TODO
}

type ChallengeMissionListReturn struct {
	Error       ErrorRetCode                    `thrift:",1,omitempty"`
	Mode        string                          `thrift:",2,omitempty"`
	Call        string                          `thrift:",3,omitempty"`
	Data        ChallengeMissionListRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData                 `thrift:",5,omitempty"`
}

type ChangeUser struct {
	Call        string             `thrift:",1,omitempty"`
	Data        ChangeUserDataInfo `thrift:",2,omitempty"`
	Common_data ParamData          `thrift:",3,omitempty"`
}

type ChangeUserDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Change_uuid string `thrift:",5,omitempty"`
}

type ChangeUserRetDataInfo struct {
	Result string `thrift:",1,omitempty"`
}

type ChangeUserReturn struct {
	Error       ErrorRetCode          `thrift:",1,omitempty"`
	Mode        string                `thrift:",2,omitempty"`
	Call        string                `thrift:",3,omitempty"`
	Data        ChangeUserRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData       `thrift:",5,omitempty"`
}

type ChannelList struct {
	Channel        int16  `thrift:",1,omitempty"`
	Url            string `thrift:",2,omitempty"`
	Port           int64  `thrift:",3,omitempty"`
	Flg            int16  `thrift:",4,omitempty"`
	Maintenance_cn string `thrift:",5,omitempty"`
	Maintenance_en string `thrift:",6,omitempty"`
	Maintenance_jp string `thrift:",7,omitempty"`
}

type CheckBuyProduct struct {
	Call string                  `thrift:",1,omitempty"`
	Data CheckBuyProductDataInfo `thrift:",2,omitempty"`
}

type CheckBuyProductDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Product_id  string `thrift:",5,omitempty"`
}

type CheckBuyProductRetDataInfo struct {
	Is_owner int16 `thrift:",1,omitempty"`
}

type CheckBuyProductReturn struct {
	Error       ErrorRetCode               `thrift:",1,omitempty"`
	Mode        string                     `thrift:",2,omitempty"`
	Call        string                     `thrift:",3,omitempty"`
	Data        CheckBuyProductRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData            `thrift:",5,omitempty"`
}

type CheckBuyShop struct {
	Call        string               `thrift:",1,omitempty"`
	Data        CheckBuyShopDataInfo `thrift:",2,omitempty"`
	Common_data ParamData            `thrift:",3,omitempty"`
}

type CheckBuyShopDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Shop_id     string `thrift:",5,omitempty"`
}

type CheckBuyShopRetDataInfo struct {
	Is_owner int16 `thrift:",1,omitempty"`
}

type CheckBuyShopReturn struct {
	Error       ErrorRetCode            `thrift:",1,omitempty"`
	Mode        string                  `thrift:",2,omitempty"`
	Call        string                  `thrift:",3,omitempty"`
	Data        CheckBuyShopRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData         `thrift:",5,omitempty"`
}

type CheckPurchased struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        CheckPurchasedDataInfo `thrift:",2,omitempty"`
	Common_data ParamData              `thrift:",3,omitempty"`
}

type CheckPurchasedDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type CheckPurchasedRetDataInfo struct {
	Is_purchased int16 `thrift:",1,omitempty"`
}

type CheckPurchasedReturn struct {
	Error       ErrorRetCode              `thrift:",1,omitempty"`
	Mode        string                    `thrift:",2,omitempty"`
	Call        string                    `thrift:",3,omitempty"`
	Data        CheckPurchasedRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData           `thrift:",5,omitempty"`
}

type CheckUser struct {
	Call        string            `thrift:",1,omitempty"`
	Data        CheckUserDataInfo `thrift:",2,omitempty"`
	Common_data ParamData         `thrift:",3,omitempty"`
}

type CheckUserDataInfo struct {
	Uuid        string `thrift:",1,omitempty"`
	Device_uuid string `thrift:",2,omitempty"`
}

type CheckUserRetDataInfo struct {
	U_seq int32  `thrift:",1,omitempty"`
	U_id  string `thrift:",2,omitempty"`
}

type CheckUserReturn struct {
	Error       ErrorRetCode         `thrift:",1,omitempty"`
	Server_time ServerTimeRet        `thrift:",2,omitempty"`
	Mode        string               `thrift:",3,omitempty"`
	Call        string               `thrift:",4,omitempty"`
	Data        CheckUserRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData      `thrift:",6,omitempty"`
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

type CollectionRetData struct {
	Music_idx int64 `thrift:",1,omitempty"`
	Cnt       int64 `thrift:",2,omitempty"`
}

type CompleteLog struct {
	Call        string              `thrift:",1,omitempty"`
	Data        CompleteLogDataInfo `thrift:",2,omitempty"`
	Common_data ParamData           `thrift:",3,omitempty"`
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

type CompleteLogRetDataInfo struct {
	Status string `thrift:",1,omitempty"`
}

type CompleteLogReturn struct {
	Error       ErrorRetCode           `thrift:",1,omitempty"`
	Server_time ServerTimeRet          `thrift:",2,omitempty"`
	Mode        string                 `thrift:",3,omitempty"`
	Call        string                 `thrift:",4,omitempty"`
	Data        CompleteLogRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData        `thrift:",6,omitempty"`
}

type DefaultSettingDataList struct {
	Setting_key   string `thrift:",1,omitempty"`
	Setting_value string `thrift:",2,omitempty"`
}

type DefaultSettingList struct {
	Call        string                     `thrift:",1,omitempty"`
	Data        DefaultSettingListDataInfo `thrift:",2,omitempty"`
	Common_data ParamData                  `thrift:",3,omitempty"`
}

type DefaultSettingListDataInfo struct {
	Device_uuid string `thrift:",1,omitempty"`
	Os          int16  `thrift:",2,omitempty"`
}

type DefaultSettingListRetDataInfo struct {
	Status       string `thrift:",1,omitempty"`
	Setting_list []any  `thrift:",2,omitempty"` // TODO
}

type DefaultSettingListReturn struct {
	Error       ErrorRetCode                  `thrift:",1,omitempty"`
	Server_time ServerTimeRet                 `thrift:",2,omitempty"`
	Mode        string                        `thrift:",3,omitempty"`
	Call        string                        `thrift:",4,omitempty"`
	Data        DefaultSettingListRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData               `thrift:",6,omitempty"`
}

type DeletePost struct {
	Call        string             `thrift:",1,omitempty"`
	Data        DeletePostDataInfo `thrift:",2,omitempty"`
	Common_data ParamData          `thrift:",3,omitempty"`
}

type DeletePostDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Idx         int64  `thrift:",5,omitempty"`
	Type        int16  `thrift:",6,omitempty"`
}

type DeletePostRetDataInfo struct {
	Idx int64 `thrift:",1,omitempty"`
}

type DeletePostReturn struct {
	Error       ErrorRetCode          `thrift:",1,omitempty"`
	Mode        string                `thrift:",2,omitempty"`
	Call        string                `thrift:",3,omitempty"`
	Data        DeletePostRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData       `thrift:",5,omitempty"`
}

type DeleteUser struct {
	Call string             `thrift:",1,omitempty"`
	Data DeleteUserDataInfo `thrift:",2,omitempty"`
}

type DeleteUserDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type DeleteUserRetDataInfo struct {
	Result string `thrift:",1,omitempty"`
}

type DeleteUserReturn struct {
	Error       ErrorRetCode          `thrift:",1,omitempty"`
	Server_time ServerTimeRet         `thrift:",2,omitempty"`
	Mode        string                `thrift:",3,omitempty"`
	Call        string                `thrift:",4,omitempty"`
	Data        DeleteUserRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData       `thrift:",6,omitempty"`
}

type EndPlayMusic struct {
	Call string               `thrift:",1,omitempty"`
	Data EndPlayMusicDataInfo `thrift:",2,omitempty"`
}

type EndPlayMusicDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Score       int32  `thrift:",5,omitempty"`
	Medal       int16  `thrift:",6,omitempty"`
	Over_flg    int16  `thrift:",7,omitempty"`
}

type EndPlayMusicRetDataInfo struct {
	Code int16 `thrift:",1,omitempty"`
}

type EndPlayMusicReturn struct {
	Error       ErrorRetCode            `thrift:",1,omitempty"`
	Mode        string                  `thrift:",2,omitempty"`
	Call        string                  `thrift:",3,omitempty"`
	Data        EndPlayMusicRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData         `thrift:",5,omitempty"`
}

type ErrorRetCode struct {
	Code   int32  `thrift:",1,omitempty"`
	Errmsg string `thrift:",2,omitempty"`
}

type EventRankRewardListArr struct {
	Event_idx    int32  `thrift:",1,omitempty"`
	Title        string `thrift:",2,omitempty"`
	Desc_kr      string `thrift:",3,omitempty"`
	Desc_en      string `thrift:",4,omitempty"`
	Desc_jp      string `thrift:",5,omitempty"`
	Reward_type  int16  `thrift:",6,omitempty"`
	Reward_value int64  `thrift:",7,omitempty"`
}

type EventRankRewardListData struct {
	Title        string `thrift:",1,omitempty"`
	Desc_cn      string `thrift:",2,omitempty"`
	Desc_en      string `thrift:",3,omitempty"`
	Desc_jp      string `thrift:",4,omitempty"`
	Reward_type  int16  `thrift:",5,omitempty"`
	Reward_value int64  `thrift:",6,omitempty"`
}

type EventRewardList struct {
	Call string                  `thrift:",1,omitempty"`
	Data EventRewardListDataInfo `thrift:",2,omitempty"`
}

type EventRewardListArr struct {
	Event_idx    int32 `thrift:",1,omitempty"`
	Gain_point   int64 `thrift:",2,omitempty"`
	Reward_type  int16 `thrift:",3,omitempty"`
	Reward_value int64 `thrift:",4,omitempty"`
}

type EventRewardListData struct {
	Gain_point   int64 `thrift:",1,omitempty"`
	Reward_type  int16 `thrift:",2,omitempty"`
	Reward_value int64 `thrift:",3,omitempty"`
}

type EventRewardListDataInfo struct {
	Event_idx   int32  `thrift:",1,omitempty"`
	Device_uuid string `thrift:",2,omitempty"`
}

type EventRewardListRetDataInfo struct {
	Event_idx        int32 `thrift:",1,omitempty"`
	Reward_list      []any `thrift:",2,omitempty"` // TODO
	Rank_reward_list []any `thrift:",3,omitempty"` // TODO
}

type EventRewardListReturn struct {
	Error       ErrorRetCode               `thrift:",1,omitempty"`
	Mode        string                     `thrift:",2,omitempty"`
	Call        string                     `thrift:",3,omitempty"`
	Data        EventRewardListRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData            `thrift:",5,omitempty"`
}

type FacebookUserJoin struct {
	Call        string                   `thrift:",1,omitempty"`
	Data        FacebookUserJoinDataInfo `thrift:",2,omitempty"`
	Common_data ParamData                `thrift:",3,omitempty"`
}

type FacebookUserJoinDataInfo struct {
	Access_token string `thrift:",1,omitempty"`
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

type GetAdLevelData struct {
	I_GroupID     int32 `thrift:",1,omitempty"`
	I_Level       int32 `thrift:",2,omitempty"`
	I_RequireEXP  int32 `thrift:",3,omitempty"`
	I_RewardGroup int32 `thrift:",4,omitempty"`
	B_IsActive    int16 `thrift:",5,omitempty"`
	I_id          int32 `thrift:",6,omitempty"`
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

type GetAlbum struct {
	Call string           `thrift:",1,omitempty"`
	Data GetAlbumDataInfo `thrift:",2,omitempty"`
}

type GetAlbumDataInfo struct {
	Null string `thrift:",1,omitempty"`
}

type GetAlbumInfoLanguage struct {
	Seq          int32  `thrift:",1,omitempty"`
	Album_title  string `thrift:",2,omitempty"`
	Company      string `thrift:",3,omitempty"`
	Album_artist string `thrift:",4,omitempty"`
}

type GetAlbumInfoList struct {
	Call string               `thrift:",1,omitempty"`
	Data GetAlbumListDataInfo `thrift:",2,omitempty"`
}

type GetAlbumInfoListDataInfo struct {
}

type GetAlbumInfoListRetDataInfo struct {
	Language map[any]any `thrift:",1,omitempty"` // TODO
}

type GetAlbumInfoListReturn struct {
	Error       ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string          `thrift:",2,omitempty"`
	Call        string          `thrift:",3,omitempty"`
	Data        map[any]any     `thrift:",4,omitempty"` // TODO
	Maintenance MaintenanceData `thrift:",5,omitempty"`
}

type GetAlbumList struct {
	Call string               `thrift:",1,omitempty"`
	Data GetAlbumListDataInfo `thrift:",2,omitempty"`
}

type GetAlbumListDataInfo struct {
}

type GetAlbumListRetDataInfo struct {
	Album_idx  int32  `thrift:",1,omitempty"`
	Country    string `thrift:",2,omitempty"`
	Date_issue string `thrift:",3,omitempty"`
	Cdn_dir    string `thrift:",4,omitempty"`
	Album_img  string `thrift:",5,omitempty"`
}

type GetAlbumListReturn struct {
	Error       ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string          `thrift:",2,omitempty"`
	Call        string          `thrift:",3,omitempty"`
	Data        []any           `thrift:",4,omitempty"` // TODO
	Maintenance MaintenanceData `thrift:",5,omitempty"`
}

type GetAlbumRetDataInfo struct {
	Album     AlbumListInfo `thrift:",1,omitempty"`
	List_info []any         `thrift:",2,omitempty"` // TODO
}

type GetAlbumReturn struct {
	Error       ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string          `thrift:",2,omitempty"`
	Call        string          `thrift:",3,omitempty"`
	Data        []any           `thrift:",4,omitempty"` // TODO
	Maintenance MaintenanceData `thrift:",5,omitempty"`
}

type GetAllPoint struct {
	Call     string              `thrift:",1,omitempty"`
	Data     GetAllPointDataInfo `thrift:",2,omitempty"`
	Sub_mode string              `thrift:",3,omitempty"`
}

type GetAllPointDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type GetAllPointRetDataInfo struct {
	U_gold           int64 `thrift:",1,omitempty"`
	U_cp             int64 `thrift:",2,omitempty"`
	U_mp             int64 `thrift:",3,omitempty"`
	U_credit         int64 `thrift:",4,omitempty"`
	Next_credit_time int64 `thrift:",5,omitempty"`
	Max_credit_time  int64 `thrift:",6,omitempty"`
}

type GetAllPointReturn struct {
	Error       ErrorRetCode           `thrift:",1,omitempty"`
	Mode        string                 `thrift:",2,omitempty"`
	Call        string                 `thrift:",3,omitempty"`
	Data        GetAllPointRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData        `thrift:",5,omitempty"`
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

type GetCache struct {
	Call string           `thrift:",1,omitempty"`
	Data GetCacheDataInfo `thrift:",2,omitempty"`
}

type GetCacheDataInfo struct {
	Type int16 `thrift:",1,omitempty"`
}

type GetCacheRetDataInfo struct {
	Uptime int32 `thrift:",1,omitempty"`
}

type GetCacheReturn struct {
	Error       ErrorRetCode        `thrift:",1,omitempty"`
	Mode        string              `thrift:",2,omitempty"`
	Call        string              `thrift:",3,omitempty"`
	Data        GetCacheRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData     `thrift:",5,omitempty"`
}

type GetChThird struct {
	Call        string             `thrift:",1,omitempty"`
	Data        GetChThirdDataInfo `thrift:",2,omitempty"`
	Common_data ParamData          `thrift:",3,omitempty"`
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

type GetChThirdDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type GetChThirdRetDataInfo struct {
	U_seq                        int32      `thrift:",1,omitempty"`
	User_ap                      UserApData `thrift:",2,omitempty"`
	User_ch_third_stage          []any      `thrift:",3,omitempty"` // TODO
	User_ch_third_chapter_reward []any      `thrift:",4,omitempty"` // TODO
}

type GetChThirdReturn struct {
	Error       ErrorRetCode          `thrift:",1,omitempty"`
	Server_time ServerTimeRet         `thrift:",2,omitempty"`
	Mode        string                `thrift:",3,omitempty"`
	Call        string                `thrift:",4,omitempty"`
	Data        GetChThirdRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData       `thrift:",6,omitempty"`
}

type GetChThirdScoreData struct {
	I_id             int32 `thrift:",1,omitempty"`
	I_Level          int32 `thrift:",2,omitempty"`
	I_CharacterScore int32 `thrift:",3,omitempty"`
	I_FollowerScore  int32 `thrift:",4,omitempty"`
	I_MusicScore     int32 `thrift:",5,omitempty"`
	B_IsActive       int16 `thrift:",6,omitempty"`
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

type GetChThirdStarReward struct {
	Call        string                       `thrift:",1,omitempty"`
	Data        GetChThirdStarRewardDataInfo `thrift:",2,omitempty"`
	Common_data ParamData                    `thrift:",3,omitempty"`
}

type GetChThirdStarRewardDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	I_id        int32  `thrift:",5,omitempty"`
	Reward_num  int32  `thrift:",6,omitempty"`
}

type GetChThirdStarRewardRetDataInfo struct {
	I_id        int32      `thrift:",1,omitempty"`
	Reward_num  int32      `thrift:",2,omitempty"`
	User_ap     UserApData `thrift:",3,omitempty"`
	Reward_data []any      `thrift:",4,omitempty"` // TODO
}

type GetChThirdStarRewardReturn struct {
	Error       ErrorRetCode                    `thrift:",1,omitempty"`
	Server_time ServerTimeRet                   `thrift:",2,omitempty"`
	Mode        string                          `thrift:",3,omitempty"`
	Call        string                          `thrift:",4,omitempty"`
	Data        GetChThirdStarRewardRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData                 `thrift:",6,omitempty"`
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

type GetChoiceUser struct {
	Call        string                `thrift:",1,omitempty"`
	Data        GetChoiceUserDataInfo `thrift:",2,omitempty"`
	Common_data ParamData             `thrift:",3,omitempty"`
}

type GetChoiceUserDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Choice_uuid string `thrift:",5,omitempty"`
}

type GetChoiceUserRetDataInfo struct {
	User        ChoiceUserData `thrift:",1,omitempty"`
	Choice_user ChoiceUserData `thrift:",2,omitempty"`
}

type GetChoiceUserReturn struct {
	Error       ErrorRetCode             `thrift:",1,omitempty"`
	Mode        string                   `thrift:",2,omitempty"`
	Call        string                   `thrift:",3,omitempty"`
	Data        GetChoiceUserRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData          `thrift:",5,omitempty"`
}

type GetCollection struct {
	Call string                `thrift:",1,omitempty"`
	Data GetCollectionDataInfo `thrift:",2,omitempty"`
}

type GetCollectionDataInfo struct {
}

type GetCollectionItem struct {
	Special_music_idx int64 `thrift:",1,omitempty"`
	Idx               int64 `thrift:",2,omitempty"`
	Item_type         int16 `thrift:",3,omitempty"`
	Item_value        int64 `thrift:",4,omitempty"`
	Buy_cnt           int64 `thrift:",5,omitempty"`
}

type GetCollectionRetDataInfo struct {
	Idx      int32 `thrift:",1,omitempty"`
	Buy_type int16 `thrift:",2,omitempty"`
	Price    int64 `thrift:",3,omitempty"`
	Item     []any `thrift:",4,omitempty"` // TODO
}

type GetCollectionReturn struct {
	Error       ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string          `thrift:",2,omitempty"`
	Call        string          `thrift:",3,omitempty"`
	Data        []any           `thrift:",4,omitempty"` // TODO
	Maintenance MaintenanceData `thrift:",5,omitempty"`
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

type GetCountryMyRank struct {
	Call string                   `thrift:",1,omitempty"`
	Data GetCountryMyRankDataInfo `thrift:",2,omitempty"`
}

type GetCountryMyRankDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Country     string `thrift:",5,omitempty"`
	Start_limit int32  `thrift:",6,omitempty"`
	End_limit   int32  `thrift:",7,omitempty"`
}

type GetCountryMyRankRet struct {
	G           int16  `thrift:",1,omitempty"`
	S           int16  `thrift:",2,omitempty"`
	B           int16  `thrift:",3,omitempty"`
	Total_score int32  `thrift:",4,omitempty"`
	U_avatar    int16  `thrift:",5,omitempty"`
	U_id        string `thrift:",6,omitempty"`
	U_nick      string `thrift:",7,omitempty"`
	U_seq       int32  `thrift:",8,omitempty"`
	U_country   string `thrift:",9,omitempty"`
	Rank        int32  `thrift:",10,omitempty"`
}

type GetCountryMyRankRetDataInfo struct {
	Select_country string              `thrift:",1,omitempty"`
	Start_limit    int16               `thrift:",2,omitempty"`
	End_limit      int16               `thrift:",3,omitempty"`
	My_rank        GetCountryMyRankRet `thrift:",4,omitempty"`
	Rank_list      []any               `thrift:",5,omitempty"` // TODO
}

type GetCountryMyRankReturn struct {
	Error       ErrorRetCode                `thrift:",1,omitempty"`
	Mode        string                      `thrift:",2,omitempty"`
	Call        string                      `thrift:",3,omitempty"`
	Data        GetCountryMyRankRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData             `thrift:",5,omitempty"`
}

type GetCountryRank struct {
	Call string                 `thrift:",1,omitempty"`
	Data GetCountryRankDataInfo `thrift:",2,omitempty"`
}

type GetCountryRankDataInfo struct {
}

type GetCountryRankRetDataInfo struct {
	Rank        int32  `thrift:",1,omitempty"`
	Country     string `thrift:",2,omitempty"`
	Gold        int32  `thrift:",3,omitempty"`
	Silver      int32  `thrift:",4,omitempty"`
	Bronze      int32  `thrift:",5,omitempty"`
	Total_score int32  `thrift:",6,omitempty"`
}

type GetCountryRankReturn struct {
	Error       ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string          `thrift:",2,omitempty"`
	Call        string          `thrift:",3,omitempty"`
	Data        []any           `thrift:",4,omitempty"` // TODO
	Maintenance MaintenanceData `thrift:",5,omitempty"`
}

type GetCountryUserRank struct {
	G           int16  `thrift:",1,omitempty"`
	S           int16  `thrift:",2,omitempty"`
	B           int16  `thrift:",3,omitempty"`
	Total_score int32  `thrift:",4,omitempty"`
	U_avatar    int16  `thrift:",5,omitempty"`
	U_nick      string `thrift:",6,omitempty"`
	U_seq       int32  `thrift:",7,omitempty"`
}

type GetCountryUserRankRet struct {
	G           int16  `thrift:",1,omitempty"`
	S           int16  `thrift:",2,omitempty"`
	B           int16  `thrift:",3,omitempty"`
	Total_score int32  `thrift:",4,omitempty"`
	U_avatar    int16  `thrift:",5,omitempty"`
	U_id        string `thrift:",6,omitempty"`
	U_nick      string `thrift:",7,omitempty"`
	U_seq       int32  `thrift:",8,omitempty"`
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

type GetDiamondBonus struct {
	Call string                  `thrift:",1,omitempty"`
	Data GetDiamondBonusDataInfo `thrift:",2,omitempty"`
}

type GetDiamondBonusDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type GetDiamondBonusRetDataInfo struct {
	Bonus1 int64 `thrift:",1,omitempty"`
	Bonus2 int64 `thrift:",2,omitempty"`
	Bonus3 int64 `thrift:",3,omitempty"`
	Bonus4 int64 `thrift:",4,omitempty"`
	Bonus5 int64 `thrift:",5,omitempty"`
	Bonus6 int64 `thrift:",6,omitempty"`
}

type GetDiamondBonusReturn struct {
	Error       ErrorRetCode               `thrift:",1,omitempty"`
	Mode        string                     `thrift:",2,omitempty"`
	Call        string                     `thrift:",3,omitempty"`
	Data        GetDiamondBonusRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData            `thrift:",5,omitempty"`
}

type GetEventListData struct {
	Idx        int32  `thrift:",1,omitempty"`
	Event_type string `thrift:",2,omitempty"`
	Start_date string `thrift:",3,omitempty"`
	End_date   string `thrift:",4,omitempty"`
	B_IsActive int16  `thrift:",5,omitempty"`
}

type GetEventMain struct {
	Call string               `thrift:",1,omitempty"`
	Data GetEventMainDataInfo `thrift:",2,omitempty"`
}

type GetEventMainDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type GetEventMainRetDataInfo struct {
	Idx         int32  `thrift:",1,omitempty"`
	Title       string `thrift:",2,omitempty"`
	Img1        string `thrift:",3,omitempty"`
	Img2        string `thrift:",4,omitempty"`
	Movie       string `thrift:",5,omitempty"`
	Utc         int64  `thrift:",6,omitempty"`
	Banner      string `thrift:",7,omitempty"`
	Desc_cn     string `thrift:",8,omitempty"`
	Desc_en     string `thrift:",9,omitempty"`
	Desc_jp     string `thrift:",10,omitempty"`
	Server_time int64  `thrift:",11,omitempty"`
	Start_time  int64  `thrift:",12,omitempty"`
	End_time    int64  `thrift:",13,omitempty"`
	My_score    int64  `thrift:",14,omitempty"`
	My_rank     int64  `thrift:",15,omitempty"`
}

type GetEventMainReturn struct {
	Error       ErrorRetCode            `thrift:",1,omitempty"`
	Mode        string                  `thrift:",2,omitempty"`
	Call        string                  `thrift:",3,omitempty"`
	Data        GetEventMainRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData         `thrift:",5,omitempty"`
}

type GetEventRank struct {
	Call string               `thrift:",1,omitempty"`
	Data GetEventRankDataInfo `thrift:",2,omitempty"`
}

type GetEventRankDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Event_idx   int32  `thrift:",5,omitempty"`
	Start_rank  int64  `thrift:",6,omitempty"`
	End_rank    int64  `thrift:",7,omitempty"`
}

type GetEventRankRetDataInfo struct {
	U_seq     int64  `thrift:",1,omitempty"`
	U_nick    string `thrift:",2,omitempty"`
	U_id      string `thrift:",3,omitempty"`
	U_avatar  int32  `thrift:",4,omitempty"`
	U_country string `thrift:",5,omitempty"`
	U_title   int32  `thrift:",6,omitempty"`
	Rank      int64  `thrift:",7,omitempty"`
	Score     int64  `thrift:",8,omitempty"`
}

type GetEventRankReturn struct {
	Error       ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string          `thrift:",2,omitempty"`
	Call        string          `thrift:",3,omitempty"`
	Data        []any           `thrift:",4,omitempty"` // TODO
	Maintenance MaintenanceData `thrift:",5,omitempty"`
}

type GetEventRewardList struct {
	Call        string                     `thrift:",1,omitempty"`
	Data        GetEventRewardListDataInfo `thrift:",2,omitempty"`
	Sub_mode    string                     `thrift:",3,omitempty"`
	Common_data ParamData                  `thrift:",4,omitempty"`
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

type GetEventRewardListDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Event_idx   int32  `thrift:",5,omitempty"`
}

type GetEventRewardListRetDataInfo struct {
	Reward_list []any `thrift:",1,omitempty"` // TODO
	Group_idx   int32 `thrift:",2,omitempty"`
}

type GetEventRewardListReturn struct {
	Error       ErrorRetCode    `thrift:",1,omitempty"`
	Server_time ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string          `thrift:",3,omitempty"`
	Call        string          `thrift:",4,omitempty"`
	Data        map[any]any     `thrift:",5,omitempty"` // TODO
	Maintenance MaintenanceData `thrift:",6,omitempty"`
}

type GetFanData struct {
	I_Area        int64 `thrift:",1,omitempty"`
	I_Grade       int64 `thrift:",2,omitempty"`
	I_FanCount    int64 `thrift:",3,omitempty"`
	I_BonusRate   int64 `thrift:",4,omitempty"`
	I_BonusRateUI int64 `thrift:",5,omitempty"`
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

type GetGameDataList struct {
	Call        string                  `thrift:",1,omitempty"`
	Data        GetGameDataListDataInfo `thrift:",2,omitempty"`
	Sub_mode    string                  `thrift:",3,omitempty"`
	Common_data ParamData               `thrift:",4,omitempty"`
}

type GetGameDataListDataInfo struct {
	Game_type   string `thrift:",1,omitempty"`
	Device_uuid string `thrift:",2,omitempty"`
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

type GetGameDataListReturn struct {
	Error       ErrorRetCode    `thrift:",1,omitempty"`
	Server_time ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string          `thrift:",3,omitempty"`
	Call        string          `thrift:",4,omitempty"`
	Data        map[any]any     `thrift:",5,omitempty"` // TODO
	Maintenance MaintenanceData `thrift:",6,omitempty"`
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

type GetLocalPush struct {
	Call string               `thrift:",1,omitempty"`
	Data GetLocalPushDataInfo `thrift:",2,omitempty"`
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

type GetLocalPushDataInfo struct {
	Device_uuid string `thrift:",1,omitempty"`
}

type GetLocalPushRetDataInfo struct {
	I_LocalPushID           int32  `thrift:",1,omitempty"`
	S_LocalPushRegisterType int16  `thrift:",2,omitempty"`
	I_TimeForSeconds        int64  `thrift:",3,omitempty"`
	S_JsonMessage_ko        string `thrift:",4,omitempty"`
	S_JsonMessage_en        string `thrift:",5,omitempty"`
	S_JsonMessage_ja        string `thrift:",6,omitempty"`
}

type GetLocalPushReturn struct {
	Error       ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string          `thrift:",2,omitempty"`
	Call        string          `thrift:",3,omitempty"`
	Data        []any           `thrift:",4,omitempty"` // TODO
	Maintenance MaintenanceData `thrift:",5,omitempty"`
}

type GetMastersRank struct {
	Call string                 `thrift:",1,omitempty"`
	Data GetMastersRankDataInfo `thrift:",2,omitempty"`
}

type GetMastersRankDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Start_limit int32  `thrift:",5,omitempty"`
	End_limit   int32  `thrift:",6,omitempty"`
}

type GetMastersRankRet struct {
	All_per   int16  `thrift:",1,omitempty"`
	All_com   int16  `thrift:",2,omitempty"`
	Score     int32  `thrift:",3,omitempty"`
	U_avatar  int16  `thrift:",4,omitempty"`
	U_id      string `thrift:",5,omitempty"`
	U_nick    string `thrift:",6,omitempty"`
	U_country string `thrift:",7,omitempty"`
	Rank      int32  `thrift:",8,omitempty"`
}

type GetMastersRankRetDataInfo struct {
	Start_limit int16             `thrift:",1,omitempty"`
	End_limit   int16             `thrift:",2,omitempty"`
	My_rank     GetMastersRankRet `thrift:",3,omitempty"`
	Rank_list   []any             `thrift:",4,omitempty"` // TODO
}

type GetMastersRankReturn struct {
	Error       ErrorRetCode              `thrift:",1,omitempty"`
	Mode        string                    `thrift:",2,omitempty"`
	Call        string                    `thrift:",3,omitempty"`
	Data        GetMastersRankRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData           `thrift:",5,omitempty"`
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

type GetMusicInfoLanguage struct {
	Seq         int32  `thrift:",1,omitempty"`
	Music_idx   int64  `thrift:",2,omitempty"`
	Music_title string `thrift:",3,omitempty"`
	Artist      string `thrift:",4,omitempty"`
	Songwriter  string `thrift:",5,omitempty"`
	Lyricist    string `thrift:",6,omitempty"`
	Description string `thrift:",7,omitempty"`
}

type GetMusicInfoList struct {
	Call string                   `thrift:",1,omitempty"`
	Data GetMusicInfoListDataInfo `thrift:",2,omitempty"`
}

type GetMusicInfoListDataInfo struct {
	Country_code string `thrift:",1,omitempty"`
}

type GetMusicInfoListRetDataInfo struct {
	Language map[any]any `thrift:",1,omitempty"` // TODO
}

type GetMusicInfoListReturn struct {
	Error       ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string          `thrift:",2,omitempty"`
	Call        string          `thrift:",3,omitempty"`
	Data        map[any]any     `thrift:",4,omitempty"` // TODO
	Maintenance MaintenanceData `thrift:",5,omitempty"`
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

type GetMusicList struct {
	Call string               `thrift:",1,omitempty"`
	Data GetMusicListDataInfo `thrift:",2,omitempty"`
}

type GetMusicListDataInfo struct {
	Country_code string `thrift:",1,omitempty"`
}

type GetMusicListRetDataInfo struct {
	Music_idx              int32        `thrift:",1,omitempty"`
	Album_idx              int32        `thrift:",2,omitempty"`
	Genre_idx              int32        `thrift:",3,omitempty"`
	Difficulty1            int16        `thrift:",4,omitempty"`
	Difficulty2            int16        `thrift:",5,omitempty"`
	Difficulty3            int16        `thrift:",6,omitempty"`
	Difficulty4            int16        `thrift:",7,omitempty"`
	Purchaselink_type      int16        `thrift:",8,omitempty"`
	Credit_cnt             int16        `thrift:",9,omitempty"`
	Buy_type               int16        `thrift:",10,omitempty"`
	Free_type              int16        `thrift:",11,omitempty"`
	Price                  float64      `thrift:",12,omitempty"`
	Is_extra               int16        `thrift:",13,omitempty"`
	Extra_buy_type1        int16        `thrift:",14,omitempty"`
	Extra_price1           int64        `thrift:",15,omitempty"`
	Extra_buy_type2        int16        `thrift:",16,omitempty"`
	Extra_price2           int64        `thrift:",17,omitempty"`
	Buy_ad                 int32        `thrift:",18,omitempty"`
	Music_ver              int32        `thrift:",19,omitempty"`
	Bpm                    float64      `thrift:",20,omitempty"`
	Cdn_dir                string       `thrift:",21,omitempty"`
	Track_type             int16        `thrift:",22,omitempty"`
	File_type              int16        `thrift:",23,omitempty"`
	New_status             int16        `thrift:",24,omitempty"`
	Update_time            int32        `thrift:",25,omitempty"`
	Point_average          float64      `thrift:",26,omitempty"`
	Use_movie              int16        `thrift:",27,omitempty"`
	Movie_timing           int32        `thrift:",28,omitempty"`
	Music_product_aos      MusicProduct `thrift:",29,omitempty"`
	Music_product_ios      MusicProduct `thrift:",30,omitempty"`
	Music_product_aos_pack MusicProduct `thrift:",31,omitempty"`
	Music_product_ios_pack MusicProduct `thrift:",32,omitempty"`
	Bga_path               string       `thrift:",33,omitempty"`
	Bg_opacity             int16        `thrift:",34,omitempty"`
}

type GetMusicListReturn struct {
	Error       ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string          `thrift:",2,omitempty"`
	Call        string          `thrift:",3,omitempty"`
	Data        []any           `thrift:",4,omitempty"` // TODO
	Maintenance MaintenanceData `thrift:",5,omitempty"`
}

type GetMusicRank struct {
	Call string               `thrift:",1,omitempty"`
	Data GetMusicRankDataInfo `thrift:",2,omitempty"`
}

type GetMusicRankDataInfo struct {
	Music_idx int32 `thrift:",1,omitempty"`
}

type GetMusicRankRetDataInfo struct {
	Id    string `thrift:",1,omitempty"`
	Score int32  `thrift:",2,omitempty"`
}

type GetMusicRankReturn struct {
	Error       ErrorRetCode            `thrift:",1,omitempty"`
	Mode        string                  `thrift:",2,omitempty"`
	Call        string                  `thrift:",3,omitempty"`
	Data        GetMusicRankRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData         `thrift:",5,omitempty"`
}

type GetMusicReward struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        GetMusicRewardDataInfo `thrift:",2,omitempty"`
	Common_data ParamData              `thrift:",3,omitempty"`
}

type GetMusicRewardDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	I_ids       []any  `thrift:",5,omitempty"` // TODO
	I_levels    []any  `thrift:",6,omitempty"` // TODO
}

type GetMusicRewardRetDataInfo struct {
	Total_reward_value    int32       `thrift:",1,omitempty"`
	Reward_music_id       []any       `thrift:",2,omitempty"` // TODO
	Reward_value          []any       `thrift:",3,omitempty"` // TODO
	User_follower_profile []any       `thrift:",4,omitempty"` // TODO
	Error_data            map[any]any `thrift:",5,omitempty"` // TODO
}

type GetMusicRewardReturn struct {
	Error       ErrorRetCode              `thrift:",1,omitempty"`
	Server_time ServerTimeRet             `thrift:",2,omitempty"`
	Mode        string                    `thrift:",3,omitempty"`
	Call        string                    `thrift:",4,omitempty"`
	Data        GetMusicRewardRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData           `thrift:",6,omitempty"`
}

type GetNoticeList struct {
	Call string                `thrift:",1,omitempty"`
	Data GetNoticeListDataInfo `thrift:",2,omitempty"`
}

type GetNoticeListDataInfo struct {
	Device_uuid string `thrift:",1,omitempty"`
}

type GetNoticeListRetDataInfo struct {
	Seq          int32  `thrift:",1,omitempty"`
	Notice_name  string `thrift:",2,omitempty"`
	Location_url string `thrift:",3,omitempty"`
	Img_url      string `thrift:",4,omitempty"`
}

type GetNoticeListReturn struct {
	Error       ErrorRetCode    `thrift:",1,omitempty"`
	Server_time ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string          `thrift:",3,omitempty"`
	Call        string          `thrift:",4,omitempty"`
	Data        []any           `thrift:",5,omitempty"` // TODO
	Maintenance MaintenanceData `thrift:",6,omitempty"`
}

type GetPassGoodsShopData struct {
	I_id            int32 `thrift:",1,omitempty"`
	I_TicketID      int32 `thrift:",2,omitempty"`
	I_GoodsTicketID int32 `thrift:",3,omitempty"`
	I_ShopID        int32 `thrift:",4,omitempty"`
	B_IsActive      int16 `thrift:",5,omitempty"`
	I_SaleShopID    int32 `thrift:",6,omitempty"`
}

type GetPercentData struct {
	I_GroupID        int32 `thrift:",1,omitempty"`
	I_RewardType     int32 `thrift:",2,omitempty"`
	I_RewardID       int32 `thrift:",3,omitempty"`
	L_RewardQuantity int32 `thrift:",4,omitempty"`
	B_IsActive       int16 `thrift:",5,omitempty"`
	I_id             int32 `thrift:",6,omitempty"`
}

type GetPost struct {
	Call        string          `thrift:",1,omitempty"`
	Data        GetPostDataInfo `thrift:",2,omitempty"`
	Common_data ParamData       `thrift:",3,omitempty"`
}

type GetPostDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type GetPostList struct {
	Idx                 int64  `thrift:",1,omitempty"`
	Notice_type         int16  `thrift:",2,omitempty"`
	Title_ko            string `thrift:",3,omitempty"`
	Memo_ko             string `thrift:",4,omitempty"`
	Title_en            string `thrift:",5,omitempty"`
	Memo_en             string `thrift:",6,omitempty"`
	Title_jp            string `thrift:",7,omitempty"`
	Memo_jp             string `thrift:",8,omitempty"`
	Title_zh_chs        string `thrift:",9,omitempty"`
	Memo_zh_chs         string `thrift:",10,omitempty"`
	Title_zh_cht        string `thrift:",11,omitempty"`
	Memo_zh_cht         string `thrift:",12,omitempty"`
	Title_vi            string `thrift:",13,omitempty"`
	Memo_vi             string `thrift:",14,omitempty"`
	Title_es            string `thrift:",15,omitempty"`
	Memo_es             string `thrift:",16,omitempty"`
	Title_it            string `thrift:",17,omitempty"`
	Memo_it             string `thrift:",18,omitempty"`
	Title_id            string `thrift:",19,omitempty"`
	Memo_id             string `thrift:",20,omitempty"`
	Title_th            string `thrift:",21,omitempty"`
	Memo_th             string `thrift:",22,omitempty"`
	Title_pt            string `thrift:",23,omitempty"`
	Memo_pt             string `thrift:",24,omitempty"`
	Title_hi            string `thrift:",25,omitempty"`
	Memo_hi             string `thrift:",26,omitempty"`
	Have_reward         int16  `thrift:",27,omitempty"`
	Status              int16  `thrift:",28,omitempty"`
	Unlimit_flg         int16  `thrift:",29,omitempty"`
	Flg                 int16  `thrift:",30,omitempty"`
	Create_time         int64  `thrift:",31,omitempty"`
	Del_time            int64  `thrift:",32,omitempty"`
	Item_list           []any  `thrift:",33,omitempty"` // TODO
	Url                 string `thrift:",34,omitempty"`
	Image_resource_name string `thrift:",35,omitempty"`
}

type GetPostRetDataInfo struct {
	Server_time int64 `thrift:",1,omitempty"`
	Post_list   []any `thrift:",2,omitempty"` // TODO
}

type GetPostReturn struct {
	Error       ErrorRetCode       `thrift:",1,omitempty"`
	Server_time ServerTimeRet      `thrift:",2,omitempty"`
	Mode        string             `thrift:",3,omitempty"`
	Call        string             `thrift:",4,omitempty"`
	Data        GetPostRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData    `thrift:",6,omitempty"`
}

type GetPostTime struct {
	Call string              `thrift:",1,omitempty"`
	Data GetPostTimeDataInfo `thrift:",2,omitempty"`
}

type GetPostTimeDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type GetPostTimeRetDataInfo struct {
	Time int64 `thrift:",1,omitempty"`
}

type GetPostTimeReturn struct {
	Error       ErrorRetCode           `thrift:",1,omitempty"`
	Mode        string                 `thrift:",2,omitempty"`
	Call        string                 `thrift:",3,omitempty"`
	Data        GetPostTimeRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData        `thrift:",5,omitempty"`
}

type GetProfile struct {
	Call string             `thrift:",1,omitempty"`
	Data GetProfileDataInfo `thrift:",2,omitempty"`
}

type GetProfileDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Get_u_seq   int64  `thrift:",5,omitempty"`
}

type GetProfileMusic struct {
	Call string                  `thrift:",1,omitempty"`
	Data GetProfileMusicDataInfo `thrift:",2,omitempty"`
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

type GetProfileMusicRetDataInfo struct {
	Music_idx    int32 `thrift:",1,omitempty"`
	Single_cnt   int64 `thrift:",2,omitempty"`
	Multi_cnt    int64 `thrift:",3,omitempty"`
	Master_cnt   int64 `thrift:",4,omitempty"`
	Single_score int64 `thrift:",5,omitempty"`
	Multi_score  int64 `thrift:",6,omitempty"`
}

type GetProfileMusicReturn struct {
	Error       ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string          `thrift:",2,omitempty"`
	Call        string          `thrift:",3,omitempty"`
	Data        []any           `thrift:",4,omitempty"` // TODO
	Maintenance MaintenanceData `thrift:",5,omitempty"`
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

type GetProfileReturn struct {
	Error       ErrorRetCode          `thrift:",1,omitempty"`
	Mode        string                `thrift:",2,omitempty"`
	Call        string                `thrift:",3,omitempty"`
	Data        GetProfileRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData       `thrift:",5,omitempty"`
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

type GetPropLevelData struct {
	I_id         int32   `thrift:",1,omitempty"`
	I_PropId     int64   `thrift:",2,omitempty"`
	I_Level      int64   `thrift:",3,omitempty"`
	D_Cost       float64 `thrift:",4,omitempty"`
	I_PerksId    int64   `thrift:",5,omitempty"`
	D_PerksValue float64 `thrift:",6,omitempty"`
	B_IsActive   int16   `thrift:",7,omitempty"`
}

type GetPurchaseDiamondLog struct {
	Call string                        `thrift:",1,omitempty"`
	Data GetPurchaseDiamondLogDataInfo `thrift:",2,omitempty"`
}

type GetPurchaseDiamondLogDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type GetPurchaseDiamondLogRetDataInfo struct {
	Bonus1 int64 `thrift:",1,omitempty"`
	Bonus2 int64 `thrift:",2,omitempty"`
	Bonus3 int64 `thrift:",3,omitempty"`
	Bonus4 int64 `thrift:",4,omitempty"`
	Bonus5 int64 `thrift:",5,omitempty"`
	Bonus6 int64 `thrift:",6,omitempty"`
}

type GetPurchaseDiamondLogReturn struct {
	Error       ErrorRetCode                     `thrift:",1,omitempty"`
	Mode        string                           `thrift:",2,omitempty"`
	Call        string                           `thrift:",3,omitempty"`
	Data        GetPurchaseDiamondLogRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData                  `thrift:",5,omitempty"`
}

type GetRankMain struct {
	Call string              `thrift:",1,omitempty"`
	Data GetRankMainDataInfo `thrift:",2,omitempty"`
}

type GetRankMainDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type GetRankMainRetDataInfo struct {
	My_country_rank       GetCountryMyRankRet    `thrift:",1,omitempty"`
	Total_country_my_rank GetTotalCountryRankRet `thrift:",2,omitempty"`
	Country_rank          []any                  `thrift:",3,omitempty"` // TODO
	Masters_rank          GetMastersRankRet      `thrift:",4,omitempty"`
}

type GetRankMainReturn struct {
	Error       ErrorRetCode           `thrift:",1,omitempty"`
	Mode        string                 `thrift:",2,omitempty"`
	Call        string                 `thrift:",3,omitempty"`
	Data        GetRankMainRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData        `thrift:",5,omitempty"`
}

type GetReviewPoint struct {
	Call string           `thrift:",1,omitempty"`
	Data GetAlbumDataInfo `thrift:",2,omitempty"`
}

type GetReviewPointDataInfo struct {
	Null string `thrift:",1,omitempty"`
}

type GetReviewPointRetDataInfo struct {
	Music_idx  int32   `thrift:",1,omitempty"`
	Difficulty int16   `thrift:",2,omitempty"`
	Point      float64 `thrift:",3,omitempty"`
}

type GetReviewPointReturn struct {
	Error       ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string          `thrift:",2,omitempty"`
	Call        string          `thrift:",3,omitempty"`
	Data        []any           `thrift:",4,omitempty"` // TODO
	Maintenance MaintenanceData `thrift:",5,omitempty"`
}

type GetRewardGroupData struct {
	I_id               int64 `thrift:",1,omitempty"`
	I_Group            int64 `thrift:",2,omitempty"`
	I_RewardType       int64 `thrift:",3,omitempty"`
	I_RewardID         int64 `thrift:",4,omitempty"`
	L_RewardQuantity   int64 `thrift:",5,omitempty"`
	I_BuyFirstQuantity int64 `thrift:",6,omitempty"`
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

type GetSamSeckList struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        GetSamSeckListDataInfo `thrift:",2,omitempty"`
	Common_data ParamData              `thrift:",3,omitempty"`
}

type GetSamSeckListDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type GetSamSeckListRetDataInfo struct {
	Event_type string `thrift:",1,omitempty"`
	RewardList []any  `thrift:",2,omitempty"` // TODO
}

type GetSamSeckListReturn struct {
	Error       ErrorRetCode              `thrift:",1,omitempty"`
	Server_time ServerTimeRet             `thrift:",2,omitempty"`
	Mode        string                    `thrift:",3,omitempty"`
	Call        string                    `thrift:",4,omitempty"`
	Data        GetSamSeckListRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData           `thrift:",6,omitempty"`
}

type GetSamSeckReward struct {
	Call        string                   `thrift:",1,omitempty"`
	Data        GetSamSeckRewardDataInfo `thrift:",2,omitempty"`
	Common_data ParamData                `thrift:",3,omitempty"`
}

type GetSamSeckRewardDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	I_id        int32  `thrift:",5,omitempty"`
}

type GetSamSeckRewardRetDataInfo struct {
	Step int16 `thrift:",1,omitempty"`
}

type GetSamSeckRewardReturn struct {
	Error       ErrorRetCode                `thrift:",1,omitempty"`
	Server_time ServerTimeRet               `thrift:",2,omitempty"`
	Mode        string                      `thrift:",3,omitempty"`
	Call        string                      `thrift:",4,omitempty"`
	Data        GetSamSeckRewardRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData             `thrift:",6,omitempty"`
}

type GetSelectRewardData struct {
	I_id               int32 `thrift:",1,omitempty"`
	I_GroupId          int32 `thrift:",2,omitempty"`
	I_RewardGroupId    int32 `thrift:",3,omitempty"`
	I_AltRewardGroupId int32 `thrift:",4,omitempty"`
	B_IsActive         int16 `thrift:",5,omitempty"`
}

type GetServerList struct {
	Call string                `thrift:",1,omitempty"`
	Data GetServerListDataInfo `thrift:",2,omitempty"`
}

type GetServerListDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Type        string `thrift:",5,omitempty"`
}

type GetServerListRetDataInfo struct {
	User_data    UserData `thrift:",1,omitempty"`
	Channel_list []any    `thrift:",2,omitempty"` // TODO
}

type GetServerListReturn struct {
	Error       ErrorRetCode             `thrift:",1,omitempty"`
	Mode        string                   `thrift:",2,omitempty"`
	Call        string                   `thrift:",3,omitempty"`
	Data        GetServerListRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData          `thrift:",5,omitempty"`
}

type GetServerTime struct {
	Call        string       `thrift:",1,omitempty"`
	Data        MainDataInfo `thrift:",2,omitempty"`
	Sub_mode    string       `thrift:",3,omitempty"`
	Common_data ParamData    `thrift:",4,omitempty"`
}

type GetServerTimeDataInfo struct {
	Device_uuid string `thrift:",1,omitempty"`
}

type GetServerTimeRetDataInfo struct {
	Time     int64 `thrift:",1,omitempty"`
	Datetime int64 `thrift:",2,omitempty"`
}

type GetServerTimeReturn struct {
	Error       ErrorRetCode             `thrift:",1,omitempty"`
	Server_time ServerTimeRet            `thrift:",2,omitempty"`
	Mode        string                   `thrift:",3,omitempty"`
	Call        string                   `thrift:",4,omitempty"`
	Data        GetServerTimeRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData          `thrift:",6,omitempty"`
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

type GetSubscribePassReward struct {
	I_id              int64 `thrift:",1,omitempty"`
	I_Group           int64 `thrift:",2,omitempty"`
	I_Step            int16 `thrift:",3,omitempty"`
	I_Goal            int64 `thrift:",4,omitempty"`
	I_FreeRewardGroup int64 `thrift:",5,omitempty"`
	I_PaidRewardGroup int64 `thrift:",6,omitempty"`
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

type GetTabList struct {
	Call string       `thrift:",1,omitempty"`
	Data InitDataInfo `thrift:",2,omitempty"`
}

type GetTabListDataInfo struct {
	Type        string `thrift:",1,omitempty"`
	Device_uuid string `thrift:",2,omitempty"`
}

type GetTabListRetDataInfo struct {
	Tab_list []any `thrift:",1,omitempty"` // TODO
}

type GetTabListReturn struct {
	Error       ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string          `thrift:",2,omitempty"`
	Call        string          `thrift:",3,omitempty"`
	Data        map[any]any     `thrift:",4,omitempty"` // TODO
	Maintenance MaintenanceData `thrift:",5,omitempty"`
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

type GetTotalCountryRank struct {
	Call string                      `thrift:",1,omitempty"`
	Data GetTotalCountryRankDataInfo `thrift:",2,omitempty"`
}

type GetTotalCountryRankDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Start_limit int32  `thrift:",5,omitempty"`
	End_limit   int32  `thrift:",6,omitempty"`
}

type GetTotalCountryRankRet struct {
	G           int16  `thrift:",1,omitempty"`
	S           int16  `thrift:",2,omitempty"`
	B           int16  `thrift:",3,omitempty"`
	Total_score int32  `thrift:",4,omitempty"`
	U_avatar    int16  `thrift:",5,omitempty"`
	U_id        string `thrift:",6,omitempty"`
	U_nick      string `thrift:",7,omitempty"`
	U_country   string `thrift:",8,omitempty"`
	Rank        int32  `thrift:",9,omitempty"`
}

type GetTotalCountryRankRetDataInfo struct {
	Start_limit int16                  `thrift:",1,omitempty"`
	End_limit   int16                  `thrift:",2,omitempty"`
	My_rank     GetTotalCountryRankRet `thrift:",3,omitempty"`
	Rank_list   []any                  `thrift:",4,omitempty"` // TODO
}

type GetTotalCountryRankReturn struct {
	Error       ErrorRetCode                   `thrift:",1,omitempty"`
	Mode        string                         `thrift:",2,omitempty"`
	Call        string                         `thrift:",3,omitempty"`
	Data        GetTotalCountryRankRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData                `thrift:",5,omitempty"`
}

type GetTotalCountryUserRankRet struct {
	G           int16  `thrift:",1,omitempty"`
	S           int16  `thrift:",2,omitempty"`
	B           int16  `thrift:",3,omitempty"`
	Total_score int32  `thrift:",4,omitempty"`
	U_avatar    int16  `thrift:",5,omitempty"`
	U_id        string `thrift:",6,omitempty"`
	U_nick      string `thrift:",7,omitempty"`
	U_country   string `thrift:",8,omitempty"`
}

type GetTotalMastersRankRet struct {
	All_per   int16  `thrift:",1,omitempty"`
	All_com   int16  `thrift:",2,omitempty"`
	Score     int32  `thrift:",3,omitempty"`
	U_avatar  int16  `thrift:",4,omitempty"`
	U_id      string `thrift:",5,omitempty"`
	U_nick    string `thrift:",6,omitempty"`
	U_country string `thrift:",7,omitempty"`
}

type GetTotalMusicMyRank struct {
	Grade string `thrift:",1,omitempty"`
	Rank  int32  `thrift:",2,omitempty"`
	Score int32  `thrift:",3,omitempty"`
}

type GetTotalMusicRank struct {
	Call string                    `thrift:",1,omitempty"`
	Data GetTotalMusicRankDataInfo `thrift:",2,omitempty"`
}

type GetTotalMusicRankDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Music_idx   int32  `thrift:",5,omitempty"`
	Start_limit int32  `thrift:",6,omitempty"`
	End_limit   int32  `thrift:",7,omitempty"`
}

type GetTotalMusicRankList struct {
	Rank    int32  `thrift:",1,omitempty"`
	Id      string `thrift:",2,omitempty"`
	Nick    string `thrift:",3,omitempty"`
	Country string `thrift:",4,omitempty"`
	Score   int32  `thrift:",5,omitempty"`
}

type GetTotalMusicRankList1 struct {
	Rank  int32  `thrift:",1,omitempty"`
	Id    string `thrift:",2,omitempty"`
	Score int32  `thrift:",3,omitempty"`
}

type GetTotalMusicRankList2 struct {
	Rank  int32  `thrift:",1,omitempty"`
	Id    string `thrift:",2,omitempty"`
	Score int32  `thrift:",3,omitempty"`
}

type GetTotalMusicRankList3 struct {
	Rank  int32  `thrift:",1,omitempty"`
	Id    string `thrift:",2,omitempty"`
	Score int32  `thrift:",3,omitempty"`
}

type GetTotalMusicRankRetData struct {
	Total_rank_list []any `thrift:",1,omitempty"` // TODO
	Rank_list1      []any `thrift:",2,omitempty"` // TODO
	Rank_list2      []any `thrift:",3,omitempty"` // TODO
	Rank_list3      []any `thrift:",4,omitempty"` // TODO
	My_rank_list    []any `thrift:",5,omitempty"` // TODO
}

type GetTotalMusicRankRetDataInfo struct {
	Total_rank_list []any `thrift:",1,omitempty"` // TODO
	Rank_list1      []any `thrift:",2,omitempty"` // TODO
	Rank_list2      []any `thrift:",3,omitempty"` // TODO
	Rank_list3      []any `thrift:",4,omitempty"` // TODO
	My_rank_list    []any `thrift:",5,omitempty"` // TODO
}

type GetTotalMusicRankReturn struct {
	Error       ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string          `thrift:",2,omitempty"`
	Call        string          `thrift:",3,omitempty"`
	Data        []any           `thrift:",4,omitempty"` // TODO
	Maintenance MaintenanceData `thrift:",5,omitempty"`
}

type GetTourList struct {
	Call string              `thrift:",1,omitempty"`
	Data GetTourListDataInfo `thrift:",2,omitempty"`
}

type GetTourListDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type GetTourListRetDataInfo struct {
	Tour_idx int32  `thrift:",1,omitempty"`
	Title    string `thrift:",2,omitempty"`
}

type GetTourListReturn struct {
	Error       ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string          `thrift:",2,omitempty"`
	Call        string          `thrift:",3,omitempty"`
	Data        []any           `thrift:",4,omitempty"` // TODO
	Maintenance MaintenanceData `thrift:",5,omitempty"`
}

type GetTourRank struct {
	Call string              `thrift:",1,omitempty"`
	Data GetTourRankDataInfo `thrift:",2,omitempty"`
}

type GetTourRankData struct {
	U_seq     int64  `thrift:",1,omitempty"`
	U_id      string `thrift:",2,omitempty"`
	U_nick    string `thrift:",3,omitempty"`
	U_avatar  int32  `thrift:",4,omitempty"`
	U_country string `thrift:",5,omitempty"`
	Rank      int64  `thrift:",6,omitempty"`
	Score     int64  `thrift:",7,omitempty"`
	Medal     []any  `thrift:",8,omitempty"` // TODO
}

type GetTourRankDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Tour_idx    int32  `thrift:",5,omitempty"`
	Track_idx   int32  `thrift:",6,omitempty"`
	Mode        string `thrift:",7,omitempty"`
	Start_limit int64  `thrift:",8,omitempty"`
	End_limit   int64  `thrift:",9,omitempty"`
}

type GetTourRankRetDataInfo struct {
	My_rank GetTourRankData `thrift:",1,omitempty"`
	Rank    []any           `thrift:",2,omitempty"` // TODO
}

type GetTourRankReturn struct {
	Error       ErrorRetCode           `thrift:",1,omitempty"`
	Mode        string                 `thrift:",2,omitempty"`
	Call        string                 `thrift:",3,omitempty"`
	Data        GetTourRankRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData        `thrift:",5,omitempty"`
}

type GetTransferId struct {
	Call string                `thrift:",1,omitempty"`
	Data GetTransferIdDataInfo `thrift:",2,omitempty"`
}

type GetTransferIdDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type GetTransferIdRetDataInfo struct {
	Transfer_id string `thrift:",1,omitempty"`
}

type GetTransferIdReturn struct {
	Error       ErrorRetCode             `thrift:",1,omitempty"`
	Mode        string                   `thrift:",2,omitempty"`
	Call        string                   `thrift:",3,omitempty"`
	Data        GetTransferIdRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData          `thrift:",5,omitempty"`
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

type GetUpdateTime struct {
	Call        string                `thrift:",1,omitempty"`
	Data        GetUpdateTimeDataInfo `thrift:",2,omitempty"`
	Common_data ParamData             `thrift:",3,omitempty"`
}

type GetUpdateTimeDataInfo struct {
	Device_uuid string `thrift:",1,omitempty"`
}

type GetUpdateTimeRetDataInfo struct {
	Upd_time int64 `thrift:",1,omitempty"`
}

type GetUpdateTimeReturn struct {
	Error       ErrorRetCode             `thrift:",1,omitempty"`
	Server_time ServerTimeRet            `thrift:",2,omitempty"`
	Mode        string                   `thrift:",3,omitempty"`
	Call        string                   `thrift:",4,omitempty"`
	Data        GetUpdateTimeRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData          `thrift:",6,omitempty"`
}

type GetUserCollection struct {
	Call string                    `thrift:",1,omitempty"`
	Data GetUserCollectionDataInfo `thrift:",2,omitempty"`
}

type GetUserCollectionDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type GetUserCollectionRetDataInfo struct {
	Music_idx   int64  `thrift:",1,omitempty"`
	Value       int64  `thrift:",2,omitempty"`
	Get_type    int16  `thrift:",3,omitempty"`
	Get_time    string `thrift:",4,omitempty"`
	Create_time string `thrift:",5,omitempty"`
}

type GetUserCollectionReturn struct {
	Error       ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string          `thrift:",2,omitempty"`
	Call        string          `thrift:",3,omitempty"`
	Data        []any           `thrift:",4,omitempty"` // TODO
	Maintenance MaintenanceData `thrift:",5,omitempty"`
}

type GetUserInfo struct {
	Call        string              `thrift:",1,omitempty"`
	Data        GetUserInfoDataInfo `thrift:",2,omitempty"`
	Common_data ParamData           `thrift:",3,omitempty"`
}

type GetUserInfoDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type GetUserInfoRetDataInfo struct {
	U_cp    int64   `thrift:",1,omitempty"`
	U_candy float64 `thrift:",2,omitempty"`
	U_like  float64 `thrift:",3,omitempty"`
	U_fans  int64   `thrift:",4,omitempty"`
}

type GetUserInfoReturn struct {
	Error       ErrorRetCode           `thrift:",1,omitempty"`
	Server_time ServerTimeRet          `thrift:",2,omitempty"`
	Mode        string                 `thrift:",3,omitempty"`
	Call        string                 `thrift:",4,omitempty"`
	Data        GetUserInfoRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData        `thrift:",6,omitempty"`
}

type GetUserItemInven struct {
	Call string                   `thrift:",1,omitempty"`
	Data GetUserItemInvenDataInfo `thrift:",2,omitempty"`
}

type GetUserItemInvenDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type GetUserItemInvenRetDataInfo struct {
	Item_inven_list map[any]any `thrift:",1,omitempty"` // TODO
	Empty           string      `thrift:",2,omitempty"`
}

type GetUserItemInvenReturn struct {
	Error       ErrorRetCode                `thrift:",1,omitempty"`
	Server_time ServerTimeRet               `thrift:",2,omitempty"`
	Mode        string                      `thrift:",3,omitempty"`
	Call        string                      `thrift:",4,omitempty"`
	Data        GetUserItemInvenRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData             `thrift:",6,omitempty"`
}

type GetVarietyStore struct {
	Call        string                  `thrift:",1,omitempty"`
	Data        GetVarietyStoreDataInfo `thrift:",2,omitempty"`
	Common_data ParamData               `thrift:",3,omitempty"`
}

type GetVarietyStoreDataInfo struct {
	Device_uuid string `thrift:",1,omitempty"`
}

type GetVarietyStoreRetDataInfo struct {
	I_id                 int32  `thrift:",1,omitempty"`
	S_ResourceName       string `thrift:",2,omitempty"`
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
	I_Cost               int32  `thrift:",27,omitempty"`
	I_RewardType         int32  `thrift:",28,omitempty"`
	I_RewardId           int32  `thrift:",29,omitempty"`
	I_RewardQuantity     int32  `thrift:",30,omitempty"`
}

type GetVarietyStoreReturn struct {
	Error       ErrorRetCode    `thrift:",1,omitempty"`
	Server_time ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string          `thrift:",3,omitempty"`
	Call        string          `thrift:",4,omitempty"`
	Data        []any           `thrift:",5,omitempty"` // TODO
	Maintenance MaintenanceData `thrift:",6,omitempty"`
}

type HibernationLog struct {
	Call string                 `thrift:",1,omitempty"`
	Data HibernationLogDataInfo `thrift:",2,omitempty"`
}

type HibernationLogDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type HibernationLogRetDataInfo struct {
	Action int32 `thrift:",1,omitempty"`
}

type HibernationLogReturn struct {
	Error       ErrorRetCode              `thrift:",1,omitempty"`
	Server_time ServerTimeRet             `thrift:",2,omitempty"`
	Mode        string                    `thrift:",3,omitempty"`
	Call        string                    `thrift:",4,omitempty"`
	Data        HibernationLogRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData           `thrift:",6,omitempty"`
}

type InitDataInfo struct {
	Type        string `thrift:",1,omitempty"`
	Os          int16  `thrift:",2,omitempty"`
	Ver         int32  `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type InitRetDataInfo struct {
	Idx      int16  `thrift:",1,omitempty"`
	Game_url string `thrift:",2,omitempty"`
	Cdn_url  string `thrift:",3,omitempty"`
}

type InitReturn struct {
	Error       ErrorRetCode    `thrift:",1,omitempty"`
	Server_time ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string          `thrift:",3,omitempty"`
	Call        string          `thrift:",4,omitempty"`
	Data        InitRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData `thrift:",6,omitempty"`
}

type ItemList struct {
	I_RewardType     int32   `thrift:",1,omitempty"`
	I_RewardId       int32   `thrift:",2,omitempty"`
	D_RewardQuantity float64 `thrift:",3,omitempty"`
}

type ItemPurchase struct {
	Call     string               `thrift:",1,omitempty"`
	Data     ItemPurchaseDataInfo `thrift:",2,omitempty"`
	Sub_mode string               `thrift:",3,omitempty"`
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

type ItemPurchaseRetDataInfo struct {
	Item_type   int16  `thrift:",1,omitempty"`
	Item_value  int64  `thrift:",2,omitempty"`
	Pay_id      string `thrift:",3,omitempty"`
	Product_id  string `thrift:",4,omitempty"`
	Purchase_id string `thrift:",5,omitempty"`
}

type ItemPurchaseReturn struct {
	Error       ErrorRetCode            `thrift:",1,omitempty"`
	Mode        string                  `thrift:",2,omitempty"`
	Call        string                  `thrift:",3,omitempty"`
	Data        ItemPurchaseRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData         `thrift:",5,omitempty"`
}

type ItemStoreList struct {
	Call        string                `thrift:",1,omitempty"`
	Data        ItemStoreListDataInfo `thrift:",2,omitempty"`
	Common_data ParamData             `thrift:",3,omitempty"`
}

type ItemStoreListDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Store_type  int16  `thrift:",5,omitempty"`
}

type ItemStoreListRetDataInfo struct {
	Bonus             int64   `thrift:",1,omitempty"`
	Idx               int64   `thrift:",2,omitempty"`
	Pp_idx            int64   `thrift:",3,omitempty"`
	Title_cn          string  `thrift:",4,omitempty"`
	Title_en          string  `thrift:",5,omitempty"`
	Title_jp          string  `thrift:",6,omitempty"`
	Desc_cn           string  `thrift:",7,omitempty"`
	Desc_en           string  `thrift:",8,omitempty"`
	Desc_jp           string  `thrift:",9,omitempty"`
	Img_url           string  `thrift:",10,omitempty"`
	Buy_type          int16   `thrift:",11,omitempty"`
	Item_type         int16   `thrift:",12,omitempty"`
	Prev_price        float64 `thrift:",13,omitempty"`
	Price             float64 `thrift:",14,omitempty"`
	Str_prev_price    string  `thrift:",15,omitempty"`
	Str_price         string  `thrift:",16,omitempty"`
	Item_value        int64   `thrift:",17,omitempty"`
	Sale_status       int16   `thrift:",18,omitempty"`
	Aos_pp_product_id string  `thrift:",19,omitempty"`
	Ios_pp_product_id string  `thrift:",20,omitempty"`
}

type ItemStoreListReturn struct {
	Error       ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string          `thrift:",2,omitempty"`
	Call        string          `thrift:",3,omitempty"`
	Data        []any           `thrift:",4,omitempty"` // TODO
	Maintenance MaintenanceData `thrift:",5,omitempty"`
}

type LastSaveTime struct {
	Call        string               `thrift:",1,omitempty"`
	Data        LastSaveTimeDataInfo `thrift:",2,omitempty"`
	Common_data ParamData            `thrift:",3,omitempty"`
}

type LastSaveTimeDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type LastSaveTimeRetDataInfo struct {
	Last_save_time int64  `thrift:",1,omitempty"`
	Device_uuid    string `thrift:",2,omitempty"`
}

type LastSaveTimeReturn struct {
	Error       ErrorRetCode            `thrift:",1,omitempty"`
	Server_time ServerTimeRet           `thrift:",2,omitempty"`
	Mode        string                  `thrift:",3,omitempty"`
	Call        string                  `thrift:",4,omitempty"`
	Data        LastSaveTimeRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData         `thrift:",6,omitempty"`
}

type LogCashComplete struct {
	Call string                  `thrift:",1,omitempty"`
	Data LogCashCompleteDataInfo `thrift:",2,omitempty"`
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

type LogCashCompleteReturn struct {
	Error       ErrorRetCode               `thrift:",1,omitempty"`
	Server_time ServerTimeRet              `thrift:",2,omitempty"`
	Mode        string                     `thrift:",3,omitempty"`
	Call        string                     `thrift:",4,omitempty"`
	Data        LogCashCompleteRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData            `thrift:",6,omitempty"`
}

type Main struct {
	Call     string       `thrift:",1,omitempty"`
	Data     MainDataInfo `thrift:",2,omitempty"`
	Sub_mode string       `thrift:",3,omitempty"`
}

type MainDataInfo struct {
	U_seq        int32  `thrift:",1,omitempty"`
	U_id         string `thrift:",2,omitempty"`
	Uuid         string `thrift:",3,omitempty"`
	Device_uuid  string `thrift:",4,omitempty"`
	Country_code string `thrift:",5,omitempty"`
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

type MainGameInfo struct {
	Album_list      []any                        `thrift:",1,omitempty"` // TODO
	Music_list      []any                        `thrift:",2,omitempty"` // TODO
	Album_language  map[any]any                  `thrift:",3,omitempty"` // TODO
	Music_language  map[any]any                  `thrift:",4,omitempty"` // TODO
	Tab             map[any]any                  `thrift:",5,omitempty"` // TODO
	Notice          []any                        `thrift:",6,omitempty"` // TODO
	Default_setting GetDefaultSettingRetDataInfo `thrift:",7,omitempty"`
}

type MainRetDataInfo struct {
	GameInfo MainGameInfo `thrift:",1,omitempty"`
	Features []any        `thrift:",2,omitempty"` // TODO
}

type MainReturn struct {
	Error       ErrorRetCode    `thrift:",1,omitempty"`
	Server_time ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string          `thrift:",3,omitempty"`
	Call        string          `thrift:",4,omitempty"`
	Data        MainRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData `thrift:",6,omitempty"`
}

type MaintenanceData struct {
	Code           int16  `thrift:",1,omitempty"`
	Title          string `thrift:",2,omitempty"`
	Description    string `thrift:",3,omitempty"`
	Utc_time       int16  `thrift:",4,omitempty"`
	Facebook_url   string `thrift:",5,omitempty"`
	Start_datetime string `thrift:",6,omitempty"`
	End_datetime   string `thrift:",7,omitempty"`
}

type MaintenanceList struct {
	Idx            int16  `thrift:",1,omitempty"`
	Title          string `thrift:",2,omitempty"`
	Start_datetime int32  `thrift:",3,omitempty"`
	End_datetime   int32  `thrift:",4,omitempty"`
	View_flg       int16  `thrift:",5,omitempty"`
}

type MoreGames struct {
	Call string            `thrift:",1,omitempty"`
	Data MoreGamesDataInfo `thrift:",2,omitempty"`
}

type MoreGamesDataInfo struct {
	Locale     string `thrift:",1,omitempty"`
	Adset_type string `thrift:",2,omitempty"`
	Uuid       string `thrift:",3,omitempty"`
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

type MoreGamesReturn struct {
	Error       ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string          `thrift:",2,omitempty"`
	Call        string          `thrift:",3,omitempty"`
	Data        map[any]any     `thrift:",4,omitempty"` // TODO
	Maintenance MaintenanceData `thrift:",5,omitempty"`
}

type MusicData struct {
	Music_idx int32 `thrift:",1,omitempty"`
	Score1    int32 `thrift:",2,omitempty"`
	Score2    int32 `thrift:",3,omitempty"`
	Score3    int32 `thrift:",4,omitempty"`
}

type MusicListInfo struct {
	Music_idx   int64       `thrift:",1,omitempty"`
	Genre_idx   int32       `thrift:",2,omitempty"`
	Difficulty1 int16       `thrift:",3,omitempty"`
	Difficulty2 int16       `thrift:",4,omitempty"`
	Difficulty3 int16       `thrift:",5,omitempty"`
	Difficulty4 int16       `thrift:",6,omitempty"`
	Buy_type    int16       `thrift:",7,omitempty"`
	Free_type   int16       `thrift:",8,omitempty"`
	Price       float64     `thrift:",9,omitempty"`
	Buy_ad      int32       `thrift:",10,omitempty"`
	Music_ver   int32       `thrift:",11,omitempty"`
	Bpm         float64     `thrift:",12,omitempty"`
	Cdn_dir     string      `thrift:",13,omitempty"`
	Track_type  int16       `thrift:",14,omitempty"`
	File_type   int16       `thrift:",15,omitempty"`
	New_status  int16       `thrift:",16,omitempty"`
	Update_time int32       `thrift:",17,omitempty"`
	Language    map[any]any `thrift:",18,omitempty"` // TODO
}

type MusicListLanguage struct {
	Title      string `thrift:",1,omitempty"`
	Artist     string `thrift:",2,omitempty"`
	Songwriter string `thrift:",3,omitempty"`
	Lyricist   string `thrift:",4,omitempty"`
}

type MusicPointReview struct {
	Call string                   `thrift:",1,omitempty"`
	Data MusicPointReviewDataInfo `thrift:",2,omitempty"`
}

type MusicPointReviewDataInfo struct {
	U_seq        int32  `thrift:",1,omitempty"`
	U_id         string `thrift:",2,omitempty"`
	Uuid         string `thrift:",3,omitempty"`
	Device_uuid  string `thrift:",4,omitempty"`
	Review_point []any  `thrift:",5,omitempty"` // TODO
}

type MusicPointReviewRetDataInfo struct {
	Music_idx  int32   `thrift:",1,omitempty"`
	Point      float64 `thrift:",2,omitempty"`
	Difficulty int16   `thrift:",3,omitempty"`
}

type MusicPointReviewReturn struct {
	Error       ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string          `thrift:",2,omitempty"`
	Call        string          `thrift:",3,omitempty"`
	Data        []any           `thrift:",4,omitempty"` // TODO
	Maintenance MaintenanceData `thrift:",5,omitempty"`
}

type MusicProduct struct {
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

type NewStoreList struct {
	Call string               `thrift:",1,omitempty"`
	Data NewStoreListDataInfo `thrift:",2,omitempty"`
}

type NewStoreListData struct {
	Store_idx int16 `thrift:",1,omitempty"`
	Music_idx int32 `thrift:",2,omitempty"`
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

type NewStoreListReturn struct {
	Error       ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string          `thrift:",2,omitempty"`
	Call        string          `thrift:",3,omitempty"`
	Data        map[any]any     `thrift:",4,omitempty"` // TODO
	Maintenance MaintenanceData `thrift:",5,omitempty"`
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

type PaidEventPoint struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        PaidEventPointDataInfo `thrift:",2,omitempty"`
	Common_data ParamData              `thrift:",3,omitempty"`
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

type PaidEventPointRetDataInfo struct {
	U_cp          int64   `thrift:",1,omitempty"`
	U_candy       float64 `thrift:",2,omitempty"`
	I_SubscribeID int64   `thrift:",3,omitempty"`
	I_Point       int64   `thrift:",4,omitempty"`
	I_Version     int32   `thrift:",5,omitempty"`
}

type PaidEventPointReturn struct {
	Error       ErrorRetCode              `thrift:",1,omitempty"`
	Server_time ServerTimeRet             `thrift:",2,omitempty"`
	Mode        string                    `thrift:",3,omitempty"`
	Call        string                    `thrift:",4,omitempty"`
	Data        PaidEventPointRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData           `thrift:",6,omitempty"`
}

type ParamData struct {
	Client_ver int32  `thrift:",1,omitempty"`
	Type       string `thrift:",2,omitempty"`
	Os         int16  `thrift:",3,omitempty"`
}

type PlayCheck struct {
	Call string            `thrift:",1,omitempty"`
	Data PlayCheckDataInfo `thrift:",2,omitempty"`
}

type PlayCheckDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Tour_idx    int32  `thrift:",5,omitempty"`
	Track_idx   int32  `thrift:",6,omitempty"`
	Music_idx   int32  `thrift:",7,omitempty"`
}

type PlayCheckRetDataInfo struct {
	Basic       int32 `thrift:",1,omitempty"`
	Pro         int32 `thrift:",2,omitempty"`
	Legend      int32 `thrift:",3,omitempty"`
	Extra       int32 `thrift:",4,omitempty"`
	Server_time int64 `thrift:",5,omitempty"`
}

type PlayCheckReturn struct {
	Error       ErrorRetCode         `thrift:",1,omitempty"`
	Mode        string               `thrift:",2,omitempty"`
	Call        string               `thrift:",3,omitempty"`
	Data        PlayCheckRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData      `thrift:",5,omitempty"`
}

type PlayMusic struct {
	Call string            `thrift:",1,omitempty"`
	Data PlayMusicDataInfo `thrift:",2,omitempty"`
}

type PlayMusicDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Music_idx   int32  `thrift:",5,omitempty"`
	Difficulty  int32  `thrift:",6,omitempty"`
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

type PlayMusicReturn struct {
	Error       ErrorRetCode         `thrift:",1,omitempty"`
	Server_time ServerTimeRet        `thrift:",2,omitempty"`
	Mode        string               `thrift:",3,omitempty"`
	Call        string               `thrift:",4,omitempty"`
	Data        PlayMusicRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData      `thrift:",6,omitempty"`
}

type ProvidePost struct {
	Call        string              `thrift:",1,omitempty"`
	Data        ProvidePostDataInfo `thrift:",2,omitempty"`
	Common_data ParamData           `thrift:",3,omitempty"`
}

type ProvidePostDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Idx         int64  `thrift:",5,omitempty"`
	Type        int16  `thrift:",6,omitempty"`
}

type ProvidePostRetDataInfo struct {
	I_RewardType     int32   `thrift:",1,omitempty"`
	I_RewardId       int32   `thrift:",2,omitempty"`
	D_RewardQuantity float64 `thrift:",3,omitempty"`
}

type ProvidePostReturn struct {
	Error       ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string          `thrift:",2,omitempty"`
	Call        string          `thrift:",3,omitempty"`
	Data        []any           `thrift:",4,omitempty"` // TODO
	Maintenance MaintenanceData `thrift:",5,omitempty"`
}

type ReadPost struct {
	Call string           `thrift:",1,omitempty"`
	Data ReadPostDataInfo `thrift:",2,omitempty"`
}

type ReadPostDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Idx         int64  `thrift:",5,omitempty"`
}

type ReadPostRetDataInfo struct {
	Idx int64 `thrift:",1,omitempty"`
}

type ReadPostReturn struct {
	Error       ErrorRetCode        `thrift:",1,omitempty"`
	Mode        string              `thrift:",2,omitempty"`
	Call        string              `thrift:",3,omitempty"`
	Data        ReadPostRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData     `thrift:",5,omitempty"`
}

type RetReward struct {
	Reward_type  int16   `thrift:",1,omitempty"`
	Reward_id    int32   `thrift:",2,omitempty"`
	Reward_value float64 `thrift:",3,omitempty"`
}

type ReviewDataList struct {
	Music_idx  int32   `thrift:",1,omitempty"`
	Point      float64 `thrift:",2,omitempty"`
	Difficulty int16   `thrift:",3,omitempty"`
}

type RewardList struct {
	I_id             int64   `thrift:",1,omitempty"`
	I_RewardType     int32   `thrift:",2,omitempty"`
	I_RewardId       int32   `thrift:",3,omitempty"`
	D_RewardQuantity float64 `thrift:",4,omitempty"`
}

type RewardListData struct {
	I_id             int64   `thrift:",1,omitempty"`
	I_RewardType     int32   `thrift:",2,omitempty"`
	I_RewardId       int32   `thrift:",3,omitempty"`
	D_RewardQuantity float64 `thrift:",4,omitempty"`
}

type SaveUserAchievement struct {
	I_id       int64   `thrift:",1,omitempty"`
	D_Quantity float64 `thrift:",2,omitempty"`
	S_Quantity string  `thrift:",3,omitempty"`
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

type SaveUserCharacter struct {
	I_id         int64 `thrift:",1,omitempty"`
	I_Level      int64 `thrift:",2,omitempty"`
	I_BonusLevel int64 `thrift:",3,omitempty"`
}

type SaveUserCostume struct {
	I_id         int64 `thrift:",1,omitempty"`
	I_Level      int64 `thrift:",2,omitempty"`
	I_BonusLevel int64 `thrift:",3,omitempty"`
}

type SaveUserDailyMission struct {
	I_id       int64 `thrift:",1,omitempty"`
	D_Quantity int64 `thrift:",2,omitempty"`
}

type SaveUserEventPoint struct {
	S_EventType string `thrift:",1,omitempty"`
	I_DataID    int64  `thrift:",2,omitempty"`
	I_Point     int64  `thrift:",3,omitempty"`
	I_Step      int64  `thrift:",4,omitempty"`
	I_Version   int32  `thrift:",5,omitempty"`
}

type SaveUserFollower struct {
	I_id         int64 `thrift:",1,omitempty"`
	I_Level      int64 `thrift:",2,omitempty"`
	I_BonusLevel int64 `thrift:",3,omitempty"`
}

type SaveUserFollowerQuest struct {
	I_CurrentID       int32   `thrift:",1,omitempty"`
	D_ConditionValue1 float64 `thrift:",2,omitempty"`
	D_ConditionValue2 float64 `thrift:",3,omitempty"`
	D_ConditionValue3 float64 `thrift:",4,omitempty"`
}

type SaveUserGuitar struct {
	I_id         int64 `thrift:",1,omitempty"`
	I_Level      int64 `thrift:",2,omitempty"`
	I_BonusLevel int64 `thrift:",3,omitempty"`
}

type SaveUserInfo struct {
	U_like                float64 `thrift:",1,omitempty"`
	U_fans                int64   `thrift:",2,omitempty"`
	U_fans_grade          int16   `thrift:",3,omitempty"`
	U_selected_costume_id int64   `thrift:",4,omitempty"`
	U_selected_music_id   int64   `thrift:",5,omitempty"`
}

type SaveUserMessenger struct {
	I_MessengerChatRoomId int64  `thrift:",1,omitempty"`
	I_LastConfirmIndex    int64  `thrift:",2,omitempty"`
	S_UnlockGroupList     string `thrift:",3,omitempty"`
	L_UpdateTimeTicks     int64  `thrift:",4,omitempty"`
}

type SaveUserMusic struct {
	I_id                    int64 `thrift:",1,omitempty"`
	I_Level                 int64 `thrift:",2,omitempty"`
	I_BonusLevel            int64 `thrift:",3,omitempty"`
	B_EncoreBonusAppear     int64 `thrift:",4,omitempty"`
	I_EncoreBonusFollowerId int64 `thrift:",5,omitempty"`
}

type SaveUserSkill struct {
	I_id               int64 `thrift:",1,omitempty"`
	B_Activate         int16 `thrift:",2,omitempty"`
	L_ActivateOnTicks  int64 `thrift:",3,omitempty"`
	L_ActivateOffTicks int64 `thrift:",4,omitempty"`
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

type ServerTimeRet struct {
	Time     int32 `thrift:",1,omitempty"`
	Datetime int64 `thrift:",2,omitempty"`
}

type SetAdReward struct {
	Call        string              `thrift:",1,omitempty"`
	Data        SetAdRewardDataInfo `thrift:",2,omitempty"`
	Common_data ParamData           `thrift:",3,omitempty"`
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

type SetAdRewardRetDataInfo struct {
	I_id                  int32               `thrift:",1,omitempty"`
	User_ad_list          UserAdList          `thrift:",2,omitempty"`
	Reward_data           []any               `thrift:",3,omitempty"` // TODO
	User_follower_profile UserFollowerProfile `thrift:",4,omitempty"`
}

type SetAdRewardReturn struct {
	Error       ErrorRetCode           `thrift:",1,omitempty"`
	Server_time ServerTimeRet          `thrift:",2,omitempty"`
	Mode        string                 `thrift:",3,omitempty"`
	Call        string                 `thrift:",4,omitempty"`
	Data        SetAdRewardRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData        `thrift:",6,omitempty"`
}

type SetAttendance struct {
	Call        string                `thrift:",1,omitempty"`
	Data        SetAttendanceDataInfo `thrift:",2,omitempty"`
	Common_data ParamData             `thrift:",3,omitempty"`
}

type SetAttendanceDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Type        string `thrift:",5,omitempty"`
}

type SetAttendanceRetDataInfo struct {
	Status                          string            `thrift:",1,omitempty"`
	User_follower_quest             UserFollowerQuest `thrift:",2,omitempty"`
	Attendance_count                int32             `thrift:",3,omitempty"`
	Attendance_date                 int32             `thrift:",4,omitempty"`
	Max_coutinuous_attendance_count int32             `thrift:",5,omitempty"`
}

type SetAttendanceReturn struct {
	Error       ErrorRetCode             `thrift:",1,omitempty"`
	Server_time ServerTimeRet            `thrift:",2,omitempty"`
	Mode        string                   `thrift:",3,omitempty"`
	Call        string                   `thrift:",4,omitempty"`
	Data        SetAttendanceRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData          `thrift:",6,omitempty"`
}

type SetBookmark struct {
	Call string              `thrift:",1,omitempty"`
	Data SetBookmarkDataInfo `thrift:",2,omitempty"`
}

type SetBookmarkDataInfo struct {
	U_seq         int32  `thrift:",1,omitempty"`
	U_id          string `thrift:",2,omitempty"`
	Uuid          string `thrift:",3,omitempty"`
	Device_uuid   string `thrift:",4,omitempty"`
	Bookmark_list []any  `thrift:",5,omitempty"` // TODO
}

type SetBookmarkDataList struct {
	Music_idx int32 `thrift:",1,omitempty"`
	Flag      int16 `thrift:",2,omitempty"`
}

type SetBookmarkRetDataInfo struct {
	Music_idx int32 `thrift:",1,omitempty"`
	Flag      int16 `thrift:",2,omitempty"`
}

type SetBookmarkReturn struct {
	Error       ErrorRetCode    `thrift:",1,omitempty"`
	Server_time ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string          `thrift:",3,omitempty"`
	Call        string          `thrift:",4,omitempty"`
	Data        []any           `thrift:",5,omitempty"` // TODO
	Maintenance MaintenanceData `thrift:",6,omitempty"`
}

type SetEventReward struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        SetEventRewardDataInfo `thrift:",2,omitempty"`
	Common_data ParamData              `thrift:",3,omitempty"`
}

type SetEventRewardDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Event_idx   int32  `thrift:",5,omitempty"`
}

type SetEventRewardItem struct {
	Item_type int32 `thrift:",1,omitempty"`
	Item_id   int32 `thrift:",2,omitempty"`
	Count     int64 `thrift:",3,omitempty"`
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

type SetEventRewardReturn struct {
	Error       ErrorRetCode              `thrift:",1,omitempty"`
	Server_time ServerTimeRet             `thrift:",2,omitempty"`
	Mode        string                    `thrift:",3,omitempty"`
	Call        string                    `thrift:",4,omitempty"`
	Data        SetEventRewardRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData           `thrift:",6,omitempty"`
}

type SetFollowerProfileGift struct {
	Call        string                         `thrift:",1,omitempty"`
	Data        SetFollowerProfileGiftDataInfo `thrift:",2,omitempty"`
	Common_data ParamData                      `thrift:",3,omitempty"`
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

type SetFollowerProfileGiftRetDataInfo struct {
	I_gift_type            int32                `thrift:",1,omitempty"`
	User_follower_giftitem UserFollowerGiftItem `thrift:",2,omitempty"`
	User_follower_profile  UserFollowerProfile  `thrift:",3,omitempty"`
}

type SetFollowerProfileGiftReturn struct {
	Error       ErrorRetCode                      `thrift:",1,omitempty"`
	Server_time ServerTimeRet                     `thrift:",2,omitempty"`
	Mode        string                            `thrift:",3,omitempty"`
	Call        string                            `thrift:",4,omitempty"`
	Data        SetFollowerProfileGiftRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData                   `thrift:",6,omitempty"`
}

type SetFollowerQuestInfinite struct {
	Call        string                           `thrift:",1,omitempty"`
	Data        SetFollowerQuestInfiniteDataInfo `thrift:",2,omitempty"`
	Common_data ParamData                        `thrift:",3,omitempty"`
}

type SetFollowerQuestInfiniteDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type SetFollowerQuestInfiniteRetDataInfo struct {
	U_seq               int32             `thrift:",1,omitempty"`
	User_follower_quest UserFollowerQuest `thrift:",2,omitempty"`
}

type SetFollowerQuestInfiniteReturn struct {
	Error       ErrorRetCode                        `thrift:",1,omitempty"`
	Server_time ServerTimeRet                       `thrift:",2,omitempty"`
	Mode        string                              `thrift:",3,omitempty"`
	Call        string                              `thrift:",4,omitempty"`
	Data        SetFollowerQuestInfiniteRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData                     `thrift:",6,omitempty"`
}

type SetGameReward struct {
	Call        string                `thrift:",1,omitempty"`
	Data        SetGameRewardDataInfo `thrift:",2,omitempty"`
	Common_data ParamData             `thrift:",3,omitempty"`
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

type SetGameRewardRetDataInfo struct {
	Type                  string              `thrift:",1,omitempty"`
	Id                    int32               `thrift:",2,omitempty"`
	Level                 int16               `thrift:",3,omitempty"`
	Reward_type           string              `thrift:",4,omitempty"`
	Reward_value          int64               `thrift:",5,omitempty"`
	Status                string              `thrift:",6,omitempty"`
	User_follower_profile UserFollowerProfile `thrift:",7,omitempty"`
}

type SetGameRewardReturn struct {
	Error       ErrorRetCode             `thrift:",1,omitempty"`
	Server_time ServerTimeRet            `thrift:",2,omitempty"`
	Mode        string                   `thrift:",3,omitempty"`
	Call        string                   `thrift:",4,omitempty"`
	Data        SetGameRewardRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData          `thrift:",6,omitempty"`
}

type SetPassReward struct {
	Call        string                `thrift:",1,omitempty"`
	Data        SetPassRewardDataInfo `thrift:",2,omitempty"`
	Common_data ParamData             `thrift:",3,omitempty"`
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

type SetPassRewardRetDataInfo struct {
	Subscribe_pass_reward UserSubscribePassReward `thrift:",1,omitempty"`
	Reward_data           []any                   `thrift:",2,omitempty"` // TODO
}

type SetPassRewardReturn struct {
	Error       ErrorRetCode             `thrift:",1,omitempty"`
	Server_time ServerTimeRet            `thrift:",2,omitempty"`
	Mode        string                   `thrift:",3,omitempty"`
	Call        string                   `thrift:",4,omitempty"`
	Data        SetPassRewardRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData          `thrift:",6,omitempty"`
}

type SetReviewPopup struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        SetReviewPopupDataInfo `thrift:",2,omitempty"`
	Common_data ParamData              `thrift:",3,omitempty"`
}

type SetReviewPopupDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type SetReviewPopupRetDataInfo struct {
	Status string `thrift:",1,omitempty"`
}

type SetReviewPopupReturn struct {
	Error       ErrorRetCode              `thrift:",1,omitempty"`
	Server_time ServerTimeRet             `thrift:",2,omitempty"`
	Mode        string                    `thrift:",3,omitempty"`
	Call        string                    `thrift:",4,omitempty"`
	Data        SetReviewPopupRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData           `thrift:",6,omitempty"`
}

type SetSelectReward struct {
	Call        string                  `thrift:",1,omitempty"`
	Data        SetSelectRewardDataInfo `thrift:",2,omitempty"`
	Common_data ParamData               `thrift:",3,omitempty"`
}

type SetSelectRewardDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	I_id        int32  `thrift:",5,omitempty"`
}

type SetSelectRewardInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	I_id        int32  `thrift:",5,omitempty"`
}

type SetSelectRewardRetDataInfo struct {
	U_seq                   int32 `thrift:",1,omitempty"`
	User_select_reward_list []any `thrift:",2,omitempty"` // TODO
}

type SetSelectRewardReturn struct {
	Error       ErrorRetCode               `thrift:",1,omitempty"`
	Server_time ServerTimeRet              `thrift:",2,omitempty"`
	Mode        string                     `thrift:",3,omitempty"`
	Call        string                     `thrift:",4,omitempty"`
	Data        SetSelectRewardRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData            `thrift:",6,omitempty"`
}

type SetSubscribe struct {
	Call        string               `thrift:",1,omitempty"`
	Data        SetSubscribeDataInfo `thrift:",2,omitempty"`
	Common_data ParamData            `thrift:",3,omitempty"`
}

type SetSubscribeDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	I_ids       []any  `thrift:",5,omitempty"` // TODO
}

type SetSubscribeRetDataInfo struct {
	U_seq               int32 `thrift:",1,omitempty"`
	User_subscribe_list []any `thrift:",2,omitempty"` // TODO
}

type SetSubscribeReturn struct {
	Error       ErrorRetCode            `thrift:",1,omitempty"`
	Server_time ServerTimeRet           `thrift:",2,omitempty"`
	Mode        string                  `thrift:",3,omitempty"`
	Call        string                  `thrift:",4,omitempty"`
	Data        SetSubscribeRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData         `thrift:",6,omitempty"`
}

type SetTutorial struct {
	Call        string              `thrift:",1,omitempty"`
	Data        SetTutorialDataInfo `thrift:",2,omitempty"`
	Common_data ParamData           `thrift:",3,omitempty"`
}

type SetTutorialDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Step        string `thrift:",5,omitempty"`
}

type SetTutorialNew struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        SetTutorialNewDataInfo `thrift:",2,omitempty"`
	Common_data ParamData              `thrift:",3,omitempty"`
}

type SetTutorialNewDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	I_ids       []any  `thrift:",5,omitempty"` // TODO
}

type SetTutorialNewRetDataInfo struct {
	U_seq    int32 `thrift:",1,omitempty"`
	Tutorial []any `thrift:",2,omitempty"` // TODO
}

type SetTutorialNewReturn struct {
	Error       ErrorRetCode              `thrift:",1,omitempty"`
	Server_time ServerTimeRet             `thrift:",2,omitempty"`
	Mode        string                    `thrift:",3,omitempty"`
	Call        string                    `thrift:",4,omitempty"`
	Data        SetTutorialNewRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData           `thrift:",6,omitempty"`
}

type SetTutorialRetDataInfo struct {
	Step string `thrift:",1,omitempty"`
}

type SetTutorialReturn struct {
	Error       ErrorRetCode           `thrift:",1,omitempty"`
	Server_time ServerTimeRet          `thrift:",2,omitempty"`
	Mode        string                 `thrift:",3,omitempty"`
	Call        string                 `thrift:",4,omitempty"`
	Data        SetTutorialRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData        `thrift:",6,omitempty"`
}

type SetUserFollowerProfileReward struct {
	Call        string                               `thrift:",1,omitempty"`
	Data        SetUserFollowerProfileRewardDataInfo `thrift:",2,omitempty"`
	Common_data ParamData                            `thrift:",3,omitempty"`
}

type SetUserFollowerProfileRewardDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	I_id        int32  `thrift:",5,omitempty"`
	S_level     string `thrift:",6,omitempty"`
}

type SetUserFollowerProfileRewardRetDataInfo struct {
	I_id                  int32               `thrift:",1,omitempty"`
	I_level               int32               `thrift:",2,omitempty"`
	Reward_data           []any               `thrift:",3,omitempty"` // TODO
	User_follower_profile UserFollowerProfile `thrift:",4,omitempty"`
}

type SetUserFollowerProfileRewardReturn struct {
	Error       ErrorRetCode    `thrift:",1,omitempty"`
	Server_time ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string          `thrift:",3,omitempty"`
	Call        string          `thrift:",4,omitempty"`
	Data        []any           `thrift:",5,omitempty"` // TODO
	Maintenance MaintenanceData `thrift:",6,omitempty"`
}

type SetUserFollowerQuest struct {
	Call        string                       `thrift:",1,omitempty"`
	Data        SetUserFollowerQuestDataInfo `thrift:",2,omitempty"`
	Common_data ParamData                    `thrift:",3,omitempty"`
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

type SetUserFollowerQuestReturn struct {
	Error       ErrorRetCode                    `thrift:",1,omitempty"`
	Server_time ServerTimeRet                   `thrift:",2,omitempty"`
	Mode        string                          `thrift:",3,omitempty"`
	Call        string                          `thrift:",4,omitempty"`
	Data        SetUserFollowerQuestRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData                 `thrift:",6,omitempty"`
}

type SetUserGP struct {
	Call string            `thrift:",1,omitempty"`
	Data SetUserGPDataInfo `thrift:",2,omitempty"`
}

type SetUserGPDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Gp          int32  `thrift:",5,omitempty"`
}

type SetUserGPRetDataInfo struct {
	Gp int32 `thrift:",1,omitempty"`
}

type SetUserGPReturn struct {
	Error       ErrorRetCode         `thrift:",1,omitempty"`
	Server_time ServerTimeRet        `thrift:",2,omitempty"`
	Mode        string               `thrift:",3,omitempty"`
	Call        string               `thrift:",4,omitempty"`
	Data        SetUserGPRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData      `thrift:",6,omitempty"`
}

type StoreList struct {
	Call string            `thrift:",1,omitempty"`
	Data StoreListDataInfo `thrift:",2,omitempty"`
}

type StoreListData struct {
	Store_idx int16 `thrift:",1,omitempty"`
	Music_idx int32 `thrift:",2,omitempty"`
}

type StoreListDataInfo struct {
	U_seq int32  `thrift:",1,omitempty"`
	U_id  string `thrift:",2,omitempty"`
	Uuid  string `thrift:",3,omitempty"`
}

type StoreListRetDataInfo struct {
	Store_idx  int16       `thrift:",1,omitempty"`
	Store_type int16       `thrift:",2,omitempty"`
	Title      string      `thrift:",3,omitempty"`
	Sub_title  string      `thrift:",4,omitempty"`
	Img_url    string      `thrift:",5,omitempty"`
	Store_list map[any]any `thrift:",6,omitempty"` // TODO
}

type StoreListReturn struct {
	Error       ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string          `thrift:",2,omitempty"`
	Call        string          `thrift:",3,omitempty"`
	Data        map[any]any     `thrift:",4,omitempty"` // TODO
	Maintenance MaintenanceData `thrift:",5,omitempty"`
}

type StoreNotice struct {
	Call string              `thrift:",1,omitempty"`
	Data StoreNoticeDataInfo `thrift:",2,omitempty"`
}

type StoreNoticeDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type StoreNoticeRetDataInfo struct {
	Idx            int32  `thrift:",1,omitempty"`
	Product_id_aos int64  `thrift:",2,omitempty"`
	Product_id_ios int64  `thrift:",3,omitempty"`
	Front_img_url  string `thrift:",4,omitempty"`
	Back_img_url   string `thrift:",5,omitempty"`
	Main_desc      string `thrift:",6,omitempty"`
	Sub_desc       string `thrift:",7,omitempty"`
}

type StoreNoticeReturn struct {
	Error       ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string          `thrift:",2,omitempty"`
	Call        string          `thrift:",3,omitempty"`
	Data        []any           `thrift:",4,omitempty"` // TODO
	Maintenance MaintenanceData `thrift:",5,omitempty"`
}

type TabListData struct {
	Tap_idx   int16 `thrift:",1,omitempty"`
	Music_idx int32 `thrift:",2,omitempty"`
}

type TitleList struct {
	Call string            `thrift:",1,omitempty"`
	Data TitleListDataInfo `thrift:",2,omitempty"`
}

type TitleListDataInfo struct {
	Type string `thrift:",1,omitempty"`
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

type TitleListReturn struct {
	Error       ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string          `thrift:",2,omitempty"`
	Call        string          `thrift:",3,omitempty"`
	Data        []any           `thrift:",4,omitempty"` // TODO
	Maintenance MaintenanceData `thrift:",5,omitempty"`
}

type TotalTopRank struct {
}

type TourMain struct {
	Call string           `thrift:",1,omitempty"`
	Data TourMainDataInfo `thrift:",2,omitempty"`
}

type TourMainDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type TourMainRetDataInfo struct {
	Idx         int32       `thrift:",1,omitempty"`
	Title       string      `thrift:",2,omitempty"`
	Start_time  int64       `thrift:",3,omitempty"`
	End_time    int64       `thrift:",4,omitempty"`
	Server_time int64       `thrift:",5,omitempty"`
	Mode        string      `thrift:",6,omitempty"`
	Tour_track  []any       `thrift:",7,omitempty"` // TODO
	Tour_rank   TourRank    `thrift:",8,omitempty"`
	Track_rank  map[any]any `thrift:",9,omitempty"` // TODO
	Utc         int16       `thrift:",10,omitempty"`
}

type TourMainReturn struct {
	Error       ErrorRetCode        `thrift:",1,omitempty"`
	Mode        string              `thrift:",2,omitempty"`
	Call        string              `thrift:",3,omitempty"`
	Data        TourMainRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData     `thrift:",5,omitempty"`
}

type TourMusicList struct {
	Music_idx int32 `thrift:",1,omitempty"`
}

type TourMyRank struct {
	Rank  int64 `thrift:",1,omitempty"`
	Score int64 `thrift:",2,omitempty"`
	Medal []any `thrift:",3,omitempty"` // TODO
}

type TourRank struct {
	My_rank  TourMyRank `thrift:",1,omitempty"`
	Top_rank []any      `thrift:",2,omitempty"` // TODO
}

type TourRankData struct {
	Track []any        `thrift:",1,omitempty"` // TODO
	Total TotalTopRank `thrift:",2,omitempty"`
}

type TourTrackList struct {
	Idx         int32  `thrift:",1,omitempty"`
	Tour_idx    int32  `thrift:",2,omitempty"`
	My_rank     int64  `thrift:",3,omitempty"`
	Music_idx   string `thrift:",4,omitempty"`
	Start_time  int32  `thrift:",5,omitempty"`
	End_time    int32  `thrift:",6,omitempty"`
	Open_flg    int16  `thrift:",7,omitempty"`
	Difficulty1 int16  `thrift:",8,omitempty"`
	Difficulty2 int16  `thrift:",9,omitempty"`
	Difficulty3 int16  `thrift:",10,omitempty"`
	Difficulty4 int16  `thrift:",11,omitempty"`
}

type TrackMyRank struct {
	Score int64 `thrift:",1,omitempty"`
	Medal int16 `thrift:",2,omitempty"`
}

type TrackMyScore struct {
	Score int64 `thrift:",1,omitempty"`
	Medal int16 `thrift:",2,omitempty"`
}

type TrackRank struct {
	My_score map[any]any `thrift:",1,omitempty"` // TODO
}

type TrackTopRank struct {
	U_seq     int64  `thrift:",1,omitempty"`
	U_id      string `thrift:",2,omitempty"`
	U_nick    string `thrift:",3,omitempty"`
	U_country string `thrift:",4,omitempty"`
	Rank      int64  `thrift:",5,omitempty"`
	Score     int64  `thrift:",6,omitempty"`
}

type TransferUser struct {
	Call string               `thrift:",1,omitempty"`
	Data TransferUserDataInfo `thrift:",2,omitempty"`
}

type TransferUserDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Email       string `thrift:",5,omitempty"`
}

type TransferUserRetDataInfo struct {
	Transfer_id string `thrift:",1,omitempty"`
}

type TransferUserReturn struct {
	Error       ErrorRetCode            `thrift:",1,omitempty"`
	Mode        string                  `thrift:",2,omitempty"`
	Call        string                  `thrift:",3,omitempty"`
	Data        TransferUserRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData         `thrift:",5,omitempty"`
}

type UpdateAchievement struct {
	Call string                    `thrift:",1,omitempty"`
	Data UpdateAchievementDataInfo `thrift:",2,omitempty"`
}

type UpdateAchievementDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Achievement int16  `thrift:",5,omitempty"`
}

type UpdateAchievementRetDataInfo struct {
	Achievement int16 `thrift:",1,omitempty"`
}

type UpdateAchievementReturn struct {
	Error       ErrorRetCode                 `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        UpdateAchievementRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData              `thrift:",5,omitempty"`
}

type UpdateAvatar struct {
	Call string               `thrift:",1,omitempty"`
	Data UpdateAvatarDataInfo `thrift:",2,omitempty"`
}

type UpdateAvatarDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	U_avatar    int16  `thrift:",5,omitempty"`
}

type UpdateAvatarRetDataInfo struct {
	U_avatar int16 `thrift:",1,omitempty"`
}

type UpdateAvatarReturn struct {
	Error       ErrorRetCode            `thrift:",1,omitempty"`
	Mode        string                  `thrift:",2,omitempty"`
	Call        string                  `thrift:",3,omitempty"`
	Data        UpdateAvatarRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData         `thrift:",5,omitempty"`
}

type UpdateNickName struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        UpdateNickNameDataInfo `thrift:",2,omitempty"`
	Common_data ParamData              `thrift:",3,omitempty"`
}

type UpdateNickNameDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Nickname    string `thrift:",5,omitempty"`
}

type UpdateNickNameRetDataInfo struct {
	Nickname string `thrift:",1,omitempty"`
}

type UpdateNickNameReturn struct {
	Error       ErrorRetCode              `thrift:",1,omitempty"`
	Mode        string                    `thrift:",2,omitempty"`
	Call        string                    `thrift:",3,omitempty"`
	Data        UpdateNickNameRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData           `thrift:",5,omitempty"`
}

type UpdateScore struct {
	Call string              `thrift:",1,omitempty"`
	Data UpdateScoreDataInfo `thrift:",2,omitempty"`
}

type UpdateScoreDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Score_list  []any  `thrift:",5,omitempty"` // TODO
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

type UpdateScoreReturn struct {
	Error       ErrorRetCode           `thrift:",1,omitempty"`
	Server_time ServerTimeRet          `thrift:",2,omitempty"`
	Mode        string                 `thrift:",3,omitempty"`
	Call        string                 `thrift:",4,omitempty"`
	Data        UpdateScoreRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData        `thrift:",6,omitempty"`
}

type UpdateUserTitle struct {
	Call string                  `thrift:",1,omitempty"`
	Data UpdateUserTitleDataInfo `thrift:",2,omitempty"`
}

type UpdateUserTitleDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	U_title     int16  `thrift:",5,omitempty"`
}

type UpdateUserTitleRetDataInfo struct {
	U_title int16 `thrift:",1,omitempty"`
}

type UpdateUserTitleReturn struct {
	Error       ErrorRetCode               `thrift:",1,omitempty"`
	Mode        string                     `thrift:",2,omitempty"`
	Call        string                     `thrift:",3,omitempty"`
	Data        UpdateUserTitleRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData            `thrift:",5,omitempty"`
}

type UseCoupon struct {
	Call string            `thrift:",1,omitempty"`
	Data UseCouponDataInfo `thrift:",2,omitempty"`
}

type UseCouponDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Send_ppid   string `thrift:",5,omitempty"`
	Coupon      string `thrift:",6,omitempty"`
}

type UseCouponRetDataInfo struct {
	Ret string `thrift:",1,omitempty"`
}

type UseCouponReturn struct {
	Error       ErrorRetCode         `thrift:",1,omitempty"`
	Mode        string               `thrift:",2,omitempty"`
	Call        string               `thrift:",3,omitempty"`
	Data        UseCouponRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData      `thrift:",5,omitempty"`
}

type UseCredit struct {
	Call string            `thrift:",1,omitempty"`
	Data UseCreditDataInfo `thrift:",2,omitempty"`
}

type UseCreditDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Credit      int64  `thrift:",5,omitempty"`
}

type UseCreditRetDataInfo struct {
	U_gold   int64 `thrift:",1,omitempty"`
	U_cp     int64 `thrift:",2,omitempty"`
	U_mp     int64 `thrift:",3,omitempty"`
	U_credit int64 `thrift:",4,omitempty"`
}

type UseCreditReturn struct {
	Error       ErrorRetCode         `thrift:",1,omitempty"`
	Mode        string               `thrift:",2,omitempty"`
	Call        string               `thrift:",3,omitempty"`
	Data        UseCreditRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData      `thrift:",5,omitempty"`
}

type UserAchievement struct {
	I_id       int64   `thrift:",1,omitempty"`
	I_Level    int64   `thrift:",2,omitempty"`
	D_Quantity float64 `thrift:",3,omitempty"`
	S_Quantity string  `thrift:",4,omitempty"`
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

type UserAvatarList struct {
	Call string                 `thrift:",1,omitempty"`
	Data UserAvatarListDataInfo `thrift:",2,omitempty"`
}

type UserAvatarListDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type UserAvatarListRetDataInfo struct {
	U_seq  int32 `thrift:",1,omitempty"`
	Avatar []any `thrift:",2,omitempty"` // TODO
}

type UserAvatarListReturn struct {
	Error       ErrorRetCode              `thrift:",1,omitempty"`
	Mode        string                    `thrift:",2,omitempty"`
	Call        string                    `thrift:",3,omitempty"`
	Data        UserAvatarListRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData           `thrift:",5,omitempty"`
}

type UserBuff struct {
	I_id         int64 `thrift:",1,omitempty"`
	I_Level      int64 `thrift:",2,omitempty"`
	I_ActiveTime int64 `thrift:",3,omitempty"`
	I_EndTime    int64 `thrift:",4,omitempty"`
}

type UserCandyShop struct {
	I_id              int64   `thrift:",1,omitempty"`
	I_CurrentBuyCount int64   `thrift:",2,omitempty"`
	I_TotalBuyCount   int64   `thrift:",3,omitempty"`
	L_LastBuyTick     float64 `thrift:",4,omitempty"`
	Upd_day           int64   `thrift:",5,omitempty"`
}

type UserCharacter struct {
	I_id         int64 `thrift:",1,omitempty"`
	I_Level      int64 `thrift:",2,omitempty"`
	I_BonusLevel int64 `thrift:",3,omitempty"`
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

type UserCostume struct {
	I_id         int64 `thrift:",1,omitempty"`
	I_Level      int64 `thrift:",2,omitempty"`
	I_BonusLevel int64 `thrift:",3,omitempty"`
}

type UserDailyMission struct {
	I_id       int64  `thrift:",1,omitempty"`
	I_Level    int64  `thrift:",2,omitempty"`
	D_Quantity int64  `thrift:",3,omitempty"`
	Upd_date   string `thrift:",4,omitempty"`
}

type UserData struct {
	U_seq int64 `thrift:",1,omitempty"`
	U_mp  int64 `thrift:",2,omitempty"`
}

type UserDel struct {
	Call        string          `thrift:",1,omitempty"`
	Data        UserDelDataInfo `thrift:",2,omitempty"`
	Common_data ParamData       `thrift:",3,omitempty"`
}

type UserDelDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type UserDelRetDataInfo struct {
	Result string `thrift:",1,omitempty"`
}

type UserDelReturn struct {
	Error       ErrorRetCode       `thrift:",1,omitempty"`
	Mode        string             `thrift:",2,omitempty"`
	Call        string             `thrift:",3,omitempty"`
	Data        UserDelRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData    `thrift:",5,omitempty"`
}

type UserEventPoint struct {
	S_EventType  string `thrift:",1,omitempty"`
	I_DataID     int64  `thrift:",2,omitempty"`
	I_Point      int64  `thrift:",3,omitempty"`
	I_Step       int64  `thrift:",4,omitempty"`
	I_ADViewTime int64  `thrift:",5,omitempty"`
	I_Version    int32  `thrift:",6,omitempty"`
}

type UserFollower struct {
	I_id         int64 `thrift:",1,omitempty"`
	I_Level      int64 `thrift:",2,omitempty"`
	I_BonusLevel int64 `thrift:",3,omitempty"`
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

type UserGuitar struct {
	I_id         int64 `thrift:",1,omitempty"`
	I_Level      int64 `thrift:",2,omitempty"`
	I_BonusLevel int64 `thrift:",3,omitempty"`
}

type UserInfo struct {
	U_like                float64 `thrift:",1,omitempty"`
	U_fans                int64   `thrift:",2,omitempty"`
	U_fans_grade          int16   `thrift:",3,omitempty"`
	U_selected_costume_id int64   `thrift:",4,omitempty"`
	U_selected_music_id   int64   `thrift:",5,omitempty"`
}

type UserItemInvenList struct {
	Item_list []any `thrift:",1,omitempty"` // TODO
}

type UserItemList struct {
	Item_type int32 `thrift:",1,omitempty"`
	Item_id   int32 `thrift:",2,omitempty"`
	Count     int64 `thrift:",3,omitempty"`
}

type UserJoin struct {
	Call        string           `thrift:",1,omitempty"`
	Data        UserJoinDataInfo `thrift:",2,omitempty"`
	Common_data ParamData        `thrift:",3,omitempty"`
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

type UserJoinRetDataInfo struct {
	U_seq int32  `thrift:",1,omitempty"`
	U_id  string `thrift:",2,omitempty"`
}

type UserJoinReturn struct {
	Error       ErrorRetCode        `thrift:",1,omitempty"`
	Server_time ServerTimeRet       `thrift:",2,omitempty"`
	Mode        string              `thrift:",3,omitempty"`
	Call        string              `thrift:",4,omitempty"`
	Data        UserJoinRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData     `thrift:",6,omitempty"`
}

type UserLevel struct {
	U_girl_level     int64 `thrift:",1,omitempty"`
	U_skill_level    int64 `thrift:",2,omitempty"`
	U_mate_level     int64 `thrift:",3,omitempty"`
	U_follower_level int64 `thrift:",4,omitempty"`
	U_play_level     int64 `thrift:",5,omitempty"`
}

type UserLoad struct {
	Call        string           `thrift:",1,omitempty"`
	Data        UserLoadDataInfo `thrift:",2,omitempty"`
	Common_data ParamData        `thrift:",3,omitempty"`
}

type UserLoadDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Type        string `thrift:",5,omitempty"`
}

type UserLoadRetDataInfo struct {
	User          UserData         `thrift:",1,omitempty"`
	Area_data     map[any]any      `thrift:",2,omitempty"` // TODO
	User_contents UserContentsData `thrift:",3,omitempty"`
}

type UserLoadReturn struct {
	Error       ErrorRetCode        `thrift:",1,omitempty"`
	Server_time ServerTimeRet       `thrift:",2,omitempty"`
	Mode        string              `thrift:",3,omitempty"`
	Call        string              `thrift:",4,omitempty"`
	Data        UserLoadRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData     `thrift:",6,omitempty"`
}

type UserLogin struct {
	Call        string            `thrift:",1,omitempty"`
	Data        UserLoginDataInfo `thrift:",2,omitempty"`
	Common_data ParamData         `thrift:",3,omitempty"`
}

type UserLoginDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	U_state     string `thrift:",5,omitempty"`
}

type UserLoginRetDataInfo struct {
	User          UserData         `thrift:",1,omitempty"`
	Area_data     map[any]any      `thrift:",2,omitempty"` // TODO
	User_contents UserContentsData `thrift:",3,omitempty"`
}

type UserLoginReturn struct {
	Error       ErrorRetCode         `thrift:",1,omitempty"`
	Server_time ServerTimeRet        `thrift:",2,omitempty"`
	Mode        string               `thrift:",3,omitempty"`
	Call        string               `thrift:",4,omitempty"`
	Data        UserLoginRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData      `thrift:",6,omitempty"`
}

type UserMessenger struct {
	I_MessengerChatRoomId int64  `thrift:",1,omitempty"`
	I_LastConfirmIndex    int64  `thrift:",2,omitempty"`
	S_UnlockGroupList     string `thrift:",3,omitempty"`
	L_UpdateTimeTicks     int64  `thrift:",4,omitempty"`
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

type UserProp struct {
	I_id    int64 `thrift:",1,omitempty"`
	I_Level int64 `thrift:",2,omitempty"`
}

type UserPurchase struct {
	Call     string               `thrift:",1,omitempty"`
	Data     UserPurchaseDataInfo `thrift:",2,omitempty"`
	Sub_mode string               `thrift:",3,omitempty"`
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

type UserPurchaseErrorLog struct {
	Call string                       `thrift:",1,omitempty"`
	Data UserPurchaseErrorLogDataInfo `thrift:",2,omitempty"`
}

type UserPurchaseErrorLogDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Error_log   string `thrift:",5,omitempty"`
}

type UserPurchaseErrorLogRetDataInfo struct {
	Ret string `thrift:",1,omitempty"`
}

type UserPurchaseErrorLogReturn struct {
	Error       ErrorRetCode                    `thrift:",1,omitempty"`
	Mode        string                          `thrift:",2,omitempty"`
	Call        string                          `thrift:",3,omitempty"`
	Data        UserPurchaseErrorLogRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData                 `thrift:",5,omitempty"`
}

type UserPurchaseMusicIdx struct {
	Music_idx int32 `thrift:",1,omitempty"`
}

type UserPurchaseRetDataInfo struct {
	Pay_id      string `thrift:",1,omitempty"`
	Product_id  string `thrift:",2,omitempty"`
	Purchase_id string `thrift:",3,omitempty"`
	Music_data  []any  `thrift:",4,omitempty"` // TODO
}

type UserPurchaseReturn struct {
	Error       ErrorRetCode            `thrift:",1,omitempty"`
	Mode        string                  `thrift:",2,omitempty"`
	Call        string                  `thrift:",3,omitempty"`
	Data        UserPurchaseRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData         `thrift:",5,omitempty"`
}

type UserSave struct {
	Call        string           `thrift:",1,omitempty"`
	Data        UserSaveDataInfo `thrift:",2,omitempty"`
	Common_data ParamData        `thrift:",3,omitempty"`
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

type UserSaveRetDataInfo struct {
	Status string `thrift:",1,omitempty"`
}

type UserSaveReturn struct {
	Error       ErrorRetCode        `thrift:",1,omitempty"`
	Server_time ServerTimeRet       `thrift:",2,omitempty"`
	Mode        string              `thrift:",3,omitempty"`
	Call        string              `thrift:",4,omitempty"`
	Data        UserSaveRetDataInfo `thrift:",5,omitempty"`
	Maintenance MaintenanceData     `thrift:",6,omitempty"`
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

type UserShop struct {
	I_id           int64 `thrift:",1,omitempty"`
	I_Count        int64 `thrift:",2,omitempty"`
	I_TotalCount   int64 `thrift:",3,omitempty"`
	I_PurchaseTime int64 `thrift:",4,omitempty"`
	Upd_day        int64 `thrift:",5,omitempty"`
}

type UserSkill struct {
	I_id               int64 `thrift:",1,omitempty"`
	I_Level            int64 `thrift:",2,omitempty"`
	B_Activate         int16 `thrift:",3,omitempty"`
	L_ActivateOnTicks  int64 `thrift:",4,omitempty"`
	L_ActivateOffTicks int64 `thrift:",5,omitempty"`
}

type UserSubscribeList struct {
	I_SubscribeID int64 `thrift:",1,omitempty"`
	I_ActiveTime  int64 `thrift:",2,omitempty"`
	I_isActive    int64 `thrift:",3,omitempty"`
}

type UserSubscribePassReward struct {
	I_SubscribeID int64 `thrift:",1,omitempty"`
	I_Type        int64 `thrift:",2,omitempty"`
	I_Step        int64 `thrift:",3,omitempty"`
	I_UpdateTime  int64 `thrift:",4,omitempty"`
	I_Version     int32 `thrift:",5,omitempty"`
}

type UserTicketCollection struct {
	I_id int64 `thrift:",1,omitempty"`
}

type UserTitleList struct {
	Call string                `thrift:",1,omitempty"`
	Data UserTitleListDataInfo `thrift:",2,omitempty"`
}

type UserTitleListDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type UserTitleListRetDataInfo struct {
	U_seq int32 `thrift:",1,omitempty"`
	Title []any `thrift:",2,omitempty"` // TODO
}

type UserTitleListReturn struct {
	Error       ErrorRetCode             `thrift:",1,omitempty"`
	Mode        string                   `thrift:",2,omitempty"`
	Call        string                   `thrift:",3,omitempty"`
	Data        UserTitleListRetDataInfo `thrift:",4,omitempty"`
	Maintenance MaintenanceData          `thrift:",5,omitempty"`
}

type UserUnit struct {
	I_id    int64 `thrift:",1,omitempty"`
	I_Level int64 `thrift:",2,omitempty"`
}
