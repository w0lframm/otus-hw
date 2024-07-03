package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	front  *ListItem
	back   *ListItem
	length int
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.length
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	item := &ListItem{
		Value: v,
	}
	if l.length == 0 {
		l.front = item
		l.back = item
	} else {
		item.Next = l.front
		l.front.Prev = item
		l.front = item
	}
	l.length++
	return l.front
}

func (l *list) PushBack(v interface{}) *ListItem {
	old := l.back
	l.back = &ListItem{
		Value: v,
		Prev:  old,
	}
	l.length++
	if l.length == 1 {
		l.front = l.back
	} else {
		old.Next = l.back
	}
	return l.back
}

func (l *list) Remove(i *ListItem) {
	l.length--
	if i.Next == nil && i.Prev == nil {
		l.front = nil
		l.back = nil
		return
	}
	prev, next := i.Prev, i.Next
	if prev == nil {
		next.Prev = nil
		l.front = next
		return
	}
	if next == nil {
		prev.Next = nil
		l.back = prev
		return
	}
	prev.Next, next.Prev = next, prev
}

func (l *list) MoveToFront(i *ListItem) {
	prev, next := i.Prev, i.Next
	if prev == nil && next == nil {
		return
	}
	if prev == nil {
		return
	}
	if next == nil {
		l.back = prev
	}
	prev.Next = next
	i.Prev = nil
	i.Next = l.front
	l.front.Prev = i
	l.front = i
}
