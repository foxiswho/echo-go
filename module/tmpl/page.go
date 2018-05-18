package tmpl

import (
	"html/template"
	"math"
	"strconv"
)

type Page struct {
	Page  int
	Size  int
	Count int
	Url   string
}

func (p *Page) Html() template.HTML {
	rollPage := 3
	nowPage := p.Page
	totalPage := int(math.Ceil(float64(p.Count) / float64(p.Size)))

	if nowPage > totalPage {
		nowPage = totalPage
	}

	nowCoolPage := float64(rollPage) / float64(2)
	nowCoolPageCeil := int(math.Ceil(nowCoolPage))
	last := totalPage

	upRow := nowPage - 1
	upPage := ""
	if upRow > 0 {
		upPage = `<li><a href="` + p.url(upRow) + `">上一页</a></li>`
	}

	downRow := nowPage + 1
	downPage := ""
	if downRow <= totalPage {
		downPage = `<li><a href="` + p.url(downRow) + `">下一页</a></li>`
	}

	firstPage := ""
	if totalPage > rollPage && (float64(nowPage)-nowCoolPage) >= 1 {
		firstPage = `<li><a href="` + p.url(1) + `">首页</a></li>`
	}

	endPage := ""
	if totalPage > rollPage && (float64(nowPage)+nowCoolPage) < float64(totalPage) {
		endPage = `<li><a href="` + p.url(last) + `">` + strconv.Itoa(last) + `</a></li>`
	}

	linkPage := ""
	for i := 1; i <= rollPage; i++ {
		page := 0
		if (float64(nowPage) - nowCoolPage) <= 0 {
			page = i
		} else if (float64(nowPage) + nowCoolPage - 1) >= float64(totalPage) {
			page = totalPage - rollPage + i
		} else {
			page = nowPage - nowCoolPageCeil + i
		}

		if page > 0 && page != nowPage {
			if page <= totalPage {
				linkPage += `<li><a href="` + p.url(page) + `">` + strconv.Itoa(page) + `</a></li>`
			} else {
				break
			}
		} else {
			if page > 0 && totalPage != 1 {
				linkPage += `<li class="active"><a href="` + p.url(page) + `">` + strconv.Itoa(page) + `</a></li>`
			}
		}
	}

	pageStr := firstPage + upPage + linkPage + downPage + endPage
	return template.HTML(`<nav> <ul class="pagination">` + pageStr + `</ul></nav>`)
}

func (p *Page) url(page int) string {
	return p.Url + "?p=" + strconv.Itoa(page) + "&s=" + strconv.Itoa(p.Size)
}
