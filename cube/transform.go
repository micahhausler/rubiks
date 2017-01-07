package cube

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"text/template"
	//"time"
)

func (c *Cube) RandomMove() *Cube {
	//rand.Seed(int64(time.Now().Second()))
	switch x := rand.Int() % 18; x {
	case 0:
		c = c.CubeUp()
	case 1:
		c = c.CubeDown()
	case 2:
		c = c.CubeSpinRight()
	case 3:
		c = c.CubeSpinLeft()
	case 4:
		c = c.CubeFlipRight()
	case 5:
		c = c.CubeFlipLeft()
	case 6:
		c = c.Right()
	case 7:
		c = c.RightInverted()
	case 8:
		c = c.Left()
	case 9:
		c = c.LeftInverted()
	case 10:
		c = c.Front()
	case 11:
		c = c.FrontInverted()
	case 12:
		c = c.Top()
	case 13:
		c = c.TopInverted()
	case 14:
		c = c.Bottom()
	case 15:
		c = c.BottomInverted()
	case 16:
		c = c.Back()
	default:
		c = c.BackInverted()
	}
	return c
}

func (c *Cube) MixUp() *Cube {
	steps := 100

	for i := 0; i < steps; i++ {
		c = c.RandomMove()
	}
	return c
}

func (c *Cube) printRows(rows [][]string) {
	for _, row := range rows {
		fmt.Println(strings.Join(row, " "))
	}
}

func (c *Cube) rightColors() [][]string {
	rows := [][]string{}
	x := 2
	for y := size; y > 0; y-- {
		row := []string{}
		for z := 0; z < size; z++ {
			row = append(row, c[x][y-1][z].ShowRight())
		}
		rows = append(rows, row)
	}
	return rows
}

func (c *Cube) ShowRight() {
	c.printRows(c.rightColors())
}

func (c *Cube) frontColors() [][]string {
	rows := [][]string{}
	z := 0
	for y := size; y > 0; y-- {
		row := []string{}
		for x := 0; x < size; x++ {
			row = append(row, c[x][y-1][z].ShowFront())
		}
		rows = append(rows, row)
	}
	return rows
}

func (c *Cube) ShowFront() {
	c.printRows(c.frontColors())
}

func (c *Cube) topColors() [][]string {
	rows := [][]string{}
	y := 2
	for z := size - 1; z > -1; z-- {
		row := []string{}
		for x := 0; x < size; x++ {
			row = append(row, c[x][y][z].ShowTop())
		}
		rows = append(rows, row)
	}
	return rows
}

func (c *Cube) ShowTop() {
	c.printRows(c.frontColors())
}

func (c *Cube) ShowLeft() {
	x := 0
	for y := size; y > 0; y-- {
		row := []string{}
		for z := size; z > 0; z-- {
			row = append(row, c[x][y-1][z-1].ShowLeft())
		}
		fmt.Println(strings.Join(row, " "))
	}
}

func ShowColors() {}

func (c *Cube) ShowCubeTR() {
	tmplData := `
       {{ index .topColors 0 0 }} {{ index .topColors 0 1}} {{ index .topColors 0 2}}
    {{ index .topColors 1 0}} {{ index .topColors 1 1}} {{ index .topColors 1 2}}    {{ index .rightColors 0 2 }}
 {{ index .topColors 2 0}} {{ index .topColors 2 1}} {{ index .topColors 2 2}}    {{ index .rightColors 0 1 }}  {{ index .rightColors 1 2}}
{{ index .frontColors 0 0 }} {{ index .frontColors 0 1 }} {{ index .frontColors 0 2 }}  {{ index .rightColors 0 0 }}  {{ index .rightColors 1 1 }}  {{ index .rightColors 2 2 }}
{{ index .frontColors 1 0 }} {{ index .frontColors 1 1 }} {{ index .frontColors 1 2 }}  {{ index .rightColors 1 0 }}  {{ index .rightColors 2 1 }}
{{ index .frontColors 2 0 }} {{ index .frontColors 2 1 }} {{ index .frontColors 2 2 }}  {{ index .rightColors 2 0 }}
`
	tmpl, err := template.New("test").Parse(tmplData)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, map[string]interface{}{
		"topColors":   c.topColors(),
		"frontColors": c.frontColors(),
		"rightColors": c.rightColors(),
	})
	if err != nil {
		panic(err)
	}

}

