package entity

type Model []float32

type Player struct {
	Model Model
}

func NewPlayer() *Player {
	p := Player{}
	p.Model = []float32{
		//  X, Y, Z, U, V

		-1.0, -1.0, 1.0, 1.0, 0.0,
		1.0, -1.0, 1.0, 0.0, 0.0,
		-1.0, 1.0, 1.0, 1.0, 1.0,
		1.0, -1.0, 1.0, 0.0, 0.0,
		1.0, 1.0, 1.0, 0.0, 1.0,
		-1.0, 1.0, 1.0, 1.0, 1.0,
	}

	return &p
}
