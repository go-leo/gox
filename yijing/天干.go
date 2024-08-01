package yijing

var (
	甲 = 天干{
		名:  "甲",
		五行: 木,
		阴阳: 阳,
	}
	乙 = 天干{
		名:  "乙",
		五行: 木,
		阴阳: 阴,
	}
	丙 = 天干{
		名:  "丙",
		五行: 火,
		阴阳: 阳,
	}
	丁 = 天干{
		名:  "丁",
		五行: 火,
		阴阳: 阴,
	}
	戊 = 天干{
		名:  "戊",
		五行: 土,
		阴阳: 阳,
	}
	己 = 天干{
		名:  "己",
		五行: 土,
		阴阳: 阴,
	}
	庚 = 天干{
		名:  "庚",
		五行: 金,
		阴阳: 阳,
	}
	辛 = 天干{
		名:  "辛",
		五行: 金,
		阴阳: 阴,
	}
	壬 = 天干{
		名:  "壬",
		五行: 水,
		阴阳: 阳,
	}
	癸 = 天干{
		名:  "癸",
		五行: 水,
		阴阳: 阴,
	}
)

type 天干 struct {
	名  string
	五行 五行
	阴阳 阴阳
}

// 生我者正印偏印，我生者伤官食神;克我者正官七杀，我克者正财偏财;同我者比肩劫财。
// 注：口诀中的“我”代表的是八字中的日元天干。
func (tg 天干) 十神(日元 天干) 十神 {
	// 同我者为比劫。同性为比肩，简称比；异性为劫财，简称劫。
	if tg.五行 == 日元.五行 && tg.阴阳 == 日元.阴阳 {
		return 比肩
	}
	if tg.五行 == 日元.五行 && tg.阴阳 != 日元.阴阳 {
		return 劫财
	}

	// 克我者为杀官。同性为偏官，或者七杀，简称杀；异性为正官，简称官
	if tg.五行.克() == 日元.五行 && tg.阴阳 == 日元.阴阳 {
		return 七杀
	}
	if tg.五行.克() == 日元.五行 && tg.阴阳 != 日元.阴阳 {
		return 正官
	}

	// 生我者为枭印。同性为偏印，也叫枭神，简称枭；异性为正印，简称印。
	if tg.五行.生() == 日元.五行 && tg.阴阳 == 日元.阴阳 {
		return 偏印
	}
	if tg.五行.生() == 日元.五行 && tg.阴阳 != 日元.阴阳 {
		return 正印
	}

	// 我克者为才财。同性为偏财，简称财；异性为正财，简称才
	if 日元.五行.克() == tg.五行 && 日元.阴阳 == tg.阴阳 {
		return 偏财
	}
	if 日元.五行.克() == tg.五行 && 日元.阴阳 != tg.阴阳 {
		return 正财
	}

	// 我生者为食伤。同性为食神，简称食；异性为伤官，简称伤
	if 日元.五行.生() == tg.五行 && 日元.阴阳 == tg.阴阳 {
		return 食神
	}
	if 日元.五行.生() == tg.五行 && 日元.阴阳 != tg.阴阳 {
		return 伤官
	}

	return 十神{}
}
