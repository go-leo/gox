package yijing

var (
	吉 = 吉凶{名: "吉"}
	凶 = 吉凶{名: "凶"}
)

type 吉凶 struct {
	名 string
}
