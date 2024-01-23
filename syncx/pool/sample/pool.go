package sample

type Pool struct{}

func (Pool) Go(f func()) error {
	go func() { f() }()
	return nil
}
