package grid

import (
	"board/utils"
	"image/color"
)

type Tile struct {
	Pos   utils.Vec2
	Color color.RGBA
}

func NewTile(Pos utils.Vec2, Color color.RGBA) (tile Tile) {
	tile.Pos = Pos
	tile.Color = Color

	return tile
}
