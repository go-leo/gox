package bazi

import (
	"fmt"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	tim := time.Unix(0, 0).Local()
	t.Log(tim.Format(time.DateTime))
	t.Log(tim.Unix())

	t.Log("--------------------")

	tim = time.Unix(-3600, 0).Local()
	t.Log(tim.Format(time.DateTime))
	t.Log(tim.Unix())

	t.Log("--------------------")

	tim = time.Time{}
	t.Log(tim.Format(time.DateTime))
	t.Log(tim.Unix())
}

func TestGet(t *testing.T) {
	var heavenlyStems = []string{
		"Jia",
		"Yi",
		"Bing",
		"Ding",
		"Wu",
		"Ji",
		"Geng",
		"Xin",
		"Ren",
		"Gui",
	}
	var earthlyBranches = []string{
		"Zi",
		"Chou",
		"Yin",
		"Mao",
		"Chen",
		"Si",
		"Wu",
		"Wei",
		"Shen",
		"You",
		"Xu",
		"Hai",
	}
	for i := 0;i < len(heavenlyStems) * len(earthlyBranches) / 2; i ++ {
		h := i % len(heavenlyStems)
		e := i % len(earthlyBranches)
		fmt.Println(i, heavenlyStems[h],earthlyBranches[e])
	}

	// var HeavenlyStems = []HeavenlyStem{
	// 	HeavenlyStemBing,
	// 	HeavenlyStemJia,
	// 	HeavenlyStemShi,
	// 	HeavenlyStemGeng,
	// 	HeavenlyStemXin,
	// 	HeavenlyStemRen,
	// 	HeavenlyStemGui,
	// }
	fmt.Println("var HeavenlyStems = []HeavenlyStem{")
	for _, v := range heavenlyStems {
		fmt.Println("HeavenlyStem"+v)
	}
	fmt.Println("}")

	fmt.Println("var earthlyBranches = []EarthlyBranche{")
	for _, v := range earthlyBranches {
		fmt.Println("EarthlyBranche"+v+",")
	}
	fmt.Println("}")
}
