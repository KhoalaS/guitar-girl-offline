package store_model

import (
	"github.com/KhoalaS/guitar-girl-offline/pkg/model/common_model"
	"github.com/KhoalaS/guitar-girl-offline/pkg/model/user_model"
)

type BuyContentsRetDataInfo struct {
	U_cp       int64                `thrift:",1,omitempty"`
	U_candy    float64              `thrift:",2,omitempty"`
	User_skill user_model.UserSkill `thrift:",3,omitempty"`
	User_unit  user_model.UserUnit  `thrift:",4,omitempty"`
}

type RetReward struct {
	Reward_type  int16 `thrift:",1,omitempty"`
	Reward_id    int32 `thrift:",2,omitempty"`
	Reward_value int64 `thrift:",3,omitempty"`
}

type StoreList struct {
	Call string            `thrift:",1,omitempty"`
	Data StoreListDataInfo `thrift:",2,omitempty"`
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

type BuyShopReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        BuyShopRetDataInfo           `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type ItemStoreListReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        []any                        `thrift:",4,omitempty"` // TODO
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type BuyContentsReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        BuyContentsRetDataInfo       `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type GetVarietyStoreDataInfo struct {
	Device_uuid string `thrift:",1,omitempty"`
}

type BuyCheckDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Idx         int32  `thrift:",5,omitempty"`
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

type ItemStoreList struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        ItemStoreListDataInfo  `thrift:",2,omitempty"`
	Common_data common_model.ParamData `thrift:",3,omitempty"`
}

type BuyShop struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        BuyShopDataInfo        `thrift:",2,omitempty"`
	Sub_mode    string                 `thrift:",3,omitempty"`
	Common_data common_model.ParamData `thrift:",4,omitempty"`
}

type StoreNotice struct {
	Call string              `thrift:",1,omitempty"`
	Data StoreNoticeDataInfo `thrift:",2,omitempty"`
}

type BuyCheckReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        BuyCheckRetDataInfo          `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type BuyShopRetDataInfo struct {
	U_cp          int64                  `thrift:",1,omitempty"`
	U_candy       float64                `thrift:",2,omitempty"`
	Reward_data   []any                  `thrift:",3,omitempty"` // TODO
	User_ad_level user_model.UserAdLevel `thrift:",4,omitempty"`
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

type ItemStoreListDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Store_type  int16  `thrift:",5,omitempty"`
}

type StoreListReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        map[any]any                  `thrift:",4,omitempty"` // TODO
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type BuyCheck struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        BuyCheckDataInfo       `thrift:",2,omitempty"`
	Common_data common_model.ParamData `thrift:",3,omitempty"`
}

type BuyCheckRetDataInfo struct {
	Result string `thrift:",1,omitempty"`
}

type BuyVarietyStore struct {
	Call        string                  `thrift:",1,omitempty"`
	Data        BuyVarietyStoreDataInfo `thrift:",2,omitempty"`
	Common_data common_model.ParamData  `thrift:",3,omitempty"`
}

type GetVarietyStore struct {
	Call        string                  `thrift:",1,omitempty"`
	Data        GetVarietyStoreDataInfo `thrift:",2,omitempty"`
	Common_data common_model.ParamData  `thrift:",3,omitempty"`
}

type StoreListDataInfo struct {
	U_seq int32  `thrift:",1,omitempty"`
	U_id  string `thrift:",2,omitempty"`
	Uuid  string `thrift:",3,omitempty"`
}

type StoreListRetDataInfo struct {
	Idx            int32   `thrift:",1,omitempty"`
	Store_type     int16   `thrift:",2,omitempty"`
	Music_idx      int32   `thrift:",3,omitempty"`
	Banner_img_url string  `thrift:",4,omitempty"`
	Location_url   string  `thrift:",5,omitempty"`
	Album_idx      int32   `thrift:",6,omitempty"`
	Difficulty1    string  `thrift:",7,omitempty"`
	Difficulty2    string  `thrift:",8,omitempty"`
	Difficulty3    string  `thrift:",9,omitempty"`
	Buy_type       int16   `thrift:",10,omitempty"`
	Buy_ad         int16   `thrift:",11,omitempty"`
	Price          float64 `thrift:",12,omitempty"`
	Cdn_dir        string  `thrift:",13,omitempty"`
	New_status     int16   `thrift:",14,omitempty"`
	Share_game1    int16   `thrift:",15,omitempty"`
	Share_game2    int16   `thrift:",16,omitempty"`
	Share_game3    int16   `thrift:",17,omitempty"`
	Share_game4    int16   `thrift:",18,omitempty"`
	Share_game5    int16   `thrift:",19,omitempty"`
	Share_game6    int16   `thrift:",20,omitempty"`
	Share_game7    int16   `thrift:",21,omitempty"`
	Share_game8    int16   `thrift:",22,omitempty"`
	Share_game9    int16   `thrift:",23,omitempty"`
	Music_title    string  `thrift:",24,omitempty"`
	Artist         string  `thrift:",25,omitempty"`
	Songwriter     string  `thrift:",26,omitempty"`
	Lyricist       string  `thrift:",27,omitempty"`
	Description    string  `thrift:",28,omitempty"`
}

type StoreNoticeReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        []any                        `thrift:",4,omitempty"` // TODO
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type BuyContents struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        BuyContentsDataInfo    `thrift:",2,omitempty"`
	Sub_mode    string                 `thrift:",3,omitempty"`
	Common_data common_model.ParamData `thrift:",4,omitempty"`
}

type BuyContentsDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Type        string `thrift:",5,omitempty"`
	Idx         int32  `thrift:",6,omitempty"`
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
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        BuyVarietyStoreRetDataInfo   `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type GetVarietyStoreReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        []any                        `thrift:",5,omitempty"` // TODO
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
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
