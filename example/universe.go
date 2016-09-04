package example

import (
    "github.com/klavs/ecs"
)

func MyUniverse() *ecs.Universe{
    universe := ecs.NewUniverse()
    universe.AddSystem(MovementSystem)
    universe.AddSystem(RenderSystem)
    universe.AddSystem(RenderSystem)

    ball := Ball{
        Position{12,24},
        Movement{3,Direction{1,2}},
    }

    item := StaticItem{
        Position{10,10},
    }

    universe.AddEntity(&ball)
    universe.AddEntity(&item)

    return universe
}

