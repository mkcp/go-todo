package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

var DEBUG = "false"

func isDebug() bool {
	return DEBUG == "true"
}

type Command string

var (
	List        = Command("list")
	Add         = Command("add")
	Complete    = Command("done")
	Remove      = Command("remove")
	Clean       = Command("clean")
	TestSpinner = Command("spinner")
)

// Displays the current state of the Todo map to users
// TODO: Use tab-something to render this as a table with headers
// e.g.
//
//	State  | Entry
//	Done   | Create a TODO app
func render(m map[Entry]State) {
	bs, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", string(bs))
}

// runTestSpinner is a helper command for main that wraps the TestSpinner command logic
func runTestSpinner() {
	spinner := NewEllipsesSpinner()

	// Simulate a ten-second runtime
	runtime := 10
	for i := 0; i < runtime; i++ {
		spinner.Animate()
		time.Sleep(100 * time.Millisecond)
	}

	spinner.Done()
	fmt.Println("All Done!")
}

func main() {
	DEBUG = os.Getenv("DEBUG")

	t, err := NewTodo()
	if err != nil {
		panic(err)
	}
	args := os.Args[1:]
	entry := NewEntry(args[1:])

	switch cmd := Command(args[0]); cmd {
	case List:
		m, err := t.List()
		if err != nil {
			panic(err)
		}
		render(m)
		return
	case Add:
		err := t.Add(entry)
		if err != nil {
			panic(err)
		}
		m, err2 := t.List()
		if err2 != nil {
			panic(err)
		}
		render(m)
		return
	case Complete:
		err := t.Done(entry)
		if err != nil {
			panic(err)
		}
		m, err2 := t.List()
		if err2 != nil {
			panic(err)
		}
		render(m)
		return
	case Remove:
		err := t.Remove(entry)
		if err != nil {
			panic(err)
		}
		m, err2 := t.List()
		if err2 != nil {
			panic(err)
		}
		render(m)
		return
	case Clean:
		err := t.Cleanup()
		if err != nil {
			panic(err)
		}
		return
	case TestSpinner:
		runTestSpinner()
		return
	default:
		fmt.Printf("Unknown command, command=%s \n", cmd)
		return
	}
}
