package BISMW

import "weather-dump/src/Meteor"

type Line struct {
	segments map[uint8]*Segment
}

func NewLine() *Line {
	e := Line{}
	e.segments = make(map[uint8]*Segment)
	return &e
}

func (e Line) RenderLine() []byte {
	var buf = make([]byte, 64*14*14)

	o := 0
	for y := 0; y < 8; y++ {
		for x := 0; x < 1568; x++ {
			if e.segments[uint8(x/112)] == nil {
				filler := [64 * 14 * 14]byte{}
				return filler[:]
			}
			//fmt.Println(uint8(x/112), o, y*112+x-((x/112)*112))
			segment := e.segments[uint8(x/112)].RenderSegment()
			buf[o] = segment[y*112+x-((x/112)*112)]
			o++
		}
	}

	return buf
}

func (e Line) GetDate() Meteor.Time {
	return Meteor.Time{}
	//return e.segments[0].GetDate()
}