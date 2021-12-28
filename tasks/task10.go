package tasks

import "fmt"

/*
	Task 10. Инвертировать односвязный список
		1 --> 2 --> 3 --> 4 --> 5
		5 --> 4 --> 3 --> 2 --> 1
*/

type Elem struct {
	value int
	next  *Elem
}

func Reverse(e *Elem) *Elem {
	elem := e

	if elem == nil {
		return nil
	}

	nextElem := elem.next
	if nextElem == nil {
		return elem
	}
	elem.next = nil

	nextNextElem := nextElem.next
	if nextNextElem == nil {
		nextElem.next = elem
		elem.next = nil

		return nextElem
	}

	for {
		nextElem.next = elem // замена ссылки в обратную сторону

		elem = nextElem
		nextElem = nextNextElem

		if nextElem.next == nil {
			nextElem.next = elem
			break
		}

		nextNextElem = nextElem.next
	}

	return nextElem
}

func printList(elem *Elem) {
	for {
		if elem != nil {
			fmt.Print(elem.value, "-->")
		} else {
			break
		}

		elem = elem.next
	}

	fmt.Println()
}

func Task10() {
	elem := &Elem{
		value: 1,
	}
	head := elem

	for i := 2; i < 10; i++ {
		elem.next = &Elem{
			value: i,
		}

		elem = elem.next
	}

	printList(head)
	reverseHead := Reverse(head)
	printList(reverseHead)
}
