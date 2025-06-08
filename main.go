package main

import (
	"board/camera"
	"board/comunication"
	"board/grid"
	"board/pieces"
	"board/ui"
	"board/utils"
	"image"

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

	for i, piece := range pieces.Pieces {
		if piece.Changed {
			changed_piece := comunication.ChangedPiece{}
			changed_piece.ID = uint8(i)

			colors := [][]color.RGBA{}

			if piece.Image != nil {
				for x := range piece.Image.Bounds().Max.X {
					colors = append(colors, []color.RGBA{})
					for y := range piece.Image.Bounds().Max.Y {
						colo := piece.Image.At(x, y)
						r, g, b, a := colo.RGBA()
						col := color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
						colors[x] = append(colors[x], col)
					}
				}
			}

			changed_piece.Image = colors
			changed_piece.Position = [2]float64{piece.Position.X, piece.Position.Y}
			comunication.Changes_Made_To_Pieces.Pieces = append(comunication.Changes_Made_To_Pieces.Pieces, changed_piece)

			piece.Changed = false
		}
	}

	for i := range comunication.Pieces_To_Change.Pieces {
		piece := comunication.Pieces_To_Change.Pieces[i]
		op := ebiten.NewImageOptions{}
		op.Unmanaged = true
		img := ebiten.NewImage(16, 16)

		for x := range piece.Image {
			for y := range piece.Image[x] {
				img.Set(x, y, piece.Image[x][y])
			}
		}

		pieces.Pieces[piece.ID] = pieces.Piece{Position: utils.Vec2{X: piece.Position[0], Y: piece.Position[1]}, Started_Click_Position: pieces.Pieces[piece.ID].Started_Click_Position, Clicked: pieces.Pieces[piece.ID].Clicked, Image: img}
	}

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

	_, icon_img, err := ebitenutil.NewImageFromFile("./art/icon.png")
	if err != nil {
		panic(err)
	}
	ebiten.SetWindowIcon([]image.Image{icon_img})

	go func() {
		for true {
			comunication.CheckChanges()
		}
	}()

	// op := ebiten.RunGameOptions{}
	// if err := ebiten.RunGameWithOptions(&Game{}, &op); err != nil {
	// panic(err)
	// }
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
