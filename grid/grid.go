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

var Brush_Color utils.IColor

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
	temp_img := ebiten.NewImage(int(grid.Size.X*32), int(grid.Size.Y*32))

	tile_img := ebiten.NewImage(32, 32)

	for y, list := range grid.Tiles {
		for x, tile := range list {
			tile_img.Fill(tile.Color)

			op := ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*32), float64(y*32))

			temp_img.DrawImage(tile_img, &op)
		}
	}

	grid.Grid_Image = temp_img
}

func (grid *Grid) Draw(screen *ebiten.Image, cam camera.Camera) {
	geom := utils.GeoM{}
	camera.DrawWithCamera(screen, cam, grid.Grid_Image, geom)
}

func (grid *Grid) Update() {
	// grid.Tiles = NewGrid(int(grid.Size.X), int(grid.Size.Y), color.RGBA{125, 125, 125, 255}, color.RGBA{255, 255, 255, 255}).Tiles

	// for y := range grid.Tiles {
	// 	for x, _ := range grid.Tiles[y] {
	// 		if (int(camera.Cam.Pos.X)+(x*32*int(camera.Cam.Zoom)))/32 == int(utils.Mouse_X)/32 {
	// 			if (int(camera.Cam.Pos.Y)+(y*32*int(camera.Cam.Zoom)))/32 == int(utils.Mouse_Y)/32 {
	// 				grid.Tiles[y][x].Color = Brush_Color.TurnToColorRGBA()
	// 				// fmt.Println(int(int(camera.Cam.Pos.X)+int(960*camera.Cam.Zoom)+(x*32*int(camera.Cam.Zoom))) / 32)
	// 				// fmt.Println(int(utils.Mouse_X / 32))
	// 			}
	// 		}
	// 	}
	// }

	// tile_to_screen_pos := (int(int(camera.Cam.Pos.X)+int(960*camera.Cam.Zoom)+(x*32*int(camera.Cam.Zoom))) / 32)
	// mouse_to_tile_pos := (int(utils.Mouse_X / 32))

	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		world_mouse_x := float64(camera.Cam.Pos.X+960-(utils.Mouse_X*camera.Cam.Zoom)) / 32
		world_mouse_y := float64(camera.Cam.Pos.Y+540-(utils.Mouse_Y*camera.Cam.Zoom)) / 32

		world_mouse_x += float64(len(grid.Tiles[0]) / 2)
		world_mouse_y += float64(len(grid.Tiles) / 2)

		fmt.Println(world_mouse_x)
		fmt.Println(world_mouse_y)

		if int(world_mouse_y) < len(grid.Tiles) && world_mouse_y > 0 {
			if int(world_mouse_x) < len(grid.Tiles[0]) && world_mouse_x > 0 {
				grid.Tiles[len(grid.Tiles)-int(world_mouse_y)-1][len(grid.Tiles[0])-int(world_mouse_x)-1].Color = Brush_Color.TurnToColorRGBA()
				grid.GenCache()
			}
		}

		fmt.Println(world_mouse_x)
		fmt.Println(world_mouse_y)
	}
}

var Temp_Grid = NewGrid(20, 20, color.RGBA{125, 125, 125, 255}, color.RGBA{255, 255, 255, 255})
