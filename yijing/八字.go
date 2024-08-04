package yijing

import "github.com/go-leo/gox/slicex"

type 八字 struct {
	年柱 柱
	月柱 柱
	日柱 柱
	时柱 柱
}

func 八字解析(年干, 年支, 月干, 月支, 日干, 日支, 时干, 时支 string) 八字 {
	return 八字{
		年柱: 柱解析(年干, 年支),
		月柱: 柱解析(月干, 月支),
		日柱: 柱解析(日干, 日支),
		时柱: 柱解析(时干, 时支),
	}
}

func (bz 八字) 日元() 天干 {
	return bz.日柱.天干
}

func (bz 八字) 五行() []五行 {
	return []五行{
		bz.年柱.天干.五行,
		bz.年柱.地支.五行,
		bz.月柱.天干.五行,
		bz.月柱.地支.五行,
		bz.日柱.天干.五行,
		bz.日柱.地支.五行,
		bz.时柱.天干.五行,
		bz.时柱.地支.五行,
	}
}

func (bz 八字) 五行缺() []五行 {
	return slicex.Difference[[]五行](全部五行, slicex.Uniq[[]五行](bz.五行()))
}

func (bz 八字) 十神() 八字十神 {
	return 算八字十神(bz)
}