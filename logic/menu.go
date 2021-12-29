package logic

import (
	"github.com/rastasi/omegon-pac/lang"
	"github.com/rastasi/omegon-pac/manifest"
	"github.com/rastasi/omegon-pac/types"
)

func MasterMenuList() {
	list := types.MenuList{
		Elements: []types.MenuListElement{
			{
				Title: lang.T("menu.master.newGame"),
			},
			{
				Title: lang.T("menu.master.credits"),
			},
			{
				Title: lang.T("menu.master.exit"),
			},
		},
	}
	manifest.Menu(list)
}
