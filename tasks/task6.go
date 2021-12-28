package tasks

import "fmt"

/*
	Task 6. Задачи из курса http://golang-book.ru/
	7.1 - Тема 7, задача 1.
*/

// 7.1 Функция sum принимает срез чисел и складывает их вместе. Как бы выглядела сигнатура этой функции?
func half(val int) (float64, bool) {
	return float64(val) / 2, val%2 == 0
}

// 7.2 Напишите функцию, которая принимает число, делит его пополам и возвращает true в случае, если исходное число чётное,
//     и false, если нечетное. Например, half(1) должна вернуть (0, false), в то время как half(2) вернет (1, true).
func sum(slice []int) int {
	summa := 0
	for _, val := range slice {
		summa += val
	}
	return summa
}

// 7.3 Напишите функцию с переменным числом параметров, которая находит наибольшее число в списке.
func maxInList(args ...int) (int, bool) {
	if len(args) == 0 {
		return 0, false
	}

	max := args[0]
	for _, val := range args {
		if max < val {
			max = val
		}
	}
	return max, true
}

// 7.4 Используя в качестве примера функцию makeEvenGenerator напишите makeOddGenerator, генерирующую нечётные числа.
func makeOddGenerator() func() int {
	i := 1
	return func() (ret int) {
		ret = i
		i += 2
		return
	}
}

// 7.5 Последовательность чисел Фибоначчи определяется как fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2).
//     Напишите рекурсивную функцию, находящую fib(n).
// 0 1 1 2 3 5 8 13 21 34
// 0 1 2 3 4 5 6 7  8   9
func fib(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

// 8.4 Какое будет значение у переменной x после выполнения программы:
func square(x *float64) {
	*x = *x * *x
}

// 8.5 Напишите программу, которая меняет местами два числа (x := 1; y := 2; swap(&x, &y) должно дать x=2 и y=1).
func swap(a, b *int) {
	*a, *b = *b, *a
}

// 9.3 Добавьте новый метод perimeter в интерфейс Shape, который будет вычислять периметр фигуры.
// Имплементируйте этот метод для Circle и Rectangle.

type Circle struct {
	r, x, y float32
}

func (c *Circle) perimeter() float32 {
	return 2.0 * 3.14 * c.r
}

func (c *Circle) area() float32 {
	return 3.14 * c.r * c.r
}

type Rectangle struct {
	x1, x2, y1, y2 float32
}

func (r *Rectangle) perimeter() float32 {
	return (r.x2 - r.x1 + r.y2 - r.y1) * 2
}

func (r *Rectangle) area() float32 {
	return (r.x2 - r.x1) * (r.y2 - r.y1)
}

type Shape interface {
	area() float32
	perimeter() float32
}

type MultiShape struct {
	shapes []Shape // интерфейсы можно использовать как тип
}

// подсчет площади среза фигур - метод MultiShape
func (m *MultiShape) area() float32 {
	var area float32
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

// подсчет периметра списка параметров типа интерфейс Shape - функция
// 0 или более интерфейсов в качестве параметра
func totalPerimeter(shapes ...Shape) float32 {
	var totalPerimeterValue float32
	for _, shape := range shapes {
		totalPerimeterValue += shape.perimeter()
	}
	return totalPerimeterValue
}

func Task6() {
	// 7.1
	slice := []int{1, 2, 3}
	fmt.Println(sum(slice))

	// 7.2
	fmt.Println(half(1))
	fmt.Println(half(2))

	// 7.3
	fmt.Println(maxInList(1, 12, 5, 6))

	// 7.4
	theOddFunc := makeOddGenerator()
	fmt.Println(theOddFunc())
	fmt.Println(theOddFunc())
	fmt.Println(theOddFunc())

	// 7.5
	n := 9
	fmt.Printf("fib(%d) = %d\n", n, fib(n))

	// 8.4
	x := 1.5
	square(&x)
	fmt.Println(x)

	// 8.5
	y := 2
	z := 3
	swap(&y, &z)
	fmt.Println(y, z)

	// 9.3
	var r Rectangle = Rectangle{x1: 0, y1: 0, x2: 10, y2: 10}
	c := Circle{r: 5, x: 0, y: 0}
	var shape Shape = &c
	fmt.Println("Circle area:", shape.area())
	fmt.Println("Circle perimeter:", shape.perimeter())
	shape = &r
	fmt.Println("Rectangle area:", shape.area())
	fmt.Println("Rectangle perimeter:", shape.perimeter())

	shapes := []Shape{&r, &c}
	manyShapes := MultiShape{
		shapes: shapes,
	}
	fmt.Println("Total area is", manyShapes.area())
}
