package ecs

import "reflect"

type Universe struct {
    systems []interface{}
    entitiesByType map[reflect.Type][]interface{}
}

func NewUniverse() *Universe{
    u := Universe{}
    u.entitiesByType = make(map[reflect.Type][]interface{})
    return &u
}

func (u *Universe) AddSystem(sys interface{}){
    u.systems = append(u.systems, sys)
}

func (u *Universe) AddEntity(e interface{}){
    eType := reflect.TypeOf(e).Elem()
    eLen := eType.NumField()
    for i:=0; i < eLen; i++ {
        fieldType := eType.Field(i).Type
        entitiesForType := u.entitiesByType[fieldType]
        if entitiesForType == nil {
            entitiesForType = make([]interface{}, 0)
        }
        u.entitiesByType[fieldType] = append(entitiesForType, e)
    } 
}

func (u *Universe) Process(){
    for _, sys := range(u.systems) {
        fnType := reflect.TypeOf(sys)
        fnLen := fnType.NumIn()
        candidates := make([][]interface{}, fnLen, fnLen)
        for i:=0; i < fnLen; i++ {
            fieldType := fnType.In(i).Elem()
            candidates[i] = u.entitiesByType[fieldType]
        }

        entities := make(map[interface{}]int)
        for _, c := range(candidates) {
            for _, e := range(c) {
                entities[e]++
            }
        }

        var matches []interface{}
        for e, count := range(entities) {
            if count == fnLen {
                matches = append(matches, e)
            } 
        }

        sysFunc := reflect.ValueOf(sys)

        for _, match := range(matches) {
            matchValue := reflect.ValueOf(match).Elem()
            var in []reflect.Value
            for i:=0; i < fnLen; i++ {
                argName := fnType.In(i).Elem().Name()
                arg := matchValue.FieldByName(argName).Addr()
                in = append(in, arg)
            }
            sysFunc.Call(in)
        }

    }
}