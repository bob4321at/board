package main

import (
	"board/camera"
	"board/grid"
	"board/pieces"
	"board/ui"
	"board/utils"
	"image/color"
	"strconv"

	"github.com/ebitengine/debugui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	debugui             debugui.DebugUI
	InputCapturingState debugui.InputCapturingState
}

var Is_Editing bool = true
var Tab_hit bool

func (g *Game) Update() error {
	mx, my := ebiten.CursorPosition()
	utils.Mouse_X = float64(mx)
	utils.Mouse_Y = float64(my)
	sx, sy := ebiten.Wheel()
	utils.Scroll_X = sx
	utils.Scroll_Y = sy

	if Is_Editing {
		temp_input_capture_state, err := g.debugui.Update(func(ctx *debugui.Context) error {
			ui.EditMenu(ctx)
			return nil
		})
		if err != nil {
			panic(err)
		}

		g.InputCapturingState = temp_input_capture_state

		if g.InputCapturingState == 0 {
			pieces.Hovering = false

			for i := range pieces.Pieces {
				piece := &pieces.Pieces[i]
				piece.Edit_Update()
			}

			if !pieces.Hovering {
				grid.Temp_Grid.Update()
			}
		}
	} else {
		for i := range pieces.Pieces {
			piece := &pieces.Pieces[i]
			piece.Game_Update()
		}

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

	for _, piece := range pieces.Pieces {
		piece.Draw(screen, camera.Cam)
	}

	ebitenutil.DebugPrintAt(screen, "Hit Tab For Tools", 10, 1080-20)
	g.debugui.Draw(screen)
	ebitenutil.DebugPrint(screen, "FPS: "+strconv.Itoa(int(ebiten.ActualFPS())))
}

func (g *Game) Layout(ow, oh int) (sw, sh int) {
	return 1920, 1080
}

func main() {
	ebiten.SetWindowSize(1920, 1080)
	ebiten.SetWindowTitle("board")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	img1, _, err := ebitenutil.NewImageFromFile("./art/checker_red.png")
	if err != nil {
		panic(err)
	}
	img2, _, err := ebitenutil.NewImageFromFile("./art/checker_blue.png")
	if err != nil {
		panic(err)
	}
	pieces.Pieces = append(pieces.Pieces, pieces.NewPiece(utils.Vec2{X: 100, Y: 0}, img1))
	pieces.Pieces = append(pieces.Pieces, pieces.NewPiece(utils.Vec2{X: 0, Y: 0}, img2))

	op := ebiten.RunGameOptions{}
	if err := ebiten.RunGameWithOptions(&Game{}, &op); err != nil {
		panic(err)
	}
}
