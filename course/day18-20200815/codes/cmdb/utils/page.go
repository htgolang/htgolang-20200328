package utils

import (
	"fmt"
	"html/template"
	"net/url"
)

type PageQueryParams struct {
	PNum   int64      `form:"pageNum"`
	PSize  int64      `form:"pageSize"`
	Inputs url.Values `form:"-"`
}

func (p *PageQueryParams) PageNum() int64 {
	if p.PNum >= 1 {
		return p.PNum
	}
	return 1
}

func (p *PageQueryParams) Offset() int64 {
	fmt.Println((p.PageNum() - 1) * p.PageSize())
	return (p.PageNum() - 1) * p.PageSize()
}

func (p *PageQueryParams) PageSize() int64 {
	if p.PSize <= 0 || p.PSize > 100 {
		return 3
	}
	return p.PSize
}

type Page struct {
	Datas    interface{}
	Total    int64
	PageSize int64
	PageNum  int64

	PrevPage int64
	NextPage int64
	Pages    []int64

	QueryParams template.URL
}

func NewPage(datas interface{}, total int64, pageSize int64, pageNum int64, inputs url.Values) *Page {

	maxPage := total / pageSize
	if total%pageSize != 0 {
		maxPage += 1
	}

	prevPage := pageNum - 1
	if prevPage <= 1 {
		prevPage = 1
	}
	nextPage := pageNum + 1
	if nextPage >= maxPage {
		nextPage = maxPage
	}
	// [5]   [1] 2 [3]
	startPage := pageNum - 2
	endPage := pageNum + 2
	if startPage <= 1 {
		startPage = 1
	}
	if endPage >= maxPage {
		endPage = maxPage
	}

	pages := []int64{}
	for page := startPage; page <= endPage; page++ {
		pages = append(pages, page)
	}

	inputs.Del("pageNum")

	return &Page{
		Datas:       datas,
		Total:       total,
		PageSize:    pageSize,
		PageNum:     pageNum,
		PrevPage:    prevPage,
		NextPage:    nextPage,
		Pages:       pages,
		QueryParams: template.URL(inputs.Encode()),
	}
}
