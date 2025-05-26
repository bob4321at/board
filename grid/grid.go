package grid

import (
	"board/camera"
	"board/utils"
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Grid struct {
	Size       utils.Vec2
	Tiles      [][]Tile
	Grid_Image *ebiten.Image
}

func NewGrid(width, height int, color_one, color_two color.RGBA) (grid Grid) {
	grid.Size = utils.Vec2{X: float64(width), Y: float64(height)}

	for y := range height {
		grid.Tiles = append(grid.Tiles, []Tile{})
		for x := range width {
			grid.Tiles[y] = append(grid.Tiles[y], Tile{})
			col := color_one
			if math.Mod(float64(y+x), 2) == 0 {
				col = color_two
			}
			grid.Tiles[y][x] = NewTile(utils.Vec2{X: float64(x), Y: float64(y)}, col)
		}
	}

	grid.GenCache()

	return grid
}

func (grid *Grid) GenCache() {
	grid.Grid_Image = ebiten.NewImage(int(grid.Size.X*32), int(grid.Size.Y*32))

	tile_img := ebiten.NewImage(32, 32)

	for y, list := range grid.Tiles {
		for x, tile := range list {
			tile_img.Fill(tile.Color)

			op := ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*32), float64(y*32))

			grid.Grid_Image.DrawImage(tile_img, &op)
		}
	}
}

func (grid *Grid) Draw(screen *ebiten.Image, cam camera.Camera) {
	geom := utils.GeoM{}
	camera.DrawWithCamera(screen, cam, grid.Grid_Image, geom)
}

func (grid *Grid) Update() {
	grid.Tiles = NewGrid(20, 20, color.RGBA{125, 125, 125, 255}, color.RGBA{255, 255, 255, 255}).Tiles
	for y := range grid.Tiles {
		for x, _ := range grid.Tiles[y] {
			fmt.Println(int(int(camera.Cam.Pos.X)+1920+(x*32*int(camera.Cam.Zoom))) / 32)
			fmt.Println(int(utils.Mouse_X / 32))
			if (int(camera.Cam.Pos.X)+1920+(x*32*int(camera.Cam.Zoom)))/32 == int(utils.Mouse_X)/32 {
				grid.Tiles[y][x].Color = color.RGBA{255, 0, 255, 255}
			}
		}
	}
	grid.GenCache()
}

var Temp_Grid = NewGrid(20, 20, color.RGBA{125, 125, 125, 255}, color.RGBA{255, 255, 255, 255})
