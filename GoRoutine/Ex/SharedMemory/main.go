package main

import "fmt"

const (
	set = iota
	get
	remove
	count
)

type SharedMap struct {
	m map[string]interface{}
	c chan command
}

type command struct {
	key		string
	value	interface{}
	action	int
	result	chan<-	interface{}
}

func (sm SharedMap) Set(k string, v interface{}) {
	sm.c <- command{action: set, key: k, value: v}
}

func (sm SharedMap) Get(k string) (interface{}, bool) {
	callback := make(chan interface{})
	sm.c <- command{action: get, key: k, result: callback}
	result := (<-callback).([2]interface{})
	return result[0], result[1].(bool)
}

func (sm SharedMap) Remove(k string) {
	sm.c <- command{action: remove, key: k}
}

func (sm SharedMap) Count() int {
	callback := make(chan interface{})
	sm.c <- command{action: count, result: callback}
	return (<-callback).(int)
}

func (sm SharedMap) run() {
	for cmd := range sm.c {
		switch cmd.action {
		case set:
			sm.m[cmd.key] = cmd.value
		case get:
			v, ok := sm.m[cmd.key]
			cmd.result <- [2]interface{}{v, ok}
		case remove:
			delete(sm.m, cmd.key)
		case count:
			cmd.result <- len(sm.m)
		}		
	}
}

func NewMap() SharedMap {
	sm := SharedMap{
		m: make(map[string]interface{}),
		c: make(chan command),
	}
	go sm.run()
	return sm
}

func main() {
	m := NewMap()

	m.Set("foo", "bar")
	fmt.Println("Count: ", m.Count())
	
	t, ok := m.Get("foo")

	if ok {
		bar := t.(string)
		fmt.Println("bar: ", bar)
	}

	m.Remove("foo")

	// Count
	fmt.Println("Count: ", m.Count())
}