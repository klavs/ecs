package ecs

import "reflect"
import "runtime"

type system struct {
    fn reflect.Value
    components []reflect.Type
    entitiesArgs [][]reflect.Value
}

func (s *system) process(){
    job := func(argSet [][]reflect.Value, c chan int){
        for _, args := range argSet {
                s.fn.Call(args)
        }
        c <- 1
    }

    var numCPU = runtime.NumCPU()
    c := make(chan int)
	part := len(s.entitiesArgs)/numCPU + 1
	for i:=0; i < numCPU; i++ {
		start := part*i
		end := part*(i+1)
		if end > len(s.entitiesArgs) {
			end = len(s.entitiesArgs)
		}
        go job(s.entitiesArgs[start:end], c)
	}

	for i:=0; i < numCPU; i++ {
        <-c
	}
}

// Universe contains and handles all systems and entities.
type Universe struct {
    systems []*system
}

// AddSystem adds a System to the Universe.
// System is function taking pointers to components as arguments.
func (u *Universe) AddSystem(sys interface{}){
    fn := reflect.ValueOf(sys)
    fnLen := fn.Type().NumIn()
    components := make([]reflect.Type, fnLen, fnLen)
    for i:=0; i < fnLen; i++ {
        components[i] = fn.Type().In(i).Elem()
    }
    u.systems = append(u.systems, &system{fn, components, nil})
}

// AddEntity takes a pointer to a struct (entity) which contains components.
// Components are extracted from entities dynamically and assigned to systems.
func (u *Universe) AddEntity(e interface{}){
    eType := reflect.TypeOf(e).Elem()
    eLen := eType.NumField()
    components := make([]reflect.Type, eLen, eLen)
    for i:=0; i < eLen; i++ {
        components[i] = eType.Field(i).Type
    }

    for _, sys := range u.systems {
        matches := 0
        for _, sysComp := range sys.components {
            for _, entityComp := range components {
                if sysComp == entityComp {
                    matches++
                }
            }
        }
        if matches == len(sys.components) {
            var in []reflect.Value
            for _, sysComp := range sys.components {
                arg := reflect.ValueOf(e).Elem().FieldByName(sysComp.Name()).Addr()
                in = append(in, arg)
            }
            sys.entitiesArgs = append(sys.entitiesArgs, in)
        }
    }
}

// Process processes systems in the order they are addedto the universe 
func (u *Universe) Process(){
    for _, sys := range u.systems {
        sys.process()
    }
}