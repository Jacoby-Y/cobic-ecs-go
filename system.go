package cobic

import (
	"reflect"
)

type NewSys struct{}

type System struct {
	Name     string
	Func     reflect.Value
	ArgTypes []reflect.Type
}

var Systems []System

func AddSystems(funcs ...any) {
	for _, f := range funcs {
		addSystem(f)
	}
}

func addSystem(f any) {
	fnType := reflect.TypeOf(f)

	// Ensure function is valid
	if fnType.Kind() != reflect.Func {
		panic("AddSystem only accepts functions")
	}

	// Extract argument types
	argTypes := make([]reflect.Type, fnType.NumIn())
	for i := 0; i < fnType.NumIn(); i++ {
		argTypes[i] = fnType.In(i)
	}

	// Register system
	Systems = append(Systems, System{
		Name:     fnType.String(),
		Func:     reflect.ValueOf(f),
		ArgTypes: argTypes,
	})
}

func RunSystems(ctx *EntityContext) {
	for _, sys := range Systems {

		qry := ctx.QueryGroupsTypes(sys.ArgTypes...)

		for _, group := range qry {
			args := make([]reflect.Value, len(sys.ArgTypes))
			for j, argType := range group {
				args[j] = reflect.ValueOf(argType)
			}

			sys.Func.Call(args)
		}
	}
}
