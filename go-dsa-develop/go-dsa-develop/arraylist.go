package main

type ArrayList struct {
	items []interface{}
	size  int
	cap   int
}

func NewArrayList() *ArrayList {
	intialCap := 10
	return &ArrayList{
		items: make([]interface{}, intialCap),
		size:  0,
		cap:   intialCap,
	}
}

func (al *ArrayList) Add(item interface{}) {
	if al.size == al.cap {
		al.resize()
	}

	al.items[al.size] = item
	al.size++
}

func (al *ArrayList) resize() {
	newCap := al.cap * 2
	newItems := make([]interface{}, newCap)
	copy(newItems, al.items)
	al.items = newItems
	al.cap = newCap
}

func (al *ArrayList) GetAll() []interface{} {
	return al.items[0:al.size]
}
