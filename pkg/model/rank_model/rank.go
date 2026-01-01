package rank_model

import "github.com/KhoalaS/guitar-girl-offline/pkg/model/common_model"

type GetMusicRankDataInfo struct {
	Music_idx int32 `thrift:",1,omitempty"`
}

type GetRankMainDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type GetRankMainReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        GetRankMainRetDataInfo       `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type GetTotalCountryRank struct {
	Call string                      `thrift:",1,omitempty"`
	Data GetTotalCountryRankDataInfo `thrift:",2,omitempty"`
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

type GetMastersRankReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        GetMastersRankRetDataInfo    `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type GetRankMainRetDataInfo struct {
	My_country_rank       GetCountryMyRankRet    `thrift:",1,omitempty"`
	Total_country_my_rank GetTotalCountryRankRet `thrift:",2,omitempty"`
	Country_rank          []any                  `thrift:",3,omitempty"` // TODO
	Masters_rank          GetMastersRankRet      `thrift:",4,omitempty"`
}

type GetTotalCountryRankReturn struct {
	Error       common_model.ErrorRetCode      `thrift:",1,omitempty"`
	Mode        string                         `thrift:",2,omitempty"`
	Call        string                         `thrift:",3,omitempty"`
	Data        GetTotalCountryRankRetDataInfo `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData   `thrift:",5,omitempty"`
}

type GetTotalMusicRankList2 struct {
	Rank  int32  `thrift:",1,omitempty"`
	Id    string `thrift:",2,omitempty"`
	Score int32  `thrift:",3,omitempty"`
}

type GetTotalMusicRankRetDataInfo struct {
	Total_rank_list []any `thrift:",1,omitempty"` // TODO
	Rank_list1      []any `thrift:",2,omitempty"` // TODO
	Rank_list2      []any `thrift:",3,omitempty"` // TODO
	Rank_list3      []any `thrift:",4,omitempty"` // TODO
	My_rank_list    []any `thrift:",5,omitempty"` // TODO
	Music_idx       int32 `thrift:",6,omitempty"`
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

type GetMastersRankDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Start_limit int32  `thrift:",5,omitempty"`
	End_limit   int32  `thrift:",6,omitempty"`
}

type GetMusicRank struct {
	Call string               `thrift:",1,omitempty"`
	Data GetMusicRankDataInfo `thrift:",2,omitempty"`
}

type GetMusicRankReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        GetMusicRankRetDataInfo      `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type GetTotalMusicRankList1 struct {
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

type GetMastersRankRetDataInfo struct {
	Start_limit int16             `thrift:",1,omitempty"`
	End_limit   int16             `thrift:",2,omitempty"`
	My_rank     GetMastersRankRet `thrift:",3,omitempty"`
	Rank_list   []any             `thrift:",4,omitempty"` // TODO
}

type GetRankMain struct {
	Call string              `thrift:",1,omitempty"`
	Data GetRankMainDataInfo `thrift:",2,omitempty"`
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

type GetTotalMusicRankList3 struct {
	Rank  int32  `thrift:",1,omitempty"`
	Id    string `thrift:",2,omitempty"`
	Score int32  `thrift:",3,omitempty"`
}

type GetCountryRank struct {
	Call string                 `thrift:",1,omitempty"`
	Data GetCountryRankDataInfo `thrift:",2,omitempty"`
}

type GetCountryRankReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        []any                        `thrift:",4,omitempty"` // TODO
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type GetTotalMusicRank struct {
	Call string                    `thrift:",1,omitempty"`
	Data GetTotalMusicRankDataInfo `thrift:",2,omitempty"`
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

type GetCountryMyRankRetDataInfo struct {
	Select_country string              `thrift:",1,omitempty"`
	Start_limit    int16               `thrift:",2,omitempty"`
	End_limit      int16               `thrift:",3,omitempty"`
	My_rank        GetCountryMyRankRet `thrift:",4,omitempty"`
	Rank_list      []any               `thrift:",5,omitempty"` // TODO
}

type GetCountryRankDataInfo struct {
}

type GetMusicRankRetDataInfo struct {
	Id    string `thrift:",1,omitempty"`
	Score int32  `thrift:",2,omitempty"`
}

type GetTotalMusicRankReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        []any                        `thrift:",4,omitempty"` // TODO
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type GetCountryMyRankReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        GetCountryMyRankRetDataInfo  `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
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

type GetTotalCountryRankDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Start_limit int32  `thrift:",5,omitempty"`
	End_limit   int32  `thrift:",6,omitempty"`
}

type GetTotalCountryRankRetDataInfo struct {
	Start_limit int16                  `thrift:",1,omitempty"`
	End_limit   int16                  `thrift:",2,omitempty"`
	My_rank     GetTotalCountryRankRet `thrift:",3,omitempty"`
	Rank_list   []any                  `thrift:",4,omitempty"` // TODO
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
	Grade   string `thrift:",1,omitempty"`
	Rank    int32  `thrift:",2,omitempty"`
	Score   int32  `thrift:",3,omitempty"`
	Lowrank int32  `thrift:",4,omitempty"`
}

type GetTotalMusicRankList struct {
	Rank    int32  `thrift:",1,omitempty"`
	Id      string `thrift:",2,omitempty"`
	Nick    string `thrift:",3,omitempty"`
	Country string `thrift:",4,omitempty"`
	Score   int32  `thrift:",5,omitempty"`
}

type GetCountryMyRank struct {
	Call string                   `thrift:",1,omitempty"`
	Data GetCountryMyRankDataInfo `thrift:",2,omitempty"`
}

type GetCountryRankRetDataInfo struct {
	Rank        int32  `thrift:",1,omitempty"`
	Country     string `thrift:",2,omitempty"`
	Gold        int32  `thrift:",3,omitempty"`
	Silver      int32  `thrift:",4,omitempty"`
	Bronze      int32  `thrift:",5,omitempty"`
	Total_score int32  `thrift:",6,omitempty"`
}

type GetMastersRank struct {
	Call string                 `thrift:",1,omitempty"`
	Data GetMastersRankDataInfo `thrift:",2,omitempty"`
}
