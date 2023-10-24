package pagex

import (
	"errors"
	"github.com/go-leo/gox/databasex/sqls"
	"github.com/go-leo/gox/stringx"
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
	// checkOrderBySqlInjection 做了sql注入检查
	checkOrderBySqlInjection bool
	// startTime 开始时间
	startTime time.Time
	// endTime 结束时间
	endTime time.Time
}

func (p *Page) init() {

}

func (p *Page) apply(opts ...Option) {
	for _, opt := range opts {
		opt(p)
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
func OrderBy(orderBy string, check ...bool) Option {
	return func(p *Page) {
		p.orderBy = orderBy
		if len(check) > 0 {
			p.checkOrderBySqlInjection = check[0]
		}
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
		return nil, errors.New("pageNum is zero")
	}
	if pageSize == 0 {
		return nil, errors.New("pageSize is zero")
	}
	p := &Page{
		pageNum:                  pageNum,
		pageSize:                 pageSize,
		offset:                   0,
		limit:                    0,
		total:                    0,
		pages:                    0,
		count:                    true,
		countColumn:              "",
		orderBy:                  "",
		checkOrderBySqlInjection: false,
		startTime:                time.Time{},
		endTime:                  time.Time{},
	}
	p.apply(opts...)
	p.init()
	if stringx.IsNotBlank(p.orderBy) && p.checkOrderBySqlInjection && sqls.CheckSqlInjection(p.orderBy, false) {
		return nil, errors.New("order by [" + p.orderBy + "] 存在 SQL 注入风险, 如想避免 SQL 注入校验，选用 UnsafeOrderBy")
	}
	// 计算出 offset 和 limit
	p.offset = (p.pageNum - 1) * p.pageSize
	p.limit = p.pageSize
	return p, nil
}

func (p *Page) PageNum() uint64 {
	return p.pageNum
}

func (p *Page) PageSize() uint64 {
	return p.pageSize
}

func (p *Page) Offset() uint64 {
	return p.offset
}

func (p *Page) Limit() uint64 {
	return p.limit
}

func (p *Page) Total() uint64 {
	return p.total
}

func (p *Page) SetTotal(total uint64) {
	p.total = total
	if total%p.pageSize == 0 {
		p.pages = total / p.pageSize
	} else {
		p.pages = total/p.pageSize + 1
	}
}

func (p *Page) StartTime() time.Time {
	return p.startTime
}

func (p *Page) EndTime() time.Time {
	return p.endTime
}

func (p *Page) Pages() uint64 {
	return p.pages
}

func (p *Page) Count() bool {
	return p.count
}

func (p *Page) CountColumn() string {
	return p.countColumn
}

func (p *Page) OrderBy() string {
	return p.orderBy
}
