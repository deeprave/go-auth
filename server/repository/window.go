package repository

import "fmt"

type Window struct {
	Limit  int64
	Offset int64
}

func NewWindow(limit int64, ofs ...int64) *Window {
	offset := int64(0)
	if len(ofs) > 0 {
		offset = ofs[0]
	}
	return &Window{
		Limit:  limit,
		Offset: offset,
	}
}

func (r *Window) Next(limit ...int64) {
	shift := r.Limit
	if len(limit) == 0 {
		shift = limit[0]
	}
	r.Offset += shift
}

func (r *Window) Previous(limit ...int64) {
	shift := r.Limit
	if len(limit) == 0 {
		shift = limit[0]
	}
	if shift < r.Offset {
		r.Offset -= shift
	} else {
		r.Offset = int64(0)
	}
}

func (r *Window) ToString() string {
	limit := ""
	if r.Limit > 0 {
		limit = fmt.Sprintf("LIMIT %d", r.Limit)
		if r.Offset > 0 {
			limit = fmt.Sprintf("%s OFFSET %d", r.Offset)
		}
	}
	return limit
}
