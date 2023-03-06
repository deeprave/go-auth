package lib

type Ptrs []any

func NewPtrs(ptrs ...any) Ptrs {
	p := Ptrs{}
	return p.Append(ptrs...)
}

func (p Ptrs) Append(ptrs ...any) Ptrs {
	P := p
	for _, ptr := range ptrs {
		if ptr != nil {
			P = append(P, ptr)
		}
	}
	return P
}
