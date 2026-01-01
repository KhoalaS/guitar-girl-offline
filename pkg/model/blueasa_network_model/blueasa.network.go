package blueasa_network_model

type Rect struct {
	Left   float64 `thrift:",1,omitempty"`
	Top    float64 `thrift:",2,omitempty"`
	Width  float64 `thrift:",3,omitempty"`
	Height float64 `thrift:",4,omitempty"`
}

type Vector3 struct {
	X float64 `thrift:",1,omitempty"`
	Y float64 `thrift:",2,omitempty"`
	Z float64 `thrift:",3,omitempty"`
}
