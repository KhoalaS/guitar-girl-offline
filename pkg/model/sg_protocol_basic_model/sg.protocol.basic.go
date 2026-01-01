package sg_protocol_basic_model

type Error struct {
	Is_success      bool   `thrift:",1,omitempty"`
	Err_code        int32  `thrift:",2,omitempty"`
	Err_message     string `thrift:",3,omitempty"`
	Is_debug        bool   `thrift:",4,omitempty"`
	Is_achieve_noti bool   `thrift:",5,omitempty"`
	Mission_notis   []any  `thrift:",6,omitempty"` // TODO
}
