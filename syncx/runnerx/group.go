package runnerx

import "golang.org/x/sync/errgroup"

type Group struct {
	Runners []Runner
}

func (g *Group) Run(ctx *Context) error {
	c, err := errgroup.WithContext(ctx)
	c.Go(func() error {})
	for _, runner := range g.Runners {
		if err := runner.Run(ctx); err != nil {
			return err
		}
	}
	return nil
}