package yijing

import (
	"fmt"
	"github.com/go-leo/gox/slicex"
	"testing"
	"time"
)

type Zhu struct {
	Tian string
	Di   string
	Cang []string
}

func (zhu Zhu) String() string {
	return fmt.Sprintf("{天:%s, 地:%s, 藏:%v}", zhu.Tian, zhu.Di, zhu.Cang)
}

func (zhu Zhu) WuXin() wuxin {
	return wuxin{
		Tian: TianWuXing[zhu.Tian],
		Di:   DiWuXing[zhu.Di],
		Cang: canganwuxin(zhu.Cang),
	}
}

func (zhu Zhu) shizheng(day Zhu) shizheng {
	dayWuxin := TianWuXing[day.Tian]
	dayYingYang := TianYingYang[day.Tian]

	zhuWuxin := TianWuXing[zhu.Tian]
	zhuYingYang := TianYingYang[zhu.Tian]

	var c []string
	for _, cang := range zhu.Cang {
		cangWuxin := TianWuXing[cang]
		cangYingYang := TianYingYang[cang]
		c = append(c, shishen(dayWuxin, dayYingYang, cangWuxin, cangYingYang))
	}
	return shizheng{
		Tian: shishen(dayWuxin, dayYingYang, zhuWuxin, zhuYingYang),
		Cang: c,
	}
}

func shishen(dayWuxin, dayYingYang, zhuWuxin, zhuYingYang string) string {
	var r string
	if dayWuxin == zhuWuxin {
		//同我、助我者为比劫。同性-比肩，异性-劫财， 兄弟
		if dayYingYang == zhuYingYang {
			r = "比肩(吉)"
		} else {
			r = "劫财(凶)"
		}
	} else if Ke(dayWuxin) == zhuWuxin {
		//我克、耗我者为财星。同性-偏财，异性-正财，钱
		if dayYingYang == zhuYingYang {
			r = "偏财(吉)"
		} else {
			r = "正财(吉)"
		}
	} else if Sheng(dayWuxin) == zhuWuxin {
		//我生、泄我者为食伤。同性-食神，异性-伤官，才华
		if dayYingYang == zhuYingYang {
			r = "食神(吉)"
		} else {
			r = "伤官(凶)"
		}
	} else if Ke(zhuWuxin) == dayWuxin {
		//克我，抑我者为官杀。同性-偏官，异性-正官，官
		if dayYingYang == zhuYingYang {
			r = "偏官(七杀)(凶)"
		} else {
			r = "正官(吉)"
		}
	} else if Sheng(zhuWuxin) == dayWuxin {
		//生我、扶我者为印星。同性-偏印，异性-正印，证件
		if dayYingYang == zhuYingYang {
			r = "偏印(凶)"
		} else {
			r = "正印(吉)"
		}
	}
	return r
}

type Bazi struct {
	Year  Zhu
	Month Zhu
	Day   Zhu
	Hour  Zhu
}

func (bz Bazi) String() string {
	return fmt.Sprintf("【年:%s, 月:%s, 日:%s, 时:%s】", bz.Year.String(), bz.Month.String(), bz.Day.String(), bz.Hour.String())
}

func (bz Bazi) WuXin() WuXin {
	return WuXin{
		Year:  bz.Year.WuXin(),
		Month: bz.Month.WuXin(),
		Day:   bz.Day.WuXin(),
		Hour:  bz.Hour.WuXin(),
	}
}

func (bz Bazi) ShiSheng() ShiSheng {
	return ShiSheng{
		Year:  bz.Year.shizheng(bz.Day),
		Month: bz.Month.shizheng(bz.Day),
		Day:   bz.Day.shizheng(bz.Day),
		Hour:  bz.Hour.shizheng(bz.Day),
	}
}

type shizheng struct {
	Tian string
	Cang []string
}

type ShiSheng struct {
	Year  shizheng
	Month shizheng
	Day   shizheng
	Hour  shizheng
}

type wuxin struct {
	Tian string
	Di   string
	Cang []string
}

func (zhu wuxin) String() string {
	return fmt.Sprintf("{天:%s, 地:%s, 藏:%v}", zhu.Tian, zhu.Di, zhu.Cang)
}

type WuXin struct {
	Year  wuxin
	Month wuxin
	Day   wuxin
	Hour  wuxin
}

func (wx WuXin) String() string {
	return fmt.Sprintf("【年:%s, 月:%s, 日:%s, 时:%s】", wx.Year.String(), wx.Month.String(), wx.Day.String(), wx.Hour.String())
}

func (wx WuXin) Que() []string {
	wxSlice := []string{wx.Year.Tian, wx.Year.Di, wx.Month.Tian, wx.Month.Di, wx.Day.Tian, wx.Day.Di, wx.Hour.Tian, wx.Hour.Di}
	return slicex.Difference(allWuXing, slicex.Uniq(wxSlice))
}

