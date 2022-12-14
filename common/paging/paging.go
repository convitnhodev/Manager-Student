package paging

import "strings"

type Paging struct {
	// focus
	// utilizing form because page, limit, total appear in url

	Page int64 `json:"page" form:"page"`
	// limit each page
	Limit int64 `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"total"`
	TotalPage int64 `json:"total_page" form:"total_page"`
	// support cursor with UID
	FakeCursor string `json:"cursor" form:"cursor"`
	NextCursor string `json:"next_cursor"`
}

func (p *Paging) Fullfill() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Limit <= 0 {
		p.Limit = 50
	}
	p.FakeCursor = strings.TrimSpace(p.FakeCursor)
}
