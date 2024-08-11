package main

import "fmt"

type Spinner struct {
	frames []string
	// Position zero is the start, we're using the empty value here
	position int
}

func NewEllipsesSpinner() *Spinner {
	return &Spinner{
		frames:   []string{"Loading.", "Loading..", "Loading..."},
		position: 0,
	}
}

// limit is used to line up the position to the length of frames, while keeping position zero-indexed.
func (s *Spinner) limit() int {
	return len(s.frames) - 1
}

// next advances the spinner's position until it reaches the end of its frames, then we restart the animation.
func (s *Spinner) next() {
	pos := s.position
	lim := s.limit()

	// See where we are at the start of next
	if isDebug() {
		fmt.Printf("start pos=%v", pos)
	}

	// Keep going if we're under the limit
	if pos < lim {
		s.position = pos + 1
		return
	}
	// We're at our limit, restart the spinner
	if pos == lim {
		s.position = 0
		return
	}
	// We're over our limit, this shouldn't happen, panic
	if lim < pos {
		panic("limit exceeded somehow")
	}
}

func (s *Spinner) render() {
	str := s.frames[s.position]
	// Erase the line
	// TOOD: Add the cleared whitespace based on max length of frame... a bit out of scope though.
	fmt.Print("\r             ")
	// Clear the eraser whitespace then write our frame
	fmt.Printf("\r%s", str)
}

// Begin starts a new line
func (s *Spinner) Begin() {
	fmt.Print("\n")
}

// Animate Runs the spinner
func (s *Spinner) Animate() {
	s.render()
	s.next()
}

// Done Starts a new line
func (s *Spinner) Done() {
	fmt.Print("\r")
}
