package main

const pi = 3.14159

type Shape interface {
	Area() float64
	Type() string
}

type Rectangle struct {
	height float64
	width  float64
}

func NewRectangle(height float64, width float64, name string) *Rectangle {
	return &Rectangle{height: height, width: width}
}

func (r Rectangle) Type() string {
	return "прямоугольник"
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Type() string {
	return "окружность"
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * pi
}

type Circle struct {
	radius float64
}

func NewCircle(radius float64, name string) *Circle {
	return &Circle{radius: radius}
}

func AreaCalculator(figures []Shape) (str string, area float64) {
	if figures == nil {
		return
	}
	for _, value := range figures {
		str = str + value.Type() + "-"
		area = area + value.Area()
	}
	str = str[0 : len(str)-1]
	return
}

func main() {
}
