package comunication

import (
	"board/grid"
	"board/pieces"
	"board/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"image/color"
	"io"
	"net/http"

	"github.com/hajimehoshi/ebiten/v2"
)

type NetworkedGrid struct {
	Size  []uint16
	Tiles [][]grid.Tile
}

type NetworkedPiece struct {
	Position []float64
	Image    [][]color.RGBA
}

type SendPiecesBackSturct struct {
	Pieces []NetworkedPiece
}

type Piece struct {
	Position utils.Vec2
	Image    [][]color.RGBA
}

type ChangeNetworkedPieceStruct struct {
	ID    uint8
	Piece NetworkedPiece
}

type ChangedPiece struct {
	ID       uint8
	Position [2]float64
	Image    [][]color.RGBA
}

type ListOfChangedPiece struct {
	Pieces []ChangedPiece
}

type User struct {
	ID          uint8
	Got_Changes bool
}

var Server_To_Join string = "localhost:8080"

var In_Server bool = false
var ID uint8 = 0
var Got_Changes bool = false

var Pieces_To_Change = ListOfChangedPiece{}
var Changes_Made_To_Pieces = ListOfChangedPiece{}

func CheckChanges() {
	if In_Server {
		for _, piece := range Changes_Made_To_Pieces.Pieces {
			temp_piece_json, err := json.Marshal(piece)
			if err != nil {
				panic(err)
			}

			fmt.Println(string(temp_piece_json))

			if _, err := http.Post("http://"+Server_To_Join+"/ChangePiece", "application/json", bytes.NewBuffer(temp_piece_json)); err != nil {
				panic(err)
			}
		}

		Changes_Made_To_Pieces.Pieces = nil

		data_in_bytes, err := json.Marshal(User{ID, false})
		if err != nil {
			panic(err)
		}

		resp, err := http.Post("http://"+Server_To_Join+"/CheckForChangeForUser", "application/json", bytes.NewBuffer(data_in_bytes))
		if err != nil {
			panic(err)
		}

		data_in_bytes, err = io.ReadAll(resp.Body)
		temp_user := User{}

		json.Unmarshal(data_in_bytes, &temp_user)

		Got_Changes = temp_user.Got_Changes

		if !temp_user.Got_Changes {
			resp, err := http.Get("http://" + Server_To_Join + "/GetPieceChanges")
			if err != nil {
				panic(err)
			}

			json_data := ListOfChangedPiece{}
			json_data_bytes, err := io.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}

			json.Unmarshal(json_data_bytes, &json_data)

			Pieces_To_Change = json_data
		}
	}
}

func SendBoard() {
	data_in_bytes, err := json.Marshal(grid.Temp_Grid)
	if err != nil {
		panic(err)
	}

	if _, err := http.Post("http://"+Server_To_Join+"/SendBoardToServer", "application/json", bytes.NewBuffer(data_in_bytes)); err != nil {
		panic(err)
	}
}

func SendPieces() {
	saved_pieces_data := []Piece{}

	for _, piece := range pieces.Pieces {
		colors := [][]color.RGBA{}

		for x := range piece.Image.Bounds().Max.X {
			colors = append(colors, []color.RGBA{})
			for y := range piece.Image.Bounds().Max.Y {
				colo := piece.Image.At(x, y)
				r, g, b, a := colo.RGBA()
				col := color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
				colors[x] = append(colors[x], col)
			}
		}

		saved_pieces_data = append(saved_pieces_data, Piece{piece.Position, colors})
	}

	data_to_send, err := json.Marshal(saved_pieces_data)
	if err != nil {
		panic(err)
	}

	if _, err := http.Post("http://"+Server_To_Join+"/GivePiecesToServer", "application/json", bytes.NewBuffer(data_to_send)); err != nil {
		panic(err)
	}
}

func AddUser() {
	resp, err := http.Get("http://" + Server_To_Join + "/AddUser")
	if err != nil {
		panic(err)
	}

	temp_user := User{}
	json_data_bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(json_data_bytes, &temp_user)

	ID = temp_user.ID
}

func GetBoard() {
	resp, err := http.Get("http://" + Server_To_Join + "/GetBoardFromServer")
	if err != nil {
		panic(err)
	}
	grid_data := NetworkedGrid{}

	grid_data_bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(grid_data_bytes, &grid_data); err != nil {
		panic(err)
	}

	grid.Temp_Grid.Size = utils.Vec2{X: float64(grid_data.Size[0]), Y: float64(grid_data.Size[1])}
	grid.Temp_Grid.Tiles = grid_data.Tiles
	grid.Temp_Grid.GenCache()
}

func GetPieces() {
	resp, err := http.Get("http://" + Server_To_Join + "/GetPiecesFromServer")
	if err != nil {
		panic(err)
	}
	pieces_data := SendPiecesBackSturct{}

	pieces_data_bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(pieces_data_bytes, &pieces_data); err != nil {
		panic(err)
	}

	saved_piece_data := []pieces.Piece{}

	for _, piece := range pieces_data.Pieces {
		img := ebiten.NewImage(16, 16)

		for x := range piece.Image {
			for y := range piece.Image[x] {
				img.Set(x, y, piece.Image[x][y])
			}
		}

		saved_piece_data = append(saved_piece_data, pieces.Piece{Position: utils.Vec2{X: piece.Position[0], Y: piece.Position[1]}, Started_Click_Position: utils.Vec2{X: 0, Y: 0}, Clicked: 0, Image: img})
	}

	pieces.Pieces = saved_piece_data
}
