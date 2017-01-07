package cube

import (
	"fmt"

	"github.com/mgutz/ansi"
)

type Color int

const (
	Black  Color = iota
	Red          = iota
	Blue         = iota
	Orange       = iota
	White        = iota
	Green        = iota
	Yellow       = iota
)

var ColorMap map[Color]string

func init() {
	ColorMap = map[Color]string{
		Red:    "Red",
		Blue:   "Blue",
		Orange: "Orange",
		White:  "White",
		Green:  "Green",
		Yellow: "Yellow",
	}
}

var size = 3

type MiniCube struct {
	Id string

	// These colors represent the original colors
	XColor,
	NegativeXColor,
	YColor,
	NegativeYColor,
	ZColor,
	NegativeZColor Color
}

var redFunc = ansi.ColorFunc("232:1")
var blueFunc = ansi.ColorFunc("232:4")
var orangeFunc = ansi.ColorFunc("232:166")
var whiteFunc = ansi.ColorFunc("232:231")
var greenFunc = ansi.ColorFunc("232:2")
var yellowFunc = ansi.ColorFunc("232:3")
var blackFunc = ansi.ColorFunc("232:232")

func showColor(c Color) string {
	box := "_"
	switch c {
	case Red:
		return redFunc(box)
	case Blue:
		return blueFunc(box)
	case Orange:
		return orangeFunc(box)
	case White:
		return whiteFunc(box)
	case Green:
		return greenFunc(box)
	case Yellow:
		return yellowFunc(box)
	default:
		return blackFunc(" ")
	}

}

func (mc *MiniCube) ShowFront() string {
	return showColor(mc.NegativeZColor)
}

func (mc *MiniCube) ShowTop() string {
	return showColor(mc.YColor)
}

func (mc *MiniCube) ShowRight() string {
	return showColor(mc.XColor)
}

func (mc *MiniCube) ShowLeft() string {
	return showColor(mc.NegativeXColor)
}

func (mc *MiniCube) ShowBottom() string {
	return showColor(mc.NegativeYColor)
}

func (mc *MiniCube) ShowBack() string {
	return showColor(mc.ZColor)
}

func (mc *MiniCube) frontInverted() {
	temp := mc.XColor
	mc.XColor = mc.NegativeYColor
	mc.NegativeYColor = mc.NegativeXColor
	mc.NegativeXColor = mc.YColor
	mc.YColor = temp
}

func (mc *MiniCube) top() {
	temp := mc.XColor
	mc.XColor = mc.ZColor
	mc.ZColor = mc.NegativeXColor
	mc.NegativeXColor = mc.NegativeZColor
	mc.NegativeZColor = temp

}
func (mc *MiniCube) right() {
	temp := mc.ZColor
	mc.ZColor = mc.YColor
	mc.YColor = mc.NegativeZColor
	mc.NegativeZColor = mc.NegativeYColor
	mc.NegativeYColor = temp
}

func (mc *MiniCube) String() string { return fmt.Sprintf("{%s}", mc.Id) }

type Cube [3][3][3]*MiniCube

func NewCube() *Cube {
	cube := Cube{}
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			for z := 0; z < size; z++ {
				s := &MiniCube{
					Id: fmt.Sprintf("%d,%d,%d", x, y, z),
				}
				if x == 0 {
					s.NegativeXColor = Green
				} else {
					s.NegativeXColor = Black
				}
				if x == 2 {
					s.XColor = Blue
				} else {
					s.XColor = Black
				}

				if y == 0 {
					s.NegativeYColor = Red
				} else {
					s.NegativeYColor = Black
				}
				if y == 2 {
					s.YColor = Orange
				} else {
					s.YColor = Black
				}

				if z == 0 {
					s.NegativeZColor = White
				} else {
					s.NegativeZColor = Black
				}
				if z == 2 {
					s.ZColor = Yellow
				} else {
					s.ZColor = Black
				}

				cube[x][y][z] = s
			}
		}
	}
	return &cube
}
