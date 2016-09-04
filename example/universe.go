package example

import (
    "github.com/klavs/ecs"
)

func MyUniverse() *ecs.Universe{
    universe := ecs.NewUniverse()
    universe.AddSystem(MovementSystem)
    universe.AddSystem(RenderSystem)
    universe.AddSystem(RenderSystem)

    counter := make([]interface{}, 512)
    for range counter {
        ball := Ball{
            Position{12,24},
            Movement{3,Direction{1,2}},
        }
        universe.AddEntity(&ball)
        item := StaticItem{
            Position{10,10},
        }
        universe.AddEntity(&item)
    }

    return universe
}

