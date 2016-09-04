package example

import (
    "github.com/klavs/ecs"
)

// MyUniverse sets up my instance of the Universe
func MyUniverse() *ecs.Universe{
    // it is a universe
    universe := &ecs.Universe{}
    // which knows how to move entities
    universe.AddSystem(MovementSystem)
    // and render them
    universe.AddSystem(RenderSystem)

    // let's add a bunch of Balls and Items
    counter := make([]interface{}, 38500)
    for range counter {
        // create a ball whith some Position and Movement 
        ball := Ball{
            Position{12,24},
            Movement{3,Direction{1,2}},
        }
        // and add it to the universe
        universe.AddEntity(&ball)

        // create a StaticItem with some position
        item := StaticItem{
            Position{10,10},
        }
        // and add it to the 
        universe.AddEntity(&item)

        // and the entity does not even need to be predefined to be used
        universe.AddEntity(&struct{
            Position
            Movement
        }{
            Position{1,2},
            Movement{7, Direction{0.3,1}},
        })
    }

    // return the universe to be used externally
    return universe
}

