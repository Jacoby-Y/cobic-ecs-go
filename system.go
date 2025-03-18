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

type BulkSystem struct {
	Name     string
	Func     reflect.Value
	ArgTypes []reflect.Type
}

var Systems []System
var BulkSystems []System

var queryCache = make(map[reflect.Value][][]Component)

func AddSystems(funcs ...any) {
	for _, f := range funcs {
		addSystem(f)
	}
}

func AddBulkSystems(funcs ...any) {
	for _, f := range funcs {
		addBulkSystem(f)
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

func addBulkSystem(f any) {
	fnType := reflect.TypeOf(f)

	// Ensure function is valid
	if fnType.Kind() != reflect.Func {
		panic("AddBulkSystem only accepts functions")
	}

	// Extract argument types
	argTypes := make([]reflect.Type, fnType.NumIn())
	for i := 0; i < fnType.NumIn(); i++ {
		argTypes[i] = fnType.In(i)
	}

	// Register system
	BulkSystems = append(Systems, System{
		Name:     fnType.String(),
		Func:     reflect.ValueOf(f),
		ArgTypes: argTypes,
	})
}

func RunSystems(ctx *EntityContext) {
	for _, sys := range Systems {

		qry, ok := queryCache[sys.Func]

		if !ok {
			qry = ctx.QueryGroupsTypes(sys.ArgTypes...)
			queryCache[sys.Func] = qry
		}

		for _, group := range qry {
			args := make([]reflect.Value, len(sys.ArgTypes))
			for j, argType := range group {
				args[j] = reflect.ValueOf(argType)
			}

			sys.Func.Call(args)
		}
	}
}

func RunBulkSystems(ctx *EntityContext) {
	for _, sys := range Systems {

		qry, ok := queryCache[sys.Func]

		if !ok {
			qry = ctx.QueryGroupsTypes(sys.ArgTypes...)
			queryCache[sys.Func] = qry
		}

		args := make([]reflect.Value, 1)

		args[0] = reflect.ValueOf(qry)

		sys.Func.Call(args)
	}
}

func ClearSystemCache() {
	queryCache = make(map[reflect.Value][][]Component)
}
