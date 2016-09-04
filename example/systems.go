package example


func MovementSystem(pos *Position, movement *Movement){
    delta := movement.Delta()
    pos.X += delta.X
    pos.Y += delta.Y
}

func RenderSystem(pos *Position){}
