package main

import (
	"github.com/anon55555/mt"

	"image/color"
)

func TntAdd(id mt.AOID, p mt.Pos) mt.AOAdd {
	return mt.AOAdd{
		ID: id,
		InitData: mt.AOInitData{
			ID:  id,
			Pos: p,

			Msgs: []mt.AOMsg{
				&mt.AOCmdProps{
					Props: Tnt(),
				},
				&mt.AOCmdAttach{},
			},
		},
	}
}

func Tnt() mt.AOProps {
	return mt.AOProps{
		Mesh:      "",
		MaxHP:     10,
		Pointable: false,
		ColBox: mt.Box{
			mt.Vec{-0.5, -0.5, -0.5},
			mt.Vec{0.5, 0.5, 0.5},
		},
		SelBox: mt.Box{
			mt.Vec{-0.5, -0.5, -0.5},
			mt.Vec{0.5, 0.5, 0.5},
		},
		Visual:          "cube",
		VisualSize:      [3]float32{1.0, 1.0, 1.0},
		Textures:        []mt.Texture{"micl2_default_tnt_top.png", "micl2_default_tnt_bottom.png", "micl2_default_tnt_side.png", "micl2_default_tnt_side.png", "micl2_default_tnt_side.png", "micl2_default_tnt_side.png"},
		DmgTextureMod:   "^[brighten",
		Shaded:          true,
		SpriteSheetSize: [2]int16{1, 1},
		SpritePos:       [2]int16{0, 0},
		Visible:         true,
		Colors:          []color.NRGBA{color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}},
		BackfaceCull:    true,
		NametagColor:    color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},
		NametagBG:       color.NRGBA{R: 0x01, G: 0x01, B: 0x01, A: 0x00},
		FaceRotateSpeed: -1,
		Infotext:        "",
		Itemstring:      "",
	}
}
