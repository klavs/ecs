package main

import (
    "github.com/klavs/ecs/example"
    "time"
    "fmt"
)


func main(){
    mu := example.MyUniverse()

    t0 := time.Now()

    for i:=0; i < 100000; i++ {
        mu.Process()
    }

    t1 := time.Now()
    fmt.Printf("The call took %v to run.\n", t1.Sub(t0))

}