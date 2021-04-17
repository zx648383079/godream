package utils

type Pager struct {
	Current   uint
	Size      uint
	Total     uint
	Pages     uint
	PageItems []uint
	Begin     uint
	End       uint
	Prev      uint
	Next      uint
	IsPrev    bool
	IsNext    bool
}

func (p Pager) Limit() int {
	return int(p.Size)
}

func (p Pager) Offset() int {
	return int(p.Begin) - 1
}

func NewPager(page, size uint, total uint) *Pager {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 20
	}
	p := new(Pager)
	p.Current = page
	p.Size = size
	p.Total = total
	p.Pages = total / size
	if total%size > 0 {
		p.Pages++
	}
	p.PageItems = make([]uint, p.Pages)
	var i uint
	for i = 1; i <= p.Pages; i++ {
		p.PageItems[i-1] = i
	}
	p.Begin = (page-1)*size + 1
	if p.Begin < 1 {
		p.Begin = 1
	}
	if p.Begin > p.Total {
		p.Begin = p.Total
	}
	p.End = page * size
	if p.End > p.Total {
		p.End = p.Total
	}
	p.Prev = p.Current - 1
	p.IsPrev = true
	if p.Prev < 1 {
		p.Prev = 1
		p.IsPrev = false
	}
	p.Next = p.Current + 1
	p.IsNext = true
	if p.Next > p.Pages {
		p.Next = p.Pages
		p.IsNext = false
	}
	return p
}
