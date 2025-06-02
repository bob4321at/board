package ui

import (
	"board/camera"
	"board/grid"
	"board/pieces"
	"board/utils"
	"image"
	"image/color"
	"strconv"

	"github.com/ebitengine/debugui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var Grid_Size string = "10"

var Brush_Color = utils.IColor{R: 255, G: 255, B: 255, A: 255}
var Tile_Color_One = utils.IColor{R: 255, G: 0, B: 255, A: 255}
var Tile_Color_Two = utils.IColor{R: 255, G: 0, B: 255, A: 255}

var Selected_Piece *pieces.Piece
var Piece_Brush_Color = utils.IColor{R: 255, G: 0, B: 255, A: 255}

func EditMenu(ctx *debugui.Context) {
	ctx.Window("Edit", image.Rect(0, 0, 200, 1000), func(layout debugui.ContainerLayout) {
		EditGridSubMenu(ctx)
		EditPieceSubMenu(ctx)
	})
}

func EditPieceSubMenu(ctx *debugui.Context) {
	ctx.Header("Piece", false, func() {
		ctx.Window("Piece Edit", image.Rect(1920-525, 0, 1920, 1080-160), func(layout debugui.ContainerLayout) {
			ctx.GridCell(func(bounds image.Rectangle) {
				ctx.DrawOnlyWidget(func(screen *ebiten.Image) {
					img, _, err := ebitenutil.NewImageFromFile("./art/edit_area.png")
					if err != nil {
						panic(err)
					}
					op := ebiten.DrawImageOptions{}
					op.GeoM.Translate(float64(bounds.Min.X), float64(bounds.Min.Y))
					screen.DrawImage(img, &op)

					if pieces.Selected_Piece != nil {
						Selected_Piece = pieces.Selected_Piece

						if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
							Selected_Piece.Image.Set(int((utils.Mouse_X-float64(bounds.Min.X))/16-8), int((utils.Mouse_Y-float64(bounds.Min.Y))/16-8), Piece_Brush_Color.TurnToColorRGBA())
						}

						op.GeoM.Reset()
						op.GeoM.Translate(-8, -8)
						op.GeoM.Scale(16, 16)
						op.GeoM.Translate(float64(bounds.Min.X)+256, float64(bounds.Min.Y)+257)
						screen.DrawImage(Selected_Piece.Image, &op)
					}
				})
			})
			ctx.Text("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n " + "Paint Color" + ":")
			ctx.Text("R:")
			ctx.Slider(&Piece_Brush_Color.R, 0, 255, 1)
			ctx.Text("G:")
			ctx.Slider(&Piece_Brush_Color.G, 0, 255, 1)
			ctx.Text("B:")
			ctx.Slider(&Piece_Brush_Color.B, 0, 255, 1)
			ctx.Text("A:")
			ctx.Slider(&Piece_Brush_Color.A, 0, 255, 1)

			ctx.GridCell(func(bounds image.Rectangle) {
				ctx.DrawOnlyWidget(func(screen *ebiten.Image) {
					img := ebiten.NewImage(100, 100)
					img.Fill(Piece_Brush_Color.TurnToColorRGBA())
					op := ebiten.DrawImageOptions{}
					op.GeoM.Translate(float64(bounds.Min.X), float64(bounds.Min.Y))
					screen.DrawImage(img, &op)
				})
			})

			ctx.Text("\n\n\n\n")

			ctx.Button("Eraser").On(func() {
				Piece_Brush_Color = utils.IColor{R: 0, G: 0, B: 0, A: 0}
			})
			ctx.Button("Clear").On(func() {
				Selected_Piece.Image.Fill(color.RGBA{0, 0, 0, 0})
			})
		})
		ctx.Button("Add Piece").On(func() {
			img := ebiten.NewImage(16, 16)
			img.Fill(color.White)
			pieces.Pieces = append(pieces.Pieces, pieces.NewPiece(utils.Vec2{X: -camera.Cam.Pos.X / camera.Cam.Zoom, Y: -camera.Cam.Pos.Y / camera.Cam.Zoom}, ebiten.NewImageFromImage(img)))
		})
	})
}

func EditGridSubMenu(ctx *debugui.Context) {
	ctx.Header("Grid", false, func() {
		ctx.Text("Edit Grid:")

		ctx.Text("\n " + "Tile Color One" + ":")
		ctx.Text("R:")
		ctx.Slider(&Tile_Color_One.R, 0, 255, 1)
		ctx.Text("G:")
		ctx.Slider(&Tile_Color_One.G, 0, 255, 1)
		ctx.Text("B:")
		ctx.Slider(&Tile_Color_One.B, 0, 255, 1)
		ctx.Text("A:")
		ctx.Slider(&Tile_Color_One.A, 0, 255, 1)

		ctx.GridCell(func(bounds image.Rectangle) {
			ctx.DrawOnlyWidget(func(screen *ebiten.Image) {
				img := ebiten.NewImage(100, 100)
				img.Fill(Tile_Color_One.TurnToColorRGBA())
				op := ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(bounds.Min.X), float64(bounds.Min.Y))
				screen.DrawImage(img, &op)
			})
		})

		ctx.Text("\n\n\n\n " + "Tile Color Two" + ":")
		ctx.Text("R:")
		ctx.Slider(&Tile_Color_Two.R, 0, 255, 1)
		ctx.Text("G:")
		ctx.Slider(&Tile_Color_Two.G, 0, 255, 1)
		ctx.Text("B:")
		ctx.Slider(&Tile_Color_Two.B, 0, 255, 1)
		ctx.Text("A:")
		ctx.Slider(&Tile_Color_Two.A, 0, 255, 1)

		ctx.GridCell(func(bounds image.Rectangle) {
			ctx.DrawOnlyWidget(func(screen *ebiten.Image) {
				img := ebiten.NewImage(100, 100)
				img.Fill(Tile_Color_Two.TurnToColorRGBA())
				op := ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(bounds.Min.X), float64(bounds.Min.Y))
				screen.DrawImage(img, &op)
			})
		})

		ctx.Text("\n\n\n\n Size: ")
		ctx.TextField(&Grid_Size)
		ctx.Button("New Grid").On(func() {
			x, err := strconv.Atoi(Grid_Size)
			if err != nil {
				panic(err)
			}
			grid.Temp_Grid = grid.NewGrid(x, x, Tile_Color_One.TurnToColorRGBA(), Tile_Color_Two.TurnToColorRGBA())
		})

		ctx.Text("Brush Color" + ":")
		ctx.Text("R:")
		ctx.Slider(&Brush_Color.R, 0, 255, 1)
		ctx.Text("G:")
		ctx.Slider(&Brush_Color.G, 0, 255, 1)
		ctx.Text("B:")
		ctx.Slider(&Brush_Color.B, 0, 255, 1)
		ctx.Text("A:")
		ctx.Slider(&Brush_Color.A, 0, 255, 1)

		grid.Brush_Color = Brush_Color

		ctx.GridCell(func(bounds image.Rectangle) {
			ctx.DrawOnlyWidget(func(screen *ebiten.Image) {
				img := ebiten.NewImage(100, 100)
				img.Fill(Brush_Color.TurnToColorRGBA())
				op := ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(bounds.Min.X), float64(bounds.Min.Y))
				screen.DrawImage(img, &op)
			})
		})
		ctx.Text("\n\n\n\n")
	})
}
