package yijing

var (
	金 = 五行{名: "金"}
	木 = 五行{名: "木"}
	水 = 五行{名: "水"}
	火 = 五行{名: "火"}
	土 = 五行{名: "土"}
)

type 五行 struct {
	名 string
}

func (wx 五行) 克() 五行 {
	switch wx {
	case 金:
		return 木
	case 木:
		return 土
	case 水:
		return 火
	case 火:
		return 金
	case 土:
		return 水
	}
	return 五行{}
}

func (wx 五行) 生() 五行 {
	switch wx {
	case 金:
		return 水
	case 木:
		return 火
	case 水:
		return 木
	case 火:
		return 土
	case 土:
		return 金
	}
	return 五行{}
}
