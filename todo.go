package main

import (
	"fmt"
	"strings"
)

type Entry string

func NewEntry(args []string) Entry {
	return Entry(strings.Join(args, " "))
}

type State bool

var (
	Pending = State(true)
	Done    = State(false)
)

// TODO: Add persistence
type Todo struct {
	m map[Entry]State
}

func NewTodo() *Todo {
	// TODO Check if exists on disk
	// TODO If it does, load
	// TODO If it doesn't, initialize
	return &Todo{
		m: make(map[Entry]State),
	}
}

func (t *Todo) List() map[Entry]State {
	return t.m
}

func (t *Todo) Add(e Entry) *Todo {
	t.m[e] = Pending
	return t
}

func (t *Todo) Done(e Entry) *Todo {
	t.m[e] = Done
	return t
}

func (t *Todo) Remove(e Entry) *Todo {
	// Error out and let the user know if we didn't find an entry
	if _, ok := t.m[e]; !ok {
		fmt.Printf("Entry not found, entry=%s", e)
		return t
	}
	delete(t.m, e)
	return t
}

// TODO: Read
// TODO: Write
// TODO: Make these async out to goroutines, set a spinner while waiting, and return.
