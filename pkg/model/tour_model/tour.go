package tour_model

import "github.com/KhoalaS/guitar-girl-offline/pkg/model/common_model"

type TourMyRank struct {
	Rank  int64 `thrift:",1,omitempty"`
	Score int64 `thrift:",2,omitempty"`
	Medal []any `thrift:",3,omitempty"` // TODO
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

type PlayCheck struct {
	Call string            `thrift:",1,omitempty"`
	Data PlayCheckDataInfo `thrift:",2,omitempty"`
}

type PlayMusicDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
	Music_idx   int32  `thrift:",5,omitempty"`
	Difficulty  int32  `thrift:",6,omitempty"`
}

type PlayMusicReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        PlayMusicRetDataInfo         `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type TrackRank struct {
	My_score map[any]any `thrift:",1,omitempty"` // TODO
}

type EndPlayMusicReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        EndPlayMusicRetDataInfo      `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type GetPostTime struct {
	Call string              `thrift:",1,omitempty"`
	Data GetPostTimeDataInfo `thrift:",2,omitempty"`
}

type GetTourRankReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        GetTourRankRetDataInfo       `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
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

type GetPostTimeRetDataInfo struct {
}

type GetTourListDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type GetTourListReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        []any                        `thrift:",4,omitempty"` // TODO
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type TourMainReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        TourMainRetDataInfo          `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type PlayCheckRetDataInfo struct {
	Basic       int32 `thrift:",1,omitempty"`
	Pro         int32 `thrift:",2,omitempty"`
	Legend      int32 `thrift:",3,omitempty"`
	Extra       int32 `thrift:",4,omitempty"`
	Server_time int64 `thrift:",5,omitempty"`
}

type PlayCheckReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        PlayCheckRetDataInfo         `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type TourRank struct {
	My_rank  TourMyRank `thrift:",1,omitempty"`
	Top_rank []any      `thrift:",2,omitempty"` // TODO
}

type TrackMyRank struct {
	Score int64 `thrift:",1,omitempty"`
	Medal int16 `thrift:",2,omitempty"`
}

type TrackTopRank struct {
	U_seq     int64  `thrift:",1,omitempty"`
	U_id      string `thrift:",2,omitempty"`
	U_nick    string `thrift:",3,omitempty"`
	U_country string `thrift:",4,omitempty"`
	Rank      int64  `thrift:",5,omitempty"`
	Score     int64  `thrift:",6,omitempty"`
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

type GetTourRankRetDataInfo struct {
	My_rank GetTourRankData `thrift:",1,omitempty"`
	Rank    []any           `thrift:",2,omitempty"` // TODO
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

type GetPostTimeReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        GetPostTimeRetDataInfo       `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type PlayMusic struct {
	Call string            `thrift:",1,omitempty"`
	Data PlayMusicDataInfo `thrift:",2,omitempty"`
}

type EndPlayMusicRetDataInfo struct {
	Code int16 `thrift:",1,omitempty"`
}

type GetTourListRetDataInfo struct {
	Tour_idx int32  `thrift:",1,omitempty"`
	Title    string `thrift:",2,omitempty"`
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

type PlayMusicRetDataInfo struct {
	Gp        int32 `thrift:",1,omitempty"`
	Music_idx int32 `thrift:",2,omitempty"`
}

type TourMain struct {
	Call string           `thrift:",1,omitempty"`
	Data TourMainDataInfo `thrift:",2,omitempty"`
}

type TrackMyScore struct {
	Score int64 `thrift:",1,omitempty"`
	Medal int16 `thrift:",2,omitempty"`
}

type GetPostTimeDataInfo struct {
	U_seq       int32  `thrift:",1,omitempty"`
	U_id        string `thrift:",2,omitempty"`
	Uuid        string `thrift:",3,omitempty"`
	Device_uuid string `thrift:",4,omitempty"`
}

type GetTourList struct {
	Call string              `thrift:",1,omitempty"`
	Data GetTourListDataInfo `thrift:",2,omitempty"`
}

type GetTourRank struct {
	Call string              `thrift:",1,omitempty"`
	Data GetTourRankDataInfo `thrift:",2,omitempty"`
}

type TotalTopRank struct {
}

type TourMusicList struct {
	Music_idx int32 `thrift:",1,omitempty"`
}
