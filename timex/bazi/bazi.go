package bazi

import (
	"fmt"
	"time"
)

var (
	// HeavenlyStemJia 甲
	HeavenlyStemJia = HeavenlyStem{"甲"}
	// HeavenlyStemYi 乙
	HeavenlyStemYi = HeavenlyStem{"乙"}
	// HeavenlyStemBing 丙
	HeavenlyStemBing = HeavenlyStem{"丙"}
	// HeavenlyStemDing 丁
	HeavenlyStemDing = HeavenlyStem{"丁"}
	// HeavenlyStemWu 戊
	HeavenlyStemWu = HeavenlyStem{"戊"}
	// HeavenlyStemJi 己
	HeavenlyStemJi = HeavenlyStem{"己"}
	// HeavenlyStemGeng 庚
	HeavenlyStemGeng = HeavenlyStem{"庚"}
	// HeavenlyStemXin 辛
	HeavenlyStemXin = HeavenlyStem{"辛"}
	// HeavenlyStemRen 壬
	HeavenlyStemRen = HeavenlyStem{"壬"}
	// HeavenlyStemGui 癸
	HeavenlyStemGui = HeavenlyStem{"癸"}
)

var heavenlyStems = []HeavenlyStem{
	HeavenlyStemJia,
	HeavenlyStemYi,
	HeavenlyStemBing,
	HeavenlyStemDing,
	HeavenlyStemWu,
	HeavenlyStemJi,
	HeavenlyStemGeng,
	HeavenlyStemXin,
	HeavenlyStemRen,
	HeavenlyStemGui,
}

var (
	// EarthlyBrancheZi 子 23:00:00 - 01:00:00
	EarthlyBrancheZi = EarthlyBranche{"子", 23, 1}
	// EarthlyBrancheChou 丑 01:00:00 - 03:00:00
	EarthlyBrancheChou = EarthlyBranche{"丑", 1, 3}
	// EarthlyBrancheYin 寅 03:00:00 - 05:00:00
	EarthlyBrancheYin = EarthlyBranche{"寅", 3, 5}
	// EarthlyBrancheMao 卯 05:00:00 - 07:00:00
	EarthlyBrancheMao = EarthlyBranche{"卯", 5, 7}
	// EarthlyBrancheChen 辰 07:00:00 - 09:00:00
	EarthlyBrancheChen = EarthlyBranche{"辰", 7, 9}
	//EarthlyBranche Si 巳 09:00:00 - 11:00:00
	EarthlyBrancheSi = EarthlyBranche{"巳", 9, 11}
	// EarthlyBrancheWu 午 11:00:00 - 13:00:00
	EarthlyBrancheWu = EarthlyBranche{"午", 11, 13}
	// EarthlyBrancheWei 未 13:00:00 - 15:00:00
	EarthlyBrancheWei = EarthlyBranche{"未", 13, 15}
	// EarthlyBrancheShen 申 15:00:00 - 17:00:00
	EarthlyBrancheShen = EarthlyBranche{"申", 15, 17}
	// EarthlyBrancheYou 酉 17:00:00 - 19:00:00
	EarthlyBrancheYou = EarthlyBranche{"酉", 17, 19}
	// EarthlyBrancheXu 戌 19:00:00 - 21:00:00
	EarthlyBrancheXu = EarthlyBranche{"戌", 19, 21}
	// EarthlyBrancheHai 亥 21:00:00 - 23:00:00
	EarthlyBrancheHai = EarthlyBranche{"亥", 21, 23}
)

var earthlyBranches = []EarthlyBranche{
	EarthlyBrancheZi,
	EarthlyBrancheChou,
	EarthlyBrancheYin,
	EarthlyBrancheMao,
	EarthlyBrancheChen,
	EarthlyBrancheSi,
	EarthlyBrancheWu,
	EarthlyBrancheWei,
	EarthlyBrancheShen,
	EarthlyBrancheYou,
	EarthlyBrancheXu,
	EarthlyBrancheHai,
}

var (
	PillarJiaZi = Pillar{
		HeavenlyStem:   HeavenlyStemGui,
		EarthlyBranche: EarthlyBrancheZi,
	}
)

type HeavenlyStem struct {
	Name string
}

type EarthlyBranche struct {
	Name  string
	Start int
	End   int
}

func (eb EarthlyBranche) String() string {
	return fmt.Sprintf("%s%s", eb.Name, eb.Name)
}

// Pillar 柱
type Pillar struct {
	// HeavenlyStem 天干
	HeavenlyStem HeavenlyStem
	// EarthlyBranche 地支
	EarthlyBranche EarthlyBranche
}

func (p Pillar) String() string {
	return fmt.Sprintf("%s%s", p.HeavenlyStem, p.EarthlyBranche.String())
}

// BaZi 八字
type BaZi struct {
	t     time.Time
	Year  Pillar
	Month Pillar
	Day   Pillar
	Hour  Pillar
}

func (b BaZi) String() string {
	return fmt.Sprintf("%s %s %s %s", b.Year.String(), b.Month.String(), b.Day.String(), b.Hour.String())
}

func (b BaZi) AsTime() time.Time {
	return b.t
}

// unix:0 -> time:1970-01-01 08:00:00 -> bazi:己酉 丙子 辛巳 壬辰
// unix:-3600 -> time:1970-01-01 07:00:00 -> bazi:己酉 丙子 辛巳 壬辰
func Get(t time.Time) BaZi {
	st := t.Unix()
	diff := st - (-3600)
	if diff < 0 {
		diff = -1 * diff
	}
	

	return BaZi{t: t}
}
