package cobic

import (
	"fmt"
	"reflect"
)

var warnEmptyQuery = false

type EntityContext struct {
	EntityIndex     int
	ComponentArrays map[reflect.Type][]Component
}

func NewContext() EntityContext {
	return EntityContext{
		EntityIndex:     0,
		ComponentArrays: make(map[reflect.Type][]Component),
	}
}

func (ctx *EntityContext) AddEntity(components ...Component) {
	for _, c := range components {
		c.SetId(ctx.EntityIndex)
		cType := ptrType(c)
		ctx.ComponentArrays[cType] = append(ctx.ComponentArrays[cType], c)
	}

	ctx.EntityIndex++
}

func (ctx *EntityContext) QueryList(components ...Component) [][]Component {
	return ctx.QueryListTypes(convertCompTypes(components...)...)
}

func (ctx *EntityContext) QueryListTypes(compTypes ...reflect.Type) [][]Component {
	var result [][]Component

	for _, compType := range compTypes {
		key := compType
		storedComponents, exists := ctx.ComponentArrays[key]
		if !exists {
			if !warnEmptyQuery {
				emptyQueryWarning(ctx.ComponentArrays, key)
			}
			return make([][]Component, 0)
		}

		// Add the filtered components for the current type to the result
		result = append(result, storedComponents)
	}

	indexes := make(map[int]struct{})
	indexesStarted := false

	for _, compList := range result {
		next_indexes := make(map[int]struct{})
		for _, comp := range compList {
			next_indexes[comp.GetId()] = struct{}{}
		}

		if !indexesStarted {
			indexesStarted = true
			indexes = next_indexes
		} else {
			indexes = intersectIndexes(indexes, next_indexes)
		}
	}

	for i, compList := range result {
		result[i] = Filter(compList, func(c Component) bool {
			_, exists := indexes[c.GetId()]
			return exists
		})
	}

	return result
}

func (ctx *EntityContext) QueryMap(components ...Component) map[reflect.Type][]Component {
	return ctx.QueryMapTypes(convertCompTypes(components...)...)
}

func (ctx *EntityContext) QueryMapTypes(components ...reflect.Type) map[reflect.Type][]Component {
	resultList := ctx.QueryListTypes(components...)
	resultMap := make(map[reflect.Type][]Component)

	for _, comps := range resultList {
		compType := ptrType(comps[0])
		resultMap[compType] = comps
	}

	return resultMap
}

func (ctx *EntityContext) QueryGroups(components ...Component) [][]Component {
	return ctx.QueryGroupsTypes(convertCompTypes(components...)...)
}

func (ctx *EntityContext) QueryGroupsTypes(components ...reflect.Type) [][]Component {

	// Step 1: Get the result as a map of component types to lists
	lists := ctx.QueryListTypes(components...)
	entities := len(lists[0])
	result := make([][]Component, entities)

	for i := range entities {
		group := make([]Component, len(lists))

		for j, comp := range lists {
			group[j] = comp[i]
		}

		result[i] = group
	}

	return result
}

func convertAnyTypes(components ...any) []reflect.Type {
	var compTypes []reflect.Type

	for _, c := range components {
		switch v := c.(type) {
		case Component:
			compTypes = append(compTypes, reflect.TypeOf(v))
		case reflect.Type:
			compTypes = append(compTypes, v)
		default:
			panic(fmt.Sprintf("unexpected component type: %T", v))
		}
	}

	return compTypes
}

func convertCompTypes(components ...Component) []reflect.Type {
	var compTypes []reflect.Type

	for _, c := range components {
		compTypes = append(compTypes, ptrType(c))
	}

	return compTypes
}

func nonPtrType(input any) reflect.Type {
	val := reflect.ValueOf(input)

	if val.Kind() == reflect.Ptr {
		return val.Elem().Type()
	}

	return val.Type()
}

func ptrType(input any) reflect.Type {
	val := reflect.ValueOf(input)

	if val.Kind() == reflect.Ptr {
		return val.Type()
	}

	return reflect.PointerTo(val.Type())
}

func intersectIndexes(setA, setB map[int]struct{}) map[int]struct{} {
	result := make(map[int]struct{})

	for key := range setA {
		if _, exists := setB[key]; exists {
			result[key] = struct{}{}
		}
	}

	return result
}

func emptyQueryWarning(arrays map[reflect.Type][]Component, key reflect.Type) {
	warnEmptyQuery = true

	keys := make([]reflect.Type, len(arrays))

	i := 0
	for k := range arrays {
		keys[i] = k
		i++
	}

	fmt.Printf("component_keys: %s\nbad_key: %s\n", keys, key)
	fmt.Print("If it looks like the bad_key should match in component_keys, then you probably forgot to make all your system arguments pointers\n\n")
}
