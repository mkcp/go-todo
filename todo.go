package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const LOCATION = "./resources/todos.json"

type Entry string

func NewEntry(args []string) Entry {
	return Entry(strings.Join(args, " "))
}

type State bool

var (
	Pending = State(true)
	Done    = State(false)
)

type Todo struct {
	m map[Entry]State
}

func NewTodo() (*Todo, error) {
	t := &Todo{
		m: make(map[Entry]State),
	}
	// Ensure file exists
	if _, err := os.Stat(LOCATION); os.IsNotExist(err) {
		_, err2 := os.Create(LOCATION)
		if err2 != nil {
			return nil, err2
		}
	}
	return t, nil
}

func (t *Todo) List() (map[Entry]State, error) {
	err := t.read()
	if err != nil {
		return nil, err
	}
	return t.m, nil
}

func (t *Todo) update(e Entry, state State) error {
	err := t.read()
	if err != nil {
		return err
	}
	t.m[e] = state
	err2 := t.write()
	if err2 != nil {
		return err2
	}
	return nil
}

func (t *Todo) Add(e Entry) error {
	return t.update(e, Pending)
}

func (t *Todo) Done(e Entry) error {
	return t.update(e, Done)
}

func (t *Todo) Remove(e Entry) error {
	// Load state
	err := t.read()
	if err != nil {
		return err
	}

	// Check if the key exists, and if so delete it
	if _, ok := t.m[e]; !ok {
		return fmt.Errorf("entry not found, entry=%s", e)
	}
	delete(t.m, e)

	// Commit the edit
	err2 := t.write()
	if err2 != nil {
		return err2
	}
	return nil
}

func (t *Todo) Cleanup() error {
	return os.Remove(LOCATION)
}

func deserialize(payload []byte) (map[Entry]State, error) {
	m := make(map[Entry]State)
	err := json.Unmarshal(payload, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func serialize(m map[Entry]State) ([]byte, error) {
	payload, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func (t *Todo) read() error {
	bs := make([]byte, 0)
	m := make(map[Entry]State)

	f, err := os.Open(LOCATION)
	if err != nil {
		return err
	}
	defer f.Close()
	// Load contents of file
	_, err2 := f.Read(bs)
	if err2 != nil {
		return err2
	}

	// parse json into map[Entry]State
	err3 := json.Unmarshal(bs, &m)
	if err3 != nil {
		// return err3
		return nil // FIXME: Can we proceed here from empty file
	}
	t.m = m
	return nil
}

func (t *Todo) write() error {
	bs, err := serialize(t.m)
	if err != nil {
		return err
	}
	f, err := os.Create(LOCATION)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(bs)
	if err != nil {
		return err
	}
	// Return nil
	return nil
}

// TODO: Make these async out to goroutines, set a spinner while waiting, and return.
