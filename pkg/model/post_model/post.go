package post_model

import "github.com/KhoalaS/guitar-girl-offline/pkg/model/common_model"

type ProvidePost struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        ProvidePostDataInfo    `thrift:",2,omitempty"`
	Common_data common_model.ParamData `thrift:",3,omitempty"`
}

type ProvidePostReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        []any                        `thrift:",4,omitempty"` // TODO
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type ReadPostDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Idx         int64  `thrift:",5,omitempty"`
}

type GetPostRetDataInfo struct {
	Server_time int64 `thrift:",1,omitempty"`
	Post_list   []any `thrift:",2,omitempty"` // TODO
}

type GetPostReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        GetPostRetDataInfo           `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type GetPostTimeReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Server_time common_model.ServerTimeRet   `thrift:",2,omitempty"`
	Mode        string                       `thrift:",3,omitempty"`
	Call        string                       `thrift:",4,omitempty"`
	Data        GetPostTimeRetDataInfo       `thrift:",5,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",6,omitempty"`
}

type ReadPost struct {
	Call string           `thrift:",1,omitempty"`
	Data ReadPostDataInfo `thrift:",2,omitempty"`
}

type ReadPostReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        ReadPostRetDataInfo          `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
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
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        DeletePostRetDataInfo        `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type GetPost struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        GetPostDataInfo        `thrift:",2,omitempty"`
	Common_data common_model.ParamData `thrift:",3,omitempty"`
}

type GetPostTime struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        GetPostTimeDataInfo    `thrift:",2,omitempty"`
	Common_data common_model.ParamData `thrift:",3,omitempty"`
}

type ItemList struct {
	I_RewardType     int32   `thrift:",1,omitempty"`
	I_RewardId       int32   `thrift:",2,omitempty"`
	D_RewardQuantity float64 `thrift:",3,omitempty"`
}

type ProvidePostDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Idx         int64  `thrift:",5,omitempty"`
	Type        int16  `thrift:",6,omitempty"`
}

type DeletePost struct {
	Call        string                 `thrift:",1,omitempty"`
	Data        DeletePostDataInfo     `thrift:",2,omitempty"`
	Common_data common_model.ParamData `thrift:",3,omitempty"`
}

type GetPostDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type GetPostTimeRetDataInfo struct {
	Time int64 `thrift:",1,omitempty"`
}

type ProvidePostRetDataInfo struct {
	I_RewardType     int32   `thrift:",1,omitempty"`
	I_RewardId       int32   `thrift:",2,omitempty"`
	D_RewardQuantity float64 `thrift:",3,omitempty"`
}

type ReadPostRetDataInfo struct {
	Idx int64 `thrift:",1,omitempty"`
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

type GetPostTimeDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}
