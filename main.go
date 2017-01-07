package main

import (
	"fmt"
	"github.com/micahhausler/rubiks/cube"
	"github.com/micahhausler/rubiks/solver"
)



func main() {
	cube := cube.NewCube()

	fmt.Println("Show cube corner")
	cube.ShowCubeTR()

	fmt.Println()
	fmt.Println("Mixed Up")
	cube.MixUp()
	cube.ShowCubeTR()
	fmt.Println()

	fmt.Println("Alligned white to front")
	cube = solver.AlignWhiteCenter(cube)
	cube.ShowCubeTR()

	fmt.Println("Alligned Blue to right")
	cube = solver.AlignBlueCenter(cube)
	cube.ShowCubeTR()

	fmt.Println("Put white on front somewhere")
	cube.Top().Top().RightInverted()
	cube.ShowCubeTR()

	fmt.Println("Put white on bottom")
	cube = solver.AllignWhiteCross(cube)
	cube.ShowCubeTR()




}
