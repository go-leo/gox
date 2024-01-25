package sample

type Gopher struct{}

func (g Gopher) Go(f func()) error {
	go f()
	return nil
}
