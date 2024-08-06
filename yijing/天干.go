package yijing

import "slices"

var (
	甲 = 天干{
		名:  "甲",
		五行: 木,
		阴阳: 阳,
		器官: 肝,
	}
	乙 = 天干{
		名:  "乙",
		五行: 木,
		阴阳: 阴,
		器官: 胆,
	}
	丙 = 天干{
		名:  "丙",
		五行: 火,
		阴阳: 阳,
		器官: 小肠,
	}
	丁 = 天干{
		名:  "丁",
		五行: 火,
		阴阳: 阴,
		器官: 心,
	}
	戊 = 天干{
		名:  "戊",
		五行: 土,
		阴阳: 阳,
		器官: 胃,
	}
	己 = 天干{
		名:  "己",
		五行: 土,
		阴阳: 阴,
		器官: 脾,
	}
	庚 = 天干{
		名:  "庚",
		五行: 金,
		阴阳: 阳,
		器官: 大肠,
	}
	辛 = 天干{
		名:  "辛",
		五行: 金,
		阴阳: 阴,
		器官: 肺,
	}
	壬 = 天干{
		名:  "壬",
		五行: 水,
		阴阳: 阳,
		器官: 膀胱,
	}
	癸 = 天干{
		名:  "癸",
		五行: 水,
		阴阳: 阴,
		器官: 肾,
	}

	全部天干 = []天干{甲, 乙, 丙, 丁, 戊, 己, 庚, 辛, 壬, 癸}
)

type 天干 struct {
	名  string
	五行 五行
	阴阳 阴阳
	器官 器官
}

func (tg 天干) String() string {
	return tg.名
}

func (tg 天干) Next() 天干 {
	index := slices.IndexFunc(全部天干, func(item 天干) bool {
		return tg.名 == item.名
	})
	return 全部天干[(index+1)%len(全部天干)]
}

func 天干解析(tg string) 天干 {
	switch tg {
	case "甲":
		return 甲
	case "乙":
		return 乙
	case "丙":
		return 丙
	case "丁":
		return 丁
	case "戊":
		return 戊
	case "己":
		return 己
	case "庚":
		return 庚
	case "辛":
		return 辛
	case "壬":
		return 壬
	case "癸":
		return 癸
	default:
		return 天干{}
	}
}
