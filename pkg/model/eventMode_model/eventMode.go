package eventMode_model

import "github.com/KhoalaS/guitar-girl-offline/pkg/model/common_model"

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

type EventRewardListArr struct {
	Event_idx    int32 `thrift:",1,omitempty"`
	Gain_point   int64 `thrift:",2,omitempty"`
	Reward_type  int16 `thrift:",3,omitempty"`
	Reward_value int64 `thrift:",4,omitempty"`
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

type GetSamSeckListReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        GetSamSeckListRetDataInfo    `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type GetSamSeckRewardReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        GetSamSeckRewardRetDataInfo  `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type RewardListData struct {
	I_id             int64   `thrift:",1,omitempty"`
	I_RewardType     int32   `thrift:",2,omitempty"`
	I_RewardId       int32   `thrift:",3,omitempty"`
	D_RewardQuantity float64 `thrift:",4,omitempty"`
}

type EventRewardList struct {
	Call string                  `thrift:",1,omitempty"`
	Data EventRewardListDataInfo `thrift:",2,omitempty"`
}

type EventRewardListDataInfo struct {
	Event_idx   int32  `thrift:",1,omitempty"`
	Device_uuid string `thrift:",2,omitempty"`
}

type EventRewardListReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        EventRewardListRetDataInfo   `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type GetEventMain struct {
	Call string               `thrift:",1,omitempty"`
	Data GetEventMainDataInfo `thrift:",2,omitempty"`
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

type GetEventRankReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        []any                        `thrift:",4,omitempty"` // TODO
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
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

type EventRewardListRetDataInfo struct {
	Event_idx        int32 `thrift:",1,omitempty"`
	Reward_list      []any `thrift:",2,omitempty"` // TODO
	Rank_reward_list []any `thrift:",3,omitempty"` // TODO
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

type EventRewardListData struct {
	Gain_point   int64 `thrift:",1,omitempty"`
	Reward_type  int16 `thrift:",2,omitempty"`
	Reward_value int64 `thrift:",3,omitempty"`
}

type GetEventMainDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type GetEventMainReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        GetEventMainRetDataInfo      `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type GetEventRank struct {
	Call string               `thrift:",1,omitempty"`
	Data GetEventRankDataInfo `thrift:",2,omitempty"`
}

type GetSamSeckList struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        GetSamSeckListDataInfo `thrift:",2,omitempty"`
	Common_data common_model.ParamData `thrift:",3,omitempty"`
}

type GetSamSeckReward struct {
	Call        string                   `thrift:",1,omitempty"`
	Data        GetSamSeckRewardDataInfo `thrift:",2,omitempty"`
	Common_data common_model.ParamData   `thrift:",3,omitempty"`
}

type RewardList struct {
	I_id             int64   `thrift:",1,omitempty"`
	I_RewardType     int32   `thrift:",2,omitempty"`
	I_RewardId       int32   `thrift:",3,omitempty"`
	D_RewardQuantity float64 `thrift:",4,omitempty"`
}
