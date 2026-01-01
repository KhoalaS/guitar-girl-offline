package common_model

type ServerTimeRet struct {
	Time     int32 `thrift:",1,omitempty"`
	Datetime int64 `thrift:",2,omitempty"`
}

type ErrorRetCode struct {
	Code   int32  `thrift:",1,omitempty"`
	Errmsg string `thrift:",2,omitempty"`
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

type ParamData struct {
	Client_ver int32  `thrift:",1,omitempty"`
	Type       string `thrift:",2,omitempty"`
	Os         int16  `thrift:",3,omitempty"`
}
