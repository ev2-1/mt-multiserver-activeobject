// this file contains some functions for creating active objects
package aoTools

import (
	"github.com/anon55555/mt"

	"image/color"
)

// AOPos creates a IDAOMsg setting the pos of ao`id` to `pos`
func AOPos(id mt.AOID, pos mt.AOPos) mt.IDAOMsg {
	return mt.IDAOMsg{
		ID: id,
		Msg: &mt.AOCmdPos{
			Pos: pos,
		},
	}
}

// AddAO creates a AOAdd with specified AOID, POS and AOProps
func AddAO(id mt.AOID, p mt.Pos, props mt.AOProps) mt.AOAdd {
	return mt.AOAdd{
		ID: id,
		InitData: mt.AOInitData{
			ID:  id,
			Pos: p,

			Msgs: []mt.AOMsg{
				&mt.AOCmdProps{
					Props: props,
				},
				&mt.AOCmdAttach{},
			},
		},
	}
}

// CubeAO creates a AOProps object for a cube
// name can be anything that has a "tnt"-like texture naming scheme
// mediaPool is the mediapool the textures are contained in
func CubeAO(mediaPool, name string) mt.AOProps {
	return mt.AOProps{
		Mesh:      "",
		MaxHP:     10,
		Pointable: true,
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
		Textures:        []mt.Texture{
			mt.Texture(mediaPool + "_default_" + name + "_top.png"),
			mt.Texture(mediaPool + "_default_" + name + "_bottom.png"),
			mt.Texture(mediaPool + "_default_" + name + "_side.png"),
			mt.Texture(mediaPool + "_default_" + name + "_side.png"),
			mt.Texture(mediaPool + "_default_" + name + "_side.png"),
			mt.Texture(mediaPool + "_default_" + name + "_side.png"),
		},
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
