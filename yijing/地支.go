package yijing

var (
	子 = 地支{
		名:  "子",
		五行: 水,
		阴阳: 阳,
		藏干: []天干{癸},
		生肖: "鼠",
		时辰: []string{"23:00:00", "01:00:00"},
	}
	丑 = 地支{
		名:  "丑",
		五行: 土,
		阴阳: 阴,
		藏干: []天干{己, 辛, 癸},
		生肖: "牛",
		时辰: []string{"01:00:00", "03:00:00"},
	}
	寅 = 地支{
		名:  "寅",
		五行: 木,
		阴阳: 阳,
		藏干: []天干{甲, 丙, 戊},
		生肖: "虎",
		时辰: []string{"03:00:00", "05:00:00"},
	}
	卯 = 地支{
		名:  "卯",
		五行: 木,
		阴阳: 阴,
		藏干: []天干{乙},
		生肖: "兔",
		时辰: []string{"05:00:00", "07:00:00"},
	}
	辰 = 地支{
		名:  "辰",
		五行: 土,
		阴阳: 阳,
		藏干: []天干{乙, 戊, 癸},
		生肖: "龙",
		时辰: []string{"07:00:00", "09:00:00"},
	}
	巳 = 地支{
		名:  "巳",
		五行: 火,
		阴阳: 阴,
		藏干: []天干{丙, 戊, 庚},
		生肖: "蛇",
		时辰: []string{"09:00:00", "11:00:00"},
	}
	午 = 地支{
		名:  "午",
		五行: 火,
		阴阳: 阳,
		藏干: []天干{丁, 己},
		生肖: "马",
		时辰: []string{"11:00:00", "13:00:00"},
	}
	未 = 地支{
		名:  "未",
		五行: 土,
		阴阳: 阴,
		藏干: []天干{乙, 丁, 己},
		生肖: "羊",
		时辰: []string{"13:00:00", "15:00:00"},
	}
	申 = 地支{
		名:  "申",
		五行: 金,
		阴阳: 阳,
		藏干: []天干{戊, 庚, 壬},
		生肖: "猴",
		时辰: []string{"15:00:00", "17:00:00"},
	}
	酉 = 地支{
		名:  "酉",
		五行: 金,
		阴阳: 阴,
		藏干: []天干{辛},
		生肖: "鸡",
		时辰: []string{"17:00:00", "19:00:00"},
	}
	戌 = 地支{
		名:  "戌",
		五行: 土,
		阴阳: 阳,
		藏干: []天干{丁, 戊, 辛},
		生肖: "狗",
		时辰: []string{"19:00:00", "21:00:00"},
	}
	亥 = 地支{
		名:  "亥",
		五行: 水,
		阴阳: 阴,
		藏干: []天干{甲, 壬},
		生肖: "猪",
		时辰: []string{"21:00:00", "23:00:00"},
	}
)

type 地支 struct {
	名  string
	五行 五行
	阴阳 阴阳
	藏干 []天干
	生肖 string
	时辰 []string
}

func (dz 地支) String() string {
	return dz.名
}

func 地支解析(dz string) 地支 {
	switch dz {
	case "子":
		return 子
	case "丑":
		return 丑
	case "寅":
		return 寅
	case "卯":
		return 卯
	case "辰":
		return 辰
	case "巳":
		return 巳
	case "午":
		return 午
	case "未":
		return 未
	case "申":
		return 申
	case "酉":
		return 酉
	case "戌":
		return 戌
	case "亥":
		return 亥
	default:
		return 地支{}
	}
}
