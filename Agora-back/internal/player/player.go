package player

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Player struct {
	RoomId string
	Position Position
	Connection *websocket.Conn
	sync.Mutex
}

func NewPlayer(wsConnection *websocket.Conn, roomId string) *Player {
	return &Player{
		Position: Position{X: 0, Y: 0},
		Connection: wsConnection,
		RoomId: roomId,
	}
}

func (p *Player) MoveUp() {
	p.Position = Position{X: p.Position.X, Y: p.Position.Y + 3}
}

func (p *Player) MoveDown() {
	p.Position = Position{X: p.Position.X, Y: p.Position.Y - 3}
}

func (p *Player) MoveLeft() {
	p.Position = Position{X: p.Position.X + 3, Y: p.Position.Y}
}

func (p *Player) MoveRight() {
	p.Position = Position{X: p.Position.X - 3, Y: p.Position.Y}
}