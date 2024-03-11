package web

type Pageable interface {
	Limit() int
	Offset() int
}

type pageable struct {
	page int
	size int
}

func (p pageable) Limit() int {
	return p.size
}

func (p pageable) Offset() int {
	return (p.page - 1) * p.size
}

func NewPageable(page, size int) Pageable {
	return pageable{
		page: page,
		size: size,
	}
}
