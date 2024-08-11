package main

import (
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
	fmt.Printf("%+v\n", m)
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

	t := NewTodo()
	args := os.Args[1:]
	entry := NewEntry(args[1:])

	switch cmd := Command(args[0]); cmd {
	case List:
		t.List()
		return
	case Add:
		t.Add(entry)
		render(t.List())
		return
	case Complete:
		t.Done(entry)
		render(t.List())
		return
	case Remove:
		t.Remove(entry)
		render(t.List())
		return
	case Clean:
		// TODO: Removes all existing disk state
	case TestSpinner:
		runTestSpinner()
		return
	default:
		fmt.Printf("Unknown command, command=%s \n", cmd)
		return
	}
}
