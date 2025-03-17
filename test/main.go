package main

import (
	"fmt"
	"reflect"
)

var resources = make(map[reflect.Type]any)

type Score struct {
	Value int
}

func SetResource(res any) {
	t := reflect.TypeOf(res)
	resources[t] = res
}

func GetResource(res any) any {
	return GetResourceType(reflect.TypeOf(res))
}

func GetResourceType(resType reflect.Type) any {
	return resources[resType]
}

func main() {
	SetResource(&Score{
		Value: 100,
	})

	score, ok := GetResource(&Score{}).(*Score)

	if !ok {
		fmt.Println("No Score resource")
		return
	}

	fmt.Printf("Score value: %d\n", score.Value)
	score.Value += 50
}
