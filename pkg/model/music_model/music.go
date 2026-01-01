package music_model

import "github.com/KhoalaS/guitar-girl-offline/pkg/model/common_model"

type GetAlbumList struct {
	Call string               `thrift:",1,omitempty"`
	Data GetAlbumListDataInfo `thrift:",2,omitempty"`
}

type GetCacheDataInfo struct {
	Type int16 `thrift:",1,omitempty"`
}

type GetMusicInfoListRetDataInfo struct {
	Language map[any]any `thrift:",1,omitempty"` // TODO
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

type AlbumListLanguage struct {
	Title        string `thrift:",1,omitempty"`
	Company      string `thrift:",2,omitempty"`
	Album_artist string `thrift:",3,omitempty"`
}

type GetAlbumListRetDataInfo struct {
	Album_idx  int32  `thrift:",1,omitempty"`
	Country    string `thrift:",2,omitempty"`
	Date_issue string `thrift:",3,omitempty"`
	Cdn_dir    string `thrift:",4,omitempty"`
	Album_img  string `thrift:",5,omitempty"`
}

type GetMusicList struct {
	Call string               `thrift:",1,omitempty"`
	Data GetMusicListDataInfo `thrift:",2,omitempty"`
}

type GetReviewPointDataInfo struct {
	Null string `thrift:",1,omitempty"`
}

type MusicListLanguage struct {
	Title      string `thrift:",1,omitempty"`
	Artist     string `thrift:",2,omitempty"`
	Songwriter string `thrift:",3,omitempty"`
	Lyricist   string `thrift:",4,omitempty"`
}

type AlbumListInfo struct {
	Album_idx  int64       `thrift:",1,omitempty"`
	Country    string      `thrift:",2,omitempty"`
	Date_issue string      `thrift:",3,omitempty"`
	Cdn_dir    string      `thrift:",4,omitempty"`
	Album_img  string      `thrift:",5,omitempty"`
	Language   map[any]any `thrift:",6,omitempty"` // TODO
}

type GetAlbumDataInfo struct {
	Null string `thrift:",1,omitempty"`
}

type GetAlbumInfoListRetDataInfo struct {
	Language map[any]any `thrift:",1,omitempty"` // TODO
}

type GetCacheRetDataInfo struct {
	Uptime int32 `thrift:",1,omitempty"`
}

type GetMusicInfoList struct {
	Call string                   `thrift:",1,omitempty"`
	Data GetMusicInfoListDataInfo `thrift:",2,omitempty"`
}

type GetMusicListReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        []any                        `thrift:",4,omitempty"` // TODO
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type GetAlbumInfoList struct {
	Call string               `thrift:",1,omitempty"`
	Data GetAlbumListDataInfo `thrift:",2,omitempty"`
}

type GetAlbumRetDataInfo struct {
	Album     AlbumListInfo `thrift:",1,omitempty"`
	List_info []any         `thrift:",2,omitempty"` // TODO
}

type GetCache struct {
	Call string           `thrift:",1,omitempty"`
	Data GetCacheDataInfo `thrift:",2,omitempty"`
}

type GetReviewPoint struct {
	Call string           `thrift:",1,omitempty"`
	Data GetAlbumDataInfo `thrift:",2,omitempty"`
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

type GetAlbum struct {
	Call string           `thrift:",1,omitempty"`
	Data GetAlbumDataInfo `thrift:",2,omitempty"`
}

type GetAlbumInfoLanguage struct {
	Seq          int32  `thrift:",1,omitempty"`
	Album_title  string `thrift:",2,omitempty"`
	Company      string `thrift:",3,omitempty"`
	Album_artist string `thrift:",4,omitempty"`
}

type GetAlbumInfoListReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        map[any]any                  `thrift:",4,omitempty"` // TODO
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type GetAlbumListReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        []any                        `thrift:",4,omitempty"` // TODO
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type GetMusicInfoListDataInfo struct {
	Country_code string `thrift:",1,omitempty"`
}

type GetMusicInfoListReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        map[any]any                  `thrift:",4,omitempty"` // TODO
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type GetReviewPointRetDataInfo struct {
	Music_idx  int32   `thrift:",1,omitempty"`
	Difficulty int16   `thrift:",2,omitempty"`
	Point      float64 `thrift:",3,omitempty"`
}

type GetAlbumReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        []any                        `thrift:",4,omitempty"` // TODO
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type GetCacheReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        GetCacheRetDataInfo          `thrift:",4,omitempty"`
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
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

type GetReviewPointReturn struct {
	Error       common_model.ErrorRetCode    `thrift:",1,omitempty"`
	Mode        string                       `thrift:",2,omitempty"`
	Call        string                       `thrift:",3,omitempty"`
	Data        []any                        `thrift:",4,omitempty"` // TODO
	Maintenance common_model.MaintenanceData `thrift:",5,omitempty"`
}

type GetAlbumInfoListDataInfo struct {
}

type GetAlbumListDataInfo struct {
}

type GetMusicListDataInfo struct {
	Country_code string `thrift:",1,omitempty"`
}
