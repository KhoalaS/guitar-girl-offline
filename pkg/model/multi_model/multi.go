package multi_model

import "github.com/KhoalaS/guitar-girl-offline/pkg/model/common_model"

type ChannelList struct {
	Channel        int16  `thrift:",1,omitempty"`
	Url            string `thrift:",2,omitempty"`
	Port           int64  `thrift:",3,omitempty"`
	Flg            int16  `thrift:",4,omitempty"`
	Maintenance_cn string `thrift:",5,omitempty"`
	Maintenance_en string `thrift:",6,omitempty"`
	Maintenance_jp string `thrift:",7,omitempty"`
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
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        GetServerListRetDataInfo     `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type UserData struct {
	U_seq int64 `thrift:",1,omitempty"`
	U_mp  int64 `thrift:",2,omitempty"`
}
