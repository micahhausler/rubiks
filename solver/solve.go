package solver

import (
	"github.com/micahhausler/rubiks/cube"
	//"fmt"
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

var (
	bottomWhite,
	leftWhite,
	rightWhite,
	topWhite bool
)


func AllignWhiteCross(c *cube.Cube) *cube.Cube {

	// search front, place the first white at bottom
	z := 0
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			//fmt.Printf("%d, %d\n",x,y )
			//fmt.Printf("x(%d) + y(%d) mod 2 = %d\n",x ,y, int(x+y)%2)
			if (x+y)%2 == 1 {
				//fmt.Printf("x(%d) + y(%d) mod 2 = %d\n",x ,y, int(x+y)%2)
				if c[x][y][z].NegativeZColor == cube.White {
					//fmt.Printf("White edge found! %d,%d,%d \n", x,y,z)
					// White edge found!

					if x != 1 && y != 0 {
						for i := 0; i < 3; i++ {
							c = c.Front()
							if c[1][0][0].NegativeZColor != cube.White {
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
