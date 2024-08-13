package tunnyx

import (
	"github.com/mdlayher/schedgroup"
)

type Gopher struct {
	Pool *schedgroup.Group
}

func (g Gopher) Go(f func()) error {
	// Delay传入的是一个time.Duration参数，它会在time.Now()+delay之后执行函数，
	//g.Pool.Delay()
	// Schedule指定明确的某个时间执行。
	//g.Pool.Schedule()
	return nil
}
