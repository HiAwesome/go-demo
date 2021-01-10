package geometry

import "math"

type Point struct {
	X, Y float64
}

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// 参数 p，叫做方法的接收器 receiver, 早期的面向对象语言留下的遗产将调用的一个方法称为 向一个对象发送消息
// 在 Go 语言中，我们并不会像其他语言那样用 this 或者 self 作为接收器，我们可以任意的选择接收器的名字
// 由于接收器的名字经常会被用到，所以保持其在方法间传递时的一致性和简短性是不错的主意。
// 这里的建议是可以使用其类型的第一个字母，比如这里是用了 Point 的首字母 p。
// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// 编译器报错: Type 'Point' has both field and method named 'X'
// func (p Point) X() {}

// Path 是一个命名的 slice 类型，而不是 Point 那样的 strut 类型
// 然而我们依然可以为它定义方法。在能够给任意类型定义方法这一点上，Go 和很多其他的面向对象的语言不太一样
// A Path is a journey connecting the points with straight lines
type Path []Point

// Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func PathDistance(path Path) float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
