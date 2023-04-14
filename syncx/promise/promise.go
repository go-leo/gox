package promise

type Promise interface {
	Then() Promise
	Catch() Promise
	Finally() Promise
}

var _ Promise = new(promise)

type promise struct {
}

func (p *promise) Then() Promise {
	return New()
}

func (p *promise) Catch() Promise {
	return New()
}

func (p *promise) Finally() Promise {
	// TODO implement me
	panic("implement me")
}

func New() Promise {
	return &promise{}
}

func All(promises ...Promise) Promise {
	return New()
}

func AllSettled(promises ...Promise) Promise {
	return New()
}

func Any(promises ...Promise) Promise {
	return New()
}

func Race(promises ...Promise) Promise {
	return New()
}

func Reject(err error) Promise {
	return New()
}

func Resolve(a any) Promise {
	return New()
}
