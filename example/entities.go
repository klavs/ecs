package example

// Ball is an entity which moves around the world
type Ball struct {
    Position
    Movement
}

// StaticItem is an entity that exists in the world but does not change it's position
type StaticItem struct {
    Position
}
