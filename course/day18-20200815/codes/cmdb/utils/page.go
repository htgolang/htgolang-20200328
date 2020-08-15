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
	if p.PNum < 1 {
		return 1
	}
	return p.PNum
}

func (p *PageQueryParams) PageSize() int64 {
	if p.PSize < 1 || p.PSize > 100 {
		return 10
	}
	return p.PSize
}

func (p *PageQueryParams) Offset() int64 {
	fmt.Println((p.PageNum() - 1) * p.PageSize())
	return (p.PageNum() - 1) * p.PageSize()
}

type Page struct {
	Total    int64
	Datas    interface{}
	PageSize int64
	PageNum  int64

	PrevPage int64
	NextPage int64
	Pages    []int64

	QueryParams template.URL
}

func NewPage(total int64, datas interface{}, pageSize int64, pageNum int64, inputs url.Values) *Page {
	pages := make([]int64, 0, 11)

	maxPage := total / pageSize
	if total%pageSize != 0 {
		maxPage += 1
	}

	startPage := pageNum - 5
	endPage := pageNum + 5
	if startPage < 1 {
		endPage -= startPage
		startPage = 1
	}

	if endPage > maxPage {
		startPage -= endPage - maxPage
		endPage = maxPage
	}
	if startPage < 1 {
		startPage = 1
	}

	prevPage := pageNum - 1
	if prevPage < 1 {
		prevPage = 1
	}
	nextPage := pageNum + 1
	if nextPage > maxPage {
		nextPage = maxPage
	}

	for page := startPage; page <= endPage; page++ {
		pages = append(pages, page)
	}

	inputs.Del("pageNum")

	return &Page{
		Total:       total,
		Datas:       datas,
		PageSize:    pageSize,
		PageNum:     pageNum,
		PrevPage:    prevPage,
		NextPage:    nextPage,
		Pages:       pages,
		QueryParams: template.URL(inputs.Encode()),
	}
}
