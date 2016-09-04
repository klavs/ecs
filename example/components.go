package example

type Position Vector2f

type Movement struct {
    Speed float64
    Direction
}

func (m *Movement) Delta() Vector2f {
    return Vector2f{
        m.X * m.Speed,
        m.Y * m.Speed,
    }
}