func TestName(t *testing.T) {
	for i := 0; i < 60; i++ {
		fmt.Printf("%s%s = 柱{天干: %s, 地支: %s}\n", Tian[i%10], Di[i%12], Tian[i%10], Di[i%12])
	}

	return

	a := Bazi{
		Year: Zhu{
			Tian: "壬",
			Di:   "申",
			Cang: canggan["申"],
		},
		Month: Zhu{
			Tian: "甲",
			Di:   "辰",
			Cang: canggan["辰"],
		},
		Day: Zhu{
			Tian: "丁",
			Di:   "丑",
			Cang: canggan["丑"],
		},
		Hour: Zhu{
			Tian: "庚",
			Di:   "戌",
			Cang: canggan["戌"],
		},
	}
	awuXin := a.WuXin()
	ashisheng := a.ShiSheng()
	fmt.Println(a)
	fmt.Println(awuXin)
	fmt.Println(ashisheng)

	b := Bazi{
		Year: Zhu{
			Tian: "癸",
			Di:   "酉",
			Cang: canggan["酉"],
		},
		Month: Zhu{
			Tian: "乙",
			Di:   "卯",
			Cang: canggan["卯"],
		},
		Day: Zhu{
			Tian: "甲",
			Di:   "辰",
			Cang: canggan["辰"],
		},
		Hour: Zhu{
			Tian: "戊",
			Di:   "辰",
			Cang: canggan["辰"],
		},
	}

	bwuXin := b.WuXin()
	bshisheng := b.ShiSheng()
	fmt.Println(b)
	fmt.Println(bwuXin)
	fmt.Println(bshisheng)

	return
	date := time.Date(2025, 1, 20, 0, 0, 0, 0, time.Local)
	bazi := Bazi{
		Year: Zhu{
			Tian: Tian[0],
			Di:   Di[4],
		},
		Month: Zhu{
			Tian: Tian[3],
			Di:   Di[1],
		},
		Day: Zhu{
			Tian: Tian[5],
			Di:   Di[1],
		},
		Hour: Zhu{
			Tian: Tian[0],
			Di:   Di[0],
		},
	}

	startT := 0
	startD := 0

	for d := 0; d < 10; d++ {
		for h := 0; h < 24; h++ {
			timestamp := date.AddDate(0, 0, d).Add(time.Duration(h) * time.Hour)
			bz := Bazi{
				Year: Zhu{
					Tian: bazi.Year.Tian,
					Di:   bazi.Year.Di,
					Cang: canggan[bazi.Year.Di],
				},
				Month: Zhu{
					Tian: bazi.Month.Tian,
					Di:   bazi.Month.Di,
					Cang: canggan[bazi.Month.Di],
				},
				Day: Zhu{
					Tian: Tian[(5+d)%10],
					Di:   Di[(1+d)%12],
					Cang: canggan[Di[(1+d)%12]],
				},
				Hour: Zhu{
					Tian: Tian[startT%10],
					Di:   Di[startD%12],
					Cang: canggan[Di[startD%12]],
				},
			}
			wx := bz.WuXin()
			que := wx.Que()
			ss := bz.ShiSheng()
			fmt.Println("=====", timestamp.Format(time.DateTime), "=====")
			fmt.Println("八字:", bz)
			fmt.Println("五行:", wx)
			fmt.Println("十神:", ss)
			fmt.Println("日元:", bz.Day.Tian, wx.Day.Tian)
			fmt.Println("五行缺:", que)
			if h%2 == 0 {
				startT++
				startD++
			}
		}
	}

}

var Tian = []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
var Di = []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}

var allWuXing = []string{"金", "木", "水", "火", "土"}

func Sheng(s string) string {
	switch s {
	case "金":
		return "水"
	case "木":
		return "火"
	case "水":
		return "木"
	case "火":
		return "土"
	case "土":
		return "金"
	}
	return ""
}

func Ke(s string) string {
	switch s {
	case "金":
		return "木"
	case "木":
		return "土"
	case "水":
		return "火"
	case "火":
		return "金"
	case "土":
		return "水"
	}
	return ""
}

var TianWuXing = map[string]string{
	"甲": "木",
	"乙": "木",
	"丙": "火",
	"丁": "火",
	"戊": "土",
	"己": "土",
	"庚": "金",
	"辛": "金",
	"壬": "水",
	"癸": "水",
}

var TianYingYang = map[string]string{
	"甲": "阳",
	"乙": "阴",
	"丙": "阳",
	"丁": "阴",
	"戊": "阳",
	"己": "阴",
	"庚": "阳",
	"辛": "阴",
	"壬": "阳",
	"癸": "阴",
}

var DiWuXing = map[string]string{
	"子": "水",
	"丑": "土",
	"寅": "木",
	"卯": "木",
	"辰": "土",
	"巳": "火",
	"午": "火",
	"未": "土",
	"申": "金",
	"酉": "金",
	"戌": "土",
	"亥": "水",
}

var DiYingYang = map[string]string{
	"子": "阳",
	"丑": "阴",
	"寅": "阳",
	"卯": "阴",
	"辰": "阳",
	"巳": "阴",
	"午": "阳",
	"未": "阴",
	"申": "阳",
	"酉": "阴",
	"戌": "阳",
	"亥": "阴",
}

var canggan = map[string][]string{
	"子": {"癸"},
	"丑": {"己", "癸", "辛"},
	"寅": {"甲", "丙", "戊"},
	"卯": {"乙"},
	"辰": {"戊", "乙", "癸"},
	"巳": {"丙", "戊", "庚"},
	"午": {"丁", "己"},
	"未": {"己", "丁", "乙"},
	"申": {"庚", "壬", "戊"},
	"酉": {"辛"},
	"戌": {"戊", "辛", "丁"},
	"亥": {"壬", "甲"},
}

func canganwuxin(s []string) []string {
	r := make([]string, 0, len(s))
	for _, s2 := range s {
		r = append(r, TianWuXing[s2])
	}
	return r
}
