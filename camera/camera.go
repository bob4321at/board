package camera

import (
	"board/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type Camera struct {
	Pos                     utils.Vec2
	Speed                   float64
	Zoom                    float64
	Is_Middle_Clicking      bool
	Zoom_Middle_Click_Start utils.Vec2
	Zoom_Middle_Click       utils.Vec2
	Started_Moving          utils.Vec2
	Moving                  bool
}

func NewCamera(Pos utils.Vec2) (camera Camera) {
	camera.Pos = Pos
	camera.Zoom = 3
	camera.Speed = 10

	return camera
}

func DrawWithCamera(screen *ebiten.Image, cam Camera, image_to_render *ebiten.Image, geom utils.GeoM) {
	offset_x := geom.X * cam.Zoom
	offset_y := geom.Y * cam.Zoom

	new_op := ebiten.DrawImageOptions{}
	new_op.GeoM.Translate(-(float64(image_to_render.Bounds().Dx()) / 2), -(float64((image_to_render.Bounds().Dy()) / 2)))
	new_op.GeoM.Scale(cam.Zoom*geom.Width, cam.Zoom*geom.Height)
	new_op.GeoM.Translate(cam.Pos.X+960+offset_x, cam.Pos.Y+540+offset_y)
	screen.DrawImage(image_to_render, &new_op)
}

func (camera *Camera) Update() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) && !camera.Moving {
		camera.Moving = true
		camera.Started_Moving = utils.Vec2{X: utils.Mouse_X - camera.Pos.X, Y: utils.Mouse_Y - camera.Pos.Y}
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) && camera.Moving {
		camera.Pos.X = -(camera.Started_Moving.X - utils.Mouse_X)
		camera.Pos.Y = -(camera.Started_Moving.Y - utils.Mouse_Y)
	}

	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) && camera.Moving {
		camera.Moving = false
	}

	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle) && camera.Is_Middle_Clicking {
		camera.Zoom_Middle_Click_Start = utils.Vec2{X: 0, Y: 0}
		camera.Zoom_Middle_Click = utils.Vec2{X: 0, Y: 0}
		camera.Is_Middle_Clicking = false
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle) && !camera.Is_Middle_Clicking {
		camera.Is_Middle_Clicking = true
		camera.Zoom_Middle_Click_Start = utils.Vec2{X: utils.Mouse_X, Y: utils.Mouse_Y}
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle) && camera.Is_Middle_Clicking {
		camera.Zoom_Middle_Click = utils.Vec2{X: utils.Mouse_X, Y: utils.Mouse_Y}
	}

	camera.Zoom -= (camera.Zoom_Middle_Click.Y - camera.Zoom_Middle_Click_Start.Y) / 10000

	camera.Zoom += utils.Scroll_Y / 50

	if camera.Zoom < 0.1 {
		camera.Zoom = 0.1
	}

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
