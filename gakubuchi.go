package main

import (
	"golang.org/x/text/width"
	"golang.org/x/xerrors"
)

type Gakubuchi struct {
	corner string
	edge string
	spaces int
}

func Newgakubuchi(conf *config) *Gakubuchi {
	return &Gakubuchi{
		corner: conf.corner,
		edge: conf.edge,
		spaces: conf.spaces,
	}
}

func (g *Gakubuchi) GetStringWidth(s string) (int, error) {
	var size int
	for _, runeValue := range s {
		p := width.LookupRune(runeValue)
		switch p.Kind() {
		case width.EastAsianWide, width.EastAsianFullwidth:
			size += 2
		case width.EastAsianNarrow, width.EastAsianHalfwidth:
			size += 1
		default:
			return 0, xerrors.New("invalid string contained")
		}
	}
	return size, nil
}
