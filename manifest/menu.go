package manifest

import (
	"fmt"

	"github.com/rastasi/omegon-pac/types"
)

func Menu(list types.MenuList) {
	fmt.Println(list)
}

func Game(grid types.Grid) {
	fmt.Println((grid))
}
