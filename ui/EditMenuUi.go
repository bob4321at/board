package ui

import (
	"board/grid"
	"board/utils"
	"image"
	"strconv"

	"github.com/ebitengine/debugui"
	"github.com/hajimehoshi/ebiten/v2"
)

var Grid_Size string = "10"

var Brush_Color = utils.IColor{R: 255, G: 255, B: 255, A: 255}
var Tile_Color_One = utils.IColor{R: 255, G: 0, B: 255, A: 255}
var Tile_Color_Two = utils.IColor{R: 255, G: 0, B: 255, A: 255}

func EditMenu(ctx *debugui.Context) {
	ctx.Window("Edit", image.Rect(0, 0, 200, 1000), func(layout debugui.ContainerLayout) {
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

		ctx.Header("Color", true, func() {
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
	})
}
