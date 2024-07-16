// Package bytebufferpool
// 它提供了校准（calibrate，用来动态调整创建元素的权重）的机制，可以“智能”地调整Pool的defaultSize和maxSize。
// 一般来说，我们使用bytebufferpool的场景比较固定，所用buffer的大小会集中在某个范围里。
// 有了校准的特性，bytebufferpool就能够偏重于创建这个范围大小的buffer，从而节省空间。
// See: https://github.com/valyala/bytebufferpool
package bytebufferpool
