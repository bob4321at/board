package pieces

import (
	"board/camera"
	"board/utils"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
)

var Holding bool = false
var Hovering bool = false

type Piece struct {
	Position               utils.Vec2
	Started_Click_Position utils.Vec2
	Clicked                int
	Image                  *ebiten.Image
}

func NewPiece(Position utils.Vec2, Image *ebiten.Image) (piece Piece) {
	piece.Position = Position
	piece.Image = ebiten.NewImageFromImage(Image)

	return piece
}

func (piece *Piece) Draw(screen *ebiten.Image, cam camera.Camera) {
	geom := utils.GeoM{}
	geom.Translate(piece.Position.X, piece.Position.Y)
	camera.DrawWithCamera(screen, cam, piece.Image, geom)
}

func (piece *Piece) Edit_Update() {
	Mouse_World_Pos := utils.Vec2{X: utils.Mouse_X, Y: utils.Mouse_Y}

	Mouse_World_Pos.X -= 960
	Mouse_World_Pos.Y -= 540

	Mouse_World_Pos.X -= camera.Cam.Pos.X
	Mouse_World_Pos.Y -= camera.Cam.Pos.Y

	Mouse_World_Pos.X /= camera.Cam.Zoom
	Mouse_World_Pos.Y /= camera.Cam.Zoom

	if Mouse_World_Pos.X > piece.Position.X-(float64(piece.Image.Bounds().Dx())/2) && Mouse_World_Pos.X < piece.Position.X+(float64(piece.Image.Bounds().Dx())/2) {
		if Mouse_World_Pos.Y > piece.Position.Y-(float64(piece.Image.Bounds().Dy())/2) && Mouse_World_Pos.Y < piece.Position.Y+(float64(piece.Image.Bounds().Dy())/2) {
			Hovering = true
			if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) && !Holding {
				Pieces = append(Pieces, NewPiece(utils.Vec2{X: piece.Position.X + (rand.Float64() * 64) - 32, Y: piece.Position.Y + (rand.Float64() * 64) - 32}, piece.Image))
				Holding = true
			}
		}
	}

	if !ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		Holding = false
	}
}

func (piece *Piece) Game_Update() {
	Mouse_World_Pos := utils.Vec2{X: utils.Mouse_X, Y: utils.Mouse_Y}

	Mouse_World_Pos.X -= 960
	Mouse_World_Pos.Y -= 540

	Mouse_World_Pos.X -= camera.Cam.Pos.X
	Mouse_World_Pos.Y -= camera.Cam.Pos.Y

	Mouse_World_Pos.X /= camera.Cam.Zoom
	Mouse_World_Pos.Y /= camera.Cam.Zoom

	if Mouse_World_Pos.X > piece.Position.X-(float64(piece.Image.Bounds().Dx())/2) && Mouse_World_Pos.X < piece.Position.X+(float64(piece.Image.Bounds().Dx())/2) {
		if Mouse_World_Pos.Y > piece.Position.Y-(float64(piece.Image.Bounds().Dy())/2) && Mouse_World_Pos.Y < piece.Position.Y+(float64(piece.Image.Bounds().Dy())/2) {
			Hovering = true
			if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) && !Holding {
				piece.Clicked = 1
				Holding = true
			}
		}
	}

	if piece.Clicked == 1 {
		piece.Clicked = 2
		// piece.Started_Click_Position = utils.Vec2{X: utils.Mouse_X - piece.Position.X, Y: utils.Mouse_Y - piece.Position.Y}
	}

	if piece.Clicked == 2 {
		piece.Position.X = -(piece.Started_Click_Position.X - Mouse_World_Pos.X)
		piece.Position.Y = -(piece.Started_Click_Position.Y - Mouse_World_Pos.Y)
	}

	if !ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		piece.Clicked = 0
		Holding = false
	}
}

var Pieces = []Piece{}
