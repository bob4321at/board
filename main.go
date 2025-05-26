package main

import (
	"board/camera"
	"board/grid"
	"board/ui"
	"board/utils"
	"image/color"

	"github.com/ebitengine/debugui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	debugui debugui.DebugUI
}

var Is_Editing bool
var Tab_hit bool

func (g *Game) Update() error {
	mx, my := ebiten.CursorPosition()
	utils.Mouse_X = float64(mx)
	utils.Mouse_Y = float64(my)

	grid.Temp_Grid.Update()

	if Is_Editing {
		if _, err := g.debugui.Update(func(ctx *debugui.Context) error {
			ui.EditMenu(ctx)
			return nil
		}); err != nil {
			panic(err)
		}
	} else {
		if _, err := g.debugui.Update(func(ctx *debugui.Context) error {
			return nil
		}); err != nil {
			panic(err)
		}
	}

	if !ebiten.IsKeyPressed(ebiten.KeyTab) {
		Tab_hit = false
	}

	if ebiten.IsKeyPressed(ebiten.KeyTab) && !Tab_hit {
		Is_Editing = !Is_Editing
		Tab_hit = true
	}

	camera.Cam.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{100, 100, 100, 255})
	grid.Temp_Grid.Draw(screen, camera.Cam)
	ebitenutil.DebugPrintAt(screen, "Hit Tab For Tools", 10, 1080-20)
	g.debugui.Draw(screen)

	test_img := ebiten.NewImage(10, 10)
	test_img.Fill(color.RGBA{255, 0, 0, 255})

	camera.DrawWithCamera(screen, camera.Cam, test_img, utils.GeoM{})
	second_point_geom := utils.GeoM{}
	second_point_geom.Translate(10*32, 10*32)
	camera.DrawWithCamera(screen, camera.Cam, test_img, second_point_geom)
}

func (g *Game) Layout(ow, oh int) (sw, sh int) {
	return 1920, 1080
}

func main() {
	ebiten.SetWindowSize(1920, 1080)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
