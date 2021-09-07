package mediatel

import (
	"fmt"
)

type foo1 struct {
	field1 string
}
type bar struct {
	field1 string
}

func (b bar) Message() string {
	return fmt.Sprintf("message: %s", b.field1)
}

/*
Задание 4. Что выведет код? Найдите ошибку в коде?

Код выведет:
message: 1
message: 2
message: 3

Ошибка может заключаться в том, что если в срезе input при инициализации экземпляра foo1 поле field
структуры foo1 будет инициализировано пустой строкой, то при вызове метода Message() получим ошибку
"invalid memory address or nil pointer dereference".
Это так, поскольку при инициализации среза result создаются указатели типа *bar, которые по умолчанию nil.
Далее в цикле for есть условие на пустоту поля field1, которое будет истиной, и цикл перейдет к
следующей итерации, без выполнения изменения элементов среза result, в котором определяется адрес
памяти "экземпляра" структуры bar.

*/

func main4() {
	// инициализация input - среза из трех "экземпляров" foo1
	input := []foo1{{field1: "1"}, {field1: "2"}, {field1: "3"}}
	// инициализация среза указателей типа *bar, длины 3, емкости 3
	result := make([]*bar, len(input)) // [<nil> <nil> <nil>]

	for i, item := range input {
		// item.field1 == "" - false для данного примера
		if item.field1 == "" {
			continue
		}

		// изменение элементов среза result - установка указателя типа *bar на адрес памяти "экземпляра" структуры bar
		result[i] = &bar{field1: item.field1}
	}

	for _, r := range result {
		fmt.Println(r.Message())
	}
}
