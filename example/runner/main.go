package main

import (
    "github.com/klavs/ecs/example"
    "time"
    "fmt"
)


func main(){
    s0 := time.Now()
    // get the universe defined in MyUniverse with all the systems and entities
    mu := example.MyUniverse()
    s1 := time.Now()
    fmt.Printf("Setup took %v to run.\n", s1.Sub(s0))

    t0 := time.Now()

    // measure 60 iterations (adjust entity count in MyUniverse and)
    for i:=0; i < 60; i++ {
        mu.Process()
    }

    t1 := time.Now()
    fmt.Printf("Processing took %v to run.\n", t1.Sub(t0))

}