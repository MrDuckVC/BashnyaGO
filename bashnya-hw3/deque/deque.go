package deque

import (
	"errors"
	"fmt"
)

type dNode struct {
	value int
	prev  *dNode
	next  *dNode
}

type Deque struct {
	head *dNode
	tail *dNode
	size int
}

func NewDeque() *Deque {
	return &Deque{nil, nil, 0}
}

func (d *Deque) PushFront(value int) {
	node := &dNode{value: value, prev: nil, next: d.head}
	if d.IsEmpty() {
		d.head = node
		d.tail = node
	} else {
		d.head.prev = node
		d.head = node
	}
	d.size++
}
func (d *Deque) PushBack(value int) {
	node := &dNode{value: value, prev: d.tail, next: nil}
	if d.IsEmpty() {
		d.head = node
		d.tail = node
	} else {
		d.tail.next = node
		d.tail = node
	}
	d.size++
}
func (d *Deque) PopFront() (int, error) {
	if d.IsEmpty() {
		return 0, errors.New("Deque is empty")
	}
	value := d.head.value
	d.head = d.head.next
	if d.head != nil {
		d.head.prev = nil
	} else {
		d.tail = nil
	}
	d.size--
	return value, nil
}
func (d *Deque) PopBack() (int, error) {
	if d.IsEmpty() {
		return 0, errors.New("Deque is empty")
	}
	value := d.tail.value
	d.tail = d.tail.prev
	if d.tail != nil {
		d.tail.next = nil
	} else {
		d.head = nil
	}
	d.size--
	return value, nil
}
func (d *Deque) IsEmpty() bool {
	return d.size == 0
}
func (d *Deque) Size() int {
	return d.size
}
func (d *Deque) Clear() {
	d.head = nil
	d.tail = nil
	d.size = 0
}

func main_deque() {
	fmt.Println("\n--- Продвинутый: Deque ---")
	deque := NewDeque()
	deque.PushBack(10) // [10]
	deque.PushFront(5) // [5, 10]
	deque.PushBack(20) // [5, 10, 20]

	fmt.Println("Размер:", deque.Size()) // 3

	item, _ := deque.PopFront()
	fmt.Println("Извлекли спереди:", item) // 5

	item, _ = deque.PopBack()
	fmt.Println("Извлекли сзади:", item) // 20

	fmt.Println("Размер:", deque.Size()) // 1
	deque.Clear()
	fmt.Println("Пустой?:", deque.IsEmpty()) // true
}
