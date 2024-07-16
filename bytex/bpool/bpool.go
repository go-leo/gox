// Package bpool
// pool最大的特色就是能够保持池子中元素的数量，一旦Put的数量多于它的阈值，就会自动丢弃，而 sync.Pool是一个没有限制的池子，只要Put就会收进去。
// bpool是基于Channel实现的，不像sync.Pool为了提高性能而做了很多优化，所以，在性能上比不过 sync.Pool。
// 不过，它提供了限制Pool容量的功能，所以，如果你想控制Pool的容量的话，可以考虑这个库
// See: https://github.com/oxtoacart/bpool
package bpool
