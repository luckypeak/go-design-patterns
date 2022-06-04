// Package visitor 将算法和数据结构分离的设计模式, 可以在不改变数据结构的前提下增加新的操作
package visitor

import "fmt"

type CarElement interface {
	Accept(visitor Visitor)
	GetName() string
}

type Engine struct {
	name string
}

func (e *Engine) GetName() string  {
	return e.name
}
func (e *Engine)Accept(visitor Visitor)  {
	visitor.Visit(e)
}

type Wheel struct {
	name string
}

func (w *Wheel)Accept(visitor Visitor){
	visitor.Visit(w)
}

func (w *Wheel)GetName() string  {
	return w.name
}

type Car struct {
	elements []CarElement
}


func NewCar() *Car{
	return &Car{
		elements: []CarElement{&Engine{name: "enging"}, &Wheel{name: "wheel"}},
	}
}

func (c *Car) Accept(visitor Visitor)  {
	for _, e := range c.elements{
		e.Accept(visitor)
	}
}
type Visitor interface {
	Visit(element CarElement)
}

type PrintVisitor struct {
}

func (ev *PrintVisitor) Visit(e CarElement) {
	fmt.Printf("visit %v print\n", e.GetName())
}

type DoVisitor struct {
}

func (ev *DoVisitor) Visit(e CarElement) {
	fmt.Printf("visit %v do\n", e.GetName())
}
