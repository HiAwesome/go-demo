package main

import (
	"fmt"
	"time"
)

/**
 * 结构体
 *
 * @author moqi
 * On 2021/1/3 14:00
 */
// 如果结构体成员的名字是以大写字母开头的，那么该成员就是可导出的；
// 这是 Go 语言导出规则决定的。一个结构体可能同时包含可导出和不可导出的成员。
type Employee struct {
	ID int
	// 通常，我们只是将相关的成员写到一起
	Name, Address string
	DoB           time.Time
	Position      string
	Salary        int
	ManagerID     int
}

// 一个命名为 S 的结构体将不能再包含 S 类型的成员：因为一个聚合的值不能包含它自身。（该限制同样适用于数组。）
// 但是 S 类型的结构体可以包含 *S 指针类型成员，这可以让我们创建递归的数据结构
type tree struct {
	value       int
	left, right *tree
}

// 结构体字面值
type Point struct {
	X, Y int
}

type address struct {
	hostname string
	port     int
}

type Circle1 struct {
	X, Y, Radius int
}

type Wheel1 struct {
	X, Y, Radius, Spokes int
}

type Point2 struct {
	X, Y int
}

type Circle2 struct {
	Center Point2
	Radius int
}

type Wheel2 struct {
	Circle Circle2
	Spokes int
}

// Go 语言有一个特性让我们只声明一个成员对应的数据类型而不声明成员的名字，这类成员就叫做匿名成员
// 匿名成员的数据类型必须是命名的类型或指向一个命名的类型的指针
type Circle3 struct {
	Point2
	Radius int
}

type Wheel3 struct {
	Circle3
	Spokes int
}

func main() {

	f1()

	f2()

	f3()

	f4()

	f5()

}

func f5() {
	var w1 Wheel1
	w1.X = 8
	w1.Y = 8
	w1.Radius = 5
	w1.Spokes = 20
	fmt.Println(w1)
	fmt.Println()

	var w2 Wheel2
	w2.Circle.Center.X = 8
	w2.Circle.Center.Y = 8
	w2.Circle.Radius = 5
	w2.Spokes = 20
	fmt.Println(w2)
	fmt.Println()

	var w3 Wheel3
	w3.X = 8
	w3.Y = 8
	w3.Radius = 5
	w3.Spokes = 20
	fmt.Println(w3)
	fmt.Println()

	w4 := Wheel3{Circle3{Point2{8, 8}, 5}, 20}
	w5 := Wheel3{
		Circle3: Circle3{
			Point2: Point2{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20, // Note: trailing comma necessary here(and at Radius)
	}
	fmt.Println(w4 == w5)
	fmt.Printf("%#v\n", w4)
	w4.X = 42
	fmt.Printf("%#v\n", w4)
	fmt.Println()
}

func f4() {
	p, q, z := Point{1, 2}, Point{2, 1}, Point{2, 1}
	fmt.Println(p.X == q.X && p.Y == q.Y)
	fmt.Println(p == q)
	fmt.Println(q == z)
	fmt.Println()

	hits := make(map[address]int)
	hits[address{"golang.ort", 443}]++
	fmt.Println(hits)
	fmt.Println()
}

func f3() {
	p := Point{1, 2}
	fmt.Println(p)
	p2 := Point{Y: 100, X: 200}
	fmt.Println(p2)
	fmt.Println()

	fmt.Println(Scale(Point{10, 20}, 5))
	fmt.Println()

	fmt.Println(Bonus(&Employee{Salary: 88}, 1000))
	fmt.Println()

	e1 := Employee{Salary: 999}
	AwardAnnualRaise(&e1)
	fmt.Println(e1)
	fmt.Println()
}

// 结构体可以作为函数的参数值和返回值
func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

// 如果考虑效率的话，较大的结构体通常会用指针的方式传入和返回
func Bonus(e *Employee, percent int) int {
	return e.Salary * percent / 100
}

// 如果要在函数内部修改结构体成员的话，用指针传入是必须的
// 因为在 Go 语言中，所有的函数参数都是值拷贝传入的
// 函数参数将不再是函数调用时的原始变量
func AwardAnnualRaise(e *Employee) {
	e.Salary = e.Salary * 105 / 100
}

func f2() {
	array1 := []int{1, 2, 9, 3, 12, 4, 7, 5}
	fmt.Println(array1)
	Sort(array1)
	fmt.Println(array1)
	fmt.Println()
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

// 根据给定的员工 ID 返回对应的员工信息结构体的指针
func EmployeeByID(id int) *Employee {
	return &Employee{ID: id}
}

func f1() {
	var dilbert Employee
	// 直接对变量赋值
	dilbert.Salary -= 5000
	// 对成员取地址，通过指针访问
	position := &dilbert.Position
	*position = "Senior " + *position
	// 点操作也可以和指向结构体点指针一起工作
	var employeeOfTheMonth = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"
	fmt.Println(dilbert)
	fmt.Println()

	fmt.Println(EmployeeByID(dilbert.ManagerID).Position)
	id := dilbert.ID
	EmployeeByID(id).Salary = 0
	fmt.Println(dilbert)
	fmt.Println()
}
