package concurrent

type Executor interface {
	Execute(command Runner)
}
