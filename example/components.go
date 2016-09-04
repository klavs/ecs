package example

// Position represents the entities position in the world
type Position Vector2f

// Movement represents speed and direction of an entity
type Movement struct {
    Speed float64
    Direction
}

// Delta calculates the change in Position when this Movement is applied
func (m *Movement) Delta() Vector2f {
    return Vector2f{
        m.X * m.Speed,
        m.Y * m.Speed,
    }
}