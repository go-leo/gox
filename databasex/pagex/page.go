package pagex

import (
	"errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Page struct {
	// pageNum 页码，从1开始
	pageNum uint64
	// pageSize 页面大小
	pageSize uint64
	// offset 跳过的行数
	offset uint64
	// limit 限制行数
	limit uint64
	// total 总行数
	total uint64
	// pages 总页数
	pages uint64
	// count 包含count查询
	count bool
	// countColumn 进行count查询的列名
	countColumn string
	// orderBy 排序,
	orderBy string
	// startTime 开始时间
	startTime time.Time
	// endTime 结束时间
	endTime time.Time
}

func (p *Page) init() *Page {
	return p
}

func (p *Page) apply(opts ...Option) *Page {
	for _, opt := range opts {
		opt(p)
	}
	return p
}

// PageNum 获取页码
func (p *Page) PageNum() uint64 {
	return p.pageNum
}

// PageSize 获取页面大小
func (p *Page) PageSize() uint64 {
	return p.pageSize
}

// Offset 获取跳过的行数
func (p *Page) Offset() uint64 {
	return p.offset
}

// Limit 获取限制行数
func (p *Page) Limit() uint64 {
	return p.limit
}

// Total 获取总行数
func (p *Page) Total() uint64 {
	return p.total
}

// SetTotal 设置总行数, 并计算总页数
func (p *Page) SetTotal(total uint64) {
	p.total = total
	p.pages = (total + p.pageSize - 1) / p.pageSize
}

// Pages 获取总页数
func (p *Page) Pages() uint64 {
	return p.pages
}

// StartTime 获取开始时间
func (p *Page) StartTime() time.Time {
	return p.startTime
}

// EndTime 获取结束时间
func (p *Page) EndTime() time.Time {
	return p.endTime
}

// Count 获取是否包含count查询
func (p *Page) Count() bool {
	return p.count
}

// CountColumn 获取进行count查询的列名
func (p *Page) CountColumn() string {
	return p.countColumn
}

// OrderBy 获取排序字段
func (p *Page) OrderBy() string {
	return p.orderBy
}

func (p *Page) AsProto() *PageProto {
	return &PageProto{
		PageNum:     p.pageNum,
		PageSize:    p.pageSize,
		Offset:      p.offset,
		Limit:       p.limit,
		Total:       p.total,
		Pages:       p.pages,
		Count:       p.count,
		CountColumn: p.countColumn,
		OrderBy:     p.orderBy,
		StartTime:   timestamppb.New(p.startTime),
		EndTime:     timestamppb.New(p.endTime),
	}
}

type Option func(p *Page)

func Count(count bool) Option {
	return func(p *Page) {
		p.count = count
	}
}

func CountColumn(countColumn string) Option {
	return func(p *Page) {
		p.countColumn = countColumn
	}
}

// OrderBy 设置排序字段
func OrderBy(orderBy string) Option {
	return func(p *Page) {
		p.orderBy = orderBy
	}
}

func StartTime(t time.Time) Option {
	return func(p *Page) {
		p.startTime = t
	}
}

func EndTime(t time.Time) Option {
	return func(p *Page) {
		p.endTime = t
	}
}

func NewPage(pageNum uint64, pageSize uint64, opts ...Option) (*Page, error) {
	if pageNum == 0 {
		return nil, errors.New("pagex: pageNum is zero")
	}
	if pageSize == 0 {
		return nil, errors.New("pagex: pageSize is zero")
	}
	// 计算出 offset 和 limit
	offset := (pageNum - 1) * pageSize
	limit := pageSize
	page := &Page{
		pageNum:  pageNum,
		pageSize: pageSize,
		offset:   offset,
		limit:    limit,
		count:    true,
	}
	p := page.apply(opts...).init()
	return p, nil
}

func FromProto(p *PageProto) *Page {
	return &Page{
		pageNum:     p.PageNum,
		pageSize:    p.PageSize,
		offset:      p.Offset,
		limit:       p.Limit,
		total:       p.Total,
		pages:       p.Pages,
		count:       p.Count,
		countColumn: p.CountColumn,
		orderBy:     p.OrderBy,
		startTime:   p.StartTime.AsTime(),
		endTime:     p.EndTime.AsTime(),
	}
}
