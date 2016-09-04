package example

// import "fmt"

// MovementSystem adjusts the Position of an entity which also has a Movement
func MovementSystem(pos *Position, movement *Movement){
    delta := movement.Delta()
    pos.X += delta.X
    pos.Y += delta.Y
}

// RenderSystem reports the position of an entity in the world
func RenderSystem(pos *Position){
    // fmt.Printf("(%v:%v)\n", pos.X, pos.Y)
}
