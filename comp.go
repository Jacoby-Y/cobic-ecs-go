package cobic

type Component interface {
	SetId(id int)
	GetId() int
}

type BaseComponent struct {
	Id int
}

func (c *BaseComponent) SetId(id int) {
	c.Id = id
}

func (c BaseComponent) GetId() int {
	return c.Id
}

type Position struct {
	BaseComponent
	X, Y float32
}

type Velocity struct {
	BaseComponent
	X, Y float32
}
