package xlsx

import (
	"testing"
)

type Person struct {
	Name string `xlsx:"name,head=姓名"`
	Age  int    `xlsx:"age,head=年龄"`
	Sex  string `xlsx:"sex,head=性别"`
}

func TestToXLSX(t *testing.T) {
	fields := []string{"姓名", "年龄", "性别"}
	rows := [][]any{
		{"张三", 18, "男"},
		{"李四", 19, "女"},
		{"王五", 20, "男"},
	}
	data, err := ToXLSX(fields, rows)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}
