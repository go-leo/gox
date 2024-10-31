package xlsx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Person struct {
	Name string `xlsx:"name,head=姓名"`
	Age  int    `xlsx:"age,head=年龄"`
	Sex  string `xlsx:"sex,head=性别"`
}

func TestToXLSX(t *testing.T) {
	p1 := Person{
		Name: "张三",
		Age:  18,
		Sex:  "男",
	}
	data, err := ToXLSX(p1)
	assert.NoError(t, err)
	t.Log(data)
}
