package groupx

// https://github.com/go-kratos/kratos/tree/v1.0.x/pkg/sync/errgroup
// 1. 控制并发goroutine的数量
// 2. cancel，失败的子任务可以cancel所有正在执行任务；
// 3. recover，而且会把panic的堆栈信息放到error中，避免子任务panic导致的程序崩溃。

// https://github.com/neilotoole/errgroup
// 它可以直接替换官方的ErrGroup，方法都一样，原有功能也一样，
// 只不过增加了可以控制并发goroutine的功能增加了可以控制并发goroutine的功能

// https://github.com/facebookarchive/errgroup
// Facebook提供的这个ErrGroup，其实并不是对Go扩展库ErrGroup的扩展，而是对标准库WaitGroup的扩展。
// 不过，因为它们的名字一样，处理的场景也类似，所以我把它也列在了这里。
// 标准库的WaitGroup只提供了Add、Done、Wait方法，而且Wait方法也没有返回子goroutine的error。
// 而Facebook提供的ErrGroup提供的Wait方法可以返回error，而且可以包含多个error。
// 子任务在调用Done之 前，可以把自己的error信息设置给ErrGroup。
// 接着，Wait在返回的时候，就会把这些error信息返回给调用者

// https://github.com/go-pkgz/syncs
// 提供了两个Group并发原语，分别是SizedGroup和ErrSizedGroup。
// SizedGroup内部是使用信号量和WaitGroup实现的，它通过信号量控制并发的goroutine数量，或者是不控
// 制goroutine数量，只控制子任务并发执行时候的数量（通过）。
// ErrSizedGroup为SizedGroup提供了error处理的功能，它的功能和Go官方扩展库的功能一样，就是等待子
// 任务完成并返回第一个出现的error。

// https://github.com/vardius/gollback
// https://github.com/AaronJan/Hunch
// 解决了ErrGroup收集子任务返回结果的痛点。
// 使用 ErrGroup时，如果你要收到子任务的结果和错误，你需要定义额外的变量收集执行结果和错误，
// 但是这个库可以提供更便利的方式。
