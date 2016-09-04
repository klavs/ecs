package ecs

import "reflect"

type system struct {
    fn reflect.Value
    components []reflect.Type
    entities []interface{}
}

func (s *system) process(){
    for _, entity := range s.entities {
        var in []reflect.Value
        for _, sysComp := range s.components {
            arg := reflect.ValueOf(entity).Elem().FieldByName(sysComp.Name()).Addr()
            in = append(in, arg)
        }
        s.fn.Call(in)
    }
}

type Universe struct {
    systems []*system
    entitiesByType map[reflect.Type][]interface{}
}

func NewUniverse() *Universe{
    u := Universe{}
    return &u
}

func (u *Universe) AddSystem(sys interface{}){
    fn := reflect.ValueOf(sys)
    fnLen := fn.Type().NumIn()
    components := make([]reflect.Type, fnLen, fnLen)
    for i:=0; i < fnLen; i++ {
        components[i] = fn.Type().In(i).Elem()
    }
    u.systems = append(u.systems, &system{fn, components, nil})
}

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
            sys.entities = append(sys.entities, e)
        }
    }
}

func (u *Universe) Process(){
    for _, sys := range u.systems {
        sys.process()
    }
}