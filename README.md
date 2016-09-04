# ecs
This is an experimental Entity Component System.
The goal is to use as much built in Golang features as possible.

An `Entity` is just a regular struct with some anonymous fields (`Components`).

A `System` is a function with pointers to `Components` as arguments.

The `Universe` is a container of `Systems` and `Entities`. `Universe` allows `Systems` and `Entities` to be added. `Universe` makes sure that each `System` works only with `Entities` that contain the `Components` defined in `System's` arguments.

The performance of a `Universe` depends directly on the number of connections from a `System` to `Entity`.

This repo contains a simple example in `example` directory.
`example/runner/main.go` is the starting point.
