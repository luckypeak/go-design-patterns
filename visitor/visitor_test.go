package visitor

import "testing"

func TestCar(t *testing.T) {
	car := NewCar()
	v1 := PrintVisitor{}
	car.Accept(&v1)
	v2 := DoVisitor{}
	car.Accept(&v2)
}