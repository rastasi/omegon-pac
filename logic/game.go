package logic

import (
	"github.com/rastasi/omegon-pac/lang"
	"github.com/rastasi/omegon-pac/manifest"
	"github.com/rastasi/omegon-pac/types"
)

func StartGame() {
	grid := types.Grid{
		Fields: []types.Field{
			{
				Name: lang.T("field.start"),
				X:    1,
				Y:    1,
			},
		},
	}
	manifest.Game(grid)
}