type Rotation func() *Cube

func (c *Cube) rotateX(x int) *Cube {
	for y := 0; y < size/2; y++ {
		for z := 0; z < size-y-1; z++ {
			var temp = c[x][y][z]

			c[x][y][z] = c[x][z][size-1-y]
			c[x][y][z].right()

			c[x][z][size-1-y] = c[x][size-1-y][size-1-z]
			c[x][z][size-1-y].right()

			c[x][size-1-y][size-1-z] = c[x][size-1-z][y]
			c[x][size-1-y][size-1-z].right()

			c[x][size-1-z][y] = temp
			c[x][size-1-z][y].right()

		}
	}
	return c
}

func (c *Cube) rotateY(y int) *Cube {
	for x := 0; x < size/2; x++ {
		for z := 0; z < size-1-x; z++ {
			var temp = c[x][y][z]

			c[x][y][z] = c[size-1-z][y][x]
			c[x][y][z].top()

			c[size-1-z][y][x] = c[size-1-x][y][size-1-z]
			c[size-1-z][y][x].top()

			c[size-1-x][y][size-1-z] = c[z][y][size-1-x]
			c[size-1-x][y][size-1-z].top()

			c[z][y][size-1-x] = temp
			c[z][y][size-1-x].top()
		}
	}
	return c
}

func (c *Cube) rotateInvertedZ(z int) *Cube {
	for x := 0; x < size/2; x++ {
		for y := 0; y < size-x-1; y++ {
			var temp = c[x][y][z]

			c[x][y][z] = c[y][size-1-x][z]
			c[x][y][z].frontInverted()

			c[y][size-1-x][z] = c[size-1-x][size-1-y][z]
			c[y][size-1-x][z].frontInverted()

			c[size-1-x][size-1-y][z] = c[size-1-y][x][z]
			c[size-1-x][size-1-y][z].frontInverted()

			c[size-1-y][x][z] = temp
			c[size-1-y][x][z].frontInverted()
		}
	}
	return c
}

func (c *Cube) CubeUp() *Cube {
	return c.rotateX(0).rotateX(1).rotateX(2)
}
func (c *Cube) CubeDown() *Cube {
	return c.CubeUp().CubeUp().CubeUp()
}

func (c *Cube) CubeSpinRight() *Cube {
	return c.rotateY(2).rotateY(1).rotateY(0)
}
func (c *Cube) CubeSpinLeft() *Cube {
	return c.CubeSpinRight().CubeSpinRight().CubeSpinRight()
}

func (c *Cube) CubeFlipRight() *Cube {
	return c.CubeFlipLeft().CubeFlipLeft().CubeFlipLeft()
}
func (c *Cube) CubeFlipLeft() *Cube {
	return c.rotateInvertedZ(0).rotateInvertedZ(1).rotateInvertedZ(2)
}

func (c *Cube) Bottom() *Cube {
	return c.BottomInverted().BottomInverted().BottomInverted()
}
func (c *Cube) BottomInverted() *Cube {
	return c.rotateY(0)
}

func (c *Cube) Top() *Cube {
	return c.rotateY(2)
}
func (c *Cube) TopInverted() *Cube {
	return c.Top().Top().Top()
}

func (c *Cube) Front() *Cube {
	return c.FrontInverted().FrontInverted().FrontInverted()
}
func (c *Cube) FrontInverted() *Cube {
	return c.rotateInvertedZ(0)
}

func (c *Cube) Back() *Cube {
	return c.rotateInvertedZ(2).rotateInvertedZ(2).rotateInvertedZ(2)
}
func (c *Cube) BackInverted() *Cube {
	return c.Back().Back().Back()
}

func (c *Cube) Right() *Cube {
	return c.rotateX(2)
}
func (c *Cube) RightInverted() *Cube {
	return c.Right().Right().Right()
}

func (c *Cube) Left() *Cube {
	return c.rotateX(0).rotateX(0).rotateX(0)
}
func (c *Cube) LeftInverted() *Cube {
	return c.Left().Left().Left()
}
