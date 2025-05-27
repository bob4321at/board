package camera

import (
	"board/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type Camera struct {
	Pos   utils.Vec2
	Speed float64
	Zoom  float64
}

func NewCamera(Pos utils.Vec2) (camera Camera) {
	camera.Pos = Pos
	camera.Zoom = 2
	camera.Speed = 10

	return camera
}

func DrawWithCamera(screen *ebiten.Image, cam Camera, image_to_render *ebiten.Image, geom utils.GeoM) {
	offset_x := geom.X * cam.Zoom
	offset_y := geom.Y * cam.Zoom

	new_op := ebiten.DrawImageOptions{}
	new_op.GeoM.Translate(-(float64(image_to_render.Bounds().Dx()) / 2), -(float64(image_to_render.Bounds().Dy() / 2)))
	new_op.GeoM.Scale(cam.Zoom, cam.Zoom)
	new_op.GeoM.Translate(cam.Pos.X+960+offset_x, cam.Pos.Y+540+offset_y)
	screen.DrawImage(image_to_render, &new_op)
}

func (camera *Camera) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyH) {
		camera.Pos.X += 1 * camera.Speed * camera.Zoom
	} else if ebiten.IsKeyPressed(ebiten.KeyL) {
		camera.Pos.X -= 1 * camera.Speed * camera.Zoom
	}

	if ebiten.IsKeyPressed(ebiten.KeyJ) {
		camera.Pos.Y -= 1 * camera.Speed * camera.Zoom
	} else if ebiten.IsKeyPressed(ebiten.KeyK) {
		camera.Pos.Y += 1 * camera.Speed * camera.Zoom
	}

	if ebiten.IsKeyPressed(ebiten.KeyI) {
		camera.Zoom += 0.01 * camera.Zoom
	} else if ebiten.IsKeyPressed(ebiten.KeyO) {
		camera.Zoom -= 0.01 * camera.Zoom
	}
}

var Cam = NewCamera(utils.Vec2{X: 0, Y: 0})
