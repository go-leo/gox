package yijing

var (
	阴 = 阴阳{名: "阴"}
	阳 = 阴阳{名: "阳"}
)

type 阴阳 struct {
	名 string
}
