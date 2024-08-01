package yijing

type WuXing struct {
	Name string
}

var Jing = WuXing{Name: "金"}
var Mu = WuXing{Name: "木"}
var Shui = WuXing{Name: "水"}
var Huo = WuXing{Name: "火"}
var Tu = WuXing{Name: "土"}

func (wx WuXing) Ke() WuXing {
	switch wx {
	case Jing:
		return Mu
	case Mu:
		return Tu
	case Shui:
		return Huo
	case Huo:
		return Jing
	case Tu:
		return Shui
	}
	return WuXing{}
}
