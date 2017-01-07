package cube

import (
	"testing"
	"strings"
	"fmt"
)


func TestRotateFrontInverted(t *testing.T) {
	cube := NewCube()

	cube.FrontInverted()
	z := 0
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			originalId := fmt.Sprintf("%d,%d,%d", 0, y, z)
			if strings.Compare(cube[0][y][z].Id, originalId) == 0 {
				t.Errorf("Square{Id: %s} didn't rotate", originalId )
			}
		}
	}



}