package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/text/width"
	"golang.org/x/xerrors"
)

type Gakubuchi struct {
	corner string
	edge string
	spaces int

	sentences string
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

func (g *Gakubuchi) Print() {
	var maxLength int
	splited := make([]string, 0, 0)
	for _, s := range strings.Split(g.sentences, "\n") {
		splited = append(splited, s)
		if maxLength < len(s) {
			maxLength = len(s)
		}
	}

	fmt.Fprintf(os.Stdout, "%v%v%v", g.corner, strings.Repeat(g.edge, maxLength+(2*g.spaces)), g.corner)
	for _, s := range splited {
		fmt.Fprintf(os.Stdout, "%v%v%v%v%v", g.corner, strings.Repeat(" ", g.spaces), strings.Repeat(" ", g.spaces), g.corner)
	}
	fmt.Fprintf(os.Stdout, "%v%v%v", g.corner, strings.Repeat(g.edge, maxLength+(2*g.spaces)), g.corner)
}
