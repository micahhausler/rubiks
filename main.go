package main

import (
	"fmt"
	"github.com/micahhausler/rubiks/cube"
	"github.com/micahhausler/rubiks/solver"
)



func main() {
	c := cube.NewCube()

	fmt.Println("Show cube corner")
	c.ShowCubeTR()

	fmt.Println()
	fmt.Println("Mixed Up")
	c.MixUp()
	c.ShowCubeTR()
	fmt.Println()

	fmt.Println("Alligned white to front")
	c = solver.AlignWhiteCenter(c)
	c.ShowCubeTR()

	fmt.Println("Alligned Blue to right")
	c = solver.AlignBlueCenter(c)
	c.ShowCubeTR()

	fmt.Println("Put white on front somewhere")
	c.Top().Top().RightInverted()
	c.ShowCubeTR()

	fmt.Println("Put white on bottom")
	c = solver.AlignFrontCross(c, cube.White)
	c.ShowCubeTR()




}
