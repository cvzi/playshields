package lru

import (
	"container/list"
	"fmt"
	"sync"
)

type Cache struct {
	max   int
	index map[interface{}]*list.Element
	*list.List
	sync.RWMutex
}

type ListData struct {
	key   interface{}
	value interface{}
}

func New(max int) *Cache {
	return &Cache{
		max:   max,
		index: make(map[interface{}]*list.Element, max+1),
		List:  list.New(),
	}
}

func (c *Cache) Get(key interface{}) (value interface{}, ok bool) {
	c.RLock()
	defer c.RUnlock()
	listElement, ok := c.index[key]
	if ok {
		c.MoveToFront(listElement)
		return listElement.Value.(*ListData).value, true
	}
	return nil, false
}

func (c *Cache) Set(key interface{}, value interface{}) {
	c.Lock()
	defer c.Unlock()

	listElement, ok := c.index[key]
	if ok {
		c.MoveToFront(listElement)
		listElement.Value.(*ListData).value = value
		return
	}
	listElement = c.PushFront(&ListData{key: key, value: value})
	c.index[key] = listElement

	if c.max != 0 && c.Len() > c.max {
		lastElement := c.Back()
		lastKey := lastElement.Value.(*ListData).key
		c.Remove(lastElement)
		delete(c.index, lastKey)
	}
}

func main() {
	fmt.Println("LRU:")

	x := New(3)

	x.Set("test1", 0)
	x.Set("test1", 1)
	x.Set("test2", 2)
	x.Set("test3", 3)
	x.Set("test4", 4)
	x.Set("test5", 5)

	fmt.Println(x.Get("test1"))
	fmt.Println(x.Get("test2"))
	fmt.Println(x.Get("test3"))
	fmt.Println(x.Get("test3"))
	x.Set("test6", 6)

	fmt.Printf("List:")
	for temp := x.List.Front(); temp != nil; temp = temp.Next() {
		fmt.Println(temp.Value)
	}
}
