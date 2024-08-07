package yijing

import "fmt"

var (
	比肩 = 十神{
		名:  "比肩",
		简称: "比",
		别名: "",
		吉凶: 吉,
	}
	劫财 = 十神{
		名:  "劫财",
		简称: "劫",
		别名: "",
		吉凶: 凶,
	}
	偏财 = 十神{
		名:  "偏财",
		简称: "财",
		别名: "",
		吉凶: 吉,
	}
	正财 = 十神{
		名:  "正财",
		简称: "才",
		别名: "",
		吉凶: 吉,
	}
	食神 = 十神{
		名:  "食神",
		简称: "食",
		别名: "",
		吉凶: 吉,
	}
	伤官 = 十神{
		名:  "伤官",
		简称: "伤",
		别名: "",
		吉凶: 凶,
	}
	七杀 = 十神{
		名:  "七杀",
		简称: "杀",
		别名: "偏官",
		吉凶: 凶,
	}
	正官 = 十神{
		名:  "正官",
		简称: "官",
		别名: "",
		吉凶: 吉,
	}
	偏印 = 十神{
		名:  "偏印",
		简称: "枭",
		别名: "枭神",
		吉凶: 凶,
	}
	正印 = 十神{
		名:  "正印",
		简称: "印",
		别名: "",
		吉凶: 吉,
	}
)

type 十神 struct {
	名  string
	简称 string
	别名 string
	吉凶 吉凶
}

func (ss 十神) String() string {
	return ss.名
}

// 生我者正印偏印，我生者伤官食神;克我者正官七杀，我克者正财偏财;同我者比肩劫财。
// 注：口诀中的“我”代表的是八字中的日元天干。
func 算十神(日元 天干, tg 天干) 十神 {
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

type 柱十神 struct {
	干神 十神
	支神 []十神
}

func (zss 柱十神) String() string {
	return fmt.Sprintf("干神:%s 支神:%v", zss.干神, zss.支神)
}

func 算柱十神(日元 天干, z 柱) 柱十神 {
	var zs []十神
	for _, tg := range z.地支.藏干 {
		zs = append(zs, 算十神(日元, tg))
	}
	return 柱十神{
		干神: 算十神(日元, z.天干),
		支神: zs,
	}
}

type 八字十神 struct {
	年柱 柱十神
	月柱 柱十神
	日柱 柱十神
	时柱 柱十神
}

func (bzss 八字十神) String() string {
	return fmt.Sprintf("年柱: %s\n月柱: %s\n日柱: %s\n时柱: %s", bzss.年柱, bzss.月柱, bzss.日柱, bzss.时柱)
}

func 算八字十神(bz 八字) 八字十神 {
	return 八字十神{
		年柱: 算柱十神(bz.日柱.天干, bz.年柱),
		月柱: 算柱十神(bz.日柱.天干, bz.月柱),
		日柱: 算柱十神(bz.日柱.天干, bz.日柱),
		时柱: 算柱十神(bz.日柱.天干, bz.时柱),
	}
}
