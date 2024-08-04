package yijing

import (
	"testing"
)

// 单元测试
func Test五行缺(t *testing.T) {
	// 构造一个八字对象
	bz := 八字{}

	// 调用五行缺函数
	missing := bz.五行缺()

	// 验证结果是否正确
	expected := 全部五行 // 预期结果
	if len(missing) != len(expected) {
		t.Errorf("missing length = %d; expected %d", len(missing), len(expected))
	}

	
}

func Test八字解析(t *testing.T) { 
	bz := 八字解析("壬","申","甲","辰","丁","丑","庚","戌")
	t.Log("五行",bz.五行())
	t.Log("五行缺",bz.五行缺())
	

}