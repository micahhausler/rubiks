package solver

import (
	"github.com/micahhausler/rubiks/cube"
	//"fmt"
	"fmt"
)

func AlignWhiteCenter(c *cube.Cube) *cube.Cube {

	if c[1][1][0].NegativeZColor != cube.White {
		for x := 0; x < 3; x++ {
			c = c.CubeSpinLeft()
			if c[1][1][0].NegativeZColor != cube.White {
				continue
			} else {
				break
			}
		}
		if c[1][1][0].NegativeZColor != cube.White {
			for x := 0; x < 3; x++ {
				c = c.CubeDown()
				if c[1][1][0].NegativeZColor != cube.White {
					continue
				} else {
					break
				}
			}
		}
	}
	return c
}

func AlignBlueCenter(c *cube.Cube) *cube.Cube {
	if c[2][1][1].XColor != cube.Blue {
		for x := 0; x < 3; x++ {
			c = c.CubeFlipRight()
			if c[1][1][0].XColor != cube.Blue {
				continue
			} else {
				break
			}
		}
	}
	return c
}

// tell how many of the front edges are aligned in a cross,
// where the top = 1,
// top, and right = 2
// top, right, bottom = 3
// all = 4
func frontEdgesNumColors(c *cube.Cube, color cube.Color) int {
	response := 0
	if c[1][2][0].NegativeZColor == color {
		return 1
	}
	if c[2][1][0].NegativeZColor == color {
		return 2
	}
	if c[1][0][0].NegativeZColor == color {
		return 3
	}
	if c[0][1][0].NegativeZColor == color {
		return 4
	}
	return response
}

var (
	bottomWhite,
	leftWhite,
	rightWhite,
	topWhite bool
)

type Axis int

const (
	X Axis = iota
	Y      = iota
	Z      = iota
)

// Return the edge colors for an axis at a certain depth (in relation to the front)
func edgeColors(c *cube.Cube, axis Axis, depth int) []cube.Color {
	colors := []cube.Color{}
	switch axis {
	case X:
		switch depth {
		case 0, 2:
			colors = append(
				colors,
				c[depth][1][0].NegativeZColor,
				c[depth][2][1].YColor,
				c[depth][1][2].ZColor,
				c[depth][0][1].NegativeYColor,
			)
		case 1:
			colors = append(
				colors,
				c[depth][0][0].NegativeZColor,
				c[depth][2][0].NegativeZColor,
				c[depth][2][0].YColor,
				c[depth][2][2].YColor,
				c[depth][2][2].ZColor,
				c[depth][0][2].ZColor,
				c[depth][0][2].NegativeYColor,
				c[depth][0][0].NegativeYColor,
			)
		default:
		}
	case Y:
		switch depth {
		case 0, 2:
			colors = append(
				colors,
				c[1][depth][0].NegativeZColor,
				c[2][depth][1].XColor,
				c[1][depth][2].ZColor,
				c[0][depth][1].NegativeXColor,
			)
		case 1:
			colors = append(
				colors,
				c[0][depth][0].NegativeZColor,
				c[2][depth][0].NegativeZColor,
				c[2][depth][0].XColor,
				c[2][depth][2].XColor,
				c[2][depth][2].ZColor,
				c[0][depth][2].ZColor,
				c[0][depth][2].NegativeXColor,
				c[0][depth][0].NegativeXColor,
			)
		default:
		}
	case Z:
		switch depth {
		case 0, 2:
			colors = append(
				colors,
				c[1][0][depth].NegativeYColor,
				c[2][1][depth].XColor,
				c[1][2][depth].YColor,
				c[0][1][depth].NegativeXColor,
			)
		case 1:
			colors = append(
				colors,
				c[0][0][depth].NegativeYColor,
				c[2][0][depth].NegativeYColor,
				c[2][0][depth].XColor,
				c[2][2][depth].XColor,
				c[2][2][depth].YColor,
				c[0][2][depth].YColor,
				c[0][2][depth].NegativeXColor,
				c[0][0][depth].NegativeXColor,
			)
		}
	default:

	}
	return colors
}

func AlignFrontCross(c *cube.Cube, color cube.Color) *cube.Cube {

	for _, c := range edgeColors(c, Y, 2){

		fmt.Println(c.String())
	}

	switch frontEdgesNumColors(c, color) {
	case 0:
		if c[2][2][1].XColor == color {
			c.Top()
			c = AlignFrontCross(c, color)
		}
	case 1:
	case 2:
	case 3:
	default:
		return c
	}

	if frontEdgesNumColors(c, color) < 4 {

		// search front, place the first white at bottom
		z := 0
		for x := 0; x < 3; x++ {
			for y := 0; y < 3; y++ {
				//fmt.Printf("%d, %d\n",x,y )
				//fmt.Printf("x(%d) + y(%d) mod 2 = %d\n",x ,y, int(x+y)%2)
				if (x+y)%2 == 1 {
					//fmt.Printf("x(%d) + y(%d) mod 2 = %d\n",x ,y, int(x+y)%2)
					if c[x][y][z].NegativeZColor == color {
						//fmt.Printf("White edge found! %d,%d,%d \n", x,y,z)
						// White edge found!

						if x != 1 && y != 0 {
							for i := 0; i < 3; i++ {
								c = c.Front()
								if c[1][0][0].NegativeZColor != color {
									continue
								} else {
									break
								}
							}
						}
						bottomWhite = true
					}
				}
			}
		}
	}

	/*

		// search top
		y := 2
		for z := 2; z > 0; z-- {
			for x := 0; x < 2; x++ {
				if x+z%2 != 0 {
					if c[x][y][z-1].YColor == cube.White {
						// White edge found!
					}
				}
			}
		}
	*/
	return c
}
