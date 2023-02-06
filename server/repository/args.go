package repository

import "strings"

type Args []string

func NewArgs(args ...string) Args {
	a := Args{}
	return a.Append(args...)
}

func (a Args) Append(args ...string) Args {
	A := a
	for _, arg := range args {
		if arg != "" {
			A = append(A, arg)
		}
	}
	return A
}

func (a Args) ToString(separator ...string) string {
	sep := " "
	if len(separator) > 0 {
		sep = separator[0]
	}
	return strings.Join(a, sep)
}

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